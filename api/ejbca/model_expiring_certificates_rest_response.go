/*
EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"encoding/json"
)

// checks if the ExpiringCertificatesRestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExpiringCertificatesRestResponse{}

// ExpiringCertificatesRestResponse struct for ExpiringCertificatesRestResponse
type ExpiringCertificatesRestResponse struct {
	PaginationRestResponseComponent *PaginationRestResponseComponent `json:"pagination_rest_response_component,omitempty"`
	CertificatesRestResponse *CertificatesRestResponse `json:"certificates_rest_response,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExpiringCertificatesRestResponse ExpiringCertificatesRestResponse

// NewExpiringCertificatesRestResponse instantiates a new ExpiringCertificatesRestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpiringCertificatesRestResponse() *ExpiringCertificatesRestResponse {
	this := ExpiringCertificatesRestResponse{}
	return &this
}

// NewExpiringCertificatesRestResponseWithDefaults instantiates a new ExpiringCertificatesRestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpiringCertificatesRestResponseWithDefaults() *ExpiringCertificatesRestResponse {
	this := ExpiringCertificatesRestResponse{}
	return &this
}

// GetPaginationRestResponseComponent returns the PaginationRestResponseComponent field value if set, zero value otherwise.
func (o *ExpiringCertificatesRestResponse) GetPaginationRestResponseComponent() PaginationRestResponseComponent {
	if o == nil || isNil(o.PaginationRestResponseComponent) {
		var ret PaginationRestResponseComponent
		return ret
	}
	return *o.PaginationRestResponseComponent
}

// GetPaginationRestResponseComponentOk returns a tuple with the PaginationRestResponseComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiringCertificatesRestResponse) GetPaginationRestResponseComponentOk() (*PaginationRestResponseComponent, bool) {
	if o == nil || isNil(o.PaginationRestResponseComponent) {
		return nil, false
	}
	return o.PaginationRestResponseComponent, true
}

// HasPaginationRestResponseComponent returns a boolean if a field has been set.
func (o *ExpiringCertificatesRestResponse) HasPaginationRestResponseComponent() bool {
	if o != nil && !isNil(o.PaginationRestResponseComponent) {
		return true
	}

	return false
}

// SetPaginationRestResponseComponent gets a reference to the given PaginationRestResponseComponent and assigns it to the PaginationRestResponseComponent field.
func (o *ExpiringCertificatesRestResponse) SetPaginationRestResponseComponent(v PaginationRestResponseComponent) {
	o.PaginationRestResponseComponent = &v
}

// GetCertificatesRestResponse returns the CertificatesRestResponse field value if set, zero value otherwise.
func (o *ExpiringCertificatesRestResponse) GetCertificatesRestResponse() CertificatesRestResponse {
	if o == nil || isNil(o.CertificatesRestResponse) {
		var ret CertificatesRestResponse
		return ret
	}
	return *o.CertificatesRestResponse
}

// GetCertificatesRestResponseOk returns a tuple with the CertificatesRestResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpiringCertificatesRestResponse) GetCertificatesRestResponseOk() (*CertificatesRestResponse, bool) {
	if o == nil || isNil(o.CertificatesRestResponse) {
		return nil, false
	}
	return o.CertificatesRestResponse, true
}

// HasCertificatesRestResponse returns a boolean if a field has been set.
func (o *ExpiringCertificatesRestResponse) HasCertificatesRestResponse() bool {
	if o != nil && !isNil(o.CertificatesRestResponse) {
		return true
	}

	return false
}

// SetCertificatesRestResponse gets a reference to the given CertificatesRestResponse and assigns it to the CertificatesRestResponse field.
func (o *ExpiringCertificatesRestResponse) SetCertificatesRestResponse(v CertificatesRestResponse) {
	o.CertificatesRestResponse = &v
}

func (o ExpiringCertificatesRestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExpiringCertificatesRestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PaginationRestResponseComponent) {
		toSerialize["pagination_rest_response_component"] = o.PaginationRestResponseComponent
	}
	if !isNil(o.CertificatesRestResponse) {
		toSerialize["certificates_rest_response"] = o.CertificatesRestResponse
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExpiringCertificatesRestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varExpiringCertificatesRestResponse := _ExpiringCertificatesRestResponse{}

	if err = json.Unmarshal(bytes, &varExpiringCertificatesRestResponse); err == nil {
		*o = ExpiringCertificatesRestResponse(varExpiringCertificatesRestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "pagination_rest_response_component")
		delete(additionalProperties, "certificates_rest_response")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExpiringCertificatesRestResponse struct {
	value *ExpiringCertificatesRestResponse
	isSet bool
}

func (v NullableExpiringCertificatesRestResponse) Get() *ExpiringCertificatesRestResponse {
	return v.value
}

func (v *NullableExpiringCertificatesRestResponse) Set(val *ExpiringCertificatesRestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExpiringCertificatesRestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExpiringCertificatesRestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpiringCertificatesRestResponse(val *ExpiringCertificatesRestResponse) *NullableExpiringCertificatesRestResponse {
	return &NullableExpiringCertificatesRestResponse{value: val, isSet: true}
}

func (v NullableExpiringCertificatesRestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpiringCertificatesRestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


