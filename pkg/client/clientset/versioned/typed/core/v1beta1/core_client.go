/*
Copyright 2018 The CDI Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"net/http"

	rest "k8s.io/client-go/rest"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
	"kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/scheme"
)

type CdiV1beta1Interface interface {
	RESTClient() rest.Interface
	CDIsGetter
	CDIConfigsGetter
	DataImportCronsGetter
	DataSourcesGetter
	DataVolumesGetter
	ObjectTransfersGetter
	StorageProfilesGetter
	VolumeImportSourcesGetter
	VolumeUploadSourcesGetter
}

// CdiV1beta1Client is used to interact with features provided by the cdi.kubevirt.io group.
type CdiV1beta1Client struct {
	restClient rest.Interface
}

func (c *CdiV1beta1Client) CDIs() CDIInterface {
	return newCDIs(c)
}

func (c *CdiV1beta1Client) CDIConfigs() CDIConfigInterface {
	return newCDIConfigs(c)
}

func (c *CdiV1beta1Client) DataImportCrons(namespace string) DataImportCronInterface {
	return newDataImportCrons(c, namespace)
}

func (c *CdiV1beta1Client) DataSources(namespace string) DataSourceInterface {
	return newDataSources(c, namespace)
}

func (c *CdiV1beta1Client) DataVolumes(namespace string) DataVolumeInterface {
	return newDataVolumes(c, namespace)
}

func (c *CdiV1beta1Client) ObjectTransfers() ObjectTransferInterface {
	return newObjectTransfers(c)
}

func (c *CdiV1beta1Client) StorageProfiles() StorageProfileInterface {
	return newStorageProfiles(c)
}

func (c *CdiV1beta1Client) VolumeImportSources(namespace string) VolumeImportSourceInterface {
	return newVolumeImportSources(c, namespace)
}

func (c *CdiV1beta1Client) VolumeUploadSources(namespace string) VolumeUploadSourceInterface {
	return newVolumeUploadSources(c, namespace)
}

// NewForConfig creates a new CdiV1beta1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*CdiV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new CdiV1beta1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*CdiV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &CdiV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new CdiV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CdiV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CdiV1beta1Client for the given RESTClient.
func New(c rest.Interface) *CdiV1beta1Client {
	return &CdiV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CdiV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
