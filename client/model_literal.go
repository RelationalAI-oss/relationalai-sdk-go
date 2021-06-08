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

// Literal struct for Literal
type Literal struct {
	SyntaxNode
	Missing bool `json:"missing"`
	Range Range `json:"range"`
	Value string `json:"value"`
}

// NewLiteral instantiates a new Literal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLiteral(missing bool, range_ Range, value string, ) *Literal {
	this := Literal{}
	this.Missing = missing
	this.Range = range_
	this.Value = value
	return &this
}

// NewLiteralWithDefaults instantiates a new Literal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLiteralWithDefaults() *Literal {
	this := Literal{}
	var missing bool = false
	this.Missing = missing
	var value string = ""
	this.Value = value
	return &this
}

// GetMissing returns the Missing field value
func (o *Literal) GetMissing() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Missing
}

// GetMissingOk returns a tuple with the Missing field value
// and a boolean to check if the value has been set.
func (o *Literal) GetMissingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Missing, true
}

// SetMissing sets field value
func (o *Literal) SetMissing(v bool) {
	o.Missing = v
}

// GetRange returns the Range field value
func (o *Literal) GetRange() Range {
	if o == nil  {
		var ret Range
		return ret
	}

	return o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *Literal) GetRangeOk() (*Range, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Range, true
}

// SetRange sets field value
func (o *Literal) SetRange(v Range) {
	o.Range = v
}

// GetValue returns the Value field value
func (o *Literal) GetValue() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Literal) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Literal) SetValue(v string) {
	o.Value = v
}

func (o Literal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedSyntaxNode, errSyntaxNode := json.Marshal(o.SyntaxNode)
	if errSyntaxNode != nil {
		return []byte{}, errSyntaxNode
	}
	errSyntaxNode = json.Unmarshal([]byte(serializedSyntaxNode), &toSerialize)
	if errSyntaxNode != nil {
		return []byte{}, errSyntaxNode
	}
	if true {
		toSerialize["missing"] = o.Missing
	}
	if true {
		toSerialize["range"] = o.Range
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableLiteral struct {
	value *Literal
	isSet bool
}

func (v NullableLiteral) Get() *Literal {
	return v.value
}

func (v *NullableLiteral) Set(val *Literal) {
	v.value = val
	v.isSet = true
}

func (v NullableLiteral) IsSet() bool {
	return v.isSet
}

func (v *NullableLiteral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLiteral(val *Literal) *NullableLiteral {
	return &NullableLiteral{value: val, isSet: true}
}

func (v NullableLiteral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLiteral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


