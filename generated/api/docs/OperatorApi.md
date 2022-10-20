# \OperatorApi

All URIs are relative to *https://api.soracom.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCoverageType**](OperatorApi.md#AddCoverageType) | **Post** /operators/{operator_id}/coverage_type/{coverage_type} | Add coverage type
[**AddOperatorContract**](OperatorApi.md#AddOperatorContract) | **Post** /operators/{operator_id}/contracts | Add Operator Contract
[**CreateCompanyInformation**](OperatorApi.md#CreateCompanyInformation) | **Post** /operators/{operator_id}/company_information | Create company information
[**CreateIndividualInformation**](OperatorApi.md#CreateIndividualInformation) | **Post** /operators/{operator_id}/individual_information | Create individual information
[**CreateOperator**](OperatorApi.md#CreateOperator) | **Post** /operators | Create Operator
[**DeleteOperatorAuthKey**](OperatorApi.md#DeleteOperatorAuthKey) | **Delete** /operators/{operator_id}/auth_keys/{auth_key_id} | Delete Operator AuthKey
[**DeleteOperatorContract**](OperatorApi.md#DeleteOperatorContract) | **Delete** /operators/{operator_id}/contracts/{contract_name} | Delete Operator Contract
[**EnableMFA**](OperatorApi.md#EnableMFA) | **Post** /operators/{operator_id}/mfa | Enable Operator&#39;s MFA.
[**GenerateAuthToken**](OperatorApi.md#GenerateAuthToken) | **Post** /operators/{operator_id}/token | Generate Authentication Token
[**GenerateOperatorAuthKey**](OperatorApi.md#GenerateOperatorAuthKey) | **Post** /operators/{operator_id}/auth_keys | Generate Operator AuthKey
[**GenerateSupportToken**](OperatorApi.md#GenerateSupportToken) | **Post** /operators/{operator_id}/support/token | Generate Token for Support Console
[**GetCompanyInformation**](OperatorApi.md#GetCompanyInformation) | **Get** /operators/{operator_id}/company_information | Get company information
[**GetIndividualInformation**](OperatorApi.md#GetIndividualInformation) | **Get** /operators/{operator_id}/individual_information | Get individual information
[**GetMFAStatus**](OperatorApi.md#GetMFAStatus) | **Get** /operators/{operator_id}/mfa | Get Operator&#39;s MFA Status.
[**GetOperator**](OperatorApi.md#GetOperator) | **Get** /operators/{operator_id} | Get Operator
[**IssueMFARevokingToken**](OperatorApi.md#IssueMFARevokingToken) | **Post** /operators/mfa_revoke_token/issue | Issue Operator&#39;s MFA Revoke Token.
[**ListOperatorAuthKeys**](OperatorApi.md#ListOperatorAuthKeys) | **Get** /operators/{operator_id}/auth_keys | List Operator AuthKeys
[**RevokeMFA**](OperatorApi.md#RevokeMFA) | **Delete** /operators/{operator_id}/mfa | Revoke Operator&#39;s MFA.
[**RevokeOperatorAuthTokens**](OperatorApi.md#RevokeOperatorAuthTokens) | **Delete** /operators/{operator_id}/tokens | ルートユーザーが発行した全ての API キーと API トークンを無効化します。
[**UpdateCompanyInformation**](OperatorApi.md#UpdateCompanyInformation) | **Put** /operators/{operator_id}/company_information | Update company information
[**UpdateIndividualInformation**](OperatorApi.md#UpdateIndividualInformation) | **Put** /operators/{operator_id}/individual_information | Update individual information
[**UpdateOperatorPassword**](OperatorApi.md#UpdateOperatorPassword) | **Post** /operators/{operator_id}/password | Update Operator Password
[**VerifyMFA**](OperatorApi.md#VerifyMFA) | **Post** /operators/{operator_id}/mfa/verify | Verify Operator&#39;s MFA OTP Code.
[**VerifyMFARevokingToken**](OperatorApi.md#VerifyMFARevokingToken) | **Post** /operators/mfa_revoke_token/verify | Verify Operator&#39;s MFA revoke token.
[**VerifyOperator**](OperatorApi.md#VerifyOperator) | **Post** /operators/verify | Verify Operator



## AddCoverageType

> AddCoverageType(ctx, operatorId, coverageType).Execute()

Add coverage type



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
    coverageType := "coverageType_example" // string | coverage_type

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.AddCoverageType(context.Background(), operatorId, coverageType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.AddCoverageType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**coverageType** | **string** | coverage_type | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCoverageTypeRequest struct via the builder pattern


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


## AddOperatorContract

> ContractUpdatedResponse AddOperatorContract(ctx, operatorId).ContractUpdatingRequest(contractUpdatingRequest).Execute()

Add Operator Contract



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
    contractUpdatingRequest := *openapiclient.NewContractUpdatingRequest() // ContractUpdatingRequest | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.AddOperatorContract(context.Background(), operatorId).ContractUpdatingRequest(contractUpdatingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.AddOperatorContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddOperatorContract`: ContractUpdatedResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.AddOperatorContract`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOperatorContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contractUpdatingRequest** | [**ContractUpdatingRequest**](ContractUpdatingRequest.md) | model | 

### Return type

[**ContractUpdatedResponse**](ContractUpdatedResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCompanyInformation

> CreateCompanyInformation(ctx, operatorId).CompanyInformationModel(companyInformationModel).Execute()

Create company information



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
    companyInformationModel := *openapiclient.NewCompanyInformationModel("CompanyName_example", "ContactPersonName_example", "CountryCode_example", "Department_example", "PhoneNumber_example", "ZipCode_example") // CompanyInformationModel | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.CreateCompanyInformation(context.Background(), operatorId).CompanyInformationModel(companyInformationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.CreateCompanyInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompanyInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyInformationModel** | [**CompanyInformationModel**](CompanyInformationModel.md) | model | 

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


## CreateIndividualInformation

> CreateIndividualInformation(ctx, operatorId).IndividualInformationModel(individualInformationModel).Execute()

Create individual information



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
    individualInformationModel := *openapiclient.NewIndividualInformationModel("CountryCode_example", "FullName_example", "PhoneNumber_example", "ZipCode_example") // IndividualInformationModel | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.CreateIndividualInformation(context.Background(), operatorId).IndividualInformationModel(individualInformationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.CreateIndividualInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIndividualInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **individualInformationModel** | [**IndividualInformationModel**](IndividualInformationModel.md) | model | 

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


## CreateOperator

> CreateOperator(ctx).RegisterOperatorsRequest(registerOperatorsRequest).Execute()

Create Operator



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
    registerOperatorsRequest := *openapiclient.NewRegisterOperatorsRequest("Email_example", "Password_example") // RegisterOperatorsRequest | email, password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.CreateOperator(context.Background()).RegisterOperatorsRequest(registerOperatorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.CreateOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerOperatorsRequest** | [**RegisterOperatorsRequest**](RegisterOperatorsRequest.md) | email, password | 

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


## DeleteOperatorAuthKey

> DeleteOperatorAuthKey(ctx, operatorId, authKeyId).Execute()

Delete Operator AuthKey



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
    authKeyId := "authKeyId_example" // string | auth_key_id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.DeleteOperatorAuthKey(context.Background(), operatorId, authKeyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.DeleteOperatorAuthKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**authKeyId** | **string** | auth_key_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOperatorAuthKeyRequest struct via the builder pattern


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


## DeleteOperatorContract

> DeleteOperatorContract(ctx, operatorId, contractName).Execute()

Delete Operator Contract



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
    contractName := "contractName_example" // string | contract_name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.DeleteOperatorContract(context.Background(), operatorId, contractName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.DeleteOperatorContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 
**contractName** | **string** | contract_name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOperatorContractRequest struct via the builder pattern


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


## EnableMFA

> EnableMFAOTPResponse EnableMFA(ctx, operatorId).Execute()

Enable Operator's MFA.



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
    resp, r, err := apiClient.OperatorApi.EnableMFA(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.EnableMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnableMFA`: EnableMFAOTPResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.EnableMFA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableMFARequest struct via the builder pattern


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


## GenerateAuthToken

> GenerateTokenResponse GenerateAuthToken(ctx, operatorId).GenerateTokenRequest(generateTokenRequest).Execute()

Generate Authentication Token



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
    generateTokenRequest := *openapiclient.NewGenerateTokenRequest() // GenerateTokenRequest | token timeout seconds

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.GenerateAuthToken(context.Background(), operatorId).GenerateTokenRequest(generateTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GenerateAuthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateAuthToken`: GenerateTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GenerateAuthToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateAuthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateTokenRequest** | [**GenerateTokenRequest**](GenerateTokenRequest.md) | token timeout seconds | 

### Return type

[**GenerateTokenResponse**](GenerateTokenResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateOperatorAuthKey

> GenerateOperatorsAuthKeyResponse GenerateOperatorAuthKey(ctx, operatorId).Execute()

Generate Operator AuthKey



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
    resp, r, err := apiClient.OperatorApi.GenerateOperatorAuthKey(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GenerateOperatorAuthKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateOperatorAuthKey`: GenerateOperatorsAuthKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GenerateOperatorAuthKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateOperatorAuthKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenerateOperatorsAuthKeyResponse**](GenerateOperatorsAuthKeyResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateSupportToken

> SupportTokenResponse GenerateSupportToken(ctx, operatorId).Execute()

Generate Token for Support Console



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
    resp, r, err := apiClient.OperatorApi.GenerateSupportToken(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GenerateSupportToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GenerateSupportToken`: SupportTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GenerateSupportToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSupportTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SupportTokenResponse**](SupportTokenResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCompanyInformation

> CompanyInformationModel GetCompanyInformation(ctx, operatorId).Execute()

Get company information



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
    resp, r, err := apiClient.OperatorApi.GetCompanyInformation(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetCompanyInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCompanyInformation`: CompanyInformationModel
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetCompanyInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCompanyInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CompanyInformationModel**](CompanyInformationModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIndividualInformation

> IndividualInformationModel GetIndividualInformation(ctx, operatorId).Execute()

Get individual information



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
    resp, r, err := apiClient.OperatorApi.GetIndividualInformation(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetIndividualInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIndividualInformation`: IndividualInformationModel
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetIndividualInformation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIndividualInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IndividualInformationModel**](IndividualInformationModel.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMFAStatus

> MFAStatusOfUseResponse GetMFAStatus(ctx, operatorId).Execute()

Get Operator's MFA Status.



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
    resp, r, err := apiClient.OperatorApi.GetMFAStatus(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetMFAStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMFAStatus`: MFAStatusOfUseResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetMFAStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMFAStatusRequest struct via the builder pattern


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


## GetOperator

> GetOperatorResponse GetOperator(ctx, operatorId).Execute()

Get Operator



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
    resp, r, err := apiClient.OperatorApi.GetOperator(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.GetOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOperator`: GetOperatorResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.GetOperator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetOperatorResponse**](GetOperatorResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IssueMFARevokingToken

> IssueMFARevokingToken(ctx).MFAIssueRevokingTokenRequest(mFAIssueRevokingTokenRequest).Execute()

Issue Operator's MFA Revoke Token.



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
    mFAIssueRevokingTokenRequest := *openapiclient.NewMFAIssueRevokingTokenRequest() // MFAIssueRevokingTokenRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.IssueMFARevokingToken(context.Background()).MFAIssueRevokingTokenRequest(mFAIssueRevokingTokenRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.IssueMFARevokingToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIssueMFARevokingTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFAIssueRevokingTokenRequest** | [**MFAIssueRevokingTokenRequest**](MFAIssueRevokingTokenRequest.md) | request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOperatorAuthKeys

> []AuthKeyResponse ListOperatorAuthKeys(ctx, operatorId).Execute()

List Operator AuthKeys



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
    resp, r, err := apiClient.OperatorApi.ListOperatorAuthKeys(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.ListOperatorAuthKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOperatorAuthKeys`: []AuthKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.ListOperatorAuthKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOperatorAuthKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]AuthKeyResponse**](AuthKeyResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeMFA

> RevokeMFA(ctx, operatorId).Execute()

Revoke Operator's MFA.



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
    resp, r, err := apiClient.OperatorApi.RevokeMFA(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.RevokeMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeMFARequest struct via the builder pattern


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


## RevokeOperatorAuthTokens

> RevokeOperatorAuthTokens(ctx, operatorId).Execute()

ルートユーザーが発行した全ての API キーと API トークンを無効化します。



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
    resp, r, err := apiClient.OperatorApi.RevokeOperatorAuthTokens(context.Background(), operatorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.RevokeOperatorAuthTokens``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRevokeOperatorAuthTokensRequest struct via the builder pattern


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


## UpdateCompanyInformation

> UpdateCompanyInformation(ctx, operatorId).CompanyInformationModel(companyInformationModel).Execute()

Update company information



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
    companyInformationModel := *openapiclient.NewCompanyInformationModel("CompanyName_example", "ContactPersonName_example", "CountryCode_example", "Department_example", "PhoneNumber_example", "ZipCode_example") // CompanyInformationModel | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.UpdateCompanyInformation(context.Background(), operatorId).CompanyInformationModel(companyInformationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.UpdateCompanyInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCompanyInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **companyInformationModel** | [**CompanyInformationModel**](CompanyInformationModel.md) | model | 

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


## UpdateIndividualInformation

> UpdateIndividualInformation(ctx, operatorId).IndividualInformationModel(individualInformationModel).Execute()

Update individual information



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
    individualInformationModel := *openapiclient.NewIndividualInformationModel("CountryCode_example", "FullName_example", "PhoneNumber_example", "ZipCode_example") // IndividualInformationModel | model

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.UpdateIndividualInformation(context.Background(), operatorId).IndividualInformationModel(individualInformationModel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.UpdateIndividualInformation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIndividualInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **individualInformationModel** | [**IndividualInformationModel**](IndividualInformationModel.md) | model | 

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


## UpdateOperatorPassword

> UpdateOperatorPassword(ctx, operatorId).UpdatePasswordRequest(updatePasswordRequest).Execute()

Update Operator Password



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
    updatePasswordRequest := *openapiclient.NewUpdatePasswordRequest("CurrentPassword_example", "NewPassword_example") // UpdatePasswordRequest | current password, new password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.UpdateOperatorPassword(context.Background(), operatorId).UpdatePasswordRequest(updatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.UpdateOperatorPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOperatorPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePasswordRequest** | [**UpdatePasswordRequest**](UpdatePasswordRequest.md) | current password, new password | 

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


## VerifyMFA

> OperatorMFAVerifyingResponse VerifyMFA(ctx, operatorId).MFAAuthenticationRequest(mFAAuthenticationRequest).Execute()

Verify Operator's MFA OTP Code.



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
    mFAAuthenticationRequest := *openapiclient.NewMFAAuthenticationRequest() // MFAAuthenticationRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.VerifyMFA(context.Background(), operatorId).MFAAuthenticationRequest(mFAAuthenticationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.VerifyMFA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VerifyMFA`: OperatorMFAVerifyingResponse
    fmt.Fprintf(os.Stdout, "Response from `OperatorApi.VerifyMFA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operatorId** | **string** | operator_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMFARequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mFAAuthenticationRequest** | [**MFAAuthenticationRequest**](MFAAuthenticationRequest.md) | request | 

### Return type

[**OperatorMFAVerifyingResponse**](OperatorMFAVerifyingResponse.md)

### Authorization

[api_key](../README.md#api_key), [api_token](../README.md#api_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyMFARevokingToken

> VerifyMFARevokingToken(ctx).MFARevokingTokenVerifyRequest(mFARevokingTokenVerifyRequest).Execute()

Verify Operator's MFA revoke token.



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
    mFARevokingTokenVerifyRequest := *openapiclient.NewMFARevokingTokenVerifyRequest() // MFARevokingTokenVerifyRequest | request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.VerifyMFARevokingToken(context.Background()).MFARevokingTokenVerifyRequest(mFARevokingTokenVerifyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.VerifyMFARevokingToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyMFARevokingTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mFARevokingTokenVerifyRequest** | [**MFARevokingTokenVerifyRequest**](MFARevokingTokenVerifyRequest.md) | request | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyOperator

> VerifyOperator(ctx).VerifyOperatorsRequest(verifyOperatorsRequest).Execute()

Verify Operator



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
    verifyOperatorsRequest := *openapiclient.NewVerifyOperatorsRequest("Token_example") // VerifyOperatorsRequest | token

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OperatorApi.VerifyOperator(context.Background()).VerifyOperatorsRequest(verifyOperatorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OperatorApi.VerifyOperator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyOperatorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyOperatorsRequest** | [**VerifyOperatorsRequest**](VerifyOperatorsRequest.md) | token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

