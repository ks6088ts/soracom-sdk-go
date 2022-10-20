# \LoraGatewayApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLoraGatewayTag**](LoraGatewayApi.md#DeleteLoraGatewayTag) | **Delete** /lora_gateways/{gateway_id}/tags/{tag_name} | Delete LoRa gateway tag.
[**DisableTerminationOnLoraGateway**](LoraGatewayApi.md#DisableTerminationOnLoraGateway) | **Post** /lora_gateways/{gateway_id}/disable_termination | Disable Termination of LoRa gateway.
[**EnableTerminationOnLoraGateway**](LoraGatewayApi.md#EnableTerminationOnLoraGateway) | **Post** /lora_gateways/{gateway_id}/enable_termination | Enable Termination of LoRa gateway.
[**GetLoraGateway**](LoraGatewayApi.md#GetLoraGateway) | **Get** /lora_gateways/{gateway_id} | Get LoRa gateway.
[**ListLoraGateways**](LoraGatewayApi.md#ListLoraGateways) | **Get** /lora_gateways | List LoRa gateways.
[**PutLoraGatewayTags**](LoraGatewayApi.md#PutLoraGatewayTags) | **Put** /lora_gateways/{gateway_id}/tags | Bulk Insert or Update LoRa gateway Tags.
[**SetLoraNetworkSet**](LoraGatewayApi.md#SetLoraNetworkSet) | **Post** /lora_gateways/{gateway_id}/set_network_set | Set Network Set ID of LoRa gateway.
[**TerminateLoraGateway**](LoraGatewayApi.md#TerminateLoraGateway) | **Post** /lora_gateways/{gateway_id}/terminate | Terminate LoRa gateway.
[**UnsetLoraNetworkSet**](LoraGatewayApi.md#UnsetLoraNetworkSet) | **Post** /lora_gateways/{gateway_id}/unset_network_set | Unset Network Set ID of LoRa gateway.



## DeleteLoraGatewayTag

> DeleteLoraGatewayTag(ctx, gatewayId, tagName).Execute()

Delete LoRa gateway tag.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID
    tagName := "tagName_example" // string | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.DeleteLoraGatewayTag(context.Background(), gatewayId, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.DeleteLoraGatewayTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 
**tagName** | **string** | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoraGatewayTagRequest struct via the builder pattern


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


## DisableTerminationOnLoraGateway

> LoraGateway DisableTerminationOnLoraGateway(ctx, gatewayId).Execute()

Disable Termination of LoRa gateway.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.DisableTerminationOnLoraGateway(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.DisableTerminationOnLoraGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableTerminationOnLoraGateway`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.DisableTerminationOnLoraGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTerminationOnLoraGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTerminationOnLoraGateway

> LoraGateway EnableTerminationOnLoraGateway(ctx, gatewayId).Execute()

Enable Termination of LoRa gateway.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.EnableTerminationOnLoraGateway(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.EnableTerminationOnLoraGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTerminationOnLoraGateway`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.EnableTerminationOnLoraGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTerminationOnLoraGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLoraGateway

> LoraGateway GetLoraGateway(ctx, gatewayId).Execute()

Get LoRa gateway.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.GetLoraGateway(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.GetLoraGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoraGateway`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.GetLoraGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoraGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoraGateways

> []LoraGateway ListLoraGateways(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List LoRa gateways.



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
    limit := int32(56) // int32 | 一度に取得する LoRa gateway の数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の LoRa gateway の ID。このパラメータを指定することで次の LoRa gateway 以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.ListLoraGateways(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.ListLoraGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoraGateways`: []LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.ListLoraGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoraGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | 検索対象にするタグの名前(完全一致)。 | 
 **tagValue** | **string** | 検索対象にするタグの検索文字列。&#x60;tag_name&#x60; を指定した場合は必須。 | 
 **tagValueMatchMode** | **string** | タグの検索条件。 | [default to &quot;exact&quot;]
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


## PutLoraGatewayTags

> LoraGateway PutLoraGatewayTags(ctx, gatewayId).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update LoRa gateway Tags.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa device の ID
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.PutLoraGatewayTags(context.Background(), gatewayId).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.PutLoraGatewayTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLoraGatewayTags`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.PutLoraGatewayTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLoraGatewayTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetLoraNetworkSet

> LoraGateway SetLoraNetworkSet(ctx, gatewayId).SetNetworkSetRequest(setNetworkSetRequest).Execute()

Set Network Set ID of LoRa gateway.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID
    setNetworkSetRequest := *openapiclient.NewSetNetworkSetRequest() // SetNetworkSetRequest | LoRa Network Set ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.SetLoraNetworkSet(context.Background(), gatewayId).SetNetworkSetRequest(setNetworkSetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.SetLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLoraNetworkSet`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.SetLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setNetworkSetRequest** | [**SetNetworkSetRequest**](SetNetworkSetRequest.md) | LoRa Network Set ID | 

### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateLoraGateway

> LoraGateway TerminateLoraGateway(ctx, gatewayId).Execute()

Terminate LoRa gateway.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.TerminateLoraGateway(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.TerminateLoraGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateLoraGateway`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.TerminateLoraGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateLoraGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetLoraNetworkSet

> LoraGateway UnsetLoraNetworkSet(ctx, gatewayId).Execute()

Unset Network Set ID of LoRa gateway.



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
    gatewayId := "gatewayId_example" // string | 対象の LoRa gateway の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraGatewayApi.UnsetLoraNetworkSet(context.Background(), gatewayId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraGatewayApi.UnsetLoraNetworkSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetLoraNetworkSet`: LoraGateway
    fmt.Fprintf(os.Stdout, "Response from `LoraGatewayApi.UnsetLoraNetworkSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gatewayId** | **string** | 対象の LoRa gateway の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetLoraNetworkSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraGateway**](LoraGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

