# QueryAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**[]Relation**](Relation.md) |  | [optional] 
**Outputs** | Pointer to **[]string** |  | [optional] 
**Persist** | Pointer to **[]string** |  | [optional] 
**Source** | [**Source**](Source.md) |  | 

## Methods

### NewQueryAction

`func NewQueryAction(source Source, ) *QueryAction`

NewQueryAction instantiates a new QueryAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryActionWithDefaults

`func NewQueryActionWithDefaults() *QueryAction`

NewQueryActionWithDefaults instantiates a new QueryAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *QueryAction) GetInputs() []Relation`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *QueryAction) GetInputsOk() (*[]Relation, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *QueryAction) SetInputs(v []Relation)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *QueryAction) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *QueryAction) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *QueryAction) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil
### GetOutputs

`func (o *QueryAction) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *QueryAction) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *QueryAction) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *QueryAction) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### SetOutputsNil

`func (o *QueryAction) SetOutputsNil(b bool)`

 SetOutputsNil sets the value for Outputs to be an explicit nil

### UnsetOutputs
`func (o *QueryAction) UnsetOutputs()`

UnsetOutputs ensures that no value is present for Outputs, not even an explicit nil
### GetPersist

`func (o *QueryAction) GetPersist() []string`

GetPersist returns the Persist field if non-nil, zero value otherwise.

### GetPersistOk

`func (o *QueryAction) GetPersistOk() (*[]string, bool)`

GetPersistOk returns a tuple with the Persist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersist

`func (o *QueryAction) SetPersist(v []string)`

SetPersist sets Persist field to given value.

### HasPersist

`func (o *QueryAction) HasPersist() bool`

HasPersist returns a boolean if a field has been set.

### SetPersistNil

`func (o *QueryAction) SetPersistNil(b bool)`

 SetPersistNil sets the value for Persist to be an explicit nil

### UnsetPersist
`func (o *QueryAction) UnsetPersist()`

UnsetPersist ensures that no value is present for Persist, not even an explicit nil
### GetSource

`func (o *QueryAction) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *QueryAction) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *QueryAction) SetSource(v Source)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


