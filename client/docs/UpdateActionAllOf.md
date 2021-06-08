# UpdateActionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Delta** | Pointer to [**[]PairAnyValueAnyValue**](PairAnyValueAnyValue.md) |  | [optional] 
**Rel** | [**RelKey**](RelKey.md) |  | 
**Updates** | Pointer to [**[]PairAnyValueAnyValue**](PairAnyValueAnyValue.md) |  | [optional] 

## Methods

### NewUpdateActionAllOf

`func NewUpdateActionAllOf(rel RelKey, ) *UpdateActionAllOf`

NewUpdateActionAllOf instantiates a new UpdateActionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateActionAllOfWithDefaults

`func NewUpdateActionAllOfWithDefaults() *UpdateActionAllOf`

NewUpdateActionAllOfWithDefaults instantiates a new UpdateActionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDelta

`func (o *UpdateActionAllOf) GetDelta() []PairAnyValueAnyValue`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *UpdateActionAllOf) GetDeltaOk() (*[]PairAnyValueAnyValue, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *UpdateActionAllOf) SetDelta(v []PairAnyValueAnyValue)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *UpdateActionAllOf) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *UpdateActionAllOf) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *UpdateActionAllOf) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetRel

`func (o *UpdateActionAllOf) GetRel() RelKey`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *UpdateActionAllOf) GetRelOk() (*RelKey, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *UpdateActionAllOf) SetRel(v RelKey)`

SetRel sets Rel field to given value.


### GetUpdates

`func (o *UpdateActionAllOf) GetUpdates() []PairAnyValueAnyValue`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *UpdateActionAllOf) GetUpdatesOk() (*[]PairAnyValueAnyValue, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *UpdateActionAllOf) SetUpdates(v []PairAnyValueAnyValue)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *UpdateActionAllOf) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### SetUpdatesNil

`func (o *UpdateActionAllOf) SetUpdatesNil(b bool)`

 SetUpdatesNil sets the value for Updates to be an explicit nil

### UnsetUpdates
`func (o *UpdateActionAllOf) UnsetUpdates()`

UnsetUpdates ensures that no value is present for Updates, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


