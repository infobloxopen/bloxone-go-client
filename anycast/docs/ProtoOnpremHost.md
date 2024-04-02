# ProtoOnpremHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnycastConfigRefs** | Pointer to [**[]ProtoAnycastConfigRef**](ProtoAnycastConfigRef.md) |  | [optional] 
**ConfigBgp** | Pointer to [**ProtoBgpConfig**](ProtoBgpConfig.md) |  | [optional] 
**ConfigOspf** | Pointer to [**ProtoOspfConfig**](ProtoOspfConfig.md) |  | [optional] 
**ConfigOspfv3** | Pointer to [**ProtoOspfv3Config**](ProtoOspfv3Config.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] [readonly] 
**IpAddress** | Pointer to **string** | IPv4 address of the on-prem host | [optional] 
**Ipv6Address** | Pointer to **string** | IPv6 address of the on-prem host | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProtoOnpremHost

`func NewProtoOnpremHost() *ProtoOnpremHost`

NewProtoOnpremHost instantiates a new ProtoOnpremHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoOnpremHostWithDefaults

`func NewProtoOnpremHostWithDefaults() *ProtoOnpremHost`

NewProtoOnpremHostWithDefaults instantiates a new ProtoOnpremHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnycastConfigRefs

`func (o *ProtoOnpremHost) GetAnycastConfigRefs() []ProtoAnycastConfigRef`

GetAnycastConfigRefs returns the AnycastConfigRefs field if non-nil, zero value otherwise.

### GetAnycastConfigRefsOk

`func (o *ProtoOnpremHost) GetAnycastConfigRefsOk() (*[]ProtoAnycastConfigRef, bool)`

GetAnycastConfigRefsOk returns a tuple with the AnycastConfigRefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastConfigRefs

`func (o *ProtoOnpremHost) SetAnycastConfigRefs(v []ProtoAnycastConfigRef)`

SetAnycastConfigRefs sets AnycastConfigRefs field to given value.

### HasAnycastConfigRefs

`func (o *ProtoOnpremHost) HasAnycastConfigRefs() bool`

HasAnycastConfigRefs returns a boolean if a field has been set.

### GetConfigBgp

`func (o *ProtoOnpremHost) GetConfigBgp() ProtoBgpConfig`

GetConfigBgp returns the ConfigBgp field if non-nil, zero value otherwise.

### GetConfigBgpOk

`func (o *ProtoOnpremHost) GetConfigBgpOk() (*ProtoBgpConfig, bool)`

GetConfigBgpOk returns a tuple with the ConfigBgp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigBgp

`func (o *ProtoOnpremHost) SetConfigBgp(v ProtoBgpConfig)`

SetConfigBgp sets ConfigBgp field to given value.

### HasConfigBgp

`func (o *ProtoOnpremHost) HasConfigBgp() bool`

HasConfigBgp returns a boolean if a field has been set.

### GetConfigOspf

`func (o *ProtoOnpremHost) GetConfigOspf() ProtoOspfConfig`

GetConfigOspf returns the ConfigOspf field if non-nil, zero value otherwise.

### GetConfigOspfOk

`func (o *ProtoOnpremHost) GetConfigOspfOk() (*ProtoOspfConfig, bool)`

GetConfigOspfOk returns a tuple with the ConfigOspf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOspf

`func (o *ProtoOnpremHost) SetConfigOspf(v ProtoOspfConfig)`

SetConfigOspf sets ConfigOspf field to given value.

### HasConfigOspf

`func (o *ProtoOnpremHost) HasConfigOspf() bool`

HasConfigOspf returns a boolean if a field has been set.

### GetConfigOspfv3

`func (o *ProtoOnpremHost) GetConfigOspfv3() ProtoOspfv3Config`

GetConfigOspfv3 returns the ConfigOspfv3 field if non-nil, zero value otherwise.

### GetConfigOspfv3Ok

`func (o *ProtoOnpremHost) GetConfigOspfv3Ok() (*ProtoOspfv3Config, bool)`

GetConfigOspfv3Ok returns a tuple with the ConfigOspfv3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigOspfv3

`func (o *ProtoOnpremHost) SetConfigOspfv3(v ProtoOspfv3Config)`

SetConfigOspfv3 sets ConfigOspfv3 field to given value.

### HasConfigOspfv3

`func (o *ProtoOnpremHost) HasConfigOspfv3() bool`

HasConfigOspfv3 returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ProtoOnpremHost) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProtoOnpremHost) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProtoOnpremHost) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ProtoOnpremHost) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ProtoOnpremHost) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProtoOnpremHost) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProtoOnpremHost) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProtoOnpremHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ProtoOnpremHost) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ProtoOnpremHost) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ProtoOnpremHost) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ProtoOnpremHost) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetIpv6Address

`func (o *ProtoOnpremHost) GetIpv6Address() string`

GetIpv6Address returns the Ipv6Address field if non-nil, zero value otherwise.

### GetIpv6AddressOk

`func (o *ProtoOnpremHost) GetIpv6AddressOk() (*string, bool)`

GetIpv6AddressOk returns a tuple with the Ipv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6Address

`func (o *ProtoOnpremHost) SetIpv6Address(v string)`

SetIpv6Address sets Ipv6Address field to given value.

### HasIpv6Address

`func (o *ProtoOnpremHost) HasIpv6Address() bool`

HasIpv6Address returns a boolean if a field has been set.

### GetName

`func (o *ProtoOnpremHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProtoOnpremHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProtoOnpremHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProtoOnpremHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ProtoOnpremHost) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProtoOnpremHost) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProtoOnpremHost) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProtoOnpremHost) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


