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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	forkliftv1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/forklift/v1beta1"
	versioned "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned"
	internalinterfaces "kubevirt.io/containerized-data-importer/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "kubevirt.io/containerized-data-importer/pkg/client/listers/forklift/v1beta1"
)

// OvirtVolumePopulatorInformer provides access to a shared informer and lister for
// OvirtVolumePopulators.
type OvirtVolumePopulatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.OvirtVolumePopulatorLister
}

type ovirtVolumePopulatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewOvirtVolumePopulatorInformer constructs a new informer for OvirtVolumePopulator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOvirtVolumePopulatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOvirtVolumePopulatorInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredOvirtVolumePopulatorInformer constructs a new informer for OvirtVolumePopulator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOvirtVolumePopulatorInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ForkliftV1beta1().OvirtVolumePopulators(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ForkliftV1beta1().OvirtVolumePopulators(namespace).Watch(context.TODO(), options)
			},
		},
		&forkliftv1beta1.OvirtVolumePopulator{},
		resyncPeriod,
		indexers,
	)
}

func (f *ovirtVolumePopulatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOvirtVolumePopulatorInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ovirtVolumePopulatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&forkliftv1beta1.OvirtVolumePopulator{}, f.defaultInformer)
}

func (f *ovirtVolumePopulatorInformer) Lister() v1beta1.OvirtVolumePopulatorLister {
	return v1beta1.NewOvirtVolumePopulatorLister(f.Informer().GetIndexer())
}
