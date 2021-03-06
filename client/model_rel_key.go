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

// RelKey struct for RelKey
type RelKey struct {
	Keys []string `json:"keys,omitempty"`
	Name string `json:"name"`
	Values []string `json:"values,omitempty"`
	Type string `json:"type"`
}

// NewRelKey instantiates a new RelKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRelKey(name string, type_ string, ) *RelKey {
	this := RelKey{}
	this.Name = name
	this.Type = type_
	return &this
}

// NewRelKeyWithDefaults instantiates a new RelKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRelKeyWithDefaults() *RelKey {
	this := RelKey{}
	var name string = ""
	this.Name = name
	var type_ string = "RelKey"
	this.Type = type_
	return &this
}

// GetKeys returns the Keys field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RelKey) GetKeys() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Keys
}

// GetKeysOk returns a tuple with the Keys field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelKey) GetKeysOk() (*[]string, bool) {
	if o == nil || o.Keys == nil {
		return nil, false
	}
	return &o.Keys, true
}

// HasKeys returns a boolean if a field has been set.
func (o *RelKey) HasKeys() bool {
	if o != nil && o.Keys != nil {
		return true
	}

	return false
}

// SetKeys gets a reference to the given []string and assigns it to the Keys field.
func (o *RelKey) SetKeys(v []string) {
	o.Keys = v
}

// GetName returns the Name field value
func (o *RelKey) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RelKey) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RelKey) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RelKey) GetValues() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RelKey) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RelKey) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *RelKey) SetValues(v []string) {
	o.Values = v
}

// GetType returns the Type field value
func (o *RelKey) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RelKey) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RelKey) SetType(v string) {
	o.Type = v
}

func (o RelKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Keys != nil {
		toSerialize["keys"] = o.Keys
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableRelKey struct {
	value *RelKey
	isSet bool
}

func (v NullableRelKey) Get() *RelKey {
	return v.value
}

func (v *NullableRelKey) Set(val *RelKey) {
	v.value = val
	v.isSet = true
}

func (v NullableRelKey) IsSet() bool {
	return v.isSet
}

func (v *NullableRelKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRelKey(val *RelKey) *NullableRelKey {
	return &NullableRelKey{value: val, isSet: true}
}

func (v NullableRelKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRelKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


