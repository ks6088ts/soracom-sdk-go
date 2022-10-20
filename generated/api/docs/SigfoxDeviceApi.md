# \SigfoxDeviceApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSigfoxDeviceTag**](SigfoxDeviceApi.md#DeleteSigfoxDeviceTag) | **Delete** /sigfox_devices/{device_id}/tags/{tag_name} | Delete Sigfox device Tag
[**DisableTerminationOnSigfoxDevice**](SigfoxDeviceApi.md#DisableTerminationOnSigfoxDevice) | **Post** /sigfox_devices/{device_id}/disable_termination | Disable Termination of Sigfox device
[**EnableTerminationOnSigfoxDevice**](SigfoxDeviceApi.md#EnableTerminationOnSigfoxDevice) | **Post** /sigfox_devices/{device_id}/enable_termination | Enable Termination of Sigfox device
[**GetDataFromSigfoxDevice**](SigfoxDeviceApi.md#GetDataFromSigfoxDevice) | **Get** /sigfox_devices/{device_id}/data | Get data sent from a Sigfox device.
[**GetSigfoxDevice**](SigfoxDeviceApi.md#GetSigfoxDevice) | **Get** /sigfox_devices/{device_id} | Get Sigfox device
[**ListSigfoxDevices**](SigfoxDeviceApi.md#ListSigfoxDevices) | **Get** /sigfox_devices | List Sigfox devices
[**PutSigfoxDeviceTags**](SigfoxDeviceApi.md#PutSigfoxDeviceTags) | **Put** /sigfox_devices/{device_id}/tags | Bulk Insert or Update Sigfox device Tags
[**RegisterSigfoxDevice**](SigfoxDeviceApi.md#RegisterSigfoxDevice) | **Post** /sigfox_devices/{device_id}/register | Register a Sigfox device.
[**SendDataToSigfoxDevice**](SigfoxDeviceApi.md#SendDataToSigfoxDevice) | **Post** /sigfox_devices/{device_id}/data | Send data to a Sigfox device.
[**SetSigfoxDeviceGroup**](SigfoxDeviceApi.md#SetSigfoxDeviceGroup) | **Post** /sigfox_devices/{device_id}/set_group | Set Group to Sigfox device
[**TerminateSigfoxDevice**](SigfoxDeviceApi.md#TerminateSigfoxDevice) | **Post** /sigfox_devices/{device_id}/terminate | Terminate Sigfox device
[**UnsetSigfoxDeviceGroup**](SigfoxDeviceApi.md#UnsetSigfoxDeviceGroup) | **Post** /sigfox_devices/{device_id}/unset_group | Unset Group to Sigfox device



## DeleteSigfoxDeviceTag

> DeleteSigfoxDeviceTag(ctx, deviceId, tagName).Execute()

Delete Sigfox device Tag



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID
    tagName := "tagName_example" // string | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.DeleteSigfoxDeviceTag(context.Background(), deviceId, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.DeleteSigfoxDeviceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 
**tagName** | **string** | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSigfoxDeviceTagRequest struct via the builder pattern


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


## DisableTerminationOnSigfoxDevice

> SigfoxDevice DisableTerminationOnSigfoxDevice(ctx, deviceId).Execute()

Disable Termination of Sigfox device



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.DisableTerminationOnSigfoxDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.DisableTerminationOnSigfoxDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableTerminationOnSigfoxDevice`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.DisableTerminationOnSigfoxDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTerminationOnSigfoxDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTerminationOnSigfoxDevice

> SigfoxDevice EnableTerminationOnSigfoxDevice(ctx, deviceId).Execute()

Enable Termination of Sigfox device



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.EnableTerminationOnSigfoxDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.EnableTerminationOnSigfoxDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTerminationOnSigfoxDevice`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.EnableTerminationOnSigfoxDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTerminationOnSigfoxDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataFromSigfoxDevice

> []DataEntry GetDataFromSigfoxDevice(ctx, deviceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get data sent from a Sigfox device.



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
    deviceId := "deviceId_example" // string | データ取得対象の Sigfox デバイスの ID
    from := int32(56) // int32 | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) (optional)
    to := int32(56) // int32 | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) (optional)
    sort := "sort_example" // string | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 (optional) (default to "desc")
    limit := int32(56) // int32 | 取得するデータエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.GetDataFromSigfoxDevice(context.Background(), deviceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.GetDataFromSigfoxDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataFromSigfoxDevice`: []DataEntry
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.GetDataFromSigfoxDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | データ取得対象の Sigfox デバイスの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFromSigfoxDeviceRequest struct via the builder pattern


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


## GetSigfoxDevice

> SigfoxDevice GetSigfoxDevice(ctx, deviceId).Execute()

Get Sigfox device



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の Device ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.GetSigfoxDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.GetSigfoxDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSigfoxDevice`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.GetSigfoxDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の Device ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSigfoxDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSigfoxDevices

> []SigfoxDevice ListSigfoxDevices(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List Sigfox devices



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
    limit := int32(56) // int32 | 取得する Sigfox device の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の Sigfox device の Device ID。このパラメータを指定することで次の Sigfox device 以降のリストを取得できる (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.ListSigfoxDevices(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.ListSigfoxDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSigfoxDevices`: []SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.ListSigfoxDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSigfoxDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | 検索対象にするタグの名前(完全一致) | 
 **tagValue** | **string** | 検索対象にするタグの検索文字列。&#x60;tag_name&#x60; を指定した場合は必須 | 
 **tagValueMatchMode** | **string** | タグの検索条件 | [default to &quot;exact&quot;]
 **limit** | **int32** | 取得する Sigfox device の上限 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の Sigfox device の Device ID。このパラメータを指定することで次の Sigfox device 以降のリストを取得できる | 

### Return type

[**[]SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSigfoxDeviceTags

> SigfoxDevice PutSigfoxDeviceTags(ctx, deviceId).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update Sigfox device Tags



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.PutSigfoxDeviceTags(context.Background(), deviceId).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.PutSigfoxDeviceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSigfoxDeviceTags`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.PutSigfoxDeviceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSigfoxDeviceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSigfoxDevice

> SigfoxDevice RegisterSigfoxDevice(ctx, deviceId).SigfoxRegistrationRequest(sigfoxRegistrationRequest).Execute()

Register a Sigfox device.



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
    deviceId := "deviceId_example" // string | 登録対象の Sigfox デバイスの ID
    sigfoxRegistrationRequest := *openapiclient.NewSigfoxRegistrationRequest() // SigfoxRegistrationRequest | Sigfox デバイスの登録リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.RegisterSigfoxDevice(context.Background(), deviceId).SigfoxRegistrationRequest(sigfoxRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.RegisterSigfoxDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterSigfoxDevice`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.RegisterSigfoxDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 登録対象の Sigfox デバイスの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSigfoxDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sigfoxRegistrationRequest** | [**SigfoxRegistrationRequest**](SigfoxRegistrationRequest.md) | Sigfox デバイスの登録リクエスト | 

### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendDataToSigfoxDevice

> SendDataToSigfoxDevice(ctx, deviceId).SigfoxData(sigfoxData).Execute()

Send data to a Sigfox device.



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
    sigfoxData := *openapiclient.NewSigfoxData() // SigfoxData | 16 進数でエンコードされたバイナリデータの文字列。バイナリデータのサイズは 8 octets（16 進数エンコードされた文字列で 16 文字）である必要がある。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.SendDataToSigfoxDevice(context.Background(), deviceId).SigfoxData(sigfoxData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.SendDataToSigfoxDevice``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSendDataToSigfoxDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sigfoxData** | [**SigfoxData**](SigfoxData.md) | 16 進数でエンコードされたバイナリデータの文字列。バイナリデータのサイズは 8 octets（16 進数エンコードされた文字列で 16 文字）である必要がある。 | 

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


## SetSigfoxDeviceGroup

> SigfoxDevice SetSigfoxDeviceGroup(ctx, deviceId).Group(group).Execute()

Set Group to Sigfox device



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID
    group := *openapiclient.NewGroup() // Group | Group（ID のみを含めばよい）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.SetSigfoxDeviceGroup(context.Background(), deviceId).Group(group).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.SetSigfoxDeviceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSigfoxDeviceGroup`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.SetSigfoxDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSigfoxDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **group** | [**Group**](Group.md) | Group（ID のみを含めばよい） | 

### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateSigfoxDevice

> SigfoxDevice TerminateSigfoxDevice(ctx, deviceId).DeleteImmediately(deleteImmediately).Execute()

Terminate Sigfox device



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID
    deleteImmediately := true // bool | Sigfox デバイスを即座に削除する (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.TerminateSigfoxDevice(context.Background(), deviceId).DeleteImmediately(deleteImmediately).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.TerminateSigfoxDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateSigfoxDevice`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.TerminateSigfoxDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateSigfoxDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteImmediately** | **bool** | Sigfox デバイスを即座に削除する | [default to false]

### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetSigfoxDeviceGroup

> SigfoxDevice UnsetSigfoxDeviceGroup(ctx, deviceId).Execute()

Unset Group to Sigfox device



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
    deviceId := "deviceId_example" // string | 対象の Sigfox device の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SigfoxDeviceApi.UnsetSigfoxDeviceGroup(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SigfoxDeviceApi.UnsetSigfoxDeviceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetSigfoxDeviceGroup`: SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `SigfoxDeviceApi.UnsetSigfoxDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象の Sigfox device の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetSigfoxDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

