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

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeProjectLoggings implements ProjectLoggingInterface
type FakeProjectLoggings struct {
	Fake *FakeManagementV3
	ns   string
}

var projectloggingsResource = schema.GroupVersionResource{Group: "management.cattle.io", Version: "v3", Resource: "projectloggings"}

var projectloggingsKind = schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectLogging"}

// Get takes name of the projectLogging, and returns the corresponding projectLogging object, and an error if there is any.
func (c *FakeProjectLoggings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ProjectLogging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(projectloggingsResource, c.ns, name), &v3.ProjectLogging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ProjectLogging), err
}

// List takes label and field selectors, and returns the list of ProjectLoggings that match those selectors.
func (c *FakeProjectLoggings) List(ctx context.Context, opts v1.ListOptions) (result *v3.ProjectLoggingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(projectloggingsResource, projectloggingsKind, c.ns, opts), &v3.ProjectLoggingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ProjectLoggingList{ListMeta: obj.(*v3.ProjectLoggingList).ListMeta}
	for _, item := range obj.(*v3.ProjectLoggingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested projectLoggings.
func (c *FakeProjectLoggings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(projectloggingsResource, c.ns, opts))

}

// Create takes the representation of a projectLogging and creates it.  Returns the server's representation of the projectLogging, and an error, if there is any.
func (c *FakeProjectLoggings) Create(ctx context.Context, projectLogging *v3.ProjectLogging, opts v1.CreateOptions) (result *v3.ProjectLogging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(projectloggingsResource, c.ns, projectLogging), &v3.ProjectLogging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ProjectLogging), err
}

// Update takes the representation of a projectLogging and updates it. Returns the server's representation of the projectLogging, and an error, if there is any.
func (c *FakeProjectLoggings) Update(ctx context.Context, projectLogging *v3.ProjectLogging, opts v1.UpdateOptions) (result *v3.ProjectLogging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(projectloggingsResource, c.ns, projectLogging), &v3.ProjectLogging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ProjectLogging), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeProjectLoggings) UpdateStatus(ctx context.Context, projectLogging *v3.ProjectLogging, opts v1.UpdateOptions) (*v3.ProjectLogging, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(projectloggingsResource, "status", c.ns, projectLogging), &v3.ProjectLogging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ProjectLogging), err
}

// Delete takes name of the projectLogging and deletes it. Returns an error if one occurs.
func (c *FakeProjectLoggings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(projectloggingsResource, c.ns, name, opts), &v3.ProjectLogging{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeProjectLoggings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(projectloggingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ProjectLoggingList{})
	return err
}

// Patch applies the patch and returns the patched projectLogging.
func (c *FakeProjectLoggings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ProjectLogging, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(projectloggingsResource, c.ns, name, pt, data, subresources...), &v3.ProjectLogging{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ProjectLogging), err
}
