# IpamsvcHardwareFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | The list of addresses to match for the hardware filter. | [optional] 
**Comment** | Pointer to **string** | The description for the hardware filter. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **NullableTime** | Time when the object has been created. | [optional] [readonly] 
**DhcpOptions** | Pointer to [**[]IpamsvcOptionItem**](IpamsvcOptionItem.md) | The list of DHCP options for the hardware filter. May be either a specific option or a group of options. | [optional] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**LeaseTime** | Pointer to **int64** | The lease lifetime duration in seconds. | [optional] 
**Name** | **string** | The name of the hardware filter. Must contain 1 to 256 characters. Can include UTF-8. | 
**Role** | Pointer to **string** | The role of DHCP filter (_values_ or _selection_).  Defaults to _values_. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the hardware filter in JSON format. | [optional] 
**UpdatedAt** | Pointer to **NullableTime** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**VendorSpecificOptionOptionSpace** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIpamsvcHardwareFilter

`func NewIpamsvcHardwareFilter(name string, ) *IpamsvcHardwareFilter`

NewIpamsvcHardwareFilter instantiates a new IpamsvcHardwareFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHardwareFilterWithDefaults

`func NewIpamsvcHardwareFilterWithDefaults() *IpamsvcHardwareFilter`

NewIpamsvcHardwareFilterWithDefaults instantiates a new IpamsvcHardwareFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *IpamsvcHardwareFilter) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *IpamsvcHardwareFilter) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *IpamsvcHardwareFilter) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *IpamsvcHardwareFilter) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcHardwareFilter) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcHardwareFilter) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcHardwareFilter) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcHardwareFilter) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcHardwareFilter) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcHardwareFilter) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcHardwareFilter) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcHardwareFilter) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *IpamsvcHardwareFilter) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IpamsvcHardwareFilter) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDhcpOptions

`func (o *IpamsvcHardwareFilter) GetDhcpOptions() []IpamsvcOptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *IpamsvcHardwareFilter) GetDhcpOptionsOk() (*[]IpamsvcOptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *IpamsvcHardwareFilter) SetDhcpOptions(v []IpamsvcOptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *IpamsvcHardwareFilter) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *IpamsvcHardwareFilter) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *IpamsvcHardwareFilter) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *IpamsvcHardwareFilter) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *IpamsvcHardwareFilter) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *IpamsvcHardwareFilter) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *IpamsvcHardwareFilter) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *IpamsvcHardwareFilter) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *IpamsvcHardwareFilter) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *IpamsvcHardwareFilter) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *IpamsvcHardwareFilter) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *IpamsvcHardwareFilter) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *IpamsvcHardwareFilter) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcHardwareFilter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcHardwareFilter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcHardwareFilter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcHardwareFilter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeaseTime

`func (o *IpamsvcHardwareFilter) GetLeaseTime() int64`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *IpamsvcHardwareFilter) GetLeaseTimeOk() (*int64, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *IpamsvcHardwareFilter) SetLeaseTime(v int64)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *IpamsvcHardwareFilter) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcHardwareFilter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcHardwareFilter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcHardwareFilter) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *IpamsvcHardwareFilter) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *IpamsvcHardwareFilter) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *IpamsvcHardwareFilter) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *IpamsvcHardwareFilter) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcHardwareFilter) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcHardwareFilter) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcHardwareFilter) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcHardwareFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcHardwareFilter) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcHardwareFilter) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcHardwareFilter) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcHardwareFilter) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *IpamsvcHardwareFilter) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *IpamsvcHardwareFilter) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetVendorSpecificOptionOptionSpace

`func (o *IpamsvcHardwareFilter) GetVendorSpecificOptionOptionSpace() string`

GetVendorSpecificOptionOptionSpace returns the VendorSpecificOptionOptionSpace field if non-nil, zero value otherwise.

### GetVendorSpecificOptionOptionSpaceOk

`func (o *IpamsvcHardwareFilter) GetVendorSpecificOptionOptionSpaceOk() (*string, bool)`

GetVendorSpecificOptionOptionSpaceOk returns a tuple with the VendorSpecificOptionOptionSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorSpecificOptionOptionSpace

`func (o *IpamsvcHardwareFilter) SetVendorSpecificOptionOptionSpace(v string)`

SetVendorSpecificOptionOptionSpace sets VendorSpecificOptionOptionSpace field to given value.

### HasVendorSpecificOptionOptionSpace

`func (o *IpamsvcHardwareFilter) HasVendorSpecificOptionOptionSpace() bool`

HasVendorSpecificOptionOptionSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


