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

// ClusterProxyConfigsGetter has a method to return a ClusterProxyConfigInterface.
// A group's client should implement this interface.
type ClusterProxyConfigsGetter interface {
	ClusterProxyConfigs(namespace string) ClusterProxyConfigInterface
}

// ClusterProxyConfigInterface has methods to work with ClusterProxyConfig resources.
type ClusterProxyConfigInterface interface {
	Create(ctx context.Context, clusterProxyConfig *v3.ClusterProxyConfig, opts v1.CreateOptions) (*v3.ClusterProxyConfig, error)
	Update(ctx context.Context, clusterProxyConfig *v3.ClusterProxyConfig, opts v1.UpdateOptions) (*v3.ClusterProxyConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.ClusterProxyConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.ClusterProxyConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterProxyConfig, err error)
	ClusterProxyConfigExpansion
}

// clusterProxyConfigs implements ClusterProxyConfigInterface
type clusterProxyConfigs struct {
	client rest.Interface
	ns     string
}

// newClusterProxyConfigs returns a ClusterProxyConfigs
func newClusterProxyConfigs(c *ManagementV3Client, namespace string) *clusterProxyConfigs {
	return &clusterProxyConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterProxyConfig, and returns the corresponding clusterProxyConfig object, and an error if there is any.
func (c *clusterProxyConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterProxyConfig, err error) {
	result = &v3.ClusterProxyConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterProxyConfigs that match those selectors.
func (c *clusterProxyConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterProxyConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.ClusterProxyConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterProxyConfigs.
func (c *clusterProxyConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterProxyConfig and creates it.  Returns the server's representation of the clusterProxyConfig, and an error, if there is any.
func (c *clusterProxyConfigs) Create(ctx context.Context, clusterProxyConfig *v3.ClusterProxyConfig, opts v1.CreateOptions) (result *v3.ClusterProxyConfig, err error) {
	result = &v3.ClusterProxyConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterProxyConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterProxyConfig and updates it. Returns the server's representation of the clusterProxyConfig, and an error, if there is any.
func (c *clusterProxyConfigs) Update(ctx context.Context, clusterProxyConfig *v3.ClusterProxyConfig, opts v1.UpdateOptions) (result *v3.ClusterProxyConfig, err error) {
	result = &v3.ClusterProxyConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		Name(clusterProxyConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterProxyConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterProxyConfig and deletes it. Returns an error if one occurs.
func (c *clusterProxyConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterProxyConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterProxyConfig.
func (c *clusterProxyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterProxyConfig, err error) {
	result = &v3.ClusterProxyConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterproxyconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
