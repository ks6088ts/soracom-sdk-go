# \SubscriberApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSubscriber**](SubscriberApi.md#ActivateSubscriber) | **Post** /subscribers/{imsi}/activate | Activate Subscriber
[**DeactivateSubscriber**](SubscriberApi.md#DeactivateSubscriber) | **Post** /subscribers/{imsi}/deactivate | Deactivate Subscriber
[**DeleteSubscriberSession**](SubscriberApi.md#DeleteSubscriberSession) | **Post** /subscribers/{imsi}/delete_session | Delete Session
[**DeleteSubscriberTag**](SubscriberApi.md#DeleteSubscriberTag) | **Delete** /subscribers/{imsi}/tags/{tag_name} | Delete Subscriber Tag
[**DeleteSubscriberTransferToken**](SubscriberApi.md#DeleteSubscriberTransferToken) | **Delete** /subscribers/transfer_token/{token} | Delete Subscribers Transfer Token
[**DisableTermination**](SubscriberApi.md#DisableTermination) | **Post** /subscribers/{imsi}/disable_termination | Disable Termination of Subscriber
[**EnableTermination**](SubscriberApi.md#EnableTermination) | **Post** /subscribers/{imsi}/enable_termination | Enable Termination of Subscriber
[**ExportSubscribers**](SubscriberApi.md#ExportSubscribers) | **Post** /subscribers/export | Export all subscribers.
[**GetDataFromSubscriber**](SubscriberApi.md#GetDataFromSubscriber) | **Get** /subscribers/{imsi}/data | Get data sent from a subscriber.
[**GetSubscriber**](SubscriberApi.md#GetSubscriber) | **Get** /subscribers/{imsi} | Get Subscriber
[**IssueSubscriberTransferToken**](SubscriberApi.md#IssueSubscriberTransferToken) | **Post** /subscribers/transfer_token/issue | Issue Subscribers Transfer Token
[**ListSessionEvents**](SubscriberApi.md#ListSessionEvents) | **Get** /subscribers/{imsi}/events/sessions | List Session Events
[**ListSubscribers**](SubscriberApi.md#ListSubscribers) | **Get** /subscribers | List Subscribers
[**PutBundles**](SubscriberApi.md#PutBundles) | **Put** /subscribers/{imsi}/bundles | Set Bundles to Subscriber.
[**PutSubscriberTags**](SubscriberApi.md#PutSubscriberTags) | **Put** /subscribers/{imsi}/tags | Bulk Insert or Update Subscriber Tags
[**RegisterSubscriber**](SubscriberApi.md#RegisterSubscriber) | **Post** /subscribers/{imsi}/register | Register Subscriber
[**ReportLocalInfo**](SubscriberApi.md#ReportLocalInfo) | **Post** /subscribers/{imsi}/report_local_info | Triggers Subscriber to report SIM local info.
[**SendSms**](SubscriberApi.md#SendSms) | **Post** /subscribers/{imsi}/send_sms | Send SMS to Subscriber
[**SendSmsByMsisdn**](SubscriberApi.md#SendSmsByMsisdn) | **Post** /subscribers/msisdn/{msisdn}/send_sms | Send SMS to Subscriber by MSISDN
[**SetExpiryTime**](SubscriberApi.md#SetExpiryTime) | **Post** /subscribers/{imsi}/set_expiry_time | Update Expiry Time of Subscriber
[**SetGroup**](SubscriberApi.md#SetGroup) | **Post** /subscribers/{imsi}/set_group | Set Group to Subscriber
[**SetImeiLock**](SubscriberApi.md#SetImeiLock) | **Post** /subscribers/{imsi}/set_imei_lock | Set IMEI lock configuration for Subscriber.
[**SetSubscriberToStandby**](SubscriberApi.md#SetSubscriberToStandby) | **Post** /subscribers/{imsi}/set_to_standby | Set Subscriber to standby mode.
[**SubscriberDownlinkPingToUserEquipment**](SubscriberApi.md#SubscriberDownlinkPingToUserEquipment) | **Post** /subscribers/{imsi}/downlink/ping | Subscriber ???????????? ping ?????????????????????????????????
[**SuspendSubscriber**](SubscriberApi.md#SuspendSubscriber) | **Post** /subscribers/{imsi}/suspend | Suspend Subscriber
[**TerminateSubscriber**](SubscriberApi.md#TerminateSubscriber) | **Post** /subscribers/{imsi}/terminate | Terminate Subscriber
[**UnsetExpiryTime**](SubscriberApi.md#UnsetExpiryTime) | **Post** /subscribers/{imsi}/unset_expiry_time | Delete Expiry Time of Subscriber
[**UnsetGroup**](SubscriberApi.md#UnsetGroup) | **Post** /subscribers/{imsi}/unset_group | Unset Group to Subscriber
[**UnsetImeiLock**](SubscriberApi.md#UnsetImeiLock) | **Post** /subscribers/{imsi}/unset_imei_lock | Unset IMEI lock configuration for Subscriber.
[**UpdateSpeedClass**](SubscriberApi.md#UpdateSpeedClass) | **Post** /subscribers/{imsi}/update_speed_class | Update Subscriber speed class
[**VerifySubscriberTransferToken**](SubscriberApi.md#VerifySubscriberTransferToken) | **Post** /subscribers/transfer_token/verify | Verify Subscriber Transfer Token



## ActivateSubscriber

> Subscriber ActivateSubscriber(ctx, imsi).Execute()

Activate Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.ActivateSubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.ActivateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateSubscriber`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.ActivateSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSubscriber

> Subscriber DeactivateSubscriber(ctx, imsi).Execute()

Deactivate Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.DeactivateSubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.DeactivateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateSubscriber`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.DeactivateSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriberSession

> Subscriber DeleteSubscriberSession(ctx, imsi).Execute()

Delete Session



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.DeleteSubscriberSession(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.DeleteSubscriberSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriberSession`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.DeleteSubscriberSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubscriberTag

> DeleteSubscriberTag(ctx, imsi, tagName).Execute()

Delete Subscriber Tag



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    tagName := "tagName_example" // string | ???????????????????????????URL ??? Path ??????????????????????????????????????????????????????????????????????????????JavaScript ?????? encodeURIComponent() ??????????????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.DeleteSubscriberTag(context.Background(), imsi, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.DeleteSubscriberTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 
**tagName** | **string** | ???????????????????????????URL ??? Path ??????????????????????????????????????????????????????????????????????????????JavaScript ?????? encodeURIComponent() ?????????????????????????????? | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberTagRequest struct via the builder pattern


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


## DeleteSubscriberTransferToken

> DeleteSubscriberTransferToken(ctx, token).Execute()

Delete Subscribers Transfer Token



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
    token := "token_example" // string | token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.DeleteSubscriberTransferToken(context.Background(), token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.DeleteSubscriberTransferToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | token | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriberTransferTokenRequest struct via the builder pattern


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


## DisableTermination

> Subscriber DisableTermination(ctx, imsi).Execute()

Disable Termination of Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.DisableTermination(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.DisableTermination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableTermination`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.DisableTermination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableTerminationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableTermination

> Subscriber EnableTermination(ctx, imsi).Execute()

Enable Termination of Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.EnableTermination(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.EnableTermination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableTermination`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.EnableTermination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableTerminationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExportSubscribers

> FileExportResponse ExportSubscribers(ctx).ExportMode(exportMode).Execute()

Export all subscribers.



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
    exportMode := "exportMode_example" // string | ???????????????????????? (?????????, ??????) (optional) (default to "sync")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.ExportSubscribers(context.Background()).ExportMode(exportMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.ExportSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExportSubscribers`: FileExportResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.ExportSubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **exportMode** | **string** | ???????????????????????? (?????????, ??????) | [default to &quot;sync&quot;]

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


## GetDataFromSubscriber

> []DataEntry GetDataFromSubscriber(ctx, imsi).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get data sent from a subscriber.



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
    imsi := "imsi_example" // string | ??????????????????????????????????????????????????? IMSI
    from := int32(56) // int32 | ????????????????????????????????? (UNIX ?????? (?????????)) (optional)
    to := int32(56) // int32 | ????????????????????????????????? (UNIX ?????? (?????????)) (optional)
    sort := "sort_example" // string | ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? (optional) (default to "desc")
    limit := int32(56) // int32 | ????????????????????????????????????????????? (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | ??????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.GetDataFromSubscriber(context.Background(), imsi).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.GetDataFromSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataFromSubscriber`: []DataEntry
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.GetDataFromSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ??????????????????????????????????????????????????? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFromSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | ????????????????????????????????? (UNIX ?????? (?????????)) | 
 **to** | **int32** | ????????????????????????????????? (UNIX ?????? (?????????)) | 
 **sort** | **string** | ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? | [default to &quot;desc&quot;]
 **limit** | **int32** | ????????????????????????????????????????????? | 
 **lastEvaluatedKey** | **string** | ??????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? | 

### Return type

[**[]DataEntry**](DataEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSubscriber

> Subscriber GetSubscriber(ctx, imsi).Execute()

Get Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.GetSubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.GetSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSubscriber`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.GetSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueSubscriberTransferToken

> IssueSubscriberTransferTokenResponse IssueSubscriberTransferToken(ctx).IssueSubscriberTransferTokenRequest(issueSubscriberTransferTokenRequest).Execute()

Issue Subscribers Transfer Token



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
    issueSubscriberTransferTokenRequest := *openapiclient.NewIssueSubscriberTransferTokenRequest("TransferDestinationOperatorEmail_example", "TransferDestinationOperatorId_example", []string{"TransferImsi_example"}) // IssueSubscriberTransferTokenRequest | ????????????????????????????????????????????????????????????????????????ID?????????????????? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.IssueSubscriberTransferToken(context.Background()).IssueSubscriberTransferTokenRequest(issueSubscriberTransferTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.IssueSubscriberTransferToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IssueSubscriberTransferToken`: IssueSubscriberTransferTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.IssueSubscriberTransferToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueSubscriberTransferTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueSubscriberTransferTokenRequest** | [**IssueSubscriberTransferTokenRequest**](IssueSubscriberTransferTokenRequest.md) | ????????????????????????????????????????????????????????????????????????ID?????????????????? IMSI | 

### Return type

[**IssueSubscriberTransferTokenResponse**](IssueSubscriberTransferTokenResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionEvents

> []SessionEvent ListSessionEvents(ctx, imsi).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List Session Events



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    from := int32(56) // int32 | ?????????????????????????????????????????????(unixtime) (optional)
    to := int32(56) // int32 | ?????????????????????????????????????????????(unixtime) (optional)
    limit := int32(56) // int32 | ???????????????????????????????????? (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.ListSessionEvents(context.Background(), imsi).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.ListSessionEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionEvents`: []SessionEvent
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.ListSessionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSessionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | ?????????????????????????????????????????????(unixtime) | 
 **to** | **int32** | ?????????????????????????????????????????????(unixtime) | 
 **limit** | **int32** | ???????????????????????????????????? | 
 **lastEvaluatedKey** | **string** | ????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????? | 

### Return type

[**[]SessionEvent**](SessionEvent.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscribers

> []Subscriber ListSubscribers(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).StatusFilter(statusFilter).SpeedClassFilter(speedClassFilter).SerialNumberFilter(serialNumberFilter).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List Subscribers



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
    tagName := "tagName_example" // string | ????????????????????????????????????(????????????)??? (optional)
    tagValue := "tagValue_example" // string | ????????????????????????????????????????????????`tag_name` ????????????????????????????????? (optional)
    tagValueMatchMode := "tagValueMatchMode_example" // string | ???????????????????????? (optional) (default to "exact")
    statusFilter := "statusFilter_example" // string | ????????????????????? status???`|`??????????????????????????????????????????????????????????????????????????????????????????????????????: `active`, `inactive`, `ready`, `instock`, `shipped`, `suspended`, `terminated` (optional)
    speedClassFilter := "speedClassFilter_example" // string | ???????????????????????????????????????`|`??????????????????????????????????????????????????????????????????????????????????????????????????????: `s1.minimum`, `s1.slow`, `s1.standard`, `s1.fast` (optional)
    serialNumberFilter := "serialNumberFilter_example" // string | ????????????????????????????????????`|`????????????????????????????????????????????????????????????????????????????????????????????????????????? Subscriber ????????????????????? (optional)
    limit := int32(56) // int32 | ???????????? Subscriber ??????????????????????????????????????? Subscriber ????????????????????????????????????????????????????????????????????? (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | ???????????????????????????????????? Subscriber ??? IMSI?????????????????????????????????????????????????????? Subscriber ??????????????????????????????????????? (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.ListSubscribers(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).StatusFilter(statusFilter).SpeedClassFilter(speedClassFilter).SerialNumberFilter(serialNumberFilter).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.ListSubscribers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscribers`: []Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.ListSubscribers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubscribersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | ????????????????????????????????????(????????????)??? | 
 **tagValue** | **string** | ????????????????????????????????????????????????&#x60;tag_name&#x60; ????????????????????????????????? | 
 **tagValueMatchMode** | **string** | ???????????????????????? | [default to &quot;exact&quot;]
 **statusFilter** | **string** | ????????????????????? status???&#x60;|&#x60;??????????????????????????????????????????????????????????????????????????????????????????????????????: &#x60;active&#x60;, &#x60;inactive&#x60;, &#x60;ready&#x60;, &#x60;instock&#x60;, &#x60;shipped&#x60;, &#x60;suspended&#x60;, &#x60;terminated&#x60; | 
 **speedClassFilter** | **string** | ???????????????????????????????????????&#x60;|&#x60;??????????????????????????????????????????????????????????????????????????????????????????????????????: &#x60;s1.minimum&#x60;, &#x60;s1.slow&#x60;, &#x60;s1.standard&#x60;, &#x60;s1.fast&#x60; | 
 **serialNumberFilter** | **string** | ????????????????????????????????????&#x60;|&#x60;????????????????????????????????????????????????????????????????????????????????????????????????????????? Subscriber ????????????????????? | 
 **limit** | **int32** | ???????????? Subscriber ??????????????????????????????????????? Subscriber ????????????????????????????????????????????????????????????????????? | 
 **lastEvaluatedKey** | **string** | ???????????????????????????????????? Subscriber ??? IMSI?????????????????????????????????????????????????????? Subscriber ??????????????????????????????????????? | 

### Return type

[**[]Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBundles

> Subscriber PutBundles(ctx, imsi).RequestBody(requestBody).Execute()

Set Bundles to Subscriber.



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    requestBody := []string{"Property_example"} // []string | ?????????????????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.PutBundles(context.Background(), imsi).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.PutBundles``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutBundles`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.PutBundles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | **[]string** | ????????????????????????????????? | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSubscriberTags

> Subscriber PutSubscriberTags(ctx, imsi).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update Subscriber Tags



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | ???????????????????????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.PutSubscriberTags(context.Background(), imsi).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.PutSubscriberTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSubscriberTags`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.PutSubscriberTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSubscriberTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | ??????????????????????????????????????? | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSubscriber

> Subscriber RegisterSubscriber(ctx, imsi).RegisterSubscribersRequest(registerSubscribersRequest).Execute()

Register Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    registerSubscribersRequest := *openapiclient.NewRegisterSubscribersRequest("RegistrationSecret_example") // RegisterSubscribersRequest | subscriber

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.RegisterSubscriber(context.Background(), imsi).RegisterSubscribersRequest(registerSubscribersRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.RegisterSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterSubscriber`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.RegisterSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerSubscribersRequest** | [**RegisterSubscribersRequest**](RegisterSubscribersRequest.md) | subscriber | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportLocalInfo

> Subscriber ReportLocalInfo(ctx, imsi).Execute()

Triggers Subscriber to report SIM local info.



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
    imsi := "imsi_example" // string | ???????????????????????????????????? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.ReportLocalInfo(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.ReportLocalInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportLocalInfo`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.ReportLocalInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ???????????????????????????????????? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportLocalInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSms

> SmsForwardingReport SendSms(ctx, imsi).SmsForwardingRequest(smsForwardingRequest).Execute()

Send SMS to Subscriber



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
    imsi := "imsi_example" // string | SMS ?????????????????????????????????????????? IMSI
    smsForwardingRequest := *openapiclient.NewSmsForwardingRequest() // SmsForwardingRequest | ????????????????????????????????????????????????????????????????????? SMS ?????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SendSms(context.Background(), imsi).SmsForwardingRequest(smsForwardingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SendSms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSms`: SmsForwardingReport
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SendSms`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | SMS ?????????????????????????????????????????? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsForwardingRequest** | [**SmsForwardingRequest**](SmsForwardingRequest.md) | ????????????????????????????????????????????????????????????????????? SMS ????????????????????? | 

### Return type

[**SmsForwardingReport**](SmsForwardingReport.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendSmsByMsisdn

> SmsForwardingReport SendSmsByMsisdn(ctx, msisdn).SmsForwardingRequest(smsForwardingRequest).Execute()

Send SMS to Subscriber by MSISDN



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
    msisdn := "msisdn_example" // string | SMS ?????????????????????????????????????????? MSISDN
    smsForwardingRequest := *openapiclient.NewSmsForwardingRequest() // SmsForwardingRequest | ????????????????????????????????????????????????????????????????????? SMS ?????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SendSmsByMsisdn(context.Background(), msisdn).SmsForwardingRequest(smsForwardingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SendSmsByMsisdn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSmsByMsisdn`: SmsForwardingReport
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SendSmsByMsisdn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**msisdn** | **string** | SMS ?????????????????????????????????????????? MSISDN | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsByMsisdnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsForwardingRequest** | [**SmsForwardingRequest**](SmsForwardingRequest.md) | ????????????????????????????????????????????????????????????????????? SMS ????????????????????? | 

### Return type

[**SmsForwardingReport**](SmsForwardingReport.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetExpiryTime

> Subscriber SetExpiryTime(ctx, imsi).ExpiryTime(expiryTime).Execute()

Update Expiry Time of Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    expiryTime := *openapiclient.NewExpiryTime(int64(123)) // ExpiryTime | ???????????????????????? (unixtime:???????????????) ??????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SetExpiryTime(context.Background(), imsi).ExpiryTime(expiryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SetExpiryTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetExpiryTime`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SetExpiryTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetExpiryTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiryTime** | [**ExpiryTime**](ExpiryTime.md) | ???????????????????????? (unixtime:???????????????) ?????????????????? | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGroup

> Subscriber SetGroup(ctx, imsi).SetGroupRequest(setGroupRequest).Execute()

Set Group to Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    setGroupRequest := *openapiclient.NewSetGroupRequest() // SetGroupRequest | Group???ID ???????????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SetGroup(context.Background(), imsi).SetGroupRequest(setGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetGroup`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setGroupRequest** | [**SetGroupRequest**](SetGroupRequest.md) | Group???ID ??????????????????????????? | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetImeiLock

> Subscriber SetImeiLock(ctx, imsi).SetImeiLockRequest(setImeiLockRequest).Execute()

Set IMEI lock configuration for Subscriber.



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    setImeiLockRequest := *openapiclient.NewSetImeiLockRequest() // SetImeiLockRequest | Subscriber ??????????????? IMEI lock ????????? ????????????????????? Subscriber ???????????? IMEI ???????????????????????? IMEI ????????????????????????????????? (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SetImeiLock(context.Background(), imsi).SetImeiLockRequest(setImeiLockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SetImeiLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetImeiLock`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SetImeiLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetImeiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setImeiLockRequest** | [**SetImeiLockRequest**](SetImeiLockRequest.md) | Subscriber ??????????????? IMEI lock ????????? ????????????????????? Subscriber ???????????? IMEI ???????????????????????? IMEI ????????????????????????????????? | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSubscriberToStandby

> Subscriber SetSubscriberToStandby(ctx, imsi).Execute()

Set Subscriber to standby mode.



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SetSubscriberToStandby(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SetSubscriberToStandby``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSubscriberToStandby`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SetSubscriberToStandby`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSubscriberToStandbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriberDownlinkPingToUserEquipment

> DownlinkPingResponse SubscriberDownlinkPingToUserEquipment(ctx, imsi).DownlinkPingRequest(downlinkPingRequest).Execute()

Subscriber ???????????? ping ?????????????????????????????????



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    downlinkPingRequest := *openapiclient.NewDownlinkPingRequest() // DownlinkPingRequest | ping ?????????????????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SubscriberDownlinkPingToUserEquipment(context.Background(), imsi).DownlinkPingRequest(downlinkPingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SubscriberDownlinkPingToUserEquipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubscriberDownlinkPingToUserEquipment`: DownlinkPingResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SubscriberDownlinkPingToUserEquipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriberDownlinkPingToUserEquipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downlinkPingRequest** | [**DownlinkPingRequest**](DownlinkPingRequest.md) | ping ????????????????????????????????? | 

### Return type

[**DownlinkPingResponse**](DownlinkPingResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendSubscriber

> Subscriber SuspendSubscriber(ctx, imsi).Execute()

Suspend Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.SuspendSubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.SuspendSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendSubscriber`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.SuspendSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateSubscriber

> Subscriber TerminateSubscriber(ctx, imsi).Execute()

Terminate Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.TerminateSubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.TerminateSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateSubscriber`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.TerminateSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetExpiryTime

> UnsetExpiryTime(ctx, imsi).Execute()

Delete Expiry Time of Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.UnsetExpiryTime(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.UnsetExpiryTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetExpiryTimeRequest struct via the builder pattern


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


## UnsetGroup

> Subscriber UnsetGroup(ctx, imsi).Execute()

Unset Group to Subscriber



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.UnsetGroup(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.UnsetGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetGroup`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.UnsetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetImeiLock

> Subscriber UnsetImeiLock(ctx, imsi).Execute()

Unset IMEI lock configuration for Subscriber.



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.UnsetImeiLock(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.UnsetImeiLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetImeiLock`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.UnsetImeiLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetImeiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSpeedClass

> Subscriber UpdateSpeedClass(ctx, imsi).UpdateSpeedClassRequest(updateSpeedClassRequest).Execute()

Update Subscriber speed class



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
    imsi := "imsi_example" // string | ????????? Subscriber ??? IMSI
    updateSpeedClassRequest := *openapiclient.NewUpdateSpeedClassRequest("SpeedClass_example") // UpdateSpeedClassRequest | speed_class

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.UpdateSpeedClass(context.Background(), imsi).UpdateSpeedClassRequest(updateSpeedClassRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.UpdateSpeedClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSpeedClass`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.UpdateSpeedClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | ????????? Subscriber ??? IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSpeedClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSpeedClassRequest** | [**UpdateSpeedClassRequest**](UpdateSpeedClassRequest.md) | speed_class | 

### Return type

[**Subscriber**](Subscriber.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySubscriberTransferToken

> VerifySubscriberTransferTokenResponse VerifySubscriberTransferToken(ctx).VerifySubscriberTransferTokenRequest(verifySubscriberTransferTokenRequest).Execute()

Verify Subscriber Transfer Token



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
    verifySubscriberTransferTokenRequest := *openapiclient.NewVerifySubscriberTransferTokenRequest("Token_example") // VerifySubscriberTransferTokenRequest | ???????????????????????????????????????

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriberApi.VerifySubscriberTransferToken(context.Background()).VerifySubscriberTransferTokenRequest(verifySubscriberTransferTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriberApi.VerifySubscriberTransferToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifySubscriberTransferToken`: VerifySubscriberTransferTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriberApi.VerifySubscriberTransferToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifySubscriberTransferTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifySubscriberTransferTokenRequest** | [**VerifySubscriberTransferTokenRequest**](VerifySubscriberTransferTokenRequest.md) | ??????????????????????????????????????? | 

### Return type

[**VerifySubscriberTransferTokenResponse**](VerifySubscriberTransferTokenResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

