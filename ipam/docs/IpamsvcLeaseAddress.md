# IpamsvcLeaseAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The IP address for the DHCP lease in IPv4 or IPv6 format within the IP space specified by _space_ field. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewIpamsvcLeaseAddress

`func NewIpamsvcLeaseAddress() *IpamsvcLeaseAddress`

NewIpamsvcLeaseAddress instantiates a new IpamsvcLeaseAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcLeaseAddressWithDefaults

`func NewIpamsvcLeaseAddressWithDefaults() *IpamsvcLeaseAddress`

NewIpamsvcLeaseAddressWithDefaults instantiates a new IpamsvcLeaseAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcLeaseAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcLeaseAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcLeaseAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpamsvcLeaseAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSpace

`func (o *IpamsvcLeaseAddress) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcLeaseAddress) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcLeaseAddress) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *IpamsvcLeaseAddress) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


