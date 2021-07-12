/*
 * RelationalAI SDK
 *
 * This is a Client SDK for RelationalAI
 *
 * API version: 1.2.0
 * Contact: support@relational.ai
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// DefaultIntegration struct for DefaultIntegration
type DefaultIntegration struct {
	Integration
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewDefaultIntegration instantiates a new DefaultIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultIntegration() *DefaultIntegration {
	this := DefaultIntegration{}
	return &this
}

// NewDefaultIntegrationWithDefaults instantiates a new DefaultIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultIntegrationWithDefaults() *DefaultIntegration {
	this := DefaultIntegration{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DefaultIntegration) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DefaultIntegration) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *DefaultIntegration) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *DefaultIntegration) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *DefaultIntegration) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *DefaultIntegration) UnsetDummy() {
	o.Dummy.Unset()
}

func (o DefaultIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedIntegration, errIntegration := json.Marshal(o.Integration)
	if errIntegration != nil {
		return []byte{}, errIntegration
	}
	errIntegration = json.Unmarshal([]byte(serializedIntegration), &toSerialize)
	if errIntegration != nil {
		return []byte{}, errIntegration
	}
	if o.Dummy.IsSet() {
		toSerialize["dummy"] = o.Dummy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableDefaultIntegration struct {
	value *DefaultIntegration
	isSet bool
}

func (v NullableDefaultIntegration) Get() *DefaultIntegration {
	return v.value
}

func (v *NullableDefaultIntegration) Set(val *DefaultIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultIntegration(val *DefaultIntegration) *NullableDefaultIntegration {
	return &NullableDefaultIntegration{value: val, isSet: true}
}

func (v NullableDefaultIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


