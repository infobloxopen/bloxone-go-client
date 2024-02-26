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


type ServerAPI interface {
	/*
	ServerCreate Create the DHCP configuration profile.

	Use this method to create a __Server__ object.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiServerCreateRequest
	*/
	ServerCreate(ctx context.Context) ApiServerCreateRequest

	// ServerCreateExecute executes the request
	//  @return IpamsvcCreateServerResponse
	ServerCreateExecute(r ApiServerCreateRequest) (*IpamsvcCreateServerResponse, *http.Response, error)
	/*
	ServerDelete Move the DHCP configuration profile to the recycle bin.

	Use this method to move a __Server__ object to the recycle bin.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiServerDeleteRequest
	*/
	ServerDelete(ctx context.Context, id string) ApiServerDeleteRequest

	// ServerDeleteExecute executes the request
	ServerDeleteExecute(r ApiServerDeleteRequest) (*http.Response, error)
	/*
	ServerList Retrieve DHCP configuration profiles.

	Use this method to retrieve __Server__ objects.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiServerListRequest
	*/
	ServerList(ctx context.Context) ApiServerListRequest

	// ServerListExecute executes the request
	//  @return IpamsvcListServerResponse
	ServerListExecute(r ApiServerListRequest) (*IpamsvcListServerResponse, *http.Response, error)
	/*
	ServerRead Retrieve the DHCP configuration profile.

	Use this method to retrieve a __Server__ object.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiServerReadRequest
	*/
	ServerRead(ctx context.Context, id string) ApiServerReadRequest

	// ServerReadExecute executes the request
	//  @return IpamsvcReadServerResponse
	ServerReadExecute(r ApiServerReadRequest) (*IpamsvcReadServerResponse, *http.Response, error)
	/*
	ServerUpdate Update the DHCP configuration profile.

	Use this method to update a __Server__ object.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiServerUpdateRequest
	*/
	ServerUpdate(ctx context.Context, id string) ApiServerUpdateRequest

	// ServerUpdateExecute executes the request
	//  @return IpamsvcUpdateServerResponse
	ServerUpdateExecute(r ApiServerUpdateRequest) (*IpamsvcUpdateServerResponse, *http.Response, error)
}

// ServerAPIService ServerAPI service
type ServerAPIService internal.Service

type ApiServerCreateRequest struct {
	ctx context.Context
	ApiService ServerAPI
	body *IpamsvcServer
	inherit *string
}

func (r ApiServerCreateRequest) Body(body IpamsvcServer) ApiServerCreateRequest {
	r.body = &body
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiServerCreateRequest) Inherit(inherit string) ApiServerCreateRequest {
	r.inherit = &inherit
	return r
}

func (r ApiServerCreateRequest) Execute() (*IpamsvcCreateServerResponse, *http.Response, error) {
	return r.ApiService.ServerCreateExecute(r)
}

/*
ServerCreate Create the DHCP configuration profile.

Use this method to create a __Server__ object.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerCreateRequest
*/
func (a *ServerAPIService) ServerCreate(ctx context.Context) ApiServerCreateRequest {
	return ApiServerCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IpamsvcCreateServerResponse
func (a *ServerAPIService) ServerCreateExecute(r ApiServerCreateRequest) (*IpamsvcCreateServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcCreateServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ServerAPIService.ServerCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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
    if r.body.Tags == nil {
    	r.body.Tags = make(map[string]interface{})
    }
    for k, v := range a.Client.Cfg.DefaultTags {
        if _, ok := r.body.Tags[k]; !ok {
            r.body.Tags[k] = v
        }
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

type ApiServerDeleteRequest struct {
	ctx context.Context
	ApiService ServerAPI
	id string
}

func (r ApiServerDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ServerDeleteExecute(r)
}

/*
ServerDelete Move the DHCP configuration profile to the recycle bin.

Use this method to move a __Server__ object to the recycle bin.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiServerDeleteRequest
*/
func (a *ServerAPIService) ServerDelete(ctx context.Context, id string) ApiServerDeleteRequest {
	return ApiServerDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ServerAPIService) ServerDeleteExecute(r ApiServerDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ServerAPIService.ServerDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/server/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := internal.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
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
		newErr := internal.NewGenericOpenAPIErrorWithBody(localVarHTTPResponse.Status, localVarBody)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiServerListRequest struct {
	ctx context.Context
	ApiService ServerAPI
	filter *string
	orderBy *string
	fields *string
	offset *int32
	limit *int32
	pageToken *string
	torderBy *string
	tfilter *string
	inherit *string
}

//   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |        
func (r ApiServerListRequest) Filter(filter string) ApiServerListRequest {
	r.filter = &filter
	return r
}

//   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.        
func (r ApiServerListRequest) OrderBy(orderBy string) ApiServerListRequest {
	r.orderBy = &orderBy
	return r
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiServerListRequest) Fields(fields string) ApiServerListRequest {
	r.fields = &fields
	return r
}

//   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.         
func (r ApiServerListRequest) Offset(offset int32) ApiServerListRequest {
	r.offset = &offset
	return r
}

//   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.         
func (r ApiServerListRequest) Limit(limit int32) ApiServerListRequest {
	r.limit = &limit
	return r
}

//   The service-defined string used to identify a page of resources. A null value indicates the first page.         
func (r ApiServerListRequest) PageToken(pageToken string) ApiServerListRequest {
	r.pageToken = &pageToken
	return r
}

// This parameter is used for sorting by tags.
func (r ApiServerListRequest) TorderBy(torderBy string) ApiServerListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiServerListRequest) Tfilter(tfilter string) ApiServerListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiServerListRequest) Inherit(inherit string) ApiServerListRequest {
	r.inherit = &inherit
	return r
}

func (r ApiServerListRequest) Execute() (*IpamsvcListServerResponse, *http.Response, error) {
	return r.ApiService.ServerListExecute(r)
}

/*
ServerList Retrieve DHCP configuration profiles.

Use this method to retrieve __Server__ objects.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiServerListRequest
*/
func (a *ServerAPIService) ServerList(ctx context.Context) ApiServerListRequest {
	return ApiServerListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IpamsvcListServerResponse
func (a *ServerAPIService) ServerListExecute(r ApiServerListRequest) (*IpamsvcListServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcListServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ServerAPIService.ServerList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/server"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
	}
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.offset != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_offset", r.offset, "")
	}
	if r.limit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_limit", r.limit, "")
	}
	if r.pageToken != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_page_token", r.pageToken, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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

type ApiServerReadRequest struct {
	ctx context.Context
	ApiService ServerAPI
	id string
	fields *string
	inherit *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiServerReadRequest) Fields(fields string) ApiServerReadRequest {
	r.fields = &fields
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiServerReadRequest) Inherit(inherit string) ApiServerReadRequest {
	r.inherit = &inherit
	return r
}

func (r ApiServerReadRequest) Execute() (*IpamsvcReadServerResponse, *http.Response, error) {
	return r.ApiService.ServerReadExecute(r)
}

/*
ServerRead Retrieve the DHCP configuration profile.

Use this method to retrieve a __Server__ object.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiServerReadRequest
*/
func (a *ServerAPIService) ServerRead(ctx context.Context, id string) ApiServerReadRequest {
	return ApiServerReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IpamsvcReadServerResponse
func (a *ServerAPIService) ServerReadExecute(r ApiServerReadRequest) (*IpamsvcReadServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcReadServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ServerAPIService.ServerRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/server/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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

type ApiServerUpdateRequest struct {
	ctx context.Context
	ApiService ServerAPI
	id string
	body *IpamsvcServer
	inherit *string
}

func (r ApiServerUpdateRequest) Body(body IpamsvcServer) ApiServerUpdateRequest {
	r.body = &body
	return r
}

// This parameter is used for getting inheritance_sources.  Allowed values: * _none_, * _partial_, * _full_.  Defaults to _none
func (r ApiServerUpdateRequest) Inherit(inherit string) ApiServerUpdateRequest {
	r.inherit = &inherit
	return r
}

func (r ApiServerUpdateRequest) Execute() (*IpamsvcUpdateServerResponse, *http.Response, error) {
	return r.ApiService.ServerUpdateExecute(r)
}

/*
ServerUpdate Update the DHCP configuration profile.

Use this method to update a __Server__ object.
A __Server__ (DHCP Config Profile) is a named configuration profile that can be shared for specified list of hosts.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiServerUpdateRequest
*/
func (a *ServerAPIService) ServerUpdate(ctx context.Context, id string) ApiServerUpdateRequest {
	return ApiServerUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IpamsvcUpdateServerResponse
func (a *ServerAPIService) ServerUpdateExecute(r ApiServerUpdateRequest) (*IpamsvcUpdateServerResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcUpdateServerResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ServerAPIService.ServerUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/server/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(internal.ParameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, internal.ReportError("body is required and must be specified")
	}

	if r.inherit != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_inherit", r.inherit, "")
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
    if r.body.Tags == nil {
    	r.body.Tags = make(map[string]interface{})
    }
    for k, v := range a.Client.Cfg.DefaultTags {
        if _, ok := r.body.Tags[k]; !ok {
            r.body.Tags[k] = v
        }
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
