# CopyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource that was requested to be copied. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JobId** | Pointer to **string** | An Unique Id to identify copy operation. | [optional] 

## Methods

### NewCopyResponse

`func NewCopyResponse() *CopyResponse`

NewCopyResponse instantiates a new CopyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCopyResponseWithDefaults

`func NewCopyResponseWithDefaults() *CopyResponse`

NewCopyResponseWithDefaults instantiates a new CopyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CopyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CopyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CopyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CopyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *CopyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CopyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CopyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CopyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *CopyResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CopyResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CopyResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CopyResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


