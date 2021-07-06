/*
 * RelationalAI SDK
 *
 * This is a Client SDK for RelationalAI
 *
 * API version: 1.1.3
 * Contact: support@relational.ai
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SetOptionsActionResult struct for SetOptionsActionResult
type SetOptionsActionResult struct {
	ActionResult
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewSetOptionsActionResult instantiates a new SetOptionsActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOptionsActionResult() *SetOptionsActionResult {
	this := SetOptionsActionResult{}
	return &this
}

// NewSetOptionsActionResultWithDefaults instantiates a new SetOptionsActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOptionsActionResultWithDefaults() *SetOptionsActionResult {
	this := SetOptionsActionResult{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetOptionsActionResult) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetOptionsActionResult) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *SetOptionsActionResult) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *SetOptionsActionResult) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *SetOptionsActionResult) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *SetOptionsActionResult) UnsetDummy() {
	o.Dummy.Unset()
}

func (o SetOptionsActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedActionResult, errActionResult := json.Marshal(o.ActionResult)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	errActionResult = json.Unmarshal([]byte(serializedActionResult), &toSerialize)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	if o.Dummy.IsSet() {
		toSerialize["dummy"] = o.Dummy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSetOptionsActionResult struct {
	value *SetOptionsActionResult
	isSet bool
}

func (v NullableSetOptionsActionResult) Get() *SetOptionsActionResult {
	return v.value
}

func (v *NullableSetOptionsActionResult) Set(val *SetOptionsActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOptionsActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOptionsActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOptionsActionResult(val *SetOptionsActionResult) *NullableSetOptionsActionResult {
	return &NullableSetOptionsActionResult{value: val, isSet: true}
}

func (v NullableSetOptionsActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOptionsActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


