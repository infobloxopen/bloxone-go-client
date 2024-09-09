# DHCPServiceInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedHosts** | Pointer to [**[]AssociatedHost**](AssociatedHost.md) |  | [optional] 
**AssociatedServer** | Pointer to [**HostAssociatedServer**](HostAssociatedServer.md) |  | [optional] 
**Comment** | Pointer to **string** | The comment for the service. | [optional] [readonly] 
**Id** | Pointer to **string** | The resource identifier. | [optional] [readonly] 
**IpSpace** | Pointer to **string** | The resource identifier. | [optional] 
**Name** | Pointer to **string** | The display name of the service. | [optional] [readonly] 
**Tags** | Pointer to **map[string]interface{}** | The tags of the service host in JSON format. | [optional] 

## Methods

### NewDHCPServiceInstance

`func NewDHCPServiceInstance() *DHCPServiceInstance`

NewDHCPServiceInstance instantiates a new DHCPServiceInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDHCPServiceInstanceWithDefaults

`func NewDHCPServiceInstanceWithDefaults() *DHCPServiceInstance`

NewDHCPServiceInstanceWithDefaults instantiates a new DHCPServiceInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociatedHosts

`func (o *DHCPServiceInstance) GetAssociatedHosts() []AssociatedHost`

GetAssociatedHosts returns the AssociatedHosts field if non-nil, zero value otherwise.

### GetAssociatedHostsOk

`func (o *DHCPServiceInstance) GetAssociatedHostsOk() (*[]AssociatedHost, bool)`

GetAssociatedHostsOk returns a tuple with the AssociatedHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedHosts

`func (o *DHCPServiceInstance) SetAssociatedHosts(v []AssociatedHost)`

SetAssociatedHosts sets AssociatedHosts field to given value.

### HasAssociatedHosts

`func (o *DHCPServiceInstance) HasAssociatedHosts() bool`

HasAssociatedHosts returns a boolean if a field has been set.

### GetAssociatedServer

`func (o *DHCPServiceInstance) GetAssociatedServer() HostAssociatedServer`

GetAssociatedServer returns the AssociatedServer field if non-nil, zero value otherwise.

### GetAssociatedServerOk

`func (o *DHCPServiceInstance) GetAssociatedServerOk() (*HostAssociatedServer, bool)`

GetAssociatedServerOk returns a tuple with the AssociatedServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociatedServer

`func (o *DHCPServiceInstance) SetAssociatedServer(v HostAssociatedServer)`

SetAssociatedServer sets AssociatedServer field to given value.

### HasAssociatedServer

`func (o *DHCPServiceInstance) HasAssociatedServer() bool`

HasAssociatedServer returns a boolean if a field has been set.

### GetComment

`func (o *DHCPServiceInstance) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *DHCPServiceInstance) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *DHCPServiceInstance) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *DHCPServiceInstance) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetId

`func (o *DHCPServiceInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DHCPServiceInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DHCPServiceInstance) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DHCPServiceInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpSpace

`func (o *DHCPServiceInstance) GetIpSpace() string`

GetIpSpace returns the IpSpace field if non-nil, zero value otherwise.

### GetIpSpaceOk

`func (o *DHCPServiceInstance) GetIpSpaceOk() (*string, bool)`

GetIpSpaceOk returns a tuple with the IpSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSpace

`func (o *DHCPServiceInstance) SetIpSpace(v string)`

SetIpSpace sets IpSpace field to given value.

### HasIpSpace

`func (o *DHCPServiceInstance) HasIpSpace() bool`

HasIpSpace returns a boolean if a field has been set.

### GetName

`func (o *DHCPServiceInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DHCPServiceInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DHCPServiceInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DHCPServiceInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTags

`func (o *DHCPServiceInstance) GetTags() map[string]interface{}`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DHCPServiceInstance) GetTagsOk() (*map[string]interface{}, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DHCPServiceInstance) SetTags(v map[string]interface{})`

SetTags sets Tags field to given value.

### HasTags

`func (o *DHCPServiceInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


