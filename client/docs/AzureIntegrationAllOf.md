# AzureIntegrationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]PairSymbolString**](PairSymbolString.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**StorageAllowedLocations** | Pointer to **[]string** |  | [optional] 
**StorageBlockedLocations** | Pointer to **[]string** |  | [optional] 
**TenantId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAzureIntegrationAllOf

`func NewAzureIntegrationAllOf() *AzureIntegrationAllOf`

NewAzureIntegrationAllOf instantiates a new AzureIntegrationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureIntegrationAllOfWithDefaults

`func NewAzureIntegrationAllOfWithDefaults() *AzureIntegrationAllOf`

NewAzureIntegrationAllOfWithDefaults instantiates a new AzureIntegrationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *AzureIntegrationAllOf) GetCredentials() []PairSymbolString`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *AzureIntegrationAllOf) GetCredentialsOk() (*[]PairSymbolString, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *AzureIntegrationAllOf) SetCredentials(v []PairSymbolString)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *AzureIntegrationAllOf) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### SetCredentialsNil

`func (o *AzureIntegrationAllOf) SetCredentialsNil(b bool)`

 SetCredentialsNil sets the value for Credentials to be an explicit nil

### UnsetCredentials
`func (o *AzureIntegrationAllOf) UnsetCredentials()`

UnsetCredentials ensures that no value is present for Credentials, not even an explicit nil
### GetName

`func (o *AzureIntegrationAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AzureIntegrationAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AzureIntegrationAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AzureIntegrationAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *AzureIntegrationAllOf) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AzureIntegrationAllOf) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStorageAllowedLocations

`func (o *AzureIntegrationAllOf) GetStorageAllowedLocations() []string`

GetStorageAllowedLocations returns the StorageAllowedLocations field if non-nil, zero value otherwise.

### GetStorageAllowedLocationsOk

`func (o *AzureIntegrationAllOf) GetStorageAllowedLocationsOk() (*[]string, bool)`

GetStorageAllowedLocationsOk returns a tuple with the StorageAllowedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAllowedLocations

`func (o *AzureIntegrationAllOf) SetStorageAllowedLocations(v []string)`

SetStorageAllowedLocations sets StorageAllowedLocations field to given value.

### HasStorageAllowedLocations

`func (o *AzureIntegrationAllOf) HasStorageAllowedLocations() bool`

HasStorageAllowedLocations returns a boolean if a field has been set.

### SetStorageAllowedLocationsNil

`func (o *AzureIntegrationAllOf) SetStorageAllowedLocationsNil(b bool)`

 SetStorageAllowedLocationsNil sets the value for StorageAllowedLocations to be an explicit nil

### UnsetStorageAllowedLocations
`func (o *AzureIntegrationAllOf) UnsetStorageAllowedLocations()`

UnsetStorageAllowedLocations ensures that no value is present for StorageAllowedLocations, not even an explicit nil
### GetStorageBlockedLocations

`func (o *AzureIntegrationAllOf) GetStorageBlockedLocations() []string`

GetStorageBlockedLocations returns the StorageBlockedLocations field if non-nil, zero value otherwise.

### GetStorageBlockedLocationsOk

`func (o *AzureIntegrationAllOf) GetStorageBlockedLocationsOk() (*[]string, bool)`

GetStorageBlockedLocationsOk returns a tuple with the StorageBlockedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBlockedLocations

`func (o *AzureIntegrationAllOf) SetStorageBlockedLocations(v []string)`

SetStorageBlockedLocations sets StorageBlockedLocations field to given value.

### HasStorageBlockedLocations

`func (o *AzureIntegrationAllOf) HasStorageBlockedLocations() bool`

HasStorageBlockedLocations returns a boolean if a field has been set.

### SetStorageBlockedLocationsNil

`func (o *AzureIntegrationAllOf) SetStorageBlockedLocationsNil(b bool)`

 SetStorageBlockedLocationsNil sets the value for StorageBlockedLocations to be an explicit nil

### UnsetStorageBlockedLocations
`func (o *AzureIntegrationAllOf) UnsetStorageBlockedLocations()`

UnsetStorageBlockedLocations ensures that no value is present for StorageBlockedLocations, not even an explicit nil
### GetTenantId

`func (o *AzureIntegrationAllOf) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureIntegrationAllOf) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureIntegrationAllOf) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AzureIntegrationAllOf) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### SetTenantIdNil

`func (o *AzureIntegrationAllOf) SetTenantIdNil(b bool)`

 SetTenantIdNil sets the value for TenantId to be an explicit nil

### UnsetTenantId
`func (o *AzureIntegrationAllOf) UnsetTenantId()`

UnsetTenantId ensures that no value is present for TenantId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


