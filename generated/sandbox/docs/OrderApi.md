# \OrderApi

All URIs are relative to *https://api-sandbox.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxShipOrder**](OrderApi.md#SandboxShipOrder) | **Post** /sandbox/orders/ship | Order を出荷済状態にする



## SandboxShipOrder

> SandboxShipOrder(ctx).SandboxShipOrderRequest(sandboxShipOrderRequest).Execute()

Order を出荷済状態にする



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
    sandboxShipOrderRequest := *openapiclient.NewSandboxShipOrderRequest("OperatorId_example", "OrderId_example") // SandboxShipOrderRequest | Shipping request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.SandboxShipOrder(context.Background()).SandboxShipOrderRequest(sandboxShipOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.SandboxShipOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxShipOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxShipOrderRequest** | [**SandboxShipOrderRequest**](SandboxShipOrderRequest.md) | Shipping request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

