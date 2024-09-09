# ReadDHCPServiceInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to [**DHCPServiceInstance**](DHCPServiceInstance.md) | The DHCP Service object. | [optional] 

## Methods

### NewReadDHCPServiceInstanceResponse

`func NewReadDHCPServiceInstanceResponse() *ReadDHCPServiceInstanceResponse`

NewReadDHCPServiceInstanceResponse instantiates a new ReadDHCPServiceInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReadDHCPServiceInstanceResponseWithDefaults

`func NewReadDHCPServiceInstanceResponseWithDefaults() *ReadDHCPServiceInstanceResponse`

NewReadDHCPServiceInstanceResponseWithDefaults instantiates a new ReadDHCPServiceInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *ReadDHCPServiceInstanceResponse) GetResult() DHCPServiceInstance`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ReadDHCPServiceInstanceResponse) GetResultOk() (*DHCPServiceInstance, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ReadDHCPServiceInstanceResponse) SetResult(v DHCPServiceInstance)`

SetResult sets Result field to given value.

### HasResult

`func (o *ReadDHCPServiceInstanceResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


