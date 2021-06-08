/*
 * Delve Client SDK
 *
 * This is a Client SDK for Delve API
 *
 * API version: 1.1.3
 * Contact: support@relational.ai
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Nil struct for Nil
type Nil struct {
	LinkedList
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewNil instantiates a new Nil object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNil() *Nil {
	this := Nil{}
	return &this
}

// NewNilWithDefaults instantiates a new Nil object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNilWithDefaults() *Nil {
	this := Nil{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Nil) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Nil) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *Nil) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *Nil) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *Nil) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *Nil) UnsetDummy() {
	o.Dummy.Unset()
}

func (o Nil) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedLinkedList, errLinkedList := json.Marshal(o.LinkedList)
	if errLinkedList != nil {
		return []byte{}, errLinkedList
	}
	errLinkedList = json.Unmarshal([]byte(serializedLinkedList), &toSerialize)
	if errLinkedList != nil {
		return []byte{}, errLinkedList
	}
	if o.Dummy.IsSet() {
		toSerialize["dummy"] = o.Dummy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableNil struct {
	value *Nil
	isSet bool
}

func (v NullableNil) Get() *Nil {
	return v.value
}

func (v *NullableNil) Set(val *Nil) {
	v.value = val
	v.isSet = true
}

func (v NullableNil) IsSet() bool {
	return v.isSet
}

func (v *NullableNil) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNil(val *Nil) *NullableNil {
	return &NullableNil{value: val, isSet: true}
}

func (v NullableNil) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNil) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


