# \DeviceApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDevice**](DeviceApi.md#CreateDevice) | **Post** /devices | 新しいデバイスを作成します。
[**DeleteDevice**](DeviceApi.md#DeleteDevice) | **Delete** /devices/{device_id} | デバイスを削除します。
[**DeleteDeviceTag**](DeviceApi.md#DeleteDeviceTag) | **Delete** /devices/{device_id}/tags/{tag_name} | デバイスのタグを削除します。
[**ExecuteDeviceResource**](DeviceApi.md#ExecuteDeviceResource) | **Post** /devices/{device_id}/{object}/{instance}/{resource}/execute | デバイスの Resource を Execute します。
[**GetDataFromDevice**](DeviceApi.md#GetDataFromDevice) | **Get** /devices/{device_id}/data | デバイスから送信されたデータを取得します。
[**GetDevice**](DeviceApi.md#GetDevice) | **Get** /devices/{device_id} | 指定されたデバイス ID のデバイスを取得します。
[**ListDevices**](DeviceApi.md#ListDevices) | **Get** /devices | デバイスのリストを取得します。
[**ObserveDeviceResource**](DeviceApi.md#ObserveDeviceResource) | **Post** /devices/{device_id}/{object}/{instance}/{resource}/observe | 指定されたデバイスの Resource の Observe を開始します。
[**ObserveDeviceResources**](DeviceApi.md#ObserveDeviceResources) | **Post** /devices/{device_id}/{object}/{instance}/observe | 指定されたデバイスの Object Instance 下の Resource 全体の Observe を開始します。
[**PutDeviceTags**](DeviceApi.md#PutDeviceTags) | **Put** /devices/{device_id}/tags | デバイスのタグを更新します。
[**ReadDeviceResource**](DeviceApi.md#ReadDeviceResource) | **Get** /devices/{device_id}/{object}/{instance}/{resource} | 指定されたデバイスの Resource を Read します。
[**ReadDeviceResources**](DeviceApi.md#ReadDeviceResources) | **Get** /devices/{device_id}/{object}/{instance} | デバイスの Instance 以下のすべての Resource を Read します。
[**SetDeviceGroup**](DeviceApi.md#SetDeviceGroup) | **Post** /devices/{device_id}/set_group | デバイスをグループに参加させます。
[**UnobserveDeviceResource**](DeviceApi.md#UnobserveDeviceResource) | **Post** /devices/{device_id}/{object}/{instance}/{resource}/unobserve | デバイスの Resource の Observe を停止します。
[**UnobserveDeviceResources**](DeviceApi.md#UnobserveDeviceResources) | **Post** /devices/{device_id}/{object}/{instance}/unobserve | 指定されたデバイスの Object Instance への Observe を停止します。
[**UnsetDeviceGroup**](DeviceApi.md#UnsetDeviceGroup) | **Post** /devices/{device_id}/unset_group | デバイスをグループから外します。
[**WriteDeviceResource**](DeviceApi.md#WriteDeviceResource) | **Put** /devices/{device_id}/{object}/{instance}/{resource} | デバイスの Resource に値を Write します。



## CreateDevice

> Device CreateDevice(ctx).Device(device).Execute()

新しいデバイスを作成します。



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
    device := *openapiclient.NewDevice() // Device | 作成するデバイス

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.CreateDevice(context.Background()).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.CreateDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.CreateDevice`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device** | [**Device**](Device.md) | 作成するデバイス | 

### Return type

[**Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDevice

> DeleteDevice(ctx, deviceId).Execute()

デバイスを削除します。



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
    deviceId := "deviceId_example" // string | 削除するデバイス

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDevice(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 削除するデバイス | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceRequest struct via the builder pattern


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


## DeleteDeviceTag

> DeleteDeviceTag(ctx, deviceId, tagName).Execute()

デバイスのタグを削除します。



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
    deviceId := "deviceId_example" // string | 更新するデバイスのデバイス ID
    tagName := "tagName_example" // string | 削除するタグの名前

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.DeleteDeviceTag(context.Background(), deviceId, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeleteDeviceTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 更新するデバイスのデバイス ID | 
**tagName** | **string** | 削除するタグの名前 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceTagRequest struct via the builder pattern


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


## ExecuteDeviceResource

> ExecuteDeviceResource(ctx, deviceId, object, instance, resource).ExecuteDeviceResourceRequest(executeDeviceResourceRequest).Execute()

デバイスの Resource を Execute します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    resource := "resource_example" // string | Resource ID
    executeDeviceResourceRequest := *openapiclient.NewExecuteDeviceResourceRequest() // ExecuteDeviceResourceRequest | Resource を実行する際に渡す引数文字列 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ExecuteDeviceResource(context.Background(), deviceId, object, instance, resource).ExecuteDeviceResourceRequest(executeDeviceResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ExecuteDeviceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 
**resource** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteDeviceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **executeDeviceResourceRequest** | [**ExecuteDeviceResourceRequest**](ExecuteDeviceResourceRequest.md) | Resource を実行する際に渡す引数文字列 | 

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


## GetDataFromDevice

> []DataEntry GetDataFromDevice(ctx, deviceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

デバイスから送信されたデータを取得します。



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
    deviceId := "deviceId_example" // string | データ取得対象のデバイスの ID
    from := int32(56) // int32 | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) (optional)
    to := int32(56) // int32 | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) (optional)
    sort := "sort_example" // string | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 (optional) (default to "desc")
    limit := int32(56) // int32 | 取得するデータエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDataFromDevice(context.Background(), deviceId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDataFromDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataFromDevice`: []DataEntry
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDataFromDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | データ取得対象のデバイスの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFromDeviceRequest struct via the builder pattern


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


## GetDevice

> Device GetDevice(ctx, deviceId).Model(model).Execute()

指定されたデバイス ID のデバイスを取得します。



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
    deviceId := "deviceId_example" // string | デバイス ID
    model := true // bool | モデル情報を取得するかどうか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.GetDevice(context.Background(), deviceId).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.GetDevice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDevice`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.GetDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | デバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **model** | **bool** | モデル情報を取得するかどうか | [default to false]

### Return type

[**Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevices

> []Device ListDevices(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

デバイスのリストを取得します。



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
    tagName := "tagName_example" // string | タグの名前 (optional)
    tagValue := "tagValue_example" // string | タグの値 (optional)
    tagValueMatchMode := "tagValueMatchMode_example" // string | タグの検索条件 (exact | prefix) (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のデバイスの ID。このパラメータを指定することで次のデバイスから始まるリストを取得できる。 (optional)
    limit := int32(56) // int32 | 取得するデバイス数の上限 (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ListDevices(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ListDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevices`: []Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ListDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | タグの名前 | 
 **tagValue** | **string** | タグの値 | 
 **tagValueMatchMode** | **string** | タグの検索条件 (exact | prefix) | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後のデバイスの ID。このパラメータを指定することで次のデバイスから始まるリストを取得できる。 | 
 **limit** | **int32** | 取得するデバイス数の上限 | [default to -1]

### Return type

[**[]Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ObserveDeviceResource

> ObserveDeviceResource(ctx, deviceId, object, instance, resource).Model(model).Execute()

指定されたデバイスの Resource の Observe を開始します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    resource := "resource_example" // string | Resource ID
    model := true // bool | モデル情報を含めるかどうか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ObserveDeviceResource(context.Background(), deviceId, object, instance, resource).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ObserveDeviceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 
**resource** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiObserveDeviceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **model** | **bool** | モデル情報を含めるかどうか | [default to false]

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


## ObserveDeviceResources

> ObserveDeviceResources(ctx, deviceId, object, instance).Model(model).Execute()

指定されたデバイスの Object Instance 下の Resource 全体の Observe を開始します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    model := true // bool | モデルの情報を追加するか否か (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ObserveDeviceResources(context.Background(), deviceId, object, instance).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ObserveDeviceResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiObserveDeviceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | **bool** | モデルの情報を追加するか否か | [default to false]

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


## PutDeviceTags

> Device PutDeviceTags(ctx, deviceId).TagUpdateRequest(tagUpdateRequest).Execute()

デバイスのタグを更新します。



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
    deviceId := "deviceId_example" // string | 更新するデバイスのデバイス ID
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.PutDeviceTags(context.Background(), deviceId).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.PutDeviceTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutDeviceTags`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.PutDeviceTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 更新するデバイスのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDeviceTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDeviceResource

> ResourceInstance ReadDeviceResource(ctx, deviceId, object, instance, resource).Model(model).Execute()

指定されたデバイスの Resource を Read します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    resource := "resource_example" // string | Resource ID
    model := true // bool | モデル情報を含めるかどうか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ReadDeviceResource(context.Background(), deviceId, object, instance, resource).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ReadDeviceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDeviceResource`: ResourceInstance
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ReadDeviceResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 
**resource** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDeviceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **model** | **bool** | モデル情報を含めるかどうか | [default to false]

### Return type

[**ResourceInstance**](ResourceInstance.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadDeviceResources

> ObjectInstance ReadDeviceResources(ctx, deviceId, object, instance).Model(model).Execute()

デバイスの Instance 以下のすべての Resource を Read します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    model := true // bool | モデル情報を含めるかどうか (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.ReadDeviceResources(context.Background(), deviceId, object, instance).Model(model).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.ReadDeviceResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadDeviceResources`: ObjectInstance
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.ReadDeviceResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadDeviceResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **model** | **bool** | モデル情報を含めるかどうか | [default to false]

### Return type

[**ObjectInstance**](ObjectInstance.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceGroup

> Device SetDeviceGroup(ctx, deviceId).SetDeviceGroupRequest(setDeviceGroupRequest).Execute()

デバイスをグループに参加させます。



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
    deviceId := "deviceId_example" // string | 更新するデバイスのデバイス ID
    setDeviceGroupRequest := *openapiclient.NewSetDeviceGroupRequest() // SetDeviceGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.SetDeviceGroup(context.Background(), deviceId).SetDeviceGroupRequest(setDeviceGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.SetDeviceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDeviceGroup`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.SetDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 更新するデバイスのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setDeviceGroupRequest** | [**SetDeviceGroupRequest**](SetDeviceGroupRequest.md) |  | 

### Return type

[**Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnobserveDeviceResource

> UnobserveDeviceResource(ctx, deviceId, object, instance, resource).Execute()

デバイスの Resource の Observe を停止します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    resource := "resource_example" // string | Resource ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnobserveDeviceResource(context.Background(), deviceId, object, instance, resource).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnobserveDeviceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 
**resource** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnobserveDeviceResourceRequest struct via the builder pattern


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


## UnobserveDeviceResources

> UnobserveDeviceResources(ctx, deviceId, object, instance).Execute()

指定されたデバイスの Object Instance への Observe を停止します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnobserveDeviceResources(context.Background(), deviceId, object, instance).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnobserveDeviceResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnobserveDeviceResourcesRequest struct via the builder pattern


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


## UnsetDeviceGroup

> Device UnsetDeviceGroup(ctx, deviceId).Execute()

デバイスをグループから外します。



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
    deviceId := "deviceId_example" // string | 更新するデバイスのデバイス ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.UnsetDeviceGroup(context.Background(), deviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.UnsetDeviceGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetDeviceGroup`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.UnsetDeviceGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 更新するデバイスのデバイス ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetDeviceGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WriteDeviceResource

> WriteDeviceResource(ctx, deviceId, object, instance, resource).WriteDeviceResourceRequest(writeDeviceResourceRequest).Execute()

デバイスの Resource に値を Write します。



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
    deviceId := "deviceId_example" // string | 対象となるデバイス
    object := "object_example" // string | Object ID
    instance := "instance_example" // string | Instance ID
    resource := "resource_example" // string | Resource ID
    writeDeviceResourceRequest := *openapiclient.NewWriteDeviceResourceRequest() // WriteDeviceResourceRequest | リソースにセットする値。この値は LwM2M の基本的なデータ型を受け入れます。もしリソースモデルが `multiple` モードの場合はプロパティ名を `value` ではなく `values` にする必要があります。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceApi.WriteDeviceResource(context.Background(), deviceId, object, instance, resource).WriteDeviceResourceRequest(writeDeviceResourceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.WriteDeviceResource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | 対象となるデバイス | 
**object** | **string** | Object ID | 
**instance** | **string** | Instance ID | 
**resource** | **string** | Resource ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiWriteDeviceResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **writeDeviceResourceRequest** | [**WriteDeviceResourceRequest**](WriteDeviceResourceRequest.md) | リソースにセットする値。この値は LwM2M の基本的なデータ型を受け入れます。もしリソースモデルが &#x60;multiple&#x60; モードの場合はプロパティ名を &#x60;value&#x60; ではなく &#x60;values&#x60; にする必要があります。 | 

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

