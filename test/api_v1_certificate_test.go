/*
EJBCA REST Interface

Testing V1CertificateApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ejbca

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/Keyfactor/ejbca-go-client-sdk/ejbca"
)

func Test_ejbca_V1CertificateApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test V1CertificateApiService CertificateRequest", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1CertificateApi.CertificateRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService EnrollKeystore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1CertificateApi.EnrollKeystore(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService EnrollPkcs10Certificate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1CertificateApi.EnrollPkcs10Certificate(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService FinalizeEnrollment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var requestId int32

		resp, httpRes, err := apiClient.V1CertificateApi.FinalizeEnrollment(context.Background(), requestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService GetCertificatesAboutToExpire", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1CertificateApi.GetCertificatesAboutToExpire(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService RevocationStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issuerDn string
		var certificateSerialNumber string

		resp, httpRes, err := apiClient.V1CertificateApi.RevocationStatus(context.Background(), issuerDn, certificateSerialNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService RevokeCertificate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issuerDn string
		var certificateSerialNumber string

		resp, httpRes, err := apiClient.V1CertificateApi.RevokeCertificate(context.Background(), issuerDn, certificateSerialNumber).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService SearchCertificates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1CertificateApi.SearchCertificates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V1CertificateApiService Status2", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V1CertificateApi.Status2(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
