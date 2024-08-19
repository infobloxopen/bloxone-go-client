# NextAvailableDelegationProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the __Delegation__. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**DelegatedTo** | Pointer to **string** | The specific IPAM service the __Delegation__ was delegated to. | [optional] 
**Name** | Pointer to **string** | The name to be provided. | [optional] 

## Methods

### NewNextAvailableDelegationProperties

`func NewNextAvailableDelegationProperties() *NextAvailableDelegationProperties`

NewNextAvailableDelegationProperties instantiates a new NextAvailableDelegationProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNextAvailableDelegationPropertiesWithDefaults

`func NewNextAvailableDelegationPropertiesWithDefaults() *NextAvailableDelegationProperties`

NewNextAvailableDelegationPropertiesWithDefaults instantiates a new NextAvailableDelegationProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *NextAvailableDelegationProperties) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NextAvailableDelegationProperties) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NextAvailableDelegationProperties) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NextAvailableDelegationProperties) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDelegatedTo

`func (o *NextAvailableDelegationProperties) GetDelegatedTo() string`

GetDelegatedTo returns the DelegatedTo field if non-nil, zero value otherwise.

### GetDelegatedToOk

`func (o *NextAvailableDelegationProperties) GetDelegatedToOk() (*string, bool)`

GetDelegatedToOk returns a tuple with the DelegatedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegatedTo

`func (o *NextAvailableDelegationProperties) SetDelegatedTo(v string)`

SetDelegatedTo sets DelegatedTo field to given value.

### HasDelegatedTo

`func (o *NextAvailableDelegationProperties) HasDelegatedTo() bool`

HasDelegatedTo returns a boolean if a field has been set.

### GetName

`func (o *NextAvailableDelegationProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NextAvailableDelegationProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NextAvailableDelegationProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NextAvailableDelegationProperties) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


