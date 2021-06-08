# ParseActionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Problems** | Pointer to [**[]AbstractProblem**](AbstractProblem.md) |  | [optional] 

## Methods

### NewParseActionResult

`func NewParseActionResult() *ParseActionResult`

NewParseActionResult instantiates a new ParseActionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseActionResultWithDefaults

`func NewParseActionResultWithDefaults() *ParseActionResult`

NewParseActionResultWithDefaults instantiates a new ParseActionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblems

`func (o *ParseActionResult) GetProblems() []AbstractProblem`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *ParseActionResult) GetProblemsOk() (*[]AbstractProblem, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *ParseActionResult) SetProblems(v []AbstractProblem)`

SetProblems sets Problems field to given value.

### HasProblems

`func (o *ParseActionResult) HasProblems() bool`

HasProblems returns a boolean if a field has been set.

### SetProblemsNil

`func (o *ParseActionResult) SetProblemsNil(b bool)`

 SetProblemsNil sets the value for Problems to be an explicit nil

### UnsetProblems
`func (o *ParseActionResult) UnsetProblems()`

UnsetProblems ensures that no value is present for Problems, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


