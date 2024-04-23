# Delegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | Optional. Comment for zone delegation. | [optional] 
**DelegationServers** | Pointer to [**[]DelegationServer**](DelegationServer.md) | Required. DNS zone delegation servers. Order is not significant. | [optional] 
**Disabled** | Pointer to **bool** | Optional. _true_ to disable object. A disabled object is effectively non-existent when generating resource records. | [optional] 
**Fqdn** | Pointer to **string** | Delegation FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Parent** | Pointer to **string** | The resource identifier. | [optional] 
**ProtocolFqdn** | Pointer to **string** | Delegation FQDN in punycode. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Tagging specifics. | [optional] 
**View** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewDelegation

`func NewDelegation() *Delegation`

NewDelegation instantiates a new Delegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelegationWithDefaults

`func NewDelegationWithDefaults() *Delegation`

NewDelegationWithDefaults instantiates a new Delegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Delegation) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Delegation) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Delegation) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Delegation) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegationServers

`func (o *Delegation) GetDelegationServers() []DelegationServer`

GetDelegationServers returns the DelegationServers field if non-nil, zero value otherwise.

### GetDelegationServersOk

`func (o *Delegation) GetDelegationServersOk() (*[]DelegationServer, bool)`

GetDelegationServersOk returns a tuple with the DelegationServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegationServers

`func (o *Delegation) SetDelegationServers(v []DelegationServer)`

SetDelegationServers sets DelegationServers field to given value.

### HasDelegationServers

`func (o *Delegation) HasDelegationServers() bool`

HasDelegationServers returns a boolean if a field has been set.

### GetDisabled

`func (o *Delegation) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *Delegation) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *Delegation) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *Delegation) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetFqdn

`func (o *Delegation) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *Delegation) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *Delegation) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *Delegation) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.

### GetId

`func (o *Delegation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Delegation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Delegation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Delegation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParent

`func (o *Delegation) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Delegation) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Delegation) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *Delegation) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetProtocolFqdn

`func (o *Delegation) GetProtocolFqdn() string`

GetProtocolFqdn returns the ProtocolFqdn field if non-nil, zero value otherwise.

### GetProtocolFqdnOk

`func (o *Delegation) GetProtocolFqdnOk() (*string, bool)`

GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolFqdn

`func (o *Delegation) SetProtocolFqdn(v string)`

SetProtocolFqdn sets ProtocolFqdn field to given value.

### HasProtocolFqdn

`func (o *Delegation) HasProtocolFqdn() bool`

HasProtocolFqdn returns a boolean if a field has been set.

### GetTags

`func (o *Delegation) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Delegation) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Delegation) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Delegation) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetView

`func (o *Delegation) GetView() string`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *Delegation) GetViewOk() (*string, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *Delegation) SetView(v string)`

SetView sets View field to given value.

### HasView

`func (o *Delegation) HasView() bool`

HasView returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


