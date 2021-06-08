# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abort** | **bool** |  | [default to false]
**Actions** | Pointer to [**[]LabeledAction**](LabeledAction.md) |  | [optional] 
**Dbname** | **string** |  | [default to ""]
**DebugLevel** | Pointer to **NullableInt32** |  | [optional] 
**Mode** | **string** |  | [default to "OPEN"]
**NowaitDurable** | **bool** |  | [default to false]
**Readonly** | **bool** |  | [default to false]
**SourceDbname** | Pointer to **NullableString** |  | [optional] 
**Version** | Pointer to **NullableInt32** |  | [optional] 
**Type** | **string** |  | [default to "Transaction"]

## Methods

### NewTransaction

`func NewTransaction(abort bool, dbname string, mode string, nowaitDurable bool, readonly bool, type_ string, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbort

`func (o *Transaction) GetAbort() bool`

GetAbort returns the Abort field if non-nil, zero value otherwise.

### GetAbortOk

`func (o *Transaction) GetAbortOk() (*bool, bool)`

GetAbortOk returns a tuple with the Abort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbort

`func (o *Transaction) SetAbort(v bool)`

SetAbort sets Abort field to given value.


### GetActions

`func (o *Transaction) GetActions() []LabeledAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Transaction) GetActionsOk() (*[]LabeledAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Transaction) SetActions(v []LabeledAction)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Transaction) HasActions() bool`

HasActions returns a boolean if a field has been set.

### SetActionsNil

`func (o *Transaction) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *Transaction) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetDbname

`func (o *Transaction) GetDbname() string`

GetDbname returns the Dbname field if non-nil, zero value otherwise.

### GetDbnameOk

`func (o *Transaction) GetDbnameOk() (*string, bool)`

GetDbnameOk returns a tuple with the Dbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbname

`func (o *Transaction) SetDbname(v string)`

SetDbname sets Dbname field to given value.


### GetDebugLevel

`func (o *Transaction) GetDebugLevel() int32`

GetDebugLevel returns the DebugLevel field if non-nil, zero value otherwise.

### GetDebugLevelOk

`func (o *Transaction) GetDebugLevelOk() (*int32, bool)`

GetDebugLevelOk returns a tuple with the DebugLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugLevel

`func (o *Transaction) SetDebugLevel(v int32)`

SetDebugLevel sets DebugLevel field to given value.

### HasDebugLevel

`func (o *Transaction) HasDebugLevel() bool`

HasDebugLevel returns a boolean if a field has been set.

### SetDebugLevelNil

`func (o *Transaction) SetDebugLevelNil(b bool)`

 SetDebugLevelNil sets the value for DebugLevel to be an explicit nil

### UnsetDebugLevel
`func (o *Transaction) UnsetDebugLevel()`

UnsetDebugLevel ensures that no value is present for DebugLevel, not even an explicit nil
### GetMode

`func (o *Transaction) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Transaction) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Transaction) SetMode(v string)`

SetMode sets Mode field to given value.


### GetNowaitDurable

`func (o *Transaction) GetNowaitDurable() bool`

GetNowaitDurable returns the NowaitDurable field if non-nil, zero value otherwise.

### GetNowaitDurableOk

`func (o *Transaction) GetNowaitDurableOk() (*bool, bool)`

GetNowaitDurableOk returns a tuple with the NowaitDurable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNowaitDurable

`func (o *Transaction) SetNowaitDurable(v bool)`

SetNowaitDurable sets NowaitDurable field to given value.


### GetReadonly

`func (o *Transaction) GetReadonly() bool`

GetReadonly returns the Readonly field if non-nil, zero value otherwise.

### GetReadonlyOk

`func (o *Transaction) GetReadonlyOk() (*bool, bool)`

GetReadonlyOk returns a tuple with the Readonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadonly

`func (o *Transaction) SetReadonly(v bool)`

SetReadonly sets Readonly field to given value.


### GetSourceDbname

`func (o *Transaction) GetSourceDbname() string`

GetSourceDbname returns the SourceDbname field if non-nil, zero value otherwise.

### GetSourceDbnameOk

`func (o *Transaction) GetSourceDbnameOk() (*string, bool)`

GetSourceDbnameOk returns a tuple with the SourceDbname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDbname

`func (o *Transaction) SetSourceDbname(v string)`

SetSourceDbname sets SourceDbname field to given value.

### HasSourceDbname

`func (o *Transaction) HasSourceDbname() bool`

HasSourceDbname returns a boolean if a field has been set.

### SetSourceDbnameNil

`func (o *Transaction) SetSourceDbnameNil(b bool)`

 SetSourceDbnameNil sets the value for SourceDbname to be an explicit nil

### UnsetSourceDbname
`func (o *Transaction) UnsetSourceDbname()`

UnsetSourceDbname ensures that no value is present for SourceDbname, not even an explicit nil
### GetVersion

`func (o *Transaction) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Transaction) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Transaction) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *Transaction) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### SetVersionNil

`func (o *Transaction) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *Transaction) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


