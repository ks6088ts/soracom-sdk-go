# \StatsApi

All URIs are relative to *https://api-sandbox.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxInsertAirStats**](StatsApi.md#SandboxInsertAirStats) | **Post** /sandbox/stats/air/subscribers/{imsi} | テスト用に Air の統計情報を登録する
[**SandboxInsertBeamStats**](StatsApi.md#SandboxInsertBeamStats) | **Post** /sandbox/stats/beam/subscribers/{imsi} | テスト用に Beam の統計情報を登録する



## SandboxInsertAirStats

> SandboxInsertAirStats(ctx, imsi).SandboxInsertAirStatsRequest(sandboxInsertAirStatsRequest).Execute()

テスト用に Air の統計情報を登録する



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
    imsi := "imsi_example" // string | IMSI
    sandboxInsertAirStatsRequest := *openapiclient.NewSandboxInsertAirStatsRequest() // SandboxInsertAirStatsRequest | ある時点のデータ通信量の統計情報

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.SandboxInsertAirStats(context.Background(), imsi).SandboxInsertAirStatsRequest(sandboxInsertAirStatsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.SandboxInsertAirStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxInsertAirStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxInsertAirStatsRequest** | [**SandboxInsertAirStatsRequest**](SandboxInsertAirStatsRequest.md) | ある時点のデータ通信量の統計情報 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxInsertBeamStats

> SandboxInsertBeamStats(ctx, imsi).SandboxInsertBeamStatsRequest(sandboxInsertBeamStatsRequest).Execute()

テスト用に Beam の統計情報を登録する



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
    imsi := "imsi_example" // string | IMSI
    sandboxInsertBeamStatsRequest := *openapiclient.NewSandboxInsertBeamStatsRequest() // SandboxInsertBeamStatsRequest | ある時点のリクエスト数の統計情報

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.SandboxInsertBeamStats(context.Background(), imsi).SandboxInsertBeamStatsRequest(sandboxInsertBeamStatsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.SandboxInsertBeamStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxInsertBeamStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxInsertBeamStatsRequest** | [**SandboxInsertBeamStatsRequest**](SandboxInsertBeamStatsRequest.md) | ある時点のリクエスト数の統計情報 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

