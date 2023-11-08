# IpamsvcCopyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the resource that was requested to be copied. | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JobId** | Pointer to **string** | An Unique Id to identify copy operation. | [optional] 

## Methods

### NewIpamsvcCopyResponse

`func NewIpamsvcCopyResponse() *IpamsvcCopyResponse`

NewIpamsvcCopyResponse instantiates a new IpamsvcCopyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcCopyResponseWithDefaults

`func NewIpamsvcCopyResponseWithDefaults() *IpamsvcCopyResponse`

NewIpamsvcCopyResponseWithDefaults instantiates a new IpamsvcCopyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IpamsvcCopyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IpamsvcCopyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IpamsvcCopyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IpamsvcCopyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcCopyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcCopyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcCopyResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcCopyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJobId

`func (o *IpamsvcCopyResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *IpamsvcCopyResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *IpamsvcCopyResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *IpamsvcCopyResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


