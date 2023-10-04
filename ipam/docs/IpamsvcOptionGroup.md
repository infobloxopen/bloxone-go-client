# IpamsvcOptionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the option group. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Time when the object has been created. | [optional] [readonly] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of DHCP options for the option group. May be either a specific option or a group of options. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | **string** | The name of the option group. Must contain 1 to 256 characters. Can include UTF-8. | 
**Protocol** | Pointer to **string** | The type of protocol (_ip4_ or _ip6_). | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the option group in JSON format. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewIpamsvcOptionGroup

`func NewIpamsvcOptionGroup(name string, ) *IpamsvcOptionGroup`

NewIpamsvcOptionGroup instantiates a new IpamsvcOptionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcOptionGroupWithDefaults

`func NewIpamsvcOptionGroupWithDefaults() *IpamsvcOptionGroup`

NewIpamsvcOptionGroupWithDefaults instantiates a new IpamsvcOptionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *IpamsvcOptionGroup) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcOptionGroup) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcOptionGroup) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcOptionGroup) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcOptionGroup) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcOptionGroup) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcOptionGroup) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcOptionGroup) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IpamsvcOptionGroup) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IpamsvcOptionGroup) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDhcpOptions

`func (o *IpamsvcOptionGroup) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcOptionGroup) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcOptionGroup) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcOptionGroup) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcOptionGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcOptionGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcOptionGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcOptionGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcOptionGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcOptionGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcOptionGroup) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *IpamsvcOptionGroup) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamsvcOptionGroup) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamsvcOptionGroup) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamsvcOptionGroup) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcOptionGroup) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcOptionGroup) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcOptionGroup) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcOptionGroup) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcOptionGroup) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcOptionGroup) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcOptionGroup) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcOptionGroup) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IpamsvcOptionGroup) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IpamsvcOptionGroup) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


