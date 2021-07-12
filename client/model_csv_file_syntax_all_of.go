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

// CSVFileSyntaxAllOf struct for CSVFileSyntaxAllOf
type CSVFileSyntaxAllOf struct {
	Datarow int32 `json:"datarow"`
	Delim string `json:"delim"`
	Escapechar string `json:"escapechar"`
	Header []string `json:"header,omitempty"`
	HeaderRow int32 `json:"header_row"`
	Ignorerepeated bool `json:"ignorerepeated"`
	Missingstrings []string `json:"missingstrings,omitempty"`
	Normalizenames bool `json:"normalizenames"`
	Quotechar string `json:"quotechar"`
}

// NewCSVFileSyntaxAllOf instantiates a new CSVFileSyntaxAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSVFileSyntaxAllOf(datarow int32, delim string, escapechar string, headerRow int32, ignorerepeated bool, normalizenames bool, quotechar string, ) *CSVFileSyntaxAllOf {
	this := CSVFileSyntaxAllOf{}
	this.Datarow = datarow
	this.Delim = delim
	this.Escapechar = escapechar
	this.HeaderRow = headerRow
	this.Ignorerepeated = ignorerepeated
	this.Normalizenames = normalizenames
	this.Quotechar = quotechar
	return &this
}

// NewCSVFileSyntaxAllOfWithDefaults instantiates a new CSVFileSyntaxAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSVFileSyntaxAllOfWithDefaults() *CSVFileSyntaxAllOf {
	this := CSVFileSyntaxAllOf{}
	var datarow int32 = 0
	this.Datarow = datarow
	var delim string = ""
	this.Delim = delim
	var escapechar string = ""
	this.Escapechar = escapechar
	var headerRow int32 = 0
	this.HeaderRow = headerRow
	var ignorerepeated bool = false
	this.Ignorerepeated = ignorerepeated
	var normalizenames bool = false
	this.Normalizenames = normalizenames
	var quotechar string = ""
	this.Quotechar = quotechar
	return &this
}

// GetDatarow returns the Datarow field value
func (o *CSVFileSyntaxAllOf) GetDatarow() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Datarow
}

// GetDatarowOk returns a tuple with the Datarow field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetDatarowOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Datarow, true
}

// SetDatarow sets field value
func (o *CSVFileSyntaxAllOf) SetDatarow(v int32) {
	o.Datarow = v
}

// GetDelim returns the Delim field value
func (o *CSVFileSyntaxAllOf) GetDelim() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Delim
}

// GetDelimOk returns a tuple with the Delim field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetDelimOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Delim, true
}

// SetDelim sets field value
func (o *CSVFileSyntaxAllOf) SetDelim(v string) {
	o.Delim = v
}

// GetEscapechar returns the Escapechar field value
func (o *CSVFileSyntaxAllOf) GetEscapechar() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Escapechar
}

// GetEscapecharOk returns a tuple with the Escapechar field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetEscapecharOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Escapechar, true
}

// SetEscapechar sets field value
func (o *CSVFileSyntaxAllOf) SetEscapechar(v string) {
	o.Escapechar = v
}

// GetHeader returns the Header field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CSVFileSyntaxAllOf) GetHeader() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CSVFileSyntaxAllOf) GetHeaderOk() (*[]string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return &o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *CSVFileSyntaxAllOf) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given []string and assigns it to the Header field.
func (o *CSVFileSyntaxAllOf) SetHeader(v []string) {
	o.Header = v
}

// GetHeaderRow returns the HeaderRow field value
func (o *CSVFileSyntaxAllOf) GetHeaderRow() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.HeaderRow
}

// GetHeaderRowOk returns a tuple with the HeaderRow field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetHeaderRowOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HeaderRow, true
}

// SetHeaderRow sets field value
func (o *CSVFileSyntaxAllOf) SetHeaderRow(v int32) {
	o.HeaderRow = v
}

// GetIgnorerepeated returns the Ignorerepeated field value
func (o *CSVFileSyntaxAllOf) GetIgnorerepeated() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Ignorerepeated
}

// GetIgnorerepeatedOk returns a tuple with the Ignorerepeated field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetIgnorerepeatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Ignorerepeated, true
}

// SetIgnorerepeated sets field value
func (o *CSVFileSyntaxAllOf) SetIgnorerepeated(v bool) {
	o.Ignorerepeated = v
}

// GetMissingstrings returns the Missingstrings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CSVFileSyntaxAllOf) GetMissingstrings() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Missingstrings
}

// GetMissingstringsOk returns a tuple with the Missingstrings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CSVFileSyntaxAllOf) GetMissingstringsOk() (*[]string, bool) {
	if o == nil || o.Missingstrings == nil {
		return nil, false
	}
	return &o.Missingstrings, true
}

// HasMissingstrings returns a boolean if a field has been set.
func (o *CSVFileSyntaxAllOf) HasMissingstrings() bool {
	if o != nil && o.Missingstrings != nil {
		return true
	}

	return false
}

// SetMissingstrings gets a reference to the given []string and assigns it to the Missingstrings field.
func (o *CSVFileSyntaxAllOf) SetMissingstrings(v []string) {
	o.Missingstrings = v
}

// GetNormalizenames returns the Normalizenames field value
func (o *CSVFileSyntaxAllOf) GetNormalizenames() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Normalizenames
}

// GetNormalizenamesOk returns a tuple with the Normalizenames field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetNormalizenamesOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Normalizenames, true
}

// SetNormalizenames sets field value
func (o *CSVFileSyntaxAllOf) SetNormalizenames(v bool) {
	o.Normalizenames = v
}

// GetQuotechar returns the Quotechar field value
func (o *CSVFileSyntaxAllOf) GetQuotechar() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Quotechar
}

// GetQuotecharOk returns a tuple with the Quotechar field value
// and a boolean to check if the value has been set.
func (o *CSVFileSyntaxAllOf) GetQuotecharOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Quotechar, true
}

// SetQuotechar sets field value
func (o *CSVFileSyntaxAllOf) SetQuotechar(v string) {
	o.Quotechar = v
}

func (o CSVFileSyntaxAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datarow"] = o.Datarow
	}
	if true {
		toSerialize["delim"] = o.Delim
	}
	if true {
		toSerialize["escapechar"] = o.Escapechar
	}
	if o.Header != nil {
		toSerialize["header"] = o.Header
	}
	if true {
		toSerialize["header_row"] = o.HeaderRow
	}
	if true {
		toSerialize["ignorerepeated"] = o.Ignorerepeated
	}
	if o.Missingstrings != nil {
		toSerialize["missingstrings"] = o.Missingstrings
	}
	if true {
		toSerialize["normalizenames"] = o.Normalizenames
	}
	if true {
		toSerialize["quotechar"] = o.Quotechar
	}
	return json.Marshal(toSerialize)
}

type NullableCSVFileSyntaxAllOf struct {
	value *CSVFileSyntaxAllOf
	isSet bool
}

func (v NullableCSVFileSyntaxAllOf) Get() *CSVFileSyntaxAllOf {
	return v.value
}

func (v *NullableCSVFileSyntaxAllOf) Set(val *CSVFileSyntaxAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCSVFileSyntaxAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCSVFileSyntaxAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSVFileSyntaxAllOf(val *CSVFileSyntaxAllOf) *NullableCSVFileSyntaxAllOf {
	return &NullableCSVFileSyntaxAllOf{value: val, isSet: true}
}

func (v NullableCSVFileSyntaxAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSVFileSyntaxAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


