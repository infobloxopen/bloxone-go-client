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

type HardwareFilterAPI interface {
	/*
			HardwareFilterCreate Create the hardware filter.

			Use this method to create a __HardwareFilter__ object.
		The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiHardwareFilterCreateRequest
	*/
	HardwareFilterCreate(ctx context.Context) ApiHardwareFilterCreateRequest

	// HardwareFilterCreateExecute executes the request
	//  @return IpamsvcCreateHardwareFilterResponse
	HardwareFilterCreateExecute(r ApiHardwareFilterCreateRequest) (*IpamsvcCreateHardwareFilterResponse, *http.Response, error)
	/*
			HardwareFilterDelete Move the hardware filter to the recycle bin.

			Use this method to move a __HardwareFilter__ object to the recycle bin.
		The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiHardwareFilterDeleteRequest
	*/
	HardwareFilterDelete(ctx context.Context, id string) ApiHardwareFilterDeleteRequest

	// HardwareFilterDeleteExecute executes the request
	HardwareFilterDeleteExecute(r ApiHardwareFilterDeleteRequest) (*http.Response, error)
	/*
			HardwareFilterList Retrieve hardware filters.

			Use this method to retrieve __HardwareFilter__ objects.
		The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiHardwareFilterListRequest
	*/
	HardwareFilterList(ctx context.Context) ApiHardwareFilterListRequest

	// HardwareFilterListExecute executes the request
	//  @return IpamsvcListHardwareFilterResponse
	HardwareFilterListExecute(r ApiHardwareFilterListRequest) (*IpamsvcListHardwareFilterResponse, *http.Response, error)
	/*
			HardwareFilterRead Retrieve the hardware filter.

			Use this method to retrieve a __HardwareFilter__ object.
		The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiHardwareFilterReadRequest
	*/
	HardwareFilterRead(ctx context.Context, id string) ApiHardwareFilterReadRequest

	// HardwareFilterReadExecute executes the request
	//  @return IpamsvcReadHardwareFilterResponse
	HardwareFilterReadExecute(r ApiHardwareFilterReadRequest) (*IpamsvcReadHardwareFilterResponse, *http.Response, error)
	/*
			HardwareFilterUpdate Update the hardware filter.

			Use this method to update a __HardwareFilter__ object.
		The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiHardwareFilterUpdateRequest
	*/
	HardwareFilterUpdate(ctx context.Context, id string) ApiHardwareFilterUpdateRequest

	// HardwareFilterUpdateExecute executes the request
	//  @return IpamsvcUpdateHardwareFilterResponse
	HardwareFilterUpdateExecute(r ApiHardwareFilterUpdateRequest) (*IpamsvcUpdateHardwareFilterResponse, *http.Response, error)
}

// HardwareFilterAPIService HardwareFilterAPI service
type HardwareFilterAPIService internal.Service

type ApiHardwareFilterCreateRequest struct {
	ctx        context.Context
	ApiService HardwareFilterAPI
	body       *IpamsvcHardwareFilter
}

func (r ApiHardwareFilterCreateRequest) Body(body IpamsvcHardwareFilter) ApiHardwareFilterCreateRequest {
	r.body = &body
	return r
}

func (r ApiHardwareFilterCreateRequest) Execute() (*IpamsvcCreateHardwareFilterResponse, *http.Response, error) {
	return r.ApiService.HardwareFilterCreateExecute(r)
}

/*
HardwareFilterCreate Create the hardware filter.

Use this method to create a __HardwareFilter__ object.
The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiHardwareFilterCreateRequest
*/
func (a *HardwareFilterAPIService) HardwareFilterCreate(ctx context.Context) ApiHardwareFilterCreateRequest {
	return ApiHardwareFilterCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcCreateHardwareFilterResponse
func (a *HardwareFilterAPIService) HardwareFilterCreateExecute(r ApiHardwareFilterCreateRequest) (*IpamsvcCreateHardwareFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcCreateHardwareFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HardwareFilterAPIService.HardwareFilterCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/hardware_filter"

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
	for k, v := range a.Client.Cfg.DefaultTags {
		if _, ok := r.body.Tags[k]; !ok {
			r.body.Tags[k] = v
		}
	}
	// body params
	localVarPostBody = r.body
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

type ApiHardwareFilterDeleteRequest struct {
	ctx        context.Context
	ApiService HardwareFilterAPI
	id         string
}

func (r ApiHardwareFilterDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.HardwareFilterDeleteExecute(r)
}

/*
HardwareFilterDelete Move the hardware filter to the recycle bin.

Use this method to move a __HardwareFilter__ object to the recycle bin.
The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiHardwareFilterDeleteRequest
*/
func (a *HardwareFilterAPIService) HardwareFilterDelete(ctx context.Context, id string) ApiHardwareFilterDeleteRequest {
	return ApiHardwareFilterDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *HardwareFilterAPIService) HardwareFilterDeleteExecute(r ApiHardwareFilterDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HardwareFilterAPIService.HardwareFilterDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/hardware_filter/{id}"
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

type ApiHardwareFilterListRequest struct {
	ctx        context.Context
	ApiService HardwareFilterAPI
	filter     *string
	orderBy    *string
	fields     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	torderBy   *string
	tfilter    *string
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r ApiHardwareFilterListRequest) Filter(filter string) ApiHardwareFilterListRequest {
	r.filter = &filter
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r ApiHardwareFilterListRequest) OrderBy(orderBy string) ApiHardwareFilterListRequest {
	r.orderBy = &orderBy
	return r
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiHardwareFilterListRequest) Fields(fields string) ApiHardwareFilterListRequest {
	r.fields = &fields
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiHardwareFilterListRequest) Offset(offset int32) ApiHardwareFilterListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiHardwareFilterListRequest) Limit(limit int32) ApiHardwareFilterListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiHardwareFilterListRequest) PageToken(pageToken string) ApiHardwareFilterListRequest {
	r.pageToken = &pageToken
	return r
}

// This parameter is used for sorting by tags.
func (r ApiHardwareFilterListRequest) TorderBy(torderBy string) ApiHardwareFilterListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiHardwareFilterListRequest) Tfilter(tfilter string) ApiHardwareFilterListRequest {
	r.tfilter = &tfilter
	return r
}

func (r ApiHardwareFilterListRequest) Execute() (*IpamsvcListHardwareFilterResponse, *http.Response, error) {
	return r.ApiService.HardwareFilterListExecute(r)
}

/*
HardwareFilterList Retrieve hardware filters.

Use this method to retrieve __HardwareFilter__ objects.
The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiHardwareFilterListRequest
*/
func (a *HardwareFilterAPIService) HardwareFilterList(ctx context.Context) ApiHardwareFilterListRequest {
	return ApiHardwareFilterListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return IpamsvcListHardwareFilterResponse
func (a *HardwareFilterAPIService) HardwareFilterListExecute(r ApiHardwareFilterListRequest) (*IpamsvcListHardwareFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcListHardwareFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HardwareFilterAPIService.HardwareFilterList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/hardware_filter"

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

type ApiHardwareFilterReadRequest struct {
	ctx        context.Context
	ApiService HardwareFilterAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiHardwareFilterReadRequest) Fields(fields string) ApiHardwareFilterReadRequest {
	r.fields = &fields
	return r
}

func (r ApiHardwareFilterReadRequest) Execute() (*IpamsvcReadHardwareFilterResponse, *http.Response, error) {
	return r.ApiService.HardwareFilterReadExecute(r)
}

/*
HardwareFilterRead Retrieve the hardware filter.

Use this method to retrieve a __HardwareFilter__ object.
The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiHardwareFilterReadRequest
*/
func (a *HardwareFilterAPIService) HardwareFilterRead(ctx context.Context, id string) ApiHardwareFilterReadRequest {
	return ApiHardwareFilterReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcReadHardwareFilterResponse
func (a *HardwareFilterAPIService) HardwareFilterReadExecute(r ApiHardwareFilterReadRequest) (*IpamsvcReadHardwareFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcReadHardwareFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HardwareFilterAPIService.HardwareFilterRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/hardware_filter/{id}"
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

type ApiHardwareFilterUpdateRequest struct {
	ctx        context.Context
	ApiService HardwareFilterAPI
	id         string
	body       *IpamsvcHardwareFilter
}

func (r ApiHardwareFilterUpdateRequest) Body(body IpamsvcHardwareFilter) ApiHardwareFilterUpdateRequest {
	r.body = &body
	return r
}

func (r ApiHardwareFilterUpdateRequest) Execute() (*IpamsvcUpdateHardwareFilterResponse, *http.Response, error) {
	return r.ApiService.HardwareFilterUpdateExecute(r)
}

/*
HardwareFilterUpdate Update the hardware filter.

Use this method to update a __HardwareFilter__ object.
The __HardwareFilter__ object applies options to clients with matching hardware addresses. It must be configured in the _filters_ list of a scope to be effective.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiHardwareFilterUpdateRequest
*/
func (a *HardwareFilterAPIService) HardwareFilterUpdate(ctx context.Context, id string) ApiHardwareFilterUpdateRequest {
	return ApiHardwareFilterUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return IpamsvcUpdateHardwareFilterResponse
func (a *HardwareFilterAPIService) HardwareFilterUpdateExecute(r ApiHardwareFilterUpdateRequest) (*IpamsvcUpdateHardwareFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *IpamsvcUpdateHardwareFilterResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "HardwareFilterAPIService.HardwareFilterUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/hardware_filter/{id}"
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
	for k, v := range a.Client.Cfg.DefaultTags {
		if _, ok := r.body.Tags[k]; !ok {
			r.body.Tags[k] = v
		}
	}
	// body params
	localVarPostBody = r.body
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
