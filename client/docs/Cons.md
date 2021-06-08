# Cons

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Head** | [**SyntaxNode**](SyntaxNode.md) |  | 
**Tail** | [**LinkedList**](LinkedList.md) |  | 

## Methods

### NewCons

`func NewCons(head SyntaxNode, tail LinkedList, ) *Cons`

NewCons instantiates a new Cons object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsWithDefaults

`func NewConsWithDefaults() *Cons`

NewConsWithDefaults instantiates a new Cons object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHead

`func (o *Cons) GetHead() SyntaxNode`

GetHead returns the Head field if non-nil, zero value otherwise.

### GetHeadOk

`func (o *Cons) GetHeadOk() (*SyntaxNode, bool)`

GetHeadOk returns a tuple with the Head field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHead

`func (o *Cons) SetHead(v SyntaxNode)`

SetHead sets Head field to given value.


### GetTail

`func (o *Cons) GetTail() LinkedList`

GetTail returns the Tail field if non-nil, zero value otherwise.

### GetTailOk

`func (o *Cons) GetTailOk() (*LinkedList, bool)`

GetTailOk returns a tuple with the Tail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTail

`func (o *Cons) SetTail(v LinkedList)`

SetTail sets Tail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


