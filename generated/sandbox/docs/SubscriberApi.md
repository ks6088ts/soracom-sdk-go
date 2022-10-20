# \SubscriberApi

All URIs are relative to *https://api-sandbox.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxCreateSubscriber**](SubscriberApi.md#SandboxCreateSubscriber) | **Post** /sandbox/subscribers/create | Subscriber を作成する



## SandboxCreateSubscriber

> SandboxCreateSubscriberResponse SandboxCreateSubscriber(ctx).SandboxCreateSubscriberRequest(sandboxCreateSubscriberRequest).Execute()

Subscriber を作成する



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
    sandboxCreateSubscriberRequest := *openapiclient.NewSandboxCreateSubscriberRequest() // SandboxCreateSubscriberRequest | Create request (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SandboxCreateSubscriber(context.Background()).SandboxCreateSubscriberRequest(sandboxCreateSubscriberRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SandboxCreateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxCreateSubscriber`: SandboxCreateSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SandboxCreateSubscriber`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxCreateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxCreateSubscriberRequest** | [**SandboxCreateSubscriberRequest**](SandboxCreateSubscriberRequest.md) | Create request | 

### Return type

[**SandboxCreateSubscriberResponse**](SandboxCreateSubscriberResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

