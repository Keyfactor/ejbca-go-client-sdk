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

// checks if the SearchCertificatesRestResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCertificatesRestResponseV2{}

// SearchCertificatesRestResponseV2 struct for SearchCertificatesRestResponseV2
type SearchCertificatesRestResponseV2 struct {
	Certificates []CertificateRestResponseV2 `json:"certificates,omitempty"`
	PaginationSummary *PaginationSummary `json:"pagination_summary,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchCertificatesRestResponseV2 SearchCertificatesRestResponseV2

// NewSearchCertificatesRestResponseV2 instantiates a new SearchCertificatesRestResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCertificatesRestResponseV2() *SearchCertificatesRestResponseV2 {
	this := SearchCertificatesRestResponseV2{}
	return &this
}

// NewSearchCertificatesRestResponseV2WithDefaults instantiates a new SearchCertificatesRestResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCertificatesRestResponseV2WithDefaults() *SearchCertificatesRestResponseV2 {
	this := SearchCertificatesRestResponseV2{}
	return &this
}

// GetCertificates returns the Certificates field value if set, zero value otherwise.
func (o *SearchCertificatesRestResponseV2) GetCertificates() []CertificateRestResponseV2 {
	if o == nil || isNil(o.Certificates) {
		var ret []CertificateRestResponseV2
		return ret
	}
	return o.Certificates
}

// GetCertificatesOk returns a tuple with the Certificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestResponseV2) GetCertificatesOk() ([]CertificateRestResponseV2, bool) {
	if o == nil || isNil(o.Certificates) {
		return nil, false
	}
	return o.Certificates, true
}

// HasCertificates returns a boolean if a field has been set.
func (o *SearchCertificatesRestResponseV2) HasCertificates() bool {
	if o != nil && !isNil(o.Certificates) {
		return true
	}

	return false
}

// SetCertificates gets a reference to the given []CertificateRestResponseV2 and assigns it to the Certificates field.
func (o *SearchCertificatesRestResponseV2) SetCertificates(v []CertificateRestResponseV2) {
	o.Certificates = v
}

// GetPaginationSummary returns the PaginationSummary field value if set, zero value otherwise.
func (o *SearchCertificatesRestResponseV2) GetPaginationSummary() PaginationSummary {
	if o == nil || isNil(o.PaginationSummary) {
		var ret PaginationSummary
		return ret
	}
	return *o.PaginationSummary
}

// GetPaginationSummaryOk returns a tuple with the PaginationSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestResponseV2) GetPaginationSummaryOk() (*PaginationSummary, bool) {
	if o == nil || isNil(o.PaginationSummary) {
		return nil, false
	}
	return o.PaginationSummary, true
}

// HasPaginationSummary returns a boolean if a field has been set.
func (o *SearchCertificatesRestResponseV2) HasPaginationSummary() bool {
	if o != nil && !isNil(o.PaginationSummary) {
		return true
	}

	return false
}

// SetPaginationSummary gets a reference to the given PaginationSummary and assigns it to the PaginationSummary field.
func (o *SearchCertificatesRestResponseV2) SetPaginationSummary(v PaginationSummary) {
	o.PaginationSummary = &v
}

func (o SearchCertificatesRestResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCertificatesRestResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Certificates) {
		toSerialize["certificates"] = o.Certificates
	}
	if !isNil(o.PaginationSummary) {
		toSerialize["pagination_summary"] = o.PaginationSummary
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchCertificatesRestResponseV2) UnmarshalJSON(bytes []byte) (err error) {
	varSearchCertificatesRestResponseV2 := _SearchCertificatesRestResponseV2{}

	if err = json.Unmarshal(bytes, &varSearchCertificatesRestResponseV2); err == nil {
		*o = SearchCertificatesRestResponseV2(varSearchCertificatesRestResponseV2)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "certificates")
		delete(additionalProperties, "pagination_summary")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchCertificatesRestResponseV2 struct {
	value *SearchCertificatesRestResponseV2
	isSet bool
}

func (v NullableSearchCertificatesRestResponseV2) Get() *SearchCertificatesRestResponseV2 {
	return v.value
}

func (v *NullableSearchCertificatesRestResponseV2) Set(val *SearchCertificatesRestResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCertificatesRestResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCertificatesRestResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCertificatesRestResponseV2(val *SearchCertificatesRestResponseV2) *NullableSearchCertificatesRestResponseV2 {
	return &NullableSearchCertificatesRestResponseV2{value: val, isSet: true}
}

func (v NullableSearchCertificatesRestResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCertificatesRestResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


