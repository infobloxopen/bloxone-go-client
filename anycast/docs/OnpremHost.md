# OnpremHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnycastConfigRefs** | Pointer to [**[]AnycastConfigRef**](AnycastConfigRef.md) |  | [optional] 
**ConfigBgp** | Pointer to [**BgpConfig**](BgpConfig.md) |  | [optional] 
**ConfigOspf** | Pointer to [**OspfConfig**](OspfConfig.md) |  | [optional] 
**ConfigOspfv3** | Pointer to [**Ospfv3Config**](Ospfv3Config.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IPv4 address of the on-prem host | [optional] 
**Ipv6Address** | Pointer to **string** | IPv6 address of the on-prem host | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewOnpremHost

`func NewOnpremHost() *OnpremHost`

NewOnpremHost instantiates a new OnpremHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnpremHostWithDefaults

`func NewOnpremHostWithDefaults() *OnpremHost`

NewOnpremHostWithDefaults instantiates a new OnpremHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnycastConfigRefs

`func (o *OnpremHost) GetAnycastConfigRefs() []AnycastConfigRef`

GetAnycastConfigRefs returns the AnycastConfigRefs field if non-nil, zero value otherwise.

### GetAnycastConfigRefsOk

`func (o *OnpremHost) GetAnycastConfigRefsOk() (*[]AnycastConfigRef, bool)`

GetAnycastConfigRefsOk returns a tuple with the AnycastConfigRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastConfigRefs

`func (o *OnpremHost) SetAnycastConfigRefs(v []AnycastConfigRef)`

SetAnycastConfigRefs sets AnycastConfigRefs field to given value.

### HasAnycastConfigRefs

`func (o *OnpremHost) HasAnycastConfigRefs() bool`

HasAnycastConfigRefs returns a boolean if a field has been set.

### GetConfigBgp

`func (o *OnpremHost) GetConfigBgp() BgpConfig`

GetConfigBgp returns the ConfigBgp field if non-nil, zero value otherwise.

### GetConfigBgpOk

`func (o *OnpremHost) GetConfigBgpOk() (*BgpConfig, bool)`

GetConfigBgpOk returns a tuple with the ConfigBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigBgp

`func (o *OnpremHost) SetConfigBgp(v BgpConfig)`

SetConfigBgp sets ConfigBgp field to given value.

### HasConfigBgp

`func (o *OnpremHost) HasConfigBgp() bool`

HasConfigBgp returns a boolean if a field has been set.

### GetConfigOspf

`func (o *OnpremHost) GetConfigOspf() OspfConfig`

GetConfigOspf returns the ConfigOspf field if non-nil, zero value otherwise.

### GetConfigOspfOk

`func (o *OnpremHost) GetConfigOspfOk() (*OspfConfig, bool)`

GetConfigOspfOk returns a tuple with the ConfigOspf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOspf

`func (o *OnpremHost) SetConfigOspf(v OspfConfig)`

SetConfigOspf sets ConfigOspf field to given value.

### HasConfigOspf

`func (o *OnpremHost) HasConfigOspf() bool`

HasConfigOspf returns a boolean if a field has been set.

### GetConfigOspfv3

`func (o *OnpremHost) GetConfigOspfv3() Ospfv3Config`

GetConfigOspfv3 returns the ConfigOspfv3 field if non-nil, zero value otherwise.

### GetConfigOspfv3Ok

`func (o *OnpremHost) GetConfigOspfv3Ok() (*Ospfv3Config, bool)`

GetConfigOspfv3Ok returns a tuple with the ConfigOspfv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOspfv3

`func (o *OnpremHost) SetConfigOspfv3(v Ospfv3Config)`

SetConfigOspfv3 sets ConfigOspfv3 field to given value.

### HasConfigOspfv3

`func (o *OnpremHost) HasConfigOspfv3() bool`

HasConfigOspfv3 returns a boolean if a field has been set.

### GetCreatedAt

`func (o *OnpremHost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OnpremHost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OnpremHost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OnpremHost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *OnpremHost) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnpremHost) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnpremHost) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OnpremHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *OnpremHost) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *OnpremHost) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *OnpremHost) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *OnpremHost) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *OnpremHost) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *OnpremHost) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *OnpremHost) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *OnpremHost) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetName

`func (o *OnpremHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OnpremHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OnpremHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OnpremHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *OnpremHost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OnpremHost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OnpremHost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *OnpremHost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


