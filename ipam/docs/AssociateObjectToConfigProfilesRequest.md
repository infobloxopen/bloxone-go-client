# AssociateObjectToConfigProfilesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigProfileIds** | **[]string** | The resource identifier. | 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**ObjectId** | **string** | The resource identifier. | 

## Methods

### NewAssociateObjectToConfigProfilesRequest

`func NewAssociateObjectToConfigProfilesRequest(configProfileIds []string, objectId string, ) *AssociateObjectToConfigProfilesRequest`

NewAssociateObjectToConfigProfilesRequest instantiates a new AssociateObjectToConfigProfilesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociateObjectToConfigProfilesRequestWithDefaults

`func NewAssociateObjectToConfigProfilesRequestWithDefaults() *AssociateObjectToConfigProfilesRequest`

NewAssociateObjectToConfigProfilesRequestWithDefaults instantiates a new AssociateObjectToConfigProfilesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigProfileIds

`func (o *AssociateObjectToConfigProfilesRequest) GetConfigProfileIds() []string`

GetConfigProfileIds returns the ConfigProfileIds field if non-nil, zero value otherwise.

### GetConfigProfileIdsOk

`func (o *AssociateObjectToConfigProfilesRequest) GetConfigProfileIdsOk() (*[]string, bool)`

GetConfigProfileIdsOk returns a tuple with the ConfigProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProfileIds

`func (o *AssociateObjectToConfigProfilesRequest) SetConfigProfileIds(v []string)`

SetConfigProfileIds sets ConfigProfileIds field to given value.


### GetFields

`func (o *AssociateObjectToConfigProfilesRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AssociateObjectToConfigProfilesRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AssociateObjectToConfigProfilesRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AssociateObjectToConfigProfilesRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetObjectId

`func (o *AssociateObjectToConfigProfilesRequest) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *AssociateObjectToConfigProfilesRequest) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *AssociateObjectToConfigProfilesRequest) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


