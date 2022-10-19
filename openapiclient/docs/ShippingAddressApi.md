# \ShippingAddressApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateShippingAddress**](ShippingAddressApi.md#CreateShippingAddress) | **Post** /operators/{operator_id}/shipping_addresses | Create shipping address
[**DeleteShippingAddress**](ShippingAddressApi.md#DeleteShippingAddress) | **Delete** /operators/{operator_id}/shipping_addresses/{shipping_address_id} | Delete shipping address
[**GetShippingAddress**](ShippingAddressApi.md#GetShippingAddress) | **Get** /operators/{operator_id}/shipping_addresses/{shipping_address_id} | Get shipping address
[**ListShippingAddresses**](ShippingAddressApi.md#ListShippingAddresses) | **Get** /operators/{operator_id}/shipping_addresses | List shipping addresses
[**UpdateShippingAddress**](ShippingAddressApi.md#UpdateShippingAddress) | **Put** /operators/{operator_id}/shipping_addresses/{shipping_address_id} | Update shipping address



## CreateShippingAddress

> GetShippingAddressResponse CreateShippingAddress(ctx, operatorId).ShippingAddressModel(shippingAddressModel).Execute()

Create shipping address



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
    shippingAddressModel := *openapiclient.NewShippingAddressModel("AddressLine1_example", "City_example", "State_example", "ZipCode_example") // ShippingAddressModel | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingAddressApi.CreateShippingAddress(context.Background(), operatorId).ShippingAddressModel(shippingAddressModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingAddressApi.CreateShippingAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShippingAddress`: GetShippingAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingAddressApi.CreateShippingAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateShippingAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **shippingAddressModel** | [**ShippingAddressModel**](ShippingAddressModel.md) | model | 

### Return type

[**GetShippingAddressResponse**](GetShippingAddressResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteShippingAddress

> DeleteShippingAddress(ctx, operatorId, shippingAddressId).Execute()

Delete shipping address



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
    shippingAddressId := "shippingAddressId_example" // string | shipping_address_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingAddressApi.DeleteShippingAddress(context.Background(), operatorId, shippingAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingAddressApi.DeleteShippingAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**shippingAddressId** | **string** | shipping_address_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteShippingAddressRequest struct via the builder pattern


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


## GetShippingAddress

> GetShippingAddressResponse GetShippingAddress(ctx, operatorId, shippingAddressId).Execute()

Get shipping address



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
    shippingAddressId := "shippingAddressId_example" // string | shipping_address_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingAddressApi.GetShippingAddress(context.Background(), operatorId, shippingAddressId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingAddressApi.GetShippingAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShippingAddress`: GetShippingAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingAddressApi.GetShippingAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**shippingAddressId** | **string** | shipping_address_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShippingAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetShippingAddressResponse**](GetShippingAddressResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListShippingAddresses

> ListShippingAddressResponse ListShippingAddresses(ctx, operatorId).Execute()

List shipping addresses



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
    resp, r, err := apiClient.ShippingAddressApi.ListShippingAddresses(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingAddressApi.ListShippingAddresses``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListShippingAddresses`: ListShippingAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `ShippingAddressApi.ListShippingAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListShippingAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListShippingAddressResponse**](ListShippingAddressResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShippingAddress

> UpdateShippingAddress(ctx, operatorId, shippingAddressId).ShippingAddressModel(shippingAddressModel).Execute()

Update shipping address



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
    operatorId := "operatorId_example" // string | Operator ID
    shippingAddressId := "shippingAddressId_example" // string | shipping_address_id
    shippingAddressModel := *openapiclient.NewShippingAddressModel("AddressLine1_example", "City_example", "State_example", "ZipCode_example") // ShippingAddressModel | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ShippingAddressApi.UpdateShippingAddress(context.Background(), operatorId, shippingAddressId).ShippingAddressModel(shippingAddressModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ShippingAddressApi.UpdateShippingAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | Operator ID | 
**shippingAddressId** | **string** | shipping_address_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShippingAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shippingAddressModel** | [**ShippingAddressModel**](ShippingAddressModel.md) | model | 

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

