package translation

import (
	"github.com/pkg/errors"
	platform "github.com/stackrox/rox/operator/apis/platform/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/utils/pointer"
)

const (
	// ResourcesKey is a key for most resources chart values.
	ResourcesKey = "resources"
	// TolerationsKey is the default tolerations key used in the charts.
	TolerationsKey = "tolerations"
)

// GetResources converts platform.Resources to chart values builder.
func GetResources(resources *corev1.ResourceRequirements) *ValuesBuilder {
	if resources == nil {
		return nil
	}
	res := NewValuesBuilder()

	if len(resources.Requests) > 0 {
		res.SetResourceList("requests", resources.Requests.DeepCopy())
	}
	if len(resources.Limits) > 0 {
		res.SetResourceList("limits", resources.Limits.DeepCopy())
	}
	return &res
}

// GetCustomize converts platform.CustomizeSpec to chart values builder.
func GetCustomize(customizeSpec *platform.CustomizeSpec) *ValuesBuilder {
	if customizeSpec == nil {
		return nil
	}

	res := NewValuesBuilder()
	res.SetStringMap("labels", customizeSpec.Labels)
	res.SetStringMap("annotations", customizeSpec.Annotations)
	envVarMap := make(map[string]interface{}, len(customizeSpec.EnvVars))
	for i := range customizeSpec.EnvVars {
		envVar := customizeSpec.EnvVars[i]
		if _, ok := envVarMap[envVar.Name]; ok {
			res.SetError(errors.Errorf("duplicate environment variable name %q", envVar.Name))
			return &res
		}

		// We need the content of the env var without the name for the Helm charts. We cannot set the name to "",
		// since it doesn't have an omitempty tag. We could create a `map[string]interface{}` with `Value` and
		// `ValueFrom` ported over, but that would break if Kubernetes ever adds to the corev1.EnvVar type.
		// Hence, rely on unstructured conversion.
		unstructuredEnvVar, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&envVar)
		if err != nil {
			res.SetError(errors.Wrapf(err, "failed parsing environment variable %q", envVar.Name))
			return &res
		}
		delete(unstructuredEnvVar, "name")
		envVarMap[envVar.Name] = unstructuredEnvVar
	}
	res.SetMap("envVars", envVarMap)
	return &res
}

// GetMisc converts platform.MiscSpec to chart values builder.
func GetMisc(miscSpec *platform.MiscSpec) *ValuesBuilder {
	if miscSpec == nil {
		return nil
	}
	if !pointer.BoolPtrDerefOr(miscSpec.CreateSCCs, false) {
		return nil
	}

	res := NewValuesBuilder()
	res.SetMap("system", map[string]interface{}{"createSCCs": true})
	return &res
}

// GetImagePullSecrets converts corev1.LocalObjectReference to a *ValuesBuilder with an "imagePullSecrets" field.
func GetImagePullSecrets(imagePullSecrets []platform.LocalSecretReference) *ValuesBuilder {
	res := NewValuesBuilder()
	if len(imagePullSecrets) > 0 {
		var ps []string
		for _, secret := range imagePullSecrets {
			ps = append(ps, secret.Name)
		}
		existing := NewValuesBuilder()
		existing.SetStringSlice("useExisting", ps)
		res.AddChild("imagePullSecrets", &existing)
	}
	return &res
}

// GetTLSConfigValues converts platform.TLSConfig to a *ValuesBuilder with an "additionalCAs" field.
func GetTLSConfigValues(tls *platform.TLSConfig) *ValuesBuilder {
	if tls == nil || len(tls.AdditionalCAs) == 0 {
		return nil
	}
	cas := NewValuesBuilder()
	for _, ca := range tls.AdditionalCAs {
		cas.SetStringValue(ca.Name, ca.Content)
	}
	res := NewValuesBuilder()
	res.AddChild("additionalCAs", &cas)
	return &res
}

// GetTolerations converts a slice of tolerations to a *ValuesBuilder object and sets the field name
// based on the key parameter.
func GetTolerations(key string, tolerations []*corev1.Toleration) *ValuesBuilder {
	v := NewValuesBuilder()

	var convertedList []interface{}
	for _, toleration := range tolerations {
		m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(toleration)
		if err != nil {
			v.SetError(errors.Wrapf(err, "failed converting %q to unstructured", key))
			break
		}
		convertedList = append(convertedList, m)
	}
	v.SetSlice(key, convertedList)

	return &v
}
