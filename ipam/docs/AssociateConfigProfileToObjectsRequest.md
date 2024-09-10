# AssociateConfigProfileToObjectsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigProfileId** | **string** | The resource identifier. | 
**Fields** | Pointer to [**ProtobufFieldMask**](ProtobufFieldMask.md) |  | [optional] 
**ObjectIds** | **[]string** | The resource identifier. | 

## Methods

### NewAssociateConfigProfileToObjectsRequest

`func NewAssociateConfigProfileToObjectsRequest(configProfileId string, objectIds []string, ) *AssociateConfigProfileToObjectsRequest`

NewAssociateConfigProfileToObjectsRequest instantiates a new AssociateConfigProfileToObjectsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociateConfigProfileToObjectsRequestWithDefaults

`func NewAssociateConfigProfileToObjectsRequestWithDefaults() *AssociateConfigProfileToObjectsRequest`

NewAssociateConfigProfileToObjectsRequestWithDefaults instantiates a new AssociateConfigProfileToObjectsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigProfileId

`func (o *AssociateConfigProfileToObjectsRequest) GetConfigProfileId() string`

GetConfigProfileId returns the ConfigProfileId field if non-nil, zero value otherwise.

### GetConfigProfileIdOk

`func (o *AssociateConfigProfileToObjectsRequest) GetConfigProfileIdOk() (*string, bool)`

GetConfigProfileIdOk returns a tuple with the ConfigProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigProfileId

`func (o *AssociateConfigProfileToObjectsRequest) SetConfigProfileId(v string)`

SetConfigProfileId sets ConfigProfileId field to given value.


### GetFields

`func (o *AssociateConfigProfileToObjectsRequest) GetFields() ProtobufFieldMask`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *AssociateConfigProfileToObjectsRequest) GetFieldsOk() (*ProtobufFieldMask, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *AssociateConfigProfileToObjectsRequest) SetFields(v ProtobufFieldMask)`

SetFields sets Fields field to given value.

### HasFields

`func (o *AssociateConfigProfileToObjectsRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetObjectIds

`func (o *AssociateConfigProfileToObjectsRequest) GetObjectIds() []string`

GetObjectIds returns the ObjectIds field if non-nil, zero value otherwise.

### GetObjectIdsOk

`func (o *AssociateConfigProfileToObjectsRequest) GetObjectIdsOk() (*[]string, bool)`

GetObjectIdsOk returns a tuple with the ObjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectIds

`func (o *AssociateConfigProfileToObjectsRequest) SetObjectIds(v []string)`

SetObjectIds sets ObjectIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


