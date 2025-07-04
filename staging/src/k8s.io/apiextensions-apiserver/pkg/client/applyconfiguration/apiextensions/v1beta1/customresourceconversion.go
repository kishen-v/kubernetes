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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

// CustomResourceConversionApplyConfiguration represents a declarative configuration of the CustomResourceConversion type for use
// with apply.
type CustomResourceConversionApplyConfiguration struct {
	Strategy                 *apiextensionsv1beta1.ConversionStrategyType `json:"strategy,omitempty"`
	WebhookClientConfig      *WebhookClientConfigApplyConfiguration       `json:"webhookClientConfig,omitempty"`
	ConversionReviewVersions []string                                     `json:"conversionReviewVersions,omitempty"`
}

// CustomResourceConversionApplyConfiguration constructs a declarative configuration of the CustomResourceConversion type for use with
// apply.
func CustomResourceConversion() *CustomResourceConversionApplyConfiguration {
	return &CustomResourceConversionApplyConfiguration{}
}
func (b CustomResourceConversionApplyConfiguration) IsApplyConfiguration() {}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *CustomResourceConversionApplyConfiguration) WithStrategy(value apiextensionsv1beta1.ConversionStrategyType) *CustomResourceConversionApplyConfiguration {
	b.Strategy = &value
	return b
}

// WithWebhookClientConfig sets the WebhookClientConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the WebhookClientConfig field is set to the value of the last call.
func (b *CustomResourceConversionApplyConfiguration) WithWebhookClientConfig(value *WebhookClientConfigApplyConfiguration) *CustomResourceConversionApplyConfiguration {
	b.WebhookClientConfig = value
	return b
}

// WithConversionReviewVersions adds the given value to the ConversionReviewVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ConversionReviewVersions field.
func (b *CustomResourceConversionApplyConfiguration) WithConversionReviewVersions(values ...string) *CustomResourceConversionApplyConfiguration {
	for i := range values {
		b.ConversionReviewVersions = append(b.ConversionReviewVersions, values[i])
	}
	return b
}
