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

// CollectProblemsActionResult struct for CollectProblemsActionResult
type CollectProblemsActionResult struct {
	ActionResult
	Problems []AbstractProblem `json:"problems,omitempty"`
}

// NewCollectProblemsActionResult instantiates a new CollectProblemsActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectProblemsActionResult() *CollectProblemsActionResult {
	this := CollectProblemsActionResult{}
	return &this
}

// NewCollectProblemsActionResultWithDefaults instantiates a new CollectProblemsActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectProblemsActionResultWithDefaults() *CollectProblemsActionResult {
	this := CollectProblemsActionResult{}
	return &this
}

// GetProblems returns the Problems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CollectProblemsActionResult) GetProblems() []AbstractProblem {
	if o == nil  {
		var ret []AbstractProblem
		return ret
	}
	return o.Problems
}

// GetProblemsOk returns a tuple with the Problems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CollectProblemsActionResult) GetProblemsOk() (*[]AbstractProblem, bool) {
	if o == nil || o.Problems == nil {
		return nil, false
	}
	return &o.Problems, true
}

// HasProblems returns a boolean if a field has been set.
func (o *CollectProblemsActionResult) HasProblems() bool {
	if o != nil && o.Problems != nil {
		return true
	}

	return false
}

// SetProblems gets a reference to the given []AbstractProblem and assigns it to the Problems field.
func (o *CollectProblemsActionResult) SetProblems(v []AbstractProblem) {
	o.Problems = v
}

func (o CollectProblemsActionResult) MarshalJSON() ([]byte, error) {
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

type NullableCollectProblemsActionResult struct {
	value *CollectProblemsActionResult
	isSet bool
}

func (v NullableCollectProblemsActionResult) Get() *CollectProblemsActionResult {
	return v.value
}

func (v *NullableCollectProblemsActionResult) Set(val *CollectProblemsActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectProblemsActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectProblemsActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectProblemsActionResult(val *CollectProblemsActionResult) *NullableCollectProblemsActionResult {
	return &NullableCollectProblemsActionResult{value: val, isSet: true}
}

func (v NullableCollectProblemsActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectProblemsActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


