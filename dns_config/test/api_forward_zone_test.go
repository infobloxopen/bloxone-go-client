/*
DNS Configuration API

Testing ForwardZoneAPIService

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

var ConfigCopyForwardZonePost = openapiclient.ConfigCopyForwardZone{
	Id:                 openapiclient.PtrString("ConfigCopyForwardZonePost"),
	Comment:            nil,
	ExternalForwarders: nil,
	Hosts:              nil,
	InternalForwarders: nil,
	Nsgs:               nil,
	Recursive:          nil,
	SkipOnError:        nil,
	TargetView:         "",
}
var ConfigForwardZonePost = openapiclient.ConfigForwardZone{
	Id:                 openapiclient.PtrString("ConfigForwardZonePost"),
	Comment:            nil,
	CreatedAt:          nil,
	Disabled:           nil,
	ExternalForwarders: nil,
	ForwardOnly:        nil,
	Fqdn:               nil,
	Hosts:              nil,
	InternalForwarders: nil,
	MappedSubnet:       nil,
	Mapping:            nil,
	Nsgs:               nil,
	Parent:             nil,
	ProtocolFqdn:       nil,
	Tags:               make(map[string]interface{}),
	UpdatedAt:          nil,
	View:               nil,
	Warnings:           nil,
}

var ConfigForwardZonePatch = openapiclient.ConfigForwardZone{
	Id:                 openapiclient.PtrString("ConfigForwardZonePatch"),
	Comment:            nil,
	CreatedAt:          nil,
	Disabled:           nil,
	ExternalForwarders: nil,
	ForwardOnly:        nil,
	Fqdn:               nil,
	Hosts:              nil,
	InternalForwarders: nil,
	MappedSubnet:       nil,
	Mapping:            nil,
	Nsgs:               nil,
	Parent:             nil,
	ProtocolFqdn:       nil,
	Tags:               make(map[string]interface{}),
	UpdatedAt:          nil,
	View:               nil,
	Warnings:           nil,
}

func TestForwardZoneAPIService(t *testing.T) {

	t.Run("Test ForwardZoneAPIService ForwardZoneCopy", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_zone/copy", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigCopyForwardZone
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigCopyForwardZonePost, reqBody)

			response := openapiclient.ConfigCopyForwardZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneCopy(context.Background()).Body(ConfigCopyForwardZonePost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardZoneAPIService ForwardZoneCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_zone", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigForwardZone
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigForwardZonePost, reqBody)

			response := openapiclient.ConfigCreateForwardZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneCreate(context.Background()).Body(ConfigForwardZonePost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardZoneAPIService ForwardZoneList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_zone", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigListForwardZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneList(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardZoneAPIService ForwardZoneRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodGet, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_zone/"+*ConfigForwardZonePost.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.ConfigReadForwardZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneRead(context.Background(), *ConfigForwardZonePost.Id).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardZoneAPIService ForwardZoneUpdate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPatch, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_zone/"+*ConfigForwardZonePatch.Id, req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			var reqBody openapiclient.ConfigForwardZone
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigForwardZonePatch, reqBody)

			response := openapiclient.ConfigUpdateForwardZoneResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneUpdate(context.Background(), *ConfigForwardZonePatch.Id).Body(ConfigForwardZonePatch).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

	t.Run("Test ForwardZoneAPIService ForwardZoneDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodDelete, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/forward_zone/"+*ConfigForwardZonePost.Id, req.URL.Path)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.ForwardZoneAPI.ForwardZoneDelete(context.Background(), *ConfigForwardZonePost.Id).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
