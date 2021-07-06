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

// CSVFileSchemaAllOf struct for CSVFileSchemaAllOf
type CSVFileSchemaAllOf struct {
	Types []string `json:"types,omitempty"`
}

// NewCSVFileSchemaAllOf instantiates a new CSVFileSchemaAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSVFileSchemaAllOf() *CSVFileSchemaAllOf {
	this := CSVFileSchemaAllOf{}
	return &this
}

// NewCSVFileSchemaAllOfWithDefaults instantiates a new CSVFileSchemaAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSVFileSchemaAllOfWithDefaults() *CSVFileSchemaAllOf {
	this := CSVFileSchemaAllOf{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CSVFileSchemaAllOf) GetTypes() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CSVFileSchemaAllOf) GetTypesOk() (*[]string, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return &o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *CSVFileSchemaAllOf) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []string and assigns it to the Types field.
func (o *CSVFileSchemaAllOf) SetTypes(v []string) {
	o.Types = v
}

func (o CSVFileSchemaAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableCSVFileSchemaAllOf struct {
	value *CSVFileSchemaAllOf
	isSet bool
}

func (v NullableCSVFileSchemaAllOf) Get() *CSVFileSchemaAllOf {
	return v.value
}

func (v *NullableCSVFileSchemaAllOf) Set(val *CSVFileSchemaAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCSVFileSchemaAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCSVFileSchemaAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSVFileSchemaAllOf(val *CSVFileSchemaAllOf) *NullableCSVFileSchemaAllOf {
	return &NullableCSVFileSchemaAllOf{value: val, isSet: true}
}

func (v NullableCSVFileSchemaAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSVFileSchemaAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


