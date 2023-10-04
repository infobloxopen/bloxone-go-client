# IpamsvcDDNSUpdateBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdnsDomain** | Pointer to **string** | The domain name for DDNS. | [optional] 
**DdnsSendUpdates** | Pointer to **bool** | Determines if DDNS updates are enabled at this level. | [optional] 

## Methods

### NewIpamsvcDDNSUpdateBlock

`func NewIpamsvcDDNSUpdateBlock() *IpamsvcDDNSUpdateBlock`

NewIpamsvcDDNSUpdateBlock instantiates a new IpamsvcDDNSUpdateBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcDDNSUpdateBlockWithDefaults

`func NewIpamsvcDDNSUpdateBlockWithDefaults() *IpamsvcDDNSUpdateBlock`

NewIpamsvcDDNSUpdateBlockWithDefaults instantiates a new IpamsvcDDNSUpdateBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdnsDomain

`func (o *IpamsvcDDNSUpdateBlock) GetDdnsDomain() string`

GetDdnsDomain returns the DdnsDomain field if non-nil, zero value otherwise.

### GetDdnsDomainOk

`func (o *IpamsvcDDNSUpdateBlock) GetDdnsDomainOk() (*string, bool)`

GetDdnsDomainOk returns a tuple with the DdnsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsDomain

`func (o *IpamsvcDDNSUpdateBlock) SetDdnsDomain(v string)`

SetDdnsDomain sets DdnsDomain field to given value.

### HasDdnsDomain

`func (o *IpamsvcDDNSUpdateBlock) HasDdnsDomain() bool`

HasDdnsDomain returns a boolean if a field has been set.

### GetDdnsSendUpdates

`func (o *IpamsvcDDNSUpdateBlock) GetDdnsSendUpdates() bool`

GetDdnsSendUpdates returns the DdnsSendUpdates field if non-nil, zero value otherwise.

### GetDdnsSendUpdatesOk

`func (o *IpamsvcDDNSUpdateBlock) GetDdnsSendUpdatesOk() (*bool, bool)`

GetDdnsSendUpdatesOk returns a tuple with the DdnsSendUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdnsSendUpdates

`func (o *IpamsvcDDNSUpdateBlock) SetDdnsSendUpdates(v bool)`

SetDdnsSendUpdates sets DdnsSendUpdates field to given value.

### HasDdnsSendUpdates

`func (o *IpamsvcDDNSUpdateBlock) HasDdnsSendUpdates() bool`

HasDdnsSendUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


