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

// JSONFileSyntax struct for JSONFileSyntax
type JSONFileSyntax struct {
	FileSyntax
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewJSONFileSyntax instantiates a new JSONFileSyntax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJSONFileSyntax() *JSONFileSyntax {
	this := JSONFileSyntax{}
	return &this
}

// NewJSONFileSyntaxWithDefaults instantiates a new JSONFileSyntax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJSONFileSyntaxWithDefaults() *JSONFileSyntax {
	this := JSONFileSyntax{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *JSONFileSyntax) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JSONFileSyntax) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *JSONFileSyntax) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *JSONFileSyntax) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *JSONFileSyntax) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *JSONFileSyntax) UnsetDummy() {
	o.Dummy.Unset()
}

func (o JSONFileSyntax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedFileSyntax, errFileSyntax := json.Marshal(o.FileSyntax)
	if errFileSyntax != nil {
		return []byte{}, errFileSyntax
	}
	errFileSyntax = json.Unmarshal([]byte(serializedFileSyntax), &toSerialize)
	if errFileSyntax != nil {
		return []byte{}, errFileSyntax
	}
	if o.Dummy.IsSet() {
		toSerialize["dummy"] = o.Dummy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableJSONFileSyntax struct {
	value *JSONFileSyntax
	isSet bool
}

func (v NullableJSONFileSyntax) Get() *JSONFileSyntax {
	return v.value
}

func (v *NullableJSONFileSyntax) Set(val *JSONFileSyntax) {
	v.value = val
	v.isSet = true
}

func (v NullableJSONFileSyntax) IsSet() bool {
	return v.isSet
}

func (v *NullableJSONFileSyntax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJSONFileSyntax(val *JSONFileSyntax) *NullableJSONFileSyntax {
	return &NullableJSONFileSyntax{value: val, isSet: true}
}

func (v NullableJSONFileSyntax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJSONFileSyntax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


