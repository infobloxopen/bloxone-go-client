# MacAddressItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The address to match for the hardware filter. | 
**HardwareFilterId** | Pointer to **string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 

## Methods

### NewMacAddressItem

`func NewMacAddressItem(address string, ) *MacAddressItem`

NewMacAddressItem instantiates a new MacAddressItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMacAddressItemWithDefaults

`func NewMacAddressItemWithDefaults() *MacAddressItem`

NewMacAddressItemWithDefaults instantiates a new MacAddressItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MacAddressItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MacAddressItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MacAddressItem) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetHardwareFilterId

`func (o *MacAddressItem) GetHardwareFilterId() string`

GetHardwareFilterId returns the HardwareFilterId field if non-nil, zero value otherwise.

### GetHardwareFilterIdOk

`func (o *MacAddressItem) GetHardwareFilterIdOk() (*string, bool)`

GetHardwareFilterIdOk returns a tuple with the HardwareFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareFilterId

`func (o *MacAddressItem) SetHardwareFilterId(v string)`

SetHardwareFilterId sets HardwareFilterId field to given value.

### HasHardwareFilterId

`func (o *MacAddressItem) HasHardwareFilterId() bool`

HasHardwareFilterId returns a boolean if a field has been set.

### GetId

`func (o *MacAddressItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MacAddressItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MacAddressItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MacAddressItem) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


