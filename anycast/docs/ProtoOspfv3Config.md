# ProtoOspfv3Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area** | Pointer to **string** | OSPF area identifier; usually in the format of an IPv4 address (although not an address itself) | [optional] 
**Cost** | Pointer to **int64** |  | [optional] 
**DeadInterval** | Pointer to **int64** |  | [optional] 
**HelloInterval** | Pointer to **int64** |  | [optional] 
**Interface** | Pointer to **string** | Name of the interface that is configured with external IP address of the host | [optional] 
**RetransmitInterval** | Pointer to **int64** |  | [optional] 
**TransmitDelay** | Pointer to **int64** |  | [optional] 

## Methods

### NewProtoOspfv3Config

`func NewProtoOspfv3Config() *ProtoOspfv3Config`

NewProtoOspfv3Config instantiates a new ProtoOspfv3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProtoOspfv3ConfigWithDefaults

`func NewProtoOspfv3ConfigWithDefaults() *ProtoOspfv3Config`

NewProtoOspfv3ConfigWithDefaults instantiates a new ProtoOspfv3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *ProtoOspfv3Config) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ProtoOspfv3Config) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ProtoOspfv3Config) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *ProtoOspfv3Config) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *ProtoOspfv3Config) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ProtoOspfv3Config) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ProtoOspfv3Config) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ProtoOspfv3Config) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *ProtoOspfv3Config) GetDeadInterval() int64`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *ProtoOspfv3Config) GetDeadIntervalOk() (*int64, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *ProtoOspfv3Config) SetDeadInterval(v int64)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *ProtoOspfv3Config) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetHelloInterval

`func (o *ProtoOspfv3Config) GetHelloInterval() int64`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *ProtoOspfv3Config) GetHelloIntervalOk() (*int64, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *ProtoOspfv3Config) SetHelloInterval(v int64)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *ProtoOspfv3Config) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetInterface

`func (o *ProtoOspfv3Config) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ProtoOspfv3Config) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ProtoOspfv3Config) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ProtoOspfv3Config) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *ProtoOspfv3Config) GetRetransmitInterval() int64`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *ProtoOspfv3Config) GetRetransmitIntervalOk() (*int64, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *ProtoOspfv3Config) SetRetransmitInterval(v int64)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *ProtoOspfv3Config) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetTransmitDelay

`func (o *ProtoOspfv3Config) GetTransmitDelay() int64`

GetTransmitDelay returns the TransmitDelay field if non-nil, zero value otherwise.

### GetTransmitDelayOk

`func (o *ProtoOspfv3Config) GetTransmitDelayOk() (*int64, bool)`

GetTransmitDelayOk returns a tuple with the TransmitDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitDelay

`func (o *ProtoOspfv3Config) SetTransmitDelay(v int64)`

SetTransmitDelay sets TransmitDelay field to given value.

### HasTransmitDelay

`func (o *ProtoOspfv3Config) HasTransmitDelay() bool`

HasTransmitDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


