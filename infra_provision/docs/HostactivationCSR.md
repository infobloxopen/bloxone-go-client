# HostactivationCSR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to [**TypesInetValue**](TypesInetValue.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Csr** | Pointer to **string** |  | [optional] 
**HostSerial** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JoinToken** | Pointer to [**HostactivationJoinToken**](HostactivationJoinToken.md) |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SrcIp** | Pointer to [**TypesInetValue**](TypesInetValue.md) |  | [optional] 
**State** | Pointer to [**HostactivationCSRState**](HostactivationCSRState.md) |  | [optional] [default to HOSTACTIVATIONCSRSTATE_UNKNOWN]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewHostactivationCSR

`func NewHostactivationCSR() *HostactivationCSR`

NewHostactivationCSR instantiates a new HostactivationCSR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostactivationCSRWithDefaults

`func NewHostactivationCSRWithDefaults() *HostactivationCSR`

NewHostactivationCSRWithDefaults instantiates a new HostactivationCSR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *HostactivationCSR) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *HostactivationCSR) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *HostactivationCSR) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.

### HasActivationCode

`func (o *HostactivationCSR) HasActivationCode() bool`

HasActivationCode returns a boolean if a field has been set.

### GetClientIp

`func (o *HostactivationCSR) GetClientIp() TypesInetValue`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *HostactivationCSR) GetClientIpOk() (*TypesInetValue, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *HostactivationCSR) SetClientIp(v TypesInetValue)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *HostactivationCSR) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *HostactivationCSR) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *HostactivationCSR) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *HostactivationCSR) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *HostactivationCSR) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCsr

`func (o *HostactivationCSR) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *HostactivationCSR) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *HostactivationCSR) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *HostactivationCSR) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetHostSerial

`func (o *HostactivationCSR) GetHostSerial() string`

GetHostSerial returns the HostSerial field if non-nil, zero value otherwise.

### GetHostSerialOk

`func (o *HostactivationCSR) GetHostSerialOk() (*string, bool)`

GetHostSerialOk returns a tuple with the HostSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSerial

`func (o *HostactivationCSR) SetHostSerial(v string)`

SetHostSerial sets HostSerial field to given value.

### HasHostSerial

`func (o *HostactivationCSR) HasHostSerial() bool`

HasHostSerial returns a boolean if a field has been set.

### GetId

`func (o *HostactivationCSR) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostactivationCSR) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostactivationCSR) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostactivationCSR) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinToken

`func (o *HostactivationCSR) GetJoinToken() HostactivationJoinToken`

GetJoinToken returns the JoinToken field if non-nil, zero value otherwise.

### GetJoinTokenOk

`func (o *HostactivationCSR) GetJoinTokenOk() (*HostactivationJoinToken, bool)`

GetJoinTokenOk returns a tuple with the JoinToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinToken

`func (o *HostactivationCSR) SetJoinToken(v HostactivationJoinToken)`

SetJoinToken sets JoinToken field to given value.

### HasJoinToken

`func (o *HostactivationCSR) HasJoinToken() bool`

HasJoinToken returns a boolean if a field has been set.

### GetOphid

`func (o *HostactivationCSR) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *HostactivationCSR) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *HostactivationCSR) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *HostactivationCSR) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetSignature

`func (o *HostactivationCSR) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *HostactivationCSR) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *HostactivationCSR) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *HostactivationCSR) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSrcIp

`func (o *HostactivationCSR) GetSrcIp() TypesInetValue`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *HostactivationCSR) GetSrcIpOk() (*TypesInetValue, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *HostactivationCSR) SetSrcIp(v TypesInetValue)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *HostactivationCSR) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetState

`func (o *HostactivationCSR) GetState() HostactivationCSRState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *HostactivationCSR) GetStateOk() (*HostactivationCSRState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *HostactivationCSR) SetState(v HostactivationCSRState)`

SetState sets State field to given value.

### HasState

`func (o *HostactivationCSR) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HostactivationCSR) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HostactivationCSR) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HostactivationCSR) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HostactivationCSR) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


