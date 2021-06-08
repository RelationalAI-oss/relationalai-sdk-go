# OutputProblemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exception** | **string** |  | [default to ""]
**ExceptionStacktrace** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | [default to ""]

## Methods

### NewOutputProblemAllOf

`func NewOutputProblemAllOf(exception string, name string, ) *OutputProblemAllOf`

NewOutputProblemAllOf instantiates a new OutputProblemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProblemAllOfWithDefaults

`func NewOutputProblemAllOfWithDefaults() *OutputProblemAllOf`

NewOutputProblemAllOfWithDefaults instantiates a new OutputProblemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetException

`func (o *OutputProblemAllOf) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *OutputProblemAllOf) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *OutputProblemAllOf) SetException(v string)`

SetException sets Exception field to given value.


### GetExceptionStacktrace

`func (o *OutputProblemAllOf) GetExceptionStacktrace() string`

GetExceptionStacktrace returns the ExceptionStacktrace field if non-nil, zero value otherwise.

### GetExceptionStacktraceOk

`func (o *OutputProblemAllOf) GetExceptionStacktraceOk() (*string, bool)`

GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStacktrace

`func (o *OutputProblemAllOf) SetExceptionStacktrace(v string)`

SetExceptionStacktrace sets ExceptionStacktrace field to given value.

### HasExceptionStacktrace

`func (o *OutputProblemAllOf) HasExceptionStacktrace() bool`

HasExceptionStacktrace returns a boolean if a field has been set.

### SetExceptionStacktraceNil

`func (o *OutputProblemAllOf) SetExceptionStacktraceNil(b bool)`

 SetExceptionStacktraceNil sets the value for ExceptionStacktrace to be an explicit nil

### UnsetExceptionStacktrace
`func (o *OutputProblemAllOf) UnsetExceptionStacktrace()`

UnsetExceptionStacktrace ensures that no value is present for ExceptionStacktrace, not even an explicit nil
### GetName

`func (o *OutputProblemAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputProblemAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputProblemAllOf) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


