/*
IP Address Management API

Testing DhcpHostAPIService

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

var IpamsvcHost_Patch = openapiclient.IpamsvcHost{
	Id:   openapiclient.PtrString("Test"),
	Tags: make(map[string]interface{}),
}

func Test_ipam_DhcpHostAPIService(t *testing.T) {

	t.Run("Test DhcpHostAPIService DhcpHostList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/host", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcListHostResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DhcpHostAPI.DhcpHostList(context.Background()).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DhcpHostAPIService DhcpHostListAssociations", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/host/"+*IpamsvcHost_Patch.Id+"/associations", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcHostAssociationsResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DhcpHostAPI.DhcpHostListAssociations(context.Background(), *IpamsvcHost_Patch.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DhcpHostAPIService DhcpHostRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/host/"+*IpamsvcHost_Patch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.IpamsvcReadHostResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DhcpHostAPI.DhcpHostRead(context.Background(), *IpamsvcHost_Patch.Id).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DhcpHostAPIService DhcpHostUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dhcp/host/"+*IpamsvcHost_Patch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.IpamsvcHost
			assert.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			assert.Equal(t, IpamsvcHost_Patch, reqBody)

			response := openapiclient.IpamsvcUpdateHostResponse{}
			body, err := json.Marshal(response)
			assert.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DhcpHostAPI.DhcpHostUpdate(context.Background(), *IpamsvcHost_Patch.Id).Body(IpamsvcHost_Patch).Execute()
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})
}
