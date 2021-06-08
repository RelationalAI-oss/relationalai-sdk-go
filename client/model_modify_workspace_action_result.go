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

// ModifyWorkspaceActionResult struct for ModifyWorkspaceActionResult
type ModifyWorkspaceActionResult struct {
	ActionResult
	DeleteEdbResult []RelKey `json:"delete_edb_result,omitempty"`
}

// NewModifyWorkspaceActionResult instantiates a new ModifyWorkspaceActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModifyWorkspaceActionResult() *ModifyWorkspaceActionResult {
	this := ModifyWorkspaceActionResult{}
	return &this
}

// NewModifyWorkspaceActionResultWithDefaults instantiates a new ModifyWorkspaceActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModifyWorkspaceActionResultWithDefaults() *ModifyWorkspaceActionResult {
	this := ModifyWorkspaceActionResult{}
	return &this
}

// GetDeleteEdbResult returns the DeleteEdbResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModifyWorkspaceActionResult) GetDeleteEdbResult() []RelKey {
	if o == nil  {
		var ret []RelKey
		return ret
	}
	return o.DeleteEdbResult
}

// GetDeleteEdbResultOk returns a tuple with the DeleteEdbResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModifyWorkspaceActionResult) GetDeleteEdbResultOk() (*[]RelKey, bool) {
	if o == nil || o.DeleteEdbResult == nil {
		return nil, false
	}
	return &o.DeleteEdbResult, true
}

// HasDeleteEdbResult returns a boolean if a field has been set.
func (o *ModifyWorkspaceActionResult) HasDeleteEdbResult() bool {
	if o != nil && o.DeleteEdbResult != nil {
		return true
	}

	return false
}

// SetDeleteEdbResult gets a reference to the given []RelKey and assigns it to the DeleteEdbResult field.
func (o *ModifyWorkspaceActionResult) SetDeleteEdbResult(v []RelKey) {
	o.DeleteEdbResult = v
}

func (o ModifyWorkspaceActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedActionResult, errActionResult := json.Marshal(o.ActionResult)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	errActionResult = json.Unmarshal([]byte(serializedActionResult), &toSerialize)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	if o.DeleteEdbResult != nil {
		toSerialize["delete_edb_result"] = o.DeleteEdbResult
	}
	return json.Marshal(toSerialize)
}

type NullableModifyWorkspaceActionResult struct {
	value *ModifyWorkspaceActionResult
	isSet bool
}

func (v NullableModifyWorkspaceActionResult) Get() *ModifyWorkspaceActionResult {
	return v.value
}

func (v *NullableModifyWorkspaceActionResult) Set(val *ModifyWorkspaceActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModifyWorkspaceActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModifyWorkspaceActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModifyWorkspaceActionResult(val *ModifyWorkspaceActionResult) *NullableModifyWorkspaceActionResult {
	return &NullableModifyWorkspaceActionResult{value: val, isSet: true}
}

func (v NullableModifyWorkspaceActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModifyWorkspaceActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


