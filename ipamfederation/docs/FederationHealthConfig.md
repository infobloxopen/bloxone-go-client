# FederationHealthConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to [**HealthConfigStatusCode**](HealthConfigStatusCode.md) |  | [optional] [default to HEALTHCONFIGSTATUSCODE_SUCCESS]
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewFederationHealthConfig

`func NewFederationHealthConfig() *FederationHealthConfig`

NewFederationHealthConfig instantiates a new FederationHealthConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationHealthConfigWithDefaults

`func NewFederationHealthConfigWithDefaults() *FederationHealthConfig`

NewFederationHealthConfigWithDefaults instantiates a new FederationHealthConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppName

`func (o *FederationHealthConfig) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *FederationHealthConfig) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *FederationHealthConfig) SetAppName(v string)`

SetAppName sets AppName field to given value.

### HasAppName

`func (o *FederationHealthConfig) HasAppName() bool`

HasAppName returns a boolean if a field has been set.

### GetMessage

`func (o *FederationHealthConfig) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FederationHealthConfig) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FederationHealthConfig) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FederationHealthConfig) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOphid

`func (o *FederationHealthConfig) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *FederationHealthConfig) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *FederationHealthConfig) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *FederationHealthConfig) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetStatusCode

`func (o *FederationHealthConfig) GetStatusCode() HealthConfigStatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *FederationHealthConfig) GetStatusCodeOk() (*HealthConfigStatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *FederationHealthConfig) SetStatusCode(v HealthConfigStatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *FederationHealthConfig) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetVersion

`func (o *FederationHealthConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FederationHealthConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FederationHealthConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FederationHealthConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


