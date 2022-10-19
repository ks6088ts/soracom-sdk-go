# \FileEntryApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDirectory**](FileEntryApi.md#DeleteDirectory) | **Delete** /files/{scope}/{path}/ | scope と path で指定されたディレクトリを削除します。
[**DeleteFile**](FileEntryApi.md#DeleteFile) | **Delete** /files/{scope}/{path} | scope と path で指定されたファイルを削除します。
[**FindFiles**](FileEntryApi.md#FindFiles) | **Get** /files | scope と prefix にマッチするファイルを探します。
[**GetFile**](FileEntryApi.md#GetFile) | **Get** /files/{scope}/{path} | scope と path で指定されたファイルをダウンロードします。
[**GetFileMetadata**](FileEntryApi.md#GetFileMetadata) | **Head** /files/{scope}/{path} | scope と path で指定されたファイルのメタデータを取得します。
[**ListFiles**](FileEntryApi.md#ListFiles) | **Get** /files/{scope}/{path}/ | scope と path で指定されたファイルやディレクトリの一覧を取得します。
[**PutFile**](FileEntryApi.md#PutFile) | **Put** /files/{scope}/{path} | 指定された scope 内の path にファイルをアップロードします。



## DeleteDirectory

> DeleteDirectory(ctx, scope, path).Execute()

scope と path で指定されたディレクトリを削除します。



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
    scope := "scope_example" // string | リクエストのスコープ (default to "private")
    path := "path_example" // string | 対象 Path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.DeleteDirectory(context.Background(), scope, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.DeleteDirectory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | リクエストのスコープ | [default to &quot;private&quot;]
**path** | **string** | 対象 Path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDirectoryRequest struct via the builder pattern


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


## DeleteFile

> DeleteFile(ctx, scope, path).Execute()

scope と path で指定されたファイルを削除します。



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
    scope := "scope_example" // string | リクエストのスコープ (default to "private")
    path := "path_example" // string | 対象 Path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.DeleteFile(context.Background(), scope, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.DeleteFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | リクエストのスコープ | [default to &quot;private&quot;]
**path** | **string** | 対象 Path | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileRequest struct via the builder pattern


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


## FindFiles

> []FileEntry FindFiles(ctx).Scope(scope).Prefix(prefix).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

scope と prefix にマッチするファイルを探します。



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
    scope := "scope_example" // string | リクエストのスコープ
    prefix := "prefix_example" // string | ファイルパスにマッチさせるプレフィックス
    limit := int32(56) // int32 | 返却するファイルエントリ数の上限 (optional) (default to 10)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のファイルエントリの filePath 。このパラメータを指定することで次のファイルエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.FindFiles(context.Background()).Scope(scope).Prefix(prefix).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.FindFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindFiles`: []FileEntry
    fmt.Fprintf(os.Stdout, "Response from `FileEntryApi.FindFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | リクエストのスコープ | 
 **prefix** | **string** | ファイルパスにマッチさせるプレフィックス | 
 **limit** | **int32** | 返却するファイルエントリ数の上限 | [default to 10]
 **lastEvaluatedKey** | **string** | 現ページで取得した最後のファイルエントリの filePath 。このパラメータを指定することで次のファイルエントリ以降のリストを取得できる。 | 

### Return type

[**[]FileEntry**](FileEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFile

> GetFile(ctx, scope, path).Execute()

scope と path で指定されたファイルをダウンロードします。



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
    scope := "scope_example" // string | リクエストのスコープ (default to "private")
    path := "path_example" // string | 対象 Path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.GetFile(context.Background(), scope, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.GetFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | リクエストのスコープ | [default to &quot;private&quot;]
**path** | **string** | 対象 Path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileRequest struct via the builder pattern


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


## GetFileMetadata

> FileEntry GetFileMetadata(ctx, scope, path).Execute()

scope と path で指定されたファイルのメタデータを取得します。



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
    scope := "scope_example" // string | リクエストのスコープ (default to "private")
    path := "path_example" // string | 対象 Path

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.GetFileMetadata(context.Background(), scope, path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.GetFileMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileMetadata`: FileEntry
    fmt.Fprintf(os.Stdout, "Response from `FileEntryApi.GetFileMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | リクエストのスコープ | [default to &quot;private&quot;]
**path** | **string** | 対象 Path | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**FileEntry**](FileEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFiles

> []FileEntry ListFiles(ctx, scope, path).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

scope と path で指定されたファイルやディレクトリの一覧を取得します。



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
    scope := "scope_example" // string | リクエストのスコープ (default to "private")
    path := "path_example" // string | 対象 Path (default to "/")
    limit := int32(56) // int32 | 返却するファイルエントリ数の上限 (optional) (default to 10)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のファイルエントリの filename 。このパラメータを指定することで次のファイルエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.ListFiles(context.Background(), scope, path).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.ListFiles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFiles`: []FileEntry
    fmt.Fprintf(os.Stdout, "Response from `FileEntryApi.ListFiles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | リクエストのスコープ | [default to &quot;private&quot;]
**path** | **string** | 対象 Path | [default to &quot;/&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiListFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | 返却するファイルエントリ数の上限 | [default to 10]
 **lastEvaluatedKey** | **string** | 現ページで取得した最後のファイルエントリの filename 。このパラメータを指定することで次のファイルエントリ以降のリストを取得できる。 | 

### Return type

[**[]FileEntry**](FileEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFile

> PutFile(ctx, scope, path).Body(body).ContentType(contentType).Execute()

指定された scope 内の path にファイルをアップロードします。



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
    scope := "scope_example" // string | リクエストのスコープ (default to "private")
    path := "path_example" // string | 対象 Path
    body := os.NewFile(1234, "some_file") // *os.File | Content of the file to upload
    contentType := "contentType_example" // string | Content type of the file to upload (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileEntryApi.PutFile(context.Background(), scope, path).Body(body).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileEntryApi.PutFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**scope** | **string** | リクエストのスコープ | [default to &quot;private&quot;]
**path** | **string** | 対象 Path | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | ***os.File** | Content of the file to upload | 
 **contentType** | **string** | Content type of the file to upload | 

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

