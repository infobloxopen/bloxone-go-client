# AtcdfpNetAddrPolicyAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddrNet** | Pointer to **string** | network address in IPv4 CIDR (address/bitmask length) string format | [optional] 
**PolicyId** | Pointer to **int32** | Identifier of the security policy associated with this address block | [optional] 

## Methods

### NewAtcdfpNetAddrPolicyAssignment

`func NewAtcdfpNetAddrPolicyAssignment() *AtcdfpNetAddrPolicyAssignment`

NewAtcdfpNetAddrPolicyAssignment instantiates a new AtcdfpNetAddrPolicyAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAtcdfpNetAddrPolicyAssignmentWithDefaults

`func NewAtcdfpNetAddrPolicyAssignmentWithDefaults() *AtcdfpNetAddrPolicyAssignment`

NewAtcdfpNetAddrPolicyAssignmentWithDefaults instantiates a new AtcdfpNetAddrPolicyAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrNet

`func (o *AtcdfpNetAddrPolicyAssignment) GetAddrNet() string`

GetAddrNet returns the AddrNet field if non-nil, zero value otherwise.

### GetAddrNetOk

`func (o *AtcdfpNetAddrPolicyAssignment) GetAddrNetOk() (*string, bool)`

GetAddrNetOk returns a tuple with the AddrNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrNet

`func (o *AtcdfpNetAddrPolicyAssignment) SetAddrNet(v string)`

SetAddrNet sets AddrNet field to given value.

### HasAddrNet

`func (o *AtcdfpNetAddrPolicyAssignment) HasAddrNet() bool`

HasAddrNet returns a boolean if a field has been set.

### GetPolicyId

`func (o *AtcdfpNetAddrPolicyAssignment) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *AtcdfpNetAddrPolicyAssignment) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *AtcdfpNetAddrPolicyAssignment) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *AtcdfpNetAddrPolicyAssignment) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


