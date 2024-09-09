# IPAMConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpServer** | Pointer to **string** |  | [optional] 
**IpSpace** | Pointer to **string** |  | [optional] 

## Methods

### NewIPAMConfig

`func NewIPAMConfig() *IPAMConfig`

NewIPAMConfig instantiates a new IPAMConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPAMConfigWithDefaults

`func NewIPAMConfigWithDefaults() *IPAMConfig`

NewIPAMConfigWithDefaults instantiates a new IPAMConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpServer

`func (o *IPAMConfig) GetDhcpServer() string`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *IPAMConfig) GetDhcpServerOk() (*string, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *IPAMConfig) SetDhcpServer(v string)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *IPAMConfig) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetIpSpace

`func (o *IPAMConfig) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *IPAMConfig) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *IPAMConfig) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *IPAMConfig) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


