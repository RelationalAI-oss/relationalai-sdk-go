# ExceptionProblemAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exception** | **string** |  | [default to ""]
**ExceptionStacktrace** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewExceptionProblemAllOf

`func NewExceptionProblemAllOf(exception string, ) *ExceptionProblemAllOf`

NewExceptionProblemAllOf instantiates a new ExceptionProblemAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionProblemAllOfWithDefaults

`func NewExceptionProblemAllOfWithDefaults() *ExceptionProblemAllOf`

NewExceptionProblemAllOfWithDefaults instantiates a new ExceptionProblemAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetException

`func (o *ExceptionProblemAllOf) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *ExceptionProblemAllOf) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *ExceptionProblemAllOf) SetException(v string)`

SetException sets Exception field to given value.


### GetExceptionStacktrace

`func (o *ExceptionProblemAllOf) GetExceptionStacktrace() string`

GetExceptionStacktrace returns the ExceptionStacktrace field if non-nil, zero value otherwise.

### GetExceptionStacktraceOk

`func (o *ExceptionProblemAllOf) GetExceptionStacktraceOk() (*string, bool)`

GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStacktrace

`func (o *ExceptionProblemAllOf) SetExceptionStacktrace(v string)`

SetExceptionStacktrace sets ExceptionStacktrace field to given value.

### HasExceptionStacktrace

`func (o *ExceptionProblemAllOf) HasExceptionStacktrace() bool`

HasExceptionStacktrace returns a boolean if a field has been set.

### SetExceptionStacktraceNil

`func (o *ExceptionProblemAllOf) SetExceptionStacktraceNil(b bool)`

 SetExceptionStacktraceNil sets the value for ExceptionStacktrace to be an explicit nil

### UnsetExceptionStacktrace
`func (o *ExceptionProblemAllOf) UnsetExceptionStacktrace()`

UnsetExceptionStacktrace ensures that no value is present for ExceptionStacktrace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


