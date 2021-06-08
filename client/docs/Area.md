# Area

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndPoint** | [**Point**](Point.md) |  | 
**StartPoint** | [**Point**](Point.md) |  | 
**Type** | **string** |  | [default to "Area"]

## Methods

### NewArea

`func NewArea(endPoint Point, startPoint Point, type_ string, ) *Area`

NewArea instantiates a new Area object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAreaWithDefaults

`func NewAreaWithDefaults() *Area`

NewAreaWithDefaults instantiates a new Area object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndPoint

`func (o *Area) GetEndPoint() Point`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *Area) GetEndPointOk() (*Point, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *Area) SetEndPoint(v Point)`

SetEndPoint sets EndPoint field to given value.


### GetStartPoint

`func (o *Area) GetStartPoint() Point`

GetStartPoint returns the StartPoint field if non-nil, zero value otherwise.

### GetStartPointOk

`func (o *Area) GetStartPointOk() (*Point, bool)`

GetStartPointOk returns a tuple with the StartPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPoint

`func (o *Area) SetStartPoint(v Point)`

SetStartPoint sets StartPoint field to given value.


### GetType

`func (o *Area) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Area) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Area) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


