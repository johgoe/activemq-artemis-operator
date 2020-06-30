/*
Copyright The Kubernetes Authors.

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

package fake

import (
	v2alpha1 "github.com/artemiscloud/activemq-artemis-operator/pkg/apis/broker/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeActiveMQArtemisScaledowns implements ActiveMQArtemisScaledownInterface
type FakeActiveMQArtemisScaledowns struct {
	Fake *FakeBrokerV2alpha1
	ns   string
}

var activemqartemisscaledownsResource = schema.GroupVersionResource{Group: "broker.amq.io", Version: "v2alpha1", Resource: "activemqartemisscaledowns"}

var activemqartemisscaledownsKind = schema.GroupVersionKind{Group: "broker.amq.io", Version: "v2alpha1", Kind: "ActiveMQArtemisScaledown"}

// Get takes name of the activeMQArtemisScaledown, and returns the corresponding activeMQArtemisScaledown object, and an error if there is any.
func (c *FakeActiveMQArtemisScaledowns) Get(name string, options v1.GetOptions) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(activemqartemisscaledownsResource, c.ns, name), &v2alpha1.ActiveMQArtemisScaledown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ActiveMQArtemisScaledown), err
}

// List takes label and field selectors, and returns the list of ActiveMQArtemisScaledowns that match those selectors.
func (c *FakeActiveMQArtemisScaledowns) List(opts v1.ListOptions) (result *v2alpha1.ActiveMQArtemisScaledownList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(activemqartemisscaledownsResource, activemqartemisscaledownsKind, c.ns, opts), &v2alpha1.ActiveMQArtemisScaledownList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.ActiveMQArtemisScaledownList{ListMeta: obj.(*v2alpha1.ActiveMQArtemisScaledownList).ListMeta}
	for _, item := range obj.(*v2alpha1.ActiveMQArtemisScaledownList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested activeMQArtemisScaledowns.
func (c *FakeActiveMQArtemisScaledowns) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(activemqartemisscaledownsResource, c.ns, opts))

}

// Create takes the representation of a activeMQArtemisScaledown and creates it.  Returns the server's representation of the activeMQArtemisScaledown, and an error, if there is any.
func (c *FakeActiveMQArtemisScaledowns) Create(activeMQArtemisScaledown *v2alpha1.ActiveMQArtemisScaledown) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(activemqartemisscaledownsResource, c.ns, activeMQArtemisScaledown), &v2alpha1.ActiveMQArtemisScaledown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ActiveMQArtemisScaledown), err
}

// Update takes the representation of a activeMQArtemisScaledown and updates it. Returns the server's representation of the activeMQArtemisScaledown, and an error, if there is any.
func (c *FakeActiveMQArtemisScaledowns) Update(activeMQArtemisScaledown *v2alpha1.ActiveMQArtemisScaledown) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(activemqartemisscaledownsResource, c.ns, activeMQArtemisScaledown), &v2alpha1.ActiveMQArtemisScaledown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ActiveMQArtemisScaledown), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeActiveMQArtemisScaledowns) UpdateStatus(activeMQArtemisScaledown *v2alpha1.ActiveMQArtemisScaledown) (*v2alpha1.ActiveMQArtemisScaledown, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(activemqartemisscaledownsResource, "status", c.ns, activeMQArtemisScaledown), &v2alpha1.ActiveMQArtemisScaledown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ActiveMQArtemisScaledown), err
}

// Delete takes name of the activeMQArtemisScaledown and deletes it. Returns an error if one occurs.
func (c *FakeActiveMQArtemisScaledowns) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(activemqartemisscaledownsResource, c.ns, name), &v2alpha1.ActiveMQArtemisScaledown{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeActiveMQArtemisScaledowns) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(activemqartemisscaledownsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2alpha1.ActiveMQArtemisScaledownList{})
	return err
}

// Patch applies the patch and returns the patched activeMQArtemisScaledown.
func (c *FakeActiveMQArtemisScaledowns) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2alpha1.ActiveMQArtemisScaledown, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(activemqartemisscaledownsResource, c.ns, name, data, subresources...), &v2alpha1.ActiveMQArtemisScaledown{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ActiveMQArtemisScaledown), err
}
