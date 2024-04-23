# CSR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationCode** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to [**TypesInetValue**](TypesInetValue.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Csr** | Pointer to **string** |  | [optional] 
**HostSerial** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**JoinToken** | Pointer to [**JoinToken**](JoinToken.md) |  | [optional] 
**Ophid** | Pointer to **string** |  | [optional] 
**Signature** | Pointer to **string** |  | [optional] 
**SrcIp** | Pointer to [**TypesInetValue**](TypesInetValue.md) |  | [optional] 
**State** | Pointer to [**CSRState**](CSRState.md) |  | [optional] [default to CSRSTATE_UNKNOWN]
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCSR

`func NewCSR() *CSR`

NewCSR instantiates a new CSR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSRWithDefaults

`func NewCSRWithDefaults() *CSR`

NewCSRWithDefaults instantiates a new CSR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationCode

`func (o *CSR) GetActivationCode() string`

GetActivationCode returns the ActivationCode field if non-nil, zero value otherwise.

### GetActivationCodeOk

`func (o *CSR) GetActivationCodeOk() (*string, bool)`

GetActivationCodeOk returns a tuple with the ActivationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationCode

`func (o *CSR) SetActivationCode(v string)`

SetActivationCode sets ActivationCode field to given value.

### HasActivationCode

`func (o *CSR) HasActivationCode() bool`

HasActivationCode returns a boolean if a field has been set.

### GetClientIp

`func (o *CSR) GetClientIp() TypesInetValue`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *CSR) GetClientIpOk() (*TypesInetValue, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *CSR) SetClientIp(v TypesInetValue)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *CSR) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CSR) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CSR) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CSR) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CSR) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCsr

`func (o *CSR) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *CSR) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *CSR) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *CSR) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### GetHostSerial

`func (o *CSR) GetHostSerial() string`

GetHostSerial returns the HostSerial field if non-nil, zero value otherwise.

### GetHostSerialOk

`func (o *CSR) GetHostSerialOk() (*string, bool)`

GetHostSerialOk returns a tuple with the HostSerial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostSerial

`func (o *CSR) SetHostSerial(v string)`

SetHostSerial sets HostSerial field to given value.

### HasHostSerial

`func (o *CSR) HasHostSerial() bool`

HasHostSerial returns a boolean if a field has been set.

### GetId

`func (o *CSR) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSR) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSR) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CSR) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinToken

`func (o *CSR) GetJoinToken() JoinToken`

GetJoinToken returns the JoinToken field if non-nil, zero value otherwise.

### GetJoinTokenOk

`func (o *CSR) GetJoinTokenOk() (*JoinToken, bool)`

GetJoinTokenOk returns a tuple with the JoinToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinToken

`func (o *CSR) SetJoinToken(v JoinToken)`

SetJoinToken sets JoinToken field to given value.

### HasJoinToken

`func (o *CSR) HasJoinToken() bool`

HasJoinToken returns a boolean if a field has been set.

### GetOphid

`func (o *CSR) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *CSR) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *CSR) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *CSR) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetSignature

`func (o *CSR) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *CSR) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *CSR) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *CSR) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSrcIp

`func (o *CSR) GetSrcIp() TypesInetValue`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *CSR) GetSrcIpOk() (*TypesInetValue, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *CSR) SetSrcIp(v TypesInetValue)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *CSR) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetState

`func (o *CSR) GetState() CSRState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CSR) GetStateOk() (*CSRState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CSR) SetState(v CSRState)`

SetState sets State field to given value.

### HasState

`func (o *CSR) HasState() bool`

HasState returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CSR) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CSR) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CSR) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CSR) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


