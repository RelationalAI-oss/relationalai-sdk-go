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

// LoadDataActionAllOf struct for LoadDataActionAllOf
type LoadDataActionAllOf struct {
	Rel string `json:"rel"`
	Value LoadData `json:"value"`
}

// NewLoadDataActionAllOf instantiates a new LoadDataActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadDataActionAllOf(rel string, value LoadData, ) *LoadDataActionAllOf {
	this := LoadDataActionAllOf{}
	this.Rel = rel
	this.Value = value
	return &this
}

// NewLoadDataActionAllOfWithDefaults instantiates a new LoadDataActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadDataActionAllOfWithDefaults() *LoadDataActionAllOf {
	this := LoadDataActionAllOf{}
	var rel string = ""
	this.Rel = rel
	return &this
}

// GetRel returns the Rel field value
func (o *LoadDataActionAllOf) GetRel() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Rel
}

// GetRelOk returns a tuple with the Rel field value
// and a boolean to check if the value has been set.
func (o *LoadDataActionAllOf) GetRelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Rel, true
}

// SetRel sets field value
func (o *LoadDataActionAllOf) SetRel(v string) {
	o.Rel = v
}

// GetValue returns the Value field value
func (o *LoadDataActionAllOf) GetValue() LoadData {
	if o == nil  {
		var ret LoadData
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *LoadDataActionAllOf) GetValueOk() (*LoadData, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *LoadDataActionAllOf) SetValue(v LoadData) {
	o.Value = v
}

func (o LoadDataActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rel"] = o.Rel
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLoadDataActionAllOf struct {
	value *LoadDataActionAllOf
	isSet bool
}

func (v NullableLoadDataActionAllOf) Get() *LoadDataActionAllOf {
	return v.value
}

func (v *NullableLoadDataActionAllOf) Set(val *LoadDataActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadDataActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadDataActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadDataActionAllOf(val *LoadDataActionAllOf) *NullableLoadDataActionAllOf {
	return &NullableLoadDataActionAllOf{value: val, isSet: true}
}

func (v NullableLoadDataActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadDataActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


