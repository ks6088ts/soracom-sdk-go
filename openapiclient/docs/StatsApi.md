# \StatsApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportAirStats**](StatsApi.md#ExportAirStats) | **Post** /stats/air/operators/{operator_id}/export | Export Air Usage Report of All Subscribers
[**ExportBeamStats**](StatsApi.md#ExportBeamStats) | **Post** /stats/beam/operators/{operator_id}/export | Export Beam Usage Report of All Subscribers
[**ExportFunkStats**](StatsApi.md#ExportFunkStats) | **Post** /stats/funk/operators/{operator_id}/export | Export Funk Usage Report of All Subscribers
[**ExportFunnelStats**](StatsApi.md#ExportFunnelStats) | **Post** /stats/funnel/operators/{operator_id}/export | Export Funnel Usage Report of All Subscribers
[**GetAirStats**](StatsApi.md#GetAirStats) | **Get** /stats/air/subscribers/{imsi} | Get Air Usage Report of Subscriber
[**GetAirStatsOfSim**](StatsApi.md#GetAirStatsOfSim) | **Get** /stats/air/sims/{sim_id} | Get Air Usage Report of SIM
[**GetBeamStats**](StatsApi.md#GetBeamStats) | **Get** /stats/beam/subscribers/{imsi} | Get Beam Usage Report of Subscriber
[**GetFunkStats**](StatsApi.md#GetFunkStats) | **Get** /stats/funk/subscribers/{imsi} | Get Funk Usage Report of Subscriber
[**GetFunnelStats**](StatsApi.md#GetFunnelStats) | **Get** /stats/funnel/subscribers/{imsi} | Get Funnel Usage Report of Subscriber
[**GetHarvestExportedDataStats**](StatsApi.md#GetHarvestExportedDataStats) | **Get** /stats/harvest/operators/{operator_id} | Harvest の利用統計情報を取得します。
[**GetHarvestStats**](StatsApi.md#GetHarvestStats) | **Get** /stats/harvest/subscribers/{imsi} | Get Harvest Usage Report of Subscriber
[**GetNapterAuditLogsExportedDataStats**](StatsApi.md#GetNapterAuditLogsExportedDataStats) | **Get** /stats/napter/audit_logs | Napter 監査ログの月次読み込みデータ量を取得する。



## ExportAirStats

> FileExportResponse ExportAirStats(ctx, operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()

Export Air Usage Report of All Subscribers



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
    exportRequest := *openapiclient.NewExportRequest() // ExportRequest | 出力するデータの期間
    exportMode := "exportMode_example" // string | export_mode (async, sync) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.ExportAirStats(context.Background(), operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.ExportAirStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportAirStats`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.ExportAirStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportAirStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportRequest** | [**ExportRequest**](ExportRequest.md) | 出力するデータの期間 | 
 **exportMode** | **string** | export_mode (async, sync) | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportBeamStats

> FileExportResponse ExportBeamStats(ctx, operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()

Export Beam Usage Report of All Subscribers



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
    operatorId := "operatorId_example" // string | operator ID
    exportRequest := *openapiclient.NewExportRequest() // ExportRequest | 出力するデータの期間
    exportMode := "exportMode_example" // string | export_mode (async, sync) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.ExportBeamStats(context.Background(), operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.ExportBeamStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportBeamStats`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.ExportBeamStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportBeamStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportRequest** | [**ExportRequest**](ExportRequest.md) | 出力するデータの期間 | 
 **exportMode** | **string** | export_mode (async, sync) | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFunkStats

> FileExportResponse ExportFunkStats(ctx, operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()

Export Funk Usage Report of All Subscribers



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
    operatorId := "operatorId_example" // string | operator ID
    exportRequest := *openapiclient.NewExportRequest() // ExportRequest | 出力するデータの期間
    exportMode := "exportMode_example" // string | export_mode (async, sync) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.ExportFunkStats(context.Background(), operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.ExportFunkStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFunkStats`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.ExportFunkStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFunkStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportRequest** | [**ExportRequest**](ExportRequest.md) | 出力するデータの期間 | 
 **exportMode** | **string** | export_mode (async, sync) | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportFunnelStats

> FileExportResponse ExportFunnelStats(ctx, operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()

Export Funnel Usage Report of All Subscribers



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
    operatorId := "operatorId_example" // string | operator ID
    exportRequest := *openapiclient.NewExportRequest() // ExportRequest | 出力するデータの期間
    exportMode := "exportMode_example" // string | export_mode (async, sync) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.ExportFunnelStats(context.Background(), operatorId).ExportRequest(exportRequest).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.ExportFunnelStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportFunnelStats`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.ExportFunnelStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportFunnelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportRequest** | [**ExportRequest**](ExportRequest.md) | 出力するデータの期間 | 
 **exportMode** | **string** | export_mode (async, sync) | 

### Return type

[**FileExportResponse**](FileExportResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAirStats

> []AirStatsResponse GetAirStats(ctx, imsi).From(from).To(to).Period(period).Execute()

Get Air Usage Report of Subscriber



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
    imsi := "imsi_example" // string | imsi
    from := int32(56) // int32 | 集計対象時刻の始まりを unixtime で与える
    to := int32(56) // int32 | 集計対象時刻の終わりを unixtime で与える
    period := "period_example" // string | 集計単位。minutes は最も細かい粒度で通信量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で通信量が記録されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetAirStats(context.Background(), imsi).From(from).To(to).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetAirStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAirStats`: []AirStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetAirStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 集計対象時刻の始まりを unixtime で与える | 
 **to** | **int32** | 集計対象時刻の終わりを unixtime で与える | 
 **period** | **string** | 集計単位。minutes は最も細かい粒度で通信量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で通信量が記録されます。 | 

### Return type

[**[]AirStatsResponse**](AirStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAirStatsOfSim

> []AirStatsResponse GetAirStatsOfSim(ctx, simId).From(from).To(to).Period(period).Execute()

Get Air Usage Report of SIM



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
    simId := "simId_example" // string | SIM ID
    from := int32(56) // int32 | 集計対象時刻の始まりを UNIX 時間 (秒) で与える
    to := int32(56) // int32 | 集計対象時刻の終わりを UNIX 時間 (秒) で与える
    period := "period_example" // string | 集計単位。minutes は最も細かい粒度で通信量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で通信量が記録されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetAirStatsOfSim(context.Background(), simId).From(from).To(to).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetAirStatsOfSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAirStatsOfSim`: []AirStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetAirStatsOfSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAirStatsOfSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 集計対象時刻の始まりを UNIX 時間 (秒) で与える | 
 **to** | **int32** | 集計対象時刻の終わりを UNIX 時間 (秒) で与える | 
 **period** | **string** | 集計単位。minutes は最も細かい粒度で通信量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で通信量が記録されます。 | 

### Return type

[**[]AirStatsResponse**](AirStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBeamStats

> []BeamStatsResponse GetBeamStats(ctx, imsi).From(from).To(to).Period(period).Execute()

Get Beam Usage Report of Subscriber



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
    imsi := "imsi_example" // string | imsi
    from := int32(56) // int32 | 集計対象時刻の始まりを unixtime で与える
    to := int32(56) // int32 | 集計対象時刻の終わりを unixtime で与える
    period := "period_example" // string | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetBeamStats(context.Background(), imsi).From(from).To(to).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetBeamStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBeamStats`: []BeamStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetBeamStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBeamStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 集計対象時刻の始まりを unixtime で与える | 
 **to** | **int32** | 集計対象時刻の終わりを unixtime で与える | 
 **period** | **string** | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。 | 

### Return type

[**[]BeamStatsResponse**](BeamStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunkStats

> []FunkStatsResponse GetFunkStats(ctx, imsi).From(from).To(to).Period(period).Execute()

Get Funk Usage Report of Subscriber



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
    imsi := "imsi_example" // string | imsi
    from := int32(56) // int32 | 集計対象時刻の始まりを unixtime で与える
    to := int32(56) // int32 | 集計対象時刻の終わりを unixtime で与える
    period := "period_example" // string | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetFunkStats(context.Background(), imsi).From(from).To(to).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetFunkStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunkStats`: []FunkStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetFunkStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunkStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 集計対象時刻の始まりを unixtime で与える | 
 **to** | **int32** | 集計対象時刻の終わりを unixtime で与える | 
 **period** | **string** | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。 | 

### Return type

[**[]FunkStatsResponse**](FunkStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFunnelStats

> []FunnelStatsResponse GetFunnelStats(ctx, imsi).From(from).To(to).Period(period).Execute()

Get Funnel Usage Report of Subscriber



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
    imsi := "imsi_example" // string | imsi
    from := int32(56) // int32 | 集計対象時刻の始まりを unixtime で与える
    to := int32(56) // int32 | 集計対象時刻の終わりを unixtime で与える
    period := "period_example" // string | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetFunnelStats(context.Background(), imsi).From(from).To(to).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetFunnelStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFunnelStats`: []FunnelStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetFunnelStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFunnelStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 集計対象時刻の始まりを unixtime で与える | 
 **to** | **int32** | 集計対象時刻の終わりを unixtime で与える | 
 **period** | **string** | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。 | 

### Return type

[**[]FunnelStatsResponse**](FunnelStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHarvestExportedDataStats

> HarvestExportedDataStatsResponse GetHarvestExportedDataStats(ctx, operatorId).YearMonth(yearMonth).Execute()

Harvest の利用統計情報を取得します。



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
    operatorId := "operatorId_example" // string | operator ID
    yearMonth := "yearMonth_example" // string | 年月を YYYYMM 形式で指定します。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetHarvestExportedDataStats(context.Background(), operatorId).YearMonth(yearMonth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetHarvestExportedDataStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHarvestExportedDataStats`: HarvestExportedDataStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetHarvestExportedDataStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHarvestExportedDataStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **yearMonth** | **string** | 年月を YYYYMM 形式で指定します。 | 

### Return type

[**HarvestExportedDataStatsResponse**](HarvestExportedDataStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetHarvestStats

> []HarvestStatsResponse GetHarvestStats(ctx, imsi).From(from).To(to).Period(period).Execute()

Get Harvest Usage Report of Subscriber



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
    imsi := "imsi_example" // string | imsi
    from := int32(56) // int32 | 集計対象時刻の始まりを unixtime で与える
    to := int32(56) // int32 | 集計対象時刻の終わりを unixtime で与える
    period := "period_example" // string | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetHarvestStats(context.Background(), imsi).From(from).To(to).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetHarvestStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHarvestStats`: []HarvestStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetHarvestStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | imsi | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHarvestStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 集計対象時刻の始まりを unixtime で与える | 
 **to** | **int32** | 集計対象時刻の終わりを unixtime で与える | 
 **period** | **string** | 集計単位。minutes は最も細かい粒度で利用量履歴を出力します。ただし、デバイスが SORACOM プラットフォームに接続している間は、約 5 分間隔で利用量が記録されます。 | 

### Return type

[**[]HarvestStatsResponse**](HarvestStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNapterAuditLogsExportedDataStats

> NapterAuditLogsExportedDataStatsResponse GetNapterAuditLogsExportedDataStats(ctx).YearMonth(yearMonth).Execute()

Napter 監査ログの月次読み込みデータ量を取得する。



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
    yearMonth := "yearMonth_example" // string | 年月を YYYYMM 形式で指定します。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StatsApi.GetNapterAuditLogsExportedDataStats(context.Background()).YearMonth(yearMonth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StatsApi.GetNapterAuditLogsExportedDataStats``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNapterAuditLogsExportedDataStats`: NapterAuditLogsExportedDataStatsResponse
    fmt.Fprintf(os.Stdout, "Response from `StatsApi.GetNapterAuditLogsExportedDataStats`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNapterAuditLogsExportedDataStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **yearMonth** | **string** | 年月を YYYYMM 形式で指定します。 | 

### Return type

[**NapterAuditLogsExportedDataStatsResponse**](NapterAuditLogsExportedDataStatsResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

