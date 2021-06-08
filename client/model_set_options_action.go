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

// SetOptionsAction struct for SetOptionsAction
type SetOptionsAction struct {
	Action
	AbortOnError NullableBool `json:"abort_on_error,omitempty"`
	Debug NullableBool `json:"debug,omitempty"`
	DebugTrace NullableBool `json:"debug_trace,omitempty"`
	Silent NullableBool `json:"silent,omitempty"`
}

// NewSetOptionsAction instantiates a new SetOptionsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetOptionsAction() *SetOptionsAction {
	this := SetOptionsAction{}
	return &this
}

// NewSetOptionsActionWithDefaults instantiates a new SetOptionsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetOptionsActionWithDefaults() *SetOptionsAction {
	this := SetOptionsAction{}
	return &this
}

// GetAbortOnError returns the AbortOnError field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetOptionsAction) GetAbortOnError() bool {
	if o == nil || o.AbortOnError.Get() == nil {
		var ret bool
		return ret
	}
	return *o.AbortOnError.Get()
}

// GetAbortOnErrorOk returns a tuple with the AbortOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetOptionsAction) GetAbortOnErrorOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AbortOnError.Get(), o.AbortOnError.IsSet()
}

// HasAbortOnError returns a boolean if a field has been set.
func (o *SetOptionsAction) HasAbortOnError() bool {
	if o != nil && o.AbortOnError.IsSet() {
		return true
	}

	return false
}

// SetAbortOnError gets a reference to the given NullableBool and assigns it to the AbortOnError field.
func (o *SetOptionsAction) SetAbortOnError(v bool) {
	o.AbortOnError.Set(&v)
}
// SetAbortOnErrorNil sets the value for AbortOnError to be an explicit nil
func (o *SetOptionsAction) SetAbortOnErrorNil() {
	o.AbortOnError.Set(nil)
}

// UnsetAbortOnError ensures that no value is present for AbortOnError, not even an explicit nil
func (o *SetOptionsAction) UnsetAbortOnError() {
	o.AbortOnError.Unset()
}

// GetDebug returns the Debug field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetOptionsAction) GetDebug() bool {
	if o == nil || o.Debug.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Debug.Get()
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetOptionsAction) GetDebugOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Debug.Get(), o.Debug.IsSet()
}

// HasDebug returns a boolean if a field has been set.
func (o *SetOptionsAction) HasDebug() bool {
	if o != nil && o.Debug.IsSet() {
		return true
	}

	return false
}

// SetDebug gets a reference to the given NullableBool and assigns it to the Debug field.
func (o *SetOptionsAction) SetDebug(v bool) {
	o.Debug.Set(&v)
}
// SetDebugNil sets the value for Debug to be an explicit nil
func (o *SetOptionsAction) SetDebugNil() {
	o.Debug.Set(nil)
}

// UnsetDebug ensures that no value is present for Debug, not even an explicit nil
func (o *SetOptionsAction) UnsetDebug() {
	o.Debug.Unset()
}

// GetDebugTrace returns the DebugTrace field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetOptionsAction) GetDebugTrace() bool {
	if o == nil || o.DebugTrace.Get() == nil {
		var ret bool
		return ret
	}
	return *o.DebugTrace.Get()
}

// GetDebugTraceOk returns a tuple with the DebugTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetOptionsAction) GetDebugTraceOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DebugTrace.Get(), o.DebugTrace.IsSet()
}

// HasDebugTrace returns a boolean if a field has been set.
func (o *SetOptionsAction) HasDebugTrace() bool {
	if o != nil && o.DebugTrace.IsSet() {
		return true
	}

	return false
}

// SetDebugTrace gets a reference to the given NullableBool and assigns it to the DebugTrace field.
func (o *SetOptionsAction) SetDebugTrace(v bool) {
	o.DebugTrace.Set(&v)
}
// SetDebugTraceNil sets the value for DebugTrace to be an explicit nil
func (o *SetOptionsAction) SetDebugTraceNil() {
	o.DebugTrace.Set(nil)
}

// UnsetDebugTrace ensures that no value is present for DebugTrace, not even an explicit nil
func (o *SetOptionsAction) UnsetDebugTrace() {
	o.DebugTrace.Unset()
}

// GetSilent returns the Silent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SetOptionsAction) GetSilent() bool {
	if o == nil || o.Silent.Get() == nil {
		var ret bool
		return ret
	}
	return *o.Silent.Get()
}

// GetSilentOk returns a tuple with the Silent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SetOptionsAction) GetSilentOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Silent.Get(), o.Silent.IsSet()
}

// HasSilent returns a boolean if a field has been set.
func (o *SetOptionsAction) HasSilent() bool {
	if o != nil && o.Silent.IsSet() {
		return true
	}

	return false
}

// SetSilent gets a reference to the given NullableBool and assigns it to the Silent field.
func (o *SetOptionsAction) SetSilent(v bool) {
	o.Silent.Set(&v)
}
// SetSilentNil sets the value for Silent to be an explicit nil
func (o *SetOptionsAction) SetSilentNil() {
	o.Silent.Set(nil)
}

// UnsetSilent ensures that no value is present for Silent, not even an explicit nil
func (o *SetOptionsAction) UnsetSilent() {
	o.Silent.Unset()
}

func (o SetOptionsAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAction, errAction := json.Marshal(o.Action)
	if errAction != nil {
		return []byte{}, errAction
	}
	errAction = json.Unmarshal([]byte(serializedAction), &toSerialize)
	if errAction != nil {
		return []byte{}, errAction
	}
	if o.AbortOnError.IsSet() {
		toSerialize["abort_on_error"] = o.AbortOnError.Get()
	}
	if o.Debug.IsSet() {
		toSerialize["debug"] = o.Debug.Get()
	}
	if o.DebugTrace.IsSet() {
		toSerialize["debug_trace"] = o.DebugTrace.Get()
	}
	if o.Silent.IsSet() {
		toSerialize["silent"] = o.Silent.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableSetOptionsAction struct {
	value *SetOptionsAction
	isSet bool
}

func (v NullableSetOptionsAction) Get() *SetOptionsAction {
	return v.value
}

func (v *NullableSetOptionsAction) Set(val *SetOptionsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSetOptionsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSetOptionsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetOptionsAction(val *SetOptionsAction) *NullableSetOptionsAction {
	return &NullableSetOptionsAction{value: val, isSet: true}
}

func (v NullableSetOptionsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetOptionsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


