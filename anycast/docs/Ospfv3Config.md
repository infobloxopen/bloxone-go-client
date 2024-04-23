# Ospfv3Config

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

### NewOspfv3Config

`func NewOspfv3Config() *Ospfv3Config`

NewOspfv3Config instantiates a new Ospfv3Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfv3ConfigWithDefaults

`func NewOspfv3ConfigWithDefaults() *Ospfv3Config`

NewOspfv3ConfigWithDefaults instantiates a new Ospfv3Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArea

`func (o *Ospfv3Config) GetArea() string`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *Ospfv3Config) GetAreaOk() (*string, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *Ospfv3Config) SetArea(v string)`

SetArea sets Area field to given value.

### HasArea

`func (o *Ospfv3Config) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCost

`func (o *Ospfv3Config) GetCost() int64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Ospfv3Config) GetCostOk() (*int64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Ospfv3Config) SetCost(v int64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Ospfv3Config) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetDeadInterval

`func (o *Ospfv3Config) GetDeadInterval() int64`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *Ospfv3Config) GetDeadIntervalOk() (*int64, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *Ospfv3Config) SetDeadInterval(v int64)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *Ospfv3Config) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetHelloInterval

`func (o *Ospfv3Config) GetHelloInterval() int64`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *Ospfv3Config) GetHelloIntervalOk() (*int64, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *Ospfv3Config) SetHelloInterval(v int64)`

SetHelloInterval sets HelloInterval field to given value.

### HasHelloInterval

`func (o *Ospfv3Config) HasHelloInterval() bool`

HasHelloInterval returns a boolean if a field has been set.

### GetInterface

`func (o *Ospfv3Config) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *Ospfv3Config) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *Ospfv3Config) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *Ospfv3Config) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetRetransmitInterval

`func (o *Ospfv3Config) GetRetransmitInterval() int64`

GetRetransmitInterval returns the RetransmitInterval field if non-nil, zero value otherwise.

### GetRetransmitIntervalOk

`func (o *Ospfv3Config) GetRetransmitIntervalOk() (*int64, bool)`

GetRetransmitIntervalOk returns a tuple with the RetransmitInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetransmitInterval

`func (o *Ospfv3Config) SetRetransmitInterval(v int64)`

SetRetransmitInterval sets RetransmitInterval field to given value.

### HasRetransmitInterval

`func (o *Ospfv3Config) HasRetransmitInterval() bool`

HasRetransmitInterval returns a boolean if a field has been set.

### GetTransmitDelay

`func (o *Ospfv3Config) GetTransmitDelay() int64`

GetTransmitDelay returns the TransmitDelay field if non-nil, zero value otherwise.

### GetTransmitDelayOk

`func (o *Ospfv3Config) GetTransmitDelayOk() (*int64, bool)`

GetTransmitDelayOk returns a tuple with the TransmitDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransmitDelay

`func (o *Ospfv3Config) SetTransmitDelay(v int64)`

SetTransmitDelay sets TransmitDelay field to given value.

### HasTransmitDelay

`func (o *Ospfv3Config) HasTransmitDelay() bool`

HasTransmitDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


