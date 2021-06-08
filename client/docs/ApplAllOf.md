# ApplAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**[]SyntaxNode**](SyntaxNode.md) |  | [optional] 
**Error** | **bool** |  | [default to false]
**Missing** | **bool** |  | [default to false]
**Range** | [**Range**](Range.md) |  | 
**Symbol** | **string** |  | [default to ""]

## Methods

### NewApplAllOf

`func NewApplAllOf(error_ bool, missing bool, range_ Range, symbol string, ) *ApplAllOf`

NewApplAllOf instantiates a new ApplAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplAllOfWithDefaults

`func NewApplAllOfWithDefaults() *ApplAllOf`

NewApplAllOfWithDefaults instantiates a new ApplAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *ApplAllOf) GetArguments() []SyntaxNode`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *ApplAllOf) GetArgumentsOk() (*[]SyntaxNode, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *ApplAllOf) SetArguments(v []SyntaxNode)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *ApplAllOf) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### SetArgumentsNil

`func (o *ApplAllOf) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *ApplAllOf) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil
### GetError

`func (o *ApplAllOf) GetError() bool`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApplAllOf) GetErrorOk() (*bool, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApplAllOf) SetError(v bool)`

SetError sets Error field to given value.


### GetMissing

`func (o *ApplAllOf) GetMissing() bool`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *ApplAllOf) GetMissingOk() (*bool, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *ApplAllOf) SetMissing(v bool)`

SetMissing sets Missing field to given value.


### GetRange

`func (o *ApplAllOf) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *ApplAllOf) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *ApplAllOf) SetRange(v Range)`

SetRange sets Range field to given value.


### GetSymbol

`func (o *ApplAllOf) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *ApplAllOf) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *ApplAllOf) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


