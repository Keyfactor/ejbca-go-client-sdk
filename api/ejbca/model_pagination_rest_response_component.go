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

// checks if the PaginationRestResponseComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationRestResponseComponent{}

// PaginationRestResponseComponent struct for PaginationRestResponseComponent
type PaginationRestResponseComponent struct {
	MoreResults *bool `json:"more_results,omitempty"`
	NextOffset *int32 `json:"next_offset,omitempty"`
	NumberOfResults *int32 `json:"number_of_results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaginationRestResponseComponent PaginationRestResponseComponent

// NewPaginationRestResponseComponent instantiates a new PaginationRestResponseComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationRestResponseComponent() *PaginationRestResponseComponent {
	this := PaginationRestResponseComponent{}
	return &this
}

// NewPaginationRestResponseComponentWithDefaults instantiates a new PaginationRestResponseComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationRestResponseComponentWithDefaults() *PaginationRestResponseComponent {
	this := PaginationRestResponseComponent{}
	return &this
}

// GetMoreResults returns the MoreResults field value if set, zero value otherwise.
func (o *PaginationRestResponseComponent) GetMoreResults() bool {
	if o == nil || isNil(o.MoreResults) {
		var ret bool
		return ret
	}
	return *o.MoreResults
}

// GetMoreResultsOk returns a tuple with the MoreResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationRestResponseComponent) GetMoreResultsOk() (*bool, bool) {
	if o == nil || isNil(o.MoreResults) {
		return nil, false
	}
	return o.MoreResults, true
}

// HasMoreResults returns a boolean if a field has been set.
func (o *PaginationRestResponseComponent) HasMoreResults() bool {
	if o != nil && !isNil(o.MoreResults) {
		return true
	}

	return false
}

// SetMoreResults gets a reference to the given bool and assigns it to the MoreResults field.
func (o *PaginationRestResponseComponent) SetMoreResults(v bool) {
	o.MoreResults = &v
}

// GetNextOffset returns the NextOffset field value if set, zero value otherwise.
func (o *PaginationRestResponseComponent) GetNextOffset() int32 {
	if o == nil || isNil(o.NextOffset) {
		var ret int32
		return ret
	}
	return *o.NextOffset
}

// GetNextOffsetOk returns a tuple with the NextOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationRestResponseComponent) GetNextOffsetOk() (*int32, bool) {
	if o == nil || isNil(o.NextOffset) {
		return nil, false
	}
	return o.NextOffset, true
}

// HasNextOffset returns a boolean if a field has been set.
func (o *PaginationRestResponseComponent) HasNextOffset() bool {
	if o != nil && !isNil(o.NextOffset) {
		return true
	}

	return false
}

// SetNextOffset gets a reference to the given int32 and assigns it to the NextOffset field.
func (o *PaginationRestResponseComponent) SetNextOffset(v int32) {
	o.NextOffset = &v
}

// GetNumberOfResults returns the NumberOfResults field value if set, zero value otherwise.
func (o *PaginationRestResponseComponent) GetNumberOfResults() int32 {
	if o == nil || isNil(o.NumberOfResults) {
		var ret int32
		return ret
	}
	return *o.NumberOfResults
}

// GetNumberOfResultsOk returns a tuple with the NumberOfResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationRestResponseComponent) GetNumberOfResultsOk() (*int32, bool) {
	if o == nil || isNil(o.NumberOfResults) {
		return nil, false
	}
	return o.NumberOfResults, true
}

// HasNumberOfResults returns a boolean if a field has been set.
func (o *PaginationRestResponseComponent) HasNumberOfResults() bool {
	if o != nil && !isNil(o.NumberOfResults) {
		return true
	}

	return false
}

// SetNumberOfResults gets a reference to the given int32 and assigns it to the NumberOfResults field.
func (o *PaginationRestResponseComponent) SetNumberOfResults(v int32) {
	o.NumberOfResults = &v
}

func (o PaginationRestResponseComponent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationRestResponseComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MoreResults) {
		toSerialize["more_results"] = o.MoreResults
	}
	if !isNil(o.NextOffset) {
		toSerialize["next_offset"] = o.NextOffset
	}
	if !isNil(o.NumberOfResults) {
		toSerialize["number_of_results"] = o.NumberOfResults
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PaginationRestResponseComponent) UnmarshalJSON(bytes []byte) (err error) {
	varPaginationRestResponseComponent := _PaginationRestResponseComponent{}

	if err = json.Unmarshal(bytes, &varPaginationRestResponseComponent); err == nil {
		*o = PaginationRestResponseComponent(varPaginationRestResponseComponent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "more_results")
		delete(additionalProperties, "next_offset")
		delete(additionalProperties, "number_of_results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaginationRestResponseComponent struct {
	value *PaginationRestResponseComponent
	isSet bool
}

func (v NullablePaginationRestResponseComponent) Get() *PaginationRestResponseComponent {
	return v.value
}

func (v *NullablePaginationRestResponseComponent) Set(val *PaginationRestResponseComponent) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationRestResponseComponent) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationRestResponseComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationRestResponseComponent(val *PaginationRestResponseComponent) *NullablePaginationRestResponseComponent {
	return &NullablePaginationRestResponseComponent{value: val, isSet: true}
}

func (v NullablePaginationRestResponseComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationRestResponseComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


