/*
DDI Keys API

Testing UploadAPIService

*/

package keys_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/infobloxopen/bloxone-go-client/internal"
	openapiclient "github.com/infobloxopen/bloxone-go-client/keys"
)

var UploadRequestPost = openapiclient.UploadRequest{
	Comment: openapiclient.PtrString("Upload Request"),
	Content: "SGVsbG8gd29ybGQ=",                     // "Hello world" in base64
	Type:    openapiclient.UPLOADCONTENTTYPE_KEYTAB, // Replace "your_content_type" with the actual content type
	Tags:    make(map[string]interface{}),
}

func TestUploadAPIService(t *testing.T) {

	t.Run("Test UploadAPIService UploadUpload", func(t *testing.T) {
		configuration := internal.NewConfiguration()
		configuration.HTTPClient = internal.NewTestClient(func(req *http.Request) *http.Response {
			require.Equal(t, http.MethodPost, req.Method)
			require.Equal(t, "/api/ddi/v1/keys/upload", req.URL.Path)
			require.Equal(t, "application/json", req.Header.Get("Content-Type"))

			var reqBody openapiclient.UploadRequest
			require.NoError(t, json.NewDecoder(req.Body).Decode(&reqBody))
			require.Equal(t, UploadRequestPost, reqBody)

			response := openapiclient.DdiuploadResponse{}
			body, err := json.Marshal(response)
			require.NoError(t, err)

			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader(body)),
				Header:     map[string][]string{"Content-Type": {"application/json"}},
			}
		})
		apiClient := openapiclient.NewAPIClient(configuration)
		resp, httpRes, err := apiClient.UploadAPI.UploadUpload(context.Background()).Body(UploadRequestPost).Execute()
		require.Nil(t, err)
		require.NotNil(t, resp)
		require.Equal(t, http.StatusOK, httpRes.StatusCode)
	})

}
