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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// VolumeImportSourceLister helps list VolumeImportSources.
// All objects returned here must be treated as read-only.
type VolumeImportSourceLister interface {
	// List lists all VolumeImportSources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeImportSource, err error)
	// VolumeImportSources returns an object that can list and get VolumeImportSources.
	VolumeImportSources(namespace string) VolumeImportSourceNamespaceLister
	VolumeImportSourceListerExpansion
}

// volumeImportSourceLister implements the VolumeImportSourceLister interface.
type volumeImportSourceLister struct {
	indexer cache.Indexer
}

// NewVolumeImportSourceLister returns a new VolumeImportSourceLister.
func NewVolumeImportSourceLister(indexer cache.Indexer) VolumeImportSourceLister {
	return &volumeImportSourceLister{indexer: indexer}
}

// List lists all VolumeImportSources in the indexer.
func (s *volumeImportSourceLister) List(selector labels.Selector) (ret []*v1beta1.VolumeImportSource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeImportSource))
	})
	return ret, err
}

// VolumeImportSources returns an object that can list and get VolumeImportSources.
func (s *volumeImportSourceLister) VolumeImportSources(namespace string) VolumeImportSourceNamespaceLister {
	return volumeImportSourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VolumeImportSourceNamespaceLister helps list and get VolumeImportSources.
// All objects returned here must be treated as read-only.
type VolumeImportSourceNamespaceLister interface {
	// List lists all VolumeImportSources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.VolumeImportSource, err error)
	// Get retrieves the VolumeImportSource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.VolumeImportSource, error)
	VolumeImportSourceNamespaceListerExpansion
}

// volumeImportSourceNamespaceLister implements the VolumeImportSourceNamespaceLister
// interface.
type volumeImportSourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VolumeImportSources in the indexer for a given namespace.
func (s volumeImportSourceNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VolumeImportSource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VolumeImportSource))
	})
	return ret, err
}

// Get retrieves the VolumeImportSource from the indexer for a given namespace and name.
func (s volumeImportSourceNamespaceLister) Get(name string) (*v1beta1.VolumeImportSource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("volumeimportsource"), name)
	}
	return obj.(*v1beta1.VolumeImportSource), nil
}
