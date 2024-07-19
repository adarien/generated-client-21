# \CoalitionAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetParticipantsByCoalitionId**](CoalitionAPI.md#GetParticipantsByCoalitionId) | **Get** /v1/coalitions/{coalitionId}/participants | Returns a list of participants in a coalition by ID



## GetParticipantsByCoalitionId

> ParticipantLoginsV1DTO GetParticipantsByCoalitionId(ctx, coalitionId).Limit(limit).Offset(offset).Execute()

Returns a list of participants in a coalition by ID

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
	coalitionId := int64(319) // int64 | Coalition ID
	limit := int32(56) // int32 | Maximum number of records to be returned (optional) (default to 50)
	offset := int32(56) // int32 | Number of records to skip for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CoalitionAPI.GetParticipantsByCoalitionId(context.Background(), coalitionId).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CoalitionAPI.GetParticipantsByCoalitionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantsByCoalitionId`: ParticipantLoginsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `CoalitionAPI.GetParticipantsByCoalitionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**coalitionId** | **int64** | Coalition ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantsByCoalitionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of records to be returned | [default to 50]
 **offset** | **int32** | Number of records to skip for pagination | [default to 0]

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

