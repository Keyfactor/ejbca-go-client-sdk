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

// checks if the CertificateRestResponseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CertificateRestResponseV2{}

// CertificateRestResponseV2 struct for CertificateRestResponseV2
type CertificateRestResponseV2 struct {
	// Certificate fingerprint
	Fingerprint *string `json:"fingerprint,omitempty"`
	// Certificate Authority fingerprint
	CAFingerprint *string `json:"cAFingerprint,omitempty"`
	// Certificate Profile Identifier
	CertificateProfileId *int32 `json:"certificateProfileId,omitempty"`
	// End Entity Profile Identifier
	EndEntityProfileId *int32 `json:"endEntityProfileId,omitempty"`
	// Date after which certificate should be considered expired
	ExpireDate *int64 `json:"expireDate,omitempty"`
	// Issuer Distinguished Name
	IssuerDN *string `json:"issuerDN,omitempty"`
	// Date at which certificate became valid
	NotBefore *int64 `json:"notBefore,omitempty"`
	// Revocation date
	RevocationDate *int64 `json:"revocationDate,omitempty"`
	// Revocation reson
	RevocationReason *int32 `json:"revocationReason,omitempty"`
	// Hex Serial Number
	SerialNumber *string `json:"serialNumber,omitempty"`
	// Certificate status
	Status *int32 `json:"status,omitempty"`
	// Subject Alternative Name (SAN)
	SubjectAltName *string `json:"subjectAltName,omitempty"`
	// Subject Distinguished Name
	SubjectDN *string `json:"subjectDN,omitempty"`
	// Subject Key Identifier
	SubjectKeyId *string `json:"subjectKeyId,omitempty"`
	Tag *string `json:"tag,omitempty"`
	Type *int32 `json:"type,omitempty"`
	// Update time
	UdpateTime *int64 `json:"udpateTime,omitempty"`
	// Username
	Username *string `json:"username,omitempty"`
	// Base64 encoded certificate
	Base64Cert *string `json:"base64Cert,omitempty"`
	// Certificate request
	CertificateRequest *string `json:"certificateRequest,omitempty"`
	// CRL partition index
	CrlPartitionIndex *int32 `json:"crlPartitionIndex,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CertificateRestResponseV2 CertificateRestResponseV2

// NewCertificateRestResponseV2 instantiates a new CertificateRestResponseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateRestResponseV2() *CertificateRestResponseV2 {
	this := CertificateRestResponseV2{}
	return &this
}

// NewCertificateRestResponseV2WithDefaults instantiates a new CertificateRestResponseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateRestResponseV2WithDefaults() *CertificateRestResponseV2 {
	this := CertificateRestResponseV2{}
	return &this
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetFingerprint() string {
	if o == nil || isNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetFingerprintOk() (*string, bool) {
	if o == nil || isNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasFingerprint() bool {
	if o != nil && !isNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *CertificateRestResponseV2) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetCAFingerprint returns the CAFingerprint field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetCAFingerprint() string {
	if o == nil || isNil(o.CAFingerprint) {
		var ret string
		return ret
	}
	return *o.CAFingerprint
}

// GetCAFingerprintOk returns a tuple with the CAFingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetCAFingerprintOk() (*string, bool) {
	if o == nil || isNil(o.CAFingerprint) {
		return nil, false
	}
	return o.CAFingerprint, true
}

// HasCAFingerprint returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasCAFingerprint() bool {
	if o != nil && !isNil(o.CAFingerprint) {
		return true
	}

	return false
}

// SetCAFingerprint gets a reference to the given string and assigns it to the CAFingerprint field.
func (o *CertificateRestResponseV2) SetCAFingerprint(v string) {
	o.CAFingerprint = &v
}

// GetCertificateProfileId returns the CertificateProfileId field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetCertificateProfileId() int32 {
	if o == nil || isNil(o.CertificateProfileId) {
		var ret int32
		return ret
	}
	return *o.CertificateProfileId
}

// GetCertificateProfileIdOk returns a tuple with the CertificateProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetCertificateProfileIdOk() (*int32, bool) {
	if o == nil || isNil(o.CertificateProfileId) {
		return nil, false
	}
	return o.CertificateProfileId, true
}

// HasCertificateProfileId returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasCertificateProfileId() bool {
	if o != nil && !isNil(o.CertificateProfileId) {
		return true
	}

	return false
}

// SetCertificateProfileId gets a reference to the given int32 and assigns it to the CertificateProfileId field.
func (o *CertificateRestResponseV2) SetCertificateProfileId(v int32) {
	o.CertificateProfileId = &v
}

// GetEndEntityProfileId returns the EndEntityProfileId field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetEndEntityProfileId() int32 {
	if o == nil || isNil(o.EndEntityProfileId) {
		var ret int32
		return ret
	}
	return *o.EndEntityProfileId
}

// GetEndEntityProfileIdOk returns a tuple with the EndEntityProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetEndEntityProfileIdOk() (*int32, bool) {
	if o == nil || isNil(o.EndEntityProfileId) {
		return nil, false
	}
	return o.EndEntityProfileId, true
}

// HasEndEntityProfileId returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasEndEntityProfileId() bool {
	if o != nil && !isNil(o.EndEntityProfileId) {
		return true
	}

	return false
}

// SetEndEntityProfileId gets a reference to the given int32 and assigns it to the EndEntityProfileId field.
func (o *CertificateRestResponseV2) SetEndEntityProfileId(v int32) {
	o.EndEntityProfileId = &v
}

// GetExpireDate returns the ExpireDate field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetExpireDate() int64 {
	if o == nil || isNil(o.ExpireDate) {
		var ret int64
		return ret
	}
	return *o.ExpireDate
}

// GetExpireDateOk returns a tuple with the ExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetExpireDateOk() (*int64, bool) {
	if o == nil || isNil(o.ExpireDate) {
		return nil, false
	}
	return o.ExpireDate, true
}

// HasExpireDate returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasExpireDate() bool {
	if o != nil && !isNil(o.ExpireDate) {
		return true
	}

	return false
}

// SetExpireDate gets a reference to the given int64 and assigns it to the ExpireDate field.
func (o *CertificateRestResponseV2) SetExpireDate(v int64) {
	o.ExpireDate = &v
}

// GetIssuerDN returns the IssuerDN field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetIssuerDN() string {
	if o == nil || isNil(o.IssuerDN) {
		var ret string
		return ret
	}
	return *o.IssuerDN
}

// GetIssuerDNOk returns a tuple with the IssuerDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetIssuerDNOk() (*string, bool) {
	if o == nil || isNil(o.IssuerDN) {
		return nil, false
	}
	return o.IssuerDN, true
}

// HasIssuerDN returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasIssuerDN() bool {
	if o != nil && !isNil(o.IssuerDN) {
		return true
	}

	return false
}

// SetIssuerDN gets a reference to the given string and assigns it to the IssuerDN field.
func (o *CertificateRestResponseV2) SetIssuerDN(v string) {
	o.IssuerDN = &v
}

// GetNotBefore returns the NotBefore field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetNotBefore() int64 {
	if o == nil || isNil(o.NotBefore) {
		var ret int64
		return ret
	}
	return *o.NotBefore
}

// GetNotBeforeOk returns a tuple with the NotBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetNotBeforeOk() (*int64, bool) {
	if o == nil || isNil(o.NotBefore) {
		return nil, false
	}
	return o.NotBefore, true
}

// HasNotBefore returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasNotBefore() bool {
	if o != nil && !isNil(o.NotBefore) {
		return true
	}

	return false
}

// SetNotBefore gets a reference to the given int64 and assigns it to the NotBefore field.
func (o *CertificateRestResponseV2) SetNotBefore(v int64) {
	o.NotBefore = &v
}

// GetRevocationDate returns the RevocationDate field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetRevocationDate() int64 {
	if o == nil || isNil(o.RevocationDate) {
		var ret int64
		return ret
	}
	return *o.RevocationDate
}

// GetRevocationDateOk returns a tuple with the RevocationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetRevocationDateOk() (*int64, bool) {
	if o == nil || isNil(o.RevocationDate) {
		return nil, false
	}
	return o.RevocationDate, true
}

// HasRevocationDate returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasRevocationDate() bool {
	if o != nil && !isNil(o.RevocationDate) {
		return true
	}

	return false
}

// SetRevocationDate gets a reference to the given int64 and assigns it to the RevocationDate field.
func (o *CertificateRestResponseV2) SetRevocationDate(v int64) {
	o.RevocationDate = &v
}

// GetRevocationReason returns the RevocationReason field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetRevocationReason() int32 {
	if o == nil || isNil(o.RevocationReason) {
		var ret int32
		return ret
	}
	return *o.RevocationReason
}

// GetRevocationReasonOk returns a tuple with the RevocationReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetRevocationReasonOk() (*int32, bool) {
	if o == nil || isNil(o.RevocationReason) {
		return nil, false
	}
	return o.RevocationReason, true
}

// HasRevocationReason returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasRevocationReason() bool {
	if o != nil && !isNil(o.RevocationReason) {
		return true
	}

	return false
}

// SetRevocationReason gets a reference to the given int32 and assigns it to the RevocationReason field.
func (o *CertificateRestResponseV2) SetRevocationReason(v int32) {
	o.RevocationReason = &v
}

// GetSerialNumber returns the SerialNumber field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetSerialNumber() string {
	if o == nil || isNil(o.SerialNumber) {
		var ret string
		return ret
	}
	return *o.SerialNumber
}

// GetSerialNumberOk returns a tuple with the SerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetSerialNumberOk() (*string, bool) {
	if o == nil || isNil(o.SerialNumber) {
		return nil, false
	}
	return o.SerialNumber, true
}

// HasSerialNumber returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasSerialNumber() bool {
	if o != nil && !isNil(o.SerialNumber) {
		return true
	}

	return false
}

// SetSerialNumber gets a reference to the given string and assigns it to the SerialNumber field.
func (o *CertificateRestResponseV2) SetSerialNumber(v string) {
	o.SerialNumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetStatus() int32 {
	if o == nil || isNil(o.Status) {
		var ret int32
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetStatusOk() (*int32, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given int32 and assigns it to the Status field.
func (o *CertificateRestResponseV2) SetStatus(v int32) {
	o.Status = &v
}

// GetSubjectAltName returns the SubjectAltName field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetSubjectAltName() string {
	if o == nil || isNil(o.SubjectAltName) {
		var ret string
		return ret
	}
	return *o.SubjectAltName
}

// GetSubjectAltNameOk returns a tuple with the SubjectAltName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetSubjectAltNameOk() (*string, bool) {
	if o == nil || isNil(o.SubjectAltName) {
		return nil, false
	}
	return o.SubjectAltName, true
}

// HasSubjectAltName returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasSubjectAltName() bool {
	if o != nil && !isNil(o.SubjectAltName) {
		return true
	}

	return false
}

// SetSubjectAltName gets a reference to the given string and assigns it to the SubjectAltName field.
func (o *CertificateRestResponseV2) SetSubjectAltName(v string) {
	o.SubjectAltName = &v
}

// GetSubjectDN returns the SubjectDN field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetSubjectDN() string {
	if o == nil || isNil(o.SubjectDN) {
		var ret string
		return ret
	}
	return *o.SubjectDN
}

// GetSubjectDNOk returns a tuple with the SubjectDN field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetSubjectDNOk() (*string, bool) {
	if o == nil || isNil(o.SubjectDN) {
		return nil, false
	}
	return o.SubjectDN, true
}

// HasSubjectDN returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasSubjectDN() bool {
	if o != nil && !isNil(o.SubjectDN) {
		return true
	}

	return false
}

// SetSubjectDN gets a reference to the given string and assigns it to the SubjectDN field.
func (o *CertificateRestResponseV2) SetSubjectDN(v string) {
	o.SubjectDN = &v
}

// GetSubjectKeyId returns the SubjectKeyId field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetSubjectKeyId() string {
	if o == nil || isNil(o.SubjectKeyId) {
		var ret string
		return ret
	}
	return *o.SubjectKeyId
}

// GetSubjectKeyIdOk returns a tuple with the SubjectKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetSubjectKeyIdOk() (*string, bool) {
	if o == nil || isNil(o.SubjectKeyId) {
		return nil, false
	}
	return o.SubjectKeyId, true
}

// HasSubjectKeyId returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasSubjectKeyId() bool {
	if o != nil && !isNil(o.SubjectKeyId) {
		return true
	}

	return false
}

// SetSubjectKeyId gets a reference to the given string and assigns it to the SubjectKeyId field.
func (o *CertificateRestResponseV2) SetSubjectKeyId(v string) {
	o.SubjectKeyId = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetTag() string {
	if o == nil || isNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetTagOk() (*string, bool) {
	if o == nil || isNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasTag() bool {
	if o != nil && !isNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *CertificateRestResponseV2) SetTag(v string) {
	o.Tag = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetType() int32 {
	if o == nil || isNil(o.Type) {
		var ret int32
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetTypeOk() (*int32, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given int32 and assigns it to the Type field.
func (o *CertificateRestResponseV2) SetType(v int32) {
	o.Type = &v
}

// GetUdpateTime returns the UdpateTime field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetUdpateTime() int64 {
	if o == nil || isNil(o.UdpateTime) {
		var ret int64
		return ret
	}
	return *o.UdpateTime
}

// GetUdpateTimeOk returns a tuple with the UdpateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetUdpateTimeOk() (*int64, bool) {
	if o == nil || isNil(o.UdpateTime) {
		return nil, false
	}
	return o.UdpateTime, true
}

// HasUdpateTime returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasUdpateTime() bool {
	if o != nil && !isNil(o.UdpateTime) {
		return true
	}

	return false
}

// SetUdpateTime gets a reference to the given int64 and assigns it to the UdpateTime field.
func (o *CertificateRestResponseV2) SetUdpateTime(v int64) {
	o.UdpateTime = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *CertificateRestResponseV2) SetUsername(v string) {
	o.Username = &v
}

// GetBase64Cert returns the Base64Cert field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetBase64Cert() string {
	if o == nil || isNil(o.Base64Cert) {
		var ret string
		return ret
	}
	return *o.Base64Cert
}

// GetBase64CertOk returns a tuple with the Base64Cert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetBase64CertOk() (*string, bool) {
	if o == nil || isNil(o.Base64Cert) {
		return nil, false
	}
	return o.Base64Cert, true
}

// HasBase64Cert returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasBase64Cert() bool {
	if o != nil && !isNil(o.Base64Cert) {
		return true
	}

	return false
}

// SetBase64Cert gets a reference to the given string and assigns it to the Base64Cert field.
func (o *CertificateRestResponseV2) SetBase64Cert(v string) {
	o.Base64Cert = &v
}

// GetCertificateRequest returns the CertificateRequest field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetCertificateRequest() string {
	if o == nil || isNil(o.CertificateRequest) {
		var ret string
		return ret
	}
	return *o.CertificateRequest
}

// GetCertificateRequestOk returns a tuple with the CertificateRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetCertificateRequestOk() (*string, bool) {
	if o == nil || isNil(o.CertificateRequest) {
		return nil, false
	}
	return o.CertificateRequest, true
}

// HasCertificateRequest returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasCertificateRequest() bool {
	if o != nil && !isNil(o.CertificateRequest) {
		return true
	}

	return false
}

// SetCertificateRequest gets a reference to the given string and assigns it to the CertificateRequest field.
func (o *CertificateRestResponseV2) SetCertificateRequest(v string) {
	o.CertificateRequest = &v
}

// GetCrlPartitionIndex returns the CrlPartitionIndex field value if set, zero value otherwise.
func (o *CertificateRestResponseV2) GetCrlPartitionIndex() int32 {
	if o == nil || isNil(o.CrlPartitionIndex) {
		var ret int32
		return ret
	}
	return *o.CrlPartitionIndex
}

// GetCrlPartitionIndexOk returns a tuple with the CrlPartitionIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateRestResponseV2) GetCrlPartitionIndexOk() (*int32, bool) {
	if o == nil || isNil(o.CrlPartitionIndex) {
		return nil, false
	}
	return o.CrlPartitionIndex, true
}

// HasCrlPartitionIndex returns a boolean if a field has been set.
func (o *CertificateRestResponseV2) HasCrlPartitionIndex() bool {
	if o != nil && !isNil(o.CrlPartitionIndex) {
		return true
	}

	return false
}

// SetCrlPartitionIndex gets a reference to the given int32 and assigns it to the CrlPartitionIndex field.
func (o *CertificateRestResponseV2) SetCrlPartitionIndex(v int32) {
	o.CrlPartitionIndex = &v
}

func (o CertificateRestResponseV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CertificateRestResponseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	// skip: fingerprint is readOnly
	// skip: cAFingerprint is readOnly
	// skip: certificateProfileId is readOnly
	// skip: endEntityProfileId is readOnly
	// skip: expireDate is readOnly
	// skip: issuerDN is readOnly
	// skip: notBefore is readOnly
	// skip: revocationDate is readOnly
	// skip: revocationReason is readOnly
	// skip: serialNumber is readOnly
	// skip: status is readOnly
	// skip: subjectAltName is readOnly
	// skip: subjectDN is readOnly
	// skip: subjectKeyId is readOnly
	// skip: tag is readOnly
	// skip: type is readOnly
	// skip: udpateTime is readOnly
	// skip: username is readOnly
	// skip: base64Cert is readOnly
	// skip: certificateRequest is readOnly
	// skip: crlPartitionIndex is readOnly

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CertificateRestResponseV2) UnmarshalJSON(bytes []byte) (err error) {
	varCertificateRestResponseV2 := _CertificateRestResponseV2{}

	if err = json.Unmarshal(bytes, &varCertificateRestResponseV2); err == nil {
		*o = CertificateRestResponseV2(varCertificateRestResponseV2)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "fingerprint")
		delete(additionalProperties, "cAFingerprint")
		delete(additionalProperties, "certificateProfileId")
		delete(additionalProperties, "endEntityProfileId")
		delete(additionalProperties, "expireDate")
		delete(additionalProperties, "issuerDN")
		delete(additionalProperties, "notBefore")
		delete(additionalProperties, "revocationDate")
		delete(additionalProperties, "revocationReason")
		delete(additionalProperties, "serialNumber")
		delete(additionalProperties, "status")
		delete(additionalProperties, "subjectAltName")
		delete(additionalProperties, "subjectDN")
		delete(additionalProperties, "subjectKeyId")
		delete(additionalProperties, "tag")
		delete(additionalProperties, "type")
		delete(additionalProperties, "udpateTime")
		delete(additionalProperties, "username")
		delete(additionalProperties, "base64Cert")
		delete(additionalProperties, "certificateRequest")
		delete(additionalProperties, "crlPartitionIndex")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCertificateRestResponseV2 struct {
	value *CertificateRestResponseV2
	isSet bool
}

func (v NullableCertificateRestResponseV2) Get() *CertificateRestResponseV2 {
	return v.value
}

func (v *NullableCertificateRestResponseV2) Set(val *CertificateRestResponseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateRestResponseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateRestResponseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateRestResponseV2(val *CertificateRestResponseV2) *NullableCertificateRestResponseV2 {
	return &NullableCertificateRestResponseV2{value: val, isSet: true}
}

func (v NullableCertificateRestResponseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateRestResponseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


