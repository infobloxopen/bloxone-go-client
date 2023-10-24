# IpamsvcDHCPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientHostname** | Pointer to **string** | The DHCP host name associated with this client. | [optional] [readonly] 
**ClientHwaddr** | Pointer to **string** | The hardware address associated with this client. | [optional] [readonly] 
**ClientId** | Pointer to **string** | The ID associated with this client. | [optional] [readonly] 
**End** | Pointer to **time.Time** | The timestamp at which the _state_, when set to _leased_, will be changed to _free_. | [optional] [readonly] 
**Fingerprint** | Pointer to **string** | The DHCP fingerprint for the associated lease. | [optional] [readonly] 
**Iaid** | Pointer to **int64** | Identity Association Identifier (IAID) of the lease. Applicable only for DHCPv6. | [optional] [readonly] 
**LeaseType** | Pointer to **string** | Lease type. Applicable only for address under DHCP control. The value can be empty for address not under DHCP control.  Valid values are: * _DHCPv6NonTemporaryAddress_: DHCPv6 non-temporary address (NA) * _DHCPv6TemporaryAddress_: DHCPv6 temporary address (TA) * _DHCPv6PrefixDelegation_: DHCPv6 prefix delegation (PD) * _DHCPv4_: DHCPv4 lease | [optional] [readonly] 
**PreferredLifetime** | Pointer to **time.Time** | The length of time that a valid address is preferred (i.e., the time until deprecation). When the preferred lifetime expires, the address becomes deprecated on the client. It is still considered leased on the server. Applicable only for DHCPv6. | [optional] [readonly] 
**Remain** | Pointer to **int64** | The remaining time, in seconds, until which the _state_, when set to _leased_, will remain in that state. | [optional] [readonly] 
**Start** | Pointer to **time.Time** | The timestamp at which _state_ was first set to _leased_. | [optional] [readonly] 
**State** | Pointer to **string** | Indicates the status of this IP address from a DHCP protocol standpoint as:   * _none_: Address is not under DHCP control.   * _free_: Address is under DHCP control but has no lease currently assigned.   * _leased_: Address is under DHCP control and has a lease currently assigned. The lease details are contained in the matching _dhcp/lease_ resource. | [optional] [readonly] 
**StateTs** | Pointer to **time.Time** | The timestamp at which the _state_ was last reported. | [optional] [readonly] 

## Methods

### NewIpamsvcDHCPInfo

`func NewIpamsvcDHCPInfo() *IpamsvcDHCPInfo`

NewIpamsvcDHCPInfo instantiates a new IpamsvcDHCPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDHCPInfoWithDefaults

`func NewIpamsvcDHCPInfoWithDefaults() *IpamsvcDHCPInfo`

NewIpamsvcDHCPInfoWithDefaults instantiates a new IpamsvcDHCPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientHostname

`func (o *IpamsvcDHCPInfo) GetClientHostname() string`

GetClientHostname returns the ClientHostname field if non-nil, zero value otherwise.

### GetClientHostnameOk

`func (o *IpamsvcDHCPInfo) GetClientHostnameOk() (*string, bool)`

GetClientHostnameOk returns a tuple with the ClientHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHostname

`func (o *IpamsvcDHCPInfo) SetClientHostname(v string)`

SetClientHostname sets ClientHostname field to given value.

### HasClientHostname

`func (o *IpamsvcDHCPInfo) HasClientHostname() bool`

HasClientHostname returns a boolean if a field has been set.

### GetClientHwaddr

`func (o *IpamsvcDHCPInfo) GetClientHwaddr() string`

GetClientHwaddr returns the ClientHwaddr field if non-nil, zero value otherwise.

### GetClientHwaddrOk

`func (o *IpamsvcDHCPInfo) GetClientHwaddrOk() (*string, bool)`

GetClientHwaddrOk returns a tuple with the ClientHwaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHwaddr

`func (o *IpamsvcDHCPInfo) SetClientHwaddr(v string)`

SetClientHwaddr sets ClientHwaddr field to given value.

### HasClientHwaddr

`func (o *IpamsvcDHCPInfo) HasClientHwaddr() bool`

HasClientHwaddr returns a boolean if a field has been set.

### GetClientId

`func (o *IpamsvcDHCPInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IpamsvcDHCPInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IpamsvcDHCPInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IpamsvcDHCPInfo) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetEnd

`func (o *IpamsvcDHCPInfo) GetEnd() time.Time`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *IpamsvcDHCPInfo) GetEndOk() (*time.Time, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *IpamsvcDHCPInfo) SetEnd(v time.Time)`

SetEnd sets End field to given value.

### HasEnd

`func (o *IpamsvcDHCPInfo) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFingerprint

`func (o *IpamsvcDHCPInfo) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *IpamsvcDHCPInfo) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *IpamsvcDHCPInfo) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *IpamsvcDHCPInfo) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetIaid

`func (o *IpamsvcDHCPInfo) GetIaid() int64`

GetIaid returns the Iaid field if non-nil, zero value otherwise.

### GetIaidOk

`func (o *IpamsvcDHCPInfo) GetIaidOk() (*int64, bool)`

GetIaidOk returns a tuple with the Iaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIaid

`func (o *IpamsvcDHCPInfo) SetIaid(v int64)`

SetIaid sets Iaid field to given value.

### HasIaid

`func (o *IpamsvcDHCPInfo) HasIaid() bool`

HasIaid returns a boolean if a field has been set.

### GetLeaseType

`func (o *IpamsvcDHCPInfo) GetLeaseType() string`

GetLeaseType returns the LeaseType field if non-nil, zero value otherwise.

### GetLeaseTypeOk

`func (o *IpamsvcDHCPInfo) GetLeaseTypeOk() (*string, bool)`

GetLeaseTypeOk returns a tuple with the LeaseType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseType

`func (o *IpamsvcDHCPInfo) SetLeaseType(v string)`

SetLeaseType sets LeaseType field to given value.

### HasLeaseType

`func (o *IpamsvcDHCPInfo) HasLeaseType() bool`

HasLeaseType returns a boolean if a field has been set.

### GetPreferredLifetime

`func (o *IpamsvcDHCPInfo) GetPreferredLifetime() time.Time`

GetPreferredLifetime returns the PreferredLifetime field if non-nil, zero value otherwise.

### GetPreferredLifetimeOk

`func (o *IpamsvcDHCPInfo) GetPreferredLifetimeOk() (*time.Time, bool)`

GetPreferredLifetimeOk returns a tuple with the PreferredLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredLifetime

`func (o *IpamsvcDHCPInfo) SetPreferredLifetime(v time.Time)`

SetPreferredLifetime sets PreferredLifetime field to given value.

### HasPreferredLifetime

`func (o *IpamsvcDHCPInfo) HasPreferredLifetime() bool`

HasPreferredLifetime returns a boolean if a field has been set.

### GetRemain

`func (o *IpamsvcDHCPInfo) GetRemain() int64`

GetRemain returns the Remain field if non-nil, zero value otherwise.

### GetRemainOk

`func (o *IpamsvcDHCPInfo) GetRemainOk() (*int64, bool)`

GetRemainOk returns a tuple with the Remain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemain

`func (o *IpamsvcDHCPInfo) SetRemain(v int64)`

SetRemain sets Remain field to given value.

### HasRemain

`func (o *IpamsvcDHCPInfo) HasRemain() bool`

HasRemain returns a boolean if a field has been set.

### GetStart

`func (o *IpamsvcDHCPInfo) GetStart() time.Time`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *IpamsvcDHCPInfo) GetStartOk() (*time.Time, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *IpamsvcDHCPInfo) SetStart(v time.Time)`

SetStart sets Start field to given value.

### HasStart

`func (o *IpamsvcDHCPInfo) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetState

`func (o *IpamsvcDHCPInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *IpamsvcDHCPInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *IpamsvcDHCPInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *IpamsvcDHCPInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateTs

`func (o *IpamsvcDHCPInfo) GetStateTs() time.Time`

GetStateTs returns the StateTs field if non-nil, zero value otherwise.

### GetStateTsOk

`func (o *IpamsvcDHCPInfo) GetStateTsOk() (*time.Time, bool)`

GetStateTsOk returns a tuple with the StateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateTs

`func (o *IpamsvcDHCPInfo) SetStateTs(v time.Time)`

SetStateTs sets StateTs field to given value.

### HasStateTs

`func (o *IpamsvcDHCPInfo) HasStateTs() bool`

HasStateTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


