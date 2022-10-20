# \SimApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSim**](SimApi.md#ActivateSim) | **Post** /sims/{sim_id}/activate | Activate SIM
[**AddSubscription**](SimApi.md#AddSubscription) | **Post** /sims/{sim_id}/profiles/{iccid}/add_subscription | 新しいセカンダリサブスクリプションを IoT SIM に追加する
[**AttachArcSimCredentials**](SimApi.md#AttachArcSimCredentials) | **Post** /sims/{sim_id}/credentials/arc | (非推奨) Arc のクレデンシャルを SIM へ付与
[**CancelSubscriptionContainerDownload**](SimApi.md#CancelSubscriptionContainerDownload) | **Post** /sims/{sim_id}/profiles/{iccid}/subscribers/{imsi}/cancel_download | Cancel subscription download to subscription container
[**CreateArcSession**](SimApi.md#CreateArcSession) | **Post** /sims/{sim_id}/sessions/arc | SORACOM Arc セッションの作成
[**CreateSim**](SimApi.md#CreateSim) | **Post** /sims | SIM の新規作成
[**CreateSimPacketCaptureSession**](SimApi.md#CreateSimPacketCaptureSession) | **Post** /sims/{sim_id}/packet_capture_sessions | Create Packet Capture Session
[**DeactivateSim**](SimApi.md#DeactivateSim) | **Post** /sims/{sim_id}/deactivate | Deactivate SIM
[**DeleteSimPacketCaptureSession**](SimApi.md#DeleteSimPacketCaptureSession) | **Delete** /sims/{sim_id}/packet_capture_sessions/{session_id} | Delete Packet Capture Session
[**DeleteSimSession**](SimApi.md#DeleteSimSession) | **Post** /sims/{sim_id}/delete_session | Delete Session
[**DeleteSimTag**](SimApi.md#DeleteSimTag) | **Delete** /sims/{sim_id}/tags/{tag_name} | Delete SIM Tag
[**DeleteSubscriptionContainerCountryMappingEntry**](SimApi.md#DeleteSubscriptionContainerCountryMappingEntry) | **Delete** /sims/{sim_id}/profiles/{iccid}/subscription_containers/country_mapping/{mcc} | Delete subscription container mapping entries
[**DisableSimTermination**](SimApi.md#DisableSimTermination) | **Post** /sims/{sim_id}/disable_termination | Disable Termination of SIM
[**EnableSimTermination**](SimApi.md#EnableSimTermination) | **Post** /sims/{sim_id}/enable_termination | Enable Termination of SIM
[**EnableSubscriptionContainer**](SimApi.md#EnableSubscriptionContainer) | **Post** /sims/{sim_id}/profiles/{iccid}/subscription_containers/{container_id}/enable | Enables a subscription container.
[**GetDataFromSim**](SimApi.md#GetDataFromSim) | **Get** /sims/{sim_id}/data | Get data sent from a SIM.
[**GetSim**](SimApi.md#GetSim) | **Get** /sims/{sim_id} | Get SIM
[**GetSimPacketCaptureSession**](SimApi.md#GetSimPacketCaptureSession) | **Get** /sims/{sim_id}/packet_capture_sessions/{session_id} | Get Packet Capture Session
[**ListSimPacketCaptureSessions**](SimApi.md#ListSimPacketCaptureSessions) | **Get** /sims/{sim_id}/packet_capture_sessions | List packet capture sessions associated with the SIM
[**ListSimSessionEvents**](SimApi.md#ListSimSessionEvents) | **Get** /sims/{sim_id}/events/sessions | List Session Events
[**ListSims**](SimApi.md#ListSims) | **Get** /sims | List SIMs
[**ListSubscriptionContainers**](SimApi.md#ListSubscriptionContainers) | **Get** /sims/{sim_id}/profiles/{iccid}/subscription_containers | Get subscription container status.
[**PutSimTags**](SimApi.md#PutSimTags) | **Put** /sims/{sim_id}/tags | Bulk Insert or Update SIM Tags
[**PutSubscriptionContainerCountryMappingEntries**](SimApi.md#PutSubscriptionContainerCountryMappingEntries) | **Put** /sims/{sim_id}/profiles/{iccid}/subscription_containers/country_mapping | Updates subscription container country mapping entries.
[**RegisterSim**](SimApi.md#RegisterSim) | **Post** /sims/{sim_id}/register | Register SIM
[**RemoveArcSimCredentials**](SimApi.md#RemoveArcSimCredentials) | **Delete** /sims/{sim_id}/credentials/arc | (非推奨) Arc のクレデンシャルを SIM から除去
[**RenewArcSimCredentials**](SimApi.md#RenewArcSimCredentials) | **Put** /sims/{sim_id}/credentials/arc | SIM に対する Arc のクレデンシャルを更新する
[**ReportSimLocalInfo**](SimApi.md#ReportSimLocalInfo) | **Post** /sims/{sim_id}/report_local_info | SIM ローカル情報レポートを取得する
[**SendDownlinkPing**](SimApi.md#SendDownlinkPing) | **Post** /sims/{sim_id}/downlink/ping | SIM に対して ping リクエストを送信する。
[**SendSmsToSim**](SimApi.md#SendSmsToSim) | **Post** /sims/{sim_id}/send_sms | Send SMS to SIM
[**SetSimExpiryTime**](SimApi.md#SetSimExpiryTime) | **Post** /sims/{sim_id}/set_expiry_time | Update Expiry Time of SIM
[**SetSimGroup**](SimApi.md#SetSimGroup) | **Post** /sims/{sim_id}/set_group | Set Group to SIM
[**SetSimImeiLock**](SimApi.md#SetSimImeiLock) | **Post** /sims/{sim_id}/set_imei_lock | Set IMEI lock configuration for SIM.
[**SetSimToStandby**](SimApi.md#SetSimToStandby) | **Post** /sims/{sim_id}/set_to_standby | Set SIM to standby mode.
[**StopSimPacketCaptureSession**](SimApi.md#StopSimPacketCaptureSession) | **Post** /sims/{sim_id}/packet_capture_sessions/{session_id}/stop | Stop SIM Packet Capture Session
[**SuspendSim**](SimApi.md#SuspendSim) | **Post** /sims/{sim_id}/suspend | Suspend SIM
[**TerminateSim**](SimApi.md#TerminateSim) | **Post** /sims/{sim_id}/terminate | Terminate SIM
[**TerminateSubscriptionContainer**](SimApi.md#TerminateSubscriptionContainer) | **Post** /sims/{sim_id}/profiles/{iccid}/subscribers/{imsi}/terminate | セカンダリサブスクリプションの利用終了
[**UnsetSimExpiryTime**](SimApi.md#UnsetSimExpiryTime) | **Post** /sims/{sim_id}/unset_expiry_time | Delete Expiry Time of SIM
[**UnsetSimGroup**](SimApi.md#UnsetSimGroup) | **Post** /sims/{sim_id}/unset_group | Unset Group to SIM
[**UnsetSimImeiLock**](SimApi.md#UnsetSimImeiLock) | **Post** /sims/{sim_id}/unset_imei_lock | Unset IMEI lock configuration for SIM.
[**UpdateSimSpeedClass**](SimApi.md#UpdateSimSpeedClass) | **Post** /sims/{sim_id}/update_speed_class | Update SIM speed class



## ActivateSim

> Sim ActivateSim(ctx, simId).Execute()

Activate SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.ActivateSim(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.ActivateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.ActivateSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddSubscription

> SimProfile AddSubscription(ctx, simId, iccid).AddSubscriptionRequest(addSubscriptionRequest).Execute()

新しいセカンダリサブスクリプションを IoT SIM に追加する



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
    simId := "simId_example" // string | 対象の SIM の ID
    iccid := "iccid_example" // string | 対象の IoT SIM の ICCID。eUICC 非対応 の IoT SIM の場合は、ICCID と SIM ID は同一です。
    addSubscriptionRequest := *openapiclient.NewAddSubscriptionRequest("Subscription_example") // AddSubscriptionRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.AddSubscription(context.Background(), simId, iccid).AddSubscriptionRequest(addSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.AddSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddSubscription`: SimProfile
    fmt.Fprintf(os.Stdout, "Response from `SimApi.AddSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM の ID | 
**iccid** | **string** | 対象の IoT SIM の ICCID。eUICC 非対応 の IoT SIM の場合は、ICCID と SIM ID は同一です。 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addSubscriptionRequest** | [**AddSubscriptionRequest**](AddSubscriptionRequest.md) |  | 

### Return type

[**SimProfile**](SimProfile.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachArcSimCredentials

> ArcCredentialAttachResponse AttachArcSimCredentials(ctx, simId).ArcCredentialAttachRequest(arcCredentialAttachRequest).Execute()

(非推奨) Arc のクレデンシャルを SIM へ付与



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
    simId := "simId_example" // string | 対象とする SIM の SIM ID
    arcCredentialAttachRequest := *openapiclient.NewArcCredentialAttachRequest() // ArcCredentialAttachRequest | Arc のクレデンシャルの付与リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.AttachArcSimCredentials(context.Background(), simId).ArcCredentialAttachRequest(arcCredentialAttachRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.AttachArcSimCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachArcSimCredentials`: ArcCredentialAttachResponse
    fmt.Fprintf(os.Stdout, "Response from `SimApi.AttachArcSimCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象とする SIM の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachArcSimCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arcCredentialAttachRequest** | [**ArcCredentialAttachRequest**](ArcCredentialAttachRequest.md) | Arc のクレデンシャルの付与リクエスト | 

### Return type

[**ArcCredentialAttachResponse**](ArcCredentialAttachResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelSubscriptionContainerDownload

> SubscriptionContainerStatus CancelSubscriptionContainerDownload(ctx, simId, iccid, imsi).Execute()

Cancel subscription download to subscription container



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
    simId := "simId_example" // string | SIM ID of the target subscription container.
    iccid := "iccid_example" // string | ICCID of the target subscription container.
    imsi := "imsi_example" // string | IMSI of the target subscription container.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.CancelSubscriptionContainerDownload(context.Background(), simId, iccid, imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.CancelSubscriptionContainerDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelSubscriptionContainerDownload`: SubscriptionContainerStatus
    fmt.Fprintf(os.Stdout, "Response from `SimApi.CancelSubscriptionContainerDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | SIM ID of the target subscription container. | 
**iccid** | **string** | ICCID of the target subscription container. | 
**imsi** | **string** | IMSI of the target subscription container. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelSubscriptionContainerDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SubscriptionContainerStatus**](SubscriptionContainerStatus.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateArcSession

> ArcSessionCreateResponse CreateArcSession(ctx, simId).Execute()

SORACOM Arc セッションの作成



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
    simId := "simId_example" // string | 対象とする SIM の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.CreateArcSession(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.CreateArcSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArcSession`: ArcSessionCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SimApi.CreateArcSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象とする SIM の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArcSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArcSessionCreateResponse**](ArcSessionCreateResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSim

> Sim CreateSim(ctx).CreateSimRequest(createSimRequest).Execute()

SIM の新規作成



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
    createSimRequest := *openapiclient.NewCreateSimRequest("Subscription_example", "Type_example") // CreateSimRequest | SIM の作成リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.CreateSim(context.Background()).CreateSimRequest(createSimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.CreateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.CreateSim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSimRequest** | [**CreateSimRequest**](CreateSimRequest.md) | SIM の作成リクエスト | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSimPacketCaptureSession

> PacketCaptureSession CreateSimPacketCaptureSession(ctx, simId).PacketCaptureSessionRequest(packetCaptureSessionRequest).Execute()

Create Packet Capture Session



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
    simId := "simId_example" // string | パケットキャプチャセッションを作成する SIM の ID
    packetCaptureSessionRequest := *openapiclient.NewPacketCaptureSessionRequest(int32(123)) // PacketCaptureSessionRequest | パケットキャプチャセッションリクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.CreateSimPacketCaptureSession(context.Background(), simId).PacketCaptureSessionRequest(packetCaptureSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.CreateSimPacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSimPacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `SimApi.CreateSimPacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | パケットキャプチャセッションを作成する SIM の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSimPacketCaptureSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **packetCaptureSessionRequest** | [**PacketCaptureSessionRequest**](PacketCaptureSessionRequest.md) | パケットキャプチャセッションリクエスト | 

### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSim

> Sim DeactivateSim(ctx, simId).Execute()

Deactivate SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.DeactivateSim(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.DeactivateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivateSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.DeactivateSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSimPacketCaptureSession

> PacketCaptureSession DeleteSimPacketCaptureSession(ctx, simId, sessionId).Execute()

Delete Packet Capture Session



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
    simId := "simId_example" // string | 対象の SIM ID
    sessionId := "sessionId_example" // string | 対象のパケットキャプチャセッション ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.DeleteSimPacketCaptureSession(context.Background(), simId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.DeleteSimPacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSimPacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `SimApi.DeleteSimPacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 
**sessionId** | **string** | 対象のパケットキャプチャセッション ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSimPacketCaptureSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSimSession

> Sim DeleteSimSession(ctx, simId).Execute()

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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.DeleteSimSession(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.DeleteSimSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSimSession`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.DeleteSimSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSimSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSimTag

> DeleteSimTag(ctx, simId, tagName).Execute()

Delete SIM Tag



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
    simId := "simId_example" // string | 対象の SIM ID
    tagName := "tagName_example" // string | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.DeleteSimTag(context.Background(), simId, tagName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.DeleteSimTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 
**tagName** | **string** | 削除対象のタグ名（URL の Path の一部になるので、パーセントエンコーディングを施す。JavaScript なら encodeURIComponent() したものを指定する） | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSimTagRequest struct via the builder pattern


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


## DeleteSubscriptionContainerCountryMappingEntry

> SubscriptionContainerStatus DeleteSubscriptionContainerCountryMappingEntry(ctx, simId, iccid, mcc).Execute()

Delete subscription container mapping entries



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
    simId := "simId_example" // string | SIM ID of the target SIM.
    iccid := "iccid_example" // string | Iccid of the target profile
    mcc := "mcc_example" // string | mobile country code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.DeleteSubscriptionContainerCountryMappingEntry(context.Background(), simId, iccid, mcc).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.DeleteSubscriptionContainerCountryMappingEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSubscriptionContainerCountryMappingEntry`: SubscriptionContainerStatus
    fmt.Fprintf(os.Stdout, "Response from `SimApi.DeleteSubscriptionContainerCountryMappingEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | SIM ID of the target SIM. | 
**iccid** | **string** | Iccid of the target profile | 
**mcc** | **string** | mobile country code | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionContainerCountryMappingEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SubscriptionContainerStatus**](SubscriptionContainerStatus.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableSimTermination

> Sim DisableSimTermination(ctx, simId).Execute()

Disable Termination of SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.DisableSimTermination(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.DisableSimTermination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableSimTermination`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.DisableSimTermination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableSimTerminationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableSimTermination

> Sim EnableSimTermination(ctx, simId).Execute()

Enable Termination of SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.EnableSimTermination(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.EnableSimTermination``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableSimTermination`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.EnableSimTermination`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSimTerminationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableSubscriptionContainer

> SubscriptionContainerStatus EnableSubscriptionContainer(ctx, simId, iccid, containerId).Execute()

Enables a subscription container.



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
    simId := "simId_example" // string | SIM ID of the target SIM.
    iccid := "iccid_example" // string | Iccid of the target profile
    containerId := "containerId_example" // string | Identifier of the target container

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.EnableSubscriptionContainer(context.Background(), simId, iccid, containerId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.EnableSubscriptionContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableSubscriptionContainer`: SubscriptionContainerStatus
    fmt.Fprintf(os.Stdout, "Response from `SimApi.EnableSubscriptionContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | SIM ID of the target SIM. | 
**iccid** | **string** | Iccid of the target profile | 
**containerId** | **string** | Identifier of the target container | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSubscriptionContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**SubscriptionContainerStatus**](SubscriptionContainerStatus.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDataFromSim

> []DataEntry GetDataFromSim(ctx, simId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

Get data sent from a SIM.



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
    simId := "simId_example" // string | データ取得対象の SIM の ID
    from := int32(56) // int32 | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) (optional)
    to := int32(56) // int32 | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) (optional)
    sort := "sort_example" // string | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 (optional) (default to "desc")
    limit := int32(56) // int32 | 取得するデータエントリ数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.GetDataFromSim(context.Background(), simId).From(from).To(to).Sort(sort).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.GetDataFromSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDataFromSim`: []DataEntry
    fmt.Fprintf(os.Stdout, "Response from `SimApi.GetDataFromSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | データ取得対象の SIM の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDataFromSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | 取得対象の期間の始まり (UNIX 時間 (ミリ秒)) | 
 **to** | **int32** | 取得対象の期間の終わり (UNIX 時間 (ミリ秒)) | 
 **sort** | **string** | データエントリのソート順。下降順（最新のデータが先頭）もしくは上昇順（最も古いデータが先頭）。 | [default to &quot;desc&quot;]
 **limit** | **int32** | 取得するデータエントリ数の上限 | 
 **lastEvaluatedKey** | **string** | 前ページで取得した最後のデータエントリのタイムスタンプ。このパラメータを指定することで次のデータエントリ以降のリストを取得できる。 | 

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


## GetSim

> Sim GetSim(ctx, simId).Execute()

Get SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.GetSim(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.GetSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.GetSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSimPacketCaptureSession

> PacketCaptureSession GetSimPacketCaptureSession(ctx, simId, sessionId).Execute()

Get Packet Capture Session



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
    simId := "simId_example" // string | 対象の SIM ID
    sessionId := "sessionId_example" // string | 対象のパケットキャプチャセッション ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.GetSimPacketCaptureSession(context.Background(), simId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.GetSimPacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSimPacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `SimApi.GetSimPacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 
**sessionId** | **string** | 対象のパケットキャプチャセッション ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSimPacketCaptureSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSimPacketCaptureSessions

> []PacketCaptureSession ListSimPacketCaptureSessions(ctx, simId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

List packet capture sessions associated with the SIM



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
    simId := "simId_example" // string | 対象の SIM ID
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のパケットキャプチャセッション。このパラメータを指定することで次のパケットキャプチャセッション以降のリストを取得できる。 (optional)
    limit := int32(56) // int32 | 取得するパケットキャプチャセッションの上限 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.ListSimPacketCaptureSessions(context.Background(), simId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.ListSimPacketCaptureSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSimPacketCaptureSessions`: []PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `SimApi.ListSimPacketCaptureSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSimPacketCaptureSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastEvaluatedKey** | **string** | 現ページで取得した最後のパケットキャプチャセッション。このパラメータを指定することで次のパケットキャプチャセッション以降のリストを取得できる。 | 
 **limit** | **int32** | 取得するパケットキャプチャセッションの上限 | [default to 10]

### Return type

[**[]PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSimSessionEvents

> []SessionEvent ListSimSessionEvents(ctx, simId).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

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
    simId := "simId_example" // string | 対象の SIM ID
    from := int32(56) // int32 | イベントの検索範囲時刻の始まり(unixtime) (optional)
    to := int32(56) // int32 | イベントの検索範囲時刻の終わり(unixtime) (optional)
    limit := int32(56) // int32 | 取得するイベント数の上限 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のイベントのタイムスタンプ。このパラメータを指定することで次のイベント以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.ListSimSessionEvents(context.Background(), simId).From(from).To(to).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.ListSimSessionEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSimSessionEvents`: []SessionEvent
    fmt.Fprintf(os.Stdout, "Response from `SimApi.ListSimSessionEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSimSessionEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | イベントの検索範囲時刻の始まり(unixtime) | 
 **to** | **int32** | イベントの検索範囲時刻の終わり(unixtime) | 
 **limit** | **int32** | 取得するイベント数の上限 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後のイベントのタイムスタンプ。このパラメータを指定することで次のイベント以降のリストを取得できる。 | 

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


## ListSims

> []Sim ListSims(ctx).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List SIMs



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
    limit := int32(56) // int32 | 取得する SIM の上限。ただし、返却される SIM の個数は指定された値を下回る可能性があります。 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の SIM ID。このパラメータを指定することで次の SIM 以降のリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.ListSims(context.Background()).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.ListSims``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSims`: []Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.ListSims`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSimsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | 取得する SIM の上限。ただし、返却される SIM の個数は指定された値を下回る可能性があります。 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の SIM ID。このパラメータを指定することで次の SIM 以降のリストを取得できる。 | 

### Return type

[**[]Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubscriptionContainers

> SubscriptionContainerStatus ListSubscriptionContainers(ctx, simId, iccid).Execute()

Get subscription container status.



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
    simId := "simId_example" // string | Sim Id of the target SIM.
    iccid := "iccid_example" // string | Iccid of the target profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.ListSubscriptionContainers(context.Background(), simId, iccid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.ListSubscriptionContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSubscriptionContainers`: SubscriptionContainerStatus
    fmt.Fprintf(os.Stdout, "Response from `SimApi.ListSubscriptionContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | Sim Id of the target SIM. | 
**iccid** | **string** | Iccid of the target profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubscriptionContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SubscriptionContainerStatus**](SubscriptionContainerStatus.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSimTags

> Sim PutSimTags(ctx, simId).TagUpdateRequest(tagUpdateRequest).Execute()

Bulk Insert or Update SIM Tags



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
    simId := "simId_example" // string | 対象の SIM ID
    tagUpdateRequest := []openapiclient.TagUpdateRequest{*openapiclient.NewTagUpdateRequest("TagName_example", "TagValue_example")} // []TagUpdateRequest | 追加・更新対象のタグの配列

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.PutSimTags(context.Background(), simId).TagUpdateRequest(tagUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.PutSimTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSimTags`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.PutSimTags`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSimTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tagUpdateRequest** | [**[]TagUpdateRequest**](TagUpdateRequest.md) | 追加・更新対象のタグの配列 | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutSubscriptionContainerCountryMappingEntries

> SubscriptionContainerStatus PutSubscriptionContainerCountryMappingEntries(ctx, simId, iccid).MappingEntries(mappingEntries).Execute()

Updates subscription container country mapping entries.



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
    simId := "simId_example" // string | SIM ID of the target SIM.
    iccid := "iccid_example" // string | Iccid of the target profile
    mappingEntries := *openapiclient.NewMappingEntries() // MappingEntries | collection of country (and optionally network) to subscription container mapping entries

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.PutSubscriptionContainerCountryMappingEntries(context.Background(), simId, iccid).MappingEntries(mappingEntries).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.PutSubscriptionContainerCountryMappingEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutSubscriptionContainerCountryMappingEntries`: SubscriptionContainerStatus
    fmt.Fprintf(os.Stdout, "Response from `SimApi.PutSubscriptionContainerCountryMappingEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | SIM ID of the target SIM. | 
**iccid** | **string** | Iccid of the target profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutSubscriptionContainerCountryMappingEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mappingEntries** | [**MappingEntries**](MappingEntries.md) | collection of country (and optionally network) to subscription container mapping entries | 

### Return type

[**SubscriptionContainerStatus**](SubscriptionContainerStatus.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterSim

> Sim RegisterSim(ctx, simId).RegisterSimRequest(registerSimRequest).Execute()

Register SIM



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
    simId := "simId_example" // string | 対象の SIM ID
    registerSimRequest := *openapiclient.NewRegisterSimRequest("RegistrationSecret_example") // RegisterSimRequest | SIM 登録リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.RegisterSim(context.Background(), simId).RegisterSimRequest(registerSimRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.RegisterSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.RegisterSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerSimRequest** | [**RegisterSimRequest**](RegisterSimRequest.md) | SIM 登録リクエスト | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveArcSimCredentials

> RemoveArcSimCredentials(ctx, simId).Execute()

(非推奨) Arc のクレデンシャルを SIM から除去



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
    simId := "simId_example" // string | 対象とする SIM の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.RemoveArcSimCredentials(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.RemoveArcSimCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象とする SIM の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveArcSimCredentialsRequest struct via the builder pattern


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


## RenewArcSimCredentials

> ArcCredentialRenewResponse RenewArcSimCredentials(ctx, simId).ArcCredentialRenewRequest(arcCredentialRenewRequest).Execute()

SIM に対する Arc のクレデンシャルを更新する



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
    simId := "simId_example" // string | 対象とする SIM の SIM ID
    arcCredentialRenewRequest := *openapiclient.NewArcCredentialRenewRequest() // ArcCredentialRenewRequest | Arc のクレデンシャルの付与リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.RenewArcSimCredentials(context.Background(), simId).ArcCredentialRenewRequest(arcCredentialRenewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.RenewArcSimCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenewArcSimCredentials`: ArcCredentialRenewResponse
    fmt.Fprintf(os.Stdout, "Response from `SimApi.RenewArcSimCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象とする SIM の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenewArcSimCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **arcCredentialRenewRequest** | [**ArcCredentialRenewRequest**](ArcCredentialRenewRequest.md) | Arc のクレデンシャルの付与リクエスト | 

### Return type

[**ArcCredentialRenewResponse**](ArcCredentialRenewResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportSimLocalInfo

> Subscriber ReportSimLocalInfo(ctx, simId).Execute()

SIM ローカル情報レポートを取得する



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
    simId := "simId_example" // string | 対象の SIM の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.ReportSimLocalInfo(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.ReportSimLocalInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportSimLocalInfo`: Subscriber
    fmt.Fprintf(os.Stdout, "Response from `SimApi.ReportSimLocalInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportSimLocalInfoRequest struct via the builder pattern


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


## SendDownlinkPing

> DownlinkPingResponse SendDownlinkPing(ctx, simId).DownlinkPingRequest(downlinkPingRequest).Execute()

SIM に対して ping リクエストを送信する。



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
    simId := "simId_example" // string | 対象の SIM の SIM ID
    downlinkPingRequest := *openapiclient.NewDownlinkPingRequest() // DownlinkPingRequest | ping リクエストのオプション

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SendDownlinkPing(context.Background(), simId).DownlinkPingRequest(downlinkPingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SendDownlinkPing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendDownlinkPing`: DownlinkPingResponse
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SendDownlinkPing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendDownlinkPingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **downlinkPingRequest** | [**DownlinkPingRequest**](DownlinkPingRequest.md) | ping リクエストのオプション | 

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


## SendSmsToSim

> SmsForwardingReport SendSmsToSim(ctx, simId).SmsForwardingRequest(smsForwardingRequest).Execute()

Send SMS to SIM



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
    simId := "simId_example" // string | SMS 送信対象の SIM の ID
    smsForwardingRequest := *openapiclient.NewSmsForwardingRequest() // SmsForwardingRequest | メッセージ本体とエンコーディングタイプからなる SMS 送信リクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SendSmsToSim(context.Background(), simId).SmsForwardingRequest(smsForwardingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SendSmsToSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSmsToSim`: SmsForwardingReport
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SendSmsToSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | SMS 送信対象の SIM の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendSmsToSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **smsForwardingRequest** | [**SmsForwardingRequest**](SmsForwardingRequest.md) | メッセージ本体とエンコーディングタイプからなる SMS 送信リクエスト | 

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


## SetSimExpiryTime

> Sim SetSimExpiryTime(ctx, simId).ExpiryTime(expiryTime).Execute()

Update Expiry Time of SIM



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
    simId := "simId_example" // string | 対象の SIM ID
    expiryTime := *openapiclient.NewExpiryTime(int64(123)) // ExpiryTime | 更新後の有効期限 (unixtime:ミリ秒単位) とアクション

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SetSimExpiryTime(context.Background(), simId).ExpiryTime(expiryTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SetSimExpiryTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSimExpiryTime`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SetSimExpiryTime`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSimExpiryTimeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expiryTime** | [**ExpiryTime**](ExpiryTime.md) | 更新後の有効期限 (unixtime:ミリ秒単位) とアクション | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSimGroup

> Sim SetSimGroup(ctx, simId).SetGroupRequest(setGroupRequest).Execute()

Set Group to SIM



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
    simId := "simId_example" // string | 対象の SIM ID
    setGroupRequest := *openapiclient.NewSetGroupRequest() // SetGroupRequest | Group（ID のみを含めばよい）

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SetSimGroup(context.Background(), simId).SetGroupRequest(setGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SetSimGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSimGroup`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SetSimGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSimGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setGroupRequest** | [**SetGroupRequest**](SetGroupRequest.md) | Group（ID のみを含めばよい） | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSimImeiLock

> Sim SetSimImeiLock(ctx, simId).SetImeiLockRequest(setImeiLockRequest).Execute()

Set IMEI lock configuration for SIM.



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
    simId := "simId_example" // string | 対象の SIM ID
    setImeiLockRequest := *openapiclient.NewSetImeiLockRequest() // SetImeiLockRequest | SIM に指定する IMEI lock の設定 （オンラインの SIM の現在の IMEI にロックするには IMEI は指定しなくてもよい） (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SetSimImeiLock(context.Background(), simId).SetImeiLockRequest(setImeiLockRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SetSimImeiLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSimImeiLock`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SetSimImeiLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSimImeiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **setImeiLockRequest** | [**SetImeiLockRequest**](SetImeiLockRequest.md) | SIM に指定する IMEI lock の設定 （オンラインの SIM の現在の IMEI にロックするには IMEI は指定しなくてもよい） | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetSimToStandby

> Sim SetSimToStandby(ctx, simId).Execute()

Set SIM to standby mode.



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SetSimToStandby(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SetSimToStandby``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetSimToStandby`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SetSimToStandby`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetSimToStandbyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopSimPacketCaptureSession

> PacketCaptureSession StopSimPacketCaptureSession(ctx, simId, sessionId).Execute()

Stop SIM Packet Capture Session



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
    simId := "simId_example" // string | 対象の SIM ID
    sessionId := "sessionId_example" // string | 対象のパケットキャプチャセッション ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.StopSimPacketCaptureSession(context.Background(), simId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.StopSimPacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopSimPacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `SimApi.StopSimPacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 
**sessionId** | **string** | 対象のパケットキャプチャセッション ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopSimPacketCaptureSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PacketCaptureSession**](PacketCaptureSession.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendSim

> Sim SuspendSim(ctx, simId).Execute()

Suspend SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.SuspendSim(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.SuspendSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.SuspendSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateSim

> Sim TerminateSim(ctx, simId).Execute()

Terminate SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.TerminateSim(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.TerminateSim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TerminateSim`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.TerminateSim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateSimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TerminateSubscriptionContainer

> TerminateSubscriptionContainer(ctx, simId, iccid, imsi).Execute()

セカンダリサブスクリプションの利用終了



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
    simId := "simId_example" // string | 対象のプライマリサブスクリプションの IoT SIM の SIM ID
    iccid := "iccid_example" // string | 対象のプライマリサブスクリプションの IoT SIM の ICCID
    imsi := "imsi_example" // string | 対象のセカンダリサブスクリプションのバーチャル SIM/Subscriber の IMSI

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.TerminateSubscriptionContainer(context.Background(), simId, iccid, imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.TerminateSubscriptionContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象のプライマリサブスクリプションの IoT SIM の SIM ID | 
**iccid** | **string** | 対象のプライマリサブスクリプションの IoT SIM の ICCID | 
**imsi** | **string** | 対象のセカンダリサブスクリプションのバーチャル SIM/Subscriber の IMSI | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateSubscriptionContainerRequest struct via the builder pattern


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


## UnsetSimExpiryTime

> UnsetSimExpiryTime(ctx, simId).Execute()

Delete Expiry Time of SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.UnsetSimExpiryTime(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.UnsetSimExpiryTime``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetSimExpiryTimeRequest struct via the builder pattern


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


## UnsetSimGroup

> Sim UnsetSimGroup(ctx, simId).Execute()

Unset Group to SIM



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.UnsetSimGroup(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.UnsetSimGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetSimGroup`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.UnsetSimGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetSimGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnsetSimImeiLock

> Sim UnsetSimImeiLock(ctx, simId).Execute()

Unset IMEI lock configuration for SIM.



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
    simId := "simId_example" // string | 対象の SIM ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.UnsetSimImeiLock(context.Background(), simId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.UnsetSimImeiLock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnsetSimImeiLock`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.UnsetSimImeiLock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetSimImeiLockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSimSpeedClass

> Sim UpdateSimSpeedClass(ctx, simId).UpdateSpeedClassRequest(updateSpeedClassRequest).Execute()

Update SIM speed class



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
    simId := "simId_example" // string | 対象の SIM ID
    updateSpeedClassRequest := *openapiclient.NewUpdateSpeedClassRequest("SpeedClass_example") // UpdateSpeedClassRequest | speed_class

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SimApi.UpdateSimSpeedClass(context.Background(), simId).UpdateSpeedClassRequest(updateSpeedClassRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SimApi.UpdateSimSpeedClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSimSpeedClass`: Sim
    fmt.Fprintf(os.Stdout, "Response from `SimApi.UpdateSimSpeedClass`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**simId** | **string** | 対象の SIM ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSimSpeedClassRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSpeedClassRequest** | [**UpdateSpeedClassRequest**](UpdateSpeedClassRequest.md) | speed_class | 

### Return type

[**Sim**](Sim.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

