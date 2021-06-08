# SyntaxError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Next** | Pointer to [**SyntaxNode**](SyntaxNode.md) |  | [optional] 
**Node** | [**SyntaxNode**](SyntaxNode.md) |  | 
**Trace** | [**LinkedList**](LinkedList.md) |  | 

## Methods

### NewSyntaxError

`func NewSyntaxError(node SyntaxNode, trace LinkedList, ) *SyntaxError`

NewSyntaxError instantiates a new SyntaxError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntaxErrorWithDefaults

`func NewSyntaxErrorWithDefaults() *SyntaxError`

NewSyntaxErrorWithDefaults instantiates a new SyntaxError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNext

`func (o *SyntaxError) GetNext() SyntaxNode`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *SyntaxError) GetNextOk() (*SyntaxNode, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *SyntaxError) SetNext(v SyntaxNode)`

SetNext sets Next field to given value.

### HasNext

`func (o *SyntaxError) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetNode

`func (o *SyntaxError) GetNode() SyntaxNode`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *SyntaxError) GetNodeOk() (*SyntaxNode, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *SyntaxError) SetNode(v SyntaxNode)`

SetNode sets Node field to given value.


### GetTrace

`func (o *SyntaxError) GetTrace() LinkedList`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *SyntaxError) GetTraceOk() (*LinkedList, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *SyntaxError) SetTrace(v LinkedList)`

SetTrace sets Trace field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


