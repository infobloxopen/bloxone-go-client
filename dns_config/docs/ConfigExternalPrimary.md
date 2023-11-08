# ConfigExternalPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Optional. Required only if _type_ is _server_. IP Address of nameserver. | [optional] 
**Fqdn** | Pointer to **string** | Optional. Required only if _type_ is _server_. FQDN of nameserver. | [optional] 
**Nsg** | Pointer to **string** | The resource identifier. | [optional] 
**ProtocolFqdn** | Pointer to **string** | FQDN of nameserver in punycode. | [optional] [readonly] 
**TsigEnabled** | Pointer to **bool** | Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary. | [optional] 
**TsigKey** | Pointer to [**ConfigTSIGKey**](ConfigTSIGKey.md) |  | [optional] 
**Type** | **string** | Allowed values: * _nsg_, * _primary_. | 

## Methods

### NewConfigExternalPrimary

`func NewConfigExternalPrimary(type_ string, ) *ConfigExternalPrimary`

NewConfigExternalPrimary instantiates a new ConfigExternalPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigExternalPrimaryWithDefaults

`func NewConfigExternalPrimaryWithDefaults() *ConfigExternalPrimary`

NewConfigExternalPrimaryWithDefaults instantiates a new ConfigExternalPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConfigExternalPrimary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConfigExternalPrimary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConfigExternalPrimary) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConfigExternalPrimary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFqdn

`func (o *ConfigExternalPrimary) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigExternalPrimary) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigExternalPrimary) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ConfigExternalPrimary) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetNsg

`func (o *ConfigExternalPrimary) GetNsg() string`

GetNsg returns the Nsg field if non-nil, zero value otherwise.

### GetNsgOk

`func (o *ConfigExternalPrimary) GetNsgOk() (*string, bool)`

GetNsgOk returns a tuple with the Nsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsg

`func (o *ConfigExternalPrimary) SetNsg(v string)`

SetNsg sets Nsg field to given value.

### HasNsg

`func (o *ConfigExternalPrimary) HasNsg() bool`

HasNsg returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *ConfigExternalPrimary) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ConfigExternalPrimary) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ConfigExternalPrimary) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ConfigExternalPrimary) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetTsigEnabled

`func (o *ConfigExternalPrimary) GetTsigEnabled() bool`

GetTsigEnabled returns the TsigEnabled field if non-nil, zero value otherwise.

### GetTsigEnabledOk

`func (o *ConfigExternalPrimary) GetTsigEnabledOk() (*bool, bool)`

GetTsigEnabledOk returns a tuple with the TsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigEnabled

`func (o *ConfigExternalPrimary) SetTsigEnabled(v bool)`

SetTsigEnabled sets TsigEnabled field to given value.

### HasTsigEnabled

`func (o *ConfigExternalPrimary) HasTsigEnabled() bool`

HasTsigEnabled returns a boolean if a field has been set.

### GetTsigKey

`func (o *ConfigExternalPrimary) GetTsigKey() ConfigTSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *ConfigExternalPrimary) GetTsigKeyOk() (*ConfigTSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *ConfigExternalPrimary) SetTsigKey(v ConfigTSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *ConfigExternalPrimary) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetType

`func (o *ConfigExternalPrimary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigExternalPrimary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigExternalPrimary) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


