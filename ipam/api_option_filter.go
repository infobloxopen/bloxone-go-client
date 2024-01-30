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

type OptionFilterAPI interface {
	/*
		OptionFilterCreate Create the DHCP option filter.

		Use this method to create an __OptionFilter__ object.
	The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiOptionFilterCreateRequest
	*/
	OptionFilterCreate(ctx context.Context) ApiOptionFilterCreateRequest

	// OptionFilterCreateExecute executes the request
	//  @return IpamsvcCreateOptionFilterResponse
	OptionFilterCreateExecute(r ApiOptionFilterCreateRequest) (*IpamsvcCreateOptionFilterResponse, *http.Response, error)
	/*
		OptionFilterDelete Move the DHCP option filter to the recycle bin.

		Use this method to move an __OptionFilter__ object to the recycle bin.
	The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id An application specific resource identity of a resource
		@return ApiOptionFilterDeleteRequest
	*/
	OptionFilterDelete(ctx context.Context, id string) ApiOptionFilterDeleteRequest

	// OptionFilterDeleteExecute executes the request
	OptionFilterDeleteExecute(r ApiOptionFilterDeleteRequest) (*http.Response, error)
	/*
		OptionFilterList Retrieve DHCP option filters.

		Use this method to retrieve __OptionFilter__ objects.
	The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiOptionFilterListRequest
	*/
	OptionFilterList(ctx context.Context) ApiOptionFilterListRequest

	// OptionFilterListExecute executes the request
	//  @return IpamsvcListOptionFilterResponse
	OptionFilterListExecute(r ApiOptionFilterListRequest) (*IpamsvcListOptionFilterResponse, *http.Response, error)
	/*
		OptionFilterRead Retrieve the DHCP option filter.

		Use this method to retrieve an __OptionFilter__ object.
	The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id An application specific resource identity of a resource
		@return ApiOptionFilterReadRequest
	*/
	OptionFilterRead(ctx context.Context, id string) ApiOptionFilterReadRequest

	// OptionFilterReadExecute executes the request
	//  @return IpamsvcReadOptionFilterResponse
	OptionFilterReadExecute(r ApiOptionFilterReadRequest) (*IpamsvcReadOptionFilterResponse, *http.Response, error)
	/*
		OptionFilterUpdate Update the DHCP option filter.

		Use this method to update an __OptionFilter__ object.
	The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param id An application specific resource identity of a resource
		@return ApiOptionFilterUpdateRequest
	*/
	OptionFilterUpdate(ctx context.Context, id string) ApiOptionFilterUpdateRequest

	// OptionFilterUpdateExecute executes the request
	//  @return IpamsvcUpdateOptionFilterResponse
	OptionFilterUpdateExecute(r ApiOptionFilterUpdateRequest) (*IpamsvcUpdateOptionFilterResponse, *http.Response, error)
}

// OptionFilterAPIService OptionFilterAPI service
type OptionFilterAPIService internal.Service

type ApiOptionFilterCreateRequest struct {
	ctx        context.Context
	ApiService OptionFilterAPI
	body       *IpamsvcOptionFilter
}

func (r ApiOptionFilterCreateRequest) Body(body IpamsvcOptionFilter) ApiOptionFilterCreateRequest {
	r.body = &body
	return r
}

func (r ApiOptionFilterCreateRequest) Execute() (*IpamsvcCreateOptionFilterResponse, *http.Response, error) {
	return r.ApiService.OptionFilterCreateExecute(r)
}

/*
OptionFilterCreate Create the DHCP option filter.

Use this method to create an __OptionFilter__ object.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOptionFilterCreateRequest
*/
func (a *OptionFilterAPIService) OptionFilterCreate(ctx context.Context) ApiOptionFilterCreateRequest {
	return ApiOptionFilterCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcCreateOptionFilterResponse
func (a *OptionFilterAPIService) OptionFilterCreateExecute(r ApiOptionFilterCreateRequest) (*IpamsvcCreateOptionFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcCreateOptionFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionFilterAPIService.OptionFilterCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_filter"

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
	if r.body.Tags == nil {
		r.body.Tags = make(map[string]interface{})
	}
	for k, v := range a.Client.Cfg.GetDefaultTags() {
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

type ApiOptionFilterDeleteRequest struct {
	ctx        context.Context
	ApiService OptionFilterAPI
	id         string
}

func (r ApiOptionFilterDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.OptionFilterDeleteExecute(r)
}

/*
OptionFilterDelete Move the DHCP option filter to the recycle bin.

Use this method to move an __OptionFilter__ object to the recycle bin.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiOptionFilterDeleteRequest
*/
func (a *OptionFilterAPIService) OptionFilterDelete(ctx context.Context, id string) ApiOptionFilterDeleteRequest {
	return ApiOptionFilterDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *OptionFilterAPIService) OptionFilterDeleteExecute(r ApiOptionFilterDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionFilterAPIService.OptionFilterDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_filter/{id}"
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

type ApiOptionFilterListRequest struct {
	ctx        context.Context
	ApiService OptionFilterAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	torderBy   *string
	tfilter    *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiOptionFilterListRequest) Fields(fields string) ApiOptionFilterListRequest {
	r.fields = &fields
	return r
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r ApiOptionFilterListRequest) Filter(filter string) ApiOptionFilterListRequest {
	r.filter = &filter
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiOptionFilterListRequest) Offset(offset int32) ApiOptionFilterListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiOptionFilterListRequest) Limit(limit int32) ApiOptionFilterListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiOptionFilterListRequest) PageToken(pageToken string) ApiOptionFilterListRequest {
	r.pageToken = &pageToken
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r ApiOptionFilterListRequest) OrderBy(orderBy string) ApiOptionFilterListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for sorting by tags.
func (r ApiOptionFilterListRequest) TorderBy(torderBy string) ApiOptionFilterListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiOptionFilterListRequest) Tfilter(tfilter string) ApiOptionFilterListRequest {
	r.tfilter = &tfilter
	return r
}

func (r ApiOptionFilterListRequest) Execute() (*IpamsvcListOptionFilterResponse, *http.Response, error) {
	return r.ApiService.OptionFilterListExecute(r)
}

/*
OptionFilterList Retrieve DHCP option filters.

Use this method to retrieve __OptionFilter__ objects.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOptionFilterListRequest
*/
func (a *OptionFilterAPIService) OptionFilterList(ctx context.Context) ApiOptionFilterListRequest {
	return ApiOptionFilterListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcListOptionFilterResponse
func (a *OptionFilterAPIService) OptionFilterListExecute(r ApiOptionFilterListRequest) (*IpamsvcListOptionFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcListOptionFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionFilterAPIService.OptionFilterList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_filter"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_fields", r.fields, "")
	}
	if r.filter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_filter", r.filter, "")
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
	if r.orderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_order_by", r.orderBy, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
	}
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
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

type ApiOptionFilterReadRequest struct {
	ctx        context.Context
	ApiService OptionFilterAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiOptionFilterReadRequest) Fields(fields string) ApiOptionFilterReadRequest {
	r.fields = &fields
	return r
}

func (r ApiOptionFilterReadRequest) Execute() (*IpamsvcReadOptionFilterResponse, *http.Response, error) {
	return r.ApiService.OptionFilterReadExecute(r)
}

/*
OptionFilterRead Retrieve the DHCP option filter.

Use this method to retrieve an __OptionFilter__ object.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiOptionFilterReadRequest
*/
func (a *OptionFilterAPIService) OptionFilterRead(ctx context.Context, id string) ApiOptionFilterReadRequest {
	return ApiOptionFilterReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcReadOptionFilterResponse
func (a *OptionFilterAPIService) OptionFilterReadExecute(r ApiOptionFilterReadRequest) (*IpamsvcReadOptionFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcReadOptionFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionFilterAPIService.OptionFilterRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_filter/{id}"
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

type ApiOptionFilterUpdateRequest struct {
	ctx        context.Context
	ApiService OptionFilterAPI
	id         string
	body       *IpamsvcOptionFilter
}

func (r ApiOptionFilterUpdateRequest) Body(body IpamsvcOptionFilter) ApiOptionFilterUpdateRequest {
	r.body = &body
	return r
}

func (r ApiOptionFilterUpdateRequest) Execute() (*IpamsvcUpdateOptionFilterResponse, *http.Response, error) {
	return r.ApiService.OptionFilterUpdateExecute(r)
}

/*
OptionFilterUpdate Update the DHCP option filter.

Use this method to update an __OptionFilter__ object.
The __OptionFilter__ object applies options to DHCP clients with matching options. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiOptionFilterUpdateRequest
*/
func (a *OptionFilterAPIService) OptionFilterUpdate(ctx context.Context, id string) ApiOptionFilterUpdateRequest {
	return ApiOptionFilterUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcUpdateOptionFilterResponse
func (a *OptionFilterAPIService) OptionFilterUpdateExecute(r ApiOptionFilterUpdateRequest) (*IpamsvcUpdateOptionFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcUpdateOptionFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionFilterAPIService.OptionFilterUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_filter/{id}"
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
	if r.body.Tags == nil {
		r.body.Tags = make(map[string]interface{})
	}
	for k, v := range a.Client.Cfg.GetDefaultTags() {
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
