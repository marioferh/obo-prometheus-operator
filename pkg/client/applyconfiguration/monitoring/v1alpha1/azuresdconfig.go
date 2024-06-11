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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	monitoringv1 "github.com/rhobs/obo-prometheus-operator/pkg/apis/monitoring/v1"
	v1 "k8s.io/api/core/v1"
)

// AzureSDConfigApplyConfiguration represents an declarative configuration of the AzureSDConfig type for use
// with apply.
type AzureSDConfigApplyConfiguration struct {
	Environment          *string                `json:"environment,omitempty"`
	AuthenticationMethod *string                `json:"authenticationMethod,omitempty"`
	SubscriptionID       *string                `json:"subscriptionID,omitempty"`
	TenantID             *string                `json:"tenantID,omitempty"`
	ClientID             *string                `json:"clientID,omitempty"`
	ClientSecret         *v1.SecretKeySelector  `json:"clientSecret,omitempty"`
	ResourceGroup        *string                `json:"resourceGroup,omitempty"`
	RefreshInterval      *monitoringv1.Duration `json:"refreshInterval,omitempty"`
	Port                 *int                   `json:"port,omitempty"`
}

// AzureSDConfigApplyConfiguration constructs an declarative configuration of the AzureSDConfig type for use with
// apply.
func AzureSDConfig() *AzureSDConfigApplyConfiguration {
	return &AzureSDConfigApplyConfiguration{}
}

// WithEnvironment sets the Environment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Environment field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithEnvironment(value string) *AzureSDConfigApplyConfiguration {
	b.Environment = &value
	return b
}

// WithAuthenticationMethod sets the AuthenticationMethod field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AuthenticationMethod field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithAuthenticationMethod(value string) *AzureSDConfigApplyConfiguration {
	b.AuthenticationMethod = &value
	return b
}

// WithSubscriptionID sets the SubscriptionID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SubscriptionID field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithSubscriptionID(value string) *AzureSDConfigApplyConfiguration {
	b.SubscriptionID = &value
	return b
}

// WithTenantID sets the TenantID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TenantID field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithTenantID(value string) *AzureSDConfigApplyConfiguration {
	b.TenantID = &value
	return b
}

// WithClientID sets the ClientID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientID field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithClientID(value string) *AzureSDConfigApplyConfiguration {
	b.ClientID = &value
	return b
}

// WithClientSecret sets the ClientSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClientSecret field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithClientSecret(value v1.SecretKeySelector) *AzureSDConfigApplyConfiguration {
	b.ClientSecret = &value
	return b
}

// WithResourceGroup sets the ResourceGroup field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceGroup field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithResourceGroup(value string) *AzureSDConfigApplyConfiguration {
	b.ResourceGroup = &value
	return b
}

// WithRefreshInterval sets the RefreshInterval field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RefreshInterval field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithRefreshInterval(value monitoringv1.Duration) *AzureSDConfigApplyConfiguration {
	b.RefreshInterval = &value
	return b
}

// WithPort sets the Port field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Port field is set to the value of the last call.
func (b *AzureSDConfigApplyConfiguration) WithPort(value int) *AzureSDConfigApplyConfiguration {
	b.Port = &value
	return b
}
