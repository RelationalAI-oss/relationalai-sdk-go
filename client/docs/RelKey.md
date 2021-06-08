# RelKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keys** | Pointer to **[]string** |  | [optional] 
**Name** | **string** |  | [default to ""]
**Values** | Pointer to **[]string** |  | [optional] 
**Type** | **string** |  | [default to "RelKey"]

## Methods

### NewRelKey

`func NewRelKey(name string, type_ string, ) *RelKey`

NewRelKey instantiates a new RelKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelKeyWithDefaults

`func NewRelKeyWithDefaults() *RelKey`

NewRelKeyWithDefaults instantiates a new RelKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeys

`func (o *RelKey) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *RelKey) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *RelKey) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *RelKey) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *RelKey) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *RelKey) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil
### GetName

`func (o *RelKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RelKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RelKey) SetName(v string)`

SetName sets Name field to given value.


### GetValues

`func (o *RelKey) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RelKey) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RelKey) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RelKey) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *RelKey) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *RelKey) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil
### GetType

`func (o *RelKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RelKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RelKey) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


