# \CellLocationApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchGetCellLocations**](CellLocationApi.md#BatchGetCellLocations) | **Post** /cell_locations | 複数の基地局の位置情報を一度に取得します。
[**GetCellLocation**](CellLocationApi.md#GetCellLocation) | **Get** /cell_locations | 基地局の位置情報を取得します。



## BatchGetCellLocations

> []CellLocation BatchGetCellLocations(ctx).CellIdentifier(cellIdentifier).Execute()

複数の基地局の位置情報を一度に取得します。



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
    cellIdentifier := []openapiclient.CellIdentifier{*openapiclient.NewCellIdentifier()} // []CellIdentifier | 基地局情報のリスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellLocationApi.BatchGetCellLocations(context.Background()).CellIdentifier(cellIdentifier).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellLocationApi.BatchGetCellLocations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BatchGetCellLocations`: []CellLocation
    fmt.Fprintf(os.Stdout, "Response from `CellLocationApi.BatchGetCellLocations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBatchGetCellLocationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cellIdentifier** | [**[]CellIdentifier**](CellIdentifier.md) | 基地局情報のリスト | 

### Return type

[**[]CellLocation**](CellLocation.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCellLocation

> CellLocation GetCellLocation(ctx).Mcc(mcc).Mnc(mnc).Lac(lac).Cid(cid).Tac(tac).Ecid(ecid).Eci(eci).Execute()

基地局の位置情報を取得します。



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
    mcc := "mcc_example" // string | MCC (Mobile Country Code) - 国コード
    mnc := "mnc_example" // string | MNC (Mobile Network Code) - ネットワークコード
    lac := "lac_example" // string | LAC (Location Area Code) - エリアコード（3G 用） (optional)
    cid := "cid_example" // string | CID (Cell ID) - セル ID（3G 用） (optional)
    tac := "tac_example" // string | TAC (Tracking Area Code) - エリアコード（LTE 用） (optional)
    ecid := "ecid_example" // string | ECID (Enhanced Cell ID) - セル ID（LTE 用） - `ecid` もしくは `eci` のいずれかを指定してください。（どちらに値を指定しても結果は同一です） (optional)
    eci := "eci_example" // string | ECID (Enhanced Cell ID) - セル ID（LTE 用） - `ecid` もしくは `eci` のいずれかを指定してください。（どちらに値を指定しても結果は同一です） (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CellLocationApi.GetCellLocation(context.Background()).Mcc(mcc).Mnc(mnc).Lac(lac).Cid(cid).Tac(tac).Ecid(ecid).Eci(eci).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CellLocationApi.GetCellLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCellLocation`: CellLocation
    fmt.Fprintf(os.Stdout, "Response from `CellLocationApi.GetCellLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCellLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mcc** | **string** | MCC (Mobile Country Code) - 国コード | 
 **mnc** | **string** | MNC (Mobile Network Code) - ネットワークコード | 
 **lac** | **string** | LAC (Location Area Code) - エリアコード（3G 用） | 
 **cid** | **string** | CID (Cell ID) - セル ID（3G 用） | 
 **tac** | **string** | TAC (Tracking Area Code) - エリアコード（LTE 用） | 
 **ecid** | **string** | ECID (Enhanced Cell ID) - セル ID（LTE 用） - &#x60;ecid&#x60; もしくは &#x60;eci&#x60; のいずれかを指定してください。（どちらに値を指定しても結果は同一です） | 
 **eci** | **string** | ECID (Enhanced Cell ID) - セル ID（LTE 用） - &#x60;ecid&#x60; もしくは &#x60;eci&#x60; のいずれかを指定してください。（どちらに値を指定しても結果は同一です） | 

### Return type

[**CellLocation**](CellLocation.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

