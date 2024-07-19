# \ClusterAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParticipantsByCoalitionId1**](ClusterAPI.md#GetParticipantsByCoalitionId1) | **Get** /v1/clusters/{clusterId}/map | Returns a cluster map



## GetParticipantsByCoalitionId1

> ClusterMapV1DTO GetParticipantsByCoalitionId1(ctx, clusterId).Limit(limit).Offset(offset).Occupied(occupied).Execute()

Returns a cluster map

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
	clusterId := int64(825) // int64 | Cluster ID
	limit := int32(56) // int32 | Maximum number of records to be returned (optional) (default to 50)
	offset := int32(56) // int32 | Number of records to skip for pagination (optional) (default to 0)
	occupied := true // bool | If true, occupied seats are returned. If false - only free seats are returned. If not specified - all seats are listed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClusterAPI.GetParticipantsByCoalitionId1(context.Background(), clusterId).Limit(limit).Offset(offset).Occupied(occupied).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClusterAPI.GetParticipantsByCoalitionId1``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantsByCoalitionId1`: ClusterMapV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ClusterAPI.GetParticipantsByCoalitionId1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clusterId** | **int64** | Cluster ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantsByCoalitionId1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of records to be returned | [default to 50]
 **offset** | **int32** | Number of records to skip for pagination | [default to 0]
 **occupied** | **bool** | If true, occupied seats are returned. If false - only free seats are returned. If not specified - all seats are listed | 

### Return type

[**ClusterMapV1DTO**](ClusterMapV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

