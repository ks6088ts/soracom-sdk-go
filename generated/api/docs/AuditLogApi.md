# \AuditLogApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApiAuditLogs**](AuditLogApi.md#GetApiAuditLogs) | **Get** /audit_logs/api | SORACOM API 呼び出しの監査ログを取得する
[**GetNapterAuditLogs**](AuditLogApi.md#GetNapterAuditLogs) | **Get** /audit_logs/napter | SORACOM Napter の監査ログを取得する



## GetApiAuditLogs

> []APIAuditLogEntry GetApiAuditLogs(ctx).ApiKind(apiKind).FromEpochMs(fromEpochMs).ToEpochMs(toEpochMs).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

SORACOM API 呼び出しの監査ログを取得する



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
    apiKind := "apiKind_example" // string | ログ取得のフィルターに用いる API 種別 (例: `/v1/auth`) (optional)
    fromEpochMs := int32(56) // int32 | ログ取得対象の期間の始まり (unixtime milliseconds) (optional)
    toEpochMs := int32(56) // int32 | ログ取得対象の期間の終わり (unixtime milliseconds) (optional)
    limit := int32(56) // int32 | 取得するログエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のログエントリのタイムスタンプ (`requestedTimeEpochMs`)。このパラメータを指定することで次のログエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.GetApiAuditLogs(context.Background()).ApiKind(apiKind).FromEpochMs(fromEpochMs).ToEpochMs(toEpochMs).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetApiAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiAuditLogs`: []APIAuditLogEntry
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetApiAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiKind** | **string** | ログ取得のフィルターに用いる API 種別 (例: &#x60;/v1/auth&#x60;) | 
 **fromEpochMs** | **int32** | ログ取得対象の期間の始まり (unixtime milliseconds) | 
 **toEpochMs** | **int32** | ログ取得対象の期間の終わり (unixtime milliseconds) | 
 **limit** | **int32** | 取得するログエントリ数の上限 | 
 **lastEvaluatedKey** | **string** | 前ページで取得した最後のログエントリのタイムスタンプ (&#x60;requestedTimeEpochMs&#x60;)。このパラメータを指定することで次のログエントリ以降のリストを取得できる。 | 

### Return type

[**[]APIAuditLogEntry**](APIAuditLogEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNapterAuditLogs

> []NapterAuditLogEntry GetNapterAuditLogs(ctx).ResourceType(resourceType).ResourceId(resourceId).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

SORACOM Napter の監査ログを取得する



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
    from := int32(56) // int32 | ログ取得対象の期間の始まり (unixtime milliseconds) (optional)
    to := int32(56) // int32 | ログ取得対象の期間の終わり (unixtime milliseconds) (optional)
    limit := int32(56) // int32 | 取得するログエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のログエントリのタイムスタンプ。このパラメータを指定することで次のログエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuditLogApi.GetNapterAuditLogs(context.Background()).ResourceType(resourceType).ResourceId(resourceId).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.GetNapterAuditLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNapterAuditLogs`: []NapterAuditLogEntry
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.GetNapterAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNapterAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | ログ取得対象のリソースのタイプ | 
 **resourceId** | **string** | ログ取得対象のリソースの ID | 
 **from** | **int32** | ログ取得対象の期間の始まり (unixtime milliseconds) | 
 **to** | **int32** | ログ取得対象の期間の終わり (unixtime milliseconds) | 
 **limit** | **int32** | 取得するログエントリ数の上限 | 
 **lastEvaluatedKey** | **string** | 前ページで取得した最後のログエントリのタイムスタンプ。このパラメータを指定することで次のログエントリ以降のリストを取得できる。 | 

### Return type

[**[]NapterAuditLogEntry**](NapterAuditLogEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

