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

// TokenAllOf struct for TokenAllOf
type TokenAllOf struct {
	Range Range `json:"range"`
	Value string `json:"value"`
}

// NewTokenAllOf instantiates a new TokenAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAllOf(range_ Range, value string, ) *TokenAllOf {
	this := TokenAllOf{}
	this.Range = range_
	this.Value = value
	return &this
}

// NewTokenAllOfWithDefaults instantiates a new TokenAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAllOfWithDefaults() *TokenAllOf {
	this := TokenAllOf{}
	var value string = ""
	this.Value = value
	return &this
}

// GetRange returns the Range field value
func (o *TokenAllOf) GetRange() Range {
	if o == nil  {
		var ret Range
		return ret
	}

	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *TokenAllOf) GetRangeOk() (*Range, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value
func (o *TokenAllOf) SetRange(v Range) {
	o.Range = v
}

// GetValue returns the Value field value
func (o *TokenAllOf) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *TokenAllOf) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *TokenAllOf) SetValue(v string) {
	o.Value = v
}

func (o TokenAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["range"] = o.Range
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableTokenAllOf struct {
	value *TokenAllOf
	isSet bool
}

func (v NullableTokenAllOf) Get() *TokenAllOf {
	return v.value
}

func (v *NullableTokenAllOf) Set(val *TokenAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAllOf(val *TokenAllOf) *NullableTokenAllOf {
	return &NullableTokenAllOf{value: val, isSet: true}
}

func (v NullableTokenAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


