# PoPRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **[]string** | PoP Region&#39;s IP addresses | [optional] 
**Id** | Pointer to **int32** | The PoP Region&#39;s serial, unique ID | [optional] [readonly] 
**Location** | Pointer to **string** | PoP Region&#39;s location | [optional] 
**Region** | Pointer to **string** | PoP Region&#39;s name | [optional] 

## Methods

### NewPoPRegion

`func NewPoPRegion() *PoPRegion`

NewPoPRegion instantiates a new PoPRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoPRegionWithDefaults

`func NewPoPRegionWithDefaults() *PoPRegion`

NewPoPRegionWithDefaults instantiates a new PoPRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *PoPRegion) GetAddresses() []string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *PoPRegion) GetAddressesOk() (*[]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *PoPRegion) SetAddresses(v []string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *PoPRegion) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetId

`func (o *PoPRegion) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoPRegion) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoPRegion) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PoPRegion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *PoPRegion) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *PoPRegion) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *PoPRegion) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *PoPRegion) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetRegion

`func (o *PoPRegion) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PoPRegion) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PoPRegion) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PoPRegion) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


