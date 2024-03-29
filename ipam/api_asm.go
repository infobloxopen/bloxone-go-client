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

type AsmAPI interface {
	/*
			AsmCreate Update subnet and ranges for Automated Scope Management.

			Use this method to update the subnet and range for Automated Scope Management.
		The __ASM__ object generates and returns the suggestions from the ASM suggestion engine and allows for updating the subnet and range.
		This method attempts to expand the scope by expanding a range or adding a new range and, if necessary, expanding the subnet.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiAsmCreateRequest
	*/
	AsmCreate(ctx context.Context) ApiAsmCreateRequest

	// AsmCreateExecute executes the request
	//  @return IpamsvcCreateASMResponse
	AsmCreateExecute(r ApiAsmCreateRequest) (*IpamsvcCreateASMResponse, *http.Response, error)
	/*
			AsmList Retrieve suggested updates for Automated Scope Management.

			Use this method to retrieve __ASM__ objects for Automated Scope Management.
		The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiAsmListRequest
	*/
	AsmList(ctx context.Context) ApiAsmListRequest

	// AsmListExecute executes the request
	//  @return IpamsvcListASMResponse
	AsmListExecute(r ApiAsmListRequest) (*IpamsvcListASMResponse, *http.Response, error)
	/*
			AsmRead Retrieve the suggested update for Automated Scope Management.

			Use this method to retrieve an __ASM__ object for Automated Scope Management.
		The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiAsmReadRequest
	*/
	AsmRead(ctx context.Context, id string) ApiAsmReadRequest

	// AsmReadExecute executes the request
	//  @return IpamsvcReadASMResponse
	AsmReadExecute(r ApiAsmReadRequest) (*IpamsvcReadASMResponse, *http.Response, error)
}

// AsmAPIService AsmAPI service
type AsmAPIService internal.Service

type ApiAsmCreateRequest struct {
	ctx        context.Context
	ApiService AsmAPI
	body       *IpamsvcASM
}

func (r ApiAsmCreateRequest) Body(body IpamsvcASM) ApiAsmCreateRequest {
	r.body = &body
	return r
}

func (r ApiAsmCreateRequest) Execute() (*IpamsvcCreateASMResponse, *http.Response, error) {
	return r.ApiService.AsmCreateExecute(r)
}

/*
AsmCreate Update subnet and ranges for Automated Scope Management.

Use this method to update the subnet and range for Automated Scope Management.
The __ASM__ object generates and returns the suggestions from the ASM suggestion engine and allows for updating the subnet and range.
This method attempts to expand the scope by expanding a range or adding a new range and, if necessary, expanding the subnet.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAsmCreateRequest
*/
func (a *AsmAPIService) AsmCreate(ctx context.Context) ApiAsmCreateRequest {
	return ApiAsmCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcCreateASMResponse
func (a *AsmAPIService) AsmCreateExecute(r ApiAsmCreateRequest) (*IpamsvcCreateASMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcCreateASMResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AsmAPIService.AsmCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipam/asm"

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

type ApiAsmListRequest struct {
	ctx        context.Context
	ApiService AsmAPI
	fields     *string
	subnetId   *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiAsmListRequest) Fields(fields string) ApiAsmListRequest {
	r.fields = &fields
	return r
}

func (r ApiAsmListRequest) SubnetId(subnetId string) ApiAsmListRequest {
	r.subnetId = &subnetId
	return r
}

func (r ApiAsmListRequest) Execute() (*IpamsvcListASMResponse, *http.Response, error) {
	return r.ApiService.AsmListExecute(r)
}

/*
AsmList Retrieve suggested updates for Automated Scope Management.

Use this method to retrieve __ASM__ objects for Automated Scope Management.
The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiAsmListRequest
*/
func (a *AsmAPIService) AsmList(ctx context.Context) ApiAsmListRequest {
	return ApiAsmListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcListASMResponse
func (a *AsmAPIService) AsmListExecute(r ApiAsmListRequest) (*IpamsvcListASMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcListASMResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AsmAPIService.AsmList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipam/asm"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.subnetId != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "subnet_id", r.subnetId, "")
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

type ApiAsmReadRequest struct {
	ctx        context.Context
	ApiService AsmAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiAsmReadRequest) Fields(fields string) ApiAsmReadRequest {
	r.fields = &fields
	return r
}

func (r ApiAsmReadRequest) Execute() (*IpamsvcReadASMResponse, *http.Response, error) {
	return r.ApiService.AsmReadExecute(r)
}

/*
AsmRead Retrieve the suggested update for Automated Scope Management.

Use this method to retrieve an __ASM__ object for Automated Scope Management.
The __ASM__ object returns the suggested updates for the subnet from the ASM suggestion engine and allows for updating the subnet and range information.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiAsmReadRequest
*/
func (a *AsmAPIService) AsmRead(ctx context.Context, id string) ApiAsmReadRequest {
	return ApiAsmReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcReadASMResponse
func (a *AsmAPIService) AsmReadExecute(r ApiAsmReadRequest) (*IpamsvcReadASMResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcReadASMResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "AsmAPIService.AsmRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/ipam/asm/{id}"
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
