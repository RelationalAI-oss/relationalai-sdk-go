# LabeledActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to ""]
**Result** | [**ActionResult**](ActionResult.md) |  | 
**Type** | **string** |  | [default to "LabeledActionResult"]

## Methods

### NewLabeledActionResult

`func NewLabeledActionResult(name string, result ActionResult, type_ string, ) *LabeledActionResult`

NewLabeledActionResult instantiates a new LabeledActionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabeledActionResultWithDefaults

`func NewLabeledActionResultWithDefaults() *LabeledActionResult`

NewLabeledActionResultWithDefaults instantiates a new LabeledActionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LabeledActionResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabeledActionResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabeledActionResult) SetName(v string)`

SetName sets Name field to given value.


### GetResult

`func (o *LabeledActionResult) GetResult() ActionResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *LabeledActionResult) GetResultOk() (*ActionResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *LabeledActionResult) SetResult(v ActionResult)`

SetResult sets Result field to given value.


### GetType

`func (o *LabeledActionResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabeledActionResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabeledActionResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


