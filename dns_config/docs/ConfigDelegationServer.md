# ConfigDelegationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Optional. IP Address of nameserver.  Only required when fqdn of a delegation server falls under delegation fqdn | [optional] 
**Fqdn** | **string** | Required. FQDN of nameserver. | 
**ProtocolFqdn** | Pointer to **string** | FQDN of nameserver in punycode. | [optional] [readonly] 

## Methods

### NewConfigDelegationServer

`func NewConfigDelegationServer(fqdn string, ) *ConfigDelegationServer`

NewConfigDelegationServer instantiates a new ConfigDelegationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDelegationServerWithDefaults

`func NewConfigDelegationServerWithDefaults() *ConfigDelegationServer`

NewConfigDelegationServerWithDefaults instantiates a new ConfigDelegationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConfigDelegationServer) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConfigDelegationServer) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConfigDelegationServer) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConfigDelegationServer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFqdn

`func (o *ConfigDelegationServer) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigDelegationServer) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigDelegationServer) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetProtocolFqdn

`func (o *ConfigDelegationServer) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ConfigDelegationServer) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ConfigDelegationServer) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ConfigDelegationServer) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


