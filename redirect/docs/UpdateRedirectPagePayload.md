# UpdateRedirectPagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | The content of the redirect page for the \&quot;custom\&quot; redirect type. | [optional] 
**RedirectIpAddress** | Pointer to **string** | The redirect IPv4 address. | [optional] 
**RedirectIpv6Address** | Pointer to **string** | The redirect IPv6 address. | [optional] 
**Smart** | Pointer to **bool** | Change the redirect page from non-proxy (smart &#x3D;&#x3D; false) to proxy (smart) | [optional] 
**Type** | Pointer to **string** | The type of the redirect page that can be \&quot;default\&quot; or \&quot;custom\&quot;. | [optional] 

## Methods

### NewUpdateRedirectPagePayload

`func NewUpdateRedirectPagePayload() *UpdateRedirectPagePayload`

NewUpdateRedirectPagePayload instantiates a new UpdateRedirectPagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRedirectPagePayloadWithDefaults

`func NewUpdateRedirectPagePayloadWithDefaults() *UpdateRedirectPagePayload`

NewUpdateRedirectPagePayloadWithDefaults instantiates a new UpdateRedirectPagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *UpdateRedirectPagePayload) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdateRedirectPagePayload) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdateRedirectPagePayload) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdateRedirectPagePayload) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRedirectIpAddress

`func (o *UpdateRedirectPagePayload) GetRedirectIpAddress() string`

GetRedirectIpAddress returns the RedirectIpAddress field if non-nil, zero value otherwise.

### GetRedirectIpAddressOk

`func (o *UpdateRedirectPagePayload) GetRedirectIpAddressOk() (*string, bool)`

GetRedirectIpAddressOk returns a tuple with the RedirectIpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectIpAddress

`func (o *UpdateRedirectPagePayload) SetRedirectIpAddress(v string)`

SetRedirectIpAddress sets RedirectIpAddress field to given value.

### HasRedirectIpAddress

`func (o *UpdateRedirectPagePayload) HasRedirectIpAddress() bool`

HasRedirectIpAddress returns a boolean if a field has been set.

### GetRedirectIpv6Address

`func (o *UpdateRedirectPagePayload) GetRedirectIpv6Address() string`

GetRedirectIpv6Address returns the RedirectIpv6Address field if non-nil, zero value otherwise.

### GetRedirectIpv6AddressOk

`func (o *UpdateRedirectPagePayload) GetRedirectIpv6AddressOk() (*string, bool)`

GetRedirectIpv6AddressOk returns a tuple with the RedirectIpv6Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectIpv6Address

`func (o *UpdateRedirectPagePayload) SetRedirectIpv6Address(v string)`

SetRedirectIpv6Address sets RedirectIpv6Address field to given value.

### HasRedirectIpv6Address

`func (o *UpdateRedirectPagePayload) HasRedirectIpv6Address() bool`

HasRedirectIpv6Address returns a boolean if a field has been set.

### GetSmart

`func (o *UpdateRedirectPagePayload) GetSmart() bool`

GetSmart returns the Smart field if non-nil, zero value otherwise.

### GetSmartOk

`func (o *UpdateRedirectPagePayload) GetSmartOk() (*bool, bool)`

GetSmartOk returns a tuple with the Smart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmart

`func (o *UpdateRedirectPagePayload) SetSmart(v bool)`

SetSmart sets Smart field to given value.

### HasSmart

`func (o *UpdateRedirectPagePayload) HasSmart() bool`

HasSmart returns a boolean if a field has been set.

### GetType

`func (o *UpdateRedirectPagePayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateRedirectPagePayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateRedirectPagePayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateRedirectPagePayload) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


