# UploadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** | The description for uploaded content. May contain 0 to 1024 characters. Can include UTF-8. | [optional] 
**Content** | **string** | Base64 encoded content. | 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags for uploaded content in JSON format. | [optional] 
**Type** | [**UploadContentType**](UploadContentType.md) |  | [default to UPLOADCONTENTTYPE_UNKNOWN]

## Methods

### NewUploadRequest

`func NewUploadRequest(content string, type_ UploadContentType, ) *UploadRequest`

NewUploadRequest instantiates a new UploadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadRequestWithDefaults

`func NewUploadRequestWithDefaults() *UploadRequest`

NewUploadRequestWithDefaults instantiates a new UploadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *UploadRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UploadRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UploadRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UploadRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetContent

`func (o *UploadRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UploadRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UploadRequest) SetContent(v string)`

SetContent sets Content field to given value.


### GetFields

`func (o *UploadRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *UploadRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *UploadRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *UploadRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetTags

`func (o *UploadRequest) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UploadRequest) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UploadRequest) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *UploadRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *UploadRequest) GetType() UploadContentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadRequest) GetTypeOk() (*UploadContentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadRequest) SetType(v UploadContentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


