# IpamsvcListFilterResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]IpamsvcFilter**](IpamsvcFilter.md) | The list of DHCP Filter objects. | [optional] 

## Methods

### NewIpamsvcListFilterResponse

`func NewIpamsvcListFilterResponse() *IpamsvcListFilterResponse`

NewIpamsvcListFilterResponse instantiates a new IpamsvcListFilterResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcListFilterResponseWithDefaults

`func NewIpamsvcListFilterResponseWithDefaults() *IpamsvcListFilterResponse`

NewIpamsvcListFilterResponseWithDefaults instantiates a new IpamsvcListFilterResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IpamsvcListFilterResponse) GetResults() []IpamsvcFilter`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IpamsvcListFilterResponse) GetResultsOk() (*[]IpamsvcFilter, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IpamsvcListFilterResponse) SetResults(v []IpamsvcFilter)`

SetResults sets Results field to given value.

### HasResults

`func (o *IpamsvcListFilterResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


