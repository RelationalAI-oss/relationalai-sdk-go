# UpdateAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delta** | Pointer to [**[]PairAnyValueAnyValue**](PairAnyValueAnyValue.md) |  | [optional] 
**Rel** | [**RelKey**](RelKey.md) |  | 
**Updates** | Pointer to [**[]PairAnyValueAnyValue**](PairAnyValueAnyValue.md) |  | [optional] 

## Methods

### NewUpdateAction

`func NewUpdateAction(rel RelKey, ) *UpdateAction`

NewUpdateAction instantiates a new UpdateAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActionWithDefaults

`func NewUpdateActionWithDefaults() *UpdateAction`

NewUpdateActionWithDefaults instantiates a new UpdateAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelta

`func (o *UpdateAction) GetDelta() []PairAnyValueAnyValue`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *UpdateAction) GetDeltaOk() (*[]PairAnyValueAnyValue, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *UpdateAction) SetDelta(v []PairAnyValueAnyValue)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *UpdateAction) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *UpdateAction) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *UpdateAction) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetRel

`func (o *UpdateAction) GetRel() RelKey`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *UpdateAction) GetRelOk() (*RelKey, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *UpdateAction) SetRel(v RelKey)`

SetRel sets Rel field to given value.


### GetUpdates

`func (o *UpdateAction) GetUpdates() []PairAnyValueAnyValue`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *UpdateAction) GetUpdatesOk() (*[]PairAnyValueAnyValue, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *UpdateAction) SetUpdates(v []PairAnyValueAnyValue)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *UpdateAction) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### SetUpdatesNil

`func (o *UpdateAction) SetUpdatesNil(b bool)`

 SetUpdatesNil sets the value for Updates to be an explicit nil

### UnsetUpdates
`func (o *UpdateAction) UnsetUpdates()`

UnsetUpdates ensures that no value is present for Updates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


