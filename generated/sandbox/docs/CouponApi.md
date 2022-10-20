# \CouponApi

All URIs are relative to *https://api-sandbox.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxCreateCoupon**](CouponApi.md#SandboxCreateCoupon) | **Post** /sandbox/coupons/create | クーポンを作成する



## SandboxCreateCoupon

> SandboxCreateCouponResponse SandboxCreateCoupon(ctx).SandboxCreateCouponRequest(sandboxCreateCouponRequest).Execute()

クーポンを作成する



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
    sandboxCreateCouponRequest := *openapiclient.NewSandboxCreateCouponRequest() // SandboxCreateCouponRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CouponApi.SandboxCreateCoupon(context.Background()).SandboxCreateCouponRequest(sandboxCreateCouponRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CouponApi.SandboxCreateCoupon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxCreateCoupon`: SandboxCreateCouponResponse
    fmt.Fprintf(os.Stdout, "Response from `CouponApi.SandboxCreateCoupon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxCreateCouponRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxCreateCouponRequest** | [**SandboxCreateCouponRequest**](SandboxCreateCouponRequest.md) | request | 

### Return type

[**SandboxCreateCouponResponse**](SandboxCreateCouponResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

