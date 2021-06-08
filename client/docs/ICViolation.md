# ICViolation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelKey** | [**RelKey**](RelKey.md) |  | 
**Source** | **string** |  | [default to ""]
**Type** | **string** |  | [default to "ICViolation"]

## Methods

### NewICViolation

`func NewICViolation(relKey RelKey, source string, type_ string, ) *ICViolation`

NewICViolation instantiates a new ICViolation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewICViolationWithDefaults

`func NewICViolationWithDefaults() *ICViolation`

NewICViolationWithDefaults instantiates a new ICViolation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelKey

`func (o *ICViolation) GetRelKey() RelKey`

GetRelKey returns the RelKey field if non-nil, zero value otherwise.

### GetRelKeyOk

`func (o *ICViolation) GetRelKeyOk() (*RelKey, bool)`

GetRelKeyOk returns a tuple with the RelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelKey

`func (o *ICViolation) SetRelKey(v RelKey)`

SetRelKey sets RelKey field to given value.


### GetSource

`func (o *ICViolation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ICViolation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ICViolation) SetSource(v string)`

SetSource sets Source field to given value.


### GetType

`func (o *ICViolation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ICViolation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ICViolation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


