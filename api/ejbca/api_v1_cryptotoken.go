/*
Copyright 2024 Keyfactor

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// V1CryptotokenApiService V1CryptotokenApi service
type V1CryptotokenApiService service

type ApiActivate1Request struct {
	ctx                              context.Context
	ApiService                       *V1CryptotokenApiService
	cryptotokenName                  string
	cryptoTokenActivationRestRequest *CryptoTokenActivationRestRequest
}

// activation code
func (r ApiActivate1Request) CryptoTokenActivationRestRequest(cryptoTokenActivationRestRequest CryptoTokenActivationRestRequest) ApiActivate1Request {
	r.cryptoTokenActivationRestRequest = &cryptoTokenActivationRestRequest
	return r
}

func (r ApiActivate1Request) Execute() (*http.Response, error) {
	return r.ApiService.Activate1Execute(r)
}

/*
Activate1 Activate a Crypto Token

Activates Crypto Token given name and activation code

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cryptotokenName Name of the token to activate
	@return ApiActivate1Request
*/
func (a *V1CryptotokenApiService) Activate1(ctx context.Context, cryptotokenName string) ApiActivate1Request {
	return ApiActivate1Request{
		ApiService:      a,
		ctx:             ctx,
		cryptotokenName: cryptotokenName,
	}
}

// Execute executes the request
func (a *V1CryptotokenApiService) Activate1Execute(r ApiActivate1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/cryptotoken/{cryptotoken_name}/activate"
	localVarPath = strings.Replace(localVarPath, "{"+"cryptotoken_name"+"}", url.PathEscape(parameterValueToString(r.cryptotokenName, "cryptotokenName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.cryptoTokenActivationRestRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiDeactivate1Request struct {
	ctx             context.Context
	ApiService      *V1CryptotokenApiService
	cryptotokenName string
}

func (r ApiDeactivate1Request) Execute() (*http.Response, error) {
	return r.ApiService.Deactivate1Execute(r)
}

/*
Deactivate1 Deactivate a Crypto Token

Deactivates Crypto Token given name

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cryptotokenName Name of the token to deactivate
	@return ApiDeactivate1Request
*/
func (a *V1CryptotokenApiService) Deactivate1(ctx context.Context, cryptotokenName string) ApiDeactivate1Request {
	return ApiDeactivate1Request{
		ApiService:      a,
		ctx:             ctx,
		cryptotokenName: cryptotokenName,
	}
}

// Execute executes the request
func (a *V1CryptotokenApiService) Deactivate1Execute(r ApiDeactivate1Request) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/cryptotoken/{cryptotoken_name}/deactivate"
	localVarPath = strings.Replace(localVarPath, "{"+"cryptotoken_name"+"}", url.PathEscape(parameterValueToString(r.cryptotokenName, "cryptotokenName")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiGenerateKeysRequest struct {
	ctx                                 context.Context
	ApiService                          *V1CryptotokenApiService
	cryptotokenName                     string
	cryptoTokenKeyGenerationRestRequest *CryptoTokenKeyGenerationRestRequest
}

func (r ApiGenerateKeysRequest) CryptoTokenKeyGenerationRestRequest(cryptoTokenKeyGenerationRestRequest CryptoTokenKeyGenerationRestRequest) ApiGenerateKeysRequest {
	r.cryptoTokenKeyGenerationRestRequest = &cryptoTokenKeyGenerationRestRequest
	return r
}

func (r ApiGenerateKeysRequest) Execute() (*http.Response, error) {
	return r.ApiService.GenerateKeysExecute(r)
}

/*
GenerateKeys Generate keys

Generates a key pair given crypto token name, key pair alias, key algorithm and key specification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cryptotokenName Name of the token to generate keys for
	@return ApiGenerateKeysRequest
*/
func (a *V1CryptotokenApiService) GenerateKeys(ctx context.Context, cryptotokenName string) ApiGenerateKeysRequest {
	return ApiGenerateKeysRequest{
		ApiService:      a,
		ctx:             ctx,
		cryptotokenName: cryptotokenName,
	}
}

// Execute executes the request
func (a *V1CryptotokenApiService) GenerateKeysExecute(r ApiGenerateKeysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/cryptotoken/{cryptotoken_name}/generatekeys"
	localVarPath = strings.Replace(localVarPath, "{"+"cryptotoken_name"+"}", url.PathEscape(parameterValueToString(r.cryptotokenName, "cryptotokenName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.cryptoTokenKeyGenerationRestRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiRemoveKeysRequest struct {
	ctx             context.Context
	ApiService      *V1CryptotokenApiService
	cryptotokenName string
	keyPairAlias    string
}

func (r ApiRemoveKeysRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveKeysExecute(r)
}

/*
RemoveKeys Remove keys

Remove a key pair given crypto token name and key pair alias to be removed.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param cryptotokenName Name of the token to remove keys for.
	@param keyPairAlias Alias for the key to be removed from the crypto token.
	@return ApiRemoveKeysRequest
*/
func (a *V1CryptotokenApiService) RemoveKeys(ctx context.Context, cryptotokenName string, keyPairAlias string) ApiRemoveKeysRequest {
	return ApiRemoveKeysRequest{
		ApiService:      a,
		ctx:             ctx,
		cryptotokenName: cryptotokenName,
		keyPairAlias:    keyPairAlias,
	}
}

// Execute executes the request
func (a *V1CryptotokenApiService) RemoveKeysExecute(r ApiRemoveKeysRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/cryptotoken/{cryptotoken_name}/{key_pair_alias}/removekeys"
	localVarPath = strings.Replace(localVarPath, "{"+"cryptotoken_name"+"}", url.PathEscape(parameterValueToString(r.cryptotokenName, "cryptotokenName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"key_pair_alias"+"}", url.PathEscape(parameterValueToString(r.keyPairAlias, "keyPairAlias")), -1)

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiStatus5Request struct {
	ctx        context.Context
	ApiService *V1CryptotokenApiService
}

func (r ApiStatus5Request) Execute() (*RestResourceStatusRestResponse, *http.Response, error) {
	return r.ApiService.Status5Execute(r)
}

/*
Status5 Get the status of this REST Resource

Returns status, API version and EJBCA version.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiStatus5Request
*/
func (a *V1CryptotokenApiService) Status5(ctx context.Context) ApiStatus5Request {
	return ApiStatus5Request{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return RestResourceStatusRestResponse
func (a *V1CryptotokenApiService) Status5Execute(r ApiStatus5Request) (*RestResourceStatusRestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RestResourceStatusRestResponse
	)

	localBasePath := "/ejbca/ejbca-rest-api"

	localVarPath := localBasePath + "/v1/cryptotoken/status"

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
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
