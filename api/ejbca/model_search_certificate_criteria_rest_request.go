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

// checks if the SearchCertificateCriteriaRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchCertificateCriteriaRestRequest{}

// SearchCertificateCriteriaRestRequest Use one of allowed values as property(see enum values below). QUERY - multiplicity [0, 1] - is used to search by SubjectDn, SubjectAn, Username or SerialNr;  Available STATUS - multiplicity [0, 12] - values are: CERT_ACTIVE, CERT_REVOKED, REVOCATION_REASON_UNSPECIFIED, REVOCATION_REASON_KEYCOMPROMISE, REVOCATION_REASON_CACOMPROMISE, REVOCATION_REASON_AFFILIATIONCHANGED, REVOCATION_REASON_SUPERSEDED, REVOCATION_REASON_CESSATIONOFOPERATION, REVOCATION_REASON_CERTIFICATEHOLD, REVOCATION_REASON_REMOVEFROMCRL, REVOCATION_REASON_PRIVILEGESWITHDRAWN, REVOCATION_REASON_AACOMPROMISE;  END_ENTITY_PROFILE, CERTIFICATE_PROFILE, CA - multiplicity [0, *) - exact match of the name for referencing End Entity Profile, Certificate Profile or CA;  ISSUED_DATE 'BEFORE' - multiplicity [0, 1] - ISO 8601 Date string;  ISSUED_DATE 'AFTER' - multiplicity [0, 1] - ISO 8601 Date string;  EXPIRE_DATE 'BEFORE' - multiplicity [0, 1] - ISO 8601 Date string;  EXPIRE_DATE 'AFTER' - multiplicity [0, 1] - ISO 8601 Date string;  REVOCATION_DATE 'BEFORE' - multiplicity [0, 1] - ISO 8601 Date string;  REVOCATION_DATE 'AFTER' - multiplicity [0, 1] - ISO 8601 Date string.  UPDATE_TIME 'BEFORE' - multiplicity [0, 1] - ISO 8601 Date string;  UPDATE_TIME 'AFTER' - multiplicity [0, 1] - ISO 8601 Date string;
type SearchCertificateCriteriaRestRequest struct {
	// A search property
	Property *string `json:"property,omitempty"`
	// A search value. This could be sting value, ISO 8601 Date string, an appropriate string name of End Entity Profile or Certificate Profile or CA
	Value *string `json:"value,omitempty"`
	// An operation for property on inserted value. 'EQUAL' for string, 'LIKE' for string value ('QUERY'), 'BEFORE' or 'AFTER' for date values
	Operation            *string `json:"operation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchCertificateCriteriaRestRequest SearchCertificateCriteriaRestRequest

// NewSearchCertificateCriteriaRestRequest instantiates a new SearchCertificateCriteriaRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchCertificateCriteriaRestRequest() *SearchCertificateCriteriaRestRequest {
	this := SearchCertificateCriteriaRestRequest{}
	return &this
}

// NewSearchCertificateCriteriaRestRequestWithDefaults instantiates a new SearchCertificateCriteriaRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchCertificateCriteriaRestRequestWithDefaults() *SearchCertificateCriteriaRestRequest {
	this := SearchCertificateCriteriaRestRequest{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SearchCertificateCriteriaRestRequest) GetProperty() string {
	if o == nil || isNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificateCriteriaRestRequest) GetPropertyOk() (*string, bool) {
	if o == nil || isNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SearchCertificateCriteriaRestRequest) HasProperty() bool {
	if o != nil && !isNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SearchCertificateCriteriaRestRequest) SetProperty(v string) {
	o.Property = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SearchCertificateCriteriaRestRequest) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificateCriteriaRestRequest) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SearchCertificateCriteriaRestRequest) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SearchCertificateCriteriaRestRequest) SetValue(v string) {
	o.Value = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SearchCertificateCriteriaRestRequest) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchCertificateCriteriaRestRequest) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SearchCertificateCriteriaRestRequest) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *SearchCertificateCriteriaRestRequest) SetOperation(v string) {
	o.Operation = &v
}

func (o SearchCertificateCriteriaRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchCertificateCriteriaRestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Property) {
		toSerialize["property"] = o.Property
	}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !isNil(o.Operation) {
		toSerialize["operation"] = o.Operation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *SearchCertificateCriteriaRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSearchCertificateCriteriaRestRequest := _SearchCertificateCriteriaRestRequest{}

	if err = json.Unmarshal(bytes, &varSearchCertificateCriteriaRestRequest); err == nil {
		*o = SearchCertificateCriteriaRestRequest(varSearchCertificateCriteriaRestRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "property")
		delete(additionalProperties, "value")
		delete(additionalProperties, "operation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSearchCertificateCriteriaRestRequest struct {
	value *SearchCertificateCriteriaRestRequest
	isSet bool
}

func (v NullableSearchCertificateCriteriaRestRequest) Get() *SearchCertificateCriteriaRestRequest {
	return v.value
}

func (v *NullableSearchCertificateCriteriaRestRequest) Set(val *SearchCertificateCriteriaRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchCertificateCriteriaRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchCertificateCriteriaRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchCertificateCriteriaRestRequest(val *SearchCertificateCriteriaRestRequest) *NullableSearchCertificateCriteriaRestRequest {
	return &NullableSearchCertificateCriteriaRestRequest{value: val, isSet: true}
}

func (v NullableSearchCertificateCriteriaRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchCertificateCriteriaRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
