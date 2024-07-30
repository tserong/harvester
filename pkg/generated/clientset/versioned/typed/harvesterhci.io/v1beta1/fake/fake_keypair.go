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

// FakeKeyPairs implements KeyPairInterface
type FakeKeyPairs struct {
	Fake *FakeHarvesterhciV1beta1
	ns   string
}

var keypairsResource = v1beta1.SchemeGroupVersion.WithResource("keypairs")

var keypairsKind = v1beta1.SchemeGroupVersion.WithKind("KeyPair")

// Get takes name of the keyPair, and returns the corresponding keyPair object, and an error if there is any.
func (c *FakeKeyPairs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KeyPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(keypairsResource, c.ns, name), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

// List takes label and field selectors, and returns the list of KeyPairs that match those selectors.
func (c *FakeKeyPairs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KeyPairList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(keypairsResource, keypairsKind, c.ns, opts), &v1beta1.KeyPairList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.KeyPairList{ListMeta: obj.(*v1beta1.KeyPairList).ListMeta}
	for _, item := range obj.(*v1beta1.KeyPairList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested keyPairs.
func (c *FakeKeyPairs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(keypairsResource, c.ns, opts))

}

// Create takes the representation of a keyPair and creates it.  Returns the server's representation of the keyPair, and an error, if there is any.
func (c *FakeKeyPairs) Create(ctx context.Context, keyPair *v1beta1.KeyPair, opts v1.CreateOptions) (result *v1beta1.KeyPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(keypairsResource, c.ns, keyPair), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

// Update takes the representation of a keyPair and updates it. Returns the server's representation of the keyPair, and an error, if there is any.
func (c *FakeKeyPairs) Update(ctx context.Context, keyPair *v1beta1.KeyPair, opts v1.UpdateOptions) (result *v1beta1.KeyPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(keypairsResource, c.ns, keyPair), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKeyPairs) UpdateStatus(ctx context.Context, keyPair *v1beta1.KeyPair, opts v1.UpdateOptions) (*v1beta1.KeyPair, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(keypairsResource, "status", c.ns, keyPair), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}

// Delete takes name of the keyPair and deletes it. Returns an error if one occurs.
func (c *FakeKeyPairs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(keypairsResource, c.ns, name, opts), &v1beta1.KeyPair{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKeyPairs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(keypairsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.KeyPairList{})
	return err
}

// Patch applies the patch and returns the patched keyPair.
func (c *FakeKeyPairs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KeyPair, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(keypairsResource, c.ns, name, pt, data, subresources...), &v1beta1.KeyPair{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KeyPair), err
}
