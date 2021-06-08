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

// FileSyntax struct for FileSyntax
type FileSyntax struct {
	Type string `json:"type"`
}

// NewFileSyntax instantiates a new FileSyntax object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSyntax(type_ string, ) *FileSyntax {
	this := FileSyntax{}
	this.Type = type_
	return &this
}

// NewFileSyntaxWithDefaults instantiates a new FileSyntax object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSyntaxWithDefaults() *FileSyntax {
	this := FileSyntax{}
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *FileSyntax) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *FileSyntax) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *FileSyntax) SetType(v string) {
	o.Type = v
}

func (o FileSyntax) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFileSyntax struct {
	value *FileSyntax
	isSet bool
}

func (v NullableFileSyntax) Get() *FileSyntax {
	return v.value
}

func (v *NullableFileSyntax) Set(val *FileSyntax) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSyntax) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSyntax) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSyntax(val *FileSyntax) *NullableFileSyntax {
	return &NullableFileSyntax{value: val, isSet: true}
}

func (v NullableFileSyntax) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSyntax) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


