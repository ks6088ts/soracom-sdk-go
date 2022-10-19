# \UserApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UserApi.md#CreateUser) | **Post** /operators/{operator_id}/users/{user_name} | Create User
[**CreateUserPassword**](UserApi.md#CreateUserPassword) | **Post** /operators/{operator_id}/users/{user_name}/password | Create Password
[**DeleteDefaultPermissions**](UserApi.md#DeleteDefaultPermissions) | **Delete** /operators/{operator_id}/users/default_permissions | すべての SAM ユーザーに適用されるデフォルト権限を削除する
[**DeleteUser**](UserApi.md#DeleteUser) | **Delete** /operators/{operator_id}/users/{user_name} | Delete User
[**DeleteUserAuthKey**](UserApi.md#DeleteUserAuthKey) | **Delete** /operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id} | Delete User AuthKey
[**DeleteUserPassword**](UserApi.md#DeleteUserPassword) | **Delete** /operators/{operator_id}/users/{user_name}/password | Delete Password
[**DeleteUserPermission**](UserApi.md#DeleteUserPermission) | **Delete** /operators/{operator_id}/users/{user_name}/permission | SAM ユーザーの権限を削除する
[**EnableUserMFA**](UserApi.md#EnableUserMFA) | **Post** /operators/{operator_id}/users/{user_name}/mfa | Enable SAM user&#39;s MFA
[**GenerateUserAuthKey**](UserApi.md#GenerateUserAuthKey) | **Post** /operators/{operator_id}/users/{user_name}/auth_keys | Generate AuthKey
[**GetDefaultPermissions**](UserApi.md#GetDefaultPermissions) | **Get** /operators/{operator_id}/users/default_permissions | すべての SAM ユーザーに適用されるデフォルト権限を取得する
[**GetUser**](UserApi.md#GetUser) | **Get** /operators/{operator_id}/users/{user_name} | Get User
[**GetUserAuthKey**](UserApi.md#GetUserAuthKey) | **Get** /operators/{operator_id}/users/{user_name}/auth_keys/{auth_key_id} | Get AuthKey
[**GetUserMFAStatus**](UserApi.md#GetUserMFAStatus) | **Get** /operators/{operator_id}/users/{user_name}/mfa | Get SAM user&#39;s MFA status
[**GetUserPermission**](UserApi.md#GetUserPermission) | **Get** /operators/{operator_id}/users/{user_name}/permission | Get User Permission
[**HasUserPassword**](UserApi.md#HasUserPassword) | **Get** /operators/{operator_id}/users/{user_name}/password | Has User Password
[**ListUserAuthKeys**](UserApi.md#ListUserAuthKeys) | **Get** /operators/{operator_id}/users/{user_name}/auth_keys | List User AuthKeys
[**ListUsers**](UserApi.md#ListUsers) | **Get** /operators/{operator_id}/users | List Users
[**RevokeUserAuthTokens**](UserApi.md#RevokeUserAuthTokens) | **Delete** /operators/{operator_id}/users/{user_name}/tokens | 指定された SAM ユーザーが発行した全ての API キーと API トークンを無効化します。
[**RevokeUserMFA**](UserApi.md#RevokeUserMFA) | **Delete** /operators/{operator_id}/users/{user_name}/mfa | Revoke SAM user&#39;s MFA
[**UpdateDefaultPermissions**](UserApi.md#UpdateDefaultPermissions) | **Put** /operators/{operator_id}/users/default_permissions | Update the default permissions
[**UpdateUser**](UserApi.md#UpdateUser) | **Put** /operators/{operator_id}/users/{user_name} | Update User
[**UpdateUserPassword**](UserApi.md#UpdateUserPassword) | **Put** /operators/{operator_id}/users/{user_name}/password | Update Password
[**UpdateUserPermission**](UserApi.md#UpdateUserPermission) | **Put** /operators/{operator_id}/users/{user_name}/permission | SAM ユーザーの権限を更新する
[**VerifyUserMFA**](UserApi.md#VerifyUserMFA) | **Post** /operators/{operator_id}/users/{user_name}/mfa/verify | Verify SAM user&#39;s MFA OTP code when MFA activation phase



## CreateUser

> CreateUser(ctx, operatorId, userName).CreateUserRequest(createUserRequest).Execute()

Create User



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
    userName := "userName_example" // string | user_name
    createUserRequest := *openapiclient.NewCreateUserRequest() // CreateUserRequest | description

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.CreateUser(context.Background(), operatorId, userName).CreateUserRequest(createUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createUserRequest** | [**CreateUserRequest**](CreateUserRequest.md) | description | 

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


## CreateUserPassword

> CreateUserPassword(ctx, operatorId, userName).CreateUserPasswordRequest(createUserPasswordRequest).Execute()

Create Password



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
    userName := "userName_example" // string | user_name
    createUserPasswordRequest := *openapiclient.NewCreateUserPasswordRequest() // CreateUserPasswordRequest | password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.CreateUserPassword(context.Background(), operatorId, userName).CreateUserPasswordRequest(createUserPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createUserPasswordRequest** | [**CreateUserPasswordRequest**](CreateUserPasswordRequest.md) | password | 

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


## DeleteDefaultPermissions

> DeleteDefaultPermissions(ctx, operatorId).Execute()

すべての SAM ユーザーに適用されるデフォルト権限を削除する



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
    operatorId := "operatorId_example" // string | オペレーターID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.DeleteDefaultPermissions(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteDefaultPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultPermissionsRequest struct via the builder pattern


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


## DeleteUser

> DeleteUser(ctx, operatorId, userName).Execute()

Delete User



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
    userName := "userName_example" // string | user_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.DeleteUser(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserAuthKey

> DeleteUserAuthKey(ctx, operatorId, userName, authKeyId).Execute()

Delete User AuthKey



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
    userName := "userName_example" // string | user_name
    authKeyId := "authKeyId_example" // string | auth_key_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.DeleteUserAuthKey(context.Background(), operatorId, userName, authKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUserAuthKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 
**authKeyId** | **string** | auth_key_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserAuthKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserPassword

> DeleteUserPassword(ctx, operatorId, userName).Execute()

Delete Password



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
    userName := "userName_example" // string | user_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.DeleteUserPassword(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserPermission

> DeleteUserPermission(ctx, operatorId, userName).Execute()

SAM ユーザーの権限を削除する



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
    operatorId := "operatorId_example" // string | オペレーター ID
    userName := "userName_example" // string | SAM ユーザー名

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.DeleteUserPermission(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.DeleteUserPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーター ID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserPermissionRequest struct via the builder pattern


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


## EnableUserMFA

> EnableMFAOTPResponse EnableUserMFA(ctx, operatorId, userName).Execute()

Enable SAM user's MFA



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
    operatorId := "operatorId_example" // string | オペレーターID
    userName := "userName_example" // string | SAM ユーザー名

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.EnableUserMFA(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.EnableUserMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableUserMFA`: EnableMFAOTPResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.EnableUserMFA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableUserMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnableMFAOTPResponse**](EnableMFAOTPResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateUserAuthKey

> GenerateUserAuthKeyResponse GenerateUserAuthKey(ctx, operatorId, userName).Execute()

Generate AuthKey



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
    userName := "userName_example" // string | user_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GenerateUserAuthKey(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GenerateUserAuthKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateUserAuthKey`: GenerateUserAuthKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GenerateUserAuthKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateUserAuthKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GenerateUserAuthKeyResponse**](GenerateUserAuthKeyResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultPermissions

> GetDefaultPermissionsResponse GetDefaultPermissions(ctx, operatorId).Execute()

すべての SAM ユーザーに適用されるデフォルト権限を取得する



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
    operatorId := "operatorId_example" // string | オペレーターID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetDefaultPermissions(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetDefaultPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDefaultPermissions`: GetDefaultPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetDefaultPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetDefaultPermissionsResponse**](GetDefaultPermissionsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> UserDetailResponse GetUser(ctx, operatorId, userName).Execute()

Get User



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
    userName := "userName_example" // string | user_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUser(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUser`: UserDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UserDetailResponse**](UserDetailResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAuthKey

> AuthKeyResponse GetUserAuthKey(ctx, operatorId, userName, authKeyId).Execute()

Get AuthKey



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
    userName := "userName_example" // string | user_name
    authKeyId := "authKeyId_example" // string | auth_key_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUserAuthKey(context.Background(), operatorId, userName, authKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserAuthKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserAuthKey`: AuthKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserAuthKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 
**authKeyId** | **string** | auth_key_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAuthKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**AuthKeyResponse**](AuthKeyResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserMFAStatus

> MFAStatusOfUseResponse GetUserMFAStatus(ctx, operatorId, userName).Execute()

Get SAM user's MFA status



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
    operatorId := "operatorId_example" // string | オペレーターID
    userName := "userName_example" // string | SAM ユーザー名

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUserMFAStatus(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserMFAStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserMFAStatus`: MFAStatusOfUseResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserMFAStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserMFAStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MFAStatusOfUseResponse**](MFAStatusOfUseResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPermission

> GetUserPermissionResponse GetUserPermission(ctx, operatorId, userName).Execute()

Get User Permission



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
    operatorId := "operatorId_example" // string | オペレーター ID
    userName := "userName_example" // string | SAM ユーザー名

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.GetUserPermission(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.GetUserPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPermission`: GetUserPermissionResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.GetUserPermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーター ID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUserPermissionResponse**](GetUserPermissionResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HasUserPassword

> GetUserPasswordResponse HasUserPassword(ctx, operatorId, userName).Execute()

Has User Password



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
    userName := "userName_example" // string | user_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.HasUserPassword(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.HasUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HasUserPassword`: GetUserPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.HasUserPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiHasUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetUserPasswordResponse**](GetUserPasswordResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserAuthKeys

> []AuthKeyResponse ListUserAuthKeys(ctx, operatorId, userName).Execute()

List User AuthKeys



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
    userName := "userName_example" // string | user_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.ListUserAuthKeys(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUserAuthKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserAuthKeys`: []AuthKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUserAuthKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserAuthKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]AuthKeyResponse**](AuthKeyResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []UserDetailResponse ListUsers(ctx, operatorId).Execute()

List Users



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
    resp, r, err := apiClient.UserApi.ListUsers(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []UserDetailResponse
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]UserDetailResponse**](UserDetailResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeUserAuthTokens

> RevokeUserAuthTokens(ctx, operatorId, userName).Execute()

指定された SAM ユーザーが発行した全ての API キーと API トークンを無効化します。



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
    operatorId := "operatorId_example" // string | オペレーターID
    userName := "userName_example" // string | SAM ユーザー名

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.RevokeUserAuthTokens(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.RevokeUserAuthTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserAuthTokensRequest struct via the builder pattern


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


## RevokeUserMFA

> RevokeUserMFA(ctx, operatorId, userName).Execute()

Revoke SAM user's MFA



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
    operatorId := "operatorId_example" // string | オペレーターID
    userName := "userName_example" // string | SAM ユーザー名

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.RevokeUserMFA(context.Background(), operatorId, userName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.RevokeUserMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeUserMFARequest struct via the builder pattern


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


## UpdateDefaultPermissions

> UpdateDefaultPermissions(ctx, operatorId).UpdateDefaultPermissionsRequest(updateDefaultPermissionsRequest).Execute()

Update the default permissions



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
    operatorId := "operatorId_example" // string | オペレーターID
    updateDefaultPermissionsRequest := *openapiclient.NewUpdateDefaultPermissionsRequest("Permissions_example") // UpdateDefaultPermissionsRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UpdateDefaultPermissions(context.Background(), operatorId).UpdateDefaultPermissionsRequest(updateDefaultPermissionsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateDefaultPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDefaultPermissionsRequest** | [**UpdateDefaultPermissionsRequest**](UpdateDefaultPermissionsRequest.md) | request | 

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


## UpdateUser

> UpdateUser(ctx, operatorId, userName).UpdateUserRequest(updateUserRequest).Execute()

Update User



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
    userName := "userName_example" // string | user_name
    updateUserRequest := *openapiclient.NewUpdateUserRequest() // UpdateUserRequest | description

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UpdateUser(context.Background(), operatorId, userName).UpdateUserRequest(updateUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateUserRequest** | [**UpdateUserRequest**](UpdateUserRequest.md) | description | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPassword

> UpdateUserPassword(ctx, operatorId, userName).UpdatePasswordRequest(updatePasswordRequest).Execute()

Update Password



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
    userName := "userName_example" // string | user_name
    updatePasswordRequest := *openapiclient.NewUpdatePasswordRequest("CurrentPassword_example", "NewPassword_example") // UpdatePasswordRequest | password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UpdateUserPassword(context.Background(), operatorId, userName).UpdatePasswordRequest(updatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**userName** | **string** | user_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updatePasswordRequest** | [**UpdatePasswordRequest**](UpdatePasswordRequest.md) | password | 

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


## UpdateUserPermission

> UpdateUserPermission(ctx, operatorId, userName).SetUserPermissionRequest(setUserPermissionRequest).Execute()

SAM ユーザーの権限を更新する



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
    operatorId := "operatorId_example" // string | オペレーター ID
    userName := "userName_example" // string | SAM ユーザー名
    setUserPermissionRequest := *openapiclient.NewSetUserPermissionRequest("Permission_example") // SetUserPermissionRequest | 権限

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.UpdateUserPermission(context.Background(), operatorId, userName).SetUserPermissionRequest(setUserPermissionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.UpdateUserPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーター ID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setUserPermissionRequest** | [**SetUserPermissionRequest**](SetUserPermissionRequest.md) | 権限 | 

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


## VerifyUserMFA

> VerifyUserMFA(ctx, operatorId, userName).MFAAuthenticationRequest(mFAAuthenticationRequest).Execute()

Verify SAM user's MFA OTP code when MFA activation phase



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
    operatorId := "operatorId_example" // string | オペレーターID
    userName := "userName_example" // string | SAM ユーザー名
    mFAAuthenticationRequest := *openapiclient.NewMFAAuthenticationRequest() // MFAAuthenticationRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.VerifyUserMFA(context.Background(), operatorId, userName).MFAAuthenticationRequest(mFAAuthenticationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.VerifyUserMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | オペレーターID | 
**userName** | **string** | SAM ユーザー名 | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyUserMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mFAAuthenticationRequest** | [**MFAAuthenticationRequest**](MFAAuthenticationRequest.md) | request | 

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

