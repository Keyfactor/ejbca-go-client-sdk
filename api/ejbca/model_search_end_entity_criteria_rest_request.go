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

// checks if the SearchEndEntityCriteriaRestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchEndEntityCriteriaRestRequest{}

// SearchEndEntityCriteriaRestRequest Use one of allowed values as property(see enum values below). QUERY - multiplicity [0, 1] - is used to search by SubjectDn, SubjectAn, Username;  Available STATUS - multiplicity [0, 9] - values are: NEW, FAILED, INITIALIZED, INPROCESS, GENERATED, REVOKED, HISTORICAL, KEYRECOVERY, WAITINGFORADDAPPROVAL;  END_ENTITY_PROFILE, CERTIFICATE_PROFILE, CA - multiplicity [0, *) - exact match of the name for referencing End Entity Profile, Certificate Profile or CA;  
type SearchEndEntityCriteriaRestRequest struct {
	// A search property
	Property *string `json:"property,omitempty"`
	// A search value. This could be string value, an appropriate string name of End Entity Profile or Certificate Profile or CA
	Value *string `json:"value,omitempty"`
	// An operation for property on inserted value. 'EQUALS' for string, 'LIKE' for string value ('QUERY')
	Operation *string `json:"operation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SearchEndEntityCriteriaRestRequest SearchEndEntityCriteriaRestRequest

// NewSearchEndEntityCriteriaRestRequest instantiates a new SearchEndEntityCriteriaRestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchEndEntityCriteriaRestRequest() *SearchEndEntityCriteriaRestRequest {
	this := SearchEndEntityCriteriaRestRequest{}
	return &this
}

// NewSearchEndEntityCriteriaRestRequestWithDefaults instantiates a new SearchEndEntityCriteriaRestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchEndEntityCriteriaRestRequestWithDefaults() *SearchEndEntityCriteriaRestRequest {
	this := SearchEndEntityCriteriaRestRequest{}
	return &this
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SearchEndEntityCriteriaRestRequest) GetProperty() string {
	if o == nil || isNil(o.Property) {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEndEntityCriteriaRestRequest) GetPropertyOk() (*string, bool) {
	if o == nil || isNil(o.Property) {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SearchEndEntityCriteriaRestRequest) HasProperty() bool {
	if o != nil && !isNil(o.Property) {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SearchEndEntityCriteriaRestRequest) SetProperty(v string) {
	o.Property = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SearchEndEntityCriteriaRestRequest) GetValue() string {
	if o == nil || isNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEndEntityCriteriaRestRequest) GetValueOk() (*string, bool) {
	if o == nil || isNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SearchEndEntityCriteriaRestRequest) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SearchEndEntityCriteriaRestRequest) SetValue(v string) {
	o.Value = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *SearchEndEntityCriteriaRestRequest) GetOperation() string {
	if o == nil || isNil(o.Operation) {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchEndEntityCriteriaRestRequest) GetOperationOk() (*string, bool) {
	if o == nil || isNil(o.Operation) {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *SearchEndEntityCriteriaRestRequest) HasOperation() bool {
	if o != nil && !isNil(o.Operation) {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *SearchEndEntityCriteriaRestRequest) SetOperation(v string) {
	o.Operation = &v
}

func (o SearchEndEntityCriteriaRestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchEndEntityCriteriaRestRequest) ToMap() (map[string]interface{}, error) {
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

func (o *SearchEndEntityCriteriaRestRequest) UnmarshalJSON(bytes []byte) (err error) {
	varSearchEndEntityCriteriaRestRequest := _SearchEndEntityCriteriaRestRequest{}

	if err = json.Unmarshal(bytes, &varSearchEndEntityCriteriaRestRequest); err == nil {
		*o = SearchEndEntityCriteriaRestRequest(varSearchEndEntityCriteriaRestRequest)
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

type NullableSearchEndEntityCriteriaRestRequest struct {
	value *SearchEndEntityCriteriaRestRequest
	isSet bool
}

func (v NullableSearchEndEntityCriteriaRestRequest) Get() *SearchEndEntityCriteriaRestRequest {
	return v.value
}

func (v *NullableSearchEndEntityCriteriaRestRequest) Set(val *SearchEndEntityCriteriaRestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchEndEntityCriteriaRestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchEndEntityCriteriaRestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchEndEntityCriteriaRestRequest(val *SearchEndEntityCriteriaRestRequest) *NullableSearchEndEntityCriteriaRestRequest {
	return &NullableSearchEndEntityCriteriaRestRequest{value: val, isSet: true}
}

func (v NullableSearchEndEntityCriteriaRestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchEndEntityCriteriaRestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


