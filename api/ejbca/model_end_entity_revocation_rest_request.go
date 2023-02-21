/*
Copyright 2022 Keyfactor
Licensed under the Apache License, Version 2.0 (the "License"); you may
not use this file except in compliance with the License.  You may obtain a
copy of the License at http://www.apache.org/licenses/LICENSE-2.0.  Unless
required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
OR CONDITIONS OF ANY KIND, either express or implied. See the License for
thespecific language governing permissions and limitations under the
License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"encoding/json"
)

// checks if the EndEntityRevocationRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndEntityRevocationRestRequest{}

// EndEntityRevocationRestRequest End Entity revocation request. Available reason codes:   0 - Unspecified,  1 - Key Compromise,  2 - CA Compromise,  3 - Affiliation Changed,  4 - Superseded,  5 - Cessation of Operation,  6 - Certificate Hold,  8 - Remove from CRL,  9 - Privileges Withdrawn,  10 - AA Compromise
type EndEntityRevocationRestRequest struct {
	// Reason code
	ReasonCode *int32 `json:"reason_code,omitempty"`
	// Delete
	Delete               *bool `json:"delete,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EndEntityRevocationRestRequest EndEntityRevocationRestRequest

// NewEndEntityRevocationRestRequest instantiates a new EndEntityRevocationRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndEntityRevocationRestRequest() *EndEntityRevocationRestRequest {
	this := EndEntityRevocationRestRequest{}
	return &this
}

// NewEndEntityRevocationRestRequestWithDefaults instantiates a new EndEntityRevocationRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndEntityRevocationRestRequestWithDefaults() *EndEntityRevocationRestRequest {
	this := EndEntityRevocationRestRequest{}
	return &this
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *EndEntityRevocationRestRequest) GetReasonCode() int32 {
	if o == nil || isNil(o.ReasonCode) {
		var ret int32
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndEntityRevocationRestRequest) GetReasonCodeOk() (*int32, bool) {
	if o == nil || isNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *EndEntityRevocationRestRequest) HasReasonCode() bool {
	if o != nil && !isNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given int32 and assigns it to the ReasonCode field.
func (o *EndEntityRevocationRestRequest) SetReasonCode(v int32) {
	o.ReasonCode = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *EndEntityRevocationRestRequest) GetDelete() bool {
	if o == nil || isNil(o.Delete) {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndEntityRevocationRestRequest) GetDeleteOk() (*bool, bool) {
	if o == nil || isNil(o.Delete) {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *EndEntityRevocationRestRequest) HasDelete() bool {
	if o != nil && !isNil(o.Delete) {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *EndEntityRevocationRestRequest) SetDelete(v bool) {
	o.Delete = &v
}

func (o EndEntityRevocationRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndEntityRevocationRestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	if !isNil(o.Delete) {
		toSerialize["delete"] = o.Delete
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EndEntityRevocationRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varEndEntityRevocationRestRequest := _EndEntityRevocationRestRequest{}

	if err = json.Unmarshal(bytes, &varEndEntityRevocationRestRequest); err == nil {
		*o = EndEntityRevocationRestRequest(varEndEntityRevocationRestRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reason_code")
		delete(additionalProperties, "delete")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEndEntityRevocationRestRequest struct {
	value *EndEntityRevocationRestRequest
	isSet bool
}

func (v NullableEndEntityRevocationRestRequest) Get() *EndEntityRevocationRestRequest {
	return v.value
}

func (v *NullableEndEntityRevocationRestRequest) Set(val *EndEntityRevocationRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEndEntityRevocationRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEndEntityRevocationRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndEntityRevocationRestRequest(val *EndEntityRevocationRestRequest) *NullableEndEntityRevocationRestRequest {
	return &NullableEndEntityRevocationRestRequest{value: val, isSet: true}
}

func (v NullableEndEntityRevocationRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndEntityRevocationRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
