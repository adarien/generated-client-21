# \ParticipantAPI

All URIs are relative to *https://edu-api.21-school.ru/services/21-school/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBadgesByLogin**](ParticipantAPI.md#GetBadgesByLogin) | **Get** /v1/participants/{login}/badges | Returns a list of participant badges by login
[**GetCoalitionByLogin**](ParticipantAPI.md#GetCoalitionByLogin) | **Get** /v1/participants/{login}/coalition | Returns participant coalition information by login
[**GetLogWeeklyAvgHoursByLoginAndDate**](ParticipantAPI.md#GetLogWeeklyAvgHoursByLoginAndDate) | **Get** /v1/participants/{login}/logtime | Returns an average week logtime by login
[**GetParticipantByLogin**](ParticipantAPI.md#GetParticipantByLogin) | **Get** /v1/participants/{login} | Returns basic participant information by login
[**GetParticipantCourseByLoginAndCourseId**](ParticipantAPI.md#GetParticipantCourseByLoginAndCourseId) | **Get** /v1/participants/{login}/courses/{courseId} | Returns participant course information by ID
[**GetParticipantCoursesByLogin**](ParticipantAPI.md#GetParticipantCoursesByLogin) | **Get** /v1/participants/{login}/courses | Returns participant courses information by login
[**GetParticipantFeedbackByLogin**](ParticipantAPI.md#GetParticipantFeedbackByLogin) | **Get** /v1/participants/{login}/feedback | Returns average participant feedback points by login
[**GetParticipantProjectByLoginAndProjectId**](ParticipantAPI.md#GetParticipantProjectByLoginAndProjectId) | **Get** /v1/participants/{login}/projects/{projectId} | Returns participant project information by ID
[**GetParticipantProjectsByLogin**](ParticipantAPI.md#GetParticipantProjectsByLogin) | **Get** /v1/participants/{login}/projects | Returns participant projects information by login
[**GetParticipantWorkstationByLogin**](ParticipantAPI.md#GetParticipantWorkstationByLogin) | **Get** /v1/participants/{login}/workstation | Returns a participant workstation by login
[**GetPointsByLogin**](ParticipantAPI.md#GetPointsByLogin) | **Get** /v1/participants/{login}/points | Returns participant points information by login
[**GetSoftSkillByLogin**](ParticipantAPI.md#GetSoftSkillByLogin) | **Get** /v1/participants/{login}/skills | Returns participant skill points by login
[**GetXpHistoryByLogin**](ParticipantAPI.md#GetXpHistoryByLogin) | **Get** /v1/participants/{login}/experience-history | Returns a list of participant XP accruals by login



## GetBadgesByLogin

> ParticipantBadgesV1DTO GetBadgesByLogin(ctx, login).Execute()

Returns a list of participant badges by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetBadgesByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetBadgesByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBadgesByLogin`: ParticipantBadgesV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetBadgesByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBadgesByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantBadgesV1DTO**](ParticipantBadgesV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCoalitionByLogin

> ParticipantCoalitionV1DTO GetCoalitionByLogin(ctx, login).Execute()

Returns participant coalition information by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetCoalitionByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetCoalitionByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCoalitionByLogin`: ParticipantCoalitionV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetCoalitionByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCoalitionByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantCoalitionV1DTO**](ParticipantCoalitionV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLogWeeklyAvgHoursByLoginAndDate

> float64 GetLogWeeklyAvgHoursByLoginAndDate(ctx, login).Date(date).Execute()

Returns an average week logtime by login

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
	login := "bibikov-lukyan" // string | Login
	date := time.Now() // string | Date. The average logtime will be calculated for the week that includes the specified date. If not specified, the current date is used by default (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetLogWeeklyAvgHoursByLoginAndDate(context.Background(), login).Date(date).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetLogWeeklyAvgHoursByLoginAndDate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLogWeeklyAvgHoursByLoginAndDate`: float64
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetLogWeeklyAvgHoursByLoginAndDate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogWeeklyAvgHoursByLoginAndDateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **date** | **string** | Date. The average logtime will be calculated for the week that includes the specified date. If not specified, the current date is used by default | 

### Return type

**float64**

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantByLogin

> ParticipantV1DTO GetParticipantByLogin(ctx, login).Execute()

Returns basic participant information by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantByLogin`: ParticipantV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantV1DTO**](ParticipantV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantCourseByLoginAndCourseId

> ParticipantCourseV1DTO GetParticipantCourseByLoginAndCourseId(ctx, login, courseId).Execute()

Returns participant course information by ID

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
	login := "bibikov-lukyan" // string | Login
	courseId := int64(12311) // int64 | Course ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantCourseByLoginAndCourseId(context.Background(), login, courseId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantCourseByLoginAndCourseId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantCourseByLoginAndCourseId`: ParticipantCourseV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantCourseByLoginAndCourseId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 
**courseId** | **int64** | Course ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantCourseByLoginAndCourseIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParticipantCourseV1DTO**](ParticipantCourseV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantCoursesByLogin

> ParticipantCoursesV1DTO GetParticipantCoursesByLogin(ctx, login).Limit(limit).Offset(offset).Status(status).Execute()

Returns participant courses information by login

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
	login := "bibikov-lukyan" // string | Login
	limit := int64(789) // int64 | Maximum number of records to be returned (optional) (default to 10)
	offset := int64(789) // int64 | Number of records to skip for pagination (optional) (default to 0)
	status := "status_example" // string | Course status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantCoursesByLogin(context.Background(), login).Limit(limit).Offset(offset).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantCoursesByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantCoursesByLogin`: ParticipantCoursesV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantCoursesByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantCoursesByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Maximum number of records to be returned | [default to 10]
 **offset** | **int64** | Number of records to skip for pagination | [default to 0]
 **status** | **string** | Course status | 

### Return type

[**ParticipantCoursesV1DTO**](ParticipantCoursesV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantFeedbackByLogin

> ParticipantFeedbackV1DTO GetParticipantFeedbackByLogin(ctx, login).Execute()

Returns average participant feedback points by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantFeedbackByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantFeedbackByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantFeedbackByLogin`: ParticipantFeedbackV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantFeedbackByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantFeedbackByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantFeedbackV1DTO**](ParticipantFeedbackV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantProjectByLoginAndProjectId

> ParticipantProjectV1DTO GetParticipantProjectByLoginAndProjectId(ctx, login, projectId).Execute()

Returns participant project information by ID

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
	login := "bibikov-lukyan" // string | Login
	projectId := int64(12311) // int64 | Project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantProjectByLoginAndProjectId(context.Background(), login, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantProjectByLoginAndProjectId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantProjectByLoginAndProjectId`: ParticipantProjectV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantProjectByLoginAndProjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 
**projectId** | **int64** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantProjectByLoginAndProjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ParticipantProjectV1DTO**](ParticipantProjectV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantProjectsByLogin

> ParticipantProjectsV1DTO GetParticipantProjectsByLogin(ctx, login).Limit(limit).Offset(offset).Status(status).Execute()

Returns participant projects information by login

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
	login := "bibikov-lukyan" // string | Login
	limit := int64(789) // int64 | Maximum number of records to be returned (optional) (default to 10)
	offset := int64(789) // int64 | Number of records to skip for pagination (optional) (default to 0)
	status := "status_example" // string | Project status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantProjectsByLogin(context.Background(), login).Limit(limit).Offset(offset).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantProjectsByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantProjectsByLogin`: ParticipantProjectsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantProjectsByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantProjectsByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Maximum number of records to be returned | [default to 10]
 **offset** | **int64** | Number of records to skip for pagination | [default to 0]
 **status** | **string** | Project status | 

### Return type

[**ParticipantProjectsV1DTO**](ParticipantProjectsV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetParticipantWorkstationByLogin

> ParticipantWorkstationV1DTO GetParticipantWorkstationByLogin(ctx, login).Execute()

Returns a participant workstation by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetParticipantWorkstationByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetParticipantWorkstationByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetParticipantWorkstationByLogin`: ParticipantWorkstationV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetParticipantWorkstationByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetParticipantWorkstationByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantWorkstationV1DTO**](ParticipantWorkstationV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPointsByLogin

> ParticipantPointsV1DTO GetPointsByLogin(ctx, login).Execute()

Returns participant points information by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetPointsByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetPointsByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPointsByLogin`: ParticipantPointsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetPointsByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPointsByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantPointsV1DTO**](ParticipantPointsV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoftSkillByLogin

> ParticipantSkillsV1DTO GetSoftSkillByLogin(ctx, login).Execute()

Returns participant skill points by login

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
	login := "bibikov-lukyan" // string | Login

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetSoftSkillByLogin(context.Background(), login).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetSoftSkillByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSoftSkillByLogin`: ParticipantSkillsV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetSoftSkillByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoftSkillByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ParticipantSkillsV1DTO**](ParticipantSkillsV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetXpHistoryByLogin

> ParticipantXpHistoryV1DTO GetXpHistoryByLogin(ctx, login).Limit(limit).Offset(offset).Execute()

Returns a list of participant XP accruals by login

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
	login := "bibikov-lukyan" // string | Login
	limit := int64(789) // int64 | Maximum number of records to be returned (optional) (default to 50)
	offset := int64(789) // int64 | Number of records to skip for pagination (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ParticipantAPI.GetXpHistoryByLogin(context.Background(), login).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ParticipantAPI.GetXpHistoryByLogin``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetXpHistoryByLogin`: ParticipantXpHistoryV1DTO
	fmt.Fprintf(os.Stdout, "Response from `ParticipantAPI.GetXpHistoryByLogin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**login** | **string** | Login | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetXpHistoryByLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int64** | Maximum number of records to be returned | [default to 50]
 **offset** | **int64** | Number of records to skip for pagination | [default to 0]

### Return type

[**ParticipantXpHistoryV1DTO**](ParticipantXpHistoryV1DTO.md)

### Authorization

[Authorization](../README.md#Authorization)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

