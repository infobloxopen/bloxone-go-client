# AtcdfpResolver

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | address that can be used as resolver | [optional] 
**IsFallback** | Pointer to **bool** | Mark it true to set default DNS resolvers that will be used in case if the BloxOne Cloud is unreachable. | [optional] 
**IsLocal** | Pointer to **bool** | Mark it true to set internal or local DNS servers&#39; IPv4 or IPv6 addresses that are used as DNS resolvers | [optional] 
**Protocols** | Pointer to [**[]AtcdfpDNSProtocol**](AtcdfpDNSProtocol.md) | The list of DNS resolver communication protocols. | [optional] 

## Methods

### NewAtcdfpResolver

`func NewAtcdfpResolver() *AtcdfpResolver`

NewAtcdfpResolver instantiates a new AtcdfpResolver object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcdfpResolverWithDefaults

`func NewAtcdfpResolverWithDefaults() *AtcdfpResolver`

NewAtcdfpResolverWithDefaults instantiates a new AtcdfpResolver object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AtcdfpResolver) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AtcdfpResolver) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AtcdfpResolver) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AtcdfpResolver) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetIsFallback

`func (o *AtcdfpResolver) GetIsFallback() bool`

GetIsFallback returns the IsFallback field if non-nil, zero value otherwise.

### GetIsFallbackOk

`func (o *AtcdfpResolver) GetIsFallbackOk() (*bool, bool)`

GetIsFallbackOk returns a tuple with the IsFallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFallback

`func (o *AtcdfpResolver) SetIsFallback(v bool)`

SetIsFallback sets IsFallback field to given value.

### HasIsFallback

`func (o *AtcdfpResolver) HasIsFallback() bool`

HasIsFallback returns a boolean if a field has been set.

### GetIsLocal

`func (o *AtcdfpResolver) GetIsLocal() bool`

GetIsLocal returns the IsLocal field if non-nil, zero value otherwise.

### GetIsLocalOk

`func (o *AtcdfpResolver) GetIsLocalOk() (*bool, bool)`

GetIsLocalOk returns a tuple with the IsLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocal

`func (o *AtcdfpResolver) SetIsLocal(v bool)`

SetIsLocal sets IsLocal field to given value.

### HasIsLocal

`func (o *AtcdfpResolver) HasIsLocal() bool`

HasIsLocal returns a boolean if a field has been set.

### GetProtocols

`func (o *AtcdfpResolver) GetProtocols() []AtcdfpDNSProtocol`

GetProtocols returns the Protocols field if non-nil, zero value otherwise.

### GetProtocolsOk

`func (o *AtcdfpResolver) GetProtocolsOk() (*[]AtcdfpDNSProtocol, bool)`

GetProtocolsOk returns a tuple with the Protocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocols

`func (o *AtcdfpResolver) SetProtocols(v []AtcdfpDNSProtocol)`

SetProtocols sets Protocols field to given value.

### HasProtocols

`func (o *AtcdfpResolver) HasProtocols() bool`

HasProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


