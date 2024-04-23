# AuthZoneExternalProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The identifier of the external provider. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the external provider. | [optional] 
**Type** | Pointer to **string** | Defines the type of external provider. Allowed values:  * _bloxone_ddi_: host type is BloxOne DDI,  * _microsoft_azure_: host type is Microsoft Azure,  * _amazon_web_service_: host type is Amazon Web Services,  * _microsoft_active_directory_: host type is Microsoft Active Directory,  * _google_cloud_platform_: host type is Google Cloud Platform. | [optional] 

## Methods

### NewAuthZoneExternalProvider

`func NewAuthZoneExternalProvider() *AuthZoneExternalProvider`

NewAuthZoneExternalProvider instantiates a new AuthZoneExternalProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthZoneExternalProviderWithDefaults

`func NewAuthZoneExternalProviderWithDefaults() *AuthZoneExternalProvider`

NewAuthZoneExternalProviderWithDefaults instantiates a new AuthZoneExternalProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AuthZoneExternalProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthZoneExternalProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthZoneExternalProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthZoneExternalProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AuthZoneExternalProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthZoneExternalProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthZoneExternalProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthZoneExternalProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *AuthZoneExternalProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthZoneExternalProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthZoneExternalProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuthZoneExternalProvider) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


