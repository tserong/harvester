/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	scheme "github.com/harvester/harvester/pkg/generated/clientset/versioned/scheme"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MultiClusterAppsGetter has a method to return a MultiClusterAppInterface.
// A group's client should implement this interface.
type MultiClusterAppsGetter interface {
	MultiClusterApps(namespace string) MultiClusterAppInterface
}

// MultiClusterAppInterface has methods to work with MultiClusterApp resources.
type MultiClusterAppInterface interface {
	Create(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.CreateOptions) (*v3.MultiClusterApp, error)
	Update(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.UpdateOptions) (*v3.MultiClusterApp, error)
	UpdateStatus(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.UpdateOptions) (*v3.MultiClusterApp, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.MultiClusterApp, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.MultiClusterAppList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.MultiClusterApp, err error)
	MultiClusterAppExpansion
}

// multiClusterApps implements MultiClusterAppInterface
type multiClusterApps struct {
	client rest.Interface
	ns     string
}

// newMultiClusterApps returns a MultiClusterApps
func newMultiClusterApps(c *ManagementV3Client, namespace string) *multiClusterApps {
	return &multiClusterApps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the multiClusterApp, and returns the corresponding multiClusterApp object, and an error if there is any.
func (c *multiClusterApps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.MultiClusterApp, err error) {
	result = &v3.MultiClusterApp{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterapps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MultiClusterApps that match those selectors.
func (c *multiClusterApps) List(ctx context.Context, opts v1.ListOptions) (result *v3.MultiClusterAppList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.MultiClusterAppList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested multiClusterApps.
func (c *multiClusterApps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("multiclusterapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a multiClusterApp and creates it.  Returns the server's representation of the multiClusterApp, and an error, if there is any.
func (c *multiClusterApps) Create(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.CreateOptions) (result *v3.MultiClusterApp, err error) {
	result = &v3.MultiClusterApp{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("multiclusterapps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterApp).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a multiClusterApp and updates it. Returns the server's representation of the multiClusterApp, and an error, if there is any.
func (c *multiClusterApps) Update(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.UpdateOptions) (result *v3.MultiClusterApp, err error) {
	result = &v3.MultiClusterApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusterapps").
		Name(multiClusterApp.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterApp).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *multiClusterApps) UpdateStatus(ctx context.Context, multiClusterApp *v3.MultiClusterApp, opts v1.UpdateOptions) (result *v3.MultiClusterApp, err error) {
	result = &v3.MultiClusterApp{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("multiclusterapps").
		Name(multiClusterApp.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(multiClusterApp).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the multiClusterApp and deletes it. Returns an error if one occurs.
func (c *multiClusterApps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusterapps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *multiClusterApps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("multiclusterapps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched multiClusterApp.
func (c *multiClusterApps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.MultiClusterApp, err error) {
	result = &v3.MultiClusterApp{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("multiclusterapps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
