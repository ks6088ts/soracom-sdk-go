# \SoraletApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSoralet**](SoraletApi.md#CreateSoralet) | **Post** /soralets | Soralet を作成します。
[**DeleteSoralet**](SoraletApi.md#DeleteSoralet) | **Delete** /soralets/{soralet_id} | Soralet を削除します。
[**DeleteSoraletVersion**](SoraletApi.md#DeleteSoraletVersion) | **Delete** /soralets/{soralet_id}/versions/{version} | Soralet のバージョンを削除します。
[**GetSoralet**](SoraletApi.md#GetSoralet) | **Get** /soralets/{soralet_id} | Soralet を取得します。
[**GetSoraletLogs**](SoraletApi.md#GetSoraletLogs) | **Get** /soralets/{soralet_id}/logs | Soralet のログメッセージを取得します。
[**ListSoraletVersions**](SoraletApi.md#ListSoraletVersions) | **Get** /soralets/{soralet_id}/versions | Soralet のバージョン一覧を取得します。
[**ListSoralets**](SoraletApi.md#ListSoralets) | **Get** /soralets | Soralet の一覧を取得します。
[**TestSoralet**](SoraletApi.md#TestSoralet) | **Post** /soralets/{soralet_id}/test | 引数を指定して Soralet をテスト実行します。
[**UploadSoraletCode**](SoraletApi.md#UploadSoraletCode) | **Post** /soralets/{soralet_id}/versions | コードをアップロードし、新しいバージョンを作成します。



## CreateSoralet

> CreateSoralet(ctx).CreateSoraletRequest(createSoraletRequest).Execute()

Soralet を作成します。



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
    createSoraletRequest := *openapiclient.NewCreateSoraletRequest("SoraletId_example") // CreateSoraletRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.CreateSoralet(context.Background()).CreateSoraletRequest(createSoraletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.CreateSoralet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSoraletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSoraletRequest** | [**CreateSoraletRequest**](CreateSoraletRequest.md) | request | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSoralet

> DeleteSoralet(ctx, soraletId).Execute()

Soralet を削除します。



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
    soraletId := "soraletId_example" // string | 削除したい Soralet の識別子

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.DeleteSoralet(context.Background(), soraletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.DeleteSoralet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | 削除したい Soralet の識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSoraletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSoraletVersion

> DeleteSoraletVersion(ctx, soraletId, version).Execute()

Soralet のバージョンを削除します。



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
    soraletId := "soraletId_example" // string | Soralet の識別子
    version := int32(56) // int32 | 削除したい Soralet のバージョン

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.DeleteSoraletVersion(context.Background(), soraletId, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.DeleteSoraletVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | Soralet の識別子 | 
**version** | **int32** | 削除したい Soralet のバージョン | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSoraletVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoralet

> Soralet GetSoralet(ctx, soraletId).Execute()

Soralet を取得します。



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
    soraletId := "soraletId_example" // string | 取得したい Soralet の識別子

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.GetSoralet(context.Background(), soraletId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.GetSoralet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoralet`: Soralet
    fmt.Fprintf(os.Stdout, "Response from `SoraletApi.GetSoralet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | 取得したい Soralet の識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Soralet**](Soralet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoraletLogs

> []SoraletLog GetSoraletLogs(ctx, soraletId).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Soralet のログメッセージを取得します。



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
    soraletId := "soraletId_example" // string | Soralet の識別子
    sort := "sort_example" // string | Sort order (optional) (default to "desc")
    limit := int32(56) // int32 | 一度のレスポンスに含まれる項目数の最大値 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現在のページで最後に取得されたログの識別子を指定します。このパラメータを指定することで、ログのリストの続きを取得できます。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.GetSoraletLogs(context.Background(), soraletId).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.GetSoraletLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoraletLogs`: []SoraletLog
    fmt.Fprintf(os.Stdout, "Response from `SoraletApi.GetSoraletLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | Soralet の識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraletLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | Sort order | [default to &quot;desc&quot;]
 **limit** | **int32** | 一度のレスポンスに含まれる項目数の最大値 | 
 **lastEvaluatedKey** | **string** | 現在のページで最後に取得されたログの識別子を指定します。このパラメータを指定することで、ログのリストの続きを取得できます。 | 

### Return type

[**[]SoraletLog**](SoraletLog.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraletVersions

> []SoraletVersion ListSoraletVersions(ctx, soraletId).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Soralet のバージョン一覧を取得します。



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
    soraletId := "soraletId_example" // string | バージョン一覧を取得したい Soralet の識別子
    sort := "sort_example" // string | Sort order (optional) (default to "desc")
    limit := int32(56) // int32 | 一度のレスポンスに含まれる項目数の最大値 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現在のページで最後に取得された Soralet の識別子を指定します。このパラメータを指定することで、Soralet のリストの続きを取得できます。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.ListSoraletVersions(context.Background(), soraletId).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.ListSoraletVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraletVersions`: []SoraletVersion
    fmt.Fprintf(os.Stdout, "Response from `SoraletApi.ListSoraletVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | バージョン一覧を取得したい Soralet の識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSoraletVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sort** | **string** | Sort order | [default to &quot;desc&quot;]
 **limit** | **int32** | 一度のレスポンスに含まれる項目数の最大値 | 
 **lastEvaluatedKey** | **string** | 現在のページで最後に取得された Soralet の識別子を指定します。このパラメータを指定することで、Soralet のリストの続きを取得できます。 | 

### Return type

[**[]SoraletVersion**](SoraletVersion.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoralets

> []Soralet ListSoralets(ctx).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Soralet の一覧を取得します。



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
    sort := "sort_example" // string | Sort order (optional) (default to "asc")
    limit := int32(56) // int32 | 一度のレスポンスに含まれる項目数の最大値 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現在のページで最後に取得された Soralet の識別子を指定します。このパラメータを指定することで、Soralet のリストの続きを取得できます。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.ListSoralets(context.Background()).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.ListSoralets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoralets`: []Soralet
    fmt.Fprintf(os.Stdout, "Response from `SoraletApi.ListSoralets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSoraletsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sort** | **string** | Sort order | [default to &quot;asc&quot;]
 **limit** | **int32** | 一度のレスポンスに含まれる項目数の最大値 | 
 **lastEvaluatedKey** | **string** | 現在のページで最後に取得された Soralet の識別子を指定します。このパラメータを指定することで、Soralet のリストの続きを取得できます。 | 

### Return type

[**[]Soralet**](Soralet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestSoralet

> ExecuteSoraletResponse TestSoralet(ctx, soraletId).ExecuteSoraletRequest(executeSoraletRequest).Execute()

引数を指定して Soralet をテスト実行します。



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
    soraletId := "soraletId_example" // string | Soralet の識別子
    executeSoraletRequest := *openapiclient.NewExecuteSoraletRequest("ContentType_example", "Direction_example", "Payload_example", map[string]SoraletDataSource{"key": *openapiclient.NewSoraletDataSource("ResourceId_example", "ResourceType_example")}, "Version_example") // ExecuteSoraletRequest | 実行リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.TestSoralet(context.Background(), soraletId).ExecuteSoraletRequest(executeSoraletRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.TestSoralet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestSoralet`: ExecuteSoraletResponse
    fmt.Fprintf(os.Stdout, "Response from `SoraletApi.TestSoralet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | Soralet の識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiTestSoraletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **executeSoraletRequest** | [**ExecuteSoraletRequest**](ExecuteSoraletRequest.md) | 実行リクエスト | 

### Return type

[**ExecuteSoraletResponse**](ExecuteSoraletResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadSoraletCode

> SoraletVersion UploadSoraletCode(ctx, soraletId).Body(body).ContentType(contentType).Execute()

コードをアップロードし、新しいバージョンを作成します。



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
    soraletId := "soraletId_example" // string | コードのアップロード先の Soralet の識別子
    body := os.NewFile(1234, "some_file") // *os.File | アップロードしたいファイルの内容
    contentType := "contentType_example" // string | アップロードするファイルの Content-Type。.wasm ファイルをアップロードする場合は `application/json` を指定してください。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraletApi.UploadSoraletCode(context.Background(), soraletId).Body(body).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraletApi.UploadSoraletCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadSoraletCode`: SoraletVersion
    fmt.Fprintf(os.Stdout, "Response from `SoraletApi.UploadSoraletCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**soraletId** | **string** | コードのアップロード先の Soralet の識別子 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadSoraletCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | ***os.File** | アップロードしたいファイルの内容 | 
 **contentType** | **string** | アップロードするファイルの Content-Type。.wasm ファイルをアップロードする場合は &#x60;application/json&#x60; を指定してください。 | 

### Return type

[**SoraletVersion**](SoraletVersion.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

