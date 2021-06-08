# LoadData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | **string** |  | [default to ""]
**Data** | Pointer to **NullableString** |  | [optional] 
**FileSchema** | Pointer to [**FileSchema**](FileSchema.md) |  | [optional] 
**FileSyntax** | Pointer to [**FileSyntax**](FileSyntax.md) |  | [optional] 
**Integration** | Pointer to [**Integration**](Integration.md) |  | [optional] 
**Key** | **interface{}** |  | 
**Path** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | [default to "LoadData"]

## Methods

### NewLoadData

`func NewLoadData(contentType string, key interface{}, type_ string, ) *LoadData`

NewLoadData instantiates a new LoadData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadDataWithDefaults

`func NewLoadDataWithDefaults() *LoadData`

NewLoadDataWithDefaults instantiates a new LoadData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *LoadData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *LoadData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *LoadData) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetData

`func (o *LoadData) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoadData) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoadData) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *LoadData) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *LoadData) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *LoadData) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetFileSchema

`func (o *LoadData) GetFileSchema() FileSchema`

GetFileSchema returns the FileSchema field if non-nil, zero value otherwise.

### GetFileSchemaOk

`func (o *LoadData) GetFileSchemaOk() (*FileSchema, bool)`

GetFileSchemaOk returns a tuple with the FileSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSchema

`func (o *LoadData) SetFileSchema(v FileSchema)`

SetFileSchema sets FileSchema field to given value.

### HasFileSchema

`func (o *LoadData) HasFileSchema() bool`

HasFileSchema returns a boolean if a field has been set.

### GetFileSyntax

`func (o *LoadData) GetFileSyntax() FileSyntax`

GetFileSyntax returns the FileSyntax field if non-nil, zero value otherwise.

### GetFileSyntaxOk

`func (o *LoadData) GetFileSyntaxOk() (*FileSyntax, bool)`

GetFileSyntaxOk returns a tuple with the FileSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileSyntax

`func (o *LoadData) SetFileSyntax(v FileSyntax)`

SetFileSyntax sets FileSyntax field to given value.

### HasFileSyntax

`func (o *LoadData) HasFileSyntax() bool`

HasFileSyntax returns a boolean if a field has been set.

### GetIntegration

`func (o *LoadData) GetIntegration() Integration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *LoadData) GetIntegrationOk() (*Integration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *LoadData) SetIntegration(v Integration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *LoadData) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetKey

`func (o *LoadData) GetKey() interface{}`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LoadData) GetKeyOk() (*interface{}, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LoadData) SetKey(v interface{})`

SetKey sets Key field to given value.


### SetKeyNil

`func (o *LoadData) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *LoadData) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetPath

`func (o *LoadData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoadData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoadData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoadData) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *LoadData) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *LoadData) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetType

`func (o *LoadData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoadData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoadData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


