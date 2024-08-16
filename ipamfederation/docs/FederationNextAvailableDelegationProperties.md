# FederationNextAvailableDelegationProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the __Delegation__. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**DelegatedTo** | Pointer to **string** | The specific IPAM service the __Delegation__ was delegated to. | [optional] 
**Name** | Pointer to **string** | The name to be provided. | [optional] 

## Methods

### NewFederationNextAvailableDelegationProperties

`func NewFederationNextAvailableDelegationProperties() *FederationNextAvailableDelegationProperties`

NewFederationNextAvailableDelegationProperties instantiates a new FederationNextAvailableDelegationProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationNextAvailableDelegationPropertiesWithDefaults

`func NewFederationNextAvailableDelegationPropertiesWithDefaults() *FederationNextAvailableDelegationProperties`

NewFederationNextAvailableDelegationPropertiesWithDefaults instantiates a new FederationNextAvailableDelegationProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *FederationNextAvailableDelegationProperties) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederationNextAvailableDelegationProperties) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederationNextAvailableDelegationProperties) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederationNextAvailableDelegationProperties) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegatedTo

`func (o *FederationNextAvailableDelegationProperties) GetDelegatedTo() string`

GetDelegatedTo returns the DelegatedTo field if non-nil, zero value otherwise.

### GetDelegatedToOk

`func (o *FederationNextAvailableDelegationProperties) GetDelegatedToOk() (*string, bool)`

GetDelegatedToOk returns a tuple with the DelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTo

`func (o *FederationNextAvailableDelegationProperties) SetDelegatedTo(v string)`

SetDelegatedTo sets DelegatedTo field to given value.

### HasDelegatedTo

`func (o *FederationNextAvailableDelegationProperties) HasDelegatedTo() bool`

HasDelegatedTo returns a boolean if a field has been set.

### GetName

`func (o *FederationNextAvailableDelegationProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationNextAvailableDelegationProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationNextAvailableDelegationProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederationNextAvailableDelegationProperties) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


