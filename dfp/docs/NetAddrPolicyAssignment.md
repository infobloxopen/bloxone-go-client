# NetAddrPolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrNet** | Pointer to **string** | network address in IPv4 CIDR (address/bitmask length) string format | [optional] 
**PolicyId** | Pointer to **int32** | Identifier of the security policy associated with this address block | [optional] 

## Methods

### NewNetAddrPolicyAssignment

`func NewNetAddrPolicyAssignment() *NetAddrPolicyAssignment`

NewNetAddrPolicyAssignment instantiates a new NetAddrPolicyAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetAddrPolicyAssignmentWithDefaults

`func NewNetAddrPolicyAssignmentWithDefaults() *NetAddrPolicyAssignment`

NewNetAddrPolicyAssignmentWithDefaults instantiates a new NetAddrPolicyAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrNet

`func (o *NetAddrPolicyAssignment) GetAddrNet() string`

GetAddrNet returns the AddrNet field if non-nil, zero value otherwise.

### GetAddrNetOk

`func (o *NetAddrPolicyAssignment) GetAddrNetOk() (*string, bool)`

GetAddrNetOk returns a tuple with the AddrNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrNet

`func (o *NetAddrPolicyAssignment) SetAddrNet(v string)`

SetAddrNet sets AddrNet field to given value.

### HasAddrNet

`func (o *NetAddrPolicyAssignment) HasAddrNet() bool`

HasAddrNet returns a boolean if a field has been set.

### GetPolicyId

`func (o *NetAddrPolicyAssignment) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *NetAddrPolicyAssignment) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *NetAddrPolicyAssignment) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *NetAddrPolicyAssignment) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


