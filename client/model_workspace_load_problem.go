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

// WorkspaceLoadProblem struct for WorkspaceLoadProblem
type WorkspaceLoadProblem struct {
	AbstractProblem
	Exception string `json:"exception"`
}

// NewWorkspaceLoadProblem instantiates a new WorkspaceLoadProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceLoadProblem(exception string, ) *WorkspaceLoadProblem {
	this := WorkspaceLoadProblem{}
	this.Exception = exception
	return &this
}

// NewWorkspaceLoadProblemWithDefaults instantiates a new WorkspaceLoadProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceLoadProblemWithDefaults() *WorkspaceLoadProblem {
	this := WorkspaceLoadProblem{}
	var exception string = ""
	this.Exception = exception
	return &this
}

// GetException returns the Exception field value
func (o *WorkspaceLoadProblem) GetException() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value
// and a boolean to check if the value has been set.
func (o *WorkspaceLoadProblem) GetExceptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Exception, true
}

// SetException sets field value
func (o *WorkspaceLoadProblem) SetException(v string) {
	o.Exception = v
}

func (o WorkspaceLoadProblem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAbstractProblem, errAbstractProblem := json.Marshal(o.AbstractProblem)
	if errAbstractProblem != nil {
		return []byte{}, errAbstractProblem
	}
	errAbstractProblem = json.Unmarshal([]byte(serializedAbstractProblem), &toSerialize)
	if errAbstractProblem != nil {
		return []byte{}, errAbstractProblem
	}
	if true {
		toSerialize["exception"] = o.Exception
	}
	return json.Marshal(toSerialize)
}

type NullableWorkspaceLoadProblem struct {
	value *WorkspaceLoadProblem
	isSet bool
}

func (v NullableWorkspaceLoadProblem) Get() *WorkspaceLoadProblem {
	return v.value
}

func (v *NullableWorkspaceLoadProblem) Set(val *WorkspaceLoadProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceLoadProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceLoadProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceLoadProblem(val *WorkspaceLoadProblem) *NullableWorkspaceLoadProblem {
	return &NullableWorkspaceLoadProblem{value: val, isSet: true}
}

func (v NullableWorkspaceLoadProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceLoadProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


