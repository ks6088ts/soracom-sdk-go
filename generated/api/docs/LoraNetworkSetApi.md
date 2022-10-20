# \LoraNetworkSetApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddPermissionToLoraNetworkSet**](LoraNetworkSetApi.md#AddPermissionToLoraNetworkSet) | **Post** /lora_network_sets/{ns_id}/add_permission | Adds permission to a LoRa network set.
[**CreateLoraNetworkSet**](LoraNetworkSetApi.md#CreateLoraNetworkSet) | **Post** /lora_network_sets | Creates a LoRa network set
[**DeleteLoraNetworkSet**](LoraNetworkSetApi.md#DeleteLoraNetworkSet) | **Delete** /lora_network_sets/{ns_id} | Delete LoRa network set.
[**DeleteLoraNetworkSetTag**](LoraNetworkSetApi.md#DeleteLoraNetworkSetTag) | **Delete** /lora_network_sets/{ns_id}/tags/{tag_name} | Delete LoRa network set tag.
[**GetLoraNetworkSet**](LoraNetworkSetApi.md#GetLoraNetworkSet) | **Get** /lora_network_sets/{ns_id} | Get LoRa network set.
[**ListGatewaysInLoraNetworkSet**](LoraNetworkSetApi.md#ListGatewaysInLoraNetworkSet) | **Get** /lora_network_sets/{ns_id}/gateways | List LoRa Gateways in a Network Set.
[**ListLoraNetworkSets**](LoraNetworkSetApi.md#ListLoraNetworkSets) | **Get** /lora_network_sets | List LoRa Network Sets.
[**PutLoraNetworkSetTags**](LoraNetworkSetApi.md#PutLoraNetworkSetTags) | **Put** /lora_network_sets/{ns_id}/tags | Bulk Insert or Update LoRa network set tags.
[**RevokePermissionFromLoraNetworkSet**](LoraNetworkSetApi.md#RevokePermissionFromLoraNetworkSet) | **Post** /lora_network_sets/{ns_id}/revoke_permission | Revokes a permission from a LoRa network set.



## AddPermissionToLoraNetworkSet

> LoraNetworkSet AddPermissionToLoraNetworkSet(ctx, nsId).UpdatePermissionRequest(updatePermissionRequest).Execute()

Adds permission to a LoRa network set.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID
    updatePermissionRequest := *openapiclient.NewUpdatePermissionRequest() // UpdatePermissionRequest | 許可された Operator のリストに追加すべき Operator の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.AddPermissionToLoraNetworkSet(context.Background(), nsId).UpdatePermissionRequest(updatePermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.AddPermissionToLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddPermissionToLoraNetworkSet`: LoraNetworkSet
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.AddPermissionToLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddPermissionToLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePermissionRequest** | [**UpdatePermissionRequest**](UpdatePermissionRequest.md) | 許可された Operator のリストに追加すべき Operator の ID | 

### Return type

[**LoraNetworkSet**](LoraNetworkSet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoraNetworkSet

> LoraNetworkSet CreateLoraNetworkSet(ctx).LoraNetworkSet(loraNetworkSet).Execute()

Creates a LoRa network set



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
    loraNetworkSet := *openapiclient.NewLoraNetworkSet() // LoraNetworkSet | タグなど作成対象の LoRa network set に設定する付加情報

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.CreateLoraNetworkSet(context.Background()).LoraNetworkSet(loraNetworkSet).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.CreateLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoraNetworkSet`: LoraNetworkSet
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.CreateLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **loraNetworkSet** | [**LoraNetworkSet**](LoraNetworkSet.md) | タグなど作成対象の LoRa network set に設定する付加情報 | 

### Return type

[**LoraNetworkSet**](LoraNetworkSet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoraNetworkSet

> DeleteLoraNetworkSet(ctx, nsId).Execute()

Delete LoRa network set.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.DeleteLoraNetworkSet(context.Background(), nsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.DeleteLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoraNetworkSetRequest struct via the builder pattern


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


## DeleteLoraNetworkSetTag

> DeleteLoraNetworkSetTag(ctx, nsId, tagName).Execute()

Delete LoRa network set tag.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID
    tagName := "tagName_example" // string | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.DeleteLoraNetworkSetTag(context.Background(), nsId, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.DeleteLoraNetworkSetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 
**tagName** | **string** | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoraNetworkSetTagRequest struct via the builder pattern


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


## GetLoraNetworkSet

> LoraNetworkSet GetLoraNetworkSet(ctx, nsId).Execute()

Get LoRa network set.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.GetLoraNetworkSet(context.Background(), nsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.GetLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoraNetworkSet`: LoraNetworkSet
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.GetLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraNetworkSet**](LoraNetworkSet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatewaysInLoraNetworkSet

> []LoraGateway ListGatewaysInLoraNetworkSet(ctx, nsId).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List LoRa Gateways in a Network Set.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID
    limit := int32(56) // int32 | 一度に取得する LoRa gateway の数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の LoRa gateway の ID。このパラメータを指定することで次の LoRa gateway 以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.ListGatewaysInLoraNetworkSet(context.Background(), nsId).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.ListGatewaysInLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGatewaysInLoraNetworkSet`: []LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.ListGatewaysInLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGatewaysInLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | 一度に取得する LoRa gateway の数の上限 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の LoRa gateway の ID。このパラメータを指定することで次の LoRa gateway 以降のリストを取得できる。 | 

### Return type

[**[]LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoraNetworkSets

> []LoraNetworkSet ListLoraNetworkSets(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List LoRa Network Sets.



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
    tagName := "tagName_example" // string | 検索対象にするタグの名前(完全一致)。 (optional)
    tagValue := "tagValue_example" // string | 検索対象にするタグの検索文字列。`tag_name` を指定した場合は必須。 (optional)
    tagValueMatchMode := "tagValueMatchMode_example" // string | タグの検索条件。 (optional) (default to "exact")
    limit := int32(56) // int32 | 一度に取得する LoRa network set の数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の LoRa network set の ID。このパラメータを指定することで次以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.ListLoraNetworkSets(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.ListLoraNetworkSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoraNetworkSets`: []LoraNetworkSet
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.ListLoraNetworkSets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoraNetworkSetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | 検索対象にするタグの名前(完全一致)。 | 
 **tagValue** | **string** | 検索対象にするタグの検索文字列。&#x60;tag_name&#x60; を指定した場合は必須。 | 
 **tagValueMatchMode** | **string** | タグの検索条件。 | [default to &quot;exact&quot;]
 **limit** | **int32** | 一度に取得する LoRa network set の数の上限 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の LoRa network set の ID。このパラメータを指定することで次以降のリストを取得できる。 | 

### Return type

[**[]LoraNetworkSet**](LoraNetworkSet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLoraNetworkSetTags

> LoraNetworkSet PutLoraNetworkSetTags(ctx, nsId).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update LoRa network set tags.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.PutLoraNetworkSetTags(context.Background(), nsId).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.PutLoraNetworkSetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLoraNetworkSetTags`: LoraNetworkSet
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.PutLoraNetworkSetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLoraNetworkSetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**LoraNetworkSet**](LoraNetworkSet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokePermissionFromLoraNetworkSet

> LoraNetworkSet RevokePermissionFromLoraNetworkSet(ctx, nsId).UpdatePermissionRequest(updatePermissionRequest).Execute()

Revokes a permission from a LoRa network set.



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
    nsId := "nsId_example" // string | 対象の LoRa network set の ID
    updatePermissionRequest := *openapiclient.NewUpdatePermissionRequest() // UpdatePermissionRequest | 許可された Operator のリストに追加すべき Operator の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraNetworkSetApi.RevokePermissionFromLoraNetworkSet(context.Background(), nsId).UpdatePermissionRequest(updatePermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraNetworkSetApi.RevokePermissionFromLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevokePermissionFromLoraNetworkSet`: LoraNetworkSet
    fmt.Fprintf(os.Stdout, "Response from `LoraNetworkSetApi.RevokePermissionFromLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**nsId** | **string** | 対象の LoRa network set の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokePermissionFromLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePermissionRequest** | [**UpdatePermissionRequest**](UpdatePermissionRequest.md) | 許可された Operator のリストに追加すべき Operator の ID | 

### Return type

[**LoraNetworkSet**](LoraNetworkSet.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

