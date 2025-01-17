/*
School21 OpenAPI Specification

Testing ClusterAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/adarien/generated-client-21"
)

func Test_openapi_ClusterAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ClusterAPIService GetParticipantsByCoalitionId1", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clusterId int64

		resp, httpRes, err := apiClient.ClusterAPI.GetParticipantsByCoalitionId1(context.Background(), clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
