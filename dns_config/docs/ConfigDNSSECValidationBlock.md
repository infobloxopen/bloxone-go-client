# ConfigDNSSECValidationBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnssecEnableValidation** | Pointer to **bool** | Optional. Field config for _dnssec_enable_validation_ field. | [optional] 
**DnssecEnabled** | Pointer to **bool** | Optional. Field config for _dnssec_enabled_ field. | [optional] 
**DnssecTrustAnchors** | Pointer to [**[]ConfigTrustAnchor**](ConfigTrustAnchor.md) | Optional. Field config for _dnssec_trust_anchors_ field. | [optional] 
**DnssecValidateExpiry** | Pointer to **bool** | Optional. Field config for _dnssec_validate_expiry_ field. | [optional] 

## Methods

### NewConfigDNSSECValidationBlock

`func NewConfigDNSSECValidationBlock() *ConfigDNSSECValidationBlock`

NewConfigDNSSECValidationBlock instantiates a new ConfigDNSSECValidationBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigDNSSECValidationBlockWithDefaults

`func NewConfigDNSSECValidationBlockWithDefaults() *ConfigDNSSECValidationBlock`

NewConfigDNSSECValidationBlockWithDefaults instantiates a new ConfigDNSSECValidationBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnssecEnableValidation

`func (o *ConfigDNSSECValidationBlock) GetDnssecEnableValidation() bool`

GetDnssecEnableValidation returns the DnssecEnableValidation field if non-nil, zero value otherwise.

### GetDnssecEnableValidationOk

`func (o *ConfigDNSSECValidationBlock) GetDnssecEnableValidationOk() (*bool, bool)`

GetDnssecEnableValidationOk returns a tuple with the DnssecEnableValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnableValidation

`func (o *ConfigDNSSECValidationBlock) SetDnssecEnableValidation(v bool)`

SetDnssecEnableValidation sets DnssecEnableValidation field to given value.

### HasDnssecEnableValidation

`func (o *ConfigDNSSECValidationBlock) HasDnssecEnableValidation() bool`

HasDnssecEnableValidation returns a boolean if a field has been set.

### GetDnssecEnabled

`func (o *ConfigDNSSECValidationBlock) GetDnssecEnabled() bool`

GetDnssecEnabled returns the DnssecEnabled field if non-nil, zero value otherwise.

### GetDnssecEnabledOk

`func (o *ConfigDNSSECValidationBlock) GetDnssecEnabledOk() (*bool, bool)`

GetDnssecEnabledOk returns a tuple with the DnssecEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecEnabled

`func (o *ConfigDNSSECValidationBlock) SetDnssecEnabled(v bool)`

SetDnssecEnabled sets DnssecEnabled field to given value.

### HasDnssecEnabled

`func (o *ConfigDNSSECValidationBlock) HasDnssecEnabled() bool`

HasDnssecEnabled returns a boolean if a field has been set.

### GetDnssecTrustAnchors

`func (o *ConfigDNSSECValidationBlock) GetDnssecTrustAnchors() []ConfigTrustAnchor`

GetDnssecTrustAnchors returns the DnssecTrustAnchors field if non-nil, zero value otherwise.

### GetDnssecTrustAnchorsOk

`func (o *ConfigDNSSECValidationBlock) GetDnssecTrustAnchorsOk() (*[]ConfigTrustAnchor, bool)`

GetDnssecTrustAnchorsOk returns a tuple with the DnssecTrustAnchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecTrustAnchors

`func (o *ConfigDNSSECValidationBlock) SetDnssecTrustAnchors(v []ConfigTrustAnchor)`

SetDnssecTrustAnchors sets DnssecTrustAnchors field to given value.

### HasDnssecTrustAnchors

`func (o *ConfigDNSSECValidationBlock) HasDnssecTrustAnchors() bool`

HasDnssecTrustAnchors returns a boolean if a field has been set.

### GetDnssecValidateExpiry

`func (o *ConfigDNSSECValidationBlock) GetDnssecValidateExpiry() bool`

GetDnssecValidateExpiry returns the DnssecValidateExpiry field if non-nil, zero value otherwise.

### GetDnssecValidateExpiryOk

`func (o *ConfigDNSSECValidationBlock) GetDnssecValidateExpiryOk() (*bool, bool)`

GetDnssecValidateExpiryOk returns a tuple with the DnssecValidateExpiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnssecValidateExpiry

`func (o *ConfigDNSSECValidationBlock) SetDnssecValidateExpiry(v bool)`

SetDnssecValidateExpiry sets DnssecValidateExpiry field to given value.

### HasDnssecValidateExpiry

`func (o *ConfigDNSSECValidationBlock) HasDnssecValidateExpiry() bool`

HasDnssecValidateExpiry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


