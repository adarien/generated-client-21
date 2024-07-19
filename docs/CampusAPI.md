# \CampusAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampuses**](CampusAPI.md#GetCampuses) | **Get** /v1/campuses | Returns a list of campuses
[**GetClustersByCampus**](CampusAPI.md#GetClustersByCampus) | **Get** /v1/campuses/{campusId}/clusters | Returns a list of clusters in a campus by ID
[**GetCoalitionsByCampus**](CampusAPI.md#GetCoalitionsByCampus) | **Get** /v1/campuses/{campusId}/coalitions | Returns a list of coalitions in a campus by ID
[**GetParticipantsByCampusId**](CampusAPI.md#GetParticipantsByCampusId) | **Get** /v1/campuses/{campusId}/participants | Returns a list of participants in a campus by ID



## GetCampuses

> CampusesV1DTO GetCampuses(ctx).Execute()

Returns a list of campuses

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampusAPI.GetCampuses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampusAPI.GetCampuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampuses`: CampusesV1DTO
	fmt.Fprintf(os.Stdout, "Response from `CampusAPI.GetCampuses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampusesRequest struct via the builder pattern


### Return type

[**CampusesV1DTO**](CampusesV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClustersByCampus

> ClustersV1DTO GetClustersByCampus(ctx, campusId).Execute()

Returns a list of clusters in a campus by ID

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
	campusId := "ff19a3a7-12f5-4332-9582-624519c3eaea" // string | Campus ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampusAPI.GetClustersByCampus(context.Background(), campusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampusAPI.GetClustersByCampus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetClustersByCampus`: ClustersV1DTO
	fmt.Fprintf(os.Stdout, "Response from `CampusAPI.GetClustersByCampus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campusId** | **string** | Campus ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetClustersByCampusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClustersV1DTO**](ClustersV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoalitionsByCampus

> CoalitionsV1DTO GetCoalitionsByCampus(ctx, campusId).Limit(limit).Offset(offset).Execute()

Returns a list of coalitions in a campus by ID

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
	campusId := "ff19a3a7-12f5-4332-9582-624519c3eaea" // string | Campus ID
	limit := int32(56) // int32 | Maximum number of records to be returned (optional) (default to 50)
	offset := int32(56) // int32 | Number of records to skip for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampusAPI.GetCoalitionsByCampus(context.Background(), campusId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampusAPI.GetCoalitionsByCampus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoalitionsByCampus`: CoalitionsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `CampusAPI.GetCoalitionsByCampus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campusId** | **string** | Campus ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoalitionsByCampusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of records to be returned | [default to 50]
 **offset** | **int32** | Number of records to skip for pagination | [default to 0]

### Return type

[**CoalitionsV1DTO**](CoalitionsV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantsByCampusId

> ParticipantLoginsV1DTO GetParticipantsByCampusId(ctx, campusId).Limit(limit).Offset(offset).Execute()

Returns a list of participants in a campus by ID

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
	campusId := "ff19a3a7-12f5-4332-9582-624519c3eaea" // string | Campus ID
	limit := int64(789) // int64 | Maximum number of records to be returned (optional) (default to 50)
	offset := int64(789) // int64 | Number of records to skip for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampusAPI.GetParticipantsByCampusId(context.Background(), campusId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampusAPI.GetParticipantsByCampusId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantsByCampusId`: ParticipantLoginsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `CampusAPI.GetParticipantsByCampusId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campusId** | **string** | Campus ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantsByCampusIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Maximum number of records to be returned | [default to 50]
 **offset** | **int64** | Number of records to skip for pagination | [default to 0]

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

