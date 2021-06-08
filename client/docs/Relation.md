# Relation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to **[][]interface{}** |  | [optional] 
**RelKey** | [**RelKey**](RelKey.md) |  | 
**Type** | **string** |  | [default to "Relation"]

## Methods

### NewRelation

`func NewRelation(relKey RelKey, type_ string, ) *Relation`

NewRelation instantiates a new Relation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelationWithDefaults

`func NewRelationWithDefaults() *Relation`

NewRelationWithDefaults instantiates a new Relation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *Relation) GetColumns() [][]interface{}`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *Relation) GetColumnsOk() (*[][]interface{}, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *Relation) SetColumns(v [][]interface{})`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *Relation) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### SetColumnsNil

`func (o *Relation) SetColumnsNil(b bool)`

 SetColumnsNil sets the value for Columns to be an explicit nil

### UnsetColumns
`func (o *Relation) UnsetColumns()`

UnsetColumns ensures that no value is present for Columns, not even an explicit nil
### GetRelKey

`func (o *Relation) GetRelKey() RelKey`

GetRelKey returns the RelKey field if non-nil, zero value otherwise.

### GetRelKeyOk

`func (o *Relation) GetRelKeyOk() (*RelKey, bool)`

GetRelKeyOk returns a tuple with the RelKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelKey

`func (o *Relation) SetRelKey(v RelKey)`

SetRelKey sets RelKey field to given value.


### GetType

`func (o *Relation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Relation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Relation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


