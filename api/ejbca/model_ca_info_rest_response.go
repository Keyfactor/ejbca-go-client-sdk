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
	"time"
)

// checks if the CaInfoRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CaInfoRestResponse{}

// CaInfoRestResponse struct for CaInfoRestResponse
type CaInfoRestResponse struct {
	// CA identifier
	Id *int32 `json:"id,omitempty"`
	// Certificate Authority (CA) name
	Name *string `json:"name,omitempty"`
	// Subject Distinguished Name
	SubjectDn *string `json:"subject_dn,omitempty"`
	// Issuer Distinguished Name
	IssuerDn *string `json:"issuer_dn,omitempty"`
	// Expiration date
	ExpirationDate       *time.Time `json:"expiration_date,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CaInfoRestResponse CaInfoRestResponse

// NewCaInfoRestResponse instantiates a new CaInfoRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCaInfoRestResponse() *CaInfoRestResponse {
	this := CaInfoRestResponse{}
	return &this
}

// NewCaInfoRestResponseWithDefaults instantiates a new CaInfoRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCaInfoRestResponseWithDefaults() *CaInfoRestResponse {
	this := CaInfoRestResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CaInfoRestResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaInfoRestResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CaInfoRestResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CaInfoRestResponse) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CaInfoRestResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaInfoRestResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CaInfoRestResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CaInfoRestResponse) SetName(v string) {
	o.Name = &v
}

// GetSubjectDn returns the SubjectDn field value if set, zero value otherwise.
func (o *CaInfoRestResponse) GetSubjectDn() string {
	if o == nil || isNil(o.SubjectDn) {
		var ret string
		return ret
	}
	return *o.SubjectDn
}

// GetSubjectDnOk returns a tuple with the SubjectDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaInfoRestResponse) GetSubjectDnOk() (*string, bool) {
	if o == nil || isNil(o.SubjectDn) {
		return nil, false
	}
	return o.SubjectDn, true
}

// HasSubjectDn returns a boolean if a field has been set.
func (o *CaInfoRestResponse) HasSubjectDn() bool {
	if o != nil && !isNil(o.SubjectDn) {
		return true
	}

	return false
}

// SetSubjectDn gets a reference to the given string and assigns it to the SubjectDn field.
func (o *CaInfoRestResponse) SetSubjectDn(v string) {
	o.SubjectDn = &v
}

// GetIssuerDn returns the IssuerDn field value if set, zero value otherwise.
func (o *CaInfoRestResponse) GetIssuerDn() string {
	if o == nil || isNil(o.IssuerDn) {
		var ret string
		return ret
	}
	return *o.IssuerDn
}

// GetIssuerDnOk returns a tuple with the IssuerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaInfoRestResponse) GetIssuerDnOk() (*string, bool) {
	if o == nil || isNil(o.IssuerDn) {
		return nil, false
	}
	return o.IssuerDn, true
}

// HasIssuerDn returns a boolean if a field has been set.
func (o *CaInfoRestResponse) HasIssuerDn() bool {
	if o != nil && !isNil(o.IssuerDn) {
		return true
	}

	return false
}

// SetIssuerDn gets a reference to the given string and assigns it to the IssuerDn field.
func (o *CaInfoRestResponse) SetIssuerDn(v string) {
	o.IssuerDn = &v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *CaInfoRestResponse) GetExpirationDate() time.Time {
	if o == nil || isNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CaInfoRestResponse) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *CaInfoRestResponse) HasExpirationDate() bool {
	if o != nil && !isNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *CaInfoRestResponse) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

func (o CaInfoRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CaInfoRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.SubjectDn) {
		toSerialize["subject_dn"] = o.SubjectDn
	}
	if !isNil(o.IssuerDn) {
		toSerialize["issuer_dn"] = o.IssuerDn
	}
	if !isNil(o.ExpirationDate) {
		toSerialize["expiration_date"] = o.ExpirationDate
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CaInfoRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCaInfoRestResponse := _CaInfoRestResponse{}

	if err = json.Unmarshal(bytes, &varCaInfoRestResponse); err == nil {
		*o = CaInfoRestResponse(varCaInfoRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "subject_dn")
		delete(additionalProperties, "issuer_dn")
		delete(additionalProperties, "expiration_date")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCaInfoRestResponse struct {
	value *CaInfoRestResponse
	isSet bool
}

func (v NullableCaInfoRestResponse) Get() *CaInfoRestResponse {
	return v.value
}

func (v *NullableCaInfoRestResponse) Set(val *CaInfoRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCaInfoRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCaInfoRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCaInfoRestResponse(val *CaInfoRestResponse) *NullableCaInfoRestResponse {
	return &NullableCaInfoRestResponse{value: val, isSet: true}
}

func (v NullableCaInfoRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCaInfoRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
