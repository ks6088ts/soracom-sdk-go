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
)


// LogApiService LogApi service
type LogApiService service

type ApiGetLogsRequest struct {
	ctx context.Context
	ApiService *LogApiService
	resourceType *string
	resourceId *string
	service *string
	from *int32
	to *int32
	limit *int32
	lastEvaluatedKey *string
}

// ログ取得対象のリソースのタイプ
func (r ApiGetLogsRequest) ResourceType(resourceType string) ApiGetLogsRequest {
	r.resourceType = &resourceType
	return r
}

// ログ取得対象のリソースの ID
func (r ApiGetLogsRequest) ResourceId(resourceId string) ApiGetLogsRequest {
	r.resourceId = &resourceId
	return r
}

// ログエントリをフィルタするためのサービス名
func (r ApiGetLogsRequest) Service(service string) ApiGetLogsRequest {
	r.service = &service
	return r
}

// ログ取得対象の期間の始まり (unixtime)
func (r ApiGetLogsRequest) From(from int32) ApiGetLogsRequest {
	r.from = &from
	return r
}

// ログ取得対象の期間の終わり (unixtime)
func (r ApiGetLogsRequest) To(to int32) ApiGetLogsRequest {
	r.to = &to
	return r
}

// 取得するログエントリ数の上限
func (r ApiGetLogsRequest) Limit(limit int32) ApiGetLogsRequest {
	r.limit = &limit
	return r
}

// 前ページで取得した最後のログエントリのタイムスタンプ。このパラメータを指定することで次のログエントリ以降のリストを取得できる。
func (r ApiGetLogsRequest) LastEvaluatedKey(lastEvaluatedKey string) ApiGetLogsRequest {
	r.lastEvaluatedKey = &lastEvaluatedKey
	return r
}

func (r ApiGetLogsRequest) Execute() ([]LogEntry, *http.Response, error) {
	return r.ApiService.GetLogsExecute(r)
}

/*
GetLogs Get Logs.

条件に合うログエントリのリストを返す。ログエントリの総数が 1 ページに収まらない場合は、次のページにアクセスするための URL を`Link`ヘッダに含めて返す。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLogsRequest
*/
func (a *LogApiService) GetLogs(ctx context.Context) ApiGetLogsRequest {
	return ApiGetLogsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []LogEntry
func (a *LogApiService) GetLogsExecute(r ApiGetLogsRequest) ([]LogEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LogEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogApiService.GetLogs")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/logs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.resourceType != nil {
		localVarQueryParams.Add("resource_type", parameterToString(*r.resourceType, ""))
	}
	if r.resourceId != nil {
		localVarQueryParams.Add("resource_id", parameterToString(*r.resourceId, ""))
	}
	if r.service != nil {
		localVarQueryParams.Add("service", parameterToString(*r.service, ""))
	}
	if r.from != nil {
		localVarQueryParams.Add("from", parameterToString(*r.from, ""))
	}
	if r.to != nil {
		localVarQueryParams.Add("to", parameterToString(*r.to, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.lastEvaluatedKey != nil {
		localVarQueryParams.Add("last_evaluated_key", parameterToString(*r.lastEvaluatedKey, ""))
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