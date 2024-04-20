# Filter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for the DHCP filter. May contain 0 to 1024 characters. Can include UTF-8. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the DHCP filter. | [optional] [readonly] 
**Protocol** | Pointer to **string** | The type of protocol of the filter (_ip4_ or _ip6_). | [optional] [readonly] 
**Role** | Pointer to **string** | The role of DHCP filter (_values_ or _selection_). | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for the DHCP filter in JSON format. | [optional] 
**Type** | Pointer to **string** | The type of DHCP filter (_hardware_ or _option_). | [optional] [readonly] 

## Methods

### NewFilter

`func NewFilter() *Filter`

NewFilter instantiates a new Filter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterWithDefaults

`func NewFilterWithDefaults() *Filter`

NewFilterWithDefaults instantiates a new Filter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Filter) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Filter) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Filter) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Filter) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *Filter) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Filter) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Filter) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Filter) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Filter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Filter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Filter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Filter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtocol

`func (o *Filter) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *Filter) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *Filter) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *Filter) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetRole

`func (o *Filter) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *Filter) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *Filter) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *Filter) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTags

`func (o *Filter) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Filter) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Filter) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *Filter) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Filter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Filter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Filter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Filter) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


