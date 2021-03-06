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

// InstallAction struct for InstallAction
type InstallAction struct {
	Action
	Sources []Source `json:"sources,omitempty"`
}

// NewInstallAction instantiates a new InstallAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstallAction() *InstallAction {
	this := InstallAction{}
	return &this
}

// NewInstallActionWithDefaults instantiates a new InstallAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstallActionWithDefaults() *InstallAction {
	this := InstallAction{}
	return &this
}

// GetSources returns the Sources field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstallAction) GetSources() []Source {
	if o == nil  {
		var ret []Source
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstallAction) GetSourcesOk() (*[]Source, bool) {
	if o == nil || o.Sources == nil {
		return nil, false
	}
	return &o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *InstallAction) HasSources() bool {
	if o != nil && o.Sources != nil {
		return true
	}

	return false
}

// SetSources gets a reference to the given []Source and assigns it to the Sources field.
func (o *InstallAction) SetSources(v []Source) {
	o.Sources = v
}

func (o InstallAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return []byte{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return []byte{}, errAction
	}
	if o.Sources != nil {
		toSerialize["sources"] = o.Sources
	}
	return json.Marshal(toSerialize)
}

type NullableInstallAction struct {
	value *InstallAction
	isSet bool
}

func (v NullableInstallAction) Get() *InstallAction {
	return v.value
}

func (v *NullableInstallAction) Set(val *InstallAction) {
	v.value = val
	v.isSet = true
}

func (v NullableInstallAction) IsSet() bool {
	return v.isSet
}

func (v *NullableInstallAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstallAction(val *InstallAction) *NullableInstallAction {
	return &NullableInstallAction{value: val, isSet: true}
}

func (v NullableInstallAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstallAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


