# AccessCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Auto generated unique Bypass Code value | [optional] 
**Activation** | Pointer to **time.Time** | The time when the Bypass Code object was activated. | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time when the Bypass Code object was created. | [optional] [readonly] 
**Description** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **time.Time** | The time when the Bypass Code object was expired. | [optional] 
**Name** | Pointer to **string** | The name of Bypass Code | [optional] 
**PolicyIds** | Pointer to **[]int32** | The list of SecurityPolicy object identifiers. | [optional] 
**Rules** | Pointer to [**[]AccessCodeRule**](AccessCodeRule.md) | The list of selected security rules | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when the Bypass Code object was last updated. | [optional] [readonly] 

## Methods

### NewAccessCode

`func NewAccessCode() *AccessCode`

NewAccessCode instantiates a new AccessCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessCodeWithDefaults

`func NewAccessCodeWithDefaults() *AccessCode`

NewAccessCodeWithDefaults instantiates a new AccessCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AccessCode) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AccessCode) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AccessCode) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AccessCode) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetActivation

`func (o *AccessCode) GetActivation() time.Time`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *AccessCode) GetActivationOk() (*time.Time, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *AccessCode) SetActivation(v time.Time)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *AccessCode) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetCreatedTime

`func (o *AccessCode) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *AccessCode) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *AccessCode) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *AccessCode) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDescription

`func (o *AccessCode) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccessCode) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccessCode) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccessCode) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *AccessCode) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *AccessCode) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *AccessCode) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *AccessCode) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetName

`func (o *AccessCode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccessCode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccessCode) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccessCode) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPolicyIds

`func (o *AccessCode) GetPolicyIds() []int32`

GetPolicyIds returns the PolicyIds field if non-nil, zero value otherwise.

### GetPolicyIdsOk

`func (o *AccessCode) GetPolicyIdsOk() (*[]int32, bool)`

GetPolicyIdsOk returns a tuple with the PolicyIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyIds

`func (o *AccessCode) SetPolicyIds(v []int32)`

SetPolicyIds sets PolicyIds field to given value.

### HasPolicyIds

`func (o *AccessCode) HasPolicyIds() bool`

HasPolicyIds returns a boolean if a field has been set.

### GetRules

`func (o *AccessCode) GetRules() []AccessCodeRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *AccessCode) GetRulesOk() (*[]AccessCodeRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *AccessCode) SetRules(v []AccessCodeRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *AccessCode) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *AccessCode) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *AccessCode) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *AccessCode) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *AccessCode) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


