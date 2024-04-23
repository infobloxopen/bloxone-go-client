package option

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

func TestWithCSPUrl(t *testing.T) {
	config := &internal.Configuration{}
	url := "http://test.com"
	opt := WithCSPUrl(url)
	opt(config)
	assert.Equal(t, url, config.CSPURL)
}

func TestWithAPIKey(t *testing.T) {
	config := &internal.Configuration{}
	apiKey := "testKey"
	opt := WithAPIKey(apiKey)
	opt(config)
	assert.Equal(t, apiKey, config.APIKey)
}

func TestWithHTTPClient(t *testing.T) {
	config := &internal.Configuration{}
	client := &http.Client{}
	opt := WithHTTPClient(client)
	opt(config)
	assert.Equal(t, client, config.HTTPClient)
}

func TestWithDefaultTags(t *testing.T) {
	config := &internal.Configuration{}
	tags := map[string]string{"tag1": "value1"}
	opt := WithDefaultTags(tags)
	opt(config)
	assert.Equal(t, tags, config.DefaultTags)
}

func TestWithClientName(t *testing.T) {
	config := &internal.Configuration{}
	name := "testClient"
	opt := WithClientName(name)
	opt(config)
	assert.Equal(t, name, config.ClientName)
}

func TestWithDebug(t *testing.T) {
	config := &internal.Configuration{}
	opt := WithDebug(true)
	opt(config)
	assert.Equal(t, true, config.Debug)
}
