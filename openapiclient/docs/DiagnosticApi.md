# \DiagnosticApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDiagnostic**](DiagnosticApi.md#GetDiagnostic) | **Get** /diagnostics/{diagnostic_id} | 診断結果を取得します。
[**SendDiagnosticRequest**](DiagnosticApi.md#SendDiagnosticRequest) | **Post** /diagnostics | 診断リクエストを送信します。



## GetDiagnostic

> DiagnosticResponse GetDiagnostic(ctx, diagnosticId).Execute()

診断結果を取得します。



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
    diagnosticId := "diagnosticId_example" // string | 診断リクエストの識別子

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiagnosticApi.GetDiagnostic(context.Background(), diagnosticId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosticApi.GetDiagnostic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiagnostic`: DiagnosticResponse
    fmt.Fprintf(os.Stdout, "Response from `DiagnosticApi.GetDiagnostic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**diagnosticId** | **string** | 診断リクエストの識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiagnosticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DiagnosticResponse**](DiagnosticResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDiagnosticRequest

> []DiagnosticResponse SendDiagnosticRequest(ctx).DiagnosticRequest(diagnosticRequest).Execute()

診断リクエストを送信します。



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
    diagnosticRequest := *openapiclient.NewDiagnosticRequest("ResourceId_example", "ResourceType_example", "Service_example") // DiagnosticRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DiagnosticApi.SendDiagnosticRequest(context.Background()).DiagnosticRequest(diagnosticRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiagnosticApi.SendDiagnosticRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendDiagnosticRequest`: []DiagnosticResponse
    fmt.Fprintf(os.Stdout, "Response from `DiagnosticApi.SendDiagnosticRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendDiagnosticRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **diagnosticRequest** | [**DiagnosticRequest**](DiagnosticRequest.md) | request | 

### Return type

[**[]DiagnosticResponse**](DiagnosticResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

