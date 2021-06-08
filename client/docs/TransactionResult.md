# TransactionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aborted** | **bool** |  | [default to false]
**Actions** | Pointer to [**[]LabeledActionResult**](LabeledActionResult.md) |  | [optional] 
**DebugLevel** | Pointer to **NullableInt32** |  | [optional] 
**Output** | Pointer to [**[]Relation**](Relation.md) |  | [optional] 
**Problems** | Pointer to [**[]AbstractProblem**](AbstractProblem.md) |  | [optional] 
**Version** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | [default to "TransactionResult"]

## Methods

### NewTransactionResult

`func NewTransactionResult(aborted bool, type_ string, ) *TransactionResult`

NewTransactionResult instantiates a new TransactionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionResultWithDefaults

`func NewTransactionResultWithDefaults() *TransactionResult`

NewTransactionResultWithDefaults instantiates a new TransactionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAborted

`func (o *TransactionResult) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *TransactionResult) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *TransactionResult) SetAborted(v bool)`

SetAborted sets Aborted field to given value.


### GetActions

`func (o *TransactionResult) GetActions() []LabeledActionResult`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *TransactionResult) GetActionsOk() (*[]LabeledActionResult, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *TransactionResult) SetActions(v []LabeledActionResult)`

SetActions sets Actions field to given value.

### HasActions

`func (o *TransactionResult) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *TransactionResult) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *TransactionResult) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetDebugLevel

`func (o *TransactionResult) GetDebugLevel() int32`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *TransactionResult) GetDebugLevelOk() (*int32, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *TransactionResult) SetDebugLevel(v int32)`

SetDebugLevel sets DebugLevel field to given value.

### HasDebugLevel

`func (o *TransactionResult) HasDebugLevel() bool`

HasDebugLevel returns a boolean if a field has been set.

### SetDebugLevelNil

`func (o *TransactionResult) SetDebugLevelNil(b bool)`

 SetDebugLevelNil sets the value for DebugLevel to be an explicit nil

### UnsetDebugLevel
`func (o *TransactionResult) UnsetDebugLevel()`

UnsetDebugLevel ensures that no value is present for DebugLevel, not even an explicit nil
### GetOutput

`func (o *TransactionResult) GetOutput() []Relation`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *TransactionResult) GetOutputOk() (*[]Relation, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *TransactionResult) SetOutput(v []Relation)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *TransactionResult) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *TransactionResult) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *TransactionResult) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetProblems

`func (o *TransactionResult) GetProblems() []AbstractProblem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *TransactionResult) GetProblemsOk() (*[]AbstractProblem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *TransactionResult) SetProblems(v []AbstractProblem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *TransactionResult) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### SetProblemsNil

`func (o *TransactionResult) SetProblemsNil(b bool)`

 SetProblemsNil sets the value for Problems to be an explicit nil

### UnsetProblems
`func (o *TransactionResult) UnsetProblems()`

UnsetProblems ensures that no value is present for Problems, not even an explicit nil
### GetVersion

`func (o *TransactionResult) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TransactionResult) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TransactionResult) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TransactionResult) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *TransactionResult) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *TransactionResult) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *TransactionResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


