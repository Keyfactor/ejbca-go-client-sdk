/*
EJBCA REST Interface

API reference documentation.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ejbca

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// V2EndentityApiService V2EndentityApi service
type V2EndentityApiService service

type ApiGetAuthorizedEndEntityProfilesRequest struct {
	ctx context.Context
	ApiService *V2EndentityApiService
}

func (r ApiGetAuthorizedEndEntityProfilesRequest) Execute() (*AuthorizedEEPsRestResponse, *http.Response, error) {
	return r.ApiService.GetAuthorizedEndEntityProfilesExecute(r)
}

/*
GetAuthorizedEndEntityProfiles List of authorized end entity profiles for the current admin.

Returns list of all authorized end entity profiles for the current admin token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAuthorizedEndEntityProfilesRequest
*/
func (a *V2EndentityApiService) GetAuthorizedEndEntityProfiles(ctx context.Context) ApiGetAuthorizedEndEntityProfilesRequest {
	return ApiGetAuthorizedEndEntityProfilesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuthorizedEEPsRestResponse
func (a *V2EndentityApiService) GetAuthorizedEndEntityProfilesExecute(r ApiGetAuthorizedEndEntityProfilesRequest) (*AuthorizedEEPsRestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizedEEPsRestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V2EndentityApiService.GetAuthorizedEndEntityProfiles")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/endentity/profiles/authorized"

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

type ApiProfileRequest struct {
	ctx context.Context
	ApiService *V2EndentityApiService
	endentityProfileName string
}

func (r ApiProfileRequest) Execute() (*EndEntityProfileResponse, *http.Response, error) {
	return r.ApiService.ProfileExecute(r)
}

/*
Profile Get End Entity Profile content

Returns End Entity Profile configurations: List of available CAs, list of available Certificate Profiles.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param endentityProfileName
 @return ApiProfileRequest
*/
func (a *V2EndentityApiService) Profile(ctx context.Context, endentityProfileName string) ApiProfileRequest {
	return ApiProfileRequest{
		ApiService: a,
		ctx: ctx,
		endentityProfileName: endentityProfileName,
	}
}

// Execute executes the request
//  @return EndEntityProfileResponse
func (a *V2EndentityApiService) ProfileExecute(r ApiProfileRequest) (*EndEntityProfileResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EndEntityProfileResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V2EndentityApiService.Profile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/endentity/profile/{endentity_profile_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"endentity_profile_name"+"}", url.PathEscape(parameterValueToString(r.endentityProfileName, "endentityProfileName")), -1)

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

type ApiSortedSearchRequest struct {
	ctx context.Context
	ApiService *V2EndentityApiService
	searchEndEntitiesRestRequestV2 *SearchEndEntitiesRestRequestV2
}

// Maximum number of results and collection of search criterias.
func (r ApiSortedSearchRequest) SearchEndEntitiesRestRequestV2(searchEndEntitiesRestRequestV2 SearchEndEntitiesRestRequestV2) ApiSortedSearchRequest {
	r.searchEndEntitiesRestRequestV2 = &searchEndEntitiesRestRequestV2
	return r
}

func (r ApiSortedSearchRequest) Execute() (*SearchEndEntitiesRestResponse, *http.Response, error) {
	return r.ApiService.SortedSearchExecute(r)
}

/*
SortedSearch Searches and sorts for end entity conforming given criteria.

Insert as many search criteria as needed and optionally a sorting criteria. A reference about allowed values for criteria could be found below, under SearchEndEntityCriteriaRestRequestV2 model.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSortedSearchRequest
*/
func (a *V2EndentityApiService) SortedSearch(ctx context.Context) ApiSortedSearchRequest {
	return ApiSortedSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return SearchEndEntitiesRestResponse
func (a *V2EndentityApiService) SortedSearchExecute(r ApiSortedSearchRequest) (*SearchEndEntitiesRestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SearchEndEntitiesRestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V2EndentityApiService.SortedSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/endentity/search"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.searchEndEntitiesRestRequestV2
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

type ApiStatus7Request struct {
	ctx context.Context
	ApiService *V2EndentityApiService
}

func (r ApiStatus7Request) Execute() (*RestResourceStatusRestResponse, *http.Response, error) {
	return r.ApiService.Status7Execute(r)
}

/*
Status7 Get the status of this REST Resource

Returns status, API version and EJBCA version.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStatus7Request
*/
func (a *V2EndentityApiService) Status7(ctx context.Context) ApiStatus7Request {
	return ApiStatus7Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RestResourceStatusRestResponse
func (a *V2EndentityApiService) Status7Execute(r ApiStatus7Request) (*RestResourceStatusRestResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RestResourceStatusRestResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "V2EndentityApiService.Status7")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/endentity/status"

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
