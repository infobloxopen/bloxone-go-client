# ForwardZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for zone configuration. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the object has been created. | [optional] [readonly] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration. | [optional] 
**ExternalForwarders** | Pointer to [**[]Forwarder**](Forwarder.md) | Optional. External DNS servers to forward to. Order is not significant. | [optional] 
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
**Warnings** | Pointer to [**[]Warning**](Warning.md) | The list of a forward zone warnings. | [optional] [readonly] 

## Methods

### NewForwardZone

`func NewForwardZone() *ForwardZone`

NewForwardZone instantiates a new ForwardZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForwardZoneWithDefaults

`func NewForwardZoneWithDefaults() *ForwardZone`

NewForwardZoneWithDefaults instantiates a new ForwardZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *ForwardZone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ForwardZone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ForwardZone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ForwardZone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ForwardZone) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ForwardZone) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ForwardZone) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ForwardZone) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *ForwardZone) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ForwardZone) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ForwardZone) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ForwardZone) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetExternalForwarders

`func (o *ForwardZone) GetExternalForwarders() []Forwarder`

GetExternalForwarders returns the ExternalForwarders field if non-nil, zero value otherwise.

### GetExternalForwardersOk

`func (o *ForwardZone) GetExternalForwardersOk() (*[]Forwarder, bool)`

GetExternalForwardersOk returns a tuple with the ExternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalForwarders

`func (o *ForwardZone) SetExternalForwarders(v []Forwarder)`

SetExternalForwarders sets ExternalForwarders field to given value.

### HasExternalForwarders

`func (o *ForwardZone) HasExternalForwarders() bool`

HasExternalForwarders returns a boolean if a field has been set.

### GetForwardOnly

`func (o *ForwardZone) GetForwardOnly() bool`

GetForwardOnly returns the ForwardOnly field if non-nil, zero value otherwise.

### GetForwardOnlyOk

`func (o *ForwardZone) GetForwardOnlyOk() (*bool, bool)`

GetForwardOnlyOk returns a tuple with the ForwardOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardOnly

`func (o *ForwardZone) SetForwardOnly(v bool)`

SetForwardOnly sets ForwardOnly field to given value.

### HasForwardOnly

`func (o *ForwardZone) HasForwardOnly() bool`

HasForwardOnly returns a boolean if a field has been set.

### GetFqdn

`func (o *ForwardZone) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *ForwardZone) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *ForwardZone) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *ForwardZone) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetHosts

`func (o *ForwardZone) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *ForwardZone) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *ForwardZone) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *ForwardZone) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *ForwardZone) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ForwardZone) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ForwardZone) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ForwardZone) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInternalForwarders

`func (o *ForwardZone) GetInternalForwarders() []string`

GetInternalForwarders returns the InternalForwarders field if non-nil, zero value otherwise.

### GetInternalForwardersOk

`func (o *ForwardZone) GetInternalForwardersOk() (*[]string, bool)`

GetInternalForwardersOk returns a tuple with the InternalForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalForwarders

`func (o *ForwardZone) SetInternalForwarders(v []string)`

SetInternalForwarders sets InternalForwarders field to given value.

### HasInternalForwarders

`func (o *ForwardZone) HasInternalForwarders() bool`

HasInternalForwarders returns a boolean if a field has been set.

### GetMappedSubnet

`func (o *ForwardZone) GetMappedSubnet() string`

GetMappedSubnet returns the MappedSubnet field if non-nil, zero value otherwise.

### GetMappedSubnetOk

`func (o *ForwardZone) GetMappedSubnetOk() (*string, bool)`

GetMappedSubnetOk returns a tuple with the MappedSubnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedSubnet

`func (o *ForwardZone) SetMappedSubnet(v string)`

SetMappedSubnet sets MappedSubnet field to given value.

### HasMappedSubnet

`func (o *ForwardZone) HasMappedSubnet() bool`

HasMappedSubnet returns a boolean if a field has been set.

### GetMapping

`func (o *ForwardZone) GetMapping() string`

GetMapping returns the Mapping field if non-nil, zero value otherwise.

### GetMappingOk

`func (o *ForwardZone) GetMappingOk() (*string, bool)`

GetMappingOk returns a tuple with the Mapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapping

`func (o *ForwardZone) SetMapping(v string)`

SetMapping sets Mapping field to given value.

### HasMapping

`func (o *ForwardZone) HasMapping() bool`

HasMapping returns a boolean if a field has been set.

### GetNsgs

`func (o *ForwardZone) GetNsgs() []string`

GetNsgs returns the Nsgs field if non-nil, zero value otherwise.

### GetNsgsOk

`func (o *ForwardZone) GetNsgsOk() (*[]string, bool)`

GetNsgsOk returns a tuple with the Nsgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsgs

`func (o *ForwardZone) SetNsgs(v []string)`

SetNsgs sets Nsgs field to given value.

### HasNsgs

`func (o *ForwardZone) HasNsgs() bool`

HasNsgs returns a boolean if a field has been set.

### GetParent

`func (o *ForwardZone) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ForwardZone) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ForwardZone) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ForwardZone) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *ForwardZone) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *ForwardZone) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *ForwardZone) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *ForwardZone) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetTags

`func (o *ForwardZone) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ForwardZone) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ForwardZone) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ForwardZone) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ForwardZone) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ForwardZone) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ForwardZone) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ForwardZone) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetView

`func (o *ForwardZone) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *ForwardZone) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *ForwardZone) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *ForwardZone) HasView() bool`

HasView returns a boolean if a field has been set.

### GetWarnings

`func (o *ForwardZone) GetWarnings() []Warning`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ForwardZone) GetWarningsOk() (*[]Warning, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ForwardZone) SetWarnings(v []Warning)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ForwardZone) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


