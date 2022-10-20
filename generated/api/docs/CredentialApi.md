# \CredentialApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredential**](CredentialApi.md#CreateCredential) | **Post** /credentials/{credentials_id} | 認証情報の作成
[**DeleteCredential**](CredentialApi.md#DeleteCredential) | **Delete** /credentials/{credentials_id} | 認証情報の削除
[**ListCredentials**](CredentialApi.md#ListCredentials) | **Get** /credentials | 認証情報の表示
[**UpdateCredential**](CredentialApi.md#UpdateCredential) | **Put** /credentials/{credentials_id} | 認証情報の更新



## CreateCredential

> CredentialsModel CreateCredential(ctx, credentialsId).CreateAndUpdateCredentialsModel(createAndUpdateCredentialsModel).Execute()

認証情報の作成



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
    credentialsId := "credentialsId_example" // string | credentials_id
    createAndUpdateCredentialsModel := *openapiclient.NewCreateAndUpdateCredentialsModel() // CreateAndUpdateCredentialsModel | credentials

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialApi.CreateCredential(context.Background(), credentialsId).CreateAndUpdateCredentialsModel(createAndUpdateCredentialsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialApi.CreateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredential`: CredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialApi.CreateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | credentials_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAndUpdateCredentialsModel** | [**CreateAndUpdateCredentialsModel**](CreateAndUpdateCredentialsModel.md) | credentials | 

### Return type

[**CredentialsModel**](CredentialsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredential

> DeleteCredential(ctx, credentialsId).Execute()

認証情報の削除



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
    credentialsId := "credentialsId_example" // string | Credentials ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialApi.DeleteCredential(context.Background(), credentialsId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialApi.DeleteCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | Credentials ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialRequest struct via the builder pattern


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


## ListCredentials

> []CredentialsModel ListCredentials(ctx).Execute()

認証情報の表示



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
    resp, r, err := apiClient.CredentialApi.ListCredentials(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialApi.ListCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCredentials`: []CredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialApi.ListCredentials`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCredentialsRequest struct via the builder pattern


### Return type

[**[]CredentialsModel**](CredentialsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredential

> CredentialsModel UpdateCredential(ctx, credentialsId).CreateAndUpdateCredentialsModel(createAndUpdateCredentialsModel).Execute()

認証情報の更新



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
    credentialsId := "credentialsId_example" // string | credentials_id
    createAndUpdateCredentialsModel := *openapiclient.NewCreateAndUpdateCredentialsModel() // CreateAndUpdateCredentialsModel | credentials

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialApi.UpdateCredential(context.Background(), credentialsId).CreateAndUpdateCredentialsModel(createAndUpdateCredentialsModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialApi.UpdateCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredential`: CredentialsModel
    fmt.Fprintf(os.Stdout, "Response from `CredentialApi.UpdateCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**credentialsId** | **string** | credentials_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createAndUpdateCredentialsModel** | [**CreateAndUpdateCredentialsModel**](CreateAndUpdateCredentialsModel.md) | credentials | 

### Return type

[**CredentialsModel**](CredentialsModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json;charset=UTF-8
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

