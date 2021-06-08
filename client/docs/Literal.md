# Literal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Missing** | **bool** |  | [default to false]
**Range** | [**Range**](Range.md) |  | 
**Value** | **string** |  | [default to ""]

## Methods

### NewLiteral

`func NewLiteral(missing bool, range_ Range, value string, ) *Literal`

NewLiteral instantiates a new Literal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiteralWithDefaults

`func NewLiteralWithDefaults() *Literal`

NewLiteralWithDefaults instantiates a new Literal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMissing

`func (o *Literal) GetMissing() bool`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *Literal) GetMissingOk() (*bool, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *Literal) SetMissing(v bool)`

SetMissing sets Missing field to given value.


### GetRange

`func (o *Literal) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *Literal) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *Literal) SetRange(v Range)`

SetRange sets Range field to given value.


### GetValue

`func (o *Literal) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Literal) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Literal) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


