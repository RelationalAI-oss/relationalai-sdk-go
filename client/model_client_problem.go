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

// ClientProblem struct for ClientProblem
type ClientProblem struct {
	AbstractProblem
	ErrorCode string `json:"error_code"`
	IsError bool `json:"is_error"`
	IsException bool `json:"is_exception"`
	Message string `json:"message"`
	Path string `json:"path"`
	Report string `json:"report"`
}

// NewClientProblem instantiates a new ClientProblem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientProblem(errorCode string, isError bool, isException bool, message string, path string, report string, ) *ClientProblem {
	this := ClientProblem{}
	this.ErrorCode = errorCode
	this.IsError = isError
	this.IsException = isException
	this.Message = message
	this.Path = path
	this.Report = report
	return &this
}

// NewClientProblemWithDefaults instantiates a new ClientProblem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientProblemWithDefaults() *ClientProblem {
	this := ClientProblem{}
	var errorCode string = ""
	this.ErrorCode = errorCode
	var isError bool = false
	this.IsError = isError
	var isException bool = false
	this.IsException = isException
	var message string = ""
	this.Message = message
	var path string = ""
	this.Path = path
	var report string = ""
	this.Report = report
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ClientProblem) GetErrorCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ClientProblem) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ClientProblem) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetIsError returns the IsError field value
func (o *ClientProblem) GetIsError() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsError
}

// GetIsErrorOk returns a tuple with the IsError field value
// and a boolean to check if the value has been set.
func (o *ClientProblem) GetIsErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsError, true
}

// SetIsError sets field value
func (o *ClientProblem) SetIsError(v bool) {
	o.IsError = v
}

// GetIsException returns the IsException field value
func (o *ClientProblem) GetIsException() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsException
}

// GetIsExceptionOk returns a tuple with the IsException field value
// and a boolean to check if the value has been set.
func (o *ClientProblem) GetIsExceptionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsException, true
}

// SetIsException sets field value
func (o *ClientProblem) SetIsException(v bool) {
	o.IsException = v
}

// GetMessage returns the Message field value
func (o *ClientProblem) GetMessage() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ClientProblem) GetMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ClientProblem) SetMessage(v string) {
	o.Message = v
}

// GetPath returns the Path field value
func (o *ClientProblem) GetPath() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ClientProblem) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ClientProblem) SetPath(v string) {
	o.Path = v
}

// GetReport returns the Report field value
func (o *ClientProblem) GetReport() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Report
}

// GetReportOk returns a tuple with the Report field value
// and a boolean to check if the value has been set.
func (o *ClientProblem) GetReportOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Report, true
}

// SetReport sets field value
func (o *ClientProblem) SetReport(v string) {
	o.Report = v
}

func (o ClientProblem) MarshalJSON() ([]byte, error) {
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
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["is_error"] = o.IsError
	}
	if true {
		toSerialize["is_exception"] = o.IsException
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["report"] = o.Report
	}
	return json.Marshal(toSerialize)
}

type NullableClientProblem struct {
	value *ClientProblem
	isSet bool
}

func (v NullableClientProblem) Get() *ClientProblem {
	return v.value
}

func (v *NullableClientProblem) Set(val *ClientProblem) {
	v.value = val
	v.isSet = true
}

func (v NullableClientProblem) IsSet() bool {
	return v.isSet
}

func (v *NullableClientProblem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientProblem(val *ClientProblem) *NullableClientProblem {
	return &NullableClientProblem{value: val, isSet: true}
}

func (v NullableClientProblem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientProblem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


