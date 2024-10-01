# CredentialConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessIdentifier** | Pointer to **string** |  | [optional] 
**AgentId** | Pointer to **string** |  | [optional] 
**Enclave** | Pointer to **string** |  | [optional] 
**ForestGuid** | Pointer to **string** |  | [optional] 
**MsServers** | Pointer to [**[]MSServers**](MSServers.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewCredentialConfig

`func NewCredentialConfig() *CredentialConfig`

NewCredentialConfig instantiates a new CredentialConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialConfigWithDefaults

`func NewCredentialConfigWithDefaults() *CredentialConfig`

NewCredentialConfigWithDefaults instantiates a new CredentialConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessIdentifier

`func (o *CredentialConfig) GetAccessIdentifier() string`

GetAccessIdentifier returns the AccessIdentifier field if non-nil, zero value otherwise.

### GetAccessIdentifierOk

`func (o *CredentialConfig) GetAccessIdentifierOk() (*string, bool)`

GetAccessIdentifierOk returns a tuple with the AccessIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIdentifier

`func (o *CredentialConfig) SetAccessIdentifier(v string)`

SetAccessIdentifier sets AccessIdentifier field to given value.

### HasAccessIdentifier

`func (o *CredentialConfig) HasAccessIdentifier() bool`

HasAccessIdentifier returns a boolean if a field has been set.

### GetAgentId

`func (o *CredentialConfig) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CredentialConfig) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CredentialConfig) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CredentialConfig) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetEnclave

`func (o *CredentialConfig) GetEnclave() string`

GetEnclave returns the Enclave field if non-nil, zero value otherwise.

### GetEnclaveOk

`func (o *CredentialConfig) GetEnclaveOk() (*string, bool)`

GetEnclaveOk returns a tuple with the Enclave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnclave

`func (o *CredentialConfig) SetEnclave(v string)`

SetEnclave sets Enclave field to given value.

### HasEnclave

`func (o *CredentialConfig) HasEnclave() bool`

HasEnclave returns a boolean if a field has been set.

### GetForestGuid

`func (o *CredentialConfig) GetForestGuid() string`

GetForestGuid returns the ForestGuid field if non-nil, zero value otherwise.

### GetForestGuidOk

`func (o *CredentialConfig) GetForestGuidOk() (*string, bool)`

GetForestGuidOk returns a tuple with the ForestGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestGuid

`func (o *CredentialConfig) SetForestGuid(v string)`

SetForestGuid sets ForestGuid field to given value.

### HasForestGuid

`func (o *CredentialConfig) HasForestGuid() bool`

HasForestGuid returns a boolean if a field has been set.

### GetMsServers

`func (o *CredentialConfig) GetMsServers() []MSServers`

GetMsServers returns the MsServers field if non-nil, zero value otherwise.

### GetMsServersOk

`func (o *CredentialConfig) GetMsServersOk() (*[]MSServers, bool)`

GetMsServersOk returns a tuple with the MsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsServers

`func (o *CredentialConfig) SetMsServers(v []MSServers)`

SetMsServers sets MsServers field to given value.

### HasMsServers

`func (o *CredentialConfig) HasMsServers() bool`

HasMsServers returns a boolean if a field has been set.

### GetRegion

`func (o *CredentialConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CredentialConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CredentialConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CredentialConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


