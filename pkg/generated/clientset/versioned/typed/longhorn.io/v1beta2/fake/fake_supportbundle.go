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

	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSupportBundles implements SupportBundleInterface
type FakeSupportBundles struct {
	Fake *FakeLonghornV1beta2
	ns   string
}

var supportbundlesResource = schema.GroupVersionResource{Group: "longhorn.io", Version: "v1beta2", Resource: "supportbundles"}

var supportbundlesKind = schema.GroupVersionKind{Group: "longhorn.io", Version: "v1beta2", Kind: "SupportBundle"}

// Get takes name of the supportBundle, and returns the corresponding supportBundle object, and an error if there is any.
func (c *FakeSupportBundles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta2.SupportBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(supportbundlesResource, c.ns, name), &v1beta2.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SupportBundle), err
}

// List takes label and field selectors, and returns the list of SupportBundles that match those selectors.
func (c *FakeSupportBundles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.SupportBundleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(supportbundlesResource, supportbundlesKind, c.ns, opts), &v1beta2.SupportBundleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.SupportBundleList{ListMeta: obj.(*v1beta2.SupportBundleList).ListMeta}
	for _, item := range obj.(*v1beta2.SupportBundleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested supportBundles.
func (c *FakeSupportBundles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(supportbundlesResource, c.ns, opts))

}

// Create takes the representation of a supportBundle and creates it.  Returns the server's representation of the supportBundle, and an error, if there is any.
func (c *FakeSupportBundles) Create(ctx context.Context, supportBundle *v1beta2.SupportBundle, opts v1.CreateOptions) (result *v1beta2.SupportBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(supportbundlesResource, c.ns, supportBundle), &v1beta2.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SupportBundle), err
}

// Update takes the representation of a supportBundle and updates it. Returns the server's representation of the supportBundle, and an error, if there is any.
func (c *FakeSupportBundles) Update(ctx context.Context, supportBundle *v1beta2.SupportBundle, opts v1.UpdateOptions) (result *v1beta2.SupportBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(supportbundlesResource, c.ns, supportBundle), &v1beta2.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SupportBundle), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSupportBundles) UpdateStatus(ctx context.Context, supportBundle *v1beta2.SupportBundle, opts v1.UpdateOptions) (*v1beta2.SupportBundle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(supportbundlesResource, "status", c.ns, supportBundle), &v1beta2.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SupportBundle), err
}

// Delete takes name of the supportBundle and deletes it. Returns an error if one occurs.
func (c *FakeSupportBundles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(supportbundlesResource, c.ns, name, opts), &v1beta2.SupportBundle{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSupportBundles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(supportbundlesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.SupportBundleList{})
	return err
}

// Patch applies the patch and returns the patched supportBundle.
func (c *FakeSupportBundles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta2.SupportBundle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(supportbundlesResource, c.ns, name, pt, data, subresources...), &v1beta2.SupportBundle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.SupportBundle), err
}
