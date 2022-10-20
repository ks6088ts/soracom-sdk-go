# \GadgetApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGadgetTag**](GadgetApi.md#DeleteGadgetTag) | **Delete** /gadgets/{product_id}/{serial_number}/tags/{tag_name} | Delete Gadget Tag
[**DisableTerminationOnGadget**](GadgetApi.md#DisableTerminationOnGadget) | **Post** /gadgets/{product_id}/{serial_number}/disable_termination | Disable Termination of Gadget
[**EnableTerminationOnGadget**](GadgetApi.md#EnableTerminationOnGadget) | **Post** /gadgets/{product_id}/{serial_number}/enable_termination | Enable Termination of Gadget
[**GetGadget**](GadgetApi.md#GetGadget) | **Get** /gadgets/{product_id}/{serial_number} | Get Gadget
[**ListGadgets**](GadgetApi.md#ListGadgets) | **Get** /gadgets | List Gadgets
[**PutGadgetTags**](GadgetApi.md#PutGadgetTags) | **Put** /gadgets/{product_id}/{serial_number}/tags | Bulk Insert or Update Gadget Tags
[**RegisterGadget**](GadgetApi.md#RegisterGadget) | **Post** /gadgets/{product_id}/{serial_number}/register | Register a Gadget.
[**SetGadgetGroup**](GadgetApi.md#SetGadgetGroup) | **Post** /gadgets/{product_id}/{serial_number}/set_group | Set Group to Gadget
[**TerminateGadget**](GadgetApi.md#TerminateGadget) | **Post** /gadgets/{product_id}/{serial_number}/terminate | Terminate Gadget
[**UnsetGadgetGroup**](GadgetApi.md#UnsetGadgetGroup) | **Post** /gadgets/{product_id}/{serial_number}/unset_group | Unset Group to Gadget



## DeleteGadgetTag

> DeleteGadgetTag(ctx, productId, serialNumber, tagName).Execute()

Delete Gadget Tag



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number
    tagName := "tagName_example" // string | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.DeleteGadgetTag(context.Background(), productId, serialNumber, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.DeleteGadgetTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 
**tagName** | **string** | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGadgetTagRequest struct via the builder pattern


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


## DisableTerminationOnGadget

> Gadget DisableTerminationOnGadget(ctx, productId, serialNumber).Execute()

Disable Termination of Gadget



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.DisableTerminationOnGadget(context.Background(), productId, serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.DisableTerminationOnGadget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableTerminationOnGadget`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.DisableTerminationOnGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTerminationOnGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTerminationOnGadget

> Gadget EnableTerminationOnGadget(ctx, productId, serialNumber).Execute()

Enable Termination of Gadget



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.EnableTerminationOnGadget(context.Background(), productId, serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.EnableTerminationOnGadget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTerminationOnGadget`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.EnableTerminationOnGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTerminationOnGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGadget

> Gadget GetGadget(ctx, productId, serialNumber).Execute()

Get Gadget



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.GetGadget(context.Background(), productId, serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.GetGadget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGadget`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.GetGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGadgets

> []Gadget ListGadgets(ctx).ProductId(productId).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List Gadgets



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
    productId := "productId_example" // string | 対象の Gadget の Product ID (optional)
    tagName := "tagName_example" // string | 検索対象にするタグの名前(完全一致) (optional)
    tagValue := "tagValue_example" // string | 検索対象にするタグの検索文字列。`tag_name` を指定した場合は必須 (optional)
    tagValueMatchMode := "tagValueMatchMode_example" // string | タグの検索条件 (optional) (default to "exact")
    limit := int32(56) // int32 | 取得する Gadget の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の Gadget の ID ({product_id}/{serial_number})。このパラメータを指定することで次の Gadget 以降のリストを取得できる (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.ListGadgets(context.Background()).ProductId(productId).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.ListGadgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGadgets`: []Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.ListGadgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGadgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **productId** | **string** | 対象の Gadget の Product ID | 
 **tagName** | **string** | 検索対象にするタグの名前(完全一致) | 
 **tagValue** | **string** | 検索対象にするタグの検索文字列。&#x60;tag_name&#x60; を指定した場合は必須 | 
 **tagValueMatchMode** | **string** | タグの検索条件 | [default to &quot;exact&quot;]
 **limit** | **int32** | 取得する Gadget の上限 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の Gadget の ID ({product_id}/{serial_number})。このパラメータを指定することで次の Gadget 以降のリストを取得できる | 

### Return type

[**[]Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutGadgetTags

> Gadget PutGadgetTags(ctx, productId, serialNumber).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update Gadget Tags



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.PutGadgetTags(context.Background(), productId, serialNumber).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.PutGadgetTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutGadgetTags`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.PutGadgetTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutGadgetTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterGadget

> Gadget RegisterGadget(ctx, productId, serialNumber).GadgetRegistrationRequest(gadgetRegistrationRequest).Execute()

Register a Gadget.



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number
    gadgetRegistrationRequest := *openapiclient.NewGadgetRegistrationRequest() // GadgetRegistrationRequest | Gadget の登録リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.RegisterGadget(context.Background(), productId, serialNumber).GadgetRegistrationRequest(gadgetRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.RegisterGadget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterGadget`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.RegisterGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **gadgetRegistrationRequest** | [**GadgetRegistrationRequest**](GadgetRegistrationRequest.md) | Gadget の登録リクエスト | 

### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGadgetGroup

> Gadget SetGadgetGroup(ctx, productId, serialNumber).Group(group).Execute()

Set Group to Gadget



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number
    group := *openapiclient.NewGroup() // Group | Group（ID のみを含めばよい）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.SetGadgetGroup(context.Background(), productId, serialNumber).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.SetGadgetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGadgetGroup`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.SetGadgetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGadgetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **group** | [**Group**](Group.md) | Group（ID のみを含めばよい） | 

### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateGadget

> Gadget TerminateGadget(ctx, productId, serialNumber).Execute()

Terminate Gadget



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.TerminateGadget(context.Background(), productId, serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.TerminateGadget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateGadget`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.TerminateGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetGadgetGroup

> Gadget UnsetGadgetGroup(ctx, productId, serialNumber).Execute()

Unset Group to Gadget



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
    productId := "productId_example" // string | 対象の Gadget の Product ID
    serialNumber := "serialNumber_example" // string | 対象の Gadget の Serial Number

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GadgetApi.UnsetGadgetGroup(context.Background(), productId, serialNumber).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GadgetApi.UnsetGadgetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetGadgetGroup`: Gadget
    fmt.Fprintf(os.Stdout, "Response from `GadgetApi.UnsetGadgetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**productId** | **string** | 対象の Gadget の Product ID | 
**serialNumber** | **string** | 対象の Gadget の Serial Number | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetGadgetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Gadget**](Gadget.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

