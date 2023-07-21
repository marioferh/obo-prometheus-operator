// Copyright The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1alpha1"
	monitoringv1alpha1 "github.com/rhobs/obo-prometheus-operator/pkg/client/applyconfiguration/monitoring/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScrapeConfigs implements ScrapeConfigInterface
type FakeScrapeConfigs struct {
	Fake *FakeMonitoringV1alpha1
	ns   string
}

var scrapeconfigsResource = schema.GroupVersionResource{Group: "monitoring.rhobs", Version: "v1alpha1", Resource: "scrapeconfigs"}

var scrapeconfigsKind = schema.GroupVersionKind{Group: "monitoring.rhobs", Version: "v1alpha1", Kind: "ScrapeConfig"}

// Get takes name of the scrapeConfig, and returns the corresponding scrapeConfig object, and an error if there is any.
func (c *FakeScrapeConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScrapeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scrapeconfigsResource, c.ns, name), &v1alpha1.ScrapeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScrapeConfig), err
}

// List takes label and field selectors, and returns the list of ScrapeConfigs that match those selectors.
func (c *FakeScrapeConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScrapeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scrapeconfigsResource, scrapeconfigsKind, c.ns, opts), &v1alpha1.ScrapeConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScrapeConfigList{ListMeta: obj.(*v1alpha1.ScrapeConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScrapeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scrapeConfigs.
func (c *FakeScrapeConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scrapeconfigsResource, c.ns, opts))

}

// Create takes the representation of a scrapeConfig and creates it.  Returns the server's representation of the scrapeConfig, and an error, if there is any.
func (c *FakeScrapeConfigs) Create(ctx context.Context, scrapeConfig *v1alpha1.ScrapeConfig, opts v1.CreateOptions) (result *v1alpha1.ScrapeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scrapeconfigsResource, c.ns, scrapeConfig), &v1alpha1.ScrapeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScrapeConfig), err
}

// Update takes the representation of a scrapeConfig and updates it. Returns the server's representation of the scrapeConfig, and an error, if there is any.
func (c *FakeScrapeConfigs) Update(ctx context.Context, scrapeConfig *v1alpha1.ScrapeConfig, opts v1.UpdateOptions) (result *v1alpha1.ScrapeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scrapeconfigsResource, c.ns, scrapeConfig), &v1alpha1.ScrapeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScrapeConfig), err
}

// Delete takes name of the scrapeConfig and deletes it. Returns an error if one occurs.
func (c *FakeScrapeConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scrapeconfigsResource, c.ns, name, opts), &v1alpha1.ScrapeConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScrapeConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scrapeconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScrapeConfigList{})
	return err
}

// Patch applies the patch and returns the patched scrapeConfig.
func (c *FakeScrapeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScrapeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scrapeconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScrapeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScrapeConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied scrapeConfig.
func (c *FakeScrapeConfigs) Apply(ctx context.Context, scrapeConfig *monitoringv1alpha1.ScrapeConfigApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.ScrapeConfig, err error) {
	if scrapeConfig == nil {
		return nil, fmt.Errorf("scrapeConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(scrapeConfig)
	if err != nil {
		return nil, err
	}
	name := scrapeConfig.Name
	if name == nil {
		return nil, fmt.Errorf("scrapeConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scrapeconfigsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.ScrapeConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScrapeConfig), err
}
