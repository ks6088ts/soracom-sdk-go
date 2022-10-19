# \SoraCamApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportSoraCamDeviceRecordedImage**](SoraCamApi.md#ExportSoraCamDeviceRecordedImage) | **Post** /sora_cam/devices/{device_id}/images/exports | クラウド常時録画で保存された録画映像から静止画をエクスポートする処理を開始する
[**ExportSoraCamDeviceRecordedVideo**](SoraCamApi.md#ExportSoraCamDeviceRecordedVideo) | **Post** /sora_cam/devices/{device_id}/videos/exports | クラウド常時録画で保存された録画映像をエクスポートする処理を開始する
[**GetSoraCamDevice**](SoraCamApi.md#GetSoraCamDevice) | **Get** /sora_cam/devices/{device_id} | ソラカメ対応カメラの情報を取得する
[**GetSoraCamDeviceExportUsage**](SoraCamApi.md#GetSoraCamDeviceExportUsage) | **Get** /sora_cam/devices/{device_id}/exports/usage | ソラカメ対応カメラの静止画のエクスポート可能枚数や録画映像のエクスポート可能時間を取得する
[**GetSoraCamDeviceExportedImage**](SoraCamApi.md#GetSoraCamDeviceExportedImage) | **Get** /sora_cam/devices/{device_id}/images/exports/{export_id} | クラウド常時録画で保存された録画映像から静止画をエクスポートする処理の現在の状況を取得する
[**GetSoraCamDeviceExportedVideo**](SoraCamApi.md#GetSoraCamDeviceExportedVideo) | **Get** /sora_cam/devices/{device_id}/videos/exports/{export_id} | クラウド常時録画で保存された録画映像をエクスポートする処理の現在の状況を取得する
[**GetSoraCamDeviceStreamingVideo**](SoraCamApi.md#GetSoraCamDeviceStreamingVideo) | **Get** /sora_cam/devices/{device_id}/stream | ストリーミング映像 (リアルタイム映像 / 録画映像) をダウンロードするための情報を取得する
[**ListSoraCamDeviceImageExports**](SoraCamApi.md#ListSoraCamDeviceImageExports) | **Get** /sora_cam/devices/images/exports | ソラカメ対応カメラで撮影した録画映像から静止画をエクスポートする処理の現在の状況をリストで取得する
[**ListSoraCamDeviceImageExportsForDevice**](SoraCamApi.md#ListSoraCamDeviceImageExportsForDevice) | **Get** /sora_cam/devices/{device_id}/images/exports | ソラカメ対応カメラで撮影した録画映像から静止画をエクスポートする処理の現在の状況をリストで取得する
[**ListSoraCamDeviceVideoExports**](SoraCamApi.md#ListSoraCamDeviceVideoExports) | **Get** /sora_cam/devices/videos/exports | ソラカメ対応カメラで撮影した録画映像をエクスポートする処理の現在の状況をリストで取得する
[**ListSoraCamDeviceVideoExportsForDevice**](SoraCamApi.md#ListSoraCamDeviceVideoExportsForDevice) | **Get** /sora_cam/devices/{device_id}/videos/exports | ソラカメ対応カメラで撮影した録画映像をエクスポートする処理の現在の状況をリストで取得する
[**ListSoraCamDevices**](SoraCamApi.md#ListSoraCamDevices) | **Get** /sora_cam/devices | ソラカメ対応カメラの一覧を取得する
[**ListSoraCamLicensePacks**](SoraCamApi.md#ListSoraCamLicensePacks) | **Get** /sora_cam/license_packs | Soracom Cloud Camera Services のライセンスパックの一覧を取得します。
[**UpdateSoraCamLicensePackQuantity**](SoraCamApi.md#UpdateSoraCamLicensePackQuantity) | **Put** /sora_cam/license_packs/{license_pack_id}/quantity | Soracom Cloud Camera Services のライセンス数を更新します。



## ExportSoraCamDeviceRecordedImage

> SoraCamImageExportInfo ExportSoraCamDeviceRecordedImage(ctx, deviceId).SoraCamImageExportSpecification(soraCamImageExportSpecification).Execute()

クラウド常時録画で保存された録画映像から静止画をエクスポートする処理を開始する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    soraCamImageExportSpecification := *openapiclient.NewSoraCamImageExportSpecification(int64(123)) // SoraCamImageExportSpecification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ExportSoraCamDeviceRecordedImage(context.Background(), deviceId).SoraCamImageExportSpecification(soraCamImageExportSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ExportSoraCamDeviceRecordedImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSoraCamDeviceRecordedImage`: SoraCamImageExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ExportSoraCamDeviceRecordedImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSoraCamDeviceRecordedImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **soraCamImageExportSpecification** | [**SoraCamImageExportSpecification**](SoraCamImageExportSpecification.md) |  | 

### Return type

[**SoraCamImageExportInfo**](SoraCamImageExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSoraCamDeviceRecordedVideo

> SoraCamVideoExportInfo ExportSoraCamDeviceRecordedVideo(ctx, deviceId).SoraCamVideoExportSpecification(soraCamVideoExportSpecification).Execute()

クラウド常時録画で保存された録画映像をエクスポートする処理を開始する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    soraCamVideoExportSpecification := *openapiclient.NewSoraCamVideoExportSpecification(int64(123), int64(123)) // SoraCamVideoExportSpecification | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ExportSoraCamDeviceRecordedVideo(context.Background(), deviceId).SoraCamVideoExportSpecification(soraCamVideoExportSpecification).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ExportSoraCamDeviceRecordedVideo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSoraCamDeviceRecordedVideo`: SoraCamVideoExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ExportSoraCamDeviceRecordedVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportSoraCamDeviceRecordedVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **soraCamVideoExportSpecification** | [**SoraCamVideoExportSpecification**](SoraCamVideoExportSpecification.md) |  | 

### Return type

[**SoraCamVideoExportInfo**](SoraCamVideoExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoraCamDevice

> SoraCamDevice GetSoraCamDevice(ctx, deviceId).Execute()

ソラカメ対応カメラの情報を取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.GetSoraCamDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.GetSoraCamDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoraCamDevice`: SoraCamDevice
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.GetSoraCamDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraCamDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoraCamDevice**](SoraCamDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoraCamDeviceExportUsage

> SoraCamExportUsage GetSoraCamDeviceExportUsage(ctx, deviceId).Execute()

ソラカメ対応カメラの静止画のエクスポート可能枚数や録画映像のエクスポート可能時間を取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.GetSoraCamDeviceExportUsage(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.GetSoraCamDeviceExportUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoraCamDeviceExportUsage`: SoraCamExportUsage
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.GetSoraCamDeviceExportUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraCamDeviceExportUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SoraCamExportUsage**](SoraCamExportUsage.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoraCamDeviceExportedImage

> SoraCamImageExportInfo GetSoraCamDeviceExportedImage(ctx, deviceId, exportId).Execute()

クラウド常時録画で保存された録画映像から静止画をエクスポートする処理の現在の状況を取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    exportId := "exportId_example" // string | 対象のエクスポート処理のエクスポート ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.GetSoraCamDeviceExportedImage(context.Background(), deviceId, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.GetSoraCamDeviceExportedImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoraCamDeviceExportedImage`: SoraCamImageExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.GetSoraCamDeviceExportedImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 
**exportId** | **string** | 対象のエクスポート処理のエクスポート ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraCamDeviceExportedImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SoraCamImageExportInfo**](SoraCamImageExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoraCamDeviceExportedVideo

> SoraCamVideoExportInfo GetSoraCamDeviceExportedVideo(ctx, deviceId, exportId).Execute()

クラウド常時録画で保存された録画映像をエクスポートする処理の現在の状況を取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    exportId := "exportId_example" // string | 対象のエクスポート処理のエクスポート ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.GetSoraCamDeviceExportedVideo(context.Background(), deviceId, exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.GetSoraCamDeviceExportedVideo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoraCamDeviceExportedVideo`: SoraCamVideoExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.GetSoraCamDeviceExportedVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 
**exportId** | **string** | 対象のエクスポート処理のエクスポート ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraCamDeviceExportedVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SoraCamVideoExportInfo**](SoraCamVideoExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSoraCamDeviceStreamingVideo

> SoraCamStreamingVideo GetSoraCamDeviceStreamingVideo(ctx, deviceId).From(from).To(to).Execute()

ストリーミング映像 (リアルタイム映像 / 録画映像) をダウンロードするための情報を取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    from := int64(789) // int64 | 録画映像の開始時刻 (UNIX 時間 (ミリ秒))  - リアルタイム映像をダウンロードするための情報を取得するには、`from` と `to` をどちらも省略します。  (optional)
    to := int64(789) // int64 | 録画映像の終了時刻 (UNIX 時間 (ミリ秒))  - リアルタイム映像をダウンロードするための情報を取得するには、`from` と `to` をどちらも省略します。 - 一回の API 呼び出しで視聴できる時間は最大 300 秒 (5 分) です。`from` と `to` の差が、300 秒を超えないようにしてください。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.GetSoraCamDeviceStreamingVideo(context.Background(), deviceId).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.GetSoraCamDeviceStreamingVideo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSoraCamDeviceStreamingVideo`: SoraCamStreamingVideo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.GetSoraCamDeviceStreamingVideo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSoraCamDeviceStreamingVideoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int64** | 録画映像の開始時刻 (UNIX 時間 (ミリ秒))  - リアルタイム映像をダウンロードするための情報を取得するには、&#x60;from&#x60; と &#x60;to&#x60; をどちらも省略します。  | 
 **to** | **int64** | 録画映像の終了時刻 (UNIX 時間 (ミリ秒))  - リアルタイム映像をダウンロードするための情報を取得するには、&#x60;from&#x60; と &#x60;to&#x60; をどちらも省略します。 - 一回の API 呼び出しで視聴できる時間は最大 300 秒 (5 分) です。&#x60;from&#x60; と &#x60;to&#x60; の差が、300 秒を超えないようにしてください。  | 

### Return type

[**SoraCamStreamingVideo**](SoraCamStreamingVideo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraCamDeviceImageExports

> []SoraCamImageExportInfo ListSoraCamDeviceImageExports(ctx).DeviceId(deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

ソラカメ対応カメラで撮影した録画映像から静止画をエクスポートする処理の現在の状況をリストで取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID。省略した場合はオペレーターが所有するすべてのデバイスが対象となります。 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 (optional)
    limit := int32(56) // int32 | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ListSoraCamDeviceImageExports(context.Background()).DeviceId(deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ListSoraCamDeviceImageExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraCamDeviceImageExports`: []SoraCamImageExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ListSoraCamDeviceImageExports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSoraCamDeviceImageExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID。省略した場合はオペレーターが所有するすべてのデバイスが対象となります。 | 
 **lastEvaluatedKey** | **string** | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 | 
 **limit** | **int32** | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 | [default to 10]

### Return type

[**[]SoraCamImageExportInfo**](SoraCamImageExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraCamDeviceImageExportsForDevice

> []SoraCamImageExportInfo ListSoraCamDeviceImageExportsForDevice(ctx, deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

ソラカメ対応カメラで撮影した録画映像から静止画をエクスポートする処理の現在の状況をリストで取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 (optional)
    limit := int32(56) // int32 | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ListSoraCamDeviceImageExportsForDevice(context.Background(), deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ListSoraCamDeviceImageExportsForDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraCamDeviceImageExportsForDevice`: []SoraCamImageExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ListSoraCamDeviceImageExportsForDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSoraCamDeviceImageExportsForDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastEvaluatedKey** | **string** | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 | 
 **limit** | **int32** | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 | [default to 10]

### Return type

[**[]SoraCamImageExportInfo**](SoraCamImageExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraCamDeviceVideoExports

> []SoraCamVideoExportInfo ListSoraCamDeviceVideoExports(ctx).DeviceId(deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

ソラカメ対応カメラで撮影した録画映像をエクスポートする処理の現在の状況をリストで取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID。省略した場合はオペレーターが所有するすべてのデバイスが対象となります。 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 (optional)
    limit := int32(56) // int32 | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ListSoraCamDeviceVideoExports(context.Background()).DeviceId(deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ListSoraCamDeviceVideoExports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraCamDeviceVideoExports`: []SoraCamVideoExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ListSoraCamDeviceVideoExports`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSoraCamDeviceVideoExportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID。省略した場合はオペレーターが所有するすべてのデバイスが対象となります。 | 
 **lastEvaluatedKey** | **string** | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 | 
 **limit** | **int32** | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 | [default to 10]

### Return type

[**[]SoraCamVideoExportInfo**](SoraCamVideoExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraCamDeviceVideoExportsForDevice

> []SoraCamVideoExportInfo ListSoraCamDeviceVideoExportsForDevice(ctx, deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

ソラカメ対応カメラで撮影した録画映像をエクスポートする処理の現在の状況をリストで取得する



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
    deviceId := "deviceId_example" // string | 対象のソラカメ対応カメラのデバイス ID
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 (optional)
    limit := int32(56) // int32 | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ListSoraCamDeviceVideoExportsForDevice(context.Background(), deviceId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ListSoraCamDeviceVideoExportsForDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraCamDeviceVideoExportsForDevice`: []SoraCamVideoExportInfo
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ListSoraCamDeviceVideoExportsForDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象のソラカメ対応カメラのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSoraCamDeviceVideoExportsForDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastEvaluatedKey** | **string** | 前回のリクエストのレスポンスに含まれていた x-soracom-next-key ヘッダーの値。このパラメーターを指定することで、前回のリクエストの続きからリストを取得できます。 | 
 **limit** | **int32** | 一回のリクエストで取得するエクスポート処理に関するデータの個数の最大値。なお、取得できるデータの個数は、指定した数を下回る可能性があります。 | [default to 10]

### Return type

[**[]SoraCamVideoExportInfo**](SoraCamVideoExportInfo.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraCamDevices

> []SoraCamDevice ListSoraCamDevices(ctx).Execute()

ソラカメ対応カメラの一覧を取得する



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ListSoraCamDevices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ListSoraCamDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraCamDevices`: []SoraCamDevice
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ListSoraCamDevices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSoraCamDevicesRequest struct via the builder pattern


### Return type

[**[]SoraCamDevice**](SoraCamDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSoraCamLicensePacks

> []SoraCamLicensePackResponse ListSoraCamLicensePacks(ctx).Execute()

Soracom Cloud Camera Services のライセンスパックの一覧を取得します。



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.ListSoraCamLicensePacks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.ListSoraCamLicensePacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSoraCamLicensePacks`: []SoraCamLicensePackResponse
    fmt.Fprintf(os.Stdout, "Response from `SoraCamApi.ListSoraCamLicensePacks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSoraCamLicensePacksRequest struct via the builder pattern


### Return type

[**[]SoraCamLicensePackResponse**](SoraCamLicensePackResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSoraCamLicensePackQuantity

> UpdateSoraCamLicensePackQuantity(ctx, licensePackId).SoraCamLicensePackQuantityUpdatingRequest(soraCamLicensePackQuantityUpdatingRequest).Execute()

Soracom Cloud Camera Services のライセンス数を更新します。



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
    licensePackId := "licensePackId_example" // string | ライセンスパックの ID
    soraCamLicensePackQuantityUpdatingRequest := *openapiclient.NewSoraCamLicensePackQuantityUpdatingRequest() // SoraCamLicensePackQuantityUpdatingRequest | ライセンス数の更新内容

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SoraCamApi.UpdateSoraCamLicensePackQuantity(context.Background(), licensePackId).SoraCamLicensePackQuantityUpdatingRequest(soraCamLicensePackQuantityUpdatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SoraCamApi.UpdateSoraCamLicensePackQuantity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**licensePackId** | **string** | ライセンスパックの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSoraCamLicensePackQuantityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **soraCamLicensePackQuantityUpdatingRequest** | [**SoraCamLicensePackQuantityUpdatingRequest**](SoraCamLicensePackQuantityUpdatingRequest.md) | ライセンス数の更新内容 | 

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

