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

// CollectProblemsAction struct for CollectProblemsAction
type CollectProblemsAction struct {
	Action
	Dummy NullableString `json:"dummy,omitempty"`
}

// NewCollectProblemsAction instantiates a new CollectProblemsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectProblemsAction() *CollectProblemsAction {
	this := CollectProblemsAction{}
	return &this
}

// NewCollectProblemsActionWithDefaults instantiates a new CollectProblemsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectProblemsActionWithDefaults() *CollectProblemsAction {
	this := CollectProblemsAction{}
	return &this
}

// GetDummy returns the Dummy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectProblemsAction) GetDummy() string {
	if o == nil || o.Dummy.Get() == nil {
		var ret string
		return ret
	}
	return *o.Dummy.Get()
}

// GetDummyOk returns a tuple with the Dummy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectProblemsAction) GetDummyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Dummy.Get(), o.Dummy.IsSet()
}

// HasDummy returns a boolean if a field has been set.
func (o *CollectProblemsAction) HasDummy() bool {
	if o != nil && o.Dummy.IsSet() {
		return true
	}

	return false
}

// SetDummy gets a reference to the given NullableString and assigns it to the Dummy field.
func (o *CollectProblemsAction) SetDummy(v string) {
	o.Dummy.Set(&v)
}
// SetDummyNil sets the value for Dummy to be an explicit nil
func (o *CollectProblemsAction) SetDummyNil() {
	o.Dummy.Set(nil)
}

// UnsetDummy ensures that no value is present for Dummy, not even an explicit nil
func (o *CollectProblemsAction) UnsetDummy() {
	o.Dummy.Unset()
}

func (o CollectProblemsAction) MarshalJSON() ([]byte, error) {
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

type NullableCollectProblemsAction struct {
	value *CollectProblemsAction
	isSet bool
}

func (v NullableCollectProblemsAction) Get() *CollectProblemsAction {
	return v.value
}

func (v *NullableCollectProblemsAction) Set(val *CollectProblemsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectProblemsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectProblemsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectProblemsAction(val *CollectProblemsAction) *NullableCollectProblemsAction {
	return &NullableCollectProblemsAction{value: val, isSet: true}
}

func (v NullableCollectProblemsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectProblemsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


