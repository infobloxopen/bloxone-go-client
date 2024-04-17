# ProtoAnycastConfigRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnycastConfigName** | Pointer to **string** |  | [optional] 
**RoutingProtocols** | Pointer to **[]string** | Routing protocols enabled for this anycast configuration, on a particular host. Valid protocol names are \&quot;BGP\&quot;, \&quot;OSPF\&quot;/\&quot;OSPFv2\&quot;, \&quot;OSPFv3\&quot;. | [optional] 

## Methods

### NewProtoAnycastConfigRef

`func NewProtoAnycastConfigRef() *ProtoAnycastConfigRef`

NewProtoAnycastConfigRef instantiates a new ProtoAnycastConfigRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoAnycastConfigRefWithDefaults

`func NewProtoAnycastConfigRefWithDefaults() *ProtoAnycastConfigRef`

NewProtoAnycastConfigRefWithDefaults instantiates a new ProtoAnycastConfigRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnycastConfigName

`func (o *ProtoAnycastConfigRef) GetAnycastConfigName() string`

GetAnycastConfigName returns the AnycastConfigName field if non-nil, zero value otherwise.

### GetAnycastConfigNameOk

`func (o *ProtoAnycastConfigRef) GetAnycastConfigNameOk() (*string, bool)`

GetAnycastConfigNameOk returns a tuple with the AnycastConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnycastConfigName

`func (o *ProtoAnycastConfigRef) SetAnycastConfigName(v string)`

SetAnycastConfigName sets AnycastConfigName field to given value.

### HasAnycastConfigName

`func (o *ProtoAnycastConfigRef) HasAnycastConfigName() bool`

HasAnycastConfigName returns a boolean if a field has been set.

### GetRoutingProtocols

`func (o *ProtoAnycastConfigRef) GetRoutingProtocols() []string`

GetRoutingProtocols returns the RoutingProtocols field if non-nil, zero value otherwise.

### GetRoutingProtocolsOk

`func (o *ProtoAnycastConfigRef) GetRoutingProtocolsOk() (*[]string, bool)`

GetRoutingProtocolsOk returns a tuple with the RoutingProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingProtocols

`func (o *ProtoAnycastConfigRef) SetRoutingProtocols(v []string)`

SetRoutingProtocols sets RoutingProtocols field to given value.

### HasRoutingProtocols

`func (o *ProtoAnycastConfigRef) HasRoutingProtocols() bool`

HasRoutingProtocols returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


