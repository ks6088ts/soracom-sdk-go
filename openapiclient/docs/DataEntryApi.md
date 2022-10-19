# \DataEntryApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDataEntry**](DataEntryApi.md#DeleteDataEntry) | **Delete** /data/{resource_type}/{resource_id}/{time} | Deletes a data entry
[**GetDataEntries**](DataEntryApi.md#GetDataEntries) | **Get** /data/{resource_type}/{resource_id} | Get data sent from a resource
[**GetDataEntry**](DataEntryApi.md#GetDataEntry) | **Get** /data/{resource_type}/{resource_id}/{time} | Gets a data entry
[**ListDataSourceResources**](DataEntryApi.md#ListDataSourceResources) | **Get** /data/resources | Get the list of data source resources



## DeleteDataEntry

> DeleteDataEntry(ctx, resourceType, resourceId, time).Execute()

Deletes a data entry



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
    resourceType := "resourceType_example" // string | データ送信元リソースの Type
    resourceId := "resourceId_example" // string | データ送信元リソースの ID
    time := int32(56) // int32 | 削除対象のデータエントリのタイムスタンプ (UNIX 時間 (ミリ秒))

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataEntryApi.DeleteDataEntry(context.Background(), resourceType, resourceId, time).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEntryApi.DeleteDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** | データ送信元リソースの Type | 
**resourceId** | **string** | データ送信元リソースの ID | 
**time** | **int32** | 削除対象のデータエントリのタイムスタンプ (UNIX 時間 (ミリ秒)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDataEntryRequest struct via the builder pattern


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


## GetDataEntries

> []DataEntry GetDataEntries(ctx, resourceType, resourceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get data sent from a resource



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
    resourceType := "resourceType_example" // string | データ送信元リソースの Type
    resourceId := "resourceId_example" // string | データ送信元リソースの ID
    from := int32(56) // int32 | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) (optional)
    to := int32(56) // int32 | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) (optional)
    sort := "sort_example" // string | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 (optional) (default to "desc")
    limit := int32(56) // int32 | 取得するデータエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataEntryApi.GetDataEntries(context.Background(), resourceType, resourceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEntryApi.GetDataEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataEntries`: []DataEntry
    fmt.Fprintf(os.Stdout, "Response from `DataEntryApi.GetDataEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** | データ送信元リソースの Type | 
**resourceId** | **string** | データ送信元リソースの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int32** | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) | 
 **to** | **int32** | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) | 
 **sort** | **string** | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 | [default to &quot;desc&quot;]
 **limit** | **int32** | 取得するデータエントリ数の上限 | 
 **lastEvaluatedKey** | **string** | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 | 

### Return type

[**[]DataEntry**](DataEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataEntry

> GetDataEntry(ctx, resourceType, resourceId, time).Execute()

Gets a data entry



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
    resourceType := "resourceType_example" // string | データ送信元リソースの Type
    resourceId := "resourceId_example" // string | データ送信元リソースの ID
    time := int32(56) // int32 | 取得対象のデータエントリのタイムスタンプ (UNIX 時間 (ミリ秒))

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataEntryApi.GetDataEntry(context.Background(), resourceType, resourceId, time).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEntryApi.GetDataEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceType** | **string** | データ送信元リソースの Type | 
**resourceId** | **string** | データ送信元リソースの ID | 
**time** | **int32** | 取得対象のデータエントリのタイムスタンプ (UNIX 時間 (ミリ秒)) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataEntryRequest struct via the builder pattern


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


## ListDataSourceResources

> []DataSourceResourceMetadata ListDataSourceResources(ctx).ResourceType(resourceType).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get the list of data source resources



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
    resourceType := "resourceType_example" // string | 対象のリソース Type (optional)
    limit := int32(56) // int32 | 取得するデータエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のデータ送信元リソースの ID。このパラメータを指定することで次のデータ送信元リソース以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DataEntryApi.ListDataSourceResources(context.Background()).ResourceType(resourceType).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataEntryApi.ListDataSourceResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDataSourceResources`: []DataSourceResourceMetadata
    fmt.Fprintf(os.Stdout, "Response from `DataEntryApi.ListDataSourceResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDataSourceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceType** | **string** | 対象のリソース Type | 
 **limit** | **int32** | 取得するデータエントリ数の上限 | 
 **lastEvaluatedKey** | **string** | 前ページで取得した最後のデータ送信元リソースの ID。このパラメータを指定することで次のデータ送信元リソース以降のリストを取得できる。 | 

### Return type

[**[]DataSourceResourceMetadata**](DataSourceResourceMetadata.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

