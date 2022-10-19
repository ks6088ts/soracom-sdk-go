# \EventHandlerApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventHandler**](EventHandlerApi.md#CreateEventHandler) | **Post** /event_handlers | Create Event Handler
[**DeleteEventHandler**](EventHandlerApi.md#DeleteEventHandler) | **Delete** /event_handlers/{handler_id} | Delete Event Handler
[**DeleteIgnoreEventHandler**](EventHandlerApi.md#DeleteIgnoreEventHandler) | **Delete** /event_handlers/{handler_id}/subscribers/{imsi}/ignore | Delete Ignore Event Handler
[**GetEventHandler**](EventHandlerApi.md#GetEventHandler) | **Get** /event_handlers/{handler_id} | Get Event Handler
[**ListEventHandlers**](EventHandlerApi.md#ListEventHandlers) | **Get** /event_handlers | List Event Handlers
[**ListEventHandlersBySubscriber**](EventHandlerApi.md#ListEventHandlersBySubscriber) | **Get** /event_handlers/subscribers/{imsi} | List Event Handlers related to Subscriber
[**SetIgnoreEventHandler**](EventHandlerApi.md#SetIgnoreEventHandler) | **Post** /event_handlers/{handler_id}/subscribers/{imsi}/ignore | Ignore Event Handler
[**UpdateEventHandler**](EventHandlerApi.md#UpdateEventHandler) | **Put** /event_handlers/{handler_id} | Update Event Handler



## CreateEventHandler

> EventHandlerModel CreateEventHandler(ctx).CreateEventHandlerRequest(createEventHandlerRequest).Execute()

Create Event Handler



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
    createEventHandlerRequest := *openapiclient.NewCreateEventHandlerRequest([]openapiclient.ActionConfig{*openapiclient.NewActionConfig(*openapiclient.NewActionConfigProperty("ExecutionDateTimeConst_example"), "Type_example")}, *openapiclient.NewRuleConfig(*openapiclient.NewRuleConfigProperty("InactiveTimeoutDateConst_example"), "Type_example"), "Status_example") // CreateEventHandlerRequest | イベントハンドラの設定内容

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.CreateEventHandler(context.Background()).CreateEventHandlerRequest(createEventHandlerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.CreateEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEventHandler`: EventHandlerModel
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerApi.CreateEventHandler`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEventHandlerRequest** | [**CreateEventHandlerRequest**](CreateEventHandlerRequest.md) | イベントハンドラの設定内容 | 

### Return type

[**EventHandlerModel**](EventHandlerModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEventHandler

> DeleteEventHandler(ctx, handlerId).Execute()

Delete Event Handler



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
    handlerId := "handlerId_example" // string | handler ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.DeleteEventHandler(context.Background(), handlerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.DeleteEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handlerId** | **string** | handler ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventHandlerRequest struct via the builder pattern


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


## DeleteIgnoreEventHandler

> DeleteIgnoreEventHandler(ctx, imsi, handlerId).Execute()

Delete Ignore Event Handler



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
    imsi := "imsi_example" // string | imsi
    handlerId := "handlerId_example" // string | handler_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.DeleteIgnoreEventHandler(context.Background(), imsi, handlerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.DeleteIgnoreEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 
**handlerId** | **string** | handler_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIgnoreEventHandlerRequest struct via the builder pattern


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


## GetEventHandler

> EventHandlerModel GetEventHandler(ctx, handlerId).Execute()

Get Event Handler



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
    handlerId := "handlerId_example" // string | handler ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.GetEventHandler(context.Background(), handlerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.GetEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventHandler`: EventHandlerModel
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerApi.GetEventHandler`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handlerId** | **string** | handler ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EventHandlerModel**](EventHandlerModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventHandlers

> []EventHandlerModel ListEventHandlers(ctx).Target(target).Execute()

List Event Handlers



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
    target := "target_example" // string | target (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.ListEventHandlers(context.Background()).Target(target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.ListEventHandlers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventHandlers`: []EventHandlerModel
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerApi.ListEventHandlers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventHandlersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **target** | **string** | target | 

### Return type

[**[]EventHandlerModel**](EventHandlerModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEventHandlersBySubscriber

> []EventHandlerModel ListEventHandlersBySubscriber(ctx, imsi).Execute()

List Event Handlers related to Subscriber



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
    imsi := "imsi_example" // string | imsi

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.ListEventHandlersBySubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.ListEventHandlersBySubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventHandlersBySubscriber`: []EventHandlerModel
    fmt.Fprintf(os.Stdout, "Response from `EventHandlerApi.ListEventHandlersBySubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventHandlersBySubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EventHandlerModel**](EventHandlerModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIgnoreEventHandler

> SetIgnoreEventHandler(ctx, imsi, handlerId).Execute()

Ignore Event Handler



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
    imsi := "imsi_example" // string | imsi
    handlerId := "handlerId_example" // string | handler_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.SetIgnoreEventHandler(context.Background(), imsi, handlerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.SetIgnoreEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 
**handlerId** | **string** | handler_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIgnoreEventHandlerRequest struct via the builder pattern


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


## UpdateEventHandler

> UpdateEventHandler(ctx, handlerId).Body(body).Execute()

Update Event Handler



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
    handlerId := "handlerId_example" // string | handler ID
    body := CreateEventHandlerRequest(987) // CreateEventHandlerRequest | イベントハンドラの設定内容

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventHandlerApi.UpdateEventHandler(context.Background(), handlerId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventHandlerApi.UpdateEventHandler``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handlerId** | **string** | handler ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventHandlerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **CreateEventHandlerRequest** | イベントハンドラの設定内容 | 

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

