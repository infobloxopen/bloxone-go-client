# BulkMacAddressItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | **[]string** | The addresses to match for the hardware filter. | 
**HardwareFilterId** | **string** | The resource identifier. | 

## Methods

### NewBulkMacAddressItem

`func NewBulkMacAddressItem(addresses []string, hardwareFilterId string, ) *BulkMacAddressItem`

NewBulkMacAddressItem instantiates a new BulkMacAddressItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMacAddressItemWithDefaults

`func NewBulkMacAddressItemWithDefaults() *BulkMacAddressItem`

NewBulkMacAddressItemWithDefaults instantiates a new BulkMacAddressItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *BulkMacAddressItem) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *BulkMacAddressItem) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *BulkMacAddressItem) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.


### GetHardwareFilterId

`func (o *BulkMacAddressItem) GetHardwareFilterId() string`

GetHardwareFilterId returns the HardwareFilterId field if non-nil, zero value otherwise.

### GetHardwareFilterIdOk

`func (o *BulkMacAddressItem) GetHardwareFilterIdOk() (*string, bool)`

GetHardwareFilterIdOk returns a tuple with the HardwareFilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareFilterId

`func (o *BulkMacAddressItem) SetHardwareFilterId(v string)`

SetHardwareFilterId sets HardwareFilterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


