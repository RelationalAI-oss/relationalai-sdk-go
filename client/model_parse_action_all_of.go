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

// ParseActionAllOf struct for ParseActionAllOf
type ParseActionAllOf struct {
	Nonterm string `json:"nonterm"`
	Source Source `json:"source"`
}

// NewParseActionAllOf instantiates a new ParseActionAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseActionAllOf(nonterm string, source Source, ) *ParseActionAllOf {
	this := ParseActionAllOf{}
	this.Nonterm = nonterm
	this.Source = source
	return &this
}

// NewParseActionAllOfWithDefaults instantiates a new ParseActionAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseActionAllOfWithDefaults() *ParseActionAllOf {
	this := ParseActionAllOf{}
	var nonterm string = ""
	this.Nonterm = nonterm
	return &this
}

// GetNonterm returns the Nonterm field value
func (o *ParseActionAllOf) GetNonterm() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Nonterm
}

// GetNontermOk returns a tuple with the Nonterm field value
// and a boolean to check if the value has been set.
func (o *ParseActionAllOf) GetNontermOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nonterm, true
}

// SetNonterm sets field value
func (o *ParseActionAllOf) SetNonterm(v string) {
	o.Nonterm = v
}

// GetSource returns the Source field value
func (o *ParseActionAllOf) GetSource() Source {
	if o == nil  {
		var ret Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ParseActionAllOf) GetSourceOk() (*Source, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ParseActionAllOf) SetSource(v Source) {
	o.Source = v
}

func (o ParseActionAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nonterm"] = o.Nonterm
	}
	if true {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableParseActionAllOf struct {
	value *ParseActionAllOf
	isSet bool
}

func (v NullableParseActionAllOf) Get() *ParseActionAllOf {
	return v.value
}

func (v *NullableParseActionAllOf) Set(val *ParseActionAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableParseActionAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableParseActionAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseActionAllOf(val *ParseActionAllOf) *NullableParseActionAllOf {
	return &NullableParseActionAllOf{value: val, isSet: true}
}

func (v NullableParseActionAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseActionAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


