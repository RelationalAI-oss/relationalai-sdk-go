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

// ParseActionResult struct for ParseActionResult
type ParseActionResult struct {
	ActionResult
	Problems []AbstractProblem `json:"problems,omitempty"`
}

// NewParseActionResult instantiates a new ParseActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParseActionResult() *ParseActionResult {
	this := ParseActionResult{}
	return &this
}

// NewParseActionResultWithDefaults instantiates a new ParseActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParseActionResultWithDefaults() *ParseActionResult {
	this := ParseActionResult{}
	return &this
}

// GetProblems returns the Problems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ParseActionResult) GetProblems() []AbstractProblem {
	if o == nil  {
		var ret []AbstractProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ParseActionResult) GetProblemsOk() (*[]AbstractProblem, bool) {
	if o == nil || o.Problems == nil {
		return nil, false
	}
	return &o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *ParseActionResult) HasProblems() bool {
	if o != nil && o.Problems != nil {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []AbstractProblem and assigns it to the Problems field.
func (o *ParseActionResult) SetProblems(v []AbstractProblem) {
	o.Problems = v
}

func (o ParseActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedActionResult, errActionResult := json.Marshal(o.ActionResult)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	errActionResult = json.Unmarshal([]byte(serializedActionResult), &toSerialize)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	if o.Problems != nil {
		toSerialize["problems"] = o.Problems
	}
	return json.Marshal(toSerialize)
}

type NullableParseActionResult struct {
	value *ParseActionResult
	isSet bool
}

func (v NullableParseActionResult) Get() *ParseActionResult {
	return v.value
}

func (v *NullableParseActionResult) Set(val *ParseActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableParseActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableParseActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParseActionResult(val *ParseActionResult) *NullableParseActionResult {
	return &NullableParseActionResult{value: val, isSet: true}
}

func (v NullableParseActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParseActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


