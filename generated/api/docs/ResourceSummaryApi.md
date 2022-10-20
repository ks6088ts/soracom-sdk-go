# \ResourceSummaryApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetResourceSummary**](ResourceSummaryApi.md#GetResourceSummary) | **Get** /resource_summaries/{resource_summary_type} | resource_summary_type で指定された種別のリソースの要約を取得します。



## GetResourceSummary

> ResourceSummary GetResourceSummary(ctx, resourceSummaryType).Execute()

resource_summary_type で指定された種別のリソースの要約を取得します。



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    resourceSummaryType := "resourceSummaryType_example" // string | リソースの要約の種別  - `simsPerStatus`: ステータスごとの IoT SIM の数 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResourceSummaryApi.GetResourceSummary(context.Background(), resourceSummaryType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResourceSummaryApi.GetResourceSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceSummary`: ResourceSummary
    fmt.Fprintf(os.Stdout, "Response from `ResourceSummaryApi.GetResourceSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceSummaryType** | **string** | リソースの要約の種別  - &#x60;simsPerStatus&#x60;: ステータスごとの IoT SIM の数  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ResourceSummary**](ResourceSummary.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

