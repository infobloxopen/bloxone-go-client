# Forwarder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Server IP address. | 
**Fqdn** | Pointer to **string** | Server FQDN. | [optional] 
**ProtocolFqdn** | Pointer to **string** | Server FQDN in punycode. | [optional] [readonly] 

## Methods

### NewForwarder

`func NewForwarder(address string, ) *Forwarder`

NewForwarder instantiates a new Forwarder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwarderWithDefaults

`func NewForwarderWithDefaults() *Forwarder`

NewForwarderWithDefaults instantiates a new Forwarder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Forwarder) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Forwarder) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Forwarder) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFqdn

`func (o *Forwarder) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Forwarder) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Forwarder) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Forwarder) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *Forwarder) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *Forwarder) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *Forwarder) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *Forwarder) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


