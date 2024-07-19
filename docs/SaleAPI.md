# \SaleAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSales**](SaleAPI.md#GetSales) | **Get** /v1/sales | Returns current sales’ statuses within parallel



## GetSales

> SalesV1DTO GetSales(ctx).Execute()

Returns current sales’ statuses within parallel

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
	resp, r, err := apiClient.SaleAPI.GetSales(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SaleAPI.GetSales``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSales`: SalesV1DTO
	fmt.Fprintf(os.Stdout, "Response from `SaleAPI.GetSales`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSalesRequest struct via the builder pattern


### Return type

[**SalesV1DTO**](SalesV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

