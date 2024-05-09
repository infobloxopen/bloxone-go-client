# CustomRedirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | Pointer to **time.Time** | The time when this Custom Redirect object was created. | [optional] [readonly] 
**Data** | Pointer to **string** | The list of csv custom IPv4/IPv6 or a single domain redirect address. | [optional] 
**Id** | Pointer to **int32** | The Custom Redirect object identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the custom redirect. | [optional] 
**PolicyIds** | Pointer to **[]int32** | The list of the security policy identifiers with which the named list is associated. | [optional] [readonly] 
**PolicyNames** | Pointer to **[]string** | The list of the security policy names with which the custom redirect is associated. | [optional] [readonly] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Custom Redirect object was last updated. | [optional] [readonly] 

## Methods

### NewCustomRedirect

`func NewCustomRedirect() *CustomRedirect`

NewCustomRedirect instantiates a new CustomRedirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomRedirectWithDefaults

`func NewCustomRedirectWithDefaults() *CustomRedirect`

NewCustomRedirectWithDefaults instantiates a new CustomRedirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTime

`func (o *CustomRedirect) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CustomRedirect) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CustomRedirect) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CustomRedirect) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetData

`func (o *CustomRedirect) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomRedirect) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomRedirect) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CustomRedirect) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *CustomRedirect) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomRedirect) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomRedirect) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CustomRedirect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CustomRedirect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomRedirect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomRedirect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomRedirect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyIds

`func (o *CustomRedirect) GetPolicyIds() []int32`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *CustomRedirect) GetPolicyIdsOk() (*[]int32, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *CustomRedirect) SetPolicyIds(v []int32)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *CustomRedirect) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### GetPolicyNames

`func (o *CustomRedirect) GetPolicyNames() []string`

GetPolicyNames returns the PolicyNames field if non-nil, zero value otherwise.

### GetPolicyNamesOk

`func (o *CustomRedirect) GetPolicyNamesOk() (*[]string, bool)`

GetPolicyNamesOk returns a tuple with the PolicyNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyNames

`func (o *CustomRedirect) SetPolicyNames(v []string)`

SetPolicyNames sets PolicyNames field to given value.

### HasPolicyNames

`func (o *CustomRedirect) HasPolicyNames() bool`

HasPolicyNames returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *CustomRedirect) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *CustomRedirect) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *CustomRedirect) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *CustomRedirect) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


