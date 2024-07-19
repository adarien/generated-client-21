# \EventAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEvents**](EventAPI.md#GetEvents) | **Get** /v1/events | Returns a list of events



## GetEvents

> EventsV1DTO GetEvents(ctx).From(from).To(to).Type_(type_).Limit(limit).Offset(offset).Execute()

Returns a list of events

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/adarien/generated-client-21"
)

func main() {
	from := time.Now() // time.Time | Date and time of start of the sample, inclusive (in UTC)
	to := time.Now() // time.Time | Date and time of end of the sample, inclusive (in UTC)
	type_ := "type__example" // string | Event type (optional)
	limit := int64(789) // int64 | Maximum number of records to be returned (optional) (default to 50)
	offset := int64(789) // int64 | Number of records to skip for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EventAPI.GetEvents(context.Background()).From(from).To(to).Type_(type_).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EventAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: EventsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `EventAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Date and time of start of the sample, inclusive (in UTC) | 
 **to** | **time.Time** | Date and time of end of the sample, inclusive (in UTC) | 
 **type_** | **string** | Event type | 
 **limit** | **int64** | Maximum number of records to be returned | [default to 50]
 **offset** | **int64** | Number of records to skip for pagination | [default to 0]

### Return type

[**EventsV1DTO**](EventsV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

