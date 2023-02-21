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

// checks if the SearchCertificatesRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCertificatesRestRequest{}

// SearchCertificatesRestRequest struct for SearchCertificatesRestRequest
type SearchCertificatesRestRequest struct {
	// Maximum number of results
	MaxNumberOfResults *int32 `json:"max_number_of_results,omitempty"`
	// A List of search criteria.
	Criteria             []SearchCertificateCriteriaRestRequest `json:"criteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchCertificatesRestRequest SearchCertificatesRestRequest

// NewSearchCertificatesRestRequest instantiates a new SearchCertificatesRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCertificatesRestRequest() *SearchCertificatesRestRequest {
	this := SearchCertificatesRestRequest{}
	return &this
}

// NewSearchCertificatesRestRequestWithDefaults instantiates a new SearchCertificatesRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCertificatesRestRequestWithDefaults() *SearchCertificatesRestRequest {
	this := SearchCertificatesRestRequest{}
	return &this
}

// GetMaxNumberOfResults returns the MaxNumberOfResults field value if set, zero value otherwise.
func (o *SearchCertificatesRestRequest) GetMaxNumberOfResults() int32 {
	if o == nil || isNil(o.MaxNumberOfResults) {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfResults
}

// GetMaxNumberOfResultsOk returns a tuple with the MaxNumberOfResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestRequest) GetMaxNumberOfResultsOk() (*int32, bool) {
	if o == nil || isNil(o.MaxNumberOfResults) {
		return nil, false
	}
	return o.MaxNumberOfResults, true
}

// HasMaxNumberOfResults returns a boolean if a field has been set.
func (o *SearchCertificatesRestRequest) HasMaxNumberOfResults() bool {
	if o != nil && !isNil(o.MaxNumberOfResults) {
		return true
	}

	return false
}

// SetMaxNumberOfResults gets a reference to the given int32 and assigns it to the MaxNumberOfResults field.
func (o *SearchCertificatesRestRequest) SetMaxNumberOfResults(v int32) {
	o.MaxNumberOfResults = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *SearchCertificatesRestRequest) GetCriteria() []SearchCertificateCriteriaRestRequest {
	if o == nil || isNil(o.Criteria) {
		var ret []SearchCertificateCriteriaRestRequest
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestRequest) GetCriteriaOk() ([]SearchCertificateCriteriaRestRequest, bool) {
	if o == nil || isNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *SearchCertificatesRestRequest) HasCriteria() bool {
	if o != nil && !isNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []SearchCertificateCriteriaRestRequest and assigns it to the Criteria field.
func (o *SearchCertificatesRestRequest) SetCriteria(v []SearchCertificateCriteriaRestRequest) {
	o.Criteria = v
}

func (o SearchCertificatesRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCertificatesRestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxNumberOfResults) {
		toSerialize["max_number_of_results"] = o.MaxNumberOfResults
	}
	if !isNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchCertificatesRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSearchCertificatesRestRequest := _SearchCertificatesRestRequest{}

	if err = json.Unmarshal(bytes, &varSearchCertificatesRestRequest); err == nil {
		*o = SearchCertificatesRestRequest(varSearchCertificatesRestRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "max_number_of_results")
		delete(additionalProperties, "criteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchCertificatesRestRequest struct {
	value *SearchCertificatesRestRequest
	isSet bool
}

func (v NullableSearchCertificatesRestRequest) Get() *SearchCertificatesRestRequest {
	return v.value
}

func (v *NullableSearchCertificatesRestRequest) Set(val *SearchCertificatesRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCertificatesRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCertificatesRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCertificatesRestRequest(val *SearchCertificatesRestRequest) *NullableSearchCertificatesRestRequest {
	return &NullableSearchCertificatesRestRequest{value: val, isSet: true}
}

func (v NullableSearchCertificatesRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCertificatesRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
