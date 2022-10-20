# \OrderApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrderApi.md#CancelOrder) | **Put** /orders/{order_id}/cancel | Cancel order.
[**ConfirmCouponOrder**](OrderApi.md#ConfirmCouponOrder) | **Put** /coupons/{order_id}/confirm | Confirm coupon order.
[**ConfirmOrder**](OrderApi.md#ConfirmOrder) | **Put** /orders/{order_id}/confirm | Confirm order.
[**ConfirmVolumeDiscountOrder**](OrderApi.md#ConfirmVolumeDiscountOrder) | **Put** /volume_discounts/{order_id}/confirm | Confirm long term discount order.
[**CreateCouponQuotation**](OrderApi.md#CreateCouponQuotation) | **Post** /coupons | Create coupon quotation.
[**CreateQuotation**](OrderApi.md#CreateQuotation) | **Post** /orders | Create Quotation
[**CreateVolumeDiscountQuotation**](OrderApi.md#CreateVolumeDiscountQuotation) | **Post** /volume_discounts | Create long term discount quotation.
[**GetOrder**](OrderApi.md#GetOrder) | **Get** /orders/{order_id} | 発注確定済みの発注情報を 1 件だけ取得する
[**ListAvailableDiscounts**](OrderApi.md#ListAvailableDiscounts) | **Get** /volume_discounts/available_discounts | List available long term discounts.
[**ListOrderedSubscribers**](OrderApi.md#ListOrderedSubscribers) | **Get** /orders/{order_id}/subscribers | List ordered subscribers.
[**ListOrders**](OrderApi.md#ListOrders) | **Get** /orders | 発注確定済みの注文履歴の一覧を取得する
[**ListProducts**](OrderApi.md#ListProducts) | **Get** /products | List products
[**RegisterOrderedSim**](OrderApi.md#RegisterOrderedSim) | **Post** /orders/{order_id}/subscribers/register | Register subscribers for operator.



## CancelOrder

> string CancelOrder(ctx, orderId).Execute()

Cancel order.



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
    orderId := "orderId_example" // string | 発注 ID。[`Order:listOrders API`](#/Order/listOrders) を呼び出すと取得できます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CancelOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrder`: string
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID。[&#x60;Order:listOrders API&#x60;](#/Order/listOrders) を呼び出すと取得できます。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmCouponOrder

> CouponResponse ConfirmCouponOrder(ctx, orderId).Execute()

Confirm coupon order.



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
    orderId := "orderId_example" // string | 発注 ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.ConfirmCouponOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ConfirmCouponOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmCouponOrder`: CouponResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.ConfirmCouponOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmCouponOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CouponResponse**](CouponResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConfirmOrder

> ConfirmOrder(ctx, orderId).Execute()

Confirm order.



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
    orderId := "orderId_example" // string | 発注 ID。[`Order:listOrders API`](#/Order/listOrders) を呼び出すと取得できます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.ConfirmOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ConfirmOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID。[&#x60;Order:listOrders API&#x60;](#/Order/listOrders) を呼び出すと取得できます。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmOrderRequest struct via the builder pattern


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


## ConfirmVolumeDiscountOrder

> GetVolumeDiscountResponse ConfirmVolumeDiscountOrder(ctx, orderId).Execute()

Confirm long term discount order.



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
    orderId := "orderId_example" // string | 発注 ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.ConfirmVolumeDiscountOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ConfirmVolumeDiscountOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConfirmVolumeDiscountOrder`: GetVolumeDiscountResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.ConfirmVolumeDiscountOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiConfirmVolumeDiscountOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVolumeDiscountResponse**](GetVolumeDiscountResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCouponQuotation

> EstimatedCouponModel CreateCouponQuotation(ctx).CreateEstimatedCouponRequest(createEstimatedCouponRequest).Execute()

Create coupon quotation.



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
    createEstimatedCouponRequest := *openapiclient.NewCreateEstimatedCouponRequest(float64(123)) // CreateEstimatedCouponRequest | クーポン詳細

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CreateCouponQuotation(context.Background()).CreateEstimatedCouponRequest(createEstimatedCouponRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CreateCouponQuotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCouponQuotation`: EstimatedCouponModel
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CreateCouponQuotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCouponQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEstimatedCouponRequest** | [**CreateEstimatedCouponRequest**](CreateEstimatedCouponRequest.md) | クーポン詳細 | 

### Return type

[**EstimatedCouponModel**](EstimatedCouponModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateQuotation

> EstimatedOrderModel CreateQuotation(ctx).CreateEstimatedOrderRequest(createEstimatedOrderRequest).Execute()

Create Quotation



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
    createEstimatedOrderRequest := *openapiclient.NewCreateEstimatedOrderRequest() // CreateEstimatedOrderRequest | 注文品リストと配送先 ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CreateQuotation(context.Background()).CreateEstimatedOrderRequest(createEstimatedOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CreateQuotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateQuotation`: EstimatedOrderModel
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CreateQuotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEstimatedOrderRequest** | [**CreateEstimatedOrderRequest**](CreateEstimatedOrderRequest.md) | 注文品リストと配送先 ID | 

### Return type

[**EstimatedOrderModel**](EstimatedOrderModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeDiscountQuotation

> EstimatedVolumeDiscountModel CreateVolumeDiscountQuotation(ctx).CreateEstimatedVolumeDiscountRequest(createEstimatedVolumeDiscountRequest).Execute()

Create long term discount quotation.



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
    createEstimatedVolumeDiscountRequest := *openapiclient.NewCreateEstimatedVolumeDiscountRequest(int32(123), int32(123), "VolumeDiscountPaymentType_example", "VolumeDiscountType_example") // CreateEstimatedVolumeDiscountRequest | 長期割引契約詳細

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.CreateVolumeDiscountQuotation(context.Background()).CreateEstimatedVolumeDiscountRequest(createEstimatedVolumeDiscountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.CreateVolumeDiscountQuotation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumeDiscountQuotation`: EstimatedVolumeDiscountModel
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.CreateVolumeDiscountQuotation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeDiscountQuotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEstimatedVolumeDiscountRequest** | [**CreateEstimatedVolumeDiscountRequest**](CreateEstimatedVolumeDiscountRequest.md) | 長期割引契約詳細 | 

### Return type

[**EstimatedVolumeDiscountModel**](EstimatedVolumeDiscountModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> GetOrderResponse GetOrder(ctx, orderId).Execute()

発注確定済みの発注情報を 1 件だけ取得する



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
    orderId := "orderId_example" // string | 発注 ID。[`Order:listOrders API`](#/Order/listOrders) を呼び出すと取得できます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.GetOrder(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: GetOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID。[&#x60;Order:listOrders API&#x60;](#/Order/listOrders) を呼び出すと取得できます。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOrderResponse**](GetOrderResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailableDiscounts

> AvailableLongTermDiscountResponse ListAvailableDiscounts(ctx).Execute()

List available long term discounts.



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
    resp, r, err := apiClient.OrderApi.ListAvailableDiscounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ListAvailableDiscounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAvailableDiscounts`: AvailableLongTermDiscountResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.ListAvailableDiscounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAvailableDiscountsRequest struct via the builder pattern


### Return type

[**AvailableLongTermDiscountResponse**](AvailableLongTermDiscountResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrderedSubscribers

> ListOrderedSubscriberResponse ListOrderedSubscribers(ctx, orderId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

List ordered subscribers.



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
    orderId := "orderId_example" // string | 発注 ID。[`Order:listOrders API`](#/Order/listOrders) を呼び出すと取得できます。
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 直前のリクエストで取得したうち、最後の Subscriber の製造番号。レスポンスヘッダの X-Soracom-Next-Key に含まれる値 (optional)
    limit := int32(56) // int32 | レスポンスに含まれる最大 Subscriber 数 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.ListOrderedSubscribers(context.Background(), orderId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ListOrderedSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrderedSubscribers`: ListOrderedSubscriberResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.ListOrderedSubscribers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID。[&#x60;Order:listOrders API&#x60;](#/Order/listOrders) を呼び出すと取得できます。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderedSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastEvaluatedKey** | **string** | 直前のリクエストで取得したうち、最後の Subscriber の製造番号。レスポンスヘッダの X-Soracom-Next-Key に含まれる値 | 
 **limit** | **int32** | レスポンスに含まれる最大 Subscriber 数 | 

### Return type

[**ListOrderedSubscriberResponse**](ListOrderedSubscriberResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrders

> ListOrderResponse ListOrders(ctx).Execute()

発注確定済みの注文履歴の一覧を取得する



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
    resp, r, err := apiClient.OrderApi.ListOrders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ListOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrders`: ListOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.ListOrders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


### Return type

[**ListOrderResponse**](ListOrderResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListProducts

> ListProductResponse ListProducts(ctx).Execute()

List products



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
    resp, r, err := apiClient.OrderApi.ListProducts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.ListProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProducts`: ListProductResponse
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.ListProducts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProductsRequest struct via the builder pattern


### Return type

[**ListProductResponse**](ListProductResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterOrderedSim

> string RegisterOrderedSim(ctx, orderId).Execute()

Register subscribers for operator.



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
    orderId := "orderId_example" // string | 発注 ID。[`Order:listOrders API`](#/Order/listOrders) を呼び出すと取得できます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrderApi.RegisterOrderedSim(context.Background(), orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrderApi.RegisterOrderedSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterOrderedSim`: string
    fmt.Fprintf(os.Stdout, "Response from `OrderApi.RegisterOrderedSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **string** | 発注 ID。[&#x60;Order:listOrders API&#x60;](#/Order/listOrders) を呼び出すと取得できます。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterOrderedSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

