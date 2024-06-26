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

// checks if the CrlRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CrlRestResponse{}

// CrlRestResponse struct for CrlRestResponse
type CrlRestResponse struct {
	// Certificate Revokation List (CRL)
	Crl *string `json:"crl,omitempty"`
	// Response format
	ResponseFormat       *string `json:"response_format,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CrlRestResponse CrlRestResponse

// NewCrlRestResponse instantiates a new CrlRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrlRestResponse() *CrlRestResponse {
	this := CrlRestResponse{}
	return &this
}

// NewCrlRestResponseWithDefaults instantiates a new CrlRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrlRestResponseWithDefaults() *CrlRestResponse {
	this := CrlRestResponse{}
	return &this
}

// GetCrl returns the Crl field value if set, zero value otherwise.
func (o *CrlRestResponse) GetCrl() string {
	if o == nil || isNil(o.Crl) {
		var ret string
		return ret
	}
	return *o.Crl
}

// GetCrlOk returns a tuple with the Crl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrlRestResponse) GetCrlOk() (*string, bool) {
	if o == nil || isNil(o.Crl) {
		return nil, false
	}
	return o.Crl, true
}

// HasCrl returns a boolean if a field has been set.
func (o *CrlRestResponse) HasCrl() bool {
	if o != nil && !isNil(o.Crl) {
		return true
	}

	return false
}

// SetCrl gets a reference to the given string and assigns it to the Crl field.
func (o *CrlRestResponse) SetCrl(v string) {
	o.Crl = &v
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *CrlRestResponse) GetResponseFormat() string {
	if o == nil || isNil(o.ResponseFormat) {
		var ret string
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrlRestResponse) GetResponseFormatOk() (*string, bool) {
	if o == nil || isNil(o.ResponseFormat) {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *CrlRestResponse) HasResponseFormat() bool {
	if o != nil && !isNil(o.ResponseFormat) {
		return true
	}

	return false
}

// SetResponseFormat gets a reference to the given string and assigns it to the ResponseFormat field.
func (o *CrlRestResponse) SetResponseFormat(v string) {
	o.ResponseFormat = &v
}

func (o CrlRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CrlRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Crl) {
		toSerialize["crl"] = o.Crl
	}
	if !isNil(o.ResponseFormat) {
		toSerialize["response_format"] = o.ResponseFormat
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CrlRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varCrlRestResponse := _CrlRestResponse{}

	if err = json.Unmarshal(bytes, &varCrlRestResponse); err == nil {
		*o = CrlRestResponse(varCrlRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "crl")
		delete(additionalProperties, "response_format")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCrlRestResponse struct {
	value *CrlRestResponse
	isSet bool
}

func (v NullableCrlRestResponse) Get() *CrlRestResponse {
	return v.value
}

func (v *NullableCrlRestResponse) Set(val *CrlRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCrlRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCrlRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrlRestResponse(val *CrlRestResponse) *NullableCrlRestResponse {
	return &NullableCrlRestResponse{value: val, isSet: true}
}

func (v NullableCrlRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrlRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
