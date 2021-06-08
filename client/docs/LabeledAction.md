# LabeledAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**Action**](Action.md) |  | 
**Name** | **string** |  | [default to ""]
**Type** | **string** |  | [default to "LabeledAction"]

## Methods

### NewLabeledAction

`func NewLabeledAction(action Action, name string, type_ string, ) *LabeledAction`

NewLabeledAction instantiates a new LabeledAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabeledActionWithDefaults

`func NewLabeledActionWithDefaults() *LabeledAction`

NewLabeledActionWithDefaults instantiates a new LabeledAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *LabeledAction) GetAction() Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *LabeledAction) GetActionOk() (*Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *LabeledAction) SetAction(v Action)`

SetAction sets Action field to given value.


### GetName

`func (o *LabeledAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LabeledAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LabeledAction) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *LabeledAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabeledAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabeledAction) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


