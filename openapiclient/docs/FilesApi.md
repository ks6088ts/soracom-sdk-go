# \FilesApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetExportedFile**](FilesApi.md#GetExportedFile) | **Get** /files/exported/{exported_file_id} | 非同期でファイルをエクスポートした場合の処理の進捗を取得する



## GetExportedFile

> GetExportedFileResponse GetExportedFile(ctx, exportedFileId).Execute()

非同期でファイルをエクスポートした場合の処理の進捗を取得する



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
    exportedFileId := "exportedFileId_example" // string | ファイルエクスポート ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FilesApi.GetExportedFile(context.Background(), exportedFileId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FilesApi.GetExportedFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExportedFile`: GetExportedFileResponse
    fmt.Fprintf(os.Stdout, "Response from `FilesApi.GetExportedFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportedFileId** | **string** | ファイルエクスポート ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportedFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetExportedFileResponse**](GetExportedFileResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

