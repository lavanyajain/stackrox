package quay

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/logging"
	"bitbucket.org/stack-rox/apollo/pkg/registries"
	"bitbucket.org/stack-rox/apollo/pkg/registries/docker"
	"bitbucket.org/stack-rox/apollo/pkg/urlfmt"
)

const (
	oauthTokenString = "$oauthtoken"
)

var (
	logger = logging.LoggerForModule()
)

// Quay is the implementation of the Docker Registry for Quay
type Quay struct {
	*docker.Registry
	config *v1.QuayConfig
}

func validate(quay *v1.QuayConfig) error {
	if quay.GetEndpoint() == "" {
		return fmt.Errorf("Quay endpoint must be specified")
	}
	// Note that the oauth token could be empty because there are public images
	return nil
}

// NewRegistryFromConfig returns a new instantiation of the Quay registry
func NewRegistryFromConfig(config *v1.QuayConfig, integration *v1.ImageIntegration) (registries.ImageRegistry, error) {
	if err := validate(config); err != nil {
		return nil, err
	}

	var username string
	if config.GetOauthToken() != "" {
		username = oauthTokenString
	}

	cfg := docker.Config{
		Username: username,
		Password: config.GetOauthToken(),
		Endpoint: config.GetEndpoint(),
	}
	dockerRegistry, err := docker.NewDockerRegistry(cfg, integration)
	if err != nil {
		return nil, err
	}
	return &Quay{
		Registry: dockerRegistry,
		config:   config,
	}, nil
}

func newRegistry(integration *v1.ImageIntegration) (registries.ImageRegistry, error) {
	quayConfig, ok := integration.IntegrationConfig.(*v1.ImageIntegration_Quay)
	if !ok {
		return nil, fmt.Errorf("Quay config must be specified")
	}
	return NewRegistryFromConfig(quayConfig.Quay, integration)
}

// Test overrides the default docker Test function because the Quay Ping endpoint requires Auth
func (q *Quay) Test() error {
	if q.config.GetOauthToken() != "" {
		return q.Registry.Test()
	}

	url, err := urlfmt.FormatURL(q.config.GetEndpoint(), urlfmt.HTTPS, urlfmt.NoTrailingSlash)
	if err != nil {
		return err
	}
	discoveryURL := url + "/api/v1/discovery"
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	resp, err := client.Get(discoveryURL)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		body, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			return fmt.Errorf("Error reaching quay.io with HTTP code %d", resp.StatusCode)
		}
		return fmt.Errorf("Error reaching quay.io with HTTP code %d: %s", resp.StatusCode, string(body))
	}
	return nil
}

func init() {
	registries.Registry["quay"] = func(integration *v1.ImageIntegration) (registries.ImageRegistry, error) {
		reg, err := newRegistry(integration)
		return reg, err
	}
}
