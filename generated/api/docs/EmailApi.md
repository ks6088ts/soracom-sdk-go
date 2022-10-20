# \EmailApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteEmail**](EmailApi.md#DeleteEmail) | **Delete** /operators/{operator_id}/emails/{email_id} | Delete email address
[**GetEmail**](EmailApi.md#GetEmail) | **Get** /operators/{operator_id}/emails/{email_id} | Get email address
[**IssueAddEmailToken**](EmailApi.md#IssueAddEmailToken) | **Post** /operators/add_email_token/issue | Issue a token to add an email address
[**ListEmails**](EmailApi.md#ListEmails) | **Get** /operators/{operator_id}/emails | List email addresses
[**VerifyAddEmailToken**](EmailApi.md#VerifyAddEmailToken) | **Post** /operators/add_email_token/verify | Verify a token to add an email address



## DeleteEmail

> DeleteEmail(ctx, operatorId, emailId).Execute()

Delete email address



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
    emailId := "emailId_example" // string | email_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailApi.DeleteEmail(context.Background(), operatorId, emailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailApi.DeleteEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**emailId** | **string** | email_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEmailRequest struct via the builder pattern


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


## GetEmail

> EmailsModel GetEmail(ctx, operatorId, emailId).Execute()

Get email address



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
    emailId := "emailId_example" // string | email_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailApi.GetEmail(context.Background(), operatorId, emailId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailApi.GetEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEmail`: EmailsModel
    fmt.Fprintf(os.Stdout, "Response from `EmailApi.GetEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**emailId** | **string** | email_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EmailsModel**](EmailsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueAddEmailToken

> IssueAddEmailToken(ctx).IssueAddEmailTokenRequest(issueAddEmailTokenRequest).Execute()

Issue a token to add an email address



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
    issueAddEmailTokenRequest := *openapiclient.NewIssueAddEmailTokenRequest("Email_example", "Password_example") // IssueAddEmailTokenRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailApi.IssueAddEmailToken(context.Background()).IssueAddEmailTokenRequest(issueAddEmailTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailApi.IssueAddEmailToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueAddEmailTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueAddEmailTokenRequest** | [**IssueAddEmailTokenRequest**](IssueAddEmailTokenRequest.md) | request | 

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


## ListEmails

> []EmailsModel ListEmails(ctx, operatorId).Execute()

List email addresses



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
    resp, r, err := apiClient.EmailApi.ListEmails(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailApi.ListEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmails`: []EmailsModel
    fmt.Fprintf(os.Stdout, "Response from `EmailApi.ListEmails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EmailsModel**](EmailsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyAddEmailToken

> VerifyAddEmailToken(ctx).VerifyAddEmailTokenRequest(verifyAddEmailTokenRequest).Execute()

Verify a token to add an email address



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
    verifyAddEmailTokenRequest := *openapiclient.NewVerifyAddEmailTokenRequest("Token_example") // VerifyAddEmailTokenRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmailApi.VerifyAddEmailToken(context.Background()).VerifyAddEmailTokenRequest(verifyAddEmailTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmailApi.VerifyAddEmailToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyAddEmailTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyAddEmailTokenRequest** | [**VerifyAddEmailTokenRequest**](VerifyAddEmailTokenRequest.md) | request | 

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

