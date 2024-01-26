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

	"github.com/infobloxopen/bloxone-go-client/internal"
)

type CacheFlushAPI interface {

	/*
			CacheFlushCreate Create the Cache Flush object.

			Use this method to create a Cache Flush object.
		The Cache Flush object is for removing entries from the DNS cache on a host. The host must be available and running DNS for this to succeed.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return ApiCacheFlushCreateRequest
	*/
	CacheFlushCreate(ctx context.Context) ApiCacheFlushCreateRequest

	// CacheFlushCreateExecute executes the request
	//  @return map[string]interface{}
	CacheFlushCreateExecute(r ApiCacheFlushCreateRequest) (map[string]interface{}, *http.Response, error)
}

// CacheFlushAPIService CacheFlushAPI service
type CacheFlushAPIService internal.Service

type ApiCacheFlushCreateRequest struct {
	ctx        context.Context
	ApiService CacheFlushAPI
	body       *ConfigCacheFlush
}

func (r ApiCacheFlushCreateRequest) Body(body ConfigCacheFlush) ApiCacheFlushCreateRequest {
	r.body = &body
	return r
}

func (r ApiCacheFlushCreateRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.CacheFlushCreateExecute(r)
}

/*
CacheFlushCreate Create the Cache Flush object.

Use this method to create a Cache Flush object.
The Cache Flush object is for removing entries from the DNS cache on a host. The host must be available and running DNS for this to succeed.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCacheFlushCreateRequest
*/
func (a *CacheFlushAPIService) CacheFlushCreate(ctx context.Context) ApiCacheFlushCreateRequest {
	return ApiCacheFlushCreateRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *CacheFlushAPIService) CacheFlushCreateExecute(r ApiCacheFlushCreateRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []internal.FormFile
		localVarReturnValue map[string]interface{}
	)

	localBasePath, err := a.Client.Cfg.ServerURLWithContext(r.ctx, "CacheFlushAPIService.CacheFlushCreate")
	if err != nil {
		return localVarReturnValue, nil, internal.NewGenericOpenAPIError(err.Error())
	}

	localVarPath := localBasePath + "/dns/cache_flush"

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
