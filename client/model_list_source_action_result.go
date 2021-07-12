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

// ListSourceActionResult struct for ListSourceActionResult
type ListSourceActionResult struct {
	ActionResult
	Sources []Source `json:"sources,omitempty"`
}

// NewListSourceActionResult instantiates a new ListSourceActionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListSourceActionResult() *ListSourceActionResult {
	this := ListSourceActionResult{}
	return &this
}

// NewListSourceActionResultWithDefaults instantiates a new ListSourceActionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListSourceActionResultWithDefaults() *ListSourceActionResult {
	this := ListSourceActionResult{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListSourceActionResult) GetSources() []Source {
	if o == nil  {
		var ret []Source
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListSourceActionResult) GetSourcesOk() (*[]Source, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *ListSourceActionResult) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []Source and assigns it to the Sources field.
func (o *ListSourceActionResult) SetSources(v []Source) {
	o.Sources = v
}

func (o ListSourceActionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedActionResult, errActionResult := json.Marshal(o.ActionResult)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	errActionResult = json.Unmarshal([]byte(serializedActionResult), &toSerialize)
	if errActionResult != nil {
		return []byte{}, errActionResult
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	return json.Marshal(toSerialize)
}

type NullableListSourceActionResult struct {
	value *ListSourceActionResult
	isSet bool
}

func (v NullableListSourceActionResult) Get() *ListSourceActionResult {
	return v.value
}

func (v *NullableListSourceActionResult) Set(val *ListSourceActionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableListSourceActionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableListSourceActionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListSourceActionResult(val *ListSourceActionResult) *NullableListSourceActionResult {
	return &NullableListSourceActionResult{value: val, isSet: true}
}

func (v NullableListSourceActionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListSourceActionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


