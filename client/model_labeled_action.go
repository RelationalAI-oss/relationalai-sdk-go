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

// LabeledAction struct for LabeledAction
type LabeledAction struct {
	Action interface{} `json:"action"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
}

// NewLabeledAction instantiates a new LabeledAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLabeledAction(action interface{}, name string, type_ string) *LabeledAction {
	this := LabeledAction{}
	this.Action = action
	this.Name = name
	this.Type = type_
	return &this
}

// NewLabeledActionWithDefaults instantiates a new LabeledAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLabeledActionWithDefaults() *LabeledAction {
	this := LabeledAction{}
	var name string = ""
	this.Name = name
	var type_ string = "LabeledAction"
	this.Type = type_
	return &this
}

// GetAction returns the Action field value
func (o *LabeledAction) GetAction() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *LabeledAction) GetActionOk() (*interface{}, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *LabeledAction) SetAction(v interface{}) {
	o.Action = v
}

// GetName returns the Name field value
func (o *LabeledAction) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *LabeledAction) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *LabeledAction) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *LabeledAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LabeledAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LabeledAction) SetType(v string) {
	o.Type = v
}

func (o LabeledAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["action"] = o.Action
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLabeledAction struct {
	value *LabeledAction
	isSet bool
}

func (v NullableLabeledAction) Get() *LabeledAction {
	return v.value
}

func (v *NullableLabeledAction) Set(val *LabeledAction) {
	v.value = val
	v.isSet = true
}

func (v NullableLabeledAction) IsSet() bool {
	return v.isSet
}

func (v *NullableLabeledAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLabeledAction(val *LabeledAction) *NullableLabeledAction {
	return &NullableLabeledAction{value: val, isSet: true}
}

func (v NullableLabeledAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLabeledAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}