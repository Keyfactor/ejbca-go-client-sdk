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

// checks if the AuthorizedEEPsRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizedEEPsRestResponse{}

// AuthorizedEEPsRestResponse struct for AuthorizedEEPsRestResponse
type AuthorizedEEPsRestResponse struct {
	EndEntitieProfiles   []EndEntityProfileRestResponse `json:"end_entitie_profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AuthorizedEEPsRestResponse AuthorizedEEPsRestResponse

// NewAuthorizedEEPsRestResponse instantiates a new AuthorizedEEPsRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizedEEPsRestResponse() *AuthorizedEEPsRestResponse {
	this := AuthorizedEEPsRestResponse{}
	return &this
}

// NewAuthorizedEEPsRestResponseWithDefaults instantiates a new AuthorizedEEPsRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizedEEPsRestResponseWithDefaults() *AuthorizedEEPsRestResponse {
	this := AuthorizedEEPsRestResponse{}
	return &this
}

// GetEndEntitieProfiles returns the EndEntitieProfiles field value if set, zero value otherwise.
func (o *AuthorizedEEPsRestResponse) GetEndEntitieProfiles() []EndEntityProfileRestResponse {
	if o == nil || isNil(o.EndEntitieProfiles) {
		var ret []EndEntityProfileRestResponse
		return ret
	}
	return o.EndEntitieProfiles
}

// GetEndEntitieProfilesOk returns a tuple with the EndEntitieProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizedEEPsRestResponse) GetEndEntitieProfilesOk() ([]EndEntityProfileRestResponse, bool) {
	if o == nil || isNil(o.EndEntitieProfiles) {
		return nil, false
	}
	return o.EndEntitieProfiles, true
}

// HasEndEntitieProfiles returns a boolean if a field has been set.
func (o *AuthorizedEEPsRestResponse) HasEndEntitieProfiles() bool {
	if o != nil && !isNil(o.EndEntitieProfiles) {
		return true
	}

	return false
}

// SetEndEntitieProfiles gets a reference to the given []EndEntityProfileRestResponse and assigns it to the EndEntitieProfiles field.
func (o *AuthorizedEEPsRestResponse) SetEndEntitieProfiles(v []EndEntityProfileRestResponse) {
	o.EndEntitieProfiles = v
}

func (o AuthorizedEEPsRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizedEEPsRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EndEntitieProfiles) {
		toSerialize["end_entitie_profiles"] = o.EndEntitieProfiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AuthorizedEEPsRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAuthorizedEEPsRestResponse := _AuthorizedEEPsRestResponse{}

	if err = json.Unmarshal(bytes, &varAuthorizedEEPsRestResponse); err == nil {
		*o = AuthorizedEEPsRestResponse(varAuthorizedEEPsRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "end_entitie_profiles")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAuthorizedEEPsRestResponse struct {
	value *AuthorizedEEPsRestResponse
	isSet bool
}

func (v NullableAuthorizedEEPsRestResponse) Get() *AuthorizedEEPsRestResponse {
	return v.value
}

func (v *NullableAuthorizedEEPsRestResponse) Set(val *AuthorizedEEPsRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizedEEPsRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizedEEPsRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizedEEPsRestResponse(val *AuthorizedEEPsRestResponse) *NullableAuthorizedEEPsRestResponse {
	return &NullableAuthorizedEEPsRestResponse{value: val, isSet: true}
}

func (v NullableAuthorizedEEPsRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizedEEPsRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
