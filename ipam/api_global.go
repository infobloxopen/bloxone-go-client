/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type GlobalAPI interface {

	/*
		GlobalRead Retrieve the global configuration.

		Use this method to retrieve the __Global__ configuration object.
	The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGlobalReadRequest
	*/
	GlobalRead(ctx context.Context) ApiGlobalReadRequest

	// GlobalReadExecute executes the request
	//  @return IpamsvcReadGlobalResponse
	GlobalReadExecute(r ApiGlobalReadRequest) (*IpamsvcReadGlobalResponse, *http.Response, error)

	/*
		GlobalRead2 Retrieve the global configuration.

		Use this method to retrieve the __Global__ configuration object.
	The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id An application specific resource identity of a resource
		@return ApiGlobalRead2Request
	*/
	GlobalRead2(ctx context.Context, id string) ApiGlobalRead2Request

	// GlobalRead2Execute executes the request
	//  @return IpamsvcReadGlobalResponse
	GlobalRead2Execute(r ApiGlobalRead2Request) (*IpamsvcReadGlobalResponse, *http.Response, error)

	/*
		GlobalUpdate Update the global configuration.

		Use this method to update the __Global__ configuration object.
	The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiGlobalUpdateRequest
	*/
	GlobalUpdate(ctx context.Context) ApiGlobalUpdateRequest

	// GlobalUpdateExecute executes the request
	//  @return IpamsvcUpdateGlobalResponse
	GlobalUpdateExecute(r ApiGlobalUpdateRequest) (*IpamsvcUpdateGlobalResponse, *http.Response, error)

	/*
		GlobalUpdate2 Update the global configuration.

		Use this method to update the __Global__ configuration object.
	The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id An application specific resource identity of a resource
		@return ApiGlobalUpdate2Request
	*/
	GlobalUpdate2(ctx context.Context, id string) ApiGlobalUpdate2Request

	// GlobalUpdate2Execute executes the request
	//  @return IpamsvcUpdateGlobalResponse
	GlobalUpdate2Execute(r ApiGlobalUpdate2Request) (*IpamsvcUpdateGlobalResponse, *http.Response, error)
}

// GlobalAPIService GlobalAPI service
type GlobalAPIService internal.Service

type ApiGlobalReadRequest struct {
	ctx        context.Context
	ApiService GlobalAPI
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiGlobalReadRequest) Fields(fields string) ApiGlobalReadRequest {
	r.fields = &fields
	return r
}

func (r ApiGlobalReadRequest) Execute() (*IpamsvcReadGlobalResponse, *http.Response, error) {
	return r.ApiService.GlobalReadExecute(r)
}

/*
GlobalRead Retrieve the global configuration.

Use this method to retrieve the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalReadRequest
*/
func (a *GlobalAPIService) GlobalRead(ctx context.Context) ApiGlobalReadRequest {
	return ApiGlobalReadRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcReadGlobalResponse
func (a *GlobalAPIService) GlobalReadExecute(r ApiGlobalReadRequest) (*IpamsvcReadGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcReadGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.GlobalRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGlobalRead2Request struct {
	ctx        context.Context
	ApiService GlobalAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiGlobalRead2Request) Fields(fields string) ApiGlobalRead2Request {
	r.fields = &fields
	return r
}

func (r ApiGlobalRead2Request) Execute() (*IpamsvcReadGlobalResponse, *http.Response, error) {
	return r.ApiService.GlobalRead2Execute(r)
}

/*
GlobalRead2 Retrieve the global configuration.

Use this method to retrieve the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiGlobalRead2Request
*/
func (a *GlobalAPIService) GlobalRead2(ctx context.Context, id string) ApiGlobalRead2Request {
	return ApiGlobalRead2Request{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcReadGlobalResponse
func (a *GlobalAPIService) GlobalRead2Execute(r ApiGlobalRead2Request) (*IpamsvcReadGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcReadGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.GlobalRead2")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGlobalUpdateRequest struct {
	ctx        context.Context
	ApiService GlobalAPI
	body       *IpamsvcGlobal
}

func (r ApiGlobalUpdateRequest) Body(body IpamsvcGlobal) ApiGlobalUpdateRequest {
	r.body = &body
	return r
}

func (r ApiGlobalUpdateRequest) Execute() (*IpamsvcUpdateGlobalResponse, *http.Response, error) {
	return r.ApiService.GlobalUpdateExecute(r)
}

/*
GlobalUpdate Update the global configuration.

Use this method to update the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGlobalUpdateRequest
*/
func (a *GlobalAPIService) GlobalUpdate(ctx context.Context) ApiGlobalUpdateRequest {
	return ApiGlobalUpdateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcUpdateGlobalResponse
func (a *GlobalAPIService) GlobalUpdateExecute(r ApiGlobalUpdateRequest) (*IpamsvcUpdateGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcUpdateGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.GlobalUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGlobalUpdate2Request struct {
	ctx        context.Context
	ApiService GlobalAPI
	id         string
	body       *IpamsvcGlobal
}

func (r ApiGlobalUpdate2Request) Body(body IpamsvcGlobal) ApiGlobalUpdate2Request {
	r.body = &body
	return r
}

func (r ApiGlobalUpdate2Request) Execute() (*IpamsvcUpdateGlobalResponse, *http.Response, error) {
	return r.ApiService.GlobalUpdate2Execute(r)
}

/*
GlobalUpdate2 Update the global configuration.

Use this method to update the __Global__ configuration object.
The service operates on single __Global__ (_dhcp/global_) object that represents parent configuration settings for inheritance.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiGlobalUpdate2Request
*/
func (a *GlobalAPIService) GlobalUpdate2(ctx context.Context, id string) ApiGlobalUpdate2Request {
	return ApiGlobalUpdate2Request{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcUpdateGlobalResponse
func (a *GlobalAPIService) GlobalUpdate2Execute(r ApiGlobalUpdate2Request) (*IpamsvcUpdateGlobalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcUpdateGlobalResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "GlobalAPIService.GlobalUpdate2")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/global/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := internal.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(internal.ContextAPIKeys).(map[string]internal.APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.Client.PrepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := internal.NewGenericOpenAPIErrorWithBody(err.Error(), localVarBody)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
