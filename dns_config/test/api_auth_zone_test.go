/*
DNS Configuration API

Testing AclAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package dns_config_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/infobloxopen/bloxone-go-client/dns_config"
	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

func Test_dns_config_AuthZoneAPIService(t *testing.T) {

	t.Run("Test AuthZoneAPIService AuthZoneCreate", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		dummyAuthZone := dns_config.ConfigAuthZone{
			Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
			Id:      openapiclient.PtrString("dummyAuthZoneId"),
		}

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "POST", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone", req.URL.Path)

			response := dns_config.ConfigCreateAuthZoneResponse{
				Result: &dummyAuthZone,
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		authZoneAPI := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		createRequest := authZoneAPI.AuthZoneAPI.AuthZoneCreate(ctx).Body(dummyAuthZone)
		resp, httpRes, err := createRequest.Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 201, httpRes.StatusCode)
		require.Equal(t, dummyAuthZone, *resp.Result)
	})

	t.Run("Test AuthZoneAPIService AuthZoneCopy", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		dummyAuthZone := dns_config.ConfigCopyAuthZone{
			Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
			Id:      openapiclient.PtrString("dummyAuthZoneId"),
		}

		dummyConfigCopyResponse := dns_config.ConfigCopyResponse{
			Description: openapiclient.PtrString("This is a copy of the dummy AuthZone for testing."),
			Id:          openapiclient.PtrString("dummyCopyId"),
			JobId:       openapiclient.PtrString("dummyJobId"),
		}

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "POST", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/copy", req.URL.Path)
			//require.Equal(t, "/api/ddi/v1/dns/auth_zone/copy", req.URL.Path)

			response := dns_config.ConfigCopyAuthZoneResponse{
				Result: &dummyConfigCopyResponse,
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: 201,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		authZoneAPI := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		copyRequest := authZoneAPI.AuthZoneAPI.AuthZoneCopy(ctx).Body(dummyAuthZone)
		resp, httpRes, err := copyRequest.Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 201, httpRes.StatusCode)
		require.Equal(t, dummyAuthZone, *resp.Result)
	})

	t.Run("Test AuthZoneAPIService AuthZoneRead", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		dummyAuthZone := dns_config.ConfigAuthZone{
			Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
			Id:      openapiclient.PtrString("dummyAuthZoneId"),
		}

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "GET", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/"+*dummyAuthZone.Id, req.URL.Path)

			response := dns_config.ConfigReadAuthZoneResponse{
				Result: &dummyAuthZone,
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		authZoneAPI := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		readRequest := authZoneAPI.AuthZoneAPI.AuthZoneRead(ctx, *dummyAuthZone.Id)
		resp, httpRes, err := readRequest.Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
		require.Equal(t, dummyAuthZone, *resp.Result)
	})

	t.Run("Test AuthZoneAPIService AuthZoneList", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		dummyAuthZone := dns_config.ConfigAuthZone{
			Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
			Id:      openapiclient.PtrString("dummyAuthZoneId"),
		}

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "GET", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone", req.URL.Path)

			response := dns_config.ConfigListAuthZoneResponse{
				Results: []dns_config.ConfigAuthZone{dummyAuthZone},
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		authZoneAPI := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		listRequest := authZoneAPI.AuthZoneAPI.AuthZoneList(ctx)
		resp, httpRes, err := listRequest.Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
		require.Equal(t, dummyAuthZone, resp.Results[0])
	})

	t.Run("Test AuthZoneAPIService AuthZoneUpdate", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		dummyAuthZone := dns_config.ConfigAuthZone{
			Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
			Id:      openapiclient.PtrString("dummyAuthZoneId"),
		}

		updatedAuthZone := dns_config.ConfigAuthZone{
			Comment: openapiclient.PtrString("This is an updated dummy AuthZone for testing."),
			Id:      dummyAuthZone.Id,
		}

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "PATCH", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/"+*dummyAuthZone.Id, req.URL.Path)

			response := dns_config.ConfigUpdateAuthZoneResponse{
				Result: &updatedAuthZone,
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		authZoneAPI := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		updateRequest := authZoneAPI.AuthZoneAPI.AuthZoneUpdate(ctx, *dummyAuthZone.Id).Body(updatedAuthZone)
		resp, httpRes, err := updateRequest.Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
		require.Equal(t, updatedAuthZone, *resp.Result)
	})

	t.Run("Test AuthZoneAPIService AuthZoneDelete", func(t *testing.T) {

		//t.Skip("skip test") // remove to run test

		dummyAuthZone := dns_config.ConfigAuthZone{
			Comment: openapiclient.PtrString("This is a dummy AuthZone for testing."),
			Id:      openapiclient.PtrString("dummyAuthZoneId"),
		}

		testClient := NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "DELETE", req.Method)
			require.Equal(t, "/api/ddi/v1/dns/auth_zone/"+*dummyAuthZone.Id, req.URL.Path)

			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})

		configuration := internal.NewConfiguration()
		configuration.HTTPClient = testClient

		authZoneAPI := dns_config.NewAPIClient(configuration)
		ctx := context.Background()

		httpRes, err := authZoneAPI.AuthZoneAPI.AuthZoneDelete(ctx, *dummyAuthZone.Id).Execute()
		require.Nil(t, err)
		require.Equal(t, 200, httpRes.StatusCode)
	})

}
