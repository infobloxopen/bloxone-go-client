# ConfigRootNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IPv4 address. | 
**Fqdn** | **string** | FQDN. | 
**ProtocolFqdn** | Pointer to **string** | FQDN in punycode. | [optional] [readonly] 

## Methods

### NewConfigRootNS

`func NewConfigRootNS(address string, fqdn string, ) *ConfigRootNS`

NewConfigRootNS instantiates a new ConfigRootNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigRootNSWithDefaults

`func NewConfigRootNSWithDefaults() *ConfigRootNS`

NewConfigRootNSWithDefaults instantiates a new ConfigRootNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConfigRootNS) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConfigRootNS) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConfigRootNS) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFqdn

`func (o *ConfigRootNS) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigRootNS) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigRootNS) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetProtocolFqdn

`func (o *ConfigRootNS) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ConfigRootNS) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ConfigRootNS) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ConfigRootNS) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


