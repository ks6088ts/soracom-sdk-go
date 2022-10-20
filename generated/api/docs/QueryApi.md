# \QueryApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchDevices**](QueryApi.md#SearchDevices) | **Get** /query/devices | SORACOM Inventory デバイスをクエリに応じて検索します。
[**SearchSigfoxDevices**](QueryApi.md#SearchSigfoxDevices) | **Get** /query/sigfox_devices | Sigfox デバイスをクエリに応じて検索します。
[**SearchSims**](QueryApi.md#SearchSims) | **Get** /query/sims | SIM をクエリに応じて検索します。
[**SearchSubscriberTrafficVolumeRanking**](QueryApi.md#SearchSubscriberTrafficVolumeRanking) | **Get** /query/subscribers/traffic_volume/ranking | Subscriber の通信量ランキングを返却します。
[**SearchSubscribers**](QueryApi.md#SearchSubscribers) | **Get** /query/subscribers | (非推奨) Subscriber をクエリに応じて検索します。



## SearchDevices

> []Device SearchDevices(ctx).Name(name).Group(group).DeviceId(deviceId).Tag(tag).Imsi(imsi).Imei(imei).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()

SORACOM Inventory デバイスをクエリに応じて検索します。



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
    name := []string{"Inner_example"} // []string | 検索したい SORACOM Inventory デバイスの名前 (optional)
    group := []string{"Inner_example"} // []string | 検索したいグループの名前 (optional)
    deviceId := []string{"Inner_example"} // []string | 検索したい SORACOM Inventory デバイスの ID (optional)
    tag := []string{"Inner_example"} // []string | 検索したいタグの値の文字列 (optional)
    imsi := []string{"Inner_example"} // []string | 検索したい SORACOM Inventory デバイスの bootstrap 時の IMSI (optional)
    imei := []string{"Inner_example"} // []string | 検索したい SORACOM Inventory デバイスの bootstrap 時の IMEI (optional)
    limit := int32(56) // int32 | 取得する結果の上限数 (optional) (default to 10)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の SORACOM Inventory デバイスの ID。このパラメータを指定することで次の SORACOM Inventory デバイス以降のリストを取得できる。 (optional)
    searchType := "searchType_example" // string | 検索の種別 (AND 検索もしくは OR 検索) (optional) (default to "and")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.SearchDevices(context.Background()).Name(name).Group(group).DeviceId(deviceId).Tag(tag).Imsi(imsi).Imei(imei).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.SearchDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchDevices`: []Device
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.SearchDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | 検索したい SORACOM Inventory デバイスの名前 | 
 **group** | **[]string** | 検索したいグループの名前 | 
 **deviceId** | **[]string** | 検索したい SORACOM Inventory デバイスの ID | 
 **tag** | **[]string** | 検索したいタグの値の文字列 | 
 **imsi** | **[]string** | 検索したい SORACOM Inventory デバイスの bootstrap 時の IMSI | 
 **imei** | **[]string** | 検索したい SORACOM Inventory デバイスの bootstrap 時の IMEI | 
 **limit** | **int32** | 取得する結果の上限数 | [default to 10]
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の SORACOM Inventory デバイスの ID。このパラメータを指定することで次の SORACOM Inventory デバイス以降のリストを取得できる。 | 
 **searchType** | **string** | 検索の種別 (AND 検索もしくは OR 検索) | [default to &quot;and&quot;]

### Return type

[**[]Device**](Device.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSigfoxDevices

> []SigfoxDevice SearchSigfoxDevices(ctx).Name(name).Group(group).DeviceId(deviceId).Tag(tag).Status(status).Registration(registration).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()

Sigfox デバイスをクエリに応じて検索します。



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
    name := []string{"Inner_example"} // []string | 検索したい Sigfox デバイスの名前 (optional)
    group := []string{"Inner_example"} // []string | 検索したいグループの名前 (optional)
    deviceId := []string{"Inner_example"} // []string | 検索したい Sigfox デバイスの ID (optional)
    tag := []string{"Inner_example"} // []string | 検索したいタグの値の文字列 (optional)
    status := "status_example" // string | 検索したい Sigfox デバイスの状態 (optional) (default to "and")
    registration := "registration_example" // string | 検索したい Sigfox デバイスの登録状態 (optional) (default to "and")
    limit := int32(56) // int32 | 取得する結果の上限数 (optional) (default to 10)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の Sigfox デバイスの ID。このパラメータを指定することで次の Sigfox デバイス以降のリストを取得できる。 (optional)
    searchType := "searchType_example" // string | 検索の種別 (AND 検索もしくは OR 検索) (optional) (default to "and")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.SearchSigfoxDevices(context.Background()).Name(name).Group(group).DeviceId(deviceId).Tag(tag).Status(status).Registration(registration).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.SearchSigfoxDevices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSigfoxDevices`: []SigfoxDevice
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.SearchSigfoxDevices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSigfoxDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | 検索したい Sigfox デバイスの名前 | 
 **group** | **[]string** | 検索したいグループの名前 | 
 **deviceId** | **[]string** | 検索したい Sigfox デバイスの ID | 
 **tag** | **[]string** | 検索したいタグの値の文字列 | 
 **status** | **string** | 検索したい Sigfox デバイスの状態 | [default to &quot;and&quot;]
 **registration** | **string** | 検索したい Sigfox デバイスの登録状態 | [default to &quot;and&quot;]
 **limit** | **int32** | 取得する結果の上限数 | [default to 10]
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の Sigfox デバイスの ID。このパラメータを指定することで次の Sigfox デバイス以降のリストを取得できる。 | 
 **searchType** | **string** | 検索の種別 (AND 検索もしくは OR 検索) | [default to &quot;and&quot;]

### Return type

[**[]SigfoxDevice**](SigfoxDevice.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSims

> []Sim SearchSims(ctx).Name(name).Group(group).SimId(simId).Imsi(imsi).Msisdn(msisdn).Iccid(iccid).SerialNumber(serialNumber).Tag(tag).Bundles(bundles).SessionStatus(sessionStatus).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()

SIM をクエリに応じて検索します。



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
    name := []string{"Inner_example"} // []string | 検索したい SIM の名前 (optional)
    group := []string{"Inner_example"} // []string | 検索したいグループの名前 (optional)
    simId := []string{"Inner_example"} // []string | 検索したい SIM ID (optional)
    imsi := []string{"Inner_example"} // []string | 検索したい IMSI (optional)
    msisdn := []string{"Inner_example"} // []string | 検索したい MSISDN (optional)
    iccid := []string{"Inner_example"} // []string | 検索したい ICCID (optional)
    serialNumber := []string{"Inner_example"} // []string | 検索したい製造番号 (optional)
    tag := []string{"Inner_example"} // []string | 検索したいタグの値の文字列 (optional)
    bundles := []string{"Inner_example"} // []string | 検索したいバンドルタイプ (optional)
    sessionStatus := "sessionStatus_example" // string | 検索したいセッションステータス (ONLINE もしくは OFFLINE) (optional) (default to "NA")
    limit := int32(56) // int32 | 取得する結果の上限数 (optional) (default to 10)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の SIM の SIM ID。このパラメータを指定することで次の SIM 以降のリストを取得できる。 (optional)
    searchType := "searchType_example" // string | 検索の種別 (AND 検索もしくは OR 検索) (optional) (default to "and")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.SearchSims(context.Background()).Name(name).Group(group).SimId(simId).Imsi(imsi).Msisdn(msisdn).Iccid(iccid).SerialNumber(serialNumber).Tag(tag).Bundles(bundles).SessionStatus(sessionStatus).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.SearchSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSims`: []Sim
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.SearchSims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | 検索したい SIM の名前 | 
 **group** | **[]string** | 検索したいグループの名前 | 
 **simId** | **[]string** | 検索したい SIM ID | 
 **imsi** | **[]string** | 検索したい IMSI | 
 **msisdn** | **[]string** | 検索したい MSISDN | 
 **iccid** | **[]string** | 検索したい ICCID | 
 **serialNumber** | **[]string** | 検索したい製造番号 | 
 **tag** | **[]string** | 検索したいタグの値の文字列 | 
 **bundles** | **[]string** | 検索したいバンドルタイプ | 
 **sessionStatus** | **string** | 検索したいセッションステータス (ONLINE もしくは OFFLINE) | [default to &quot;NA&quot;]
 **limit** | **int32** | 取得する結果の上限数 | [default to 10]
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の SIM の SIM ID。このパラメータを指定することで次の SIM 以降のリストを取得できる。 | 
 **searchType** | **string** | 検索の種別 (AND 検索もしくは OR 検索) | [default to &quot;and&quot;]

### Return type

[**[]Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSubscriberTrafficVolumeRanking

> []TrafficVolumeRanking SearchSubscriberTrafficVolumeRanking(ctx).From(from).To(to).Limit(limit).Order(order).Execute()

Subscriber の通信量ランキングを返却します。



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
    from := int64(789) // int64 | 検索期間の始点 (UNIX 時間 (ミリ秒))
    to := int64(789) // int64 | 検索期間の終点 (UNIX 時間 (ミリ秒))
    limit := int32(56) // int32 | 取得する結果の上限数 (optional) (default to 10)
    order := "order_example" // string | ランキングの順序 (optional) (default to "desc")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.SearchSubscriberTrafficVolumeRanking(context.Background()).From(from).To(to).Limit(limit).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.SearchSubscriberTrafficVolumeRanking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSubscriberTrafficVolumeRanking`: []TrafficVolumeRanking
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.SearchSubscriberTrafficVolumeRanking`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSubscriberTrafficVolumeRankingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int64** | 検索期間の始点 (UNIX 時間 (ミリ秒)) | 
 **to** | **int64** | 検索期間の終点 (UNIX 時間 (ミリ秒)) | 
 **limit** | **int32** | 取得する結果の上限数 | [default to 10]
 **order** | **string** | ランキングの順序 | [default to &quot;desc&quot;]

### Return type

[**[]TrafficVolumeRanking**](TrafficVolumeRanking.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSubscribers

> []Subscriber SearchSubscribers(ctx).Name(name).Group(group).Imsi(imsi).Msisdn(msisdn).Iccid(iccid).SerialNumber(serialNumber).Tag(tag).Subscription(subscription).ModuleType(moduleType).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()

(非推奨) Subscriber をクエリに応じて検索します。



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
    name := []string{"Inner_example"} // []string | 検索したい Subscriber の名前 (optional)
    group := []string{"Inner_example"} // []string | 検索したいグループの名前 (optional)
    imsi := []string{"Inner_example"} // []string | 検索したい IMSI (optional)
    msisdn := []string{"Inner_example"} // []string | 検索したい MSISDN (optional)
    iccid := []string{"Inner_example"} // []string | 検索したい ICCID (optional)
    serialNumber := []string{"Inner_example"} // []string | 検索したい製造番号 (optional)
    tag := []string{"Inner_example"} // []string | 検索したいタグの値の文字列 (optional)
    subscription := []string{"Inner_example"} // []string | 検索したいサブスクリプション (例: `plan01s`) (optional)
    moduleType := []string{"Inner_example"} // []string | 検索したいモジュールタイプ (例: `mini`, `virtual`) (optional)
    limit := int32(56) // int32 | 取得する結果の上限数 (optional) (default to 10)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の Subscriber の IMSI。このパラメータを指定することで次の Subscriber 以降のリストを取得できる。 (optional)
    searchType := "searchType_example" // string | 検索の種別 (AND 検索もしくは OR 検索) (optional) (default to "and")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.QueryApi.SearchSubscribers(context.Background()).Name(name).Group(group).Imsi(imsi).Msisdn(msisdn).Iccid(iccid).SerialNumber(serialNumber).Tag(tag).Subscription(subscription).ModuleType(moduleType).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).SearchType(searchType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QueryApi.SearchSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchSubscribers`: []Subscriber
    fmt.Fprintf(os.Stdout, "Response from `QueryApi.SearchSubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | 検索したい Subscriber の名前 | 
 **group** | **[]string** | 検索したいグループの名前 | 
 **imsi** | **[]string** | 検索したい IMSI | 
 **msisdn** | **[]string** | 検索したい MSISDN | 
 **iccid** | **[]string** | 検索したい ICCID | 
 **serialNumber** | **[]string** | 検索したい製造番号 | 
 **tag** | **[]string** | 検索したいタグの値の文字列 | 
 **subscription** | **[]string** | 検索したいサブスクリプション (例: &#x60;plan01s&#x60;) | 
 **moduleType** | **[]string** | 検索したいモジュールタイプ (例: &#x60;mini&#x60;, &#x60;virtual&#x60;) | 
 **limit** | **int32** | 取得する結果の上限数 | [default to 10]
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の Subscriber の IMSI。このパラメータを指定することで次の Subscriber 以降のリストを取得できる。 | 
 **searchType** | **string** | 検索の種別 (AND 検索もしくは OR 検索) | [default to &quot;and&quot;]

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=UTF-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

