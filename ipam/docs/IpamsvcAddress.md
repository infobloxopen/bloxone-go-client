# IpamsvcAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address in form \&quot;a.b.c.d\&quot;. | 
**Comment** | Pointer to **string** | The description for the address object. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when the object has been created. | [optional] [readonly] 
**DhcpInfo** | Pointer to [**IpamsvcDHCPInfo**](IpamsvcDHCPInfo.md) |  | [optional] 
**DisableDhcp** | Pointer to **bool** | Read only. Represent the value of the same field in the associated _dhcp/fixed_address_ object. | [optional] [readonly] 
**DiscoveryAttrs** | Pointer to **map[string]interface{}** | The discovery attributes for this address in JSON format. | [optional] [readonly] 
**DiscoveryMetadata** | Pointer to **map[string]interface{}** | The discovery metadata for this address in JSON format. | [optional] [readonly] 
**Host** | Pointer to **string** | The resource identifier. | [optional] 
**Hwaddr** | Pointer to **string** | The hardware address associated with this IP address. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Interface** | Pointer to **string** | The name of the network interface card (NIC) associated with the address, if any. | [optional] 
**Names** | Pointer to [**[]IpamsvcName**](IpamsvcName.md) | The list of all names associated with this address. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**Protocol** | Pointer to **string** | The type of protocol (_ip4_ or _ip6_). | [optional] [readonly] 
**Range** | Pointer to **string** | The resource identifier. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 
**State** | Pointer to **string** | The state of the address (_used_ or _free_). | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags for this address in JSON format. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Time when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**Usage** | Pointer to **[]string** | The usage is a combination of indicators, each tracking a specific associated use. Listed below are usage indicators with their meaning:  usage indicator        | description  ---------------------- | --------------------------------  _IPAM_                 |  Address was created by the IPAM component.  _IPAM_, _RESERVED_     |  Address was created by the API call _ipam/address_ or _ipam/host_.  _IPAM_, _NETWORK_      |  Address was automatically created by the IPAM component and is the network address of the parent subnet.  _IPAM_, _BROADCAST_    |  Address was automatically created by the IPAM component and is the broadcast address of the parent subnet.  _DHCP_                 |  Address was created by the DHCP component.  _DHCP_, _FIXEDADDRESS_ |  Address was created by the API call _dhcp/fixed_address_.  _DHCP_, _LEASED_       |  An active lease for that address was issued by a DHCP server.  _DHCP_, _DISABLED_     |  Address is disabled.  _DNS_                  |  Address is used by one or more DNS records.  _DISCOVERED_           |  Address is discovered by some network discovery probe like Network Insight or NetMRI in NIOS. | [optional] [readonly] 

## Methods

### NewIpamsvcAddress

`func NewIpamsvcAddress(address string, ) *IpamsvcAddress`

NewIpamsvcAddress instantiates a new IpamsvcAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcAddressWithDefaults

`func NewIpamsvcAddressWithDefaults() *IpamsvcAddress`

NewIpamsvcAddressWithDefaults instantiates a new IpamsvcAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetComment

`func (o *IpamsvcAddress) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcAddress) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcAddress) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcAddress) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IpamsvcAddress) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IpamsvcAddress) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IpamsvcAddress) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IpamsvcAddress) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDhcpInfo

`func (o *IpamsvcAddress) GetDhcpInfo() IpamsvcDHCPInfo`

GetDhcpInfo returns the DhcpInfo field if non-nil, zero value otherwise.

### GetDhcpInfoOk

`func (o *IpamsvcAddress) GetDhcpInfoOk() (*IpamsvcDHCPInfo, bool)`

GetDhcpInfoOk returns a tuple with the DhcpInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpInfo

`func (o *IpamsvcAddress) SetDhcpInfo(v IpamsvcDHCPInfo)`

SetDhcpInfo sets DhcpInfo field to given value.

### HasDhcpInfo

`func (o *IpamsvcAddress) HasDhcpInfo() bool`

HasDhcpInfo returns a boolean if a field has been set.

### GetDisableDhcp

`func (o *IpamsvcAddress) GetDisableDhcp() bool`

GetDisableDhcp returns the DisableDhcp field if non-nil, zero value otherwise.

### GetDisableDhcpOk

`func (o *IpamsvcAddress) GetDisableDhcpOk() (*bool, bool)`

GetDisableDhcpOk returns a tuple with the DisableDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDhcp

`func (o *IpamsvcAddress) SetDisableDhcp(v bool)`

SetDisableDhcp sets DisableDhcp field to given value.

### HasDisableDhcp

`func (o *IpamsvcAddress) HasDisableDhcp() bool`

HasDisableDhcp returns a boolean if a field has been set.

### GetDiscoveryAttrs

`func (o *IpamsvcAddress) GetDiscoveryAttrs() map[string]interface{}`

GetDiscoveryAttrs returns the DiscoveryAttrs field if non-nil, zero value otherwise.

### GetDiscoveryAttrsOk

`func (o *IpamsvcAddress) GetDiscoveryAttrsOk() (*map[string]interface{}, bool)`

GetDiscoveryAttrsOk returns a tuple with the DiscoveryAttrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryAttrs

`func (o *IpamsvcAddress) SetDiscoveryAttrs(v map[string]interface{})`

SetDiscoveryAttrs sets DiscoveryAttrs field to given value.

### HasDiscoveryAttrs

`func (o *IpamsvcAddress) HasDiscoveryAttrs() bool`

HasDiscoveryAttrs returns a boolean if a field has been set.

### GetDiscoveryMetadata

`func (o *IpamsvcAddress) GetDiscoveryMetadata() map[string]interface{}`

GetDiscoveryMetadata returns the DiscoveryMetadata field if non-nil, zero value otherwise.

### GetDiscoveryMetadataOk

`func (o *IpamsvcAddress) GetDiscoveryMetadataOk() (*map[string]interface{}, bool)`

GetDiscoveryMetadataOk returns a tuple with the DiscoveryMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryMetadata

`func (o *IpamsvcAddress) SetDiscoveryMetadata(v map[string]interface{})`

SetDiscoveryMetadata sets DiscoveryMetadata field to given value.

### HasDiscoveryMetadata

`func (o *IpamsvcAddress) HasDiscoveryMetadata() bool`

HasDiscoveryMetadata returns a boolean if a field has been set.

### GetHost

`func (o *IpamsvcAddress) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *IpamsvcAddress) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *IpamsvcAddress) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *IpamsvcAddress) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetHwaddr

`func (o *IpamsvcAddress) GetHwaddr() string`

GetHwaddr returns the Hwaddr field if non-nil, zero value otherwise.

### GetHwaddrOk

`func (o *IpamsvcAddress) GetHwaddrOk() (*string, bool)`

GetHwaddrOk returns a tuple with the Hwaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHwaddr

`func (o *IpamsvcAddress) SetHwaddr(v string)`

SetHwaddr sets Hwaddr field to given value.

### HasHwaddr

`func (o *IpamsvcAddress) HasHwaddr() bool`

HasHwaddr returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcAddress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcAddress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcAddress) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcAddress) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterface

`func (o *IpamsvcAddress) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *IpamsvcAddress) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *IpamsvcAddress) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *IpamsvcAddress) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetNames

`func (o *IpamsvcAddress) GetNames() []IpamsvcName`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *IpamsvcAddress) GetNamesOk() (*[]IpamsvcName, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *IpamsvcAddress) SetNames(v []IpamsvcName)`

SetNames sets Names field to given value.

### HasNames

`func (o *IpamsvcAddress) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetParent

`func (o *IpamsvcAddress) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *IpamsvcAddress) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *IpamsvcAddress) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *IpamsvcAddress) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocol

`func (o *IpamsvcAddress) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpamsvcAddress) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpamsvcAddress) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpamsvcAddress) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRange

`func (o *IpamsvcAddress) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *IpamsvcAddress) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *IpamsvcAddress) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *IpamsvcAddress) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcAddress) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcAddress) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcAddress) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *IpamsvcAddress) HasSpace() bool`

HasSpace returns a boolean if a field has been set.

### GetState

`func (o *IpamsvcAddress) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IpamsvcAddress) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IpamsvcAddress) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IpamsvcAddress) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcAddress) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcAddress) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcAddress) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcAddress) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IpamsvcAddress) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IpamsvcAddress) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IpamsvcAddress) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IpamsvcAddress) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUsage

`func (o *IpamsvcAddress) GetUsage() []string`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *IpamsvcAddress) GetUsageOk() (*[]string, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *IpamsvcAddress) SetUsage(v []string)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *IpamsvcAddress) HasUsage() bool`

HasUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


