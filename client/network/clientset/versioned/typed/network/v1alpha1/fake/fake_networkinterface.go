/*
Copyright 2024 Google LLC
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/network/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkInterfaces implements NetworkInterfaceInterface
type FakeNetworkInterfaces struct {
	Fake *FakeNetworkingV1alpha1
	ns   string
}

var networkinterfacesResource = v1alpha1.SchemeGroupVersion.WithResource("networkinterfaces")

var networkinterfacesKind = v1alpha1.SchemeGroupVersion.WithKind("NetworkInterface")

// Get takes name of the networkInterface, and returns the corresponding networkInterface object, and an error if there is any.
func (c *FakeNetworkInterfaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkinterfacesResource, c.ns, name), &v1alpha1.NetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// List takes label and field selectors, and returns the list of NetworkInterfaces that match those selectors.
func (c *FakeNetworkInterfaces) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NetworkInterfaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkinterfacesResource, networkinterfacesKind, c.ns, opts), &v1alpha1.NetworkInterfaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NetworkInterfaceList{ListMeta: obj.(*v1alpha1.NetworkInterfaceList).ListMeta}
	for _, item := range obj.(*v1alpha1.NetworkInterfaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networkInterfaces.
func (c *FakeNetworkInterfaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkinterfacesResource, c.ns, opts))

}

// Create takes the representation of a networkInterface and creates it.  Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *FakeNetworkInterfaces) Create(ctx context.Context, networkInterface *v1alpha1.NetworkInterface, opts v1.CreateOptions) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkinterfacesResource, c.ns, networkInterface), &v1alpha1.NetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Update takes the representation of a networkInterface and updates it. Returns the server's representation of the networkInterface, and an error, if there is any.
func (c *FakeNetworkInterfaces) Update(ctx context.Context, networkInterface *v1alpha1.NetworkInterface, opts v1.UpdateOptions) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkinterfacesResource, c.ns, networkInterface), &v1alpha1.NetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNetworkInterfaces) UpdateStatus(ctx context.Context, networkInterface *v1alpha1.NetworkInterface, opts v1.UpdateOptions) (*v1alpha1.NetworkInterface, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networkinterfacesResource, "status", c.ns, networkInterface), &v1alpha1.NetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}

// Delete takes name of the networkInterface and deletes it. Returns an error if one occurs.
func (c *FakeNetworkInterfaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(networkinterfacesResource, c.ns, name, opts), &v1alpha1.NetworkInterface{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNetworkInterfaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkinterfacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NetworkInterfaceList{})
	return err
}

// Patch applies the patch and returns the patched networkInterface.
func (c *FakeNetworkInterfaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NetworkInterface, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkinterfacesResource, c.ns, name, pt, data, subresources...), &v1alpha1.NetworkInterface{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NetworkInterface), err
}
