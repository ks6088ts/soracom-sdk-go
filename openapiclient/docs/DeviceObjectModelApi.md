# \DeviceObjectModelApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceObjectModel**](DeviceObjectModelApi.md#CreateDeviceObjectModel) | **Post** /device_object_models | 新しいデバイスオブジェクトモデルを作成します。
[**DeleteDeviceObjectModel**](DeviceObjectModelApi.md#DeleteDeviceObjectModel) | **Delete** /device_object_models/{model_id} | 対象のデバイスオブジェクトモデルを削除します。
[**GetDeviceObjectModel**](DeviceObjectModelApi.md#GetDeviceObjectModel) | **Get** /device_object_models/{model_id} | デバイスオブジェクトモデルを取得します。
[**ListDeviceObjectModels**](DeviceObjectModelApi.md#ListDeviceObjectModels) | **Get** /device_object_models | デバイスオブジェクトモデルのリストを返します。
[**SetDeviceObjectModelScope**](DeviceObjectModelApi.md#SetDeviceObjectModelScope) | **Post** /device_object_models/{model_id}/set_scope | 対象のデバイスオブジェクトモデルが適用される Scope をセットします。
[**UpdateDeviceObjectModel**](DeviceObjectModelApi.md#UpdateDeviceObjectModel) | **Post** /device_object_models/{model_id} | デバイスオブジェクトモデルを更新します。



## CreateDeviceObjectModel

> DeviceObjectModel CreateDeviceObjectModel(ctx).DeviceObjectModel(deviceObjectModel).Execute()

新しいデバイスオブジェクトモデルを作成します。



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
    deviceObjectModel := *openapiclient.NewDeviceObjectModel() // DeviceObjectModel | Device object model definition in either XML or JSON format.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceObjectModelApi.CreateDeviceObjectModel(context.Background()).DeviceObjectModel(deviceObjectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceObjectModelApi.CreateDeviceObjectModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceObjectModel`: DeviceObjectModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceObjectModelApi.CreateDeviceObjectModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceObjectModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceObjectModel** | [**DeviceObjectModel**](DeviceObjectModel.md) | Device object model definition in either XML or JSON format. | 

### Return type

[**DeviceObjectModel**](DeviceObjectModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceObjectModel

> DeleteDeviceObjectModel(ctx, modelId).Execute()

対象のデバイスオブジェクトモデルを削除します。



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
    modelId := "modelId_example" // string | 対象のデバイスオブジェクトモデルの ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceObjectModelApi.DeleteDeviceObjectModel(context.Background(), modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceObjectModelApi.DeleteDeviceObjectModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | 対象のデバイスオブジェクトモデルの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceObjectModelRequest struct via the builder pattern


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


## GetDeviceObjectModel

> DeviceObjectModel GetDeviceObjectModel(ctx, modelId).Execute()

デバイスオブジェクトモデルを取得します。



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
    modelId := "modelId_example" // string | 対象のデバイスオブジェクトモデルの ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceObjectModelApi.GetDeviceObjectModel(context.Background(), modelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceObjectModelApi.GetDeviceObjectModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceObjectModel`: DeviceObjectModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceObjectModelApi.GetDeviceObjectModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | 対象のデバイスオブジェクトモデルの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceObjectModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeviceObjectModel**](DeviceObjectModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceObjectModels

> []DeviceObjectModel ListDeviceObjectModels(ctx).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

デバイスオブジェクトモデルのリストを返します。



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
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のデバイスオブジェクトモデルの ID。このパラメータを指定することで次のデバイスオブジェクトモデル以降のリストを取得できる。 (optional)
    limit := int32(56) // int32 | 取得する結果の上限数 (optional) (default to -1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceObjectModelApi.ListDeviceObjectModels(context.Background()).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceObjectModelApi.ListDeviceObjectModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceObjectModels`: []DeviceObjectModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceObjectModelApi.ListDeviceObjectModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceObjectModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lastEvaluatedKey** | **string** | 現ページで取得した最後のデバイスオブジェクトモデルの ID。このパラメータを指定することで次のデバイスオブジェクトモデル以降のリストを取得できる。 | 
 **limit** | **int32** | 取得する結果の上限数 | [default to -1]

### Return type

[**[]DeviceObjectModel**](DeviceObjectModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceObjectModelScope

> DeviceObjectModel SetDeviceObjectModelScope(ctx, modelId).SetDeviceObjectModelScopeRequest(setDeviceObjectModelScopeRequest).Execute()

対象のデバイスオブジェクトモデルが適用される Scope をセットします。



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
    modelId := "modelId_example" // string | 対象のデバイスオブジェクトモデルの ID
    setDeviceObjectModelScopeRequest := *openapiclient.NewSetDeviceObjectModelScopeRequest() // SetDeviceObjectModelScopeRequest | 対象のデバイスオブジェクトモデルにセットする Scope の値

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceObjectModelApi.SetDeviceObjectModelScope(context.Background(), modelId).SetDeviceObjectModelScopeRequest(setDeviceObjectModelScopeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceObjectModelApi.SetDeviceObjectModelScope``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDeviceObjectModelScope`: DeviceObjectModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceObjectModelApi.SetDeviceObjectModelScope`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | 対象のデバイスオブジェクトモデルの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceObjectModelScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setDeviceObjectModelScopeRequest** | [**SetDeviceObjectModelScopeRequest**](SetDeviceObjectModelScopeRequest.md) | 対象のデバイスオブジェクトモデルにセットする Scope の値 | 

### Return type

[**DeviceObjectModel**](DeviceObjectModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceObjectModel

> DeviceObjectModel UpdateDeviceObjectModel(ctx, modelId).DeviceObjectModel(deviceObjectModel).Execute()

デバイスオブジェクトモデルを更新します。



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
    modelId := "modelId_example" // string | 対象のデバイスオブジェクトモデルの ID
    deviceObjectModel := *openapiclient.NewDeviceObjectModel() // DeviceObjectModel | XML か JSON により表現されたデバイスオブジェクトモデルの表記。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeviceObjectModelApi.UpdateDeviceObjectModel(context.Background(), modelId).DeviceObjectModel(deviceObjectModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceObjectModelApi.UpdateDeviceObjectModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceObjectModel`: DeviceObjectModel
    fmt.Fprintf(os.Stdout, "Response from `DeviceObjectModelApi.UpdateDeviceObjectModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelId** | **string** | 対象のデバイスオブジェクトモデルの ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceObjectModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceObjectModel** | [**DeviceObjectModel**](DeviceObjectModel.md) | XML か JSON により表現されたデバイスオブジェクトモデルの表記。 | 

### Return type

[**DeviceObjectModel**](DeviceObjectModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

