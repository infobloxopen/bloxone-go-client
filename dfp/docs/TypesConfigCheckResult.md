# TypesConfigCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInfo** | Pointer to **string** | Provides more info about the potential problem. | [optional] 
**ConfigCheckType** | Pointer to **string** | Service check type. | [optional] 
**ResourceUri** | Pointer to **string** | URI of the resource that was checked. | [optional] 
**ResultCode** | Pointer to **string** | The check result. | [optional] 

## Methods

### NewTypesConfigCheckResult

`func NewTypesConfigCheckResult() *TypesConfigCheckResult`

NewTypesConfigCheckResult instantiates a new TypesConfigCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTypesConfigCheckResultWithDefaults

`func NewTypesConfigCheckResultWithDefaults() *TypesConfigCheckResult`

NewTypesConfigCheckResultWithDefaults instantiates a new TypesConfigCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInfo

`func (o *TypesConfigCheckResult) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *TypesConfigCheckResult) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *TypesConfigCheckResult) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *TypesConfigCheckResult) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetConfigCheckType

`func (o *TypesConfigCheckResult) GetConfigCheckType() string`

GetConfigCheckType returns the ConfigCheckType field if non-nil, zero value otherwise.

### GetConfigCheckTypeOk

`func (o *TypesConfigCheckResult) GetConfigCheckTypeOk() (*string, bool)`

GetConfigCheckTypeOk returns a tuple with the ConfigCheckType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigCheckType

`func (o *TypesConfigCheckResult) SetConfigCheckType(v string)`

SetConfigCheckType sets ConfigCheckType field to given value.

### HasConfigCheckType

`func (o *TypesConfigCheckResult) HasConfigCheckType() bool`

HasConfigCheckType returns a boolean if a field has been set.

### GetResourceUri

`func (o *TypesConfigCheckResult) GetResourceUri() string`

GetResourceUri returns the ResourceUri field if non-nil, zero value otherwise.

### GetResourceUriOk

`func (o *TypesConfigCheckResult) GetResourceUriOk() (*string, bool)`

GetResourceUriOk returns a tuple with the ResourceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUri

`func (o *TypesConfigCheckResult) SetResourceUri(v string)`

SetResourceUri sets ResourceUri field to given value.

### HasResourceUri

`func (o *TypesConfigCheckResult) HasResourceUri() bool`

HasResourceUri returns a boolean if a field has been set.

### GetResultCode

`func (o *TypesConfigCheckResult) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TypesConfigCheckResult) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TypesConfigCheckResult) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TypesConfigCheckResult) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


