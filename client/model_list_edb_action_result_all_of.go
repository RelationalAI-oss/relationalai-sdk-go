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

// ListEdbActionResultAllOf struct for ListEdbActionResultAllOf
type ListEdbActionResultAllOf struct {
	Rels []RelKey `json:"rels,omitempty"`
}

// NewListEdbActionResultAllOf instantiates a new ListEdbActionResultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListEdbActionResultAllOf() *ListEdbActionResultAllOf {
	this := ListEdbActionResultAllOf{}
	return &this
}

// NewListEdbActionResultAllOfWithDefaults instantiates a new ListEdbActionResultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListEdbActionResultAllOfWithDefaults() *ListEdbActionResultAllOf {
	this := ListEdbActionResultAllOf{}
	return &this
}

// GetRels returns the Rels field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ListEdbActionResultAllOf) GetRels() []RelKey {
	if o == nil  {
		var ret []RelKey
		return ret
	}
	return o.Rels
}

// GetRelsOk returns a tuple with the Rels field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListEdbActionResultAllOf) GetRelsOk() (*[]RelKey, bool) {
	if o == nil || o.Rels == nil {
		return nil, false
	}
	return &o.Rels, true
}

// HasRels returns a boolean if a field has been set.
func (o *ListEdbActionResultAllOf) HasRels() bool {
	if o != nil && o.Rels != nil {
		return true
	}

	return false
}

// SetRels gets a reference to the given []RelKey and assigns it to the Rels field.
func (o *ListEdbActionResultAllOf) SetRels(v []RelKey) {
	o.Rels = v
}

func (o ListEdbActionResultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Rels != nil {
		toSerialize["rels"] = o.Rels
	}
	return json.Marshal(toSerialize)
}

type NullableListEdbActionResultAllOf struct {
	value *ListEdbActionResultAllOf
	isSet bool
}

func (v NullableListEdbActionResultAllOf) Get() *ListEdbActionResultAllOf {
	return v.value
}

func (v *NullableListEdbActionResultAllOf) Set(val *ListEdbActionResultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableListEdbActionResultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableListEdbActionResultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListEdbActionResultAllOf(val *ListEdbActionResultAllOf) *NullableListEdbActionResultAllOf {
	return &NullableListEdbActionResultAllOf{value: val, isSet: true}
}

func (v NullableListEdbActionResultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListEdbActionResultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

