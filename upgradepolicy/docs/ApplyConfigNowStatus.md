# ApplyConfigNowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostid** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to [**StatusCode**](StatusCode.md) |  | [optional] [default to STATUSCODE_SUCCESS]

## Methods

### NewApplyConfigNowStatus

`func NewApplyConfigNowStatus() *ApplyConfigNowStatus`

NewApplyConfigNowStatus instantiates a new ApplyConfigNowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyConfigNowStatusWithDefaults

`func NewApplyConfigNowStatusWithDefaults() *ApplyConfigNowStatus`

NewApplyConfigNowStatusWithDefaults instantiates a new ApplyConfigNowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostid

`func (o *ApplyConfigNowStatus) GetHostid() string`

GetHostid returns the Hostid field if non-nil, zero value otherwise.

### GetHostidOk

`func (o *ApplyConfigNowStatus) GetHostidOk() (*string, bool)`

GetHostidOk returns a tuple with the Hostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostid

`func (o *ApplyConfigNowStatus) SetHostid(v string)`

SetHostid sets Hostid field to given value.

### HasHostid

`func (o *ApplyConfigNowStatus) HasHostid() bool`

HasHostid returns a boolean if a field has been set.

### GetOphid

`func (o *ApplyConfigNowStatus) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ApplyConfigNowStatus) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ApplyConfigNowStatus) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ApplyConfigNowStatus) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetStatusCode

`func (o *ApplyConfigNowStatus) GetStatusCode() StatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ApplyConfigNowStatus) GetStatusCodeOk() (*StatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ApplyConfigNowStatus) SetStatusCode(v StatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ApplyConfigNowStatus) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


