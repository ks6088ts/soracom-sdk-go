# \SystemNotificationApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSystemNotification**](SystemNotificationApi.md#DeleteSystemNotification) | **Delete** /operators/{operator_id}/system_notifications/{type} | Delete system notification
[**GetSystemNotification**](SystemNotificationApi.md#GetSystemNotification) | **Get** /operators/{operator_id}/system_notifications/{type} | Get system notification
[**ListSystemNotifications**](SystemNotificationApi.md#ListSystemNotifications) | **Get** /operators/{operator_id}/system_notifications | List system notifications
[**SetSystemNotification**](SystemNotificationApi.md#SetSystemNotification) | **Post** /operators/{operator_id}/system_notifications/{type} | Set system notification



## DeleteSystemNotification

> DeleteSystemNotification(ctx, operatorId, type_).Execute()

Delete system notification



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
    operatorId := "operatorId_example" // string | operator_id
    type_ := "type__example" // string | system notification type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemNotificationApi.DeleteSystemNotification(context.Background(), operatorId, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemNotificationApi.DeleteSystemNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**type_** | **string** | system notification type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSystemNotificationRequest struct via the builder pattern


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


## GetSystemNotification

> SystemNotificationsModel GetSystemNotification(ctx, operatorId, type_).Execute()

Get system notification



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
    operatorId := "operatorId_example" // string | operator_id
    type_ := "type__example" // string | system notification type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemNotificationApi.GetSystemNotification(context.Background(), operatorId, type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemNotificationApi.GetSystemNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSystemNotification`: SystemNotificationsModel
    fmt.Fprintf(os.Stdout, "Response from `SystemNotificationApi.GetSystemNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**type_** | **string** | system notification type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSystemNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SystemNotificationsModel**](SystemNotificationsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSystemNotifications

> []SystemNotificationsModel ListSystemNotifications(ctx, operatorId).Execute()

List system notifications



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
    operatorId := "operatorId_example" // string | operator_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemNotificationApi.ListSystemNotifications(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemNotificationApi.ListSystemNotifications``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSystemNotifications`: []SystemNotificationsModel
    fmt.Fprintf(os.Stdout, "Response from `SystemNotificationApi.ListSystemNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSystemNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SystemNotificationsModel**](SystemNotificationsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSystemNotification

> SystemNotificationsModel SetSystemNotification(ctx, operatorId, type_).SetSystemNotificationsRequest(setSystemNotificationsRequest).Execute()

Set system notification



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
    operatorId := "operatorId_example" // string | operator_id
    type_ := "type__example" // string | system notification type
    setSystemNotificationsRequest := *openapiclient.NewSetSystemNotificationsRequest([]string{"EmailIdList_example"}) // SetSystemNotificationsRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SystemNotificationApi.SetSystemNotification(context.Background(), operatorId, type_).SetSystemNotificationsRequest(setSystemNotificationsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SystemNotificationApi.SetSystemNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSystemNotification`: SystemNotificationsModel
    fmt.Fprintf(os.Stdout, "Response from `SystemNotificationApi.SetSystemNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**type_** | **string** | system notification type | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSystemNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setSystemNotificationsRequest** | [**SetSystemNotificationsRequest**](SetSystemNotificationsRequest.md) | request | 

### Return type

[**SystemNotificationsModel**](SystemNotificationsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

