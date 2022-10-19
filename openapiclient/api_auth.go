/*
SORACOM API

SORACOM API v1

API version: VERSION_PLACEHOLDER
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapiclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// AuthApiService AuthApi service
type AuthApiService service

type ApiAuthRequest struct {
	ctx context.Context
	ApiService *AuthApiService
	authRequest *AuthRequest
}

// 認証リクエスト
func (r ApiAuthRequest) AuthRequest(authRequest AuthRequest) ApiAuthRequest {
	r.authRequest = &authRequest
	return r
}

func (r ApiAuthRequest) Execute() (*AuthResponse, *http.Response, error) {
	return r.ApiService.AuthExecute(r)
}

/*
Auth API アクセスの認証を行い、API キーと API トークンを発行します。

API アクセスの認証を行います。以下の情報をリクエストに含めてください。

- ルートユーザーで認証する場合: email と password
- ルートユーザーまたは SAM ユーザーの認証キーで認証する場合: authKeyId と authKey
- SAM ユーザー認証の場合: operatorId, userName, password

認証が成功した場合は、API キーと API トークンが返されます。以降の認証が必要なリクエスト時にはこの API キーと API トークンをヘッダーに付与してリクエストを行ってください。


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuthRequest
*/
func (a *AuthApiService) Auth(ctx context.Context) ApiAuthRequest {
	return ApiAuthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthResponse
func (a *AuthApiService) AuthExecute(r ApiAuthRequest) (*AuthResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.Auth")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.authRequest == nil {
		return localVarReturnValue, nil, reportError("authRequest is required and must be specified")
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
	localVarPostBody = r.authRequest
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

type ApiIssuePasswordResetTokenRequest struct {
	ctx context.Context
	ApiService *AuthApiService
	issuePasswordResetTokenRequest *IssuePasswordResetTokenRequest
}

// email address
func (r ApiIssuePasswordResetTokenRequest) IssuePasswordResetTokenRequest(issuePasswordResetTokenRequest IssuePasswordResetTokenRequest) ApiIssuePasswordResetTokenRequest {
	r.issuePasswordResetTokenRequest = &issuePasswordResetTokenRequest
	return r
}

func (r ApiIssuePasswordResetTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.IssuePasswordResetTokenExecute(r)
}

/*
IssuePasswordResetToken パスワードリセット用のトークンを発行します。

パスワードをリセットするためのワンタイムトークンを生成して、オペレーターのメールアドレスに送付します。メールに記載されているトークンを取得し、/v1/auth/password_reset_token/verify API を呼び出して新しいパスワードを設定します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuePasswordResetTokenRequest
*/
func (a *AuthApiService) IssuePasswordResetToken(ctx context.Context) ApiIssuePasswordResetTokenRequest {
	return ApiIssuePasswordResetTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) IssuePasswordResetTokenExecute(r ApiIssuePasswordResetTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.IssuePasswordResetToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/password_reset_token/issue"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.issuePasswordResetTokenRequest == nil {
		return nil, reportError("issuePasswordResetTokenRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.issuePasswordResetTokenRequest
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

type ApiLogoutRequest struct {
	ctx context.Context
	ApiService *AuthApiService
}

func (r ApiLogoutRequest) Execute() (*http.Response, error) {
	return r.ApiService.LogoutExecute(r)
}

/*
Logout SORACOM API にアクセスするための API キーと API トークンを無効化します。

この API を呼び出すときに指定された API キーと API トークンを無効化します。無効化したあとは、有効期間に関わらず API キーと API トークンを用いて SORACOM API を呼び出すことはできません。この API は、権限設定に関わらず呼び出すことができます。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLogoutRequest
*/
func (a *AuthApiService) Logout(ctx context.Context) ApiLogoutRequest {
	return ApiLogoutRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) LogoutExecute(r ApiLogoutRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.Logout")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/logout"

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

type ApiVerifyPasswordResetTokenRequest struct {
	ctx context.Context
	ApiService *AuthApiService
	verifyPasswordResetTokenRequest *VerifyPasswordResetTokenRequest
}

// token, password
func (r ApiVerifyPasswordResetTokenRequest) VerifyPasswordResetTokenRequest(verifyPasswordResetTokenRequest VerifyPasswordResetTokenRequest) ApiVerifyPasswordResetTokenRequest {
	r.verifyPasswordResetTokenRequest = &verifyPasswordResetTokenRequest
	return r
}

func (r ApiVerifyPasswordResetTokenRequest) Execute() (*http.Response, error) {
	return r.ApiService.VerifyPasswordResetTokenExecute(r)
}

/*
VerifyPasswordResetToken パスワードリセット用のトークンを検証し、パスワードを更新します。

パスワードリセット用のトークンを検証し、検証に成功した場合オペレーターのパスワードを更新します。

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyPasswordResetTokenRequest
*/
func (a *AuthApiService) VerifyPasswordResetToken(ctx context.Context) ApiVerifyPasswordResetTokenRequest {
	return ApiVerifyPasswordResetTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *AuthApiService) VerifyPasswordResetTokenExecute(r ApiVerifyPasswordResetTokenRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.VerifyPasswordResetToken")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auth/password_reset_token/verify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.verifyPasswordResetTokenRequest == nil {
		return nil, reportError("verifyPasswordResetTokenRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.verifyPasswordResetTokenRequest
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