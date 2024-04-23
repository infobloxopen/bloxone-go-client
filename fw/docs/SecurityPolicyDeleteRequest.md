# SecurityPolicyDeleteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int32** | The list of Security Policy object identifiers. | [optional] 

## Methods

### NewSecurityPolicyDeleteRequest

`func NewSecurityPolicyDeleteRequest() *SecurityPolicyDeleteRequest`

NewSecurityPolicyDeleteRequest instantiates a new SecurityPolicyDeleteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityPolicyDeleteRequestWithDefaults

`func NewSecurityPolicyDeleteRequestWithDefaults() *SecurityPolicyDeleteRequest`

NewSecurityPolicyDeleteRequestWithDefaults instantiates a new SecurityPolicyDeleteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *SecurityPolicyDeleteRequest) GetIds() []int32`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SecurityPolicyDeleteRequest) GetIdsOk() (*[]int32, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SecurityPolicyDeleteRequest) SetIds(v []int32)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SecurityPolicyDeleteRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


