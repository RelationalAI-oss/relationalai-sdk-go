# Range

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | [**Area**](Area.md) |  | 
**EndByte** | **int32** |  | [default to 0]
**Input** | Pointer to **interface{}** |  | [optional] 
**StartByte** | **int32** |  | [default to 0]
**Type** | **string** |  | [default to "Range"]

## Methods

### NewRange

`func NewRange(area Area, endByte int32, startByte int32, type_ string, ) *Range`

NewRange instantiates a new Range object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRangeWithDefaults

`func NewRangeWithDefaults() *Range`

NewRangeWithDefaults instantiates a new Range object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *Range) GetArea() Area`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *Range) GetAreaOk() (*Area, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *Range) SetArea(v Area)`

SetArea sets Area field to given value.


### GetEndByte

`func (o *Range) GetEndByte() int32`

GetEndByte returns the EndByte field if non-nil, zero value otherwise.

### GetEndByteOk

`func (o *Range) GetEndByteOk() (*int32, bool)`

GetEndByteOk returns a tuple with the EndByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndByte

`func (o *Range) SetEndByte(v int32)`

SetEndByte sets EndByte field to given value.


### GetInput

`func (o *Range) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Range) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Range) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *Range) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *Range) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *Range) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetStartByte

`func (o *Range) GetStartByte() int32`

GetStartByte returns the StartByte field if non-nil, zero value otherwise.

### GetStartByteOk

`func (o *Range) GetStartByteOk() (*int32, bool)`

GetStartByteOk returns a tuple with the StartByte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartByte

`func (o *Range) SetStartByte(v int32)`

SetStartByte sets StartByte field to given value.


### GetType

`func (o *Range) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Range) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Range) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


