package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stackrox/rox/generated/internalapi/sensor"
	"github.com/stackrox/rox/generated/storage"
	policiesTesting "github.com/stackrox/rox/pkg/defaults/policies/testing"
	"github.com/stackrox/rox/pkg/uuid"
	"github.com/stackrox/rox/sensor/admission-control/manager"
	managerTesting "github.com/stackrox/rox/sensor/admission-control/manager/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	admission "k8s.io/api/admission/v1beta1"
)

const (
	ExecIntoPodPolicyName       = "Kubernetes Actions: Exec into Pod"
	LatestTagPolicyName         = "Latest tag"
	admissionControllerEndpoint = "https://admission-control.stackrox:443"
)

type admissionReviewType string

const (
	// eventAdmissionReview is used to validate kube events
	eventAdmissionReview admissionReviewType = "events"
	// validateAdmissionReview is used to validate deployments, pods etc
	validateAdmissionReview admissionReviewType = "validate"
)

func (a admissionReviewType) String() string {
	return string(a)
}

func TestExecIntoPodNameEventPolicy(t *testing.T) {
	policy, err := policiesTesting.GetDefaultPolicy(t, ExecIntoPodPolicyName)
	require.NoError(t, err)

	mgr := managerTesting.NewTestManager(t,
		managerTesting.TestManagerOptions{Policy: policy},
	)

	err = mgr.Start()
	require.NoError(t, err)
	defer mgr.Stop()

	managerTesting.ProcessDeploymentEvent(t, mgr, &storage.Deployment{
		Id:        "f3237faf-8350-4c39-b045-ff4c493ddb71",
		Name:      "sensor",
		Type:      "Deployment",
		Namespace: "stackrox",
	})
	managerTesting.ProcessPodEvent(t, mgr, &storage.Pod{
		Id:           "64a1d6ee-2425-5f19-990e-a2d8b18c1e4c",
		Name:         "sensor-74f6965874-qckz6",
		DeploymentId: "f3237faf-8350-4c39-b045-ff4c493ddb71",
		Namespace:    "stackrox",
	})

	r := serviceTestRun{
		reviewType:        eventAdmissionReview,
		mgr:               mgr,
		reviewRequestPath: "testdata/review_requests/pod_exec_event_review.json",
		assertionFunc: func(t *testing.T, review admission.AdmissionReview, alerts []*storage.Alert) {
			require.NotNil(t, alerts)
			require.Len(t, alerts, 1)
			assert.Equal(t, "Kubernetes Actions: Exec into Pod", alerts[0].GetPolicy().GetName())

			violations := alerts[0].GetViolations()
			require.Len(t, violations, 1)
			assert.Equal(t, "Kubernetes API received exec '/bin/sh' request into pod 'sensor-74f6965874-qckz6' container 'sensor'", violations[0].GetMessage())
			assert.Equal(t, "sensor", alerts[0].GetDeployment().GetName())

			assert.True(t, review.Response.Allowed)
		},
		t: t,
	}

	r.execute()
}

func TestLatestTagPolicyAdmissionReview(t *testing.T) {
	policy, err := policiesTesting.GetDefaultPolicy(t, LatestTagPolicyName)
	require.NoError(t, err)

	policy.EnforcementActions = []storage.EnforcementAction{
		storage.EnforcementAction_SCALE_TO_ZERO_ENFORCEMENT,
	}

	mgr := managerTesting.NewTestManager(t, managerTesting.TestManagerOptions{
		AdmissionControllerSettings: &sensor.AdmissionControlSettings{
			ClusterId: uuid.NewDummy().String(),
			ClusterConfig: &storage.DynamicClusterConfig{
				AdmissionControllerConfig: &storage.AdmissionControllerConfig{
					EnforceOnUpdates: true,
					Enabled:          true,
				},
			},
		},
		Policy: policy,
		ImageServiceResponse: &sensor.GetImageResponse{
			Image: &storage.Image{
				Id: "sha256:e66b2e83961df8f87a4a20c0365b1404d60cdd58798f4db5763332fe0ac235ea",
				Name: &storage.ImageName{
					Registry: "docker.io",
					Remote:   "library/nginx",
					Tag:      "latest",
					FullName: "docker.io/library/nginx:latest",
				},
			},
		},
	})

	err = mgr.Start()
	require.NoError(t, err)
	defer mgr.Stop()

	r := serviceTestRun{
		mgr:               mgr,
		reviewType:        validateAdmissionReview,
		reviewRequestPath: "testdata/review_requests/latest_tag_admission_review.json",
		assertionFunc: func(t *testing.T, review admission.AdmissionReview, alerts []*storage.Alert) {
			const latestTagErrMessage = "Container 'nginx' has image with tag 'latest'"

			require.NotNil(t, alerts)
			require.Len(t, alerts, 1)
			assert.Equal(t, LatestTagPolicyName, alerts[0].GetPolicy().GetName())
			require.Len(t, alerts[0].GetViolations(), 1)
			assert.Equal(t, latestTagErrMessage, alerts[0].GetViolations()[0].Message)

			assert.Contains(t, review.Response.Result.Message, latestTagErrMessage)
			assert.False(t, review.Response.Allowed)
		},
		t: t,
	}

	r.execute()
}

type serviceTestRun struct {
	mgr               manager.Manager
	reviewType        admissionReviewType
	reviewRequestPath string
	assertionFunc     func(t *testing.T, review admission.AdmissionReview, alerts []*storage.Alert)
	t                 *testing.T
}

func (r serviceTestRun) execute() {
	require.NotNil(r.t, r.mgr)
	require.True(r.t, r.mgr.IsReady())

	s := service{
		mgr: r.mgr,
	}

	requestBody, err := os.ReadFile(r.reviewRequestPath)
	require.NoError(r.t, err)

	req := httptest.NewRequest(http.MethodPost, path.Join(admissionControllerEndpoint, r.reviewType.String()), bytes.NewBuffer(requestBody))
	resp := httptest.NewRecorder()

	switch r.reviewType {
	case eventAdmissionReview:
		s.handleK8sEvents(resp, req)
	case validateAdmissionReview:
		s.handleValidate(resp, req)
	default:
		require.Failf(r.t, "Review type %q is not supported", r.reviewType.String())
	}

	require.NotNil(r.t, resp)
	assert.Equal(r.t, http.StatusOK, resp.Code)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(r.t, err)
	review := admission.AdmissionReview{}
	err = json.Unmarshal(respBody, &review)
	require.NoError(r.t, err)

	select {
	case <-time.After(1 * time.Second):
		assert.Fail(r.t, "Did not receive any alerts before timeout expired, but expected some")
	case alerts := <-r.mgr.Alerts():
		r.assertionFunc(r.t, review, alerts)
	}
}