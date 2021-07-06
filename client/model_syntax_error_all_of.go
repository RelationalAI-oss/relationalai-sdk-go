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

// SyntaxErrorAllOf struct for SyntaxErrorAllOf
type SyntaxErrorAllOf struct {
	Next *SyntaxNode `json:"next,omitempty"`
	Node SyntaxNode `json:"node"`
	Trace LinkedList `json:"trace"`
}

// NewSyntaxErrorAllOf instantiates a new SyntaxErrorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntaxErrorAllOf(node SyntaxNode, trace LinkedList, ) *SyntaxErrorAllOf {
	this := SyntaxErrorAllOf{}
	this.Node = node
	this.Trace = trace
	return &this
}

// NewSyntaxErrorAllOfWithDefaults instantiates a new SyntaxErrorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntaxErrorAllOfWithDefaults() *SyntaxErrorAllOf {
	this := SyntaxErrorAllOf{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *SyntaxErrorAllOf) GetNext() SyntaxNode {
	if o == nil || o.Next == nil {
		var ret SyntaxNode
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntaxErrorAllOf) GetNextOk() (*SyntaxNode, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *SyntaxErrorAllOf) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given SyntaxNode and assigns it to the Next field.
func (o *SyntaxErrorAllOf) SetNext(v SyntaxNode) {
	o.Next = &v
}

// GetNode returns the Node field value
func (o *SyntaxErrorAllOf) GetNode() SyntaxNode {
	if o == nil  {
		var ret SyntaxNode
		return ret
	}

	return o.Node
}

// GetNodeOk returns a tuple with the Node field value
// and a boolean to check if the value has been set.
func (o *SyntaxErrorAllOf) GetNodeOk() (*SyntaxNode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Node, true
}

// SetNode sets field value
func (o *SyntaxErrorAllOf) SetNode(v SyntaxNode) {
	o.Node = v
}

// GetTrace returns the Trace field value
func (o *SyntaxErrorAllOf) GetTrace() LinkedList {
	if o == nil  {
		var ret LinkedList
		return ret
	}

	return o.Trace
}

// GetTraceOk returns a tuple with the Trace field value
// and a boolean to check if the value has been set.
func (o *SyntaxErrorAllOf) GetTraceOk() (*LinkedList, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Trace, true
}

// SetTrace sets field value
func (o *SyntaxErrorAllOf) SetTrace(v LinkedList) {
	o.Trace = v
}

func (o SyntaxErrorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["node"] = o.Node
	}
	if true {
		toSerialize["trace"] = o.Trace
	}
	return json.Marshal(toSerialize)
}

type NullableSyntaxErrorAllOf struct {
	value *SyntaxErrorAllOf
	isSet bool
}

func (v NullableSyntaxErrorAllOf) Get() *SyntaxErrorAllOf {
	return v.value
}

func (v *NullableSyntaxErrorAllOf) Set(val *SyntaxErrorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntaxErrorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntaxErrorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntaxErrorAllOf(val *SyntaxErrorAllOf) *NullableSyntaxErrorAllOf {
	return &NullableSyntaxErrorAllOf{value: val, isSet: true}
}

func (v NullableSyntaxErrorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntaxErrorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


