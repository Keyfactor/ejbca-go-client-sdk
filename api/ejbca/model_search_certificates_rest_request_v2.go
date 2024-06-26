/*
Copyright 2024 Keyfactor

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"encoding/json"
)

// checks if the SearchCertificatesRestRequestV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCertificatesRestRequestV2{}

// SearchCertificatesRestRequestV2 struct for SearchCertificatesRestRequestV2
type SearchCertificatesRestRequestV2 struct {
	Pagination *Pagination                       `json:"pagination,omitempty"`
	Sort       *SearchCertificateSortRestRequest `json:"sort,omitempty"`
	// A List of search criteria.
	Criteria             []SearchCertificateCriteriaRestRequest `json:"criteria,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchCertificatesRestRequestV2 SearchCertificatesRestRequestV2

// NewSearchCertificatesRestRequestV2 instantiates a new SearchCertificatesRestRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCertificatesRestRequestV2() *SearchCertificatesRestRequestV2 {
	this := SearchCertificatesRestRequestV2{}
	return &this
}

// NewSearchCertificatesRestRequestV2WithDefaults instantiates a new SearchCertificatesRestRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCertificatesRestRequestV2WithDefaults() *SearchCertificatesRestRequestV2 {
	this := SearchCertificatesRestRequestV2{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *SearchCertificatesRestRequestV2) GetPagination() Pagination {
	if o == nil || isNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestRequestV2) GetPaginationOk() (*Pagination, bool) {
	if o == nil || isNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *SearchCertificatesRestRequestV2) HasPagination() bool {
	if o != nil && !isNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *SearchCertificatesRestRequestV2) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetSort returns the Sort field value if set, zero value otherwise.
func (o *SearchCertificatesRestRequestV2) GetSort() SearchCertificateSortRestRequest {
	if o == nil || isNil(o.Sort) {
		var ret SearchCertificateSortRestRequest
		return ret
	}
	return *o.Sort
}

// GetSortOk returns a tuple with the Sort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestRequestV2) GetSortOk() (*SearchCertificateSortRestRequest, bool) {
	if o == nil || isNil(o.Sort) {
		return nil, false
	}
	return o.Sort, true
}

// HasSort returns a boolean if a field has been set.
func (o *SearchCertificatesRestRequestV2) HasSort() bool {
	if o != nil && !isNil(o.Sort) {
		return true
	}

	return false
}

// SetSort gets a reference to the given SearchCertificateSortRestRequest and assigns it to the Sort field.
func (o *SearchCertificatesRestRequestV2) SetSort(v SearchCertificateSortRestRequest) {
	o.Sort = &v
}

// GetCriteria returns the Criteria field value if set, zero value otherwise.
func (o *SearchCertificatesRestRequestV2) GetCriteria() []SearchCertificateCriteriaRestRequest {
	if o == nil || isNil(o.Criteria) {
		var ret []SearchCertificateCriteriaRestRequest
		return ret
	}
	return o.Criteria
}

// GetCriteriaOk returns a tuple with the Criteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificatesRestRequestV2) GetCriteriaOk() ([]SearchCertificateCriteriaRestRequest, bool) {
	if o == nil || isNil(o.Criteria) {
		return nil, false
	}
	return o.Criteria, true
}

// HasCriteria returns a boolean if a field has been set.
func (o *SearchCertificatesRestRequestV2) HasCriteria() bool {
	if o != nil && !isNil(o.Criteria) {
		return true
	}

	return false
}

// SetCriteria gets a reference to the given []SearchCertificateCriteriaRestRequest and assigns it to the Criteria field.
func (o *SearchCertificatesRestRequestV2) SetCriteria(v []SearchCertificateCriteriaRestRequest) {
	o.Criteria = v
}

func (o SearchCertificatesRestRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCertificatesRestRequestV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !isNil(o.Sort) {
		toSerialize["sort"] = o.Sort
	}
	if !isNil(o.Criteria) {
		toSerialize["criteria"] = o.Criteria
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchCertificatesRestRequestV2) UnmarshalJSON(bytes []byte) (err error) {
	varSearchCertificatesRestRequestV2 := _SearchCertificatesRestRequestV2{}

	if err = json.Unmarshal(bytes, &varSearchCertificatesRestRequestV2); err == nil {
		*o = SearchCertificatesRestRequestV2(varSearchCertificatesRestRequestV2)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination")
		delete(additionalProperties, "sort")
		delete(additionalProperties, "criteria")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchCertificatesRestRequestV2 struct {
	value *SearchCertificatesRestRequestV2
	isSet bool
}

func (v NullableSearchCertificatesRestRequestV2) Get() *SearchCertificatesRestRequestV2 {
	return v.value
}

func (v *NullableSearchCertificatesRestRequestV2) Set(val *SearchCertificatesRestRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCertificatesRestRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCertificatesRestRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCertificatesRestRequestV2(val *SearchCertificatesRestRequestV2) *NullableSearchCertificatesRestRequestV2 {
	return &NullableSearchCertificatesRestRequestV2{value: val, isSet: true}
}

func (v NullableSearchCertificatesRestRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCertificatesRestRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
