/*
DDI Keys API

The DDI Keys application is a BloxOne DDI service for managing TSIG keys and GSS-TSIG (Kerberos) keys which are used by other BloxOne DDI applications. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package keys

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type KerberosAPI interface {

	/*
			KerberosDelete Delete the Kerberos key.

			Use this method to delete a __KerberosKey__ object.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiKerberosDeleteRequest
	*/
	KerberosDelete(ctx context.Context, id string) ApiKerberosDeleteRequest

	// KerberosDeleteExecute executes the request
	KerberosDeleteExecute(r ApiKerberosDeleteRequest) (*http.Response, error)

	/*
			KerberosList Retrieve Kerberos keys.

			Use this method to retrieve __KerberosKey__ objects.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiKerberosListRequest
	*/
	KerberosList(ctx context.Context) ApiKerberosListRequest

	// KerberosListExecute executes the request
	//  @return KeysListKerberosKeyResponse
	KerberosListExecute(r ApiKerberosListRequest) (*KeysListKerberosKeyResponse, *http.Response, error)

	/*
			KerberosRead Retrieve the Kerberos key.

			Use this method to retrieve a __KerberosKey__ object.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiKerberosReadRequest
	*/
	KerberosRead(ctx context.Context, id string) ApiKerberosReadRequest

	// KerberosReadExecute executes the request
	//  @return KeysReadKerberosKeyResponse
	KerberosReadExecute(r ApiKerberosReadRequest) (*KeysReadKerberosKeyResponse, *http.Response, error)

	/*
			KerberosUpdate Update the Kerberos key.

			Use this method to update a __KerberosKey__ object.
		A __KerberosKey__ object represents a Kerberos key.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param id An application specific resource identity of a resource
			@return ApiKerberosUpdateRequest
	*/
	KerberosUpdate(ctx context.Context, id string) ApiKerberosUpdateRequest

	// KerberosUpdateExecute executes the request
	//  @return KeysUpdateKerberosKeyResponse
	KerberosUpdateExecute(r ApiKerberosUpdateRequest) (*KeysUpdateKerberosKeyResponse, *http.Response, error)

	/*
		KeysKerberosPost Method for KeysKerberosPost

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiKeysKerberosPostRequest
	*/
	KeysKerberosPost(ctx context.Context) ApiKeysKerberosPostRequest

	// KeysKerberosPostExecute executes the request
	//  @return KeysListKerberosKeyResponse
	KeysKerberosPostExecute(r ApiKeysKerberosPostRequest) (*KeysListKerberosKeyResponse, *http.Response, error)
}

// KerberosAPIService KerberosAPI service
type KerberosAPIService internal.Service

type ApiKerberosDeleteRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	id         string
}

func (r ApiKerberosDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.KerberosDeleteExecute(r)
}

/*
KerberosDelete Delete the Kerberos key.

Use this method to delete a __KerberosKey__ object.
A __KerberosKey__ object represents a Kerberos key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiKerberosDeleteRequest
*/
func (a *KerberosAPIService) KerberosDelete(ctx context.Context, id string) ApiKerberosDeleteRequest {
	return ApiKerberosDeleteRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
func (a *KerberosAPIService) KerberosDeleteExecute(r ApiKerberosDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []internal.FormFile
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.KerberosDelete")
	if err != nil {
		return nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos/{id}"
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

type ApiKerberosListRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	fields     *string
	filter     *string
	offset     *int32
	limit      *int32
	pageToken  *string
	orderBy    *string
	tfilter    *string
	torderBy   *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiKerberosListRequest) Fields(fields string) ApiKerberosListRequest {
	r.fields = &fields
	return r
}

//   A collection of response resources can be filtered by a logical expression string that includes JSON tag references to values in each resource, literal values, and logical operators. If a resource does not have the specified tag, its value is assumed to be null.  Literal values include numbers (integer and floating-point), and quoted (both single- or double-quoted) literal strings, and &#39;null&#39;. The following operators are commonly used in filter expressions:  |  Op   |  Description               |  |  --   |  -----------               |  |  &#x3D;&#x3D;   |  Equal                     |  |  !&#x3D;   |  Not Equal                 |  |  &gt;    |  Greater Than              |  |   &gt;&#x3D;  |  Greater Than or Equal To  |  |  &lt;    |  Less Than                 |  |  &lt;&#x3D;   |  Less Than or Equal To     |  |  and  |  Logical AND               |  |  ~    |  Matches Regex             |  |  !~   |  Does Not Match Regex      |  |  or   |  Logical OR                |  |  not  |  Logical NOT               |  |  ()   |  Groupping Operators       |
func (r ApiKerberosListRequest) Filter(filter string) ApiKerberosListRequest {
	r.filter = &filter
	return r
}

//   The integer index (zero-origin) of the offset into a collection of resources. If omitted or null the value is assumed to be &#39;0&#39;.
func (r ApiKerberosListRequest) Offset(offset int32) ApiKerberosListRequest {
	r.offset = &offset
	return r
}

//   The integer number of resources to be returned in the response. The service may impose maximum value. If omitted the service may impose a default value.
func (r ApiKerberosListRequest) Limit(limit int32) ApiKerberosListRequest {
	r.limit = &limit
	return r
}

//   The service-defined string used to identify a page of resources. A null value indicates the first page.
func (r ApiKerberosListRequest) PageToken(pageToken string) ApiKerberosListRequest {
	r.pageToken = &pageToken
	return r
}

//   A collection of response resources can be sorted by their JSON tags. For a &#39;flat&#39; resource, the tag name is straightforward. If sorting is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, its value is assumed to be null.)  Specify this parameter as a comma-separated list of JSON tag names. The sort direction can be specified by a suffix separated by whitespace before the tag name. The suffix &#39;asc&#39; sorts the data in ascending order. The suffix &#39;desc&#39; sorts the data in descending order. If no suffix is specified the data is sorted in ascending order.
func (r ApiKerberosListRequest) OrderBy(orderBy string) ApiKerberosListRequest {
	r.orderBy = &orderBy
	return r
}

// This parameter is used for filtering by tags.
func (r ApiKerberosListRequest) Tfilter(tfilter string) ApiKerberosListRequest {
	r.tfilter = &tfilter
	return r
}

// This parameter is used for sorting by tags.
func (r ApiKerberosListRequest) TorderBy(torderBy string) ApiKerberosListRequest {
	r.torderBy = &torderBy
	return r
}

func (r ApiKerberosListRequest) Execute() (*KeysListKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.KerberosListExecute(r)
}

/*
KerberosList Retrieve Kerberos keys.

Use this method to retrieve __KerberosKey__ objects.
A __KerberosKey__ object represents a Kerberos key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiKerberosListRequest
*/
func (a *KerberosAPIService) KerberosList(ctx context.Context) ApiKerberosListRequest {
	return ApiKerberosListRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return KeysListKerberosKeyResponse
func (a *KerberosAPIService) KerberosListExecute(r ApiKerberosListRequest) (*KeysListKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *KeysListKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.KerberosList")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos"

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

type ApiKerberosReadRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	id         string
	fields     *string
}

//   A collection of response resources can be transformed by specifying a set of JSON tags to be returned. For a “flat” resource, the tag name is straightforward. If field selection is allowed on non-flat hierarchical resources, the service should implement a qualified naming scheme such as dot-qualification to reference data down the hierarchy. If a resource does not have the specified tag, the tag does not appear in the output resource.  Specify this parameter as a comma-separated list of JSON tag names.
func (r ApiKerberosReadRequest) Fields(fields string) ApiKerberosReadRequest {
	r.fields = &fields
	return r
}

func (r ApiKerberosReadRequest) Execute() (*KeysReadKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.KerberosReadExecute(r)
}

/*
KerberosRead Retrieve the Kerberos key.

Use this method to retrieve a __KerberosKey__ object.
A __KerberosKey__ object represents a Kerberos key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiKerberosReadRequest
*/
func (a *KerberosAPIService) KerberosRead(ctx context.Context, id string) ApiKerberosReadRequest {
	return ApiKerberosReadRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return KeysReadKerberosKeyResponse
func (a *KerberosAPIService) KerberosReadExecute(r ApiKerberosReadRequest) (*KeysReadKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *KeysReadKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.KerberosRead")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos/{id}"
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

type ApiKerberosUpdateRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	id         string
	body       *KerberosKey
}

func (r ApiKerberosUpdateRequest) Body(body KerberosKey) ApiKerberosUpdateRequest {
	r.body = &body
	return r
}

func (r ApiKerberosUpdateRequest) Execute() (*KeysUpdateKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.KerberosUpdateExecute(r)
}

/*
KerberosUpdate Update the Kerberos key.

Use this method to update a __KerberosKey__ object.
A __KerberosKey__ object represents a Kerberos key.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id An application specific resource identity of a resource
 @return ApiKerberosUpdateRequest
*/
func (a *KerberosAPIService) KerberosUpdate(ctx context.Context, id string) ApiKerberosUpdateRequest {
	return ApiKerberosUpdateRequest{
		ApiService: a,
		ctx:        ctx,
		id:         id,
	}
}

// Execute executes the request
//  @return KeysUpdateKerberosKeyResponse
func (a *KerberosAPIService) KerberosUpdateExecute(r ApiKerberosUpdateRequest) (*KeysUpdateKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *KeysUpdateKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.KerberosUpdate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos/{id}"
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

type ApiKeysKerberosPostRequest struct {
	ctx        context.Context
	ApiService KerberosAPI
	body       *KerberosKey
}

func (r ApiKeysKerberosPostRequest) Body(body KerberosKey) ApiKeysKerberosPostRequest {
	r.body = &body
	return r
}

func (r ApiKeysKerberosPostRequest) Execute() (*KeysListKerberosKeyResponse, *http.Response, error) {
	return r.ApiService.KeysKerberosPostExecute(r)
}

/*
KeysKerberosPost Method for KeysKerberosPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiKeysKerberosPostRequest
*/
func (a *KerberosAPIService) KeysKerberosPost(ctx context.Context) ApiKeysKerberosPostRequest {
	return ApiKeysKerberosPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return KeysListKerberosKeyResponse
func (a *KerberosAPIService) KeysKerberosPostExecute(r ApiKeysKerberosPostRequest) (*KeysListKerberosKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue *KeysListKerberosKeyResponse
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "KerberosAPIService.KeysKerberosPost")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/keys/kerberos"

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
