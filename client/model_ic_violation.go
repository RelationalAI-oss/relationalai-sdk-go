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

// ICViolation struct for ICViolation
type ICViolation struct {
	RelKey RelKey `json:"rel_key"`
	Source string `json:"source"`
	Type string `json:"type"`
}

// NewICViolation instantiates a new ICViolation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewICViolation(relKey RelKey, source string, type_ string, ) *ICViolation {
	this := ICViolation{}
	this.RelKey = relKey
	this.Source = source
	this.Type = type_
	return &this
}

// NewICViolationWithDefaults instantiates a new ICViolation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewICViolationWithDefaults() *ICViolation {
	this := ICViolation{}
	var source string = ""
	this.Source = source
	var type_ string = "ICViolation"
	this.Type = type_
	return &this
}

// GetRelKey returns the RelKey field value
func (o *ICViolation) GetRelKey() RelKey {
	if o == nil  {
		var ret RelKey
		return ret
	}

	return o.RelKey
}

// GetRelKeyOk returns a tuple with the RelKey field value
// and a boolean to check if the value has been set.
func (o *ICViolation) GetRelKeyOk() (*RelKey, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RelKey, true
}

// SetRelKey sets field value
func (o *ICViolation) SetRelKey(v RelKey) {
	o.RelKey = v
}

// GetSource returns the Source field value
func (o *ICViolation) GetSource() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ICViolation) GetSourceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ICViolation) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *ICViolation) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ICViolation) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ICViolation) SetType(v string) {
	o.Type = v
}

func (o ICViolation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rel_key"] = o.RelKey
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableICViolation struct {
	value *ICViolation
	isSet bool
}

func (v NullableICViolation) Get() *ICViolation {
	return v.value
}

func (v *NullableICViolation) Set(val *ICViolation) {
	v.value = val
	v.isSet = true
}

func (v NullableICViolation) IsSet() bool {
	return v.isSet
}

func (v *NullableICViolation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableICViolation(val *ICViolation) *NullableICViolation {
	return &NullableICViolation{value: val, isSet: true}
}

func (v NullableICViolation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableICViolation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


