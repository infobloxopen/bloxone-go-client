# IpamsvcHostName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **bool** | When _true_, the name is treated as an alias. | [optional] 
**Name** | **string** | A name for the host. | 
**PrimaryName** | Pointer to **bool** | When _true_, the name field is treated as primary name. There must be one and only one primary name in the list of host names. The primary name will be treated as the canonical name for all the aliases. PTR record will be generated only for the primary name. | [optional] 
**Zone** | **string** | The resource identifier. | 

## Methods

### NewIpamsvcHostName

`func NewIpamsvcHostName(name string, zone string, ) *IpamsvcHostName`

NewIpamsvcHostName instantiates a new IpamsvcHostName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHostNameWithDefaults

`func NewIpamsvcHostNameWithDefaults() *IpamsvcHostName`

NewIpamsvcHostNameWithDefaults instantiates a new IpamsvcHostName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *IpamsvcHostName) GetAlias() bool`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *IpamsvcHostName) GetAliasOk() (*bool, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *IpamsvcHostName) SetAlias(v bool)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *IpamsvcHostName) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcHostName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcHostName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcHostName) SetName(v string)`

SetName sets Name field to given value.


### GetPrimaryName

`func (o *IpamsvcHostName) GetPrimaryName() bool`

GetPrimaryName returns the PrimaryName field if non-nil, zero value otherwise.

### GetPrimaryNameOk

`func (o *IpamsvcHostName) GetPrimaryNameOk() (*bool, bool)`

GetPrimaryNameOk returns a tuple with the PrimaryName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryName

`func (o *IpamsvcHostName) SetPrimaryName(v bool)`

SetPrimaryName sets PrimaryName field to given value.

### HasPrimaryName

`func (o *IpamsvcHostName) HasPrimaryName() bool`

HasPrimaryName returns a boolean if a field has been set.

### GetZone

`func (o *IpamsvcHostName) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *IpamsvcHostName) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *IpamsvcHostName) SetZone(v string)`

SetZone sets Zone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


