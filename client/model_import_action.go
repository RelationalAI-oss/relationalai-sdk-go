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

// ImportAction struct for ImportAction
type ImportAction struct {
	Action
	Inputs []Relation `json:"inputs,omitempty"`
}

// NewImportAction instantiates a new ImportAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportAction() *ImportAction {
	this := ImportAction{}
	return &this
}

// NewImportActionWithDefaults instantiates a new ImportAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportActionWithDefaults() *ImportAction {
	this := ImportAction{}
	return &this
}

// GetInputs returns the Inputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ImportAction) GetInputs() []Relation {
	if o == nil  {
		var ret []Relation
		return ret
	}
	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ImportAction) GetInputsOk() (*[]Relation, bool) {
	if o == nil || o.Inputs == nil {
		return nil, false
	}
	return &o.Inputs, true
}

// HasInputs returns a boolean if a field has been set.
func (o *ImportAction) HasInputs() bool {
	if o != nil && o.Inputs != nil {
		return true
	}

	return false
}

// SetInputs gets a reference to the given []Relation and assigns it to the Inputs field.
func (o *ImportAction) SetInputs(v []Relation) {
	o.Inputs = v
}

func (o ImportAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return []byte{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return []byte{}, errAction
	}
	if o.Inputs != nil {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableImportAction struct {
	value *ImportAction
	isSet bool
}

func (v NullableImportAction) Get() *ImportAction {
	return v.value
}

func (v *NullableImportAction) Set(val *ImportAction) {
	v.value = val
	v.isSet = true
}

func (v NullableImportAction) IsSet() bool {
	return v.isSet
}

func (v *NullableImportAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportAction(val *ImportAction) *NullableImportAction {
	return &NullableImportAction{value: val, isSet: true}
}

func (v NullableImportAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


