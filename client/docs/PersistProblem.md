# PersistProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exception** | **string** |  | [default to ""]
**ExceptionStacktrace** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPersistProblem

`func NewPersistProblem(exception string, ) *PersistProblem`

NewPersistProblem instantiates a new PersistProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersistProblemWithDefaults

`func NewPersistProblemWithDefaults() *PersistProblem`

NewPersistProblemWithDefaults instantiates a new PersistProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetException

`func (o *PersistProblem) GetException() string`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *PersistProblem) GetExceptionOk() (*string, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *PersistProblem) SetException(v string)`

SetException sets Exception field to given value.


### GetExceptionStacktrace

`func (o *PersistProblem) GetExceptionStacktrace() string`

GetExceptionStacktrace returns the ExceptionStacktrace field if non-nil, zero value otherwise.

### GetExceptionStacktraceOk

`func (o *PersistProblem) GetExceptionStacktraceOk() (*string, bool)`

GetExceptionStacktraceOk returns a tuple with the ExceptionStacktrace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionStacktrace

`func (o *PersistProblem) SetExceptionStacktrace(v string)`

SetExceptionStacktrace sets ExceptionStacktrace field to given value.

### HasExceptionStacktrace

`func (o *PersistProblem) HasExceptionStacktrace() bool`

HasExceptionStacktrace returns a boolean if a field has been set.

### SetExceptionStacktraceNil

`func (o *PersistProblem) SetExceptionStacktraceNil(b bool)`

 SetExceptionStacktraceNil sets the value for ExceptionStacktrace to be an explicit nil

### UnsetExceptionStacktrace
`func (o *PersistProblem) UnsetExceptionStacktrace()`

UnsetExceptionStacktrace ensures that no value is present for ExceptionStacktrace, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


