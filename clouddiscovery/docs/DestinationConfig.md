# DestinationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dns** | Pointer to [**DNSConfig**](DNSConfig.md) |  | [optional] 
**Ipam** | Pointer to [**IPAMConfig**](IPAMConfig.md) |  | [optional] 

## Methods

### NewDestinationConfig

`func NewDestinationConfig() *DestinationConfig`

NewDestinationConfig instantiates a new DestinationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationConfigWithDefaults

`func NewDestinationConfigWithDefaults() *DestinationConfig`

NewDestinationConfigWithDefaults instantiates a new DestinationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDns

`func (o *DestinationConfig) GetDns() DNSConfig`

GetDns returns the Dns field if non-nil, zero value otherwise.

### GetDnsOk

`func (o *DestinationConfig) GetDnsOk() (*DNSConfig, bool)`

GetDnsOk returns a tuple with the Dns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDns

`func (o *DestinationConfig) SetDns(v DNSConfig)`

SetDns sets Dns field to given value.

### HasDns

`func (o *DestinationConfig) HasDns() bool`

HasDns returns a boolean if a field has been set.

### GetIpam

`func (o *DestinationConfig) GetIpam() IPAMConfig`

GetIpam returns the Ipam field if non-nil, zero value otherwise.

### GetIpamOk

`func (o *DestinationConfig) GetIpamOk() (*IPAMConfig, bool)`

GetIpamOk returns a tuple with the Ipam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpam

`func (o *DestinationConfig) SetIpam(v IPAMConfig)`

SetIpam sets Ipam field to given value.

### HasIpam

`func (o *DestinationConfig) HasIpam() bool`

HasIpam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


