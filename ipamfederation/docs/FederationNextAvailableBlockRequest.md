# FederationNextAvailableBlockRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | Pointer to **int64** | The CIDR of the federated block. This is required, if _address_ does not specify it in its input. | [optional] 
**Comment** | Pointer to **string** | The description for the _federation/federated_block_. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Count** | Pointer to **int64** | The count of __Block__ required. If not provided, it will default to 1. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name to be provided. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the federated block in JSON format. | [optional] 

## Methods

### NewFederationNextAvailableBlockRequest

`func NewFederationNextAvailableBlockRequest() *FederationNextAvailableBlockRequest`

NewFederationNextAvailableBlockRequest instantiates a new FederationNextAvailableBlockRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationNextAvailableBlockRequestWithDefaults

`func NewFederationNextAvailableBlockRequestWithDefaults() *FederationNextAvailableBlockRequest`

NewFederationNextAvailableBlockRequestWithDefaults instantiates a new FederationNextAvailableBlockRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *FederationNextAvailableBlockRequest) GetCidr() int64`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *FederationNextAvailableBlockRequest) GetCidrOk() (*int64, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *FederationNextAvailableBlockRequest) SetCidr(v int64)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *FederationNextAvailableBlockRequest) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetComment

`func (o *FederationNextAvailableBlockRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *FederationNextAvailableBlockRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *FederationNextAvailableBlockRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *FederationNextAvailableBlockRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCount

`func (o *FederationNextAvailableBlockRequest) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *FederationNextAvailableBlockRequest) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *FederationNextAvailableBlockRequest) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *FederationNextAvailableBlockRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetId

`func (o *FederationNextAvailableBlockRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationNextAvailableBlockRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationNextAvailableBlockRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederationNextAvailableBlockRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FederationNextAvailableBlockRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationNextAvailableBlockRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationNextAvailableBlockRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederationNextAvailableBlockRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *FederationNextAvailableBlockRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *FederationNextAvailableBlockRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *FederationNextAvailableBlockRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *FederationNextAvailableBlockRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


