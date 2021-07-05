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

// PersistProblem struct for PersistProblem
type PersistProblem struct {
	AbstractProblem
	Exception string `json:"exception"`
	ExceptionStacktrace NullableString `json:"exception_stacktrace,omitempty"`
}

// NewPersistProblem instantiates a new PersistProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersistProblem(exception string, ) *PersistProblem {
	this := PersistProblem{}
	this.Exception = exception
	return &this
}

// NewPersistProblemWithDefaults instantiates a new PersistProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersistProblemWithDefaults() *PersistProblem {
	this := PersistProblem{}
	var exception string = ""
	this.Exception = exception
	return &this
}

// GetException returns the Exception field value
func (o *PersistProblem) GetException() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value
// and a boolean to check if the value has been set.
func (o *PersistProblem) GetExceptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Exception, true
}

// SetException sets field value
func (o *PersistProblem) SetException(v string) {
	o.Exception = v
}

// GetExceptionStacktrace returns the ExceptionStacktrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PersistProblem) GetExceptionStacktrace() string {
	if o == nil || o.ExceptionStacktrace.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExceptionStacktrace.Get()
}

// GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PersistProblem) GetExceptionStacktraceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExceptionStacktrace.Get(), o.ExceptionStacktrace.IsSet()
}

// HasExceptionStacktrace returns a boolean if a field has been set.
func (o *PersistProblem) HasExceptionStacktrace() bool {
	if o != nil && o.ExceptionStacktrace.IsSet() {
		return true
	}

	return false
}

// SetExceptionStacktrace gets a reference to the given NullableString and assigns it to the ExceptionStacktrace field.
func (o *PersistProblem) SetExceptionStacktrace(v string) {
	o.ExceptionStacktrace.Set(&v)
}
// SetExceptionStacktraceNil sets the value for ExceptionStacktrace to be an explicit nil
func (o *PersistProblem) SetExceptionStacktraceNil() {
	o.ExceptionStacktrace.Set(nil)
}

// UnsetExceptionStacktrace ensures that no value is present for ExceptionStacktrace, not even an explicit nil
func (o *PersistProblem) UnsetExceptionStacktrace() {
	o.ExceptionStacktrace.Unset()
}

func (o PersistProblem) MarshalJSON() ([]byte, error) {
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
	if o.ExceptionStacktrace.IsSet() {
		toSerialize["exception_stacktrace"] = o.ExceptionStacktrace.Get()
	}
	return json.Marshal(toSerialize)
}

type NullablePersistProblem struct {
	value *PersistProblem
	isSet bool
}

func (v NullablePersistProblem) Get() *PersistProblem {
	return v.value
}

func (v *NullablePersistProblem) Set(val *PersistProblem) {
	v.value = val
	v.isSet = true
}

func (v NullablePersistProblem) IsSet() bool {
	return v.isSet
}

func (v *NullablePersistProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersistProblem(val *PersistProblem) *NullablePersistProblem {
	return &NullablePersistProblem{value: val, isSet: true}
}

func (v NullablePersistProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersistProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


