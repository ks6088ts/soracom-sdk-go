/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// DeviceObjectModelApiService DeviceObjectModelApi service
type DeviceObjectModelApiService service

type ApiCreateDeviceObjectModelRequest struct {
	ctx context.Context
	ApiService *DeviceObjectModelApiService
	deviceObjectModel *DeviceObjectModel
}

// Device object model definition in either XML or JSON format.
func (r ApiCreateDeviceObjectModelRequest) DeviceObjectModel(deviceObjectModel DeviceObjectModel) ApiCreateDeviceObjectModelRequest {
	r.deviceObjectModel = &deviceObjectModel
	return r
}

func (r ApiCreateDeviceObjectModelRequest) Execute() (*DeviceObjectModel, *http.Response, error) {
	return r.ApiService.CreateDeviceObjectModelExecute(r)
}

/*
CreateDeviceObjectModel 新しいデバイスオブジェクトモデルを作成します。

新しいデバイスオブジェクトモデルを作成します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDeviceObjectModelRequest
*/
func (a *DeviceObjectModelApiService) CreateDeviceObjectModel(ctx context.Context) ApiCreateDeviceObjectModelRequest {
	return ApiCreateDeviceObjectModelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DeviceObjectModel
func (a *DeviceObjectModelApiService) CreateDeviceObjectModelExecute(r ApiCreateDeviceObjectModelRequest) (*DeviceObjectModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceObjectModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceObjectModelApiService.CreateDeviceObjectModel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device_object_models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceObjectModel == nil {
		return localVarReturnValue, nil, reportError("deviceObjectModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deviceObjectModel
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-API-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteDeviceObjectModelRequest struct {
	ctx context.Context
	ApiService *DeviceObjectModelApiService
	modelId string
}

func (r ApiDeleteDeviceObjectModelRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDeviceObjectModelExecute(r)
}

/*
DeleteDeviceObjectModel 対象のデバイスオブジェクトモデルを削除します。

対象のデバイスオブジェクトモデルを削除します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param modelId 対象のデバイスオブジェクトモデルの ID
 @return ApiDeleteDeviceObjectModelRequest
*/
func (a *DeviceObjectModelApiService) DeleteDeviceObjectModel(ctx context.Context, modelId string) ApiDeleteDeviceObjectModelRequest {
	return ApiDeleteDeviceObjectModelRequest{
		ApiService: a,
		ctx: ctx,
		modelId: modelId,
	}
}

// Execute executes the request
func (a *DeviceObjectModelApiService) DeleteDeviceObjectModelExecute(r ApiDeleteDeviceObjectModelRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceObjectModelApiService.DeleteDeviceObjectModel")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device_object_models/{model_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"model_id"+"}", url.PathEscape(parameterToString(r.modelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-API-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetDeviceObjectModelRequest struct {
	ctx context.Context
	ApiService *DeviceObjectModelApiService
	modelId string
}

func (r ApiGetDeviceObjectModelRequest) Execute() (*DeviceObjectModel, *http.Response, error) {
	return r.ApiService.GetDeviceObjectModelExecute(r)
}

/*
GetDeviceObjectModel デバイスオブジェクトモデルを取得します。

デバイスオブジェクトモデルを取得します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param modelId 対象のデバイスオブジェクトモデルの ID
 @return ApiGetDeviceObjectModelRequest
*/
func (a *DeviceObjectModelApiService) GetDeviceObjectModel(ctx context.Context, modelId string) ApiGetDeviceObjectModelRequest {
	return ApiGetDeviceObjectModelRequest{
		ApiService: a,
		ctx: ctx,
		modelId: modelId,
	}
}

// Execute executes the request
//  @return DeviceObjectModel
func (a *DeviceObjectModelApiService) GetDeviceObjectModelExecute(r ApiGetDeviceObjectModelRequest) (*DeviceObjectModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceObjectModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceObjectModelApiService.GetDeviceObjectModel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device_object_models/{model_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"model_id"+"}", url.PathEscape(parameterToString(r.modelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-API-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListDeviceObjectModelsRequest struct {
	ctx context.Context
	ApiService *DeviceObjectModelApiService
	lastEvaluatedKey *string
	limit *int32
}

// 現ページで取得した最後のデバイスオブジェクトモデルの ID。このパラメータを指定することで次のデバイスオブジェクトモデル以降のリストを取得できる。
func (r ApiListDeviceObjectModelsRequest) LastEvaluatedKey(lastEvaluatedKey string) ApiListDeviceObjectModelsRequest {
	r.lastEvaluatedKey = &lastEvaluatedKey
	return r
}

// 取得する結果の上限数
func (r ApiListDeviceObjectModelsRequest) Limit(limit int32) ApiListDeviceObjectModelsRequest {
	r.limit = &limit
	return r
}

func (r ApiListDeviceObjectModelsRequest) Execute() ([]DeviceObjectModel, *http.Response, error) {
	return r.ApiService.ListDeviceObjectModelsExecute(r)
}

/*
ListDeviceObjectModels デバイスオブジェクトモデルのリストを返します。

デバイスオブジェクトモデルのリストを返します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListDeviceObjectModelsRequest
*/
func (a *DeviceObjectModelApiService) ListDeviceObjectModels(ctx context.Context) ApiListDeviceObjectModelsRequest {
	return ApiListDeviceObjectModelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []DeviceObjectModel
func (a *DeviceObjectModelApiService) ListDeviceObjectModelsExecute(r ApiListDeviceObjectModelsRequest) ([]DeviceObjectModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []DeviceObjectModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceObjectModelApiService.ListDeviceObjectModels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device_object_models"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.lastEvaluatedKey != nil {
		localVarQueryParams.Add("last_evaluated_key", parameterToString(*r.lastEvaluatedKey, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-API-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSetDeviceObjectModelScopeRequest struct {
	ctx context.Context
	ApiService *DeviceObjectModelApiService
	modelId string
	setDeviceObjectModelScopeRequest *SetDeviceObjectModelScopeRequest
}

// 対象のデバイスオブジェクトモデルにセットする Scope の値
func (r ApiSetDeviceObjectModelScopeRequest) SetDeviceObjectModelScopeRequest(setDeviceObjectModelScopeRequest SetDeviceObjectModelScopeRequest) ApiSetDeviceObjectModelScopeRequest {
	r.setDeviceObjectModelScopeRequest = &setDeviceObjectModelScopeRequest
	return r
}

func (r ApiSetDeviceObjectModelScopeRequest) Execute() (*DeviceObjectModel, *http.Response, error) {
	return r.ApiService.SetDeviceObjectModelScopeExecute(r)
}

/*
SetDeviceObjectModelScope 対象のデバイスオブジェクトモデルが適用される Scope をセットします。

対象のデバイスオブジェクトモデルが適用される Scope をセットします。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param modelId 対象のデバイスオブジェクトモデルの ID
 @return ApiSetDeviceObjectModelScopeRequest
*/
func (a *DeviceObjectModelApiService) SetDeviceObjectModelScope(ctx context.Context, modelId string) ApiSetDeviceObjectModelScopeRequest {
	return ApiSetDeviceObjectModelScopeRequest{
		ApiService: a,
		ctx: ctx,
		modelId: modelId,
	}
}

// Execute executes the request
//  @return DeviceObjectModel
func (a *DeviceObjectModelApiService) SetDeviceObjectModelScopeExecute(r ApiSetDeviceObjectModelScopeRequest) (*DeviceObjectModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceObjectModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceObjectModelApiService.SetDeviceObjectModelScope")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device_object_models/{model_id}/set_scope"
	localVarPath = strings.Replace(localVarPath, "{"+"model_id"+"}", url.PathEscape(parameterToString(r.modelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.setDeviceObjectModelScopeRequest == nil {
		return localVarReturnValue, nil, reportError("setDeviceObjectModelScopeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.setDeviceObjectModelScopeRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-API-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateDeviceObjectModelRequest struct {
	ctx context.Context
	ApiService *DeviceObjectModelApiService
	modelId string
	deviceObjectModel *DeviceObjectModel
}

// XML か JSON により表現されたデバイスオブジェクトモデルの表記。
func (r ApiUpdateDeviceObjectModelRequest) DeviceObjectModel(deviceObjectModel DeviceObjectModel) ApiUpdateDeviceObjectModelRequest {
	r.deviceObjectModel = &deviceObjectModel
	return r
}

func (r ApiUpdateDeviceObjectModelRequest) Execute() (*DeviceObjectModel, *http.Response, error) {
	return r.ApiService.UpdateDeviceObjectModelExecute(r)
}

/*
UpdateDeviceObjectModel デバイスオブジェクトモデルを更新します。

デバイスオブジェクトモデルを更新します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param modelId 対象のデバイスオブジェクトモデルの ID
 @return ApiUpdateDeviceObjectModelRequest
*/
func (a *DeviceObjectModelApiService) UpdateDeviceObjectModel(ctx context.Context, modelId string) ApiUpdateDeviceObjectModelRequest {
	return ApiUpdateDeviceObjectModelRequest{
		ApiService: a,
		ctx: ctx,
		modelId: modelId,
	}
}

// Execute executes the request
//  @return DeviceObjectModel
func (a *DeviceObjectModelApiService) UpdateDeviceObjectModelExecute(r ApiUpdateDeviceObjectModelRequest) (*DeviceObjectModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DeviceObjectModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DeviceObjectModelApiService.UpdateDeviceObjectModel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/device_object_models/{model_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"model_id"+"}", url.PathEscape(parameterToString(r.modelId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.deviceObjectModel == nil {
		return localVarReturnValue, nil, reportError("deviceObjectModel is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/xml"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.deviceObjectModel
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-API-Key"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_token"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Soracom-Token"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
