/*
DNS Configuration API

Testing ForwardNsgAPIService

*/

package dns_config_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	openapiclient "github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/internal"
)

var ConfigForwardNSGPost = openapiclient.ConfigForwardNSG{
	Id:                 openapiclient.PtrString("ConfigForwardNSGPost"),
	Comment:            nil,
	ExternalForwarders: nil,
	ForwardersOnly:     nil,
	Hosts:              nil,
	InternalForwarders: nil,
	Name:               "",
	Nsgs:               nil,
	Tags:               make(map[string]interface{}),
}
var ConfigForwardNSGPatch = openapiclient.ConfigForwardNSG{
	Id:                 openapiclient.PtrString("ConfigForwardNSGPatch"),
	Comment:            nil,
	ExternalForwarders: nil,
	ForwardersOnly:     nil,
	Hosts:              nil,
	InternalForwarders: nil,
	Name:               "",
	Nsgs:               nil,
	Tags:               make(map[string]interface{}),
}

func TestForwardNsgAPIService(t *testing.T) {

	t.Run("Test ForwardNsgAPIService ForwardNsgCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_nsg", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigForwardNSG
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigForwardNSGPost, reqBody)

			response := openapiclient.ConfigCreateForwardNSGResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardNsgAPI.ForwardNsgCreate(context.Background()).Body(ConfigForwardNSGPost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardNsgAPIService ForwardNsgDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_nsg/"+*ConfigForwardNSGPost.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.ForwardNsgAPI.ForwardNsgDelete(context.Background(), *ConfigForwardNSGPost.Id).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardNsgAPIService ForwardNsgList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_nsg", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigListForwardNSGResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardNsgAPI.ForwardNsgList(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardNsgAPIService ForwardNsgRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_nsg/"+*ConfigForwardNSGPost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigReadForwardNSGResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardNsgAPI.ForwardNsgRead(context.Background(), *ConfigForwardNSGPost.Id).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardNsgAPIService ForwardNsgUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_nsg/"+*ConfigForwardNSGPatch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigForwardNSG
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigForwardNSGPatch, reqBody)

			response := openapiclient.ConfigUpdateForwardNSGResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardNsgAPI.ForwardNsgUpdate(context.Background(), *ConfigForwardNSGPatch.Id).Body(ConfigForwardNSGPatch).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
