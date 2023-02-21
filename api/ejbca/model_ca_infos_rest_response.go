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

// checks if the CaInfosRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaInfosRestResponse{}

// CaInfosRestResponse struct for CaInfosRestResponse
type CaInfosRestResponse struct {
	CertificateAuthorities []CaInfoRestResponse `json:"certificate_authorities,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaInfosRestResponse CaInfosRestResponse

// NewCaInfosRestResponse instantiates a new CaInfosRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaInfosRestResponse() *CaInfosRestResponse {
	this := CaInfosRestResponse{}
	return &this
}

// NewCaInfosRestResponseWithDefaults instantiates a new CaInfosRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaInfosRestResponseWithDefaults() *CaInfosRestResponse {
	this := CaInfosRestResponse{}
	return &this
}

// GetCertificateAuthorities returns the CertificateAuthorities field value if set, zero value otherwise.
func (o *CaInfosRestResponse) GetCertificateAuthorities() []CaInfoRestResponse {
	if o == nil || isNil(o.CertificateAuthorities) {
		var ret []CaInfoRestResponse
		return ret
	}
	return o.CertificateAuthorities
}

// GetCertificateAuthoritiesOk returns a tuple with the CertificateAuthorities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaInfosRestResponse) GetCertificateAuthoritiesOk() ([]CaInfoRestResponse, bool) {
	if o == nil || isNil(o.CertificateAuthorities) {
		return nil, false
	}
	return o.CertificateAuthorities, true
}

// HasCertificateAuthorities returns a boolean if a field has been set.
func (o *CaInfosRestResponse) HasCertificateAuthorities() bool {
	if o != nil && !isNil(o.CertificateAuthorities) {
		return true
	}

	return false
}

// SetCertificateAuthorities gets a reference to the given []CaInfoRestResponse and assigns it to the CertificateAuthorities field.
func (o *CaInfosRestResponse) SetCertificateAuthorities(v []CaInfoRestResponse) {
	o.CertificateAuthorities = v
}

func (o CaInfosRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaInfosRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CertificateAuthorities) {
		toSerialize["certificate_authorities"] = o.CertificateAuthorities
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaInfosRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCaInfosRestResponse := _CaInfosRestResponse{}

	if err = json.Unmarshal(bytes, &varCaInfosRestResponse); err == nil {
		*o = CaInfosRestResponse(varCaInfosRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "certificate_authorities")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaInfosRestResponse struct {
	value *CaInfosRestResponse
	isSet bool
}

func (v NullableCaInfosRestResponse) Get() *CaInfosRestResponse {
	return v.value
}

func (v *NullableCaInfosRestResponse) Set(val *CaInfosRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCaInfosRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCaInfosRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaInfosRestResponse(val *CaInfosRestResponse) *NullableCaInfosRestResponse {
	return &NullableCaInfosRestResponse{value: val, isSet: true}
}

func (v NullableCaInfosRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaInfosRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


