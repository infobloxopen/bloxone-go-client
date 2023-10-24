# IpamsvcIpamHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]IpamsvcHostAddress**](IpamsvcHostAddress.md) | The list of all addresses associated with the IPAM host, which may be in different IP spaces. | [optional] 
**AutoGenerateRecords** | Pointer to **bool** | This flag specifies if resource records have to be auto generated for the host. | [optional] 
**Comment** | Pointer to **string** | The description for the IPAM host. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**HostNames** | Pointer to [**[]IpamsvcHostName**](IpamsvcHostName.md) | The name records to be generated for the host.  This field is required if _auto_generate_records_ is true. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the IPAM host. Must contain 1 to 256 characters. Can include UTF-8. | 
**Tags** | Pointer to **map[string]interface{}** | The tags for the IPAM host in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewIpamsvcIpamHost

`func NewIpamsvcIpamHost(name string, ) *IpamsvcIpamHost`

NewIpamsvcIpamHost instantiates a new IpamsvcIpamHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcIpamHostWithDefaults

`func NewIpamsvcIpamHostWithDefaults() *IpamsvcIpamHost`

NewIpamsvcIpamHostWithDefaults instantiates a new IpamsvcIpamHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IpamsvcIpamHost) GetAddresses() []IpamsvcHostAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IpamsvcIpamHost) GetAddressesOk() (*[]IpamsvcHostAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IpamsvcIpamHost) SetAddresses(v []IpamsvcHostAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IpamsvcIpamHost) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetAutoGenerateRecords

`func (o *IpamsvcIpamHost) GetAutoGenerateRecords() bool`

GetAutoGenerateRecords returns the AutoGenerateRecords field if non-nil, zero value otherwise.

### GetAutoGenerateRecordsOk

`func (o *IpamsvcIpamHost) GetAutoGenerateRecordsOk() (*bool, bool)`

GetAutoGenerateRecordsOk returns a tuple with the AutoGenerateRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoGenerateRecords

`func (o *IpamsvcIpamHost) SetAutoGenerateRecords(v bool)`

SetAutoGenerateRecords sets AutoGenerateRecords field to given value.

### HasAutoGenerateRecords

`func (o *IpamsvcIpamHost) HasAutoGenerateRecords() bool`

HasAutoGenerateRecords returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcIpamHost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcIpamHost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcIpamHost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcIpamHost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcIpamHost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcIpamHost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcIpamHost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcIpamHost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHostNames

`func (o *IpamsvcIpamHost) GetHostNames() []IpamsvcHostName`

GetHostNames returns the HostNames field if non-nil, zero value otherwise.

### GetHostNamesOk

`func (o *IpamsvcIpamHost) GetHostNamesOk() (*[]IpamsvcHostName, bool)`

GetHostNamesOk returns a tuple with the HostNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNames

`func (o *IpamsvcIpamHost) SetHostNames(v []IpamsvcHostName)`

SetHostNames sets HostNames field to given value.

### HasHostNames

`func (o *IpamsvcIpamHost) HasHostNames() bool`

HasHostNames returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcIpamHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcIpamHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcIpamHost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcIpamHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcIpamHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcIpamHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcIpamHost) SetName(v string)`

SetName sets Name field to given value.


### GetTags

`func (o *IpamsvcIpamHost) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcIpamHost) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcIpamHost) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcIpamHost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcIpamHost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcIpamHost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcIpamHost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcIpamHost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


