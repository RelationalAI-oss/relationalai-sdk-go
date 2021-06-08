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

// ListSourceAction struct for ListSourceAction
type ListSourceAction struct {
	Action
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewListSourceAction instantiates a new ListSourceAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSourceAction() *ListSourceAction {
	this := ListSourceAction{}
	return &this
}

// NewListSourceActionWithDefaults instantiates a new ListSourceAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSourceActionWithDefaults() *ListSourceAction {
	this := ListSourceAction{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListSourceAction) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSourceAction) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *ListSourceAction) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *ListSourceAction) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *ListSourceAction) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *ListSourceAction) UnsetDummy() {
	o.Dummy.Unset()
}

func (o ListSourceAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return []byte{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return []byte{}, errAction
	}
	if o.Dummy.IsSet() {
		toSerialize["dummy"] = o.Dummy.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableListSourceAction struct {
	value *ListSourceAction
	isSet bool
}

func (v NullableListSourceAction) Get() *ListSourceAction {
	return v.value
}

func (v *NullableListSourceAction) Set(val *ListSourceAction) {
	v.value = val
	v.isSet = true
}

func (v NullableListSourceAction) IsSet() bool {
	return v.isSet
}

func (v *NullableListSourceAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSourceAction(val *ListSourceAction) *NullableListSourceAction {
	return &NullableListSourceAction{value: val, isSet: true}
}

func (v NullableListSourceAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSourceAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


