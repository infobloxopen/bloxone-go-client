# ECSZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | Access control for zone.  Allowed values: * _allow_, * _deny_. | 
**Fqdn** | **string** | Zone FQDN. | 
**ProtocolFqdn** | Pointer to **string** | Zone FQDN in punycode. | [optional] [readonly] 

## Methods

### NewECSZone

`func NewECSZone(access string, fqdn string, ) *ECSZone`

NewECSZone instantiates a new ECSZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewECSZoneWithDefaults

`func NewECSZoneWithDefaults() *ECSZone`

NewECSZoneWithDefaults instantiates a new ECSZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ECSZone) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ECSZone) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ECSZone) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetFqdn

`func (o *ECSZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ECSZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ECSZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetProtocolFqdn

`func (o *ECSZone) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ECSZone) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ECSZone) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ECSZone) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


