# DTCPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | __DTC Policy__ display name. | [optional] [readonly] 
**PolicyId** | Pointer to **string** | The resource identifier. | [optional] 

## Methods

### NewDTCPolicy

`func NewDTCPolicy() *DTCPolicy`

NewDTCPolicy instantiates a new DTCPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDTCPolicyWithDefaults

`func NewDTCPolicyWithDefaults() *DTCPolicy`

NewDTCPolicyWithDefaults instantiates a new DTCPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DTCPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DTCPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DTCPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DTCPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyId

`func (o *DTCPolicy) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DTCPolicy) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DTCPolicy) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *DTCPolicy) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


