/*
DNS Configuration API

Testing CacheFlushAPIService

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

var ConfigCacheFlushPost = openapiclient.ConfigCacheFlush{
	FlushSubdomains: openapiclient.PtrBool(true),
	Fqdn:            openapiclient.PtrString("dummyFqdn"),
	Ophid:           openapiclient.PtrString("dummyOphid"),
	ServiceId:       openapiclient.PtrString("dummyServiceId"),
	Ttl:             openapiclient.PtrInt64(120),
	ViewName:        openapiclient.PtrString("dummyViewName"),
}

func TestCacheFlushAPIService(t *testing.T) {

	t.Run("Test CacheFlushAPIService CacheFlushCreate", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/dns/cache_flush", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.ConfigCacheFlush
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, ConfigCacheFlushPost, reqBody)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte{})),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		_, httpRes, err := apiClient.CacheFlushAPI.CacheFlushCreate(context.Background()).Body(ConfigCacheFlushPost).Execute()
		require.Nil(t, err)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
