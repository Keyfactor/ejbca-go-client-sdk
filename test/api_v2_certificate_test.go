/*
EJBCA REST Interface

Testing V2CertificateApiService

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

func Test_ejbca_V2CertificateApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test V2CertificateApiService GetCertificateProfileInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var profileName string

		resp, httpRes, err := apiClient.V2CertificateApi.GetCertificateProfileInfo(context.Background(), profileName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V2CertificateApiService SearchCertificates1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V2CertificateApi.SearchCertificates1(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V2CertificateApiService Status3", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.V2CertificateApi.Status3(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
