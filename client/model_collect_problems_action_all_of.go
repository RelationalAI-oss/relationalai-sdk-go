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

// CollectProblemsActionAllOf struct for CollectProblemsActionAllOf
type CollectProblemsActionAllOf struct {
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewCollectProblemsActionAllOf instantiates a new CollectProblemsActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectProblemsActionAllOf() *CollectProblemsActionAllOf {
	this := CollectProblemsActionAllOf{}
	return &this
}

// NewCollectProblemsActionAllOfWithDefaults instantiates a new CollectProblemsActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectProblemsActionAllOfWithDefaults() *CollectProblemsActionAllOf {
	this := CollectProblemsActionAllOf{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectProblemsActionAllOf) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectProblemsActionAllOf) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *CollectProblemsActionAllOf) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *CollectProblemsActionAllOf) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *CollectProblemsActionAllOf) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *CollectProblemsActionAllOf) UnsetDummy() {
	o.Dummy.Unset()
}

func (o CollectProblemsActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dummy.IsSet() {
		toSerialize["dummy"] = o.Dummy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCollectProblemsActionAllOf struct {
	value *CollectProblemsActionAllOf
	isSet bool
}

func (v NullableCollectProblemsActionAllOf) Get() *CollectProblemsActionAllOf {
	return v.value
}

func (v *NullableCollectProblemsActionAllOf) Set(val *CollectProblemsActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectProblemsActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectProblemsActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectProblemsActionAllOf(val *CollectProblemsActionAllOf) *NullableCollectProblemsActionAllOf {
	return &NullableCollectProblemsActionAllOf{value: val, isSet: true}
}

func (v NullableCollectProblemsActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectProblemsActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


