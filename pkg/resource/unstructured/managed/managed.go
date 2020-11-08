/*
Copyright 2020 The Crossplane Authors.

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

// Package managed contains an unstructured managed resource.
package managed

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"github.com/crossplane/crossplane-runtime/pkg/fieldpath"
)

// An Option modifies an unstructured managed resource.
type Option func(*Unstructured)

// WithGroupVersionKind sets the GroupVersionKind of the unstructured managed
// resource.
func WithGroupVersionKind(gvk schema.GroupVersionKind) Option {
	return func(c *Unstructured) {
		c.SetGroupVersionKind(gvk)
	}
}

// WithConditions returns an Option that sets the supplied conditions on an
// unstructured managed resource.
func WithConditions(c ...v1alpha1.Condition) Option {
	return func(cr *Unstructured) {
		cr.SetConditions(c...)
	}
}

// New returns a new unstructured managed resource.
func New(opts ...Option) *Unstructured {
	c := &Unstructured{unstructured.Unstructured{Object: make(map[string]interface{})}}
	for _, f := range opts {
		f(c)
	}
	return c
}

// An Unstructured managed resource.
type Unstructured struct {
	unstructured.Unstructured
}

// GetUnstructured returns the underlying *unstructured.Unstructured.
func (c *Unstructured) GetUnstructured() *unstructured.Unstructured {
	return &c.Unstructured
}

// GetWriteConnectionSecretToReference of this managed resource.
func (c *Unstructured) GetWriteConnectionSecretToReference() *v1alpha1.SecretReference {
	out := &v1alpha1.SecretReference{}
	if err := fieldpath.Pave(c.Object).GetValueInto("spec.writeConnectionSecretToRef", out); err != nil {
		return nil
	}
	return out
}

// SetWriteConnectionSecretToReference of this managed resource.
func (c *Unstructured) SetWriteConnectionSecretToReference(ref *v1alpha1.SecretReference) {
	_ = fieldpath.Pave(c.Object).SetValue("spec.writeConnectionSecretToRef", ref)
}

// GetCondition of this managed resource.
func (c *Unstructured) GetCondition(ct v1alpha1.ConditionType) v1alpha1.Condition {
	conditioned := v1alpha1.ConditionedStatus{}
	// The path is directly `status` because conditions are inline.
	if err := fieldpath.Pave(c.Object).GetValueInto("status", &conditioned); err != nil {
		return v1alpha1.Condition{}
	}
	return conditioned.GetCondition(ct)
}

// SetConditions of this managed resource.
func (c *Unstructured) SetConditions(conditions ...v1alpha1.Condition) {
	conditioned := v1alpha1.ConditionedStatus{}
	// The path is directly `status` because conditions are inline.
	_ = fieldpath.Pave(c.Object).GetValueInto("status", &conditioned)
	conditioned.SetConditions(conditions...)
	_ = fieldpath.Pave(c.Object).SetValue("status.conditions", conditioned.Conditions)
}

// GetDeletionPolicy of this managed resource.
func (c *Unstructured) GetDeletionPolicy() v1alpha1.DeletionPolicy {
	out := ""
	if err := fieldpath.Pave(c.Object).GetValueInto("spec.writeConnectionSecretToRef", out); err != nil {
		return v1alpha1.DeletionPolicy("")
	}
	return v1alpha1.DeletionPolicy(out)
}

// SetDeletionPolicy of this managed resource.
func (c *Unstructured) SetDeletionPolicy(r v1alpha1.DeletionPolicy) {
	_ = fieldpath.Pave(c.Object).SetValue("spec.deletionPolicy", r)
}

// GetProviderConfigReference of this  managed resource.
func (c *Unstructured) GetProviderConfigReference() *v1alpha1.Reference {
	out := &v1alpha1.Reference{}
	if err := fieldpath.Pave(c.Object).GetValueInto("spec.providerConfigRef", out); err != nil {
		return nil
	}
	return out
}

// SetProviderConfigReference of this  managed resource.
func (c *Unstructured) SetProviderConfigReference(r *v1alpha1.Reference) {
	_ = fieldpath.Pave(c.Object).SetValue("spec.providerConfigRef", r)
}

/*
GetProviderReference of this  managed resource.
Deprecated: Use GetProviderConfigReference.
*/
func (c *Unstructured) GetProviderReference() *v1alpha1.Reference {
	out := &v1alpha1.Reference{}
	if err := fieldpath.Pave(c.Object).GetValueInto("spec.providerRef", out); err != nil {
		return nil
	}
	return out
}

/*
SetProviderReference of this GKECluster.
Deprecated: Use SetProviderConfigReference.
*/
func (c *Unstructured) SetProviderReference(r *v1alpha1.Reference) {
	_ = fieldpath.Pave(c.Object).SetValue("spec.providerRef", r)
}
