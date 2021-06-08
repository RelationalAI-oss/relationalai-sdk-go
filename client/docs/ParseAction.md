# ParseAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonterm** | **string** |  | [default to ""]
**Source** | [**Source**](Source.md) |  | 

## Methods

### NewParseAction

`func NewParseAction(nonterm string, source Source, ) *ParseAction`

NewParseAction instantiates a new ParseAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseActionWithDefaults

`func NewParseActionWithDefaults() *ParseAction`

NewParseActionWithDefaults instantiates a new ParseAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNonterm

`func (o *ParseAction) GetNonterm() string`

GetNonterm returns the Nonterm field if non-nil, zero value otherwise.

### GetNontermOk

`func (o *ParseAction) GetNontermOk() (*string, bool)`

GetNontermOk returns a tuple with the Nonterm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonterm

`func (o *ParseAction) SetNonterm(v string)`

SetNonterm sets Nonterm field to given value.


### GetSource

`func (o *ParseAction) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ParseAction) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ParseAction) SetSource(v Source)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


