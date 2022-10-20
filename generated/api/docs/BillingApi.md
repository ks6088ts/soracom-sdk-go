# \BillingApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportBilling**](BillingApi.md#ExportBilling) | **Post** /bills/{yyyyMM}/export | 指定月の課金詳細情報 CSV をダウンロードするための URL を発行する
[**ExportLatestBilling**](BillingApi.md#ExportLatestBilling) | **Post** /bills/latest/export | Export latest billing CSV file to S3
[**GetBilling**](BillingApi.md#GetBilling) | **Get** /bills/{yyyyMM} | 指定した月の確定した利用料金を取得する
[**GetBillingHistory**](BillingApi.md#GetBillingHistory) | **Get** /bills | 過去の確定した利用料金を取得する
[**GetBillingPerDay**](BillingApi.md#GetBillingPerDay) | **Get** /bills/{yyyyMM}/daily | 日ごとの利用料金を取得する
[**GetLatestBilling**](BillingApi.md#GetLatestBilling) | **Get** /bills/latest | Get latest bill



## ExportBilling

> FileExportResponse ExportBilling(ctx, yyyyMM).ExportMode(exportMode).Execute()

指定月の課金詳細情報 CSV をダウンロードするための URL を発行する



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
    yyyyMM := "yyyyMM_example" // string | 対象年月
    exportMode := "exportMode_example" // string | 課金詳細情報 CSV をダウンロードするための URL を取得する方法を指定します。  - `async`: SORACOM プラットフォームで URL が発行されるのを待たずに、`exportedFieldId` を取得します。この `exportedFieldId` を [`Files:getExportedFile API`](#/Files/getExportedFile) で指定すると、URL を取得できます。課金詳細情報 CSV のファイルサイズが大きい場合は、`async` を利用してください。 - `sync` (デフォルト): SORACOM プラットフォームで URL が発行されるのを待ちます。ただし、課金詳細情報 CSV のファイルサイズが大きい場合など、タイムアウトして URL を取得できないことがあります。タイムアウトする場合は、`async` を指定してください。  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.ExportBilling(context.Background(), yyyyMM).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ExportBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBilling`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ExportBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**yyyyMM** | **string** | 対象年月 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportMode** | **string** | 課金詳細情報 CSV をダウンロードするための URL を取得する方法を指定します。  - &#x60;async&#x60;: SORACOM プラットフォームで URL が発行されるのを待たずに、&#x60;exportedFieldId&#x60; を取得します。この &#x60;exportedFieldId&#x60; を [&#x60;Files:getExportedFile API&#x60;](#/Files/getExportedFile) で指定すると、URL を取得できます。課金詳細情報 CSV のファイルサイズが大きい場合は、&#x60;async&#x60; を利用してください。 - &#x60;sync&#x60; (デフォルト): SORACOM プラットフォームで URL が発行されるのを待ちます。ただし、課金詳細情報 CSV のファイルサイズが大きい場合など、タイムアウトして URL を取得できないことがあります。タイムアウトする場合は、&#x60;async&#x60; を指定してください。  | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportLatestBilling

> FileExportResponse ExportLatestBilling(ctx).ExportMode(exportMode).Execute()

Export latest billing CSV file to S3



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
    exportMode := "exportMode_example" // string | export_mode (async, sync) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.ExportLatestBilling(context.Background()).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ExportLatestBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportLatestBilling`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ExportLatestBilling`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportLatestBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportMode** | **string** | export_mode (async, sync) | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBilling

> MonthlyBill GetBilling(ctx, yyyyMM).Execute()

指定した月の確定した利用料金を取得する



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
    yyyyMM := "yyyyMM_example" // string | 対象年月

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.GetBilling(context.Background(), yyyyMM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBilling`: MonthlyBill
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBilling`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**yyyyMM** | **string** | 対象年月 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MonthlyBill**](MonthlyBill.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingHistory

> GetBillingHistoryResponse GetBillingHistory(ctx).Execute()

過去の確定した利用料金を取得する



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
    resp, r, err := apiClient.BillingApi.GetBillingHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBillingHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingHistory`: GetBillingHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBillingHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingHistoryRequest struct via the builder pattern


### Return type

[**GetBillingHistoryResponse**](GetBillingHistoryResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBillingPerDay

> DailyBillResponse GetBillingPerDay(ctx, yyyyMM).Execute()

日ごとの利用料金を取得する



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
    yyyyMM := "yyyyMM_example" // string | 対象年月

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.GetBillingPerDay(context.Background(), yyyyMM).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetBillingPerDay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBillingPerDay`: DailyBillResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetBillingPerDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**yyyyMM** | **string** | 対象年月 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillingPerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DailyBillResponse**](DailyBillResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLatestBilling

> GetLatestBill GetLatestBilling(ctx).Execute()

Get latest bill



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
    resp, r, err := apiClient.BillingApi.GetLatestBilling(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.GetLatestBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestBilling`: GetLatestBill
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.GetLatestBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestBillingRequest struct via the builder pattern


### Return type

[**GetLatestBill**](GetLatestBill.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

