# IpamsvcCreateNextAvailableIPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]Address**](Address.md) | The list of allocated IP address objects. | [optional] 

## Methods

### NewIpamsvcCreateNextAvailableIPResponse

`func NewIpamsvcCreateNextAvailableIPResponse() *IpamsvcCreateNextAvailableIPResponse`

NewIpamsvcCreateNextAvailableIPResponse instantiates a new IpamsvcCreateNextAvailableIPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcCreateNextAvailableIPResponseWithDefaults

`func NewIpamsvcCreateNextAvailableIPResponseWithDefaults() *IpamsvcCreateNextAvailableIPResponse`

NewIpamsvcCreateNextAvailableIPResponseWithDefaults instantiates a new IpamsvcCreateNextAvailableIPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IpamsvcCreateNextAvailableIPResponse) GetResults() []Address`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IpamsvcCreateNextAvailableIPResponse) GetResultsOk() (*[]Address, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IpamsvcCreateNextAvailableIPResponse) SetResults(v []Address)`

SetResults sets Results field to given value.

### HasResults

`func (o *IpamsvcCreateNextAvailableIPResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


