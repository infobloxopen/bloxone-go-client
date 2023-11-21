# ConfigForwardZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for zone configuration. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the object has been created. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**ExternalForwarders** | Pointer to [**[]ConfigForwarder**](ConfigForwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
**ForwardOnly** | Pointer to **bool** | Optional. _true_ to only forward. | [optional] 
**Fqdn** | Pointer to **string** | Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation. | [optional] 
**Hosts** | Pointer to **[]string** | The resource identifier. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InternalForwarders** | Pointer to **[]string** | The resource identifier. | [optional] 
**MappedSubnet** | Pointer to **string** | Reverse zone network address in the following format: \&quot;ip-address/cidr\&quot;. Defaults to empty. | [optional] [readonly] 
**Mapping** | Pointer to **string** | Read-only. Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to _forward_. | [optional] [readonly] 
**Nsgs** | Pointer to **[]string** | The resource identifier. | [optional] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**ProtocolFqdn** | Pointer to **string** | Zone FQDN in punycode. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation. | [optional] [readonly] 
**View** | Pointer to **string** | The resource identifier. | [optional] 
**Warnings** | Pointer to [**[]ConfigWarning**](ConfigWarning.md) | The list of a forward zone warnings. | [optional] [readonly] 

## Methods

### NewConfigForwardZone

`func NewConfigForwardZone() *ConfigForwardZone`

NewConfigForwardZone instantiates a new ConfigForwardZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigForwardZoneWithDefaults

`func NewConfigForwardZoneWithDefaults() *ConfigForwardZone`

NewConfigForwardZoneWithDefaults instantiates a new ConfigForwardZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ConfigForwardZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigForwardZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigForwardZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigForwardZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ConfigForwardZone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ConfigForwardZone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ConfigForwardZone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ConfigForwardZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *ConfigForwardZone) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ConfigForwardZone) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ConfigForwardZone) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ConfigForwardZone) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExternalForwarders

`func (o *ConfigForwardZone) GetExternalForwarders() []ConfigForwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ConfigForwardZone) GetExternalForwardersOk() (*[]ConfigForwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ConfigForwardZone) SetExternalForwarders(v []ConfigForwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ConfigForwardZone) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetForwardOnly

`func (o *ConfigForwardZone) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *ConfigForwardZone) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *ConfigForwardZone) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *ConfigForwardZone) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetFqdn

`func (o *ConfigForwardZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ConfigForwardZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ConfigForwardZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ConfigForwardZone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetHosts

`func (o *ConfigForwardZone) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ConfigForwardZone) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ConfigForwardZone) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ConfigForwardZone) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *ConfigForwardZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigForwardZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigForwardZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigForwardZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ConfigForwardZone) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ConfigForwardZone) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ConfigForwardZone) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ConfigForwardZone) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetMappedSubnet

`func (o *ConfigForwardZone) GetMappedSubnet() string`

GetMappedSubnet returns the MappedSubnet field if non-nil, zero value otherwise.

### GetMappedSubnetOk

`func (o *ConfigForwardZone) GetMappedSubnetOk() (*string, bool)`

GetMappedSubnetOk returns a tuple with the MappedSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedSubnet

`func (o *ConfigForwardZone) SetMappedSubnet(v string)`

SetMappedSubnet sets MappedSubnet field to given value.

### HasMappedSubnet

`func (o *ConfigForwardZone) HasMappedSubnet() bool`

HasMappedSubnet returns a boolean if a field has been set.

### GetMapping

`func (o *ConfigForwardZone) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *ConfigForwardZone) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *ConfigForwardZone) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *ConfigForwardZone) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetNsgs

`func (o *ConfigForwardZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ConfigForwardZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ConfigForwardZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ConfigForwardZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetParent

`func (o *ConfigForwardZone) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ConfigForwardZone) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ConfigForwardZone) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ConfigForwardZone) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *ConfigForwardZone) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ConfigForwardZone) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ConfigForwardZone) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ConfigForwardZone) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetTags

`func (o *ConfigForwardZone) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigForwardZone) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigForwardZone) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigForwardZone) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ConfigForwardZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ConfigForwardZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ConfigForwardZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ConfigForwardZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetView

`func (o *ConfigForwardZone) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ConfigForwardZone) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ConfigForwardZone) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ConfigForwardZone) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWarnings

`func (o *ConfigForwardZone) GetWarnings() []ConfigWarning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ConfigForwardZone) GetWarningsOk() (*[]ConfigWarning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ConfigForwardZone) SetWarnings(v []ConfigWarning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ConfigForwardZone) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


