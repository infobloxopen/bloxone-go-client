package keys_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"testing"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func Test_keys_TsigAPIService(t *testing.T) {
	t.Run("Test TsigAPIService TsigCreate", func(t *testing.T) {
		dummyKey := openapiclient.KeysTSIGKey{
			Algorithm: openapiclient.PtrString("hmac_sha256"),
			Name:      "dummyKey36",
			Secret:    "XOJvQkcX6Og0CHFg+rQ27pqAB+EhSjVoI4Bs/JWegBc=",
		}
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "POST", req.Method)
			require.Equal(t, "/api/ddi/v1/keys/tsig", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))
			var reqBody openapiclient.KeysTSIGKey
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, dummyKey, reqBody)

			response := openapiclient.KeysCreateTSIGKeyResponse{
				Result: &dummyKey,
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.TsigAPI.TsigCreate(context.Background()).Body(dummyKey).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TsigAPIService TsigList", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "GET", req.Method)
			require.Equal(t, "/api/ddi/v1/keys/tsig", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.KeysListTSIGKeyResponse{
				Results: []openapiclient.KeysTSIGKey{
					{
						Algorithm: openapiclient.PtrString("hmac_sha256"),
						Name:      "dummyKey36",
						Secret:    "XOJvQkcX6Og0CHFg+rQ27pqAB+EhSjVoI4Bs/JWegBc=",
					},
				},
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.TsigAPI.TsigList(context.Background()).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
		require.NotEmpty(t, resp.Results)
	})

	t.Run("Test TsigAPIService TsigRead", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "GET", req.Method)
			require.Equal(t, "/api/ddi/v1/keys/tsig/dummyKey36", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Accept"))

			response := openapiclient.KeysListTSIGKeyResponse{
				Results: []openapiclient.KeysTSIGKey{
					{
						Algorithm: openapiclient.PtrString("hmac_sha256"),
						Name:      "dummyKey36",
						Secret:    "XOJvQkcX6Og0CHFg+rQ27pqAB+EhSjVoI4Bs/JWegBc=",
					},
				},
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.TsigAPI.TsigRead(context.Background(), "dummyKey36").Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
	})
	t.Run("Test TsigAPIService TsigUpdate", func(t *testing.T) {
		dummyKey := openapiclient.KeysTSIGKey{
			Algorithm: openapiclient.PtrString("hmac_sha256"),
			Name:      "dummyKey36",
			Secret:    "XOJvQkcX6Og0CHFg+rQ27pqAB+EhSjVoI4Bs/JWegBc=",
		}
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "PATCH", req.Method)
			require.Equal(t, "/api/ddi/v1/keys/tsig/dummyKey36", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))
			var reqBody openapiclient.KeysTSIGKey
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, dummyKey, reqBody)

			response := openapiclient.KeysUpdateTSIGKeyResponse{
				Result: &dummyKey,
			}
			body, err := json.Marshal(response)
			require.NoError(t, err)
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		updateRequest := apiClient.TsigAPI.TsigUpdate(context.Background(), "dummyKey36").Body(dummyKey)
		resp, httpRes, err := updateRequest.Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TsigAPIService TsigDelete", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, "DELETE", req.Method)
			require.Equal(t, "/api/ddi/v1/keys/tsig/dummyKey36", req.URL.Path)
			return &http.Response{
				StatusCode: 204,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     make(http.Header),
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		httpRes, err := apiClient.TsigAPI.TsigDelete(context.Background(), "dummyKey36").Execute()
		require.Nil(t, err)
		require.Equal(t, 204, httpRes.StatusCode)
	})
}
