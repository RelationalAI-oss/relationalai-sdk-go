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

// CardinalityAction struct for CardinalityAction
type CardinalityAction struct {
	Action
	Relname NullableString `json:"relname,omitempty"`
}

// NewCardinalityAction instantiates a new CardinalityAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardinalityAction() *CardinalityAction {
	this := CardinalityAction{}
	return &this
}

// NewCardinalityActionWithDefaults instantiates a new CardinalityAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardinalityActionWithDefaults() *CardinalityAction {
	this := CardinalityAction{}
	return &this
}

// GetRelname returns the Relname field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CardinalityAction) GetRelname() string {
	if o == nil || o.Relname.Get() == nil {
		var ret string
		return ret
	}
	return *o.Relname.Get()
}

// GetRelnameOk returns a tuple with the Relname field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CardinalityAction) GetRelnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Relname.Get(), o.Relname.IsSet()
}

// HasRelname returns a boolean if a field has been set.
func (o *CardinalityAction) HasRelname() bool {
	if o != nil && o.Relname.IsSet() {
		return true
	}

	return false
}

// SetRelname gets a reference to the given NullableString and assigns it to the Relname field.
func (o *CardinalityAction) SetRelname(v string) {
	o.Relname.Set(&v)
}
// SetRelnameNil sets the value for Relname to be an explicit nil
func (o *CardinalityAction) SetRelnameNil() {
	o.Relname.Set(nil)
}

// UnsetRelname ensures that no value is present for Relname, not even an explicit nil
func (o *CardinalityAction) UnsetRelname() {
	o.Relname.Unset()
}

func (o CardinalityAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return []byte{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return []byte{}, errAction
	}
	if o.Relname.IsSet() {
		toSerialize["relname"] = o.Relname.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableCardinalityAction struct {
	value *CardinalityAction
	isSet bool
}

func (v NullableCardinalityAction) Get() *CardinalityAction {
	return v.value
}

func (v *NullableCardinalityAction) Set(val *CardinalityAction) {
	v.value = val
	v.isSet = true
}

func (v NullableCardinalityAction) IsSet() bool {
	return v.isSet
}

func (v *NullableCardinalityAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardinalityAction(val *CardinalityAction) *NullableCardinalityAction {
	return &NullableCardinalityAction{value: val, isSet: true}
}

func (v NullableCardinalityAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardinalityAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


