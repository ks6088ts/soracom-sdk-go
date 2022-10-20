# \LoraDeviceApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLoraDeviceTag**](LoraDeviceApi.md#DeleteLoraDeviceTag) | **Delete** /lora_devices/{device_id}/tags/{tag_name} | Delete LoRa device Tag
[**DisableTerminationOnLoraDevice**](LoraDeviceApi.md#DisableTerminationOnLoraDevice) | **Post** /lora_devices/{device_id}/disable_termination | Disable Termination of LoRa device
[**EnableTerminationOnLoraDevice**](LoraDeviceApi.md#EnableTerminationOnLoraDevice) | **Post** /lora_devices/{device_id}/enable_termination | Enable Termination of LoRa device
[**GetDataFromLoraDevice**](LoraDeviceApi.md#GetDataFromLoraDevice) | **Get** /lora_devices/{device_id}/data | Get data sent from a LoRa device.
[**GetLoraDevice**](LoraDeviceApi.md#GetLoraDevice) | **Get** /lora_devices/{device_id} | Get LoRa device
[**ListLoraDevices**](LoraDeviceApi.md#ListLoraDevices) | **Get** /lora_devices | List LoRa devices
[**PutLoraDeviceTags**](LoraDeviceApi.md#PutLoraDeviceTags) | **Put** /lora_devices/{device_id}/tags | Bulk Insert or Update LoRa device Tags
[**RegisterLoraDevice**](LoraDeviceApi.md#RegisterLoraDevice) | **Post** /lora_devices/{device_id}/register | Register LoRa device
[**SendDataToLoraDevice**](LoraDeviceApi.md#SendDataToLoraDevice) | **Post** /lora_devices/{device_id}/data | Send data to a LoRa device.
[**SetLoraDeviceGroup**](LoraDeviceApi.md#SetLoraDeviceGroup) | **Post** /lora_devices/{device_id}/set_group | Set Group to LoRa device
[**TerminateLoraDevice**](LoraDeviceApi.md#TerminateLoraDevice) | **Post** /lora_devices/{device_id}/terminate | Terminate LoRa device
[**UnsetLoraDeviceGroup**](LoraDeviceApi.md#UnsetLoraDeviceGroup) | **Post** /lora_devices/{device_id}/unset_group | Unset Group to LoRa device



## DeleteLoraDeviceTag

> DeleteLoraDeviceTag(ctx, deviceId, tagName).Execute()

Delete LoRa device Tag



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID
    tagName := "tagName_example" // string | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.DeleteLoraDeviceTag(context.Background(), deviceId, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.DeleteLoraDeviceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 
**tagName** | **string** | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoraDeviceTagRequest struct via the builder pattern


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


## DisableTerminationOnLoraDevice

> LoraDevice DisableTerminationOnLoraDevice(ctx, deviceId).Execute()

Disable Termination of LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.DisableTerminationOnLoraDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.DisableTerminationOnLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableTerminationOnLoraDevice`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.DisableTerminationOnLoraDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTerminationOnLoraDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTerminationOnLoraDevice

> LoraDevice EnableTerminationOnLoraDevice(ctx, deviceId).Execute()

Enable Termination of LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.EnableTerminationOnLoraDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.EnableTerminationOnLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTerminationOnLoraDevice`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.EnableTerminationOnLoraDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTerminationOnLoraDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataFromLoraDevice

> []DataEntry GetDataFromLoraDevice(ctx, deviceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get data sent from a LoRa device.



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
    deviceId := "deviceId_example" // string | データ取得対象の LoRa デバイスの ID
    from := int32(56) // int32 | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) (optional)
    to := int32(56) // int32 | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) (optional)
    sort := "sort_example" // string | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 (optional) (default to "desc")
    limit := int32(56) // int32 | 取得するデータエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.GetDataFromLoraDevice(context.Background(), deviceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.GetDataFromLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataFromLoraDevice`: []DataEntry
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.GetDataFromLoraDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | データ取得対象の LoRa デバイスの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFromLoraDeviceRequest struct via the builder pattern


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


## GetLoraDevice

> LoraDevice GetLoraDevice(ctx, deviceId).Execute()

Get LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.GetLoraDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.GetLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoraDevice`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.GetLoraDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoraDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoraDevices

> []LoraDevice ListLoraDevices(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List LoRa devices



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
    tagName := "tagName_example" // string | 検索対象にするタグの名前(完全一致) (optional)
    tagValue := "tagValue_example" // string | 検索対象にするタグの検索文字列。`tag_name` を指定した場合は必須 (optional)
    tagValueMatchMode := "tagValueMatchMode_example" // string | タグの検索条件 (optional) (default to "exact")
    limit := int32(56) // int32 | 取得する LoRa device の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の LoRa device の Device ID。このパラメータを指定することで次の LoRa device 以降のリストを取得できる (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.ListLoraDevices(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.ListLoraDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoraDevices`: []LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.ListLoraDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoraDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | 検索対象にするタグの名前(完全一致) | 
 **tagValue** | **string** | 検索対象にするタグの検索文字列。&#x60;tag_name&#x60; を指定した場合は必須 | 
 **tagValueMatchMode** | **string** | タグの検索条件 | [default to &quot;exact&quot;]
 **limit** | **int32** | 取得する LoRa device の上限 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の LoRa device の Device ID。このパラメータを指定することで次の LoRa device 以降のリストを取得できる | 

### Return type

[**[]LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLoraDeviceTags

> LoraDevice PutLoraDeviceTags(ctx, deviceId).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update LoRa device Tags



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.PutLoraDeviceTags(context.Background(), deviceId).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.PutLoraDeviceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutLoraDeviceTags`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.PutLoraDeviceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLoraDeviceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterLoraDevice

> LoraDevice RegisterLoraDevice(ctx, deviceId).RegisterLoraDeviceRequest(registerLoraDeviceRequest).Execute()

Register LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の Device ID
    registerLoraDeviceRequest := *openapiclient.NewRegisterLoraDeviceRequest() // RegisterLoraDeviceRequest | LoRa device

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.RegisterLoraDevice(context.Background(), deviceId).RegisterLoraDeviceRequest(registerLoraDeviceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.RegisterLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterLoraDevice`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.RegisterLoraDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterLoraDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerLoraDeviceRequest** | [**RegisterLoraDeviceRequest**](RegisterLoraDeviceRequest.md) | LoRa device | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDataToLoraDevice

> SendDataToLoraDevice(ctx, deviceId).LoraData(loraData).Execute()

Send data to a LoRa device.



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
    deviceId := "deviceId_example" // string | 送信先デバイスの ID
    loraData := *openapiclient.NewLoraData() // LoraData | 16 進数でエンコードされたバイナリデータの文字列。バイナリデータの最大サイズは 11 octets（16 進数エンコードされた文字列で 22 文字）。fPort を任意で指定できます。fPort は 0 以上でなければなりません。0 はコントロールプレーンに使用され、通常は 1 以上を使用します。無効な値が与えられた場合には特定ベンダーのデバイスで問題が起こらないように設定されたデフォルト値である 2 が使用されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.SendDataToLoraDevice(context.Background(), deviceId).LoraData(loraData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.SendDataToLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 送信先デバイスの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendDataToLoraDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **loraData** | [**LoraData**](LoraData.md) | 16 進数でエンコードされたバイナリデータの文字列。バイナリデータの最大サイズは 11 octets（16 進数エンコードされた文字列で 22 文字）。fPort を任意で指定できます。fPort は 0 以上でなければなりません。0 はコントロールプレーンに使用され、通常は 1 以上を使用します。無効な値が与えられた場合には特定ベンダーのデバイスで問題が起こらないように設定されたデフォルト値である 2 が使用されます。 | 

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


## SetLoraDeviceGroup

> LoraDevice SetLoraDeviceGroup(ctx, deviceId).Group(group).Execute()

Set Group to LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID
    group := *openapiclient.NewGroup() // Group | Group（ID のみを含めばよい）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.SetLoraDeviceGroup(context.Background(), deviceId).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.SetLoraDeviceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetLoraDeviceGroup`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.SetLoraDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetLoraDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) | Group（ID のみを含めばよい） | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateLoraDevice

> LoraDevice TerminateLoraDevice(ctx, deviceId).Execute()

Terminate LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.TerminateLoraDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.TerminateLoraDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateLoraDevice`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.TerminateLoraDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateLoraDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetLoraDeviceGroup

> LoraDevice UnsetLoraDeviceGroup(ctx, deviceId).Execute()

Unset Group to LoRa device



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
    deviceId := "deviceId_example" // string | 対象の LoRa device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LoraDeviceApi.UnsetLoraDeviceGroup(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoraDeviceApi.UnsetLoraDeviceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetLoraDeviceGroup`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `LoraDeviceApi.UnsetLoraDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の LoRa device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetLoraDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

