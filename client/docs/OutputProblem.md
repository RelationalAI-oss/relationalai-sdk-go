# OutputProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exception** | **string** |  | [default to ""]
**ExceptionStacktrace** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | [default to ""]

## Methods

### NewOutputProblem

`func NewOutputProblem(exception string, name string, ) *OutputProblem`

NewOutputProblem instantiates a new OutputProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputProblemWithDefaults

`func NewOutputProblemWithDefaults() *OutputProblem`

NewOutputProblemWithDefaults instantiates a new OutputProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetException

`func (o *OutputProblem) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *OutputProblem) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *OutputProblem) SetException(v string)`

SetException sets Exception field to given value.


### GetExceptionStacktrace

`func (o *OutputProblem) GetExceptionStacktrace() string`

GetExceptionStacktrace returns the ExceptionStacktrace field if non-nil, zero value otherwise.

### GetExceptionStacktraceOk

`func (o *OutputProblem) GetExceptionStacktraceOk() (*string, bool)`

GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStacktrace

`func (o *OutputProblem) SetExceptionStacktrace(v string)`

SetExceptionStacktrace sets ExceptionStacktrace field to given value.

### HasExceptionStacktrace

`func (o *OutputProblem) HasExceptionStacktrace() bool`

HasExceptionStacktrace returns a boolean if a field has been set.

### SetExceptionStacktraceNil

`func (o *OutputProblem) SetExceptionStacktraceNil(b bool)`

 SetExceptionStacktraceNil sets the value for ExceptionStacktrace to be an explicit nil

### UnsetExceptionStacktrace
`func (o *OutputProblem) UnsetExceptionStacktrace()`

UnsetExceptionStacktrace ensures that no value is present for ExceptionStacktrace, not even an explicit nil
### GetName

`func (o *OutputProblem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OutputProblem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OutputProblem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


