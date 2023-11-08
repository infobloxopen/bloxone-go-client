# InheritanceAssignedHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | The human-readable display name for the host referred to by _ophid_. | [optional] [readonly] 
**Host** | Pointer to **string** | The resource identifier. | [optional] 
**Ophid** | Pointer to **string** | The on-prem host ID. | [optional] [readonly] 

## Methods

### NewInheritanceAssignedHost

`func NewInheritanceAssignedHost() *InheritanceAssignedHost`

NewInheritanceAssignedHost instantiates a new InheritanceAssignedHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInheritanceAssignedHostWithDefaults

`func NewInheritanceAssignedHostWithDefaults() *InheritanceAssignedHost`

NewInheritanceAssignedHostWithDefaults instantiates a new InheritanceAssignedHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *InheritanceAssignedHost) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InheritanceAssignedHost) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InheritanceAssignedHost) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InheritanceAssignedHost) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHost

`func (o *InheritanceAssignedHost) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *InheritanceAssignedHost) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *InheritanceAssignedHost) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *InheritanceAssignedHost) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetOphid

`func (o *InheritanceAssignedHost) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *InheritanceAssignedHost) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *InheritanceAssignedHost) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *InheritanceAssignedHost) HasOphid() bool`

HasOphid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


