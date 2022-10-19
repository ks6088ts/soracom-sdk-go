# \PaymentApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePaymentMethod**](PaymentApi.md#ActivatePaymentMethod) | **Post** /payment_methods/current/activate | Activate payment method.
[**ExportPaymentStatement**](PaymentApi.md#ExportPaymentStatement) | **Post** /payment_statements/{payment_statement_id}/export | Export payment statement.
[**GetPayerInformation**](PaymentApi.md#GetPayerInformation) | **Get** /payment_statements/payer_information | Export payer information.
[**GetPaymentMethod**](PaymentApi.md#GetPaymentMethod) | **Get** /payment_methods/current | Get payment method information.
[**GetPaymentTransaction**](PaymentApi.md#GetPaymentTransaction) | **Get** /payment_history/transactions/{payment_transaction_id} | Get payment transaction result.
[**GetVolumeDiscount**](PaymentApi.md#GetVolumeDiscount) | **Get** /volume_discounts/{contract_id} | Get long term discount.
[**ListCoupons**](PaymentApi.md#ListCoupons) | **Get** /coupons | List coupons.
[**ListPaymentStatements**](PaymentApi.md#ListPaymentStatements) | **Get** /payment_statements | List payment statements.
[**ListVolumeDiscounts**](PaymentApi.md#ListVolumeDiscounts) | **Get** /volume_discounts | List long term discounts.
[**RegisterCoupon**](PaymentApi.md#RegisterCoupon) | **Post** /coupons/{coupon_code}/register | Register Coupon.
[**RegisterPayerInformation**](PaymentApi.md#RegisterPayerInformation) | **Post** /payment_statements/payer_information | Register payer information.



## ActivatePaymentMethod

> map[string]interface{} ActivatePaymentMethod(ctx).Execute()

Activate payment method.



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
    resp, r, err := apiClient.PaymentApi.ActivatePaymentMethod(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ActivatePaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivatePaymentMethod`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ActivatePaymentMethod`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePaymentMethodRequest struct via the builder pattern


### Return type

**map[string]interface{}**

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportPaymentStatement

> FileExportResponse ExportPaymentStatement(ctx, paymentStatementId).ExportMode(exportMode).Execute()

Export payment statement.



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
    paymentStatementId := "paymentStatementId_example" // string | 課金明細 ID
    exportMode := "exportMode_example" // string | 出力モード (非同期, 同期) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.ExportPaymentStatement(context.Background(), paymentStatementId).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ExportPaymentStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportPaymentStatement`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ExportPaymentStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentStatementId** | **string** | 課金明細 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportPaymentStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportMode** | **string** | 出力モード (非同期, 同期) | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayerInformation

> RegisterPayerInformationModel GetPayerInformation(ctx).Execute()

Export payer information.



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
    resp, r, err := apiClient.PaymentApi.GetPayerInformation(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetPayerInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayerInformation`: RegisterPayerInformationModel
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetPayerInformation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayerInformationRequest struct via the builder pattern


### Return type

[**RegisterPayerInformationModel**](RegisterPayerInformationModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentMethod

> GetPaymentMethodResult GetPaymentMethod(ctx).Execute()

Get payment method information.



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
    resp, r, err := apiClient.PaymentApi.GetPaymentMethod(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetPaymentMethod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentMethod`: GetPaymentMethodResult
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetPaymentMethod`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentMethodRequest struct via the builder pattern


### Return type

[**GetPaymentMethodResult**](GetPaymentMethodResult.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPaymentTransaction

> GetPaymentTransactionResult GetPaymentTransaction(ctx, paymentTransactionId).Execute()

Get payment transaction result.



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
    paymentTransactionId := "paymentTransactionId_example" // string | 課金処理 ID。[Billing:getBillingHistory API](#/Billing/getBillingHistory) や [Billing:getBilling API](#/Billing/getBilling) で取得できます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.GetPaymentTransaction(context.Background(), paymentTransactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetPaymentTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPaymentTransaction`: GetPaymentTransactionResult
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetPaymentTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentTransactionId** | **string** | 課金処理 ID。[Billing:getBillingHistory API](#/Billing/getBillingHistory) や [Billing:getBilling API](#/Billing/getBilling) で取得できます。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetPaymentTransactionResult**](GetPaymentTransactionResult.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeDiscount

> GetVolumeDiscountResponse GetVolumeDiscount(ctx, contractId).Execute()

Get long term discount.



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
    contractId := "contractId_example" // string | contract_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.GetVolumeDiscount(context.Background(), contractId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.GetVolumeDiscount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumeDiscount`: GetVolumeDiscountResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.GetVolumeDiscount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contractId** | **string** | contract_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeDiscountRequest struct via the builder pattern


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


## ListCoupons

> ListCouponResponse ListCoupons(ctx).Execute()

List coupons.



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
    resp, r, err := apiClient.PaymentApi.ListCoupons(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ListCoupons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCoupons`: ListCouponResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ListCoupons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCouponsRequest struct via the builder pattern


### Return type

[**ListCouponResponse**](ListCouponResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPaymentStatements

> ListPaymentStatementResponse ListPaymentStatements(ctx).Execute()

List payment statements.



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
    resp, r, err := apiClient.PaymentApi.ListPaymentStatements(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ListPaymentStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPaymentStatements`: ListPaymentStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ListPaymentStatements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPaymentStatementsRequest struct via the builder pattern


### Return type

[**ListPaymentStatementResponse**](ListPaymentStatementResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeDiscounts

> ListVolumeDiscountResponse ListVolumeDiscounts(ctx).Execute()

List long term discounts.



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
    resp, r, err := apiClient.PaymentApi.ListVolumeDiscounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.ListVolumeDiscounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVolumeDiscounts`: ListVolumeDiscountResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.ListVolumeDiscounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeDiscountsRequest struct via the builder pattern


### Return type

[**ListVolumeDiscountResponse**](ListVolumeDiscountResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCoupon

> CouponResponse RegisterCoupon(ctx, couponCode).Execute()

Register Coupon.



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
    couponCode := "couponCode_example" // string | クーポンコード

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.RegisterCoupon(context.Background(), couponCode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.RegisterCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterCoupon`: CouponResponse
    fmt.Fprintf(os.Stdout, "Response from `PaymentApi.RegisterCoupon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCode** | **string** | クーポンコード | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCouponRequest struct via the builder pattern


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


## RegisterPayerInformation

> RegisterPayerInformation(ctx).RegisterPayerInformationModel(registerPayerInformationModel).Execute()

Register payer information.



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
    registerPayerInformationModel := *openapiclient.NewRegisterPayerInformationModel() // RegisterPayerInformationModel | 課金明細に登録する支払い者情報

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PaymentApi.RegisterPayerInformation(context.Background()).RegisterPayerInformationModel(registerPayerInformationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentApi.RegisterPayerInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPayerInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerPayerInformationModel** | [**RegisterPayerInformationModel**](RegisterPayerInformationModel.md) | 課金明細に登録する支払い者情報 | 

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

