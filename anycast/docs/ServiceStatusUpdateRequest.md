# ServiceStatusUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to [**ServiceStatusCode**](ServiceStatusCode.md) |  | [optional] [default to SERVICESTATUSCODE_SUCCESS]
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceStatusUpdateRequest

`func NewServiceStatusUpdateRequest() *ServiceStatusUpdateRequest`

NewServiceStatusUpdateRequest instantiates a new ServiceStatusUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusUpdateRequestWithDefaults

`func NewServiceStatusUpdateRequestWithDefaults() *ServiceStatusUpdateRequest`

NewServiceStatusUpdateRequestWithDefaults instantiates a new ServiceStatusUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *ServiceStatusUpdateRequest) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ServiceStatusUpdateRequest) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ServiceStatusUpdateRequest) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *ServiceStatusUpdateRequest) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetMessage

`func (o *ServiceStatusUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceStatusUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceStatusUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceStatusUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOphid

`func (o *ServiceStatusUpdateRequest) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ServiceStatusUpdateRequest) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ServiceStatusUpdateRequest) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ServiceStatusUpdateRequest) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceStatusUpdateRequest) GetStatusCode() ServiceStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceStatusUpdateRequest) GetStatusCodeOk() (*ServiceStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceStatusUpdateRequest) SetStatusCode(v ServiceStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceStatusUpdateRequest) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetVersion

`func (o *ServiceStatusUpdateRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServiceStatusUpdateRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServiceStatusUpdateRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServiceStatusUpdateRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


