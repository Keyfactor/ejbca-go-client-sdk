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

// checks if the FinalizeRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinalizeRestRequest{}

// FinalizeRestRequest struct for FinalizeRestRequest
type FinalizeRestRequest struct {
	// Response format
	ResponseFormat *string `json:"response_format,omitempty"`
	// Password
	Password *string `json:"password,omitempty"`
	// Key algorithm
	KeyAlg *string `json:"key_alg,omitempty"`
	// Key specification
	KeySpec              *string `json:"key_spec,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FinalizeRestRequest FinalizeRestRequest

// NewFinalizeRestRequest instantiates a new FinalizeRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinalizeRestRequest() *FinalizeRestRequest {
	this := FinalizeRestRequest{}
	return &this
}

// NewFinalizeRestRequestWithDefaults instantiates a new FinalizeRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinalizeRestRequestWithDefaults() *FinalizeRestRequest {
	this := FinalizeRestRequest{}
	return &this
}

// GetResponseFormat returns the ResponseFormat field value if set, zero value otherwise.
func (o *FinalizeRestRequest) GetResponseFormat() string {
	if o == nil || isNil(o.ResponseFormat) {
		var ret string
		return ret
	}
	return *o.ResponseFormat
}

// GetResponseFormatOk returns a tuple with the ResponseFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalizeRestRequest) GetResponseFormatOk() (*string, bool) {
	if o == nil || isNil(o.ResponseFormat) {
		return nil, false
	}
	return o.ResponseFormat, true
}

// HasResponseFormat returns a boolean if a field has been set.
func (o *FinalizeRestRequest) HasResponseFormat() bool {
	if o != nil && !isNil(o.ResponseFormat) {
		return true
	}

	return false
}

// SetResponseFormat gets a reference to the given string and assigns it to the ResponseFormat field.
func (o *FinalizeRestRequest) SetResponseFormat(v string) {
	o.ResponseFormat = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *FinalizeRestRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalizeRestRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *FinalizeRestRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *FinalizeRestRequest) SetPassword(v string) {
	o.Password = &v
}

// GetKeyAlg returns the KeyAlg field value if set, zero value otherwise.
func (o *FinalizeRestRequest) GetKeyAlg() string {
	if o == nil || isNil(o.KeyAlg) {
		var ret string
		return ret
	}
	return *o.KeyAlg
}

// GetKeyAlgOk returns a tuple with the KeyAlg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalizeRestRequest) GetKeyAlgOk() (*string, bool) {
	if o == nil || isNil(o.KeyAlg) {
		return nil, false
	}
	return o.KeyAlg, true
}

// HasKeyAlg returns a boolean if a field has been set.
func (o *FinalizeRestRequest) HasKeyAlg() bool {
	if o != nil && !isNil(o.KeyAlg) {
		return true
	}

	return false
}

// SetKeyAlg gets a reference to the given string and assigns it to the KeyAlg field.
func (o *FinalizeRestRequest) SetKeyAlg(v string) {
	o.KeyAlg = &v
}

// GetKeySpec returns the KeySpec field value if set, zero value otherwise.
func (o *FinalizeRestRequest) GetKeySpec() string {
	if o == nil || isNil(o.KeySpec) {
		var ret string
		return ret
	}
	return *o.KeySpec
}

// GetKeySpecOk returns a tuple with the KeySpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinalizeRestRequest) GetKeySpecOk() (*string, bool) {
	if o == nil || isNil(o.KeySpec) {
		return nil, false
	}
	return o.KeySpec, true
}

// HasKeySpec returns a boolean if a field has been set.
func (o *FinalizeRestRequest) HasKeySpec() bool {
	if o != nil && !isNil(o.KeySpec) {
		return true
	}

	return false
}

// SetKeySpec gets a reference to the given string and assigns it to the KeySpec field.
func (o *FinalizeRestRequest) SetKeySpec(v string) {
	o.KeySpec = &v
}

func (o FinalizeRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinalizeRestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResponseFormat) {
		toSerialize["response_format"] = o.ResponseFormat
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.KeyAlg) {
		toSerialize["key_alg"] = o.KeyAlg
	}
	if !isNil(o.KeySpec) {
		toSerialize["key_spec"] = o.KeySpec
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FinalizeRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varFinalizeRestRequest := _FinalizeRestRequest{}

	if err = json.Unmarshal(bytes, &varFinalizeRestRequest); err == nil {
		*o = FinalizeRestRequest(varFinalizeRestRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "response_format")
		delete(additionalProperties, "password")
		delete(additionalProperties, "key_alg")
		delete(additionalProperties, "key_spec")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFinalizeRestRequest struct {
	value *FinalizeRestRequest
	isSet bool
}

func (v NullableFinalizeRestRequest) Get() *FinalizeRestRequest {
	return v.value
}

func (v *NullableFinalizeRestRequest) Set(val *FinalizeRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFinalizeRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFinalizeRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinalizeRestRequest(val *FinalizeRestRequest) *NullableFinalizeRestRequest {
	return &NullableFinalizeRestRequest{value: val, isSet: true}
}

func (v NullableFinalizeRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinalizeRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
