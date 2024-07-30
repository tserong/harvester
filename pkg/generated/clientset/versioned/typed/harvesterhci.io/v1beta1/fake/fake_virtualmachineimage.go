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

	v1beta1 "github.com/harvester/harvester/pkg/apis/harvesterhci.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMachineImages implements VirtualMachineImageInterface
type FakeVirtualMachineImages struct {
	Fake *FakeHarvesterhciV1beta1
	ns   string
}

var virtualmachineimagesResource = v1beta1.SchemeGroupVersion.WithResource("virtualmachineimages")

var virtualmachineimagesKind = v1beta1.SchemeGroupVersion.WithKind("VirtualMachineImage")

// Get takes name of the virtualMachineImage, and returns the corresponding virtualMachineImage object, and an error if there is any.
func (c *FakeVirtualMachineImages) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(virtualmachineimagesResource, c.ns, name), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

// List takes label and field selectors, and returns the list of VirtualMachineImages that match those selectors.
func (c *FakeVirtualMachineImages) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VirtualMachineImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(virtualmachineimagesResource, virtualmachineimagesKind, c.ns, opts), &v1beta1.VirtualMachineImageList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VirtualMachineImageList{ListMeta: obj.(*v1beta1.VirtualMachineImageList).ListMeta}
	for _, item := range obj.(*v1beta1.VirtualMachineImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachineImages.
func (c *FakeVirtualMachineImages) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(virtualmachineimagesResource, c.ns, opts))

}

// Create takes the representation of a virtualMachineImage and creates it.  Returns the server's representation of the virtualMachineImage, and an error, if there is any.
func (c *FakeVirtualMachineImages) Create(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.CreateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(virtualmachineimagesResource, c.ns, virtualMachineImage), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

// Update takes the representation of a virtualMachineImage and updates it. Returns the server's representation of the virtualMachineImage, and an error, if there is any.
func (c *FakeVirtualMachineImages) Update(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (result *v1beta1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(virtualmachineimagesResource, c.ns, virtualMachineImage), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVirtualMachineImages) UpdateStatus(ctx context.Context, virtualMachineImage *v1beta1.VirtualMachineImage, opts v1.UpdateOptions) (*v1beta1.VirtualMachineImage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(virtualmachineimagesResource, "status", c.ns, virtualMachineImage), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}

// Delete takes name of the virtualMachineImage and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachineImages) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(virtualmachineimagesResource, c.ns, name, opts), &v1beta1.VirtualMachineImage{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachineImages) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(virtualmachineimagesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VirtualMachineImageList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachineImage.
func (c *FakeVirtualMachineImages) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(virtualmachineimagesResource, c.ns, name, pt, data, subresources...), &v1beta1.VirtualMachineImage{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VirtualMachineImage), err
}
