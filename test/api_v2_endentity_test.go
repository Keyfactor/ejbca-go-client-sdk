/*
EJBCA REST Interface

Testing V2EndentityApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ejbca

import (
	"context"
	openapiclient "github.com/Keyfactor/ejbca-go-client-sdk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_ejbca_V2EndentityApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test V2EndentityApiService GetAuthorizedEndEntityProfiles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.V2EndentityApi.GetAuthorizedEndEntityProfiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V2EndentityApiService Profile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var endentityProfileName string

		resp, httpRes, err := apiClient.V2EndentityApi.Profile(context.Background(), endentityProfileName).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V2EndentityApiService SortedSearch", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.V2EndentityApi.SortedSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test V2EndentityApiService Status7", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.V2EndentityApi.Status7(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
