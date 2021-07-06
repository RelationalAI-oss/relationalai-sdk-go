/*
 * RelationalAI SDK
 *
 * This is a Client SDK for RelationalAI
 *
 * API version: 1.1.3
 * Contact: support@relational.ai
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SyntaxNode struct for SyntaxNode
type SyntaxNode struct {
	Type string `json:"type"`
}

// NewSyntaxNode instantiates a new SyntaxNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntaxNode(type_ string, ) *SyntaxNode {
	this := SyntaxNode{}
	this.Type = type_
	return &this
}

// NewSyntaxNodeWithDefaults instantiates a new SyntaxNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntaxNodeWithDefaults() *SyntaxNode {
	this := SyntaxNode{}
	var type_ string = ""
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SyntaxNode) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntaxNode) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntaxNode) SetType(v string) {
	o.Type = v
}

func (o SyntaxNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntaxNode struct {
	value *SyntaxNode
	isSet bool
}

func (v NullableSyntaxNode) Get() *SyntaxNode {
	return v.value
}

func (v *NullableSyntaxNode) Set(val *SyntaxNode) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntaxNode) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntaxNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntaxNode(val *SyntaxNode) *NullableSyntaxNode {
	return &NullableSyntaxNode{value: val, isSet: true}
}

func (v NullableSyntaxNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntaxNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


