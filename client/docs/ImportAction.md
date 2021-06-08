# ImportAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inputs** | Pointer to [**[]Relation**](Relation.md) |  | [optional] 

## Methods

### NewImportAction

`func NewImportAction() *ImportAction`

NewImportAction instantiates a new ImportAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportActionWithDefaults

`func NewImportActionWithDefaults() *ImportAction`

NewImportActionWithDefaults instantiates a new ImportAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputs

`func (o *ImportAction) GetInputs() []Relation`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *ImportAction) GetInputsOk() (*[]Relation, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *ImportAction) SetInputs(v []Relation)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *ImportAction) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### SetInputsNil

`func (o *ImportAction) SetInputsNil(b bool)`

 SetInputsNil sets the value for Inputs to be an explicit nil

### UnsetInputs
`func (o *ImportAction) UnsetInputs()`

UnsetInputs ensures that no value is present for Inputs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


