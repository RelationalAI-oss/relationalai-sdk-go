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

// LoadDataAction struct for LoadDataAction
type LoadDataAction struct {
	Action
	Rel string `json:"rel"`
	Value LoadData `json:"value"`
}

// NewLoadDataAction instantiates a new LoadDataAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadDataAction(rel string, value LoadData, ) *LoadDataAction {
	this := LoadDataAction{}
	this.Rel = rel
	this.Value = value
	return &this
}

// NewLoadDataActionWithDefaults instantiates a new LoadDataAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadDataActionWithDefaults() *LoadDataAction {
	this := LoadDataAction{}
	var rel string = ""
	this.Rel = rel
	return &this
}

// GetRel returns the Rel field value
func (o *LoadDataAction) GetRel() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Rel
}

// GetRelOk returns a tuple with the Rel field value
// and a boolean to check if the value has been set.
func (o *LoadDataAction) GetRelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rel, true
}

// SetRel sets field value
func (o *LoadDataAction) SetRel(v string) {
	o.Rel = v
}

// GetValue returns the Value field value
func (o *LoadDataAction) GetValue() LoadData {
	if o == nil  {
		var ret LoadData
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LoadDataAction) GetValueOk() (*LoadData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LoadDataAction) SetValue(v LoadData) {
	o.Value = v
}

func (o LoadDataAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return []byte{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return []byte{}, errAction
	}
	if true {
		toSerialize["rel"] = o.Rel
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLoadDataAction struct {
	value *LoadDataAction
	isSet bool
}

func (v NullableLoadDataAction) Get() *LoadDataAction {
	return v.value
}

func (v *NullableLoadDataAction) Set(val *LoadDataAction) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadDataAction) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadDataAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadDataAction(val *LoadDataAction) *NullableLoadDataAction {
	return &NullableLoadDataAction{value: val, isSet: true}
}

func (v NullableLoadDataAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadDataAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


