# AnycastConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **int64** |  | [optional] 
**AnycastIpAddress** | Pointer to **string** |  | [optional] 
**AnycastIpv6Address** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Fields** | Pointer to [**BufFieldMask**](BufFieldMask.md) |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IsConfigured** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OnpremHosts** | Pointer to [**[]OnpremHostRef**](OnpremHostRef.md) |  | [optional] 
**RuntimeStatus** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAnycastConfig

`func NewAnycastConfig() *AnycastConfig`

NewAnycastConfig instantiates a new AnycastConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnycastConfigWithDefaults

`func NewAnycastConfigWithDefaults() *AnycastConfig`

NewAnycastConfigWithDefaults instantiates a new AnycastConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AnycastConfig) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AnycastConfig) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AnycastConfig) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AnycastConfig) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAnycastIpAddress

`func (o *AnycastConfig) GetAnycastIpAddress() string`

GetAnycastIpAddress returns the AnycastIpAddress field if non-nil, zero value otherwise.

### GetAnycastIpAddressOk

`func (o *AnycastConfig) GetAnycastIpAddressOk() (*string, bool)`

GetAnycastIpAddressOk returns a tuple with the AnycastIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastIpAddress

`func (o *AnycastConfig) SetAnycastIpAddress(v string)`

SetAnycastIpAddress sets AnycastIpAddress field to given value.

### HasAnycastIpAddress

`func (o *AnycastConfig) HasAnycastIpAddress() bool`

HasAnycastIpAddress returns a boolean if a field has been set.

### GetAnycastIpv6Address

`func (o *AnycastConfig) GetAnycastIpv6Address() string`

GetAnycastIpv6Address returns the AnycastIpv6Address field if non-nil, zero value otherwise.

### GetAnycastIpv6AddressOk

`func (o *AnycastConfig) GetAnycastIpv6AddressOk() (*string, bool)`

GetAnycastIpv6AddressOk returns a tuple with the AnycastIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastIpv6Address

`func (o *AnycastConfig) SetAnycastIpv6Address(v string)`

SetAnycastIpv6Address sets AnycastIpv6Address field to given value.

### HasAnycastIpv6Address

`func (o *AnycastConfig) HasAnycastIpv6Address() bool`

HasAnycastIpv6Address returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnycastConfig) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnycastConfig) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnycastConfig) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AnycastConfig) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *AnycastConfig) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnycastConfig) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnycastConfig) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnycastConfig) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFields

`func (o *AnycastConfig) GetFields() BufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AnycastConfig) GetFieldsOk() (*BufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AnycastConfig) SetFields(v BufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AnycastConfig) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetId

`func (o *AnycastConfig) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnycastConfig) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnycastConfig) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AnycastConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsConfigured

`func (o *AnycastConfig) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *AnycastConfig) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *AnycastConfig) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *AnycastConfig) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetName

`func (o *AnycastConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnycastConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnycastConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AnycastConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnpremHosts

`func (o *AnycastConfig) GetOnpremHosts() []OnpremHostRef`

GetOnpremHosts returns the OnpremHosts field if non-nil, zero value otherwise.

### GetOnpremHostsOk

`func (o *AnycastConfig) GetOnpremHostsOk() (*[]OnpremHostRef, bool)`

GetOnpremHostsOk returns a tuple with the OnpremHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnpremHosts

`func (o *AnycastConfig) SetOnpremHosts(v []OnpremHostRef)`

SetOnpremHosts sets OnpremHosts field to given value.

### HasOnpremHosts

`func (o *AnycastConfig) HasOnpremHosts() bool`

HasOnpremHosts returns a boolean if a field has been set.

### GetRuntimeStatus

`func (o *AnycastConfig) GetRuntimeStatus() string`

GetRuntimeStatus returns the RuntimeStatus field if non-nil, zero value otherwise.

### GetRuntimeStatusOk

`func (o *AnycastConfig) GetRuntimeStatusOk() (*string, bool)`

GetRuntimeStatusOk returns a tuple with the RuntimeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeStatus

`func (o *AnycastConfig) SetRuntimeStatus(v string)`

SetRuntimeStatus sets RuntimeStatus field to given value.

### HasRuntimeStatus

`func (o *AnycastConfig) HasRuntimeStatus() bool`

HasRuntimeStatus returns a boolean if a field has been set.

### GetService

`func (o *AnycastConfig) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AnycastConfig) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AnycastConfig) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AnycastConfig) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTags

`func (o *AnycastConfig) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AnycastConfig) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AnycastConfig) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *AnycastConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AnycastConfig) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AnycastConfig) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AnycastConfig) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AnycastConfig) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


