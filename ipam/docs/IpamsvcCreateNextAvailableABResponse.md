# IpamsvcCreateNextAvailableABResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Results** | Pointer to [**[]IpamsvcAddressBlock**](IpamsvcAddressBlock.md) | The list of Next Available Address Block objects. | [optional] 

## Methods

### NewIpamsvcCreateNextAvailableABResponse

`func NewIpamsvcCreateNextAvailableABResponse() *IpamsvcCreateNextAvailableABResponse`

NewIpamsvcCreateNextAvailableABResponse instantiates a new IpamsvcCreateNextAvailableABResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcCreateNextAvailableABResponseWithDefaults

`func NewIpamsvcCreateNextAvailableABResponseWithDefaults() *IpamsvcCreateNextAvailableABResponse`

NewIpamsvcCreateNextAvailableABResponseWithDefaults instantiates a new IpamsvcCreateNextAvailableABResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResults

`func (o *IpamsvcCreateNextAvailableABResponse) GetResults() []IpamsvcAddressBlock`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *IpamsvcCreateNextAvailableABResponse) GetResultsOk() (*[]IpamsvcAddressBlock, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *IpamsvcCreateNextAvailableABResponse) SetResults(v []IpamsvcAddressBlock)`

SetResults sets Results field to given value.

### HasResults

`func (o *IpamsvcCreateNextAvailableABResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


