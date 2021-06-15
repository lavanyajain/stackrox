package translation

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"github.com/stackrox/rox/generated/storage"
	securedcluster "github.com/stackrox/rox/operator/api/securedcluster/v1alpha1"
	"github.com/stackrox/rox/operator/pkg/values/translation"
	"github.com/stackrox/rox/pkg/buildinfo"
	"helm.sh/helm/v3/pkg/chartutil"
	k8sErrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	supportedOperandNamespace     = "stackrox"
	supportedResourceName         = "stackrox-secured-cluster-services"
	sensorTLSSecretName           = "sensor-tls"
	admissionControlTLSSecretName = "admission-control-tls"
	collectorTLSSecretName        = "collector-tls"
)

// NewTranslator creates a translator
func NewTranslator(config *rest.Config) Translator {
	// TODO(ROX-7251): make sure that the client we create here is kosher
	return Translator{clientSet: kubernetes.NewForConfigOrDie(config)}
}

// Translator translates and enriches helm values
type Translator struct {
	clientSet kubernetes.Interface
}

// Translate translates and enriches helm values
func (t Translator) Translate(ctx context.Context, u *unstructured.Unstructured) (chartutil.Values, error) {
	sc := securedcluster.SecuredCluster{}
	err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &sc)
	if err != nil {
		return nil, err
	}

	if err := validate(sc); err != nil {
		return nil, err
	}

	return t.translate(ctx, sc)
}

// Translate translates a SecuredCluster CR into helm values.
func (t Translator) translate(ctx context.Context, sc securedcluster.SecuredCluster) (chartutil.Values, error) {
	v := translation.NewValuesBuilder()
	if !buildinfo.ReleaseBuild {
		v.AddAllFrom(t.getHelmDevelopmentDefaults(sc))
	}

	// TODO(ROX-7125): prevent/allow cluster name change?
	v.SetStringValue("clusterName", sc.Spec.ClusterName)

	v.SetString("centralEndpoint", sc.Spec.CentralEndpoint)

	v.AddAllFrom(t.getTLSValues(ctx, sc))

	v.AddAllFrom(translation.GetImagePullSecrets(sc.Spec.ImagePullSecrets))

	// TODO(ROX-7178): support explicit env.openshift and env.istio setting
	// TODO(ROX-7150): support setting/overriding images

	customize := translation.NewValuesBuilder()

	if sc.Spec.Sensor != nil {
		v.AddChild("sensor", getSensorValues(ctx, t.clientSet, sc.Namespace, sc.Spec.Sensor))
		customize.AddChild("sensor", translation.GetCustomize(sc.Spec.Sensor.Customize))
	}

	if sc.Spec.AdmissionControl != nil {
		v.AddChild("admissionControl", t.getAdmissionControlValues(ctx, sc.Namespace, sc.Spec.AdmissionControl))
		customize.AddChild("admission-control", translation.GetCustomize(sc.Spec.AdmissionControl.Customize))
	}

	if sc.Spec.Collector != nil {
		v.AddChild("collector", t.getCollectorValues(ctx, sc.Namespace, sc.Spec.Collector))
		customize.AddChild("collector", translation.GetCustomize(sc.Spec.Collector.Customize))
	}

	customize.AddAllFrom(translation.GetCustomize(sc.Spec.Customize))
	v.AddChild("customize", &customize)

	return v.Build()
}

func validate(sc securedcluster.SecuredCluster) error {
	if sc.Name != supportedResourceName {
		return errors.Errorf("invalid metadata.name: currently only %q is supported", supportedResourceName)
	}
	if sc.Namespace != supportedOperandNamespace {
		return errors.Errorf("invalid metadata.namespace: currently only %q is supported", supportedOperandNamespace)
	}
	return nil
}

// getTLSValues reads TLS configuration and looks up CA certificate from secrets.
func (t Translator) getTLSValues(ctx context.Context, sc securedcluster.SecuredCluster) *translation.ValuesBuilder {
	v := translation.NewValuesBuilder()
	if err := t.checkRequiredTLSSecrets(ctx, sc); err != nil {
		return v.SetError(err)
	}

	v.SetBoolValue("createSecrets", false)
	v.AddAllFrom(translation.GetTLSConfig(sc.Spec.TLS))
	sensorSecret, err := t.clientSet.CoreV1().Secrets(sc.Namespace).Get(ctx, sensorTLSSecretName, metav1.GetOptions{})
	if err != nil {
		return v.SetError(errors.Wrapf(err, "failed reading %q secret", sensorTLSSecretName))
	}

	centralCA, ok := sensorSecret.Data["ca.pem"]
	if !ok {
		return v.SetError(errors.Errorf("could not find centrals ca certificate 'ca.pem' in secret/%s", sensorTLSSecretName))
	}
	v.SetStringMap("ca", map[string]string{"cert": string(centralCA)})

	return &v
}

func (t Translator) checkRequiredTLSSecrets(ctx context.Context, sc securedcluster.SecuredCluster) error {
	var finalErr error
	for _, name := range []string{sensorTLSSecretName, admissionControlTLSSecretName, collectorTLSSecretName} {
		if err := t.checkInitBundleSecret(ctx, sc, name); err != nil {
			finalErr = multierror.Append(finalErr, err)
		}
	}
	return finalErr
}

func (t Translator) checkInitBundleSecret(ctx context.Context, sc securedcluster.SecuredCluster, secretName string) error {
	if _, err := t.clientSet.CoreV1().Secrets(sc.Namespace).Get(ctx, secretName, metav1.GetOptions{}); err != nil {
		if k8sErrors.IsNotFound(err) {
			return errors.Wrapf(err, "init-bundle secret %q does not exist, please make sure you have downloaded init-bundle secrets (from UI or with roxctl) and created corresponding resources in the cluster", secretName)
		}
		return errors.Wrapf(err, "failed receiving secret %q", secretName)
	}
	return nil
}

func getSensorValues(ctx context.Context, clientSet kubernetes.Interface, namespace string, sensor *securedcluster.SensorComponentSpec) *translation.ValuesBuilder {
	sv := translation.NewValuesBuilder()

	sv.SetPullPolicy("imagePullPolicy", sensor.ImagePullPolicy)
	sv.AddChild(translation.ResourcesKey, translation.GetResources(sensor.Resources))
	sv.AddAllFrom(translation.GetServiceTLS(ctx, clientSet, namespace, sensor.ServiceTLS, "spec.sensor.serviceTLS"))
	sv.SetStringMap("nodeSelector", sensor.NodeSelector)
	sv.SetString("endpoint", sensor.Endpoint)

	return &sv
}

func (t Translator) getAdmissionControlValues(ctx context.Context, namespace string, admissionControl *securedcluster.AdmissionControlComponentSpec) *translation.ValuesBuilder {
	acv := translation.NewValuesBuilder()

	acv.SetPullPolicy("imagePullPolicy", admissionControl.ImagePullPolicy)
	acv.AddChild(translation.ResourcesKey, translation.GetResources(admissionControl.Resources))
	acv.AddAllFrom(translation.GetServiceTLS(ctx, t.clientSet, namespace, admissionControl.ServiceTLS, "spec.admissionControl.serviceTLS"))
	acv.SetBool("listenOnCreates", admissionControl.ListenOnCreates)
	acv.SetBool("listenOnUpdates", admissionControl.ListenOnUpdates)
	acv.SetBool("listenOnEvents", admissionControl.ListenOnEvents)

	return &acv
}

func (t Translator) getCollectorValues(ctx context.Context, namespace string, collector *securedcluster.CollectorComponentSpec) *translation.ValuesBuilder {
	cv := translation.NewValuesBuilder()

	if collector.Collection != nil {
		switch *collector.Collection {
		case securedcluster.CollectionEBPF:
			cv.SetStringValue("collectionMethod", storage.CollectionMethod_EBPF.String())
		case securedcluster.CollectionKernelModule:
			cv.SetStringValue("collectionMethod", storage.CollectionMethod_KERNEL_MODULE.String())
		case securedcluster.CollectionNone:
			cv.SetStringValue("collectionMethod", storage.CollectionMethod_NO_COLLECTION.String())
		default:
			return cv.SetError(fmt.Errorf("invalid spec.collector.collection %q", *collector.Collection))
		}
	}

	if collector.TaintToleration != nil {
		switch *collector.TaintToleration {
		case securedcluster.TaintTolerate:
			cv.SetBoolValue("disableTaintTolerations", false)
		case securedcluster.TaintAvoid:
			cv.SetBoolValue("disableTaintTolerations", true)
		default:
			return cv.SetError(fmt.Errorf("invalid spec.collector.taintToleration %q", *collector.TaintToleration))
		}
	}

	cv.AddAllFrom(t.getCollectorContainerValues(collector.Collector))
	cv.AddAllFrom(t.getComplianceContainerValues(collector.Compliance))
	cv.AddAllFrom(translation.GetServiceTLS(ctx, t.clientSet, namespace, collector.ServiceTLS, "spec.collector.serviceTLS"))

	return &cv
}

func (t Translator) getCollectorContainerValues(collectorContainerSpec *securedcluster.CollectorContainerSpec) *translation.ValuesBuilder {
	if collectorContainerSpec == nil {
		return nil
	}

	cv := translation.NewValuesBuilder()
	if collectorContainerSpec.ImageFlavor != nil {
		switch *collectorContainerSpec.ImageFlavor {
		case securedcluster.ImageFlavorSlim:
			cv.SetBoolValue("slimMode", true)
		case securedcluster.ImageFlavorRegular:
			cv.SetBoolValue("slimMode", false)
		default:
			return cv.SetError(fmt.Errorf("invalid spec.collector.collector.imageFlavor %q", *collectorContainerSpec.ImageFlavor))
		}
	}

	cv.SetPullPolicy("imagePullPolicy", collectorContainerSpec.ImagePullPolicy)
	cv.AddChild(translation.ResourcesKey, translation.GetResources(collectorContainerSpec.Resources))

	// TODO(ROX-7176): make "customize" work for collector container
	return &cv
}

func (t Translator) getComplianceContainerValues(compliance *securedcluster.ContainerSpec) *translation.ValuesBuilder {
	if compliance == nil {
		return nil
	}

	cv := translation.NewValuesBuilder()
	cv.SetPullPolicy("complianceImagePullPolicy", compliance.ImagePullPolicy)
	cv.AddChild("complianceResources", translation.GetResources(compliance.Resources))

	// TODO(ROX-7176): make "customize" work for compliance container
	return &cv
}

func (t Translator) getHelmDevelopmentDefaults(sc securedcluster.SecuredCluster) *translation.ValuesBuilder {
	rootBuilder := translation.NewValuesBuilder()

	// Set development image from docker hub
	imageBuilder := translation.NewValuesBuilder()
	if os.Getenv("MAIN_IMAGE_TAG") != "" {
		imageBuilder.SetStringValue("tag", os.Getenv("MAIN_IMAGE_TAG"))
	}
	rootBuilder.AddChild("image", &imageBuilder)

	// Set development image pull secrets if no secret was referenced in Spec
	if len(sc.Spec.ImagePullSecrets) == 0 {
		rootBuilder.SetStringMap("imagePullSecrets", map[string]string{
			"username": os.Getenv("REGISTRY_USERNAME"),
			"password": os.Getenv("REGISTRY_PASSWORD"),
		})
	}

	return &rootBuilder
}