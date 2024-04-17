/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type ForwardZoneAPI interface {
	/*
			ForwardZoneCopy Copies the __ForwardZone__ object.

			Use this method to copy an __ForwardZone__ object to a different __View__.
		This object (_dns/forward_zone_) represents a forwarding zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiForwardZoneCopyRequest
	*/
	ForwardZoneCopy(ctx context.Context) ApiForwardZoneCopyRequest

	// ForwardZoneCopyExecute executes the request
	//  @return ConfigCopyForwardZoneResponse
	ForwardZoneCopyExecute(r ApiForwardZoneCopyRequest) (*ConfigCopyForwardZoneResponse, *http.Response, error)
	/*
			ForwardZoneCreate Create the ForwardZone object.

			Use this method to create a ForwardZone object.
		This object (_dns/forward_zone_) represents a forwarding zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiForwardZoneCreateRequest
	*/
	ForwardZoneCreate(ctx context.Context) ApiForwardZoneCreateRequest

	// ForwardZoneCreateExecute executes the request
	//  @return ConfigCreateForwardZoneResponse
	ForwardZoneCreateExecute(r ApiForwardZoneCreateRequest) (*ConfigCreateForwardZoneResponse, *http.Response, error)
	/*
			ForwardZoneDelete Move the Forward Zone object to Recyclebin.

			Use this method to move a Forward Zone object to Recyclebin.
		This object (_dns/forward_zone_) represents a forwarding zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiForwardZoneDeleteRequest
	*/
	ForwardZoneDelete(ctx context.Context, id string) ApiForwardZoneDeleteRequest

	// ForwardZoneDeleteExecute executes the request
	ForwardZoneDeleteExecute(r ApiForwardZoneDeleteRequest) (*http.Response, error)
	/*
			ForwardZoneList List Forward Zone objects.

			Use this method to list Forward Zone objects.
		This object (_dns/forward_zone_) represents a forwarding zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiForwardZoneListRequest
	*/
	ForwardZoneList(ctx context.Context) ApiForwardZoneListRequest

	// ForwardZoneListExecute executes the request
	//  @return ConfigListForwardZoneResponse
	ForwardZoneListExecute(r ApiForwardZoneListRequest) (*ConfigListForwardZoneResponse, *http.Response, error)
	/*
			ForwardZoneRead Read the Forward Zone object.

			Use this method to read a Forward Zone object.
		This object (_dns/forward_zone_) represents a forwarding zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiForwardZoneReadRequest
	*/
	ForwardZoneRead(ctx context.Context, id string) ApiForwardZoneReadRequest

	// ForwardZoneReadExecute executes the request
	//  @return ConfigReadForwardZoneResponse
	ForwardZoneReadExecute(r ApiForwardZoneReadRequest) (*ConfigReadForwardZoneResponse, *http.Response, error)
	/*
			ForwardZoneUpdate Update the Forward Zone object.

			Use this method to update a Forward Zone object.
		This object (_dns/forward_zone_) represents a forwarding zone.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiForwardZoneUpdateRequest
	*/
	ForwardZoneUpdate(ctx context.Context, id string) ApiForwardZoneUpdateRequest

	// ForwardZoneUpdateExecute executes the request
	//  @return ConfigUpdateForwardZoneResponse
	ForwardZoneUpdateExecute(r ApiForwardZoneUpdateRequest) (*ConfigUpdateForwardZoneResponse, *http.Response, error)
}

// ForwardZoneAPIService ForwardZoneAPI service
type ForwardZoneAPIService internal.Service

type ApiForwardZoneCopyRequest struct {
	ctx        context.Context
	ApiService ForwardZoneAPI
	body       *ConfigCopyForwardZone
}

func (r ApiForwardZoneCopyRequest) Body(body ConfigCopyForwardZone) ApiForwardZoneCopyRequest {
	r.body = &body
	return r
}

func (r ApiForwardZoneCopyRequest) Execute() (*ConfigCopyForwardZoneResponse, *http.Response, error) {
	return r.ApiService.ForwardZoneCopyExecute(r)
}

/*
ForwardZoneCopy Copies the __ForwardZone__ object.

Use this method to copy an __ForwardZone__ object to a different __View__.
This object (_dns/forward_zone_) represents a forwarding zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiForwardZoneCopyRequest
*/
func (a *ForwardZoneAPIService) ForwardZoneCopy(ctx context.Context) ApiForwardZoneCopyRequest {
	return ApiForwardZoneCopyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ConfigCopyForwardZoneResponse
func (a *ForwardZoneAPIService) ForwardZoneCopyExecute(r ApiForwardZoneCopyRequest) (*ConfigCopyForwardZoneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ConfigCopyForwardZoneResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardZoneAPIService.ForwardZoneCopy")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_zone/copy"

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

type ApiForwardZoneCreateRequest struct {
	ctx        context.Context
	ApiService ForwardZoneAPI
	body       *ConfigForwardZone
}

func (r ApiForwardZoneCreateRequest) Body(body ConfigForwardZone) ApiForwardZoneCreateRequest {
	r.body = &body
	return r
}

func (r ApiForwardZoneCreateRequest) Execute() (*ConfigCreateForwardZoneResponse, *http.Response, error) {
	return r.ApiService.ForwardZoneCreateExecute(r)
}

/*
ForwardZoneCreate Create the ForwardZone object.

Use this method to create a ForwardZone object.
This object (_dns/forward_zone_) represents a forwarding zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiForwardZoneCreateRequest
*/
func (a *ForwardZoneAPIService) ForwardZoneCreate(ctx context.Context) ApiForwardZoneCreateRequest {
	return ApiForwardZoneCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ConfigCreateForwardZoneResponse
func (a *ForwardZoneAPIService) ForwardZoneCreateExecute(r ApiForwardZoneCreateRequest) (*ConfigCreateForwardZoneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ConfigCreateForwardZoneResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardZoneAPIService.ForwardZoneCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_zone"

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

type ApiForwardZoneDeleteRequest struct {
	ctx        context.Context
	ApiService ForwardZoneAPI
	id         string
}

func (r ApiForwardZoneDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ForwardZoneDeleteExecute(r)
}

/*
ForwardZoneDelete Move the Forward Zone object to Recyclebin.

Use this method to move a Forward Zone object to Recyclebin.
This object (_dns/forward_zone_) represents a forwarding zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiForwardZoneDeleteRequest
*/
func (a *ForwardZoneAPIService) ForwardZoneDelete(ctx context.Context, id string) ApiForwardZoneDeleteRequest {
	return ApiForwardZoneDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *ForwardZoneAPIService) ForwardZoneDeleteExecute(r ApiForwardZoneDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardZoneAPIService.ForwardZoneDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_zone/{id}"
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

type ApiForwardZoneListRequest struct {
	ctx        context.Context
	ApiService ForwardZoneAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	tfilter    *string
	torderBy   *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiForwardZoneListRequest) Fields(fields string) ApiForwardZoneListRequest {
	r.fields = &fields
	return r
}

// A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r ApiForwardZoneListRequest) Filter(filter string) ApiForwardZoneListRequest {
	r.filter = &filter
	return r
}

// The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiForwardZoneListRequest) Offset(offset int32) ApiForwardZoneListRequest {
	r.offset = &offset
	return r
}

// The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiForwardZoneListRequest) Limit(limit int32) ApiForwardZoneListRequest {
	r.limit = &limit
	return r
}

// The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiForwardZoneListRequest) PageToken(pageToken string) ApiForwardZoneListRequest {
	r.pageToken = &pageToken
	return r
}

// A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r ApiForwardZoneListRequest) OrderBy(orderBy string) ApiForwardZoneListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiForwardZoneListRequest) Tfilter(tfilter string) ApiForwardZoneListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for sorting by tags.
func (r ApiForwardZoneListRequest) TorderBy(torderBy string) ApiForwardZoneListRequest {
	r.torderBy = &torderBy
	return r
}

func (r ApiForwardZoneListRequest) Execute() (*ConfigListForwardZoneResponse, *http.Response, error) {
	return r.ApiService.ForwardZoneListExecute(r)
}

/*
ForwardZoneList List Forward Zone objects.

Use this method to list Forward Zone objects.
This object (_dns/forward_zone_) represents a forwarding zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiForwardZoneListRequest
*/
func (a *ForwardZoneAPIService) ForwardZoneList(ctx context.Context) ApiForwardZoneListRequest {
	return ApiForwardZoneListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return ConfigListForwardZoneResponse
func (a *ForwardZoneAPIService) ForwardZoneListExecute(r ApiForwardZoneListRequest) (*ConfigListForwardZoneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ConfigListForwardZoneResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardZoneAPIService.ForwardZoneList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_zone"

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
	if r.tfilter != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_tfilter", r.tfilter, "")
	}
	if r.torderBy != nil {
		internal.ParameterAddToHeaderOrQuery(localVarQueryParams, "_torder_by", r.torderBy, "")
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

type ApiForwardZoneReadRequest struct {
	ctx        context.Context
	ApiService ForwardZoneAPI
	id         string
	fields     *string
}

// A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiForwardZoneReadRequest) Fields(fields string) ApiForwardZoneReadRequest {
	r.fields = &fields
	return r
}

func (r ApiForwardZoneReadRequest) Execute() (*ConfigReadForwardZoneResponse, *http.Response, error) {
	return r.ApiService.ForwardZoneReadExecute(r)
}

/*
ForwardZoneRead Read the Forward Zone object.

Use this method to read a Forward Zone object.
This object (_dns/forward_zone_) represents a forwarding zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiForwardZoneReadRequest
*/
func (a *ForwardZoneAPIService) ForwardZoneRead(ctx context.Context, id string) ApiForwardZoneReadRequest {
	return ApiForwardZoneReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ConfigReadForwardZoneResponse
func (a *ForwardZoneAPIService) ForwardZoneReadExecute(r ApiForwardZoneReadRequest) (*ConfigReadForwardZoneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ConfigReadForwardZoneResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardZoneAPIService.ForwardZoneRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_zone/{id}"
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

type ApiForwardZoneUpdateRequest struct {
	ctx        context.Context
	ApiService ForwardZoneAPI
	id         string
	body       *ConfigForwardZone
}

func (r ApiForwardZoneUpdateRequest) Body(body ConfigForwardZone) ApiForwardZoneUpdateRequest {
	r.body = &body
	return r
}

func (r ApiForwardZoneUpdateRequest) Execute() (*ConfigUpdateForwardZoneResponse, *http.Response, error) {
	return r.ApiService.ForwardZoneUpdateExecute(r)
}

/*
ForwardZoneUpdate Update the Forward Zone object.

Use this method to update a Forward Zone object.
This object (_dns/forward_zone_) represents a forwarding zone.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiForwardZoneUpdateRequest
*/
func (a *ForwardZoneAPIService) ForwardZoneUpdate(ctx context.Context, id string) ApiForwardZoneUpdateRequest {
	return ApiForwardZoneUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//
//	@return ConfigUpdateForwardZoneResponse
func (a *ForwardZoneAPIService) ForwardZoneUpdateExecute(r ApiForwardZoneUpdateRequest) (*ConfigUpdateForwardZoneResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *ConfigUpdateForwardZoneResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "ForwardZoneAPIService.ForwardZoneUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/forward_zone/{id}"
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
