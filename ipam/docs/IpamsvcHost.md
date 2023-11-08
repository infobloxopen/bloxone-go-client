# IpamsvcHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The primary IP address of the on-prem host. | [optional] [readonly] 
**AnycastAddresses** | Pointer to **[]string** | Anycast address configured to the host. Order is not significant. | [optional] [readonly] 
**AssociatedServer** | Pointer to [**IpamsvcHostAssociatedServer**](IpamsvcHostAssociatedServer.md) |  | [optional] 
**Comment** | Pointer to **string** | The description for the on-prem host. | [optional] [readonly] 
**CurrentVersion** | Pointer to **string** | Current dhcp application version of the host. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IpSpace** | Pointer to **string** | The resource identifier. | [optional] 
**Name** | Pointer to **string** | The display name of the on-prem host. | [optional] [readonly] 
**Ophid** | Pointer to **string** | The on-prem host ID. | [optional] [readonly] 
**ProviderId** | Pointer to **string** | External provider identifier. | [optional] [readonly] 
**Server** | Pointer to **string** | The resource identifier. | [optional] 
**Tags** | Pointer to **map[string]interface{}** | The tags of the on-prem host in JSON format. | [optional] 
**Type** | Pointer to **string** | Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services.  * _microsoft_active_directory_: host type is Microsoft Active Directory. | [optional] [readonly] 

## Methods

### NewIpamsvcHost

`func NewIpamsvcHost() *IpamsvcHost`

NewIpamsvcHost instantiates a new IpamsvcHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamsvcHostWithDefaults

`func NewIpamsvcHostWithDefaults() *IpamsvcHost`

NewIpamsvcHostWithDefaults instantiates a new IpamsvcHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *IpamsvcHost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *IpamsvcHost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *IpamsvcHost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *IpamsvcHost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAnycastAddresses

`func (o *IpamsvcHost) GetAnycastAddresses() []string`

GetAnycastAddresses returns the AnycastAddresses field if non-nil, zero value otherwise.

### GetAnycastAddressesOk

`func (o *IpamsvcHost) GetAnycastAddressesOk() (*[]string, bool)`

GetAnycastAddressesOk returns a tuple with the AnycastAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastAddresses

`func (o *IpamsvcHost) SetAnycastAddresses(v []string)`

SetAnycastAddresses sets AnycastAddresses field to given value.

### HasAnycastAddresses

`func (o *IpamsvcHost) HasAnycastAddresses() bool`

HasAnycastAddresses returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *IpamsvcHost) GetAssociatedServer() IpamsvcHostAssociatedServer`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *IpamsvcHost) GetAssociatedServerOk() (*IpamsvcHostAssociatedServer, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *IpamsvcHost) SetAssociatedServer(v IpamsvcHostAssociatedServer)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *IpamsvcHost) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetComment

`func (o *IpamsvcHost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *IpamsvcHost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *IpamsvcHost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *IpamsvcHost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *IpamsvcHost) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *IpamsvcHost) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *IpamsvcHost) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *IpamsvcHost) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetId

`func (o *IpamsvcHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpamsvcHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpamsvcHost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpamsvcHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpSpace

`func (o *IpamsvcHost) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *IpamsvcHost) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *IpamsvcHost) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *IpamsvcHost) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetName

`func (o *IpamsvcHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IpamsvcHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IpamsvcHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IpamsvcHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOphid

`func (o *IpamsvcHost) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *IpamsvcHost) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *IpamsvcHost) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *IpamsvcHost) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetProviderId

`func (o *IpamsvcHost) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *IpamsvcHost) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *IpamsvcHost) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *IpamsvcHost) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetServer

`func (o *IpamsvcHost) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *IpamsvcHost) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *IpamsvcHost) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *IpamsvcHost) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetTags

`func (o *IpamsvcHost) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *IpamsvcHost) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *IpamsvcHost) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *IpamsvcHost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *IpamsvcHost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpamsvcHost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpamsvcHost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IpamsvcHost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


