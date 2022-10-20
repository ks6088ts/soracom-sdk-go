# \VirtualPrivateGatewayApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloseGate**](VirtualPrivateGatewayApi.md#CloseGate) | **Post** /virtual_private_gateways/{vpg_id}/gate/close | Close SORACOM Gate.
[**CreateMirroringPeer**](VirtualPrivateGatewayApi.md#CreateMirroringPeer) | **Post** /virtual_private_gateways/{vpg_id}/junction/mirroring/peers | Junction ミラーリング peer を追加します。
[**CreatePacketCaptureSession**](VirtualPrivateGatewayApi.md#CreatePacketCaptureSession) | **Post** /virtual_private_gateways/{vpg_id}/packet_capture_sessions | Create Packet Capture Session
[**CreateVirtualPrivateGateway**](VirtualPrivateGatewayApi.md#CreateVirtualPrivateGateway) | **Post** /virtual_private_gateways | Create Virtual Private Gateway
[**CreateVpcPeeringConnection**](VirtualPrivateGatewayApi.md#CreateVpcPeeringConnection) | **Post** /virtual_private_gateways/{vpg_id}/vpc_peering_connections | Create VPC Peering Connection
[**DeleteMirroringPeer**](VirtualPrivateGatewayApi.md#DeleteMirroringPeer) | **Delete** /virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr} | Junction ミラーリング peer を削除します。
[**DeletePacketCaptureSession**](VirtualPrivateGatewayApi.md#DeletePacketCaptureSession) | **Delete** /virtual_private_gateways/{vpg_id}/packet_capture_sessions/{session_id} | Delete Packet Capture Session
[**DeleteVirtualPrivateGatewayIpAddressMapEntry**](VirtualPrivateGatewayApi.md#DeleteVirtualPrivateGatewayIpAddressMapEntry) | **Delete** /virtual_private_gateways/{vpg_id}/ip_address_map/{key} | Delete VPG IP address map entry
[**DeleteVpcPeeringConnection**](VirtualPrivateGatewayApi.md#DeleteVpcPeeringConnection) | **Delete** /virtual_private_gateways/{vpg_id}/vpc_peering_connections/{pcx_id} | Delete VPC Peering Connection
[**DisableGatePrivacySeparator**](VirtualPrivateGatewayApi.md#DisableGatePrivacySeparator) | **Post** /virtual_private_gateways/{vpg_id}/gate/disable_privacy_separator | 指定された VPG に対して SORACOM Gate のプライバシーセパレーター機能を無効にする。
[**EnableGatePrivacySeparator**](VirtualPrivateGatewayApi.md#EnableGatePrivacySeparator) | **Post** /virtual_private_gateways/{vpg_id}/gate/enable_privacy_separator | 指定された VPG に対して SORACOM Gate のプライバシーセパレーター機能を有効にする。
[**GetPacketCaptureSession**](VirtualPrivateGatewayApi.md#GetPacketCaptureSession) | **Get** /virtual_private_gateways/{vpg_id}/packet_capture_sessions/{session_id} | Get Packet Capture Session
[**GetVirtualPrivateGateway**](VirtualPrivateGatewayApi.md#GetVirtualPrivateGateway) | **Get** /virtual_private_gateways/{vpg_id} | Get Virtual Private Gateway
[**ListGatePeers**](VirtualPrivateGatewayApi.md#ListGatePeers) | **Get** /virtual_private_gateways/{vpg_id}/gate/peers | List VPG Gate peers
[**ListPacketCaptureSessions**](VirtualPrivateGatewayApi.md#ListPacketCaptureSessions) | **Get** /virtual_private_gateways/{vpg_id}/packet_capture_sessions | List Packet Capture Sessions
[**ListVirtualPrivateGatewayIpAddressMapEntries**](VirtualPrivateGatewayApi.md#ListVirtualPrivateGatewayIpAddressMapEntries) | **Get** /virtual_private_gateways/{vpg_id}/ip_address_map | List VPG IP address map entries
[**ListVirtualPrivateGateways**](VirtualPrivateGatewayApi.md#ListVirtualPrivateGateways) | **Get** /virtual_private_gateways | List Virtual Private Gateways
[**OpenGate**](VirtualPrivateGatewayApi.md#OpenGate) | **Post** /virtual_private_gateways/{vpg_id}/gate/open | Open SORACOM Gate
[**PutVirtualPrivateGatewayIpAddressMapEntry**](VirtualPrivateGatewayApi.md#PutVirtualPrivateGatewayIpAddressMapEntry) | **Post** /virtual_private_gateways/{vpg_id}/ip_address_map | Put VPG IP address map entry
[**RegisterGatePeer**](VirtualPrivateGatewayApi.md#RegisterGatePeer) | **Post** /virtual_private_gateways/{vpg_id}/gate/peers | Register VPG Gate Peer
[**SetInspectionConfiguration**](VirtualPrivateGatewayApi.md#SetInspectionConfiguration) | **Post** /virtual_private_gateways/{vpg_id}/junction/set_inspection | Junction インスペクション機能の設定を行います。
[**SetRedirectionConfiguration**](VirtualPrivateGatewayApi.md#SetRedirectionConfiguration) | **Post** /virtual_private_gateways/{vpg_id}/junction/set_redirection | Junction リダイレクション機能の設定を行います。
[**SetRoutingFilter**](VirtualPrivateGatewayApi.md#SetRoutingFilter) | **Post** /virtual_private_gateways/{vpg_id}/set_routing_filter | Sets Virtual Private Gateway outbound routing filter.
[**StopPacketCaptureSession**](VirtualPrivateGatewayApi.md#StopPacketCaptureSession) | **Post** /virtual_private_gateways/{vpg_id}/packet_capture_sessions/{session_id}/stop | Stop Packet Capture Session
[**TerminateVirtualPrivateGateway**](VirtualPrivateGatewayApi.md#TerminateVirtualPrivateGateway) | **Post** /virtual_private_gateways/{vpg_id}/terminate | Terminate Virtual Private Gateway
[**UnregisterGatePeer**](VirtualPrivateGatewayApi.md#UnregisterGatePeer) | **Delete** /virtual_private_gateways/{vpg_id}/gate/peers/{outer_ip_address} | Unregister VPG Gate Peer
[**UnsetInspectionConfiguration**](VirtualPrivateGatewayApi.md#UnsetInspectionConfiguration) | **Post** /virtual_private_gateways/{vpg_id}/junction/unset_inspection | Junction インスペクション機能の設定を解除します。
[**UnsetRedirectionConfiguration**](VirtualPrivateGatewayApi.md#UnsetRedirectionConfiguration) | **Post** /virtual_private_gateways/{vpg_id}/junction/unset_redirection | Junction リダイレクション機能の設定を行います。
[**UpdateMirroringPeer**](VirtualPrivateGatewayApi.md#UpdateMirroringPeer) | **Put** /virtual_private_gateways/{vpg_id}/junction/mirroring/peers/{ipaddr} | Junction ミラーリング peer を更新します。



## CloseGate

> CloseGate(ctx, vpgId).Execute()

Close SORACOM Gate.



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.CloseGate(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.CloseGate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloseGateRequest struct via the builder pattern


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


## CreateMirroringPeer

> CreateMirroringPeer(ctx, vpgId).JunctionMirroringPeer(junctionMirroringPeer).Execute()

Junction ミラーリング peer を追加します。



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
    vpgId := "vpgId_example" // string | VPG ID
    junctionMirroringPeer := *openapiclient.NewJunctionMirroringPeer() // JunctionMirroringPeer | ミラーリング peer に関する情報

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.CreateMirroringPeer(context.Background(), vpgId).JunctionMirroringPeer(junctionMirroringPeer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.CreateMirroringPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMirroringPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **junctionMirroringPeer** | [**JunctionMirroringPeer**](JunctionMirroringPeer.md) | ミラーリング peer に関する情報 | 

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


## CreatePacketCaptureSession

> PacketCaptureSession CreatePacketCaptureSession(ctx, vpgId).PacketCaptureSessionRequest(packetCaptureSessionRequest).Execute()

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
    vpgId := "vpgId_example" // string | 対象の VPG ID
    packetCaptureSessionRequest := *openapiclient.NewPacketCaptureSessionRequest(int32(123)) // PacketCaptureSessionRequest | パケットキャプチャセッションリクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.CreatePacketCaptureSession(context.Background(), vpgId).PacketCaptureSessionRequest(packetCaptureSessionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.CreatePacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.CreatePacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePacketCaptureSessionRequest struct via the builder pattern


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


## CreateVirtualPrivateGateway

> VirtualPrivateGateway CreateVirtualPrivateGateway(ctx).CreateVirtualPrivateGatewayRequest(createVirtualPrivateGatewayRequest).Execute()

Create Virtual Private Gateway



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
    createVirtualPrivateGatewayRequest := *openapiclient.NewCreateVirtualPrivateGatewayRequest(int32(123)) // CreateVirtualPrivateGatewayRequest | 作成対象の VPG の情報を含むリクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.CreateVirtualPrivateGateway(context.Background()).CreateVirtualPrivateGatewayRequest(createVirtualPrivateGatewayRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.CreateVirtualPrivateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVirtualPrivateGateway`: VirtualPrivateGateway
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.CreateVirtualPrivateGateway`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVirtualPrivateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVirtualPrivateGatewayRequest** | [**CreateVirtualPrivateGatewayRequest**](CreateVirtualPrivateGatewayRequest.md) | 作成対象の VPG の情報を含むリクエスト | 

### Return type

[**VirtualPrivateGateway**](VirtualPrivateGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVpcPeeringConnection

> CreateVpcPeeringConnectionRequest CreateVpcPeeringConnection(ctx, vpgId).CreateVpcPeeringConnectionRequest(createVpcPeeringConnectionRequest).Execute()

Create VPC Peering Connection



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    createVpcPeeringConnectionRequest := *openapiclient.NewCreateVpcPeeringConnectionRequest() // CreateVpcPeeringConnectionRequest | 作成対象の VPC Peering Connection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.CreateVpcPeeringConnection(context.Background(), vpgId).CreateVpcPeeringConnectionRequest(createVpcPeeringConnectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.CreateVpcPeeringConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVpcPeeringConnection`: CreateVpcPeeringConnectionRequest
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.CreateVpcPeeringConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcPeeringConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVpcPeeringConnectionRequest** | [**CreateVpcPeeringConnectionRequest**](CreateVpcPeeringConnectionRequest.md) | 作成対象の VPC Peering Connection | 

### Return type

[**CreateVpcPeeringConnectionRequest**](CreateVpcPeeringConnectionRequest.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMirroringPeer

> DeleteMirroringPeer(ctx, vpgId, ipaddr).Execute()

Junction ミラーリング peer を削除します。



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
    vpgId := "vpgId_example" // string | VPG ID
    ipaddr := "ipaddr_example" // string | 削除したいミラーリング peer の IP アドレス

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.DeleteMirroringPeer(context.Background(), vpgId, ipaddr).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.DeleteMirroringPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 
**ipaddr** | **string** | 削除したいミラーリング peer の IP アドレス | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMirroringPeerRequest struct via the builder pattern


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


## DeletePacketCaptureSession

> PacketCaptureSession DeletePacketCaptureSession(ctx, vpgId, sessionId).Execute()

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
    vpgId := "vpgId_example" // string | 対象の VPG ID
    sessionId := "sessionId_example" // string | 対象のパケットキャプチャセッション ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.DeletePacketCaptureSession(context.Background(), vpgId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.DeletePacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.DeletePacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG ID | 
**sessionId** | **string** | 対象のパケットキャプチャセッション ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePacketCaptureSessionRequest struct via the builder pattern


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


## DeleteVirtualPrivateGatewayIpAddressMapEntry

> IpAddressMapEntry DeleteVirtualPrivateGatewayIpAddressMapEntry(ctx, vpgId, key).Execute()

Delete VPG IP address map entry



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    key := "key_example" // string | 対象の Key

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.DeleteVirtualPrivateGatewayIpAddressMapEntry(context.Background(), vpgId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.DeleteVirtualPrivateGatewayIpAddressMapEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVirtualPrivateGatewayIpAddressMapEntry`: IpAddressMapEntry
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.DeleteVirtualPrivateGatewayIpAddressMapEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 
**key** | **string** | 対象の Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVirtualPrivateGatewayIpAddressMapEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IpAddressMapEntry**](IpAddressMapEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcPeeringConnection

> DeleteVpcPeeringConnection(ctx, vpgId, pcxId).Execute()

Delete VPC Peering Connection



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    pcxId := "pcxId_example" // string | 削除対象の VPC Peering Connection の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.DeleteVpcPeeringConnection(context.Background(), vpgId, pcxId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.DeleteVpcPeeringConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 
**pcxId** | **string** | 削除対象の VPC Peering Connection の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcPeeringConnectionRequest struct via the builder pattern


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


## DisableGatePrivacySeparator

> VirtualPrivateGateway DisableGatePrivacySeparator(ctx, vpgId).Execute()

指定された VPG に対して SORACOM Gate のプライバシーセパレーター機能を無効にする。



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
    vpgId := "vpgId_example" // string | VPG ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.DisableGatePrivacySeparator(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.DisableGatePrivacySeparator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisableGatePrivacySeparator`: VirtualPrivateGateway
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.DisableGatePrivacySeparator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableGatePrivacySeparatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualPrivateGateway**](VirtualPrivateGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnableGatePrivacySeparator

> VirtualPrivateGateway EnableGatePrivacySeparator(ctx, vpgId).Execute()

指定された VPG に対して SORACOM Gate のプライバシーセパレーター機能を有効にする。



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
    vpgId := "vpgId_example" // string | VPG ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.EnableGatePrivacySeparator(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.EnableGatePrivacySeparator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableGatePrivacySeparator`: VirtualPrivateGateway
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.EnableGatePrivacySeparator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableGatePrivacySeparatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualPrivateGateway**](VirtualPrivateGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPacketCaptureSession

> PacketCaptureSession GetPacketCaptureSession(ctx, vpgId, sessionId).Execute()

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
    vpgId := "vpgId_example" // string | 対象の VPG ID
    sessionId := "sessionId_example" // string | パケットキャプチャセッションリクエスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.GetPacketCaptureSession(context.Background(), vpgId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.GetPacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.GetPacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG ID | 
**sessionId** | **string** | パケットキャプチャセッションリクエスト | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPacketCaptureSessionRequest struct via the builder pattern


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


## GetVirtualPrivateGateway

> VirtualPrivateGateway GetVirtualPrivateGateway(ctx, vpgId).Execute()

Get Virtual Private Gateway



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.GetVirtualPrivateGateway(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.GetVirtualPrivateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualPrivateGateway`: VirtualPrivateGateway
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.GetVirtualPrivateGateway`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualPrivateGatewayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VirtualPrivateGateway**](VirtualPrivateGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatePeers

> []GatePeer ListGatePeers(ctx, vpgId).Execute()

List VPG Gate peers



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.ListGatePeers(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.ListGatePeers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGatePeers`: []GatePeer
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.ListGatePeers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListGatePeersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GatePeer**](GatePeer.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPacketCaptureSessions

> []PacketCaptureSession ListPacketCaptureSessions(ctx, vpgId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()

List Packet Capture Sessions



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
    vpgId := "vpgId_example" // string | 対象の VPG ID
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後のパケットキャプチャセッション。このパラメータを指定することで次のパケットキャプチャセッション以降のリストを取得できる。 (optional) (default to "null")
    limit := int32(56) // int32 | 取得するパケットキャプチャセッションの上限 (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.ListPacketCaptureSessions(context.Background(), vpgId).LastEvaluatedKey(lastEvaluatedKey).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.ListPacketCaptureSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPacketCaptureSessions`: []PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.ListPacketCaptureSessions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPacketCaptureSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lastEvaluatedKey** | **string** | 現ページで取得した最後のパケットキャプチャセッション。このパラメータを指定することで次のパケットキャプチャセッション以降のリストを取得できる。 | [default to &quot;null&quot;]
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


## ListVirtualPrivateGatewayIpAddressMapEntries

> []IpAddressMapEntry ListVirtualPrivateGatewayIpAddressMapEntries(ctx, vpgId).Execute()

List VPG IP address map entries



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.ListVirtualPrivateGatewayIpAddressMapEntries(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.ListVirtualPrivateGatewayIpAddressMapEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualPrivateGatewayIpAddressMapEntries`: []IpAddressMapEntry
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.ListVirtualPrivateGatewayIpAddressMapEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualPrivateGatewayIpAddressMapEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IpAddressMapEntry**](IpAddressMapEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualPrivateGateways

> []VirtualPrivateGateway ListVirtualPrivateGateways(ctx).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List Virtual Private Gateways



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
    tagName := "tagName_example" // string | VPG についた Tag 名。完全一致する Tag 名が検索対象となる。tag_name を指定した場合、tag_value が必須。 (optional)
    tagValue := "tagValue_example" // string | VPG についた Tag の値。 (optional)
    tagValueMatchMode := "tagValueMatchMode_example" // string | タグの検索条件。 (optional) (default to "exact")
    limit := int32(56) // int32 | レスポンス 1 ページあたりの最大数 (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | 現ページで取得した最後の Group の ID。このパラメータを指定することで次の VPG から始まるリストを取得できる。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.ListVirtualPrivateGateways(context.Background()).TagName(tagName).TagValue(tagValue).TagValueMatchMode(tagValueMatchMode).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.ListVirtualPrivateGateways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualPrivateGateways`: []VirtualPrivateGateway
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.ListVirtualPrivateGateways`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualPrivateGatewaysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tagName** | **string** | VPG についた Tag 名。完全一致する Tag 名が検索対象となる。tag_name を指定した場合、tag_value が必須。 | 
 **tagValue** | **string** | VPG についた Tag の値。 | 
 **tagValueMatchMode** | **string** | タグの検索条件。 | [default to &quot;exact&quot;]
 **limit** | **int32** | レスポンス 1 ページあたりの最大数 | 
 **lastEvaluatedKey** | **string** | 現ページで取得した最後の Group の ID。このパラメータを指定することで次の VPG から始まるリストを取得できる。 | 

### Return type

[**[]VirtualPrivateGateway**](VirtualPrivateGateway.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OpenGate

> OpenGate(ctx, vpgId).OpenGateRequest(openGateRequest).Execute()

Open SORACOM Gate



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    openGateRequest := *openapiclient.NewOpenGateRequest() // OpenGateRequest | オプショナルな Gate の設定パラメータ。 (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.OpenGate(context.Background(), vpgId).OpenGateRequest(openGateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.OpenGate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiOpenGateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **openGateRequest** | [**OpenGateRequest**](OpenGateRequest.md) | オプショナルな Gate の設定パラメータ。 | 

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


## PutVirtualPrivateGatewayIpAddressMapEntry

> IpAddressMapEntry PutVirtualPrivateGatewayIpAddressMapEntry(ctx, vpgId).PutIpAddressMapEntryRequest(putIpAddressMapEntryRequest).Execute()

Put VPG IP address map entry



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    putIpAddressMapEntryRequest := *openapiclient.NewPutIpAddressMapEntryRequest("IpAddress_example", "Key_example") // PutIpAddressMapEntryRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.PutVirtualPrivateGatewayIpAddressMapEntry(context.Background(), vpgId).PutIpAddressMapEntryRequest(putIpAddressMapEntryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.PutVirtualPrivateGatewayIpAddressMapEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutVirtualPrivateGatewayIpAddressMapEntry`: IpAddressMapEntry
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.PutVirtualPrivateGatewayIpAddressMapEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVirtualPrivateGatewayIpAddressMapEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putIpAddressMapEntryRequest** | [**PutIpAddressMapEntryRequest**](PutIpAddressMapEntryRequest.md) |  | 

### Return type

[**IpAddressMapEntry**](IpAddressMapEntry.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterGatePeer

> GatePeer RegisterGatePeer(ctx, vpgId).RegisterGatePeerRequest(registerGatePeerRequest).Execute()

Register VPG Gate Peer



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    registerGatePeerRequest := *openapiclient.NewRegisterGatePeerRequest("OuterIpAddress_example") // RegisterGatePeerRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.RegisterGatePeer(context.Background(), vpgId).RegisterGatePeerRequest(registerGatePeerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.RegisterGatePeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterGatePeer`: GatePeer
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.RegisterGatePeer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterGatePeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registerGatePeerRequest** | [**RegisterGatePeerRequest**](RegisterGatePeerRequest.md) |  | 

### Return type

[**GatePeer**](GatePeer.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetInspectionConfiguration

> SetInspectionConfiguration(ctx, vpgId).JunctionInspectionConfiguration(junctionInspectionConfiguration).Execute()

Junction インスペクション機能の設定を行います。



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
    vpgId := "vpgId_example" // string | VPG ID
    junctionInspectionConfiguration := *openapiclient.NewJunctionInspectionConfiguration() // JunctionInspectionConfiguration | インスペクション設定内容

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.SetInspectionConfiguration(context.Background(), vpgId).JunctionInspectionConfiguration(junctionInspectionConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.SetInspectionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetInspectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **junctionInspectionConfiguration** | [**JunctionInspectionConfiguration**](JunctionInspectionConfiguration.md) | インスペクション設定内容 | 

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


## SetRedirectionConfiguration

> SetRedirectionConfiguration(ctx, vpgId).JunctionRedirectionConfiguration(junctionRedirectionConfiguration).Execute()

Junction リダイレクション機能の設定を行います。



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
    vpgId := "vpgId_example" // string | VPG ID
    junctionRedirectionConfiguration := *openapiclient.NewJunctionRedirectionConfiguration() // JunctionRedirectionConfiguration | リダイレクション設定内容

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.SetRedirectionConfiguration(context.Background(), vpgId).JunctionRedirectionConfiguration(junctionRedirectionConfiguration).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.SetRedirectionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRedirectionConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **junctionRedirectionConfiguration** | [**JunctionRedirectionConfiguration**](JunctionRedirectionConfiguration.md) | リダイレクション設定内容 | 

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


## SetRoutingFilter

> SetRoutingFilter(ctx, vpgId).RoutingFilterEntry(routingFilterEntry).Execute()

Sets Virtual Private Gateway outbound routing filter.



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    routingFilterEntry := []openapiclient.RoutingFilterEntry{*openapiclient.NewRoutingFilterEntry("Action_example", "IpRange_example")} // []RoutingFilterEntry | ルーティングフィルタエントリのリスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.SetRoutingFilter(context.Background(), vpgId).RoutingFilterEntry(routingFilterEntry).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.SetRoutingFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetRoutingFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **routingFilterEntry** | [**[]RoutingFilterEntry**](RoutingFilterEntry.md) | ルーティングフィルタエントリのリスト | 

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


## StopPacketCaptureSession

> PacketCaptureSession StopPacketCaptureSession(ctx, vpgId, sessionId).Execute()

Stop Packet Capture Session



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
    vpgId := "vpgId_example" // string | 対象の VPG ID
    sessionId := "sessionId_example" // string | 対象のパケットキャプチャセッション ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.StopPacketCaptureSession(context.Background(), vpgId, sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.StopPacketCaptureSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopPacketCaptureSession`: PacketCaptureSession
    fmt.Fprintf(os.Stdout, "Response from `VirtualPrivateGatewayApi.StopPacketCaptureSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG ID | 
**sessionId** | **string** | 対象のパケットキャプチャセッション ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPacketCaptureSessionRequest struct via the builder pattern


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


## TerminateVirtualPrivateGateway

> TerminateVirtualPrivateGateway(ctx, vpgId).Execute()

Terminate Virtual Private Gateway



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.TerminateVirtualPrivateGateway(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.TerminateVirtualPrivateGateway``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTerminateVirtualPrivateGatewayRequest struct via the builder pattern


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


## UnregisterGatePeer

> UnregisterGatePeer(ctx, vpgId, outerIpAddress).Execute()

Unregister VPG Gate Peer



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
    vpgId := "vpgId_example" // string | 対象の VPG の ID
    outerIpAddress := "outerIpAddress_example" // string | 対象の Peer の ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.UnregisterGatePeer(context.Background(), vpgId, outerIpAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.UnregisterGatePeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | 対象の VPG の ID | 
**outerIpAddress** | **string** | 対象の Peer の ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterGatePeerRequest struct via the builder pattern


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


## UnsetInspectionConfiguration

> UnsetInspectionConfiguration(ctx, vpgId).Execute()

Junction インスペクション機能の設定を解除します。



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
    vpgId := "vpgId_example" // string | VPG ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.UnsetInspectionConfiguration(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.UnsetInspectionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetInspectionConfigurationRequest struct via the builder pattern


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


## UnsetRedirectionConfiguration

> UnsetRedirectionConfiguration(ctx, vpgId).Execute()

Junction リダイレクション機能の設定を行います。



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
    vpgId := "vpgId_example" // string | VPG ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.UnsetRedirectionConfiguration(context.Background(), vpgId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.UnsetRedirectionConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnsetRedirectionConfigurationRequest struct via the builder pattern


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


## UpdateMirroringPeer

> UpdateMirroringPeer(ctx, vpgId, ipaddr).AttributeUpdate(attributeUpdate).Execute()

Junction ミラーリング peer を更新します。



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
    vpgId := "vpgId_example" // string | VPG ID
    ipaddr := "ipaddr_example" // string | 更新したいミラーリング peer の IP アドレス
    attributeUpdate := []openapiclient.AttributeUpdate{*openapiclient.NewAttributeUpdate()} // []AttributeUpdate | 更新する属性のリスト

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VirtualPrivateGatewayApi.UpdateMirroringPeer(context.Background(), vpgId, ipaddr).AttributeUpdate(attributeUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VirtualPrivateGatewayApi.UpdateMirroringPeer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpgId** | **string** | VPG ID | 
**ipaddr** | **string** | 更新したいミラーリング peer の IP アドレス | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMirroringPeerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attributeUpdate** | [**[]AttributeUpdate**](AttributeUpdate.md) | 更新する属性のリスト | 

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

