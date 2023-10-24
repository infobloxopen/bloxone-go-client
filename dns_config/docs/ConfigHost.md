# ConfigHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbsoluteName** | Pointer to **string** | Host FQDN. | [optional] [readonly] 
**Address** | Pointer to **string** | Host&#39;s primary IP Address. | [optional] [readonly] 
**AnycastAddresses** | Pointer to **[]string** | Anycast address configured to the host. Order is not significant. | [optional] [readonly] 
**AssociatedServer** | Pointer to [**ConfigHostAssociatedServer**](ConfigHostAssociatedServer.md) |  | [optional] 
**Comment** | Pointer to **string** | Host description. | [optional] [readonly] 
**CurrentVersion** | Pointer to **string** | Host current version. | [optional] [readonly] 
**Dfp** | Pointer to **bool** | Below _dfp_ field is deprecated and not supported anymore. The indication whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host will be migrated into the new _dfp_service_ field. | [optional] [readonly] 
**DfpService** | Pointer to **string** | DFP service indicates whether or not BloxOne DDI DNS and BloxOne TD DFP are both active on the host. If so, BloxOne DDI DNS will augment recursive queries and forward them to BloxOne TD DFP. Allowed values:  * _unavailable_: BloxOne TD DFP application is not available,  * _enabled_: BloxOne TD DFP application is available and enabled,  * _disabled_: BloxOne TD DFP application is available but disabled. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**InheritanceSources** | Pointer to [**ConfigHostInheritance**](ConfigHostInheritance.md) |  | [optional] 
**KerberosKeys** | Pointer to [**[]ConfigKerberosKey**](ConfigKerberosKey.md) | Optional. _kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty. | [optional] 
**Name** | Pointer to **string** | Host display name. | [optional] [readonly] 
**Ophid** | Pointer to **string** | On-Prem Host ID. | [optional] [readonly] 
**ProtocolAbsoluteName** | Pointer to **string** | Host FQDN in punycode. | [optional] [readonly] 
**ProviderId** | Pointer to **string** | External provider identifier. | [optional] [readonly] 
**Server** | Pointer to **string** | The resource identifier. | [optional] 
**SiteId** | Pointer to **string** | Host site ID. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | Host tagging specifics. | [optional] 
**Type** | Pointer to **string** | Defines the type of host. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform. | [optional] [readonly] 

## Methods

### NewConfigHost

`func NewConfigHost() *ConfigHost`

NewConfigHost instantiates a new ConfigHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigHostWithDefaults

`func NewConfigHostWithDefaults() *ConfigHost`

NewConfigHostWithDefaults instantiates a new ConfigHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbsoluteName

`func (o *ConfigHost) GetAbsoluteName() string`

GetAbsoluteName returns the AbsoluteName field if non-nil, zero value otherwise.

### GetAbsoluteNameOk

`func (o *ConfigHost) GetAbsoluteNameOk() (*string, bool)`

GetAbsoluteNameOk returns a tuple with the AbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbsoluteName

`func (o *ConfigHost) SetAbsoluteName(v string)`

SetAbsoluteName sets AbsoluteName field to given value.

### HasAbsoluteName

`func (o *ConfigHost) HasAbsoluteName() bool`

HasAbsoluteName returns a boolean if a field has been set.

### GetAddress

`func (o *ConfigHost) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ConfigHost) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ConfigHost) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ConfigHost) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAnycastAddresses

`func (o *ConfigHost) GetAnycastAddresses() []string`

GetAnycastAddresses returns the AnycastAddresses field if non-nil, zero value otherwise.

### GetAnycastAddressesOk

`func (o *ConfigHost) GetAnycastAddressesOk() (*[]string, bool)`

GetAnycastAddressesOk returns a tuple with the AnycastAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastAddresses

`func (o *ConfigHost) SetAnycastAddresses(v []string)`

SetAnycastAddresses sets AnycastAddresses field to given value.

### HasAnycastAddresses

`func (o *ConfigHost) HasAnycastAddresses() bool`

HasAnycastAddresses returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *ConfigHost) GetAssociatedServer() ConfigHostAssociatedServer`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *ConfigHost) GetAssociatedServerOk() (*ConfigHostAssociatedServer, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *ConfigHost) SetAssociatedServer(v ConfigHostAssociatedServer)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *ConfigHost) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetComment

`func (o *ConfigHost) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ConfigHost) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ConfigHost) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ConfigHost) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *ConfigHost) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *ConfigHost) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *ConfigHost) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *ConfigHost) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDfp

`func (o *ConfigHost) GetDfp() bool`

GetDfp returns the Dfp field if non-nil, zero value otherwise.

### GetDfpOk

`func (o *ConfigHost) GetDfpOk() (*bool, bool)`

GetDfpOk returns a tuple with the Dfp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfp

`func (o *ConfigHost) SetDfp(v bool)`

SetDfp sets Dfp field to given value.

### HasDfp

`func (o *ConfigHost) HasDfp() bool`

HasDfp returns a boolean if a field has been set.

### GetDfpService

`func (o *ConfigHost) GetDfpService() string`

GetDfpService returns the DfpService field if non-nil, zero value otherwise.

### GetDfpServiceOk

`func (o *ConfigHost) GetDfpServiceOk() (*string, bool)`

GetDfpServiceOk returns a tuple with the DfpService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDfpService

`func (o *ConfigHost) SetDfpService(v string)`

SetDfpService sets DfpService field to given value.

### HasDfpService

`func (o *ConfigHost) HasDfpService() bool`

HasDfpService returns a boolean if a field has been set.

### GetId

`func (o *ConfigHost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConfigHost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConfigHost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConfigHost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInheritanceSources

`func (o *ConfigHost) GetInheritanceSources() ConfigHostInheritance`

GetInheritanceSources returns the InheritanceSources field if non-nil, zero value otherwise.

### GetInheritanceSourcesOk

`func (o *ConfigHost) GetInheritanceSourcesOk() (*ConfigHostInheritance, bool)`

GetInheritanceSourcesOk returns a tuple with the InheritanceSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInheritanceSources

`func (o *ConfigHost) SetInheritanceSources(v ConfigHostInheritance)`

SetInheritanceSources sets InheritanceSources field to given value.

### HasInheritanceSources

`func (o *ConfigHost) HasInheritanceSources() bool`

HasInheritanceSources returns a boolean if a field has been set.

### GetKerberosKeys

`func (o *ConfigHost) GetKerberosKeys() []ConfigKerberosKey`

GetKerberosKeys returns the KerberosKeys field if non-nil, zero value otherwise.

### GetKerberosKeysOk

`func (o *ConfigHost) GetKerberosKeysOk() (*[]ConfigKerberosKey, bool)`

GetKerberosKeysOk returns a tuple with the KerberosKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeys

`func (o *ConfigHost) SetKerberosKeys(v []ConfigKerberosKey)`

SetKerberosKeys sets KerberosKeys field to given value.

### HasKerberosKeys

`func (o *ConfigHost) HasKerberosKeys() bool`

HasKerberosKeys returns a boolean if a field has been set.

### GetName

`func (o *ConfigHost) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConfigHost) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConfigHost) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConfigHost) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOphid

`func (o *ConfigHost) GetOphid() string`

GetOphid returns the Ophid field if non-nil, zero value otherwise.

### GetOphidOk

`func (o *ConfigHost) GetOphidOk() (*string, bool)`

GetOphidOk returns a tuple with the Ophid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOphid

`func (o *ConfigHost) SetOphid(v string)`

SetOphid sets Ophid field to given value.

### HasOphid

`func (o *ConfigHost) HasOphid() bool`

HasOphid returns a boolean if a field has been set.

### GetProtocolAbsoluteName

`func (o *ConfigHost) GetProtocolAbsoluteName() string`

GetProtocolAbsoluteName returns the ProtocolAbsoluteName field if non-nil, zero value otherwise.

### GetProtocolAbsoluteNameOk

`func (o *ConfigHost) GetProtocolAbsoluteNameOk() (*string, bool)`

GetProtocolAbsoluteNameOk returns a tuple with the ProtocolAbsoluteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolAbsoluteName

`func (o *ConfigHost) SetProtocolAbsoluteName(v string)`

SetProtocolAbsoluteName sets ProtocolAbsoluteName field to given value.

### HasProtocolAbsoluteName

`func (o *ConfigHost) HasProtocolAbsoluteName() bool`

HasProtocolAbsoluteName returns a boolean if a field has been set.

### GetProviderId

`func (o *ConfigHost) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *ConfigHost) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *ConfigHost) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *ConfigHost) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetServer

`func (o *ConfigHost) GetServer() string`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ConfigHost) GetServerOk() (*string, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ConfigHost) SetServer(v string)`

SetServer sets Server field to given value.

### HasServer

`func (o *ConfigHost) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSiteId

`func (o *ConfigHost) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ConfigHost) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ConfigHost) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ConfigHost) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTags

`func (o *ConfigHost) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ConfigHost) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ConfigHost) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *ConfigHost) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *ConfigHost) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConfigHost) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConfigHost) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConfigHost) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


