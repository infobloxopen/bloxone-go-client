# ExternalPrimary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Optional. Required only if _type_ is _server_. IP Address of nameserver. | [optional] 
**Fqdn** | Pointer to **string** | Optional. Required only if _type_ is _server_. FQDN of nameserver. | [optional] 
**Nsg** | Pointer to **string** | The resource identifier. | [optional] 
**ProtocolFqdn** | Pointer to **string** | FQDN of nameserver in punycode. | [optional] [readonly] 
**TsigEnabled** | Pointer to **bool** | Optional. If enabled, secondaries will use the configured TSIG key when requesting a zone transfer from this primary. | [optional] 
**TsigKey** | Pointer to [**TSIGKey**](TSIGKey.md) |  | [optional] 
**Type** | **string** | Allowed values: * _nsg_, * _primary_. | 

## Methods

### NewExternalPrimary

`func NewExternalPrimary(type_ string, ) *ExternalPrimary`

NewExternalPrimary instantiates a new ExternalPrimary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPrimaryWithDefaults

`func NewExternalPrimaryWithDefaults() *ExternalPrimary`

NewExternalPrimaryWithDefaults instantiates a new ExternalPrimary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ExternalPrimary) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ExternalPrimary) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ExternalPrimary) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ExternalPrimary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFqdn

`func (o *ExternalPrimary) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ExternalPrimary) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ExternalPrimary) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ExternalPrimary) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetNsg

`func (o *ExternalPrimary) GetNsg() string`

GetNsg returns the Nsg field if non-nil, zero value otherwise.

### GetNsgOk

`func (o *ExternalPrimary) GetNsgOk() (*string, bool)`

GetNsgOk returns a tuple with the Nsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsg

`func (o *ExternalPrimary) SetNsg(v string)`

SetNsg sets Nsg field to given value.

### HasNsg

`func (o *ExternalPrimary) HasNsg() bool`

HasNsg returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *ExternalPrimary) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ExternalPrimary) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ExternalPrimary) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ExternalPrimary) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetTsigEnabled

`func (o *ExternalPrimary) GetTsigEnabled() bool`

GetTsigEnabled returns the TsigEnabled field if non-nil, zero value otherwise.

### GetTsigEnabledOk

`func (o *ExternalPrimary) GetTsigEnabledOk() (*bool, bool)`

GetTsigEnabledOk returns a tuple with the TsigEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigEnabled

`func (o *ExternalPrimary) SetTsigEnabled(v bool)`

SetTsigEnabled sets TsigEnabled field to given value.

### HasTsigEnabled

`func (o *ExternalPrimary) HasTsigEnabled() bool`

HasTsigEnabled returns a boolean if a field has been set.

### GetTsigKey

`func (o *ExternalPrimary) GetTsigKey() TSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *ExternalPrimary) GetTsigKeyOk() (*TSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *ExternalPrimary) SetTsigKey(v TSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *ExternalPrimary) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.

### GetType

`func (o *ExternalPrimary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExternalPrimary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExternalPrimary) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


