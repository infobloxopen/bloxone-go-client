# ConfigACLItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | **string** | Access permission for _element_.  Allowed values:  * _allow_,  * _deny_. | 
**Acl** | Pointer to **string** | The resource identifier. | [optional] 
**Address** | Pointer to **string** | Optional. Data for _ip_ _element_.  Must be empty if _element_ is not _ip_. | [optional] 
**Element** | **string** | Type of element.  Allowed values:  * _any_,  * _ip_,  * _acl_,  * _tsig_key_. | 
**TsigKey** | Pointer to [**ConfigTSIGKey**](ConfigTSIGKey.md) |  | [optional] 

## Methods

### NewConfigACLItem

`func NewConfigACLItem(access string, element string, ) *ConfigACLItem`

NewConfigACLItem instantiates a new ConfigACLItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigACLItemWithDefaults

`func NewConfigACLItemWithDefaults() *ConfigACLItem`

NewConfigACLItemWithDefaults instantiates a new ConfigACLItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ConfigACLItem) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ConfigACLItem) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ConfigACLItem) SetAccess(v string)`

SetAccess sets Access field to given value.


### GetAcl

`func (o *ConfigACLItem) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *ConfigACLItem) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *ConfigACLItem) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *ConfigACLItem) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetAddress

`func (o *ConfigACLItem) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConfigACLItem) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConfigACLItem) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConfigACLItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetElement

`func (o *ConfigACLItem) GetElement() string`

GetElement returns the Element field if non-nil, zero value otherwise.

### GetElementOk

`func (o *ConfigACLItem) GetElementOk() (*string, bool)`

GetElementOk returns a tuple with the Element field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElement

`func (o *ConfigACLItem) SetElement(v string)`

SetElement sets Element field to given value.


### GetTsigKey

`func (o *ConfigACLItem) GetTsigKey() ConfigTSIGKey`

GetTsigKey returns the TsigKey field if non-nil, zero value otherwise.

### GetTsigKeyOk

`func (o *ConfigACLItem) GetTsigKeyOk() (*ConfigTSIGKey, bool)`

GetTsigKeyOk returns a tuple with the TsigKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTsigKey

`func (o *ConfigACLItem) SetTsigKey(v ConfigTSIGKey)`

SetTsigKey sets TsigKey field to given value.

### HasTsigKey

`func (o *ConfigACLItem) HasTsigKey() bool`

HasTsigKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


