/*
Infrastructure Management API

The **Infrastructure Management API** provides a RESTful interface to manage Infrastructure Hosts and Services objects.  The following is a list of the different Services and their string types (the string types are to be used with the APIs for the `service_type` field):  | Service name | Service type |   | ------ | ------ |   | Access Authentication | authn |   | Anycast | anycast |   | Data Connector | cdc |   | DHCP | dhcp |   | DNS | dns |   | DNS Forwarding Proxy | dfp |   | NIOS Grid Connector | orpheus |   | MS AD Sync | msad |   | NTP | ntp |   | BGP | bgp |   | RIP | rip |   | OSPF | ospf |    ---   ### Hosts API  The Hosts API is used to manage the Infrastructure Host resources. These include various operations related to hosts such as viewing, creating, updating, replacing, disconnecting, and deleting Hosts. Management of Hosts is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Hosts tab.  ---   ### Services API  The Services API is used to manage the Infrastructure Service resources (a.k.a. BloxOne applications). These include various operations related to hosts such as viewing, creating, updating, starting/stopping, configuring, and deleting Services. Management of Services is done from the Cloud Services Portal (CSP) by navigating to the Manage -> Infrastructure -> Services tab.  ---   ### Detail APIs  The Detail APIs are read-only APIs used to list all the Infrastructure resources (Hosts and Services). Each resource record returned also contains information about its other associated resources and the status information for itself and the associated resource(s) (i.e., Host/Service status).  ---   

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infra_mgmt

import (
	"encoding/json"
)

// checks if the InfraApplicationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraApplicationsResponse{}

// InfraApplicationsResponse Applications response.
type InfraApplicationsResponse struct {
	Results *InfraApplications `json:"results,omitempty"`
}

// NewInfraApplicationsResponse instantiates a new InfraApplicationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraApplicationsResponse() *InfraApplicationsResponse {
	this := InfraApplicationsResponse{}
	return &this
}

// NewInfraApplicationsResponseWithDefaults instantiates a new InfraApplicationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraApplicationsResponseWithDefaults() *InfraApplicationsResponse {
	this := InfraApplicationsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *InfraApplicationsResponse) GetResults() InfraApplications {
	if o == nil || IsNil(o.Results) {
		var ret InfraApplications
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraApplicationsResponse) GetResultsOk() (*InfraApplications, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *InfraApplicationsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given InfraApplications and assigns it to the Results field.
func (o *InfraApplicationsResponse) SetResults(v InfraApplications) {
	o.Results = &v
}

func (o InfraApplicationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraApplicationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableInfraApplicationsResponse struct {
	value *InfraApplicationsResponse
	isSet bool
}

func (v NullableInfraApplicationsResponse) Get() *InfraApplicationsResponse {
	return v.value
}

func (v *NullableInfraApplicationsResponse) Set(val *InfraApplicationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraApplicationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraApplicationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraApplicationsResponse(val *InfraApplicationsResponse) *NullableInfraApplicationsResponse {
	return &NullableInfraApplicationsResponse{value: val, isSet: true}
}

func (v NullableInfraApplicationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraApplicationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


