# HealthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to [**HealthConfigStatusCode**](HealthConfigStatusCode.md) |  | [optional] [default to HEALTHCONFIGSTATUSCODE_SUCCESS]
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthConfig

`func NewHealthConfig() *HealthConfig`

NewHealthConfig instantiates a new HealthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthConfigWithDefaults

`func NewHealthConfigWithDefaults() *HealthConfig`

NewHealthConfigWithDefaults instantiates a new HealthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *HealthConfig) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *HealthConfig) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *HealthConfig) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *HealthConfig) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetMessage

`func (o *HealthConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOphid

`func (o *HealthConfig) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *HealthConfig) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *HealthConfig) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *HealthConfig) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetStatusCode

`func (o *HealthConfig) GetStatusCode() HealthConfigStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *HealthConfig) GetStatusCodeOk() (*HealthConfigStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *HealthConfig) SetStatusCode(v HealthConfigStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *HealthConfig) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetVersion

`func (o *HealthConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *HealthConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *HealthConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *HealthConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


