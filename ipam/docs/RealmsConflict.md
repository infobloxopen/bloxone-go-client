# RealmsConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**FederatedRealms** | Pointer to **[]string** | List of __FederatedRealm__ object ids. | [optional] 
**IpSpace** | Pointer to **string** | The resource identifier. | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewRealmsConflict

`func NewRealmsConflict() *RealmsConflict`

NewRealmsConflict instantiates a new RealmsConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRealmsConflictWithDefaults

`func NewRealmsConflictWithDefaults() *RealmsConflict`

NewRealmsConflictWithDefaults instantiates a new RealmsConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *RealmsConflict) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RealmsConflict) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RealmsConflict) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RealmsConflict) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFederatedRealms

`func (o *RealmsConflict) GetFederatedRealms() []string`

GetFederatedRealms returns the FederatedRealms field if non-nil, zero value otherwise.

### GetFederatedRealmsOk

`func (o *RealmsConflict) GetFederatedRealmsOk() (*[]string, bool)`

GetFederatedRealmsOk returns a tuple with the FederatedRealms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedRealms

`func (o *RealmsConflict) SetFederatedRealms(v []string)`

SetFederatedRealms sets FederatedRealms field to given value.

### HasFederatedRealms

`func (o *RealmsConflict) HasFederatedRealms() bool`

HasFederatedRealms returns a boolean if a field has been set.

### GetIpSpace

`func (o *RealmsConflict) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *RealmsConflict) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *RealmsConflict) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *RealmsConflict) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetType

`func (o *RealmsConflict) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RealmsConflict) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RealmsConflict) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RealmsConflict) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


