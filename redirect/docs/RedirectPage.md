# RedirectPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The content of the redirect page for the \&quot;custom\&quot; redirect type. | [optional] 
**CreatedTime** | Pointer to **time.Time** | The time when this Redirect Page object was created. | [optional] [readonly] 
**RedirectIpAddress** | Pointer to **string** | The redirect IPv4 address. | [optional] 
**RedirectIpv6Address** | Pointer to **string** | The redirect IPv6 address. | [optional] 
**Smart** | Pointer to **bool** | Whether the redirect type is smart | [optional] 
**Type** | Pointer to **string** | The type of the redirect page that can be \&quot;default\&quot; or \&quot;custom\&quot;. | [optional] 
**UpdatedTime** | Pointer to **time.Time** | The time when this Redirect Page object was last updated. | [optional] [readonly] 

## Methods

### NewRedirectPage

`func NewRedirectPage() *RedirectPage`

NewRedirectPage instantiates a new RedirectPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedirectPageWithDefaults

`func NewRedirectPageWithDefaults() *RedirectPage`

NewRedirectPageWithDefaults instantiates a new RedirectPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *RedirectPage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *RedirectPage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *RedirectPage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *RedirectPage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedTime

`func (o *RedirectPage) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *RedirectPage) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *RedirectPage) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *RedirectPage) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetRedirectIpAddress

`func (o *RedirectPage) GetRedirectIpAddress() string`

GetRedirectIpAddress returns the RedirectIpAddress field if non-nil, zero value otherwise.

### GetRedirectIpAddressOk

`func (o *RedirectPage) GetRedirectIpAddressOk() (*string, bool)`

GetRedirectIpAddressOk returns a tuple with the RedirectIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectIpAddress

`func (o *RedirectPage) SetRedirectIpAddress(v string)`

SetRedirectIpAddress sets RedirectIpAddress field to given value.

### HasRedirectIpAddress

`func (o *RedirectPage) HasRedirectIpAddress() bool`

HasRedirectIpAddress returns a boolean if a field has been set.

### GetRedirectIpv6Address

`func (o *RedirectPage) GetRedirectIpv6Address() string`

GetRedirectIpv6Address returns the RedirectIpv6Address field if non-nil, zero value otherwise.

### GetRedirectIpv6AddressOk

`func (o *RedirectPage) GetRedirectIpv6AddressOk() (*string, bool)`

GetRedirectIpv6AddressOk returns a tuple with the RedirectIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectIpv6Address

`func (o *RedirectPage) SetRedirectIpv6Address(v string)`

SetRedirectIpv6Address sets RedirectIpv6Address field to given value.

### HasRedirectIpv6Address

`func (o *RedirectPage) HasRedirectIpv6Address() bool`

HasRedirectIpv6Address returns a boolean if a field has been set.

### GetSmart

`func (o *RedirectPage) GetSmart() bool`

GetSmart returns the Smart field if non-nil, zero value otherwise.

### GetSmartOk

`func (o *RedirectPage) GetSmartOk() (*bool, bool)`

GetSmartOk returns a tuple with the Smart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmart

`func (o *RedirectPage) SetSmart(v bool)`

SetSmart sets Smart field to given value.

### HasSmart

`func (o *RedirectPage) HasSmart() bool`

HasSmart returns a boolean if a field has been set.

### GetType

`func (o *RedirectPage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RedirectPage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RedirectPage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RedirectPage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *RedirectPage) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *RedirectPage) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *RedirectPage) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *RedirectPage) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


