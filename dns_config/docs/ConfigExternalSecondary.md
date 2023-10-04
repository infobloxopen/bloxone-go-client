# ConfigExternalSecondary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | IP Address of nameserver. | 
**Fqdn** | **string** | FQDN of nameserver. | 
**ProtocolFqdn** | Pointer to **string** | FQDN of nameserver in punycode. | [optional] [readonly] 
**Stealth** | Pointer to **bool** | If enabled, the NS record and glue record will NOT be automatically generated according to secondaries nameserver assignment.  Default: _false_ | [optional] 
**TsigEnabled** | Pointer to **bool** | If enabled, secondaries will use the configured TSIG key when requesting a zone transfer.  Default: _false_ | [optional] 
**TsigKey** | Pointer to [**ConfigTSIGKey**](ConfigTSIGKey.md) |  | [optional] 

## Methods

### NewConfigExternalSecondary

`func NewConfigExternalSecondary(address string, fqdn string, ) *ConfigExternalSecondary`

NewConfigExternalSecondary instantiates a new ConfigExternalSecondary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigExternalSecondaryWithDefaults

`func NewConfigExternalSecondaryWithDefaults() *ConfigExternalSecondary`

NewConfigExternalSecondaryWithDefaults instantiates a new ConfigExternalSecondary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ConfigExternalSecondary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConfigExternalSecondary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConfigExternalSecondary) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetFqdn

`func (o *ConfigExternalSecondary) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigExternalSecondary) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigExternalSecondary) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.


### GetProtocolFqdn

`func (o *ConfigExternalSecondary) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ConfigExternalSecondary) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ConfigExternalSecondary) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ConfigExternalSecondary) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetStealth

`func (o *ConfigExternalSecondary) GetStealth() bool`

GetStealth returns the Stealth field if non-nil, zero value otherwise.

### GetStealthOk

`func (o *ConfigExternalSecondary) GetStealthOk() (*bool, bool)`

GetStealthOk returns a tuple with the Stealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStealth

`func (o *ConfigExternalSecondary) SetStealth(v bool)`

SetStealth sets Stealth field to given value.

### HasStealth

`func (o *ConfigExternalSecondary) HasStealth() bool`

HasStealth returns a boolean if a field has been set.

### GetTsigEnabled

`func (o *ConfigExternalSecondary) GetTsigEnabled() bool`

GetTsigEnabled returns the TsigEnabled field if non-nil, zero value otherwise.

### GetTsigEnabledOk

`func (o *ConfigExternalSecondary) GetTsigEnabledOk() (*bool, bool)`

GetTsigEnabledOk returns a tuple with the TsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigEnabled

`func (o *ConfigExternalSecondary) SetTsigEnabled(v bool)`

SetTsigEnabled sets TsigEnabled field to given value.

### HasTsigEnabled

`func (o *ConfigExternalSecondary) HasTsigEnabled() bool`

HasTsigEnabled returns a boolean if a field has been set.

### GetTsigKey

`func (o *ConfigExternalSecondary) GetTsigKey() ConfigTSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *ConfigExternalSecondary) GetTsigKeyOk() (*ConfigTSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *ConfigExternalSecondary) SetTsigKey(v ConfigTSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *ConfigExternalSecondary) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


