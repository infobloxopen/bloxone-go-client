# HostAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_. | [optional] 
**Ref** | Pointer to **string** | The resource identifier. | [optional] 
**Space** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewHostAddress

`func NewHostAddress() *HostAddress`

NewHostAddress instantiates a new HostAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostAddressWithDefaults

`func NewHostAddressWithDefaults() *HostAddress`

NewHostAddressWithDefaults instantiates a new HostAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *HostAddress) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *HostAddress) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *HostAddress) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *HostAddress) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetRef

`func (o *HostAddress) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *HostAddress) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *HostAddress) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *HostAddress) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSpace

`func (o *HostAddress) GetSpace() string`

GetSpace returns the Space field if non-nil, zero value otherwise.

### GetSpaceOk

`func (o *HostAddress) GetSpaceOk() (*string, bool)`

GetSpaceOk returns a tuple with the Space field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpace

`func (o *HostAddress) SetSpace(v string)`

SetSpace sets Space field to given value.

### HasSpace

`func (o *HostAddress) HasSpace() bool`

HasSpace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


