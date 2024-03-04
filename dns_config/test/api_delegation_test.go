/*
DNS Configuration API

Testing DelegationAPIService

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

var ConfigDelegationPost = openapiclient.ConfigDelegation{
	Id:                openapiclient.PtrString("ConfigDelegationPost"),
	Comment:           nil,
	DelegationServers: nil,
	Disabled:          nil,
	Fqdn:              nil,
	Parent:            nil,
	ProtocolFqdn:      nil,
	Tags:              make(map[string]interface{}),
	View:              nil,
}
var ConfigDelegationPatch = openapiclient.ConfigDelegation{
	Id:                openapiclient.PtrString("ConfigDelegationPatch"),
	Comment:           nil,
	DelegationServers: nil,
	Disabled:          nil,
	Fqdn:              nil,
	Parent:            nil,
	ProtocolFqdn:      nil,
	Tags:              make(map[string]interface{}),
	View:              nil,
}

func TestDelegationAPIService(t *testing.T) {

	t.Run("Test DelegationAPIService DelegationCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/delegation", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigDelegation
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigDelegationPost, reqBody)

			response := openapiclient.ConfigCreateDelegationResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DelegationAPI.DelegationCreate(context.Background()).Body(ConfigDelegationPost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DelegationAPIService DelegationList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/delegation", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigListDelegationResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DelegationAPI.DelegationList(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DelegationAPIService DelegationRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/delegation/"+*ConfigDelegationPost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigReadDelegationResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DelegationAPI.DelegationRead(context.Background(), *ConfigDelegationPost.Id).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DelegationAPIService DelegationUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/delegation/"+*ConfigDelegationPatch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			var reqBody openapiclient.ConfigDelegation
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigDelegationPatch, reqBody)

			response := openapiclient.ConfigUpdateDelegationResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.DelegationAPI.DelegationUpdate(context.Background(), *ConfigDelegationPatch.Id).Body(ConfigDelegationPatch).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test DelegationAPIService DelegationDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/delegation/"+*ConfigDelegationPost.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.DelegationAPI.DelegationDelete(context.Background(), *ConfigDelegationPost.Id).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
