# QueryActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Output** | Pointer to [**[]Relation**](Relation.md) |  | [optional] 

## Methods

### NewQueryActionResult

`func NewQueryActionResult() *QueryActionResult`

NewQueryActionResult instantiates a new QueryActionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryActionResultWithDefaults

`func NewQueryActionResultWithDefaults() *QueryActionResult`

NewQueryActionResultWithDefaults instantiates a new QueryActionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutput

`func (o *QueryActionResult) GetOutput() []Relation`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *QueryActionResult) GetOutputOk() (*[]Relation, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *QueryActionResult) SetOutput(v []Relation)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *QueryActionResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *QueryActionResult) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *QueryActionResult) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


