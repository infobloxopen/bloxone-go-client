# FixedAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The reserved address. | 
**Comment** | Pointer to **string** | The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DhcpOptions** | Pointer to [**[]OptionItem**](OptionItem.md) | The list of DHCP options. May be either a specific option or a group of options. | [optional] 
**DisableDhcp** | Pointer to **bool** | Optional. _true_ to disable object. The fixed address is converted to an exclusion when generating configuration.  Defaults to _false_. | [optional] 
**HeaderOptionFilename** | Pointer to **string** | The configuration for header option filename field. | [optional] 
**HeaderOptionServerAddress** | Pointer to **string** | The configuration for header option server address field. | [optional] 
**HeaderOptionServerName** | Pointer to **string** | The configuration for header option server name field. | [optional] 
**Hostname** | Pointer to **string** | The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceAssignedHosts** | Pointer to [**[]InheritanceAssignedHost**](InheritanceAssignedHost.md) | The list of the inheritance assigned hosts of the object. | [optional] [readonly] 
**InheritanceParent** | Pointer to **string** | The resource identifier. | [optional] 
**InheritanceSources** | Pointer to [**FixedAddressInheritance**](FixedAddressInheritance.md) |  | [optional] 
**IpSpace** | Pointer to **string** | The resource identifier. | [optional] 
**MatchType** | **string** | Indicates how to match the client:  * _mac_: match the client MAC address for both IPv4 and IPv6,  * _client_text_ or _client_hex_: match the client identifier for IPv4 only,  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82) for IPv4 only,  * _duid_: match the DHCP unique identifier, currently match only for IPv6 protocol. | 
**MatchValue** | **string** | The value to match. | 
**Name** | Pointer to **string** | The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the fixed address in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 

## Methods

### NewFixedAddress

`func NewFixedAddress(address string, matchType string, matchValue string, ) *FixedAddress`

NewFixedAddress instantiates a new FixedAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFixedAddressWithDefaults

`func NewFixedAddressWithDefaults() *FixedAddress`

NewFixedAddressWithDefaults instantiates a new FixedAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FixedAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FixedAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FixedAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetComment

`func (o *FixedAddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FixedAddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FixedAddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FixedAddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FixedAddress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FixedAddress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FixedAddress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FixedAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDhcpOptions

`func (o *FixedAddress) GetDhcpOptions() []OptionItem`

GetDhcpOptions returns the DhcpOptions field if non-nil, zero value otherwise.

### GetDhcpOptionsOk

`func (o *FixedAddress) GetDhcpOptionsOk() (*[]OptionItem, bool)`

GetDhcpOptionsOk returns a tuple with the DhcpOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpOptions

`func (o *FixedAddress) SetDhcpOptions(v []OptionItem)`

SetDhcpOptions sets DhcpOptions field to given value.

### HasDhcpOptions

`func (o *FixedAddress) HasDhcpOptions() bool`

HasDhcpOptions returns a boolean if a field has been set.

### GetDisableDhcp

`func (o *FixedAddress) GetDisableDhcp() bool`

GetDisableDhcp returns the DisableDhcp field if non-nil, zero value otherwise.

### GetDisableDhcpOk

`func (o *FixedAddress) GetDisableDhcpOk() (*bool, bool)`

GetDisableDhcpOk returns a tuple with the DisableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDhcp

`func (o *FixedAddress) SetDisableDhcp(v bool)`

SetDisableDhcp sets DisableDhcp field to given value.

### HasDisableDhcp

`func (o *FixedAddress) HasDisableDhcp() bool`

HasDisableDhcp returns a boolean if a field has been set.

### GetHeaderOptionFilename

`func (o *FixedAddress) GetHeaderOptionFilename() string`

GetHeaderOptionFilename returns the HeaderOptionFilename field if non-nil, zero value otherwise.

### GetHeaderOptionFilenameOk

`func (o *FixedAddress) GetHeaderOptionFilenameOk() (*string, bool)`

GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionFilename

`func (o *FixedAddress) SetHeaderOptionFilename(v string)`

SetHeaderOptionFilename sets HeaderOptionFilename field to given value.

### HasHeaderOptionFilename

`func (o *FixedAddress) HasHeaderOptionFilename() bool`

HasHeaderOptionFilename returns a boolean if a field has been set.

### GetHeaderOptionServerAddress

`func (o *FixedAddress) GetHeaderOptionServerAddress() string`

GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field if non-nil, zero value otherwise.

### GetHeaderOptionServerAddressOk

`func (o *FixedAddress) GetHeaderOptionServerAddressOk() (*string, bool)`

GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerAddress

`func (o *FixedAddress) SetHeaderOptionServerAddress(v string)`

SetHeaderOptionServerAddress sets HeaderOptionServerAddress field to given value.

### HasHeaderOptionServerAddress

`func (o *FixedAddress) HasHeaderOptionServerAddress() bool`

HasHeaderOptionServerAddress returns a boolean if a field has been set.

### GetHeaderOptionServerName

`func (o *FixedAddress) GetHeaderOptionServerName() string`

GetHeaderOptionServerName returns the HeaderOptionServerName field if non-nil, zero value otherwise.

### GetHeaderOptionServerNameOk

`func (o *FixedAddress) GetHeaderOptionServerNameOk() (*string, bool)`

GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderOptionServerName

`func (o *FixedAddress) SetHeaderOptionServerName(v string)`

SetHeaderOptionServerName sets HeaderOptionServerName field to given value.

### HasHeaderOptionServerName

`func (o *FixedAddress) HasHeaderOptionServerName() bool`

HasHeaderOptionServerName returns a boolean if a field has been set.

### GetHostname

`func (o *FixedAddress) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *FixedAddress) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *FixedAddress) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *FixedAddress) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetId

`func (o *FixedAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FixedAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FixedAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FixedAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceAssignedHosts

`func (o *FixedAddress) GetInheritanceAssignedHosts() []InheritanceAssignedHost`

GetInheritanceAssignedHosts returns the InheritanceAssignedHosts field if non-nil, zero value otherwise.

### GetInheritanceAssignedHostsOk

`func (o *FixedAddress) GetInheritanceAssignedHostsOk() (*[]InheritanceAssignedHost, bool)`

GetInheritanceAssignedHostsOk returns a tuple with the InheritanceAssignedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceAssignedHosts

`func (o *FixedAddress) SetInheritanceAssignedHosts(v []InheritanceAssignedHost)`

SetInheritanceAssignedHosts sets InheritanceAssignedHosts field to given value.

### HasInheritanceAssignedHosts

`func (o *FixedAddress) HasInheritanceAssignedHosts() bool`

HasInheritanceAssignedHosts returns a boolean if a field has been set.

### GetInheritanceParent

`func (o *FixedAddress) GetInheritanceParent() string`

GetInheritanceParent returns the InheritanceParent field if non-nil, zero value otherwise.

### GetInheritanceParentOk

`func (o *FixedAddress) GetInheritanceParentOk() (*string, bool)`

GetInheritanceParentOk returns a tuple with the InheritanceParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceParent

`func (o *FixedAddress) SetInheritanceParent(v string)`

SetInheritanceParent sets InheritanceParent field to given value.

### HasInheritanceParent

`func (o *FixedAddress) HasInheritanceParent() bool`

HasInheritanceParent returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *FixedAddress) GetInheritanceSources() FixedAddressInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *FixedAddress) GetInheritanceSourcesOk() (*FixedAddressInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *FixedAddress) SetInheritanceSources(v FixedAddressInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *FixedAddress) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetIpSpace

`func (o *FixedAddress) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *FixedAddress) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *FixedAddress) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *FixedAddress) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetMatchType

`func (o *FixedAddress) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *FixedAddress) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *FixedAddress) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetMatchValue

`func (o *FixedAddress) GetMatchValue() string`

GetMatchValue returns the MatchValue field if non-nil, zero value otherwise.

### GetMatchValueOk

`func (o *FixedAddress) GetMatchValueOk() (*string, bool)`

GetMatchValueOk returns a tuple with the MatchValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchValue

`func (o *FixedAddress) SetMatchValue(v string)`

SetMatchValue sets MatchValue field to given value.


### GetName

`func (o *FixedAddress) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FixedAddress) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FixedAddress) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FixedAddress) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *FixedAddress) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *FixedAddress) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *FixedAddress) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *FixedAddress) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetTags

`func (o *FixedAddress) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FixedAddress) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FixedAddress) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FixedAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FixedAddress) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FixedAddress) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FixedAddress) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FixedAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


