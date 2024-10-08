/*
BloxOne Anycast API

Anycast capability enables HA (High Availability) configuration of BloxOne applications that run on equipment located on customer's premises (on-prem hosts). Anycast supports DNS, as well as DNS-forwarding services.  Anycast-enabled application setups use multiple on-premises installations for one particular application type. Multiple application instances are configured to use the same endpoint address. Anycast capability is collocated with such application instance, monitoring the local application instance and advertising to the upstream router (a customer equipment) a per-instance, local route to the common application endpoint address, as long as the local application instance is available. Depending on the type of the upstream router, the customer may configure local route advertisement via either BGP (Boarder Gateway Protocol) or OSPF (Open Shortest Path First) routing protocols. Both protocols may be enabled as well. Multiple routes to the common application service address provide redundancy without the need to reconfigure application clients.  Should an application instance become unavailable, the local route advertisements stop, resulting in withdrawal of the route (in the upstream router) to the application instance that has gone out of service and ensuring that subsequent application requests thus get routed to the remaining available application instances.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anycast

import (
	"encoding/json"
)

// checks if the BgpConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpConfig{}

// BgpConfig struct for BgpConfig
type BgpConfig struct {
	Asn *int64 `json:"asn,omitempty"`
	// Examples:     ASDOT        ASPLAIN     INTEGER     VALID/INVALID     0.1          1           1           Valid     1            1           1           Valid     65535        65535       65535       Valid     0.65535      65535       65535       Valid     1.0          65536       65536       Valid     1.1          65537       65537       Valid     1.65535      131071      131071      Valid     65535.0      4294901760  4294901760  Valid     65535.1      4294901761  4294901761  Valid     65535.65535  4294967295  4294967295  Valid      0.65536                              Invalid     65535.655536                         Invalid     65536.0                              Invalid     65536.65535                          Invalid                  4294967296              Invalid
	AsnText       *string            `json:"asn_text,omitempty"`
	Fields        *ProtobufFieldMask `json:"fields,omitempty"`
	HolddownSecs  *int64             `json:"holddown_secs,omitempty"`
	KeepAliveSecs *int64             `json:"keep_alive_secs,omitempty"`
	LinkDetect    *bool              `json:"link_detect,omitempty"`
	Neighbors     []BgpNeighbor      `json:"neighbors,omitempty"`
	// Any predefined BGP configuration, with embedded new lines; the preamble will be prepended to the generated BGP configuration.
	Preamble             *string `json:"preamble,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpConfig BgpConfig

// NewBgpConfig instantiates a new BgpConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpConfig() *BgpConfig {
	this := BgpConfig{}
	return &this
}

// NewBgpConfigWithDefaults instantiates a new BgpConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpConfigWithDefaults() *BgpConfig {
	this := BgpConfig{}
	return &this
}

// GetAsn returns the Asn field value if set, zero value otherwise.
func (o *BgpConfig) GetAsn() int64 {
	if o == nil || IsNil(o.Asn) {
		var ret int64
		return ret
	}
	return *o.Asn
}

// GetAsnOk returns a tuple with the Asn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetAsnOk() (*int64, bool) {
	if o == nil || IsNil(o.Asn) {
		return nil, false
	}
	return o.Asn, true
}

// HasAsn returns a boolean if a field has been set.
func (o *BgpConfig) HasAsn() bool {
	if o != nil && !IsNil(o.Asn) {
		return true
	}

	return false
}

// SetAsn gets a reference to the given int64 and assigns it to the Asn field.
func (o *BgpConfig) SetAsn(v int64) {
	o.Asn = &v
}

// GetAsnText returns the AsnText field value if set, zero value otherwise.
func (o *BgpConfig) GetAsnText() string {
	if o == nil || IsNil(o.AsnText) {
		var ret string
		return ret
	}
	return *o.AsnText
}

// GetAsnTextOk returns a tuple with the AsnText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetAsnTextOk() (*string, bool) {
	if o == nil || IsNil(o.AsnText) {
		return nil, false
	}
	return o.AsnText, true
}

// HasAsnText returns a boolean if a field has been set.
func (o *BgpConfig) HasAsnText() bool {
	if o != nil && !IsNil(o.AsnText) {
		return true
	}

	return false
}

// SetAsnText gets a reference to the given string and assigns it to the AsnText field.
func (o *BgpConfig) SetAsnText(v string) {
	o.AsnText = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *BgpConfig) GetFields() ProtobufFieldMask {
	if o == nil || IsNil(o.Fields) {
		var ret ProtobufFieldMask
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetFieldsOk() (*ProtobufFieldMask, bool) {
	if o == nil || IsNil(o.Fields) {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *BgpConfig) HasFields() bool {
	if o != nil && !IsNil(o.Fields) {
		return true
	}

	return false
}

// SetFields gets a reference to the given ProtobufFieldMask and assigns it to the Fields field.
func (o *BgpConfig) SetFields(v ProtobufFieldMask) {
	o.Fields = &v
}

// GetHolddownSecs returns the HolddownSecs field value if set, zero value otherwise.
func (o *BgpConfig) GetHolddownSecs() int64 {
	if o == nil || IsNil(o.HolddownSecs) {
		var ret int64
		return ret
	}
	return *o.HolddownSecs
}

// GetHolddownSecsOk returns a tuple with the HolddownSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetHolddownSecsOk() (*int64, bool) {
	if o == nil || IsNil(o.HolddownSecs) {
		return nil, false
	}
	return o.HolddownSecs, true
}

// HasHolddownSecs returns a boolean if a field has been set.
func (o *BgpConfig) HasHolddownSecs() bool {
	if o != nil && !IsNil(o.HolddownSecs) {
		return true
	}

	return false
}

// SetHolddownSecs gets a reference to the given int64 and assigns it to the HolddownSecs field.
func (o *BgpConfig) SetHolddownSecs(v int64) {
	o.HolddownSecs = &v
}

// GetKeepAliveSecs returns the KeepAliveSecs field value if set, zero value otherwise.
func (o *BgpConfig) GetKeepAliveSecs() int64 {
	if o == nil || IsNil(o.KeepAliveSecs) {
		var ret int64
		return ret
	}
	return *o.KeepAliveSecs
}

// GetKeepAliveSecsOk returns a tuple with the KeepAliveSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetKeepAliveSecsOk() (*int64, bool) {
	if o == nil || IsNil(o.KeepAliveSecs) {
		return nil, false
	}
	return o.KeepAliveSecs, true
}

// HasKeepAliveSecs returns a boolean if a field has been set.
func (o *BgpConfig) HasKeepAliveSecs() bool {
	if o != nil && !IsNil(o.KeepAliveSecs) {
		return true
	}

	return false
}

// SetKeepAliveSecs gets a reference to the given int64 and assigns it to the KeepAliveSecs field.
func (o *BgpConfig) SetKeepAliveSecs(v int64) {
	o.KeepAliveSecs = &v
}

// GetLinkDetect returns the LinkDetect field value if set, zero value otherwise.
func (o *BgpConfig) GetLinkDetect() bool {
	if o == nil || IsNil(o.LinkDetect) {
		var ret bool
		return ret
	}
	return *o.LinkDetect
}

// GetLinkDetectOk returns a tuple with the LinkDetect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetLinkDetectOk() (*bool, bool) {
	if o == nil || IsNil(o.LinkDetect) {
		return nil, false
	}
	return o.LinkDetect, true
}

// HasLinkDetect returns a boolean if a field has been set.
func (o *BgpConfig) HasLinkDetect() bool {
	if o != nil && !IsNil(o.LinkDetect) {
		return true
	}

	return false
}

// SetLinkDetect gets a reference to the given bool and assigns it to the LinkDetect field.
func (o *BgpConfig) SetLinkDetect(v bool) {
	o.LinkDetect = &v
}

// GetNeighbors returns the Neighbors field value if set, zero value otherwise.
func (o *BgpConfig) GetNeighbors() []BgpNeighbor {
	if o == nil || IsNil(o.Neighbors) {
		var ret []BgpNeighbor
		return ret
	}
	return o.Neighbors
}

// GetNeighborsOk returns a tuple with the Neighbors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetNeighborsOk() ([]BgpNeighbor, bool) {
	if o == nil || IsNil(o.Neighbors) {
		return nil, false
	}
	return o.Neighbors, true
}

// HasNeighbors returns a boolean if a field has been set.
func (o *BgpConfig) HasNeighbors() bool {
	if o != nil && !IsNil(o.Neighbors) {
		return true
	}

	return false
}

// SetNeighbors gets a reference to the given []BgpNeighbor and assigns it to the Neighbors field.
func (o *BgpConfig) SetNeighbors(v []BgpNeighbor) {
	o.Neighbors = v
}

// GetPreamble returns the Preamble field value if set, zero value otherwise.
func (o *BgpConfig) GetPreamble() string {
	if o == nil || IsNil(o.Preamble) {
		var ret string
		return ret
	}
	return *o.Preamble
}

// GetPreambleOk returns a tuple with the Preamble field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpConfig) GetPreambleOk() (*string, bool) {
	if o == nil || IsNil(o.Preamble) {
		return nil, false
	}
	return o.Preamble, true
}

// HasPreamble returns a boolean if a field has been set.
func (o *BgpConfig) HasPreamble() bool {
	if o != nil && !IsNil(o.Preamble) {
		return true
	}

	return false
}

// SetPreamble gets a reference to the given string and assigns it to the Preamble field.
func (o *BgpConfig) SetPreamble(v string) {
	o.Preamble = &v
}

func (o BgpConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Asn) {
		toSerialize["asn"] = o.Asn
	}
	if !IsNil(o.AsnText) {
		toSerialize["asn_text"] = o.AsnText
	}
	if !IsNil(o.Fields) {
		toSerialize["fields"] = o.Fields
	}
	if !IsNil(o.HolddownSecs) {
		toSerialize["holddown_secs"] = o.HolddownSecs
	}
	if !IsNil(o.KeepAliveSecs) {
		toSerialize["keep_alive_secs"] = o.KeepAliveSecs
	}
	if !IsNil(o.LinkDetect) {
		toSerialize["link_detect"] = o.LinkDetect
	}
	if !IsNil(o.Neighbors) {
		toSerialize["neighbors"] = o.Neighbors
	}
	if !IsNil(o.Preamble) {
		toSerialize["preamble"] = o.Preamble
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpConfig) UnmarshalJSON(data []byte) (err error) {
	varBgpConfig := _BgpConfig{}

	err = json.Unmarshal(data, &varBgpConfig)

	if err != nil {
		return err
	}

	*o = BgpConfig(varBgpConfig)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "asn")
		delete(additionalProperties, "asn_text")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "holddown_secs")
		delete(additionalProperties, "keep_alive_secs")
		delete(additionalProperties, "link_detect")
		delete(additionalProperties, "neighbors")
		delete(additionalProperties, "preamble")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpConfig struct {
	value *BgpConfig
	isSet bool
}

func (v NullableBgpConfig) Get() *BgpConfig {
	return v.value
}

func (v *NullableBgpConfig) Set(val *BgpConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpConfig(val *BgpConfig) *NullableBgpConfig {
	return &NullableBgpConfig{value: val, isSet: true}
}

func (v NullableBgpConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
