# \LagoonApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLagoonUser**](LagoonApi.md#CreateLagoonUser) | **Post** /lagoon/users | SORACOM Lagoon のユーザーを新たに作成します。
[**DeleteLagoonUser**](LagoonApi.md#DeleteLagoonUser) | **Delete** /lagoon/users/{lagoon_user_id} | SORACOM Lagoon のユーザーを削除します。
[**GetImageLink**](LagoonApi.md#GetImageLink) | **Get** /lagoon/image/link | SORACOM Lagoon の提供するカスタムアイコンのイメージへのリンク提供
[**GetLagoonMigrationInfo**](LagoonApi.md#GetLagoonMigrationInfo) | **Get** /lagoon/migration | (廃止されました) この API は常に &#x60;410 GONE&#x60; を返却します: SORACOM Lagoon のバージョン移行に関連する情報の取得
[**InitializeLagoonDashboardPermissions**](LagoonApi.md#InitializeLagoonDashboardPermissions) | **Post** /lagoon/dashboards/{dashboard_id}/permissions/init | SORACOM Lagoon のダッシュボードの権限初期化
[**ListLagoonDashboardsPermissions**](LagoonApi.md#ListLagoonDashboardsPermissions) | **Get** /lagoon/dashboards/permissions | SORACOM Lagoon の全ダッシュボードの権限取得
[**ListLagoonLicensePackStatus**](LagoonApi.md#ListLagoonLicensePackStatus) | **Get** /lagoon/license_packs | SORACOM Lagoon のライセンスパックのステータス取得
[**ListLagoonUsers**](LagoonApi.md#ListLagoonUsers) | **Get** /lagoon/users | オペレーターに属する SORACOM Lagoon のユーザー一覧を取得します。
[**MigrateLagoon**](LagoonApi.md#MigrateLagoon) | **Post** /lagoon/migration | (廃止されました) この API は常に &#x60;410 GONE&#x60; を返却します: SORACOM Lagoon のバージョン移行の実行
[**RegisterLagoon**](LagoonApi.md#RegisterLagoon) | **Post** /lagoon/register | SORACOM Lagoon サービスを有効にします。
[**TerminateLagoon**](LagoonApi.md#TerminateLagoon) | **Post** /lagoon/terminate | SORACOM Lagoon サービスの使用を終了します。
[**UpdateLagoonDashboardPermissions**](LagoonApi.md#UpdateLagoonDashboardPermissions) | **Put** /lagoon/dashboards/{dashboard_id}/permissions | SORACOM Lagoon のダッシュボード権限更新
[**UpdateLagoonLicensePack**](LagoonApi.md#UpdateLagoonLicensePack) | **Put** /lagoon/license_packs | SORACOM Lagoon の追加ライセンスパック更新・変更
[**UpdateLagoonPlan**](LagoonApi.md#UpdateLagoonPlan) | **Put** /lagoon/plan | SORACOM Lagoon のプラン変更
[**UpdateLagoonUserEmail**](LagoonApi.md#UpdateLagoonUserEmail) | **Put** /lagoon/users/{lagoon_user_id}/email | SORACOM Lagoon のユーザーのメールアドレスを変更します。
[**UpdateLagoonUserPassword**](LagoonApi.md#UpdateLagoonUserPassword) | **Put** /lagoon/users/{lagoon_user_id}/password | SORACOM Lagoon のユーザーのパスワードを変更します。
[**UpdateLagoonUserPermission**](LagoonApi.md#UpdateLagoonUserPermission) | **Put** /lagoon/users/{lagoon_user_id}/permission | SORACOM Lagoon のユーザーの権限を変更します。



## CreateLagoonUser

> LagoonUserCreationResponse CreateLagoonUser(ctx).LagoonUserCreationRequest(lagoonUserCreationRequest).Execute()

SORACOM Lagoon のユーザーを新たに作成します。



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
    lagoonUserCreationRequest := *openapiclient.NewLagoonUserCreationRequest() // LagoonUserCreationRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.CreateLagoonUser(context.Background()).LagoonUserCreationRequest(lagoonUserCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.CreateLagoonUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLagoonUser`: LagoonUserCreationResponse
    fmt.Fprintf(os.Stdout, "Response from `LagoonApi.CreateLagoonUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLagoonUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lagoonUserCreationRequest** | [**LagoonUserCreationRequest**](LagoonUserCreationRequest.md) | request | 

### Return type

[**LagoonUserCreationResponse**](LagoonUserCreationResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLagoonUser

> DeleteLagoonUser(ctx, lagoonUserId).Execute()

SORACOM Lagoon のユーザーを削除します。



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
    lagoonUserId := int32(56) // int32 | Target ID of the lagoon user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.DeleteLagoonUser(context.Background(), lagoonUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.DeleteLagoonUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoonUserId** | **int32** | Target ID of the lagoon user | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLagoonUserRequest struct via the builder pattern


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


## GetImageLink

> string GetImageLink(ctx).Classic(classic).Execute()

SORACOM Lagoon の提供するカスタムアイコンのイメージへのリンク提供



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
    classic := true // bool | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.GetImageLink(context.Background()).Classic(classic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.GetImageLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetImageLink`: string
    fmt.Fprintf(os.Stdout, "Response from `LagoonApi.GetImageLink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetImageLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classic** | **bool** | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. | 

### Return type

**string**

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLagoonMigrationInfo

> GetLagoonMigrationInfo(ctx).LagoonPlanChangingRequest(lagoonPlanChangingRequest).Execute()

(廃止されました) この API は常に `410 GONE` を返却します: SORACOM Lagoon のバージョン移行に関連する情報の取得



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
    lagoonPlanChangingRequest := *openapiclient.NewLagoonPlanChangingRequest() // LagoonPlanChangingRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.GetLagoonMigrationInfo(context.Background()).LagoonPlanChangingRequest(lagoonPlanChangingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.GetLagoonMigrationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLagoonMigrationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lagoonPlanChangingRequest** | [**LagoonPlanChangingRequest**](LagoonPlanChangingRequest.md) | req | 

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


## InitializeLagoonDashboardPermissions

> InitializeLagoonDashboardPermissions(ctx, dashboardId).Classic(classic).Execute()

SORACOM Lagoon のダッシュボードの権限初期化



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
    dashboardId := int64(789) // int64 | dashboard_id
    classic := true // bool | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.InitializeLagoonDashboardPermissions(context.Background(), dashboardId).Classic(classic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.InitializeLagoonDashboardPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | dashboard_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiInitializeLagoonDashboardPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **classic** | **bool** | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. | 

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


## ListLagoonDashboardsPermissions

> []LagoonDashboardPermissionsResponse ListLagoonDashboardsPermissions(ctx).Classic(classic).Execute()

SORACOM Lagoon の全ダッシュボードの権限取得



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
    classic := true // bool | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.ListLagoonDashboardsPermissions(context.Background()).Classic(classic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.ListLagoonDashboardsPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLagoonDashboardsPermissions`: []LagoonDashboardPermissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `LagoonApi.ListLagoonDashboardsPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLagoonDashboardsPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classic** | **bool** | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. | 

### Return type

[**[]LagoonDashboardPermissionsResponse**](LagoonDashboardPermissionsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLagoonLicensePackStatus

> []LagoonLicensePackStatusResponse ListLagoonLicensePackStatus(ctx).Execute()

SORACOM Lagoon のライセンスパックのステータス取得



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
    resp, r, err := apiClient.LagoonApi.ListLagoonLicensePackStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.ListLagoonLicensePackStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLagoonLicensePackStatus`: []LagoonLicensePackStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `LagoonApi.ListLagoonLicensePackStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLagoonLicensePackStatusRequest struct via the builder pattern


### Return type

[**[]LagoonLicensePackStatusResponse**](LagoonLicensePackStatusResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLagoonUsers

> []LagoonUser ListLagoonUsers(ctx).Classic(classic).Execute()

オペレーターに属する SORACOM Lagoon のユーザー一覧を取得します。



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
    classic := true // bool | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.ListLagoonUsers(context.Background()).Classic(classic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.ListLagoonUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLagoonUsers`: []LagoonUser
    fmt.Fprintf(os.Stdout, "Response from `LagoonApi.ListLagoonUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLagoonUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **classic** | **bool** | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. | 

### Return type

[**[]LagoonUser**](LagoonUser.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateLagoon

> MigrateLagoon(ctx).LagoonMigrationFromClassicRequest(lagoonMigrationFromClassicRequest).Execute()

(廃止されました) この API は常に `410 GONE` を返却します: SORACOM Lagoon のバージョン移行の実行



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
    lagoonMigrationFromClassicRequest := *openapiclient.NewLagoonMigrationFromClassicRequest() // LagoonMigrationFromClassicRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.MigrateLagoon(context.Background()).LagoonMigrationFromClassicRequest(lagoonMigrationFromClassicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.MigrateLagoon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateLagoonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lagoonMigrationFromClassicRequest** | [**LagoonMigrationFromClassicRequest**](LagoonMigrationFromClassicRequest.md) | req | 

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


## RegisterLagoon

> LagoonRegistrationResponse RegisterLagoon(ctx).LagoonRegistrationRequest(lagoonRegistrationRequest).Execute()

SORACOM Lagoon サービスを有効にします。



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
    lagoonRegistrationRequest := *openapiclient.NewLagoonRegistrationRequest() // LagoonRegistrationRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.RegisterLagoon(context.Background()).LagoonRegistrationRequest(lagoonRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.RegisterLagoon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterLagoon`: LagoonRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `LagoonApi.RegisterLagoon`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterLagoonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lagoonRegistrationRequest** | [**LagoonRegistrationRequest**](LagoonRegistrationRequest.md) | request | 

### Return type

[**LagoonRegistrationResponse**](LagoonRegistrationResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateLagoon

> TerminateLagoon(ctx).Execute()

SORACOM Lagoon サービスの使用を終了します。



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
    resp, r, err := apiClient.LagoonApi.TerminateLagoon(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.TerminateLagoon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateLagoonRequest struct via the builder pattern


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


## UpdateLagoonDashboardPermissions

> UpdateLagoonDashboardPermissions(ctx, dashboardId).LagoonDashboardPermissionsUpdatingRequest(lagoonDashboardPermissionsUpdatingRequest).Classic(classic).Execute()

SORACOM Lagoon のダッシュボード権限更新



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
    dashboardId := int64(789) // int64 | dashboard_id
    lagoonDashboardPermissionsUpdatingRequest := *openapiclient.NewLagoonDashboardPermissionsUpdatingRequest() // LagoonDashboardPermissionsUpdatingRequest | req
    classic := true // bool | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.UpdateLagoonDashboardPermissions(context.Background(), dashboardId).LagoonDashboardPermissionsUpdatingRequest(lagoonDashboardPermissionsUpdatingRequest).Classic(classic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.UpdateLagoonDashboardPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | dashboard_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLagoonDashboardPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lagoonDashboardPermissionsUpdatingRequest** | [**LagoonDashboardPermissionsUpdatingRequest**](LagoonDashboardPermissionsUpdatingRequest.md) | req | 
 **classic** | **bool** | If the value is true, a request will be issued to Lagoon Classic. This is only valid if both Lagoon and Lagoon Classic are enabled. | 

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


## UpdateLagoonLicensePack

> UpdateLagoonLicensePack(ctx).LagoonLicensePacksUpdatingRequest(lagoonLicensePacksUpdatingRequest).Execute()

SORACOM Lagoon の追加ライセンスパック更新・変更



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
    lagoonLicensePacksUpdatingRequest := *openapiclient.NewLagoonLicensePacksUpdatingRequest() // LagoonLicensePacksUpdatingRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.UpdateLagoonLicensePack(context.Background()).LagoonLicensePacksUpdatingRequest(lagoonLicensePacksUpdatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.UpdateLagoonLicensePack``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLagoonLicensePackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lagoonLicensePacksUpdatingRequest** | [**LagoonLicensePacksUpdatingRequest**](LagoonLicensePacksUpdatingRequest.md) | req | 

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


## UpdateLagoonPlan

> UpdateLagoonPlan(ctx).LagoonPlanChangingRequest(lagoonPlanChangingRequest).Execute()

SORACOM Lagoon のプラン変更



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
    lagoonPlanChangingRequest := *openapiclient.NewLagoonPlanChangingRequest() // LagoonPlanChangingRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.UpdateLagoonPlan(context.Background()).LagoonPlanChangingRequest(lagoonPlanChangingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.UpdateLagoonPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLagoonPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lagoonPlanChangingRequest** | [**LagoonPlanChangingRequest**](LagoonPlanChangingRequest.md) | req | 

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


## UpdateLagoonUserEmail

> UpdateLagoonUserEmail(ctx, lagoonUserId).LagoonUserEmailUpdatingRequest(lagoonUserEmailUpdatingRequest).Execute()

SORACOM Lagoon のユーザーのメールアドレスを変更します。



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
    lagoonUserId := int32(56) // int32 | Target ID of the lagoon user
    lagoonUserEmailUpdatingRequest := *openapiclient.NewLagoonUserEmailUpdatingRequest() // LagoonUserEmailUpdatingRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.UpdateLagoonUserEmail(context.Background(), lagoonUserId).LagoonUserEmailUpdatingRequest(lagoonUserEmailUpdatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.UpdateLagoonUserEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoonUserId** | **int32** | Target ID of the lagoon user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLagoonUserEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lagoonUserEmailUpdatingRequest** | [**LagoonUserEmailUpdatingRequest**](LagoonUserEmailUpdatingRequest.md) | request | 

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


## UpdateLagoonUserPassword

> UpdateLagoonUserPassword(ctx, lagoonUserId).LagoonUserPasswordUpdatingRequest(lagoonUserPasswordUpdatingRequest).Execute()

SORACOM Lagoon のユーザーのパスワードを変更します。



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
    lagoonUserId := int32(56) // int32 | Target ID of the lagoon user
    lagoonUserPasswordUpdatingRequest := *openapiclient.NewLagoonUserPasswordUpdatingRequest() // LagoonUserPasswordUpdatingRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.UpdateLagoonUserPassword(context.Background(), lagoonUserId).LagoonUserPasswordUpdatingRequest(lagoonUserPasswordUpdatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.UpdateLagoonUserPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoonUserId** | **int32** | Target ID of the lagoon user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLagoonUserPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lagoonUserPasswordUpdatingRequest** | [**LagoonUserPasswordUpdatingRequest**](LagoonUserPasswordUpdatingRequest.md) | request | 

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


## UpdateLagoonUserPermission

> UpdateLagoonUserPermission(ctx, lagoonUserId).LagoonUserPermissionUpdatingRequest(lagoonUserPermissionUpdatingRequest).Execute()

SORACOM Lagoon のユーザーの権限を変更します。



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
    lagoonUserId := int32(56) // int32 | Target ID of the lagoon user
    lagoonUserPermissionUpdatingRequest := *openapiclient.NewLagoonUserPermissionUpdatingRequest() // LagoonUserPermissionUpdatingRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LagoonApi.UpdateLagoonUserPermission(context.Background(), lagoonUserId).LagoonUserPermissionUpdatingRequest(lagoonUserPermissionUpdatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LagoonApi.UpdateLagoonUserPermission``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lagoonUserId** | **int32** | Target ID of the lagoon user | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLagoonUserPermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lagoonUserPermissionUpdatingRequest** | [**LagoonUserPermissionUpdatingRequest**](LagoonUserPermissionUpdatingRequest.md) | request | 

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

