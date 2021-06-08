# CSVFileSyntax

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datarow** | **int32** |  | [default to 0]
**Delim** | **string** |  | [default to ""]
**Escapechar** | **string** |  | [default to ""]
**Header** | Pointer to **[]string** |  | [optional] 
**HeaderRow** | **int32** |  | [default to 0]
**Ignorerepeated** | **bool** |  | [default to false]
**Missingstrings** | Pointer to **[]string** |  | [optional] 
**Normalizenames** | **bool** |  | [default to false]
**Quotechar** | **string** |  | [default to ""]

## Methods

### NewCSVFileSyntax

`func NewCSVFileSyntax(datarow int32, delim string, escapechar string, headerRow int32, ignorerepeated bool, normalizenames bool, quotechar string, ) *CSVFileSyntax`

NewCSVFileSyntax instantiates a new CSVFileSyntax object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSVFileSyntaxWithDefaults

`func NewCSVFileSyntaxWithDefaults() *CSVFileSyntax`

NewCSVFileSyntaxWithDefaults instantiates a new CSVFileSyntax object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatarow

`func (o *CSVFileSyntax) GetDatarow() int32`

GetDatarow returns the Datarow field if non-nil, zero value otherwise.

### GetDatarowOk

`func (o *CSVFileSyntax) GetDatarowOk() (*int32, bool)`

GetDatarowOk returns a tuple with the Datarow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatarow

`func (o *CSVFileSyntax) SetDatarow(v int32)`

SetDatarow sets Datarow field to given value.


### GetDelim

`func (o *CSVFileSyntax) GetDelim() string`

GetDelim returns the Delim field if non-nil, zero value otherwise.

### GetDelimOk

`func (o *CSVFileSyntax) GetDelimOk() (*string, bool)`

GetDelimOk returns a tuple with the Delim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelim

`func (o *CSVFileSyntax) SetDelim(v string)`

SetDelim sets Delim field to given value.


### GetEscapechar

`func (o *CSVFileSyntax) GetEscapechar() string`

GetEscapechar returns the Escapechar field if non-nil, zero value otherwise.

### GetEscapecharOk

`func (o *CSVFileSyntax) GetEscapecharOk() (*string, bool)`

GetEscapecharOk returns a tuple with the Escapechar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscapechar

`func (o *CSVFileSyntax) SetEscapechar(v string)`

SetEscapechar sets Escapechar field to given value.


### GetHeader

`func (o *CSVFileSyntax) GetHeader() []string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CSVFileSyntax) GetHeaderOk() (*[]string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CSVFileSyntax) SetHeader(v []string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CSVFileSyntax) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### SetHeaderNil

`func (o *CSVFileSyntax) SetHeaderNil(b bool)`

 SetHeaderNil sets the value for Header to be an explicit nil

### UnsetHeader
`func (o *CSVFileSyntax) UnsetHeader()`

UnsetHeader ensures that no value is present for Header, not even an explicit nil
### GetHeaderRow

`func (o *CSVFileSyntax) GetHeaderRow() int32`

GetHeaderRow returns the HeaderRow field if non-nil, zero value otherwise.

### GetHeaderRowOk

`func (o *CSVFileSyntax) GetHeaderRowOk() (*int32, bool)`

GetHeaderRowOk returns a tuple with the HeaderRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderRow

`func (o *CSVFileSyntax) SetHeaderRow(v int32)`

SetHeaderRow sets HeaderRow field to given value.


### GetIgnorerepeated

`func (o *CSVFileSyntax) GetIgnorerepeated() bool`

GetIgnorerepeated returns the Ignorerepeated field if non-nil, zero value otherwise.

### GetIgnorerepeatedOk

`func (o *CSVFileSyntax) GetIgnorerepeatedOk() (*bool, bool)`

GetIgnorerepeatedOk returns a tuple with the Ignorerepeated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorerepeated

`func (o *CSVFileSyntax) SetIgnorerepeated(v bool)`

SetIgnorerepeated sets Ignorerepeated field to given value.


### GetMissingstrings

`func (o *CSVFileSyntax) GetMissingstrings() []string`

GetMissingstrings returns the Missingstrings field if non-nil, zero value otherwise.

### GetMissingstringsOk

`func (o *CSVFileSyntax) GetMissingstringsOk() (*[]string, bool)`

GetMissingstringsOk returns a tuple with the Missingstrings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingstrings

`func (o *CSVFileSyntax) SetMissingstrings(v []string)`

SetMissingstrings sets Missingstrings field to given value.

### HasMissingstrings

`func (o *CSVFileSyntax) HasMissingstrings() bool`

HasMissingstrings returns a boolean if a field has been set.

### SetMissingstringsNil

`func (o *CSVFileSyntax) SetMissingstringsNil(b bool)`

 SetMissingstringsNil sets the value for Missingstrings to be an explicit nil

### UnsetMissingstrings
`func (o *CSVFileSyntax) UnsetMissingstrings()`

UnsetMissingstrings ensures that no value is present for Missingstrings, not even an explicit nil
### GetNormalizenames

`func (o *CSVFileSyntax) GetNormalizenames() bool`

GetNormalizenames returns the Normalizenames field if non-nil, zero value otherwise.

### GetNormalizenamesOk

`func (o *CSVFileSyntax) GetNormalizenamesOk() (*bool, bool)`

GetNormalizenamesOk returns a tuple with the Normalizenames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizenames

`func (o *CSVFileSyntax) SetNormalizenames(v bool)`

SetNormalizenames sets Normalizenames field to given value.


### GetQuotechar

`func (o *CSVFileSyntax) GetQuotechar() string`

GetQuotechar returns the Quotechar field if non-nil, zero value otherwise.

### GetQuotecharOk

`func (o *CSVFileSyntax) GetQuotecharOk() (*string, bool)`

GetQuotecharOk returns a tuple with the Quotechar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotechar

`func (o *CSVFileSyntax) SetQuotechar(v string)`

SetQuotechar sets Quotechar field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


