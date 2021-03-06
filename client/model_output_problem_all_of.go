/*
 * RelationalAI SDK
 *
 * This is a Client SDK for RelationalAI
 *
 * API version: 1.2.0
 * Contact: support@relational.ai
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OutputProblemAllOf struct for OutputProblemAllOf
type OutputProblemAllOf struct {
	Exception string `json:"exception"`
	ExceptionStacktrace NullableString `json:"exception_stacktrace,omitempty"`
	Name string `json:"name"`
}

// NewOutputProblemAllOf instantiates a new OutputProblemAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputProblemAllOf(exception string, name string, ) *OutputProblemAllOf {
	this := OutputProblemAllOf{}
	this.Exception = exception
	this.Name = name
	return &this
}

// NewOutputProblemAllOfWithDefaults instantiates a new OutputProblemAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputProblemAllOfWithDefaults() *OutputProblemAllOf {
	this := OutputProblemAllOf{}
	var exception string = ""
	this.Exception = exception
	var name string = ""
	this.Name = name
	return &this
}

// GetException returns the Exception field value
func (o *OutputProblemAllOf) GetException() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Exception
}

// GetExceptionOk returns a tuple with the Exception field value
// and a boolean to check if the value has been set.
func (o *OutputProblemAllOf) GetExceptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Exception, true
}

// SetException sets field value
func (o *OutputProblemAllOf) SetException(v string) {
	o.Exception = v
}

// GetExceptionStacktrace returns the ExceptionStacktrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutputProblemAllOf) GetExceptionStacktrace() string {
	if o == nil || o.ExceptionStacktrace.Get() == nil {
		var ret string
		return ret
	}
	return *o.ExceptionStacktrace.Get()
}

// GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutputProblemAllOf) GetExceptionStacktraceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ExceptionStacktrace.Get(), o.ExceptionStacktrace.IsSet()
}

// HasExceptionStacktrace returns a boolean if a field has been set.
func (o *OutputProblemAllOf) HasExceptionStacktrace() bool {
	if o != nil && o.ExceptionStacktrace.IsSet() {
		return true
	}

	return false
}

// SetExceptionStacktrace gets a reference to the given NullableString and assigns it to the ExceptionStacktrace field.
func (o *OutputProblemAllOf) SetExceptionStacktrace(v string) {
	o.ExceptionStacktrace.Set(&v)
}
// SetExceptionStacktraceNil sets the value for ExceptionStacktrace to be an explicit nil
func (o *OutputProblemAllOf) SetExceptionStacktraceNil() {
	o.ExceptionStacktrace.Set(nil)
}

// UnsetExceptionStacktrace ensures that no value is present for ExceptionStacktrace, not even an explicit nil
func (o *OutputProblemAllOf) UnsetExceptionStacktrace() {
	o.ExceptionStacktrace.Unset()
}

// GetName returns the Name field value
func (o *OutputProblemAllOf) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *OutputProblemAllOf) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *OutputProblemAllOf) SetName(v string) {
	o.Name = v
}

func (o OutputProblemAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["exception"] = o.Exception
	}
	if o.ExceptionStacktrace.IsSet() {
		toSerialize["exception_stacktrace"] = o.ExceptionStacktrace.Get()
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableOutputProblemAllOf struct {
	value *OutputProblemAllOf
	isSet bool
}

func (v NullableOutputProblemAllOf) Get() *OutputProblemAllOf {
	return v.value
}

func (v *NullableOutputProblemAllOf) Set(val *OutputProblemAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputProblemAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputProblemAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputProblemAllOf(val *OutputProblemAllOf) *NullableOutputProblemAllOf {
	return &NullableOutputProblemAllOf{value: val, isSet: true}
}

func (v NullableOutputProblemAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputProblemAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


