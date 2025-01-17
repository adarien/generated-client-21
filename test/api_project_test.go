/*
School21 OpenAPI Specification

Testing ProjectAPIService

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

func Test_openapi_ProjectAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectAPIService GetLoginsByProjectId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int64

		resp, httpRes, err := apiClient.ProjectAPI.GetLoginsByProjectId(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService GetProjectByProjectId", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId int64

		resp, httpRes, err := apiClient.ProjectAPI.GetProjectByProjectId(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
