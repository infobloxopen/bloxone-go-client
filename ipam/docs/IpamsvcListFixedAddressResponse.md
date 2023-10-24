# IpamsvcListFixedAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]IpamsvcFixedAddress**](IpamsvcFixedAddress.md) | The list of FixedAddress objects. | [optional] 

## Methods

### NewIpamsvcListFixedAddressResponse

`func NewIpamsvcListFixedAddressResponse() *IpamsvcListFixedAddressResponse`

NewIpamsvcListFixedAddressResponse instantiates a new IpamsvcListFixedAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcListFixedAddressResponseWithDefaults

`func NewIpamsvcListFixedAddressResponseWithDefaults() *IpamsvcListFixedAddressResponse`

NewIpamsvcListFixedAddressResponseWithDefaults instantiates a new IpamsvcListFixedAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IpamsvcListFixedAddressResponse) GetResults() []IpamsvcFixedAddress`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IpamsvcListFixedAddressResponse) GetResultsOk() (*[]IpamsvcFixedAddress, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IpamsvcListFixedAddressResponse) SetResults(v []IpamsvcFixedAddress)`

SetResults sets Results field to given value.

### HasResults

`func (o *IpamsvcListFixedAddressResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


