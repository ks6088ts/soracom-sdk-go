# \AuthApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Auth**](AuthApi.md#Auth) | **Post** /auth | API アクセスの認証を行い、API キーと API トークンを発行します。
[**IssuePasswordResetToken**](AuthApi.md#IssuePasswordResetToken) | **Post** /auth/password_reset_token/issue | パスワードリセット用のトークンを発行します。
[**Logout**](AuthApi.md#Logout) | **Post** /auth/logout | SORACOM API にアクセスするための API キーと API トークンを無効化します。
[**VerifyPasswordResetToken**](AuthApi.md#VerifyPasswordResetToken) | **Post** /auth/password_reset_token/verify | パスワードリセット用のトークンを検証し、パスワードを更新します。



## Auth

> AuthResponse Auth(ctx).AuthRequest(authRequest).Execute()

API アクセスの認証を行い、API キーと API トークンを発行します。



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
    authRequest := *openapiclient.NewAuthRequest() // AuthRequest | 認証リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.Auth(context.Background()).AuthRequest(authRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.Auth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Auth`: AuthResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.Auth`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authRequest** | [**AuthRequest**](AuthRequest.md) | 認証リクエスト | 

### Return type

[**AuthResponse**](AuthResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssuePasswordResetToken

> IssuePasswordResetToken(ctx).IssuePasswordResetTokenRequest(issuePasswordResetTokenRequest).Execute()

パスワードリセット用のトークンを発行します。



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
    issuePasswordResetTokenRequest := *openapiclient.NewIssuePasswordResetTokenRequest("Email_example") // IssuePasswordResetTokenRequest | email address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.IssuePasswordResetToken(context.Background()).IssuePasswordResetTokenRequest(issuePasswordResetTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.IssuePasswordResetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssuePasswordResetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issuePasswordResetTokenRequest** | [**IssuePasswordResetTokenRequest**](IssuePasswordResetTokenRequest.md) | email address | 

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


## Logout

> Logout(ctx).Execute()

SORACOM API にアクセスするための API キーと API トークンを無効化します。



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
    resp, r, err := apiClient.AuthApi.Logout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.Logout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiLogoutRequest struct via the builder pattern


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


## VerifyPasswordResetToken

> VerifyPasswordResetToken(ctx).VerifyPasswordResetTokenRequest(verifyPasswordResetTokenRequest).Execute()

パスワードリセット用のトークンを検証し、パスワードを更新します。



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
    verifyPasswordResetTokenRequest := *openapiclient.NewVerifyPasswordResetTokenRequest("Password_example", "Token_example") // VerifyPasswordResetTokenRequest | token, password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.VerifyPasswordResetToken(context.Background()).VerifyPasswordResetTokenRequest(verifyPasswordResetTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.VerifyPasswordResetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPasswordResetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyPasswordResetTokenRequest** | [**VerifyPasswordResetTokenRequest**](VerifyPasswordResetTokenRequest.md) | token, password | 

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

