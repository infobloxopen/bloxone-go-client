# InternalDomains

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this Internal Domain lists object was created. | [optional] [readonly] 
**Description** | Pointer to **string** | The brief description for the internal domain lists . | [optional] 
**Id** | Pointer to **int32** | The Internal Domain object identifier. | [optional] [readonly] 
**InternalDomains** | Pointer to **[]string** | The list of internal domains, should be unique to each other and has to be read-only from the API level. | [optional] 
**IsDefault** | Pointer to **bool** | True if name is &#39;Default Bypass Domains/CIDRs&#39; otherwise false. | [optional] 
**Name** | Pointer to **string** | The name of the internal domain lists. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | Enables tag support for resource where tags attribute contains user-defined key value pairs | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Internal domain lists object was last updated. | [optional] [readonly] 

## Methods

### NewInternalDomains

`func NewInternalDomains() *InternalDomains`

NewInternalDomains instantiates a new InternalDomains object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalDomainsWithDefaults

`func NewInternalDomainsWithDefaults() *InternalDomains`

NewInternalDomainsWithDefaults instantiates a new InternalDomains object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *InternalDomains) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *InternalDomains) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *InternalDomains) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *InternalDomains) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *InternalDomains) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InternalDomains) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InternalDomains) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InternalDomains) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *InternalDomains) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalDomains) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalDomains) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalDomains) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalDomains

`func (o *InternalDomains) GetInternalDomains() []string`

GetInternalDomains returns the InternalDomains field if non-nil, zero value otherwise.

### GetInternalDomainsOk

`func (o *InternalDomains) GetInternalDomainsOk() (*[]string, bool)`

GetInternalDomainsOk returns a tuple with the InternalDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalDomains

`func (o *InternalDomains) SetInternalDomains(v []string)`

SetInternalDomains sets InternalDomains field to given value.

### HasInternalDomains

`func (o *InternalDomains) HasInternalDomains() bool`

HasInternalDomains returns a boolean if a field has been set.

### GetIsDefault

`func (o *InternalDomains) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *InternalDomains) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *InternalDomains) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *InternalDomains) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *InternalDomains) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalDomains) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalDomains) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalDomains) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *InternalDomains) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InternalDomains) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InternalDomains) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *InternalDomains) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *InternalDomains) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *InternalDomains) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *InternalDomains) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *InternalDomains) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


