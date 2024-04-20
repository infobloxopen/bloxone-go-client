# IpamHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]HostAddress**](HostAddress.md) | The list of all addresses associated with the IPAM host, which may be in different IP spaces. | [optional] 
**AutoGenerateRecords** | Pointer to **bool** | This flag specifies if resource records have to be auto generated for the host. | [optional] 
**Comment** | Pointer to **string** | The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**HostNames** | Pointer to [**[]HostName**](HostName.md) | The name records to be generated for the host.  This field is required if _auto_generate_records_ is true. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the IPAM host in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewIpamHost

`func NewIpamHost(name string, ) *IpamHost`

NewIpamHost instantiates a new IpamHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamHostWithDefaults

`func NewIpamHostWithDefaults() *IpamHost`

NewIpamHostWithDefaults instantiates a new IpamHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IpamHost) GetAddresses() []HostAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IpamHost) GetAddressesOk() (*[]HostAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IpamHost) SetAddresses(v []HostAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IpamHost) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAutoGenerateRecords

`func (o *IpamHost) GetAutoGenerateRecords() bool`

GetAutoGenerateRecords returns the AutoGenerateRecords field if non-nil, zero value otherwise.

### GetAutoGenerateRecordsOk

`func (o *IpamHost) GetAutoGenerateRecordsOk() (*bool, bool)`

GetAutoGenerateRecordsOk returns a tuple with the AutoGenerateRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGenerateRecords

`func (o *IpamHost) SetAutoGenerateRecords(v bool)`

SetAutoGenerateRecords sets AutoGenerateRecords field to given value.

### HasAutoGenerateRecords

`func (o *IpamHost) HasAutoGenerateRecords() bool`

HasAutoGenerateRecords returns a boolean if a field has been set.

### GetComment

`func (o *IpamHost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamHost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamHost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamHost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamHost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamHost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamHost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamHost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHostNames

`func (o *IpamHost) GetHostNames() []HostName`

GetHostNames returns the HostNames field if non-nil, zero value otherwise.

### GetHostNamesOk

`func (o *IpamHost) GetHostNamesOk() (*[]HostName, bool)`

GetHostNamesOk returns a tuple with the HostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNames

`func (o *IpamHost) SetHostNames(v []HostName)`

SetHostNames sets HostNames field to given value.

### HasHostNames

`func (o *IpamHost) HasHostNames() bool`

HasHostNames returns a boolean if a field has been set.

### GetId

`func (o *IpamHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamHost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamHost) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *IpamHost) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamHost) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamHost) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamHost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamHost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamHost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamHost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamHost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


