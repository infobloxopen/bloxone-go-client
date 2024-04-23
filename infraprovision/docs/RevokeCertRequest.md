# RevokeCertRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertSerial** | Pointer to **string** |  | [optional] 
**Ophid** | Pointer to **string** | On-prem host ID which can be obtained either from on-prem or BloxOne UI portal(Manage &gt; Infrastructure &gt; Hosts &gt; Select the onprem &gt; click on 3 dots on top right side &gt; General Information &gt; Ophid) . | [optional] 
**RevokeReason** | Pointer to **string** |  | [optional] 

## Methods

### NewRevokeCertRequest

`func NewRevokeCertRequest() *RevokeCertRequest`

NewRevokeCertRequest instantiates a new RevokeCertRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRevokeCertRequestWithDefaults

`func NewRevokeCertRequestWithDefaults() *RevokeCertRequest`

NewRevokeCertRequestWithDefaults instantiates a new RevokeCertRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertSerial

`func (o *RevokeCertRequest) GetCertSerial() string`

GetCertSerial returns the CertSerial field if non-nil, zero value otherwise.

### GetCertSerialOk

`func (o *RevokeCertRequest) GetCertSerialOk() (*string, bool)`

GetCertSerialOk returns a tuple with the CertSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertSerial

`func (o *RevokeCertRequest) SetCertSerial(v string)`

SetCertSerial sets CertSerial field to given value.

### HasCertSerial

`func (o *RevokeCertRequest) HasCertSerial() bool`

HasCertSerial returns a boolean if a field has been set.

### GetOphid

`func (o *RevokeCertRequest) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *RevokeCertRequest) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *RevokeCertRequest) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *RevokeCertRequest) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetRevokeReason

`func (o *RevokeCertRequest) GetRevokeReason() string`

GetRevokeReason returns the RevokeReason field if non-nil, zero value otherwise.

### GetRevokeReasonOk

`func (o *RevokeCertRequest) GetRevokeReasonOk() (*string, bool)`

GetRevokeReasonOk returns a tuple with the RevokeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeReason

`func (o *RevokeCertRequest) SetRevokeReason(v string)`

SetRevokeReason sets RevokeReason field to given value.

### HasRevokeReason

`func (o *RevokeCertRequest) HasRevokeReason() bool`

HasRevokeReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


