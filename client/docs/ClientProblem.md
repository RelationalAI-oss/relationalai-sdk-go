# ClientProblem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | [default to ""]
**IsError** | **bool** |  | [default to false]
**IsException** | **bool** |  | [default to false]
**Message** | **string** |  | [default to ""]
**Path** | **string** |  | [default to ""]
**Report** | **string** |  | [default to ""]

## Methods

### NewClientProblem

`func NewClientProblem(errorCode string, isError bool, isException bool, message string, path string, report string, ) *ClientProblem`

NewClientProblem instantiates a new ClientProblem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientProblemWithDefaults

`func NewClientProblemWithDefaults() *ClientProblem`

NewClientProblemWithDefaults instantiates a new ClientProblem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *ClientProblem) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ClientProblem) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ClientProblem) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetIsError

`func (o *ClientProblem) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *ClientProblem) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *ClientProblem) SetIsError(v bool)`

SetIsError sets IsError field to given value.


### GetIsException

`func (o *ClientProblem) GetIsException() bool`

GetIsException returns the IsException field if non-nil, zero value otherwise.

### GetIsExceptionOk

`func (o *ClientProblem) GetIsExceptionOk() (*bool, bool)`

GetIsExceptionOk returns a tuple with the IsException field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsException

`func (o *ClientProblem) SetIsException(v bool)`

SetIsException sets IsException field to given value.


### GetMessage

`func (o *ClientProblem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClientProblem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClientProblem) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPath

`func (o *ClientProblem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ClientProblem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ClientProblem) SetPath(v string)`

SetPath sets Path field to given value.


### GetReport

`func (o *ClientProblem) GetReport() string`

GetReport returns the Report field if non-nil, zero value otherwise.

### GetReportOk

`func (o *ClientProblem) GetReportOk() (*string, bool)`

GetReportOk returns a tuple with the Report field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReport

`func (o *ClientProblem) SetReport(v string)`

SetReport sets Report field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


