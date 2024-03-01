/*
IP Address Management API

Testing OptionSpaceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ipam_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/ipam"
)

var IpamsvcOptionSpace_Post = openapiclient.IpamsvcOptionSpace{
	Id:   openapiclient.PtrString("Test Post"),
	Tags: make(map[string]interface{}),
}
var IpamsvcOptionSpace_Patch = openapiclient.IpamsvcOptionSpace{
	Id:   openapiclient.PtrString("Test Patch"),
	Tags: make(map[string]interface{}),
}

func Test_ipam_OptionSpaceAPIService(t *testing.T) {

	t.Run("Test OptionSpaceAPIService OptionSpaceCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_space", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.IpamsvcOptionSpace
			assert.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			assert.Equal(t, IpamsvcOptionSpace_Post, reqBody)

			response := openapiclient.IpamsvcCreateOptionSpaceResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionSpaceAPI.OptionSpaceCreate(context.Background()).Body(IpamsvcOptionSpace_Post).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionSpaceAPIService OptionSpaceDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_space/"+*IpamsvcOptionSpace_Post.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.OptionSpaceAPI.OptionSpaceDelete(context.Background(), *IpamsvcOptionSpace_Post.Id).Execute()
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionSpaceAPIService OptionSpaceList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_space", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcListOptionSpaceResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionSpaceAPI.OptionSpaceList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionSpaceAPIService OptionSpaceRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_space/"+*IpamsvcOptionSpace_Post.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcReadOptionSpaceResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionSpaceAPI.OptionSpaceRead(context.Background(), *IpamsvcOptionSpace_Post.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test OptionSpaceAPIService OptionSpaceUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/option_space/"+*IpamsvcOptionSpace_Patch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.IpamsvcOptionSpace
			assert.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			assert.Equal(t, IpamsvcOptionSpace_Patch, reqBody)

			response := openapiclient.IpamsvcUpdateOptionSpaceResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.OptionSpaceAPI.OptionSpaceUpdate(context.Background(), *IpamsvcOptionSpace_Patch.Id).Body(IpamsvcOptionSpace_Patch).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})
}
