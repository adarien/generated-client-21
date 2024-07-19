# \CourseAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCourseByCourseId**](CourseAPI.md#GetCourseByCourseId) | **Get** /v1/courses/{courseId} | Returns course information by ID



## GetCourseByCourseId

> CourseV1DTO GetCourseByCourseId(ctx, courseId).Execute()

Returns course information by ID

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
	courseId := int64(123131) // int64 | Course ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CourseAPI.GetCourseByCourseId(context.Background(), courseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CourseAPI.GetCourseByCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCourseByCourseId`: CourseV1DTO
	fmt.Fprintf(os.Stdout, "Response from `CourseAPI.GetCourseByCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**courseId** | **int64** | Course ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCourseByCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CourseV1DTO**](CourseV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

