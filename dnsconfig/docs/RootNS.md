# RootNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IPv4 address. | 
**Fqdn** | **string** | FQDN. | 
**ProtocolFqdn** | Pointer to **string** | FQDN in punycode. | [optional] [readonly] 

## Methods

### NewRootNS

`func NewRootNS(address string, fqdn string, ) *RootNS`

NewRootNS instantiates a new RootNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRootNSWithDefaults

`func NewRootNSWithDefaults() *RootNS`

NewRootNSWithDefaults instantiates a new RootNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RootNS) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RootNS) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RootNS) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFqdn

`func (o *RootNS) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *RootNS) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *RootNS) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetProtocolFqdn

`func (o *RootNS) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *RootNS) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *RootNS) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *RootNS) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


