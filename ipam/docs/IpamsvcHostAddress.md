# IpamsvcHostAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_. | 
**Ref** | **string** | The resource identifier. | 
**Space** | **string** | The resource identifier. | 

## Methods

### NewIpamsvcHostAddress

`func NewIpamsvcHostAddress(address string, ref string, space string, ) *IpamsvcHostAddress`

NewIpamsvcHostAddress instantiates a new IpamsvcHostAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHostAddressWithDefaults

`func NewIpamsvcHostAddressWithDefaults() *IpamsvcHostAddress`

NewIpamsvcHostAddressWithDefaults instantiates a new IpamsvcHostAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcHostAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcHostAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcHostAddress) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetRef

`func (o *IpamsvcHostAddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *IpamsvcHostAddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *IpamsvcHostAddress) SetRef(v string)`

SetRef sets Ref field to given value.


### GetSpace

`func (o *IpamsvcHostAddress) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *IpamsvcHostAddress) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *IpamsvcHostAddress) SetSpace(v string)`

SetSpace sets Space field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


