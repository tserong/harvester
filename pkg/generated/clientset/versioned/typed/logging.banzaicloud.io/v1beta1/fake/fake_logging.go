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

package fake

import (
	"context"

	v1beta1 "github.com/kube-logging/logging-operator/pkg/sdk/logging/api/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLoggings implements LoggingInterface
type FakeLoggings struct {
	Fake *FakeLoggingV1beta1
}

var loggingsResource = v1beta1.SchemeGroupVersion.WithResource("loggings")

var loggingsKind = v1beta1.SchemeGroupVersion.WithKind("Logging")

// Get takes name of the logging, and returns the corresponding logging object, and an error if there is any.
func (c *FakeLoggings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(loggingsResource, name), &v1beta1.Logging{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Logging), err
}

// List takes label and field selectors, and returns the list of Loggings that match those selectors.
func (c *FakeLoggings) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.LoggingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(loggingsResource, loggingsKind, opts), &v1beta1.LoggingList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.LoggingList{ListMeta: obj.(*v1beta1.LoggingList).ListMeta}
	for _, item := range obj.(*v1beta1.LoggingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested loggings.
func (c *FakeLoggings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(loggingsResource, opts))
}

// Create takes the representation of a logging and creates it.  Returns the server's representation of the logging, and an error, if there is any.
func (c *FakeLoggings) Create(ctx context.Context, logging *v1beta1.Logging, opts v1.CreateOptions) (result *v1beta1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(loggingsResource, logging), &v1beta1.Logging{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Logging), err
}

// Update takes the representation of a logging and updates it. Returns the server's representation of the logging, and an error, if there is any.
func (c *FakeLoggings) Update(ctx context.Context, logging *v1beta1.Logging, opts v1.UpdateOptions) (result *v1beta1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(loggingsResource, logging), &v1beta1.Logging{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Logging), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLoggings) UpdateStatus(ctx context.Context, logging *v1beta1.Logging, opts v1.UpdateOptions) (*v1beta1.Logging, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(loggingsResource, "status", logging), &v1beta1.Logging{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Logging), err
}

// Delete takes name of the logging and deletes it. Returns an error if one occurs.
func (c *FakeLoggings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(loggingsResource, name, opts), &v1beta1.Logging{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLoggings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(loggingsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.LoggingList{})
	return err
}

// Patch applies the patch and returns the patched logging.
func (c *FakeLoggings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Logging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(loggingsResource, name, pt, data, subresources...), &v1beta1.Logging{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Logging), err
}
