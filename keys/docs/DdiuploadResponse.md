# DdiuploadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KerberosKeys** | Pointer to [**KerberosKeys**](KerberosKeys.md) |  | [optional] 
**Warning** | Pointer to **string** | May contain any non-critical warning messages after processing the content. | [optional] 

## Methods

### NewDdiuploadResponse

`func NewDdiuploadResponse() *DdiuploadResponse`

NewDdiuploadResponse instantiates a new DdiuploadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdiuploadResponseWithDefaults

`func NewDdiuploadResponseWithDefaults() *DdiuploadResponse`

NewDdiuploadResponseWithDefaults instantiates a new DdiuploadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKerberosKeys

`func (o *DdiuploadResponse) GetKerberosKeys() KerberosKeys`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *DdiuploadResponse) GetKerberosKeysOk() (*KerberosKeys, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *DdiuploadResponse) SetKerberosKeys(v KerberosKeys)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *DdiuploadResponse) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetWarning

`func (o *DdiuploadResponse) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *DdiuploadResponse) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *DdiuploadResponse) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *DdiuploadResponse) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


