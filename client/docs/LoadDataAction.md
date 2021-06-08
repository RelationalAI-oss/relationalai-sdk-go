# LoadDataAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rel** | **string** |  | [default to ""]
**Value** | [**LoadData**](LoadData.md) |  | 

## Methods

### NewLoadDataAction

`func NewLoadDataAction(rel string, value LoadData, ) *LoadDataAction`

NewLoadDataAction instantiates a new LoadDataAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadDataActionWithDefaults

`func NewLoadDataActionWithDefaults() *LoadDataAction`

NewLoadDataActionWithDefaults instantiates a new LoadDataAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRel

`func (o *LoadDataAction) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *LoadDataAction) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *LoadDataAction) SetRel(v string)`

SetRel sets Rel field to given value.


### GetValue

`func (o *LoadDataAction) GetValue() LoadData`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoadDataAction) GetValueOk() (*LoadData, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoadDataAction) SetValue(v LoadData)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


