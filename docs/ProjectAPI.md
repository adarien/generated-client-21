# \ProjectAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLoginsByProjectId**](ProjectAPI.md#GetLoginsByProjectId) | **Get** /v1/projects/{projectId}/participants | Returns a list of participants by project ID
[**GetProjectByProjectId**](ProjectAPI.md#GetProjectByProjectId) | **Get** /v1/projects/{projectId} | Returns project information by ID



## GetLoginsByProjectId

> ParticipantLoginsV1DTO GetLoginsByProjectId(ctx, projectId).Limit(limit).Offset(offset).Status(status).CampusId(campusId).Execute()

Returns a list of participants by project ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/adarien/generated-client-21"
)

func main() {
	projectId := int64(123131) // int64 | Project ID
	limit := int64(789) // int64 | Maximum number of records to be returned (optional) (default to 50)
	offset := int64(789) // int64 | Number of records to skip for pagination (optional) (default to 0)
	status := "status_example" // string | Project status (optional)
	campusId := "96098f4b-5708-4c42-a62c-6893419169b3" // string | Campus ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetLoginsByProjectId(context.Background(), projectId).Limit(limit).Offset(offset).Status(status).CampusId(campusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetLoginsByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoginsByProjectId`: ParticipantLoginsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetLoginsByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoginsByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Maximum number of records to be returned | [default to 50]
 **offset** | **int64** | Number of records to skip for pagination | [default to 0]
 **status** | **string** | Project status | 
 **campusId** | **string** | Campus ID | 

### Return type

[**ParticipantLoginsV1DTO**](ParticipantLoginsV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectByProjectId

> ProjectV1DTO GetProjectByProjectId(ctx, projectId).Execute()

Returns project information by ID

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/adarien/generated-client-21"
)

func main() {
	projectId := int64(123131) // int64 | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAPI.GetProjectByProjectId(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAPI.GetProjectByProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectByProjectId`: ProjectV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ProjectAPI.GetProjectByProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectByProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectV1DTO**](ProjectV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

