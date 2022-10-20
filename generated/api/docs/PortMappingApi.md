# \PortMappingApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePortMapping**](PortMappingApi.md#CreatePortMapping) | **Post** /port_mappings | Create Port Mapping.
[**DeletePortMapping**](PortMappingApi.md#DeletePortMapping) | **Delete** /port_mappings/{ip_address}/{port} | Delete PortMapping.
[**ListPortMappings**](PortMappingApi.md#ListPortMappings) | **Get** /port_mappings | List Port Mapping Entries.
[**ListPortMappingsForSubscriber**](PortMappingApi.md#ListPortMappingsForSubscriber) | **Get** /port_mappings/subscribers/{imsi} | Get Port Mapping entries for a subscriber.



## CreatePortMapping

> PortMapping CreatePortMapping(ctx).CreatePortMappingRequest(createPortMappingRequest).Execute()

Create Port Mapping.



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
    createPortMappingRequest := *openapiclient.NewCreatePortMappingRequest(*openapiclient.NewPortMappingDestination("Imsi_example", float32(123))) // CreatePortMappingRequest | Port mapping to be created.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortMappingApi.CreatePortMapping(context.Background()).CreatePortMappingRequest(createPortMappingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortMappingApi.CreatePortMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePortMapping`: PortMapping
    fmt.Fprintf(os.Stdout, "Response from `PortMappingApi.CreatePortMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePortMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPortMappingRequest** | [**CreatePortMappingRequest**](CreatePortMappingRequest.md) | Port mapping to be created. | 

### Return type

[**PortMapping**](PortMapping.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePortMapping

> DeletePortMapping(ctx, ipAddress, port).Execute()

Delete PortMapping.



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
    ipAddress := "ipAddress_example" // string | IP address of the target port mapping entry
    port := "port_example" // string | Port of the target port mapping entry

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortMappingApi.DeletePortMapping(context.Background(), ipAddress, port).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortMappingApi.DeletePortMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ipAddress** | **string** | IP address of the target port mapping entry | 
**port** | **string** | Port of the target port mapping entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePortMappingRequest struct via the builder pattern


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


## ListPortMappings

> []PortMapping ListPortMappings(ctx).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()

List Port Mapping Entries.



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
    limit := int32(56) // int32 | Maximum number of results per response page. (optional)
    lastEvaluatedKey := "lastEvaluatedKey_example" // string | The last Port Mapping ID retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next group onward. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortMappingApi.ListPortMappings(context.Background()).Limit(limit).LastEvaluatedKey(lastEvaluatedKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortMappingApi.ListPortMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPortMappings`: []PortMapping
    fmt.Fprintf(os.Stdout, "Response from `PortMappingApi.ListPortMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPortMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of results per response page. | 
 **lastEvaluatedKey** | **string** | The last Port Mapping ID retrieved on the current page. By specifying this parameter, you can continue to retrieve the list from the next group onward. | 

### Return type

[**[]PortMapping**](PortMapping.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPortMappingsForSubscriber

> PortMapping ListPortMappingsForSubscriber(ctx, imsi).Execute()

Get Port Mapping entries for a subscriber.



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
    imsi := "imsi_example" // string | Target subscriber IMSI.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortMappingApi.ListPortMappingsForSubscriber(context.Background(), imsi).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortMappingApi.ListPortMappingsForSubscriber``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPortMappingsForSubscriber`: PortMapping
    fmt.Fprintf(os.Stdout, "Response from `PortMappingApi.ListPortMappingsForSubscriber`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imsi** | **string** | Target subscriber IMSI. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPortMappingsForSubscriberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PortMapping**](PortMapping.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

