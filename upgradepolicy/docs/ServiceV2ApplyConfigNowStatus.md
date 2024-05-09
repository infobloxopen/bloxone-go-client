# ServiceV2ApplyConfigNowStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostid** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to [**ServiceV2StatusCode**](ServiceV2StatusCode.md) |  | [optional] [default to SERVICEV2STATUSCODE_SUCCESS]

## Methods

### NewServiceV2ApplyConfigNowStatus

`func NewServiceV2ApplyConfigNowStatus() *ServiceV2ApplyConfigNowStatus`

NewServiceV2ApplyConfigNowStatus instantiates a new ServiceV2ApplyConfigNowStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceV2ApplyConfigNowStatusWithDefaults

`func NewServiceV2ApplyConfigNowStatusWithDefaults() *ServiceV2ApplyConfigNowStatus`

NewServiceV2ApplyConfigNowStatusWithDefaults instantiates a new ServiceV2ApplyConfigNowStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostid

`func (o *ServiceV2ApplyConfigNowStatus) GetHostid() string`

GetHostid returns the Hostid field if non-nil, zero value otherwise.

### GetHostidOk

`func (o *ServiceV2ApplyConfigNowStatus) GetHostidOk() (*string, bool)`

GetHostidOk returns a tuple with the Hostid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostid

`func (o *ServiceV2ApplyConfigNowStatus) SetHostid(v string)`

SetHostid sets Hostid field to given value.

### HasHostid

`func (o *ServiceV2ApplyConfigNowStatus) HasHostid() bool`

HasHostid returns a boolean if a field has been set.

### GetOphid

`func (o *ServiceV2ApplyConfigNowStatus) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ServiceV2ApplyConfigNowStatus) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ServiceV2ApplyConfigNowStatus) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ServiceV2ApplyConfigNowStatus) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetStatusCode

`func (o *ServiceV2ApplyConfigNowStatus) GetStatusCode() ServiceV2StatusCode`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ServiceV2ApplyConfigNowStatus) GetStatusCodeOk() (*ServiceV2StatusCode, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ServiceV2ApplyConfigNowStatus) SetStatusCode(v ServiceV2StatusCode)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ServiceV2ApplyConfigNowStatus) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


