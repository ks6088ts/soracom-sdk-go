# \LogApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLogs**](LogApi.md#GetLogs) | **Get** /logs | Get Logs.



## GetLogs

> []LogEntry GetLogs(ctx).ResourceType(resourceType).ResourceId(resourceId).Service(service).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get Logs.



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
    resourceType := "resourceType_example" // string | ログ取得対象のリソースのタイプ (optional)
    resourceId := "resourceId_example" // string | ログ取得対象のリソースの ID (optional)
    service := "service_example" // string | ログエントリをフィルタするためのサービス名 (optional)
    from := int32(56) // int32 | ログ取得対象の期間の始まり (unixtime) (optional)
    to := int32(56) // int32 | ログ取得対象の期間の終わり (unixtime) (optional)
    limit := int32(56) // int32 | 取得するログエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のログエントリのタイムスタンプ。このパラメータを指定することで次のログエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogApi.GetLogs(context.Background()).ResourceType(resourceType).ResourceId(resourceId).Service(service).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogApi.GetLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLogs`: []LogEntry
    fmt.Fprintf(os.Stdout, "Response from `LogApi.GetLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | ログ取得対象のリソースのタイプ | 
 **resourceId** | **string** | ログ取得対象のリソースの ID | 
 **service** | **string** | ログエントリをフィルタするためのサービス名 | 
 **from** | **int32** | ログ取得対象の期間の始まり (unixtime) | 
 **to** | **int32** | ログ取得対象の期間の終わり (unixtime) | 
 **limit** | **int32** | 取得するログエントリ数の上限 | 
 **lastEvaluatedKey** | **string** | 前ページで取得した最後のログエントリのタイムスタンプ。このパラメータを指定することで次のログエントリ以降のリストを取得できる。 | 

### Return type

[**[]LogEntry**](LogEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

