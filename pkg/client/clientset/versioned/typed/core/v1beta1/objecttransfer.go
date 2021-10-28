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
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
	scheme "kubevirt.io/containerized-data-importer/pkg/client/clientset/versioned/scheme"
)

// ObjectTransfersGetter has a method to return a ObjectTransferInterface.
// A group's client should implement this interface.
type ObjectTransfersGetter interface {
	ObjectTransfers() ObjectTransferInterface
}

// ObjectTransferInterface has methods to work with ObjectTransfer resources.
type ObjectTransferInterface interface {
	Create(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.CreateOptions) (*v1beta1.ObjectTransfer, error)
	Update(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.UpdateOptions) (*v1beta1.ObjectTransfer, error)
	UpdateStatus(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.UpdateOptions) (*v1beta1.ObjectTransfer, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ObjectTransfer, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ObjectTransferList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ObjectTransfer, err error)
	ObjectTransferExpansion
}

// objectTransfers implements ObjectTransferInterface
type objectTransfers struct {
	client rest.Interface
}

// newObjectTransfers returns a ObjectTransfers
func newObjectTransfers(c *CdiV1beta1Client) *objectTransfers {
	return &objectTransfers{
		client: c.RESTClient(),
	}
}

// Get takes name of the objectTransfer, and returns the corresponding objectTransfer object, and an error if there is any.
func (c *objectTransfers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ObjectTransfer, err error) {
	result = &v1beta1.ObjectTransfer{}
	err = c.client.Get().
		Resource("objecttransfers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ObjectTransfers that match those selectors.
func (c *objectTransfers) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ObjectTransferList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ObjectTransferList{}
	err = c.client.Get().
		Resource("objecttransfers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested objectTransfers.
func (c *objectTransfers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("objecttransfers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a objectTransfer and creates it.  Returns the server's representation of the objectTransfer, and an error, if there is any.
func (c *objectTransfers) Create(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.CreateOptions) (result *v1beta1.ObjectTransfer, err error) {
	result = &v1beta1.ObjectTransfer{}
	err = c.client.Post().
		Resource("objecttransfers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(objectTransfer).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a objectTransfer and updates it. Returns the server's representation of the objectTransfer, and an error, if there is any.
func (c *objectTransfers) Update(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.UpdateOptions) (result *v1beta1.ObjectTransfer, err error) {
	result = &v1beta1.ObjectTransfer{}
	err = c.client.Put().
		Resource("objecttransfers").
		Name(objectTransfer.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(objectTransfer).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *objectTransfers) UpdateStatus(ctx context.Context, objectTransfer *v1beta1.ObjectTransfer, opts v1.UpdateOptions) (result *v1beta1.ObjectTransfer, err error) {
	result = &v1beta1.ObjectTransfer{}
	err = c.client.Put().
		Resource("objecttransfers").
		Name(objectTransfer.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(objectTransfer).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the objectTransfer and deletes it. Returns an error if one occurs.
func (c *objectTransfers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("objecttransfers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *objectTransfers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("objecttransfers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched objectTransfer.
func (c *objectTransfers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ObjectTransfer, err error) {
	result = &v1beta1.ObjectTransfer{}
	err = c.client.Patch(pt).
		Resource("objecttransfers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
