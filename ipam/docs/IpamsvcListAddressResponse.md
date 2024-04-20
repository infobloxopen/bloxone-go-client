# IpamsvcListAddressResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Address**](Address.md) | The list of Address objects. | [optional] 

## Methods

### NewIpamsvcListAddressResponse

`func NewIpamsvcListAddressResponse() *IpamsvcListAddressResponse`

NewIpamsvcListAddressResponse instantiates a new IpamsvcListAddressResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcListAddressResponseWithDefaults

`func NewIpamsvcListAddressResponseWithDefaults() *IpamsvcListAddressResponse`

NewIpamsvcListAddressResponseWithDefaults instantiates a new IpamsvcListAddressResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IpamsvcListAddressResponse) GetResults() []Address`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IpamsvcListAddressResponse) GetResultsOk() (*[]Address, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IpamsvcListAddressResponse) SetResults(v []Address)`

SetResults sets Results field to given value.

### HasResults

`func (o *IpamsvcListAddressResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


