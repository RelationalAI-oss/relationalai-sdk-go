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

// LoadData struct for LoadData
type LoadData struct {
	ContentType string `json:"content_type"`
	Data NullableString `json:"data,omitempty"`
	FileSchema *FileSchema `json:"file_schema,omitempty"`
	FileSyntax *FileSyntax `json:"file_syntax,omitempty"`
	Integration *Integration `json:"integration,omitempty"`
	Key interface{} `json:"key"`
	Path NullableString `json:"path,omitempty"`
	Type string `json:"type"`
}

// NewLoadData instantiates a new LoadData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoadData(contentType string, key interface{}, type_ string, ) *LoadData {
	this := LoadData{}
	this.ContentType = contentType
	this.Key = key
	this.Type = type_
	return &this
}

// NewLoadDataWithDefaults instantiates a new LoadData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoadDataWithDefaults() *LoadData {
	this := LoadData{}
	var contentType string = ""
	this.ContentType = contentType
	var type_ string = "LoadData"
	this.Type = type_
	return &this
}

// GetContentType returns the ContentType field value
func (o *LoadData) GetContentType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value
// and a boolean to check if the value has been set.
func (o *LoadData) GetContentTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentType, true
}

// SetContentType sets field value
func (o *LoadData) SetContentType(v string) {
	o.ContentType = v
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadData) GetData() string {
	if o == nil || o.Data.Get() == nil {
		var ret string
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadData) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *LoadData) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableString and assigns it to the Data field.
func (o *LoadData) SetData(v string) {
	o.Data.Set(&v)
}
// SetDataNil sets the value for Data to be an explicit nil
func (o *LoadData) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *LoadData) UnsetData() {
	o.Data.Unset()
}

// GetFileSchema returns the FileSchema field value if set, zero value otherwise.
func (o *LoadData) GetFileSchema() FileSchema {
	if o == nil || o.FileSchema == nil {
		var ret FileSchema
		return ret
	}
	return *o.FileSchema
}

// GetFileSchemaOk returns a tuple with the FileSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadData) GetFileSchemaOk() (*FileSchema, bool) {
	if o == nil || o.FileSchema == nil {
		return nil, false
	}
	return o.FileSchema, true
}

// HasFileSchema returns a boolean if a field has been set.
func (o *LoadData) HasFileSchema() bool {
	if o != nil && o.FileSchema != nil {
		return true
	}

	return false
}

// SetFileSchema gets a reference to the given FileSchema and assigns it to the FileSchema field.
func (o *LoadData) SetFileSchema(v FileSchema) {
	o.FileSchema = &v
}

// GetFileSyntax returns the FileSyntax field value if set, zero value otherwise.
func (o *LoadData) GetFileSyntax() FileSyntax {
	if o == nil || o.FileSyntax == nil {
		var ret FileSyntax
		return ret
	}
	return *o.FileSyntax
}

// GetFileSyntaxOk returns a tuple with the FileSyntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadData) GetFileSyntaxOk() (*FileSyntax, bool) {
	if o == nil || o.FileSyntax == nil {
		return nil, false
	}
	return o.FileSyntax, true
}

// HasFileSyntax returns a boolean if a field has been set.
func (o *LoadData) HasFileSyntax() bool {
	if o != nil && o.FileSyntax != nil {
		return true
	}

	return false
}

// SetFileSyntax gets a reference to the given FileSyntax and assigns it to the FileSyntax field.
func (o *LoadData) SetFileSyntax(v FileSyntax) {
	o.FileSyntax = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *LoadData) GetIntegration() Integration {
	if o == nil || o.Integration == nil {
		var ret Integration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoadData) GetIntegrationOk() (*Integration, bool) {
	if o == nil || o.Integration == nil {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *LoadData) HasIntegration() bool {
	if o != nil && o.Integration != nil {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given Integration and assigns it to the Integration field.
func (o *LoadData) SetIntegration(v Integration) {
	o.Integration = &v
}

// GetKey returns the Key field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *LoadData) GetKey() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadData) GetKeyOk() (*interface{}, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *LoadData) SetKey(v interface{}) {
	o.Key = v
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *LoadData) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LoadData) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *LoadData) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *LoadData) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *LoadData) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *LoadData) UnsetPath() {
	o.Path.Unset()
}

// GetType returns the Type field value
func (o *LoadData) GetType() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *LoadData) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *LoadData) SetType(v string) {
	o.Type = v
}

func (o LoadData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["content_type"] = o.ContentType
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.FileSchema != nil {
		toSerialize["file_schema"] = o.FileSchema
	}
	if o.FileSyntax != nil {
		toSerialize["file_syntax"] = o.FileSyntax
	}
	if o.Integration != nil {
		toSerialize["integration"] = o.Integration
	}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableLoadData struct {
	value *LoadData
	isSet bool
}

func (v NullableLoadData) Get() *LoadData {
	return v.value
}

func (v *NullableLoadData) Set(val *LoadData) {
	v.value = val
	v.isSet = true
}

func (v NullableLoadData) IsSet() bool {
	return v.isSet
}

func (v *NullableLoadData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoadData(val *LoadData) *NullableLoadData {
	return &NullableLoadData{value: val, isSet: true}
}

func (v NullableLoadData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoadData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


