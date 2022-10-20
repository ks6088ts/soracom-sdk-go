# \OperatorApi

All URIs are relative to *https://api-sandbox.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SandboxDeleteOperator**](OperatorApi.md#SandboxDeleteOperator) | **Delete** /sandbox/operators/{operator_id} | Operator を削除する
[**SandboxGetSignupToken**](OperatorApi.md#SandboxGetSignupToken) | **Post** /sandbox/operators/token/{email} | サインアップトークンを取得する
[**SandboxInitializeOperator**](OperatorApi.md#SandboxInitializeOperator) | **Post** /sandbox/init | Operator を初期化する



## SandboxDeleteOperator

> SandboxDeleteOperator(ctx, operatorId).Execute()

Operator を削除する



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
    resp, r, err := apiClient.OperatorApi.SandboxDeleteOperator(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.SandboxDeleteOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxDeleteOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxGetSignupToken

> SandboxGetSignupTokenResponse SandboxGetSignupToken(ctx, email).SandboxGetSignupTokenRequest(sandboxGetSignupTokenRequest).Execute()

サインアップトークンを取得する



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
    email := "email_example" // string | email
    sandboxGetSignupTokenRequest := *openapiclient.NewSandboxGetSignupTokenRequest() // SandboxGetSignupTokenRequest | 認証リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.SandboxGetSignupToken(context.Background(), email).SandboxGetSignupTokenRequest(sandboxGetSignupTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.SandboxGetSignupToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxGetSignupToken`: SandboxGetSignupTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.SandboxGetSignupToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**email** | **string** | email | 

### Other Parameters

Other parameters are passed through a pointer to a apiSandboxGetSignupTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sandboxGetSignupTokenRequest** | [**SandboxGetSignupTokenRequest**](SandboxGetSignupTokenRequest.md) | 認証リクエスト | 

### Return type

[**SandboxGetSignupTokenResponse**](SandboxGetSignupTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxInitializeOperator

> SandboxAuthResponse SandboxInitializeOperator(ctx).SandboxInitRequest(sandboxInitRequest).Execute()

Operator を初期化する



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
    sandboxInitRequest := *openapiclient.NewSandboxInitRequest("AuthKey_example", "AuthKeyId_example", "Email_example", "Password_example") // SandboxInitRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.SandboxInitializeOperator(context.Background()).SandboxInitRequest(sandboxInitRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.SandboxInitializeOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxInitializeOperator`: SandboxAuthResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.SandboxInitializeOperator`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxInitializeOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxInitRequest** | [**SandboxInitRequest**](SandboxInitRequest.md) | request | 

### Return type

[**SandboxAuthResponse**](SandboxAuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

