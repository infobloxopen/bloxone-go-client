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


type OptionGroupAPI interface {
	/*
	OptionGroupCreate Create the DHCP option group.

	Use this method to create an __OptionGroup__ object.
The __OptionGroup__ object is a named collection of options.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOptionGroupCreateRequest
	*/
	OptionGroupCreate(ctx context.Context) ApiOptionGroupCreateRequest

	// OptionGroupCreateExecute executes the request
	//  @return IpamsvcCreateOptionGroupResponse
	OptionGroupCreateExecute(r ApiOptionGroupCreateRequest) (*IpamsvcCreateOptionGroupResponse, *http.Response, error)
	/*
	OptionGroupDelete Move the DHCP option group to the recycle bin.

	Use this method to move an __OptionGroup__ object to the recycle bin.
The __OptionGroup__ object is a named collection of options.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiOptionGroupDeleteRequest
	*/
	OptionGroupDelete(ctx context.Context, id string) ApiOptionGroupDeleteRequest

	// OptionGroupDeleteExecute executes the request
	OptionGroupDeleteExecute(r ApiOptionGroupDeleteRequest) (*http.Response, error)
	/*
	OptionGroupList Retrieve DHCP option groups.

	Use this method to retrieve __OptionGroup__ objects.
The __OptionGroup__ object is a named collection of options.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiOptionGroupListRequest
	*/
	OptionGroupList(ctx context.Context) ApiOptionGroupListRequest

	// OptionGroupListExecute executes the request
	//  @return IpamsvcListOptionGroupResponse
	OptionGroupListExecute(r ApiOptionGroupListRequest) (*IpamsvcListOptionGroupResponse, *http.Response, error)
	/*
	OptionGroupRead Retrieve the DHCP option group.

	Use this method to retrieve an __OptionGroup__ object.
The __OptionGroup__ object is a named collection of options.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiOptionGroupReadRequest
	*/
	OptionGroupRead(ctx context.Context, id string) ApiOptionGroupReadRequest

	// OptionGroupReadExecute executes the request
	//  @return IpamsvcReadOptionGroupResponse
	OptionGroupReadExecute(r ApiOptionGroupReadRequest) (*IpamsvcReadOptionGroupResponse, *http.Response, error)
	/*
	OptionGroupUpdate Update the DHCP option group.

	Use this method to update an __OptionGroup__ object.
The __OptionGroup__ object is a named collection of options.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id An application specific resource identity of a resource
	@return ApiOptionGroupUpdateRequest
	*/
	OptionGroupUpdate(ctx context.Context, id string) ApiOptionGroupUpdateRequest

	// OptionGroupUpdateExecute executes the request
	//  @return IpamsvcUpdateOptionGroupResponse
	OptionGroupUpdateExecute(r ApiOptionGroupUpdateRequest) (*IpamsvcUpdateOptionGroupResponse, *http.Response, error)
}

// OptionGroupAPIService OptionGroupAPI service
type OptionGroupAPIService internal.Service

type ApiOptionGroupCreateRequest struct {
	ctx context.Context
	ApiService OptionGroupAPI
	body *IpamsvcOptionGroup
}

func (r ApiOptionGroupCreateRequest) Body(body IpamsvcOptionGroup) ApiOptionGroupCreateRequest {
	r.body = &body
	return r
}

func (r ApiOptionGroupCreateRequest) Execute() (*IpamsvcCreateOptionGroupResponse, *http.Response, error) {
	return r.ApiService.OptionGroupCreateExecute(r)
}

/*
OptionGroupCreate Create the DHCP option group.

Use this method to create an __OptionGroup__ object.
The __OptionGroup__ object is a named collection of options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOptionGroupCreateRequest
*/
func (a *OptionGroupAPIService) OptionGroupCreate(ctx context.Context) ApiOptionGroupCreateRequest {
	return ApiOptionGroupCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IpamsvcCreateOptionGroupResponse
func (a *OptionGroupAPIService) OptionGroupCreateExecute(r ApiOptionGroupCreateRequest) (*IpamsvcCreateOptionGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcCreateOptionGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionGroupAPIService.OptionGroupCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_group"

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

type ApiOptionGroupDeleteRequest struct {
	ctx context.Context
	ApiService OptionGroupAPI
	id string
}

func (r ApiOptionGroupDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.OptionGroupDeleteExecute(r)
}

/*
OptionGroupDelete Move the DHCP option group to the recycle bin.

Use this method to move an __OptionGroup__ object to the recycle bin.
The __OptionGroup__ object is a named collection of options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiOptionGroupDeleteRequest
*/
func (a *OptionGroupAPIService) OptionGroupDelete(ctx context.Context, id string) ApiOptionGroupDeleteRequest {
	return ApiOptionGroupDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *OptionGroupAPIService) OptionGroupDeleteExecute(r ApiOptionGroupDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionGroupAPIService.OptionGroupDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_group/{id}"
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

type ApiOptionGroupListRequest struct {
	ctx context.Context
	ApiService OptionGroupAPI
	fields *string
	filter *string
	offset *int32
	limit *int32
	pageToken *string
	orderBy *string
	torderBy *string
	tfilter *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiOptionGroupListRequest) Fields(fields string) ApiOptionGroupListRequest {
	r.fields = &fields
	return r
}

//   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |        
func (r ApiOptionGroupListRequest) Filter(filter string) ApiOptionGroupListRequest {
	r.filter = &filter
	return r
}

//   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.         
func (r ApiOptionGroupListRequest) Offset(offset int32) ApiOptionGroupListRequest {
	r.offset = &offset
	return r
}

//   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.         
func (r ApiOptionGroupListRequest) Limit(limit int32) ApiOptionGroupListRequest {
	r.limit = &limit
	return r
}

//   The service-defined string used to identify a page of resources. A null value indicates the first page.         
func (r ApiOptionGroupListRequest) PageToken(pageToken string) ApiOptionGroupListRequest {
	r.pageToken = &pageToken
	return r
}

//   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.        
func (r ApiOptionGroupListRequest) OrderBy(orderBy string) ApiOptionGroupListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for sorting by tags.
func (r ApiOptionGroupListRequest) TorderBy(torderBy string) ApiOptionGroupListRequest {
	r.torderBy = &torderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiOptionGroupListRequest) Tfilter(tfilter string) ApiOptionGroupListRequest {
	r.tfilter = &tfilter
	return r
}

func (r ApiOptionGroupListRequest) Execute() (*IpamsvcListOptionGroupResponse, *http.Response, error) {
	return r.ApiService.OptionGroupListExecute(r)
}

/*
OptionGroupList Retrieve DHCP option groups.

Use this method to retrieve __OptionGroup__ objects.
The __OptionGroup__ object is a named collection of options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOptionGroupListRequest
*/
func (a *OptionGroupAPIService) OptionGroupList(ctx context.Context) ApiOptionGroupListRequest {
	return ApiOptionGroupListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IpamsvcListOptionGroupResponse
func (a *OptionGroupAPIService) OptionGroupListExecute(r ApiOptionGroupListRequest) (*IpamsvcListOptionGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcListOptionGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionGroupAPIService.OptionGroupList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_group"

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

type ApiOptionGroupReadRequest struct {
	ctx context.Context
	ApiService OptionGroupAPI
	id string
	fields *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.        
func (r ApiOptionGroupReadRequest) Fields(fields string) ApiOptionGroupReadRequest {
	r.fields = &fields
	return r
}

func (r ApiOptionGroupReadRequest) Execute() (*IpamsvcReadOptionGroupResponse, *http.Response, error) {
	return r.ApiService.OptionGroupReadExecute(r)
}

/*
OptionGroupRead Retrieve the DHCP option group.

Use this method to retrieve an __OptionGroup__ object.
The __OptionGroup__ object is a named collection of options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiOptionGroupReadRequest
*/
func (a *OptionGroupAPIService) OptionGroupRead(ctx context.Context, id string) ApiOptionGroupReadRequest {
	return ApiOptionGroupReadRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IpamsvcReadOptionGroupResponse
func (a *OptionGroupAPIService) OptionGroupReadExecute(r ApiOptionGroupReadRequest) (*IpamsvcReadOptionGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcReadOptionGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionGroupAPIService.OptionGroupRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_group/{id}"
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

type ApiOptionGroupUpdateRequest struct {
	ctx context.Context
	ApiService OptionGroupAPI
	id string
	body *IpamsvcOptionGroup
}

func (r ApiOptionGroupUpdateRequest) Body(body IpamsvcOptionGroup) ApiOptionGroupUpdateRequest {
	r.body = &body
	return r
}

func (r ApiOptionGroupUpdateRequest) Execute() (*IpamsvcUpdateOptionGroupResponse, *http.Response, error) {
	return r.ApiService.OptionGroupUpdateExecute(r)
}

/*
OptionGroupUpdate Update the DHCP option group.

Use this method to update an __OptionGroup__ object.
The __OptionGroup__ object is a named collection of options.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiOptionGroupUpdateRequest
*/
func (a *OptionGroupAPIService) OptionGroupUpdate(ctx context.Context, id string) ApiOptionGroupUpdateRequest {
	return ApiOptionGroupUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return IpamsvcUpdateOptionGroupResponse
func (a *OptionGroupAPIService) OptionGroupUpdateExecute(r ApiOptionGroupUpdateRequest) (*IpamsvcUpdateOptionGroupResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []internal.FormFile
		localVarReturnValue  *IpamsvcUpdateOptionGroupResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "OptionGroupAPIService.OptionGroupUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dhcp/option_group/{id}"
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
