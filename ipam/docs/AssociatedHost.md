# AssociatedHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The primary IP address of the on-prem host. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The DHCP host name. | [optional] [readonly] 
**Ophid** | Pointer to **string** | The on-prem host ID. | [optional] [readonly] 

## Methods

### NewAssociatedHost

`func NewAssociatedHost() *AssociatedHost`

NewAssociatedHost instantiates a new AssociatedHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociatedHostWithDefaults

`func NewAssociatedHostWithDefaults() *AssociatedHost`

NewAssociatedHostWithDefaults instantiates a new AssociatedHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AssociatedHost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AssociatedHost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AssociatedHost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AssociatedHost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *AssociatedHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssociatedHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssociatedHost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssociatedHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AssociatedHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssociatedHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssociatedHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssociatedHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOphid

`func (o *AssociatedHost) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *AssociatedHost) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *AssociatedHost) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *AssociatedHost) HasOphid() bool`

HasOphid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


