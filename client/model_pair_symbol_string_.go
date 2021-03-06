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

// PairSymbolString struct for PairSymbolString
type PairSymbolString struct {
	First string `json:"first"`
	Second string `json:"second"`
	Type string `json:"type"`
}

// NewPairSymbolString instantiates a new PairSymbolString object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPairSymbolString(first string, second string, type_ string, ) *PairSymbolString {
	this := PairSymbolString{}
	this.First = first
	this.Second = second
	this.Type = type_
	return &this
}

// NewPairSymbolStringWithDefaults instantiates a new PairSymbolString object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPairSymbolStringWithDefaults() *PairSymbolString {
	this := PairSymbolString{}
	var first string = ""
	this.First = first
	var second string = ""
	this.Second = second
	var type_ string = "Pair_Symbol_String_"
	this.Type = type_
	return &this
}

// GetFirst returns the First field value
func (o *PairSymbolString) GetFirst() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *PairSymbolString) GetFirstOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value
func (o *PairSymbolString) SetFirst(v string) {
	o.First = v
}

// GetSecond returns the Second field value
func (o *PairSymbolString) GetSecond() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Second
}

// GetSecondOk returns a tuple with the Second field value
// and a boolean to check if the value has been set.
func (o *PairSymbolString) GetSecondOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Second, true
}

// SetSecond sets field value
func (o *PairSymbolString) SetSecond(v string) {
	o.Second = v
}

// GetType returns the Type field value
func (o *PairSymbolString) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PairSymbolString) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PairSymbolString) SetType(v string) {
	o.Type = v
}

func (o PairSymbolString) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["first"] = o.First
	}
	if true {
		toSerialize["second"] = o.Second
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullablePairSymbolString struct {
	value *PairSymbolString
	isSet bool
}

func (v NullablePairSymbolString) Get() *PairSymbolString {
	return v.value
}

func (v *NullablePairSymbolString) Set(val *PairSymbolString) {
	v.value = val
	v.isSet = true
}

func (v NullablePairSymbolString) IsSet() bool {
	return v.isSet
}

func (v *NullablePairSymbolString) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePairSymbolString(val *PairSymbolString) *NullablePairSymbolString {
	return &NullablePairSymbolString{value: val, isSet: true}
}

func (v NullablePairSymbolString) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePairSymbolString) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


