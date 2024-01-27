package client

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

const (
	ENVBloxOneCSPURL = "BLOXONE_CSP_URL"
	ENVBloxOneAPIKey = "BLOXONE_API_KEY"
)

const (
	HeaderClient        = "x-infoblox-client"
	HeaderSDK           = "x-infoblox-sdk"
	HeaderAuthorization = "Authorization"
)

const version = "0.1"
const sdkIdentifier = "golang-sdk"
const clientIdentifier = "automation"

// Configuration stores the configuration of the API client
type Configuration struct {
	// ClientName is the name of the client using the SDK.
	// Required.
	ClientName string

	// CSPURL is the URL for BloxOne Cloud Services Portal.
	// Can also be configured using the `BLOXONE_CSP_URL` environment variable.
	// Optional. Default is https://csp.infoblox.com
	CSPURL string

	// APIKey for accessing the BloxOne API.
	// Can also be configured by using the `BLOXONE_API_KEY` environment variable.
	// https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys
	// Required.
	APIKey string

	// HTTPClient to use for the SDK.
	// Optional. The default HTTPClient will be used if not provided.
	HTTPClient *http.Client

	// Default global tags the client can set for all requests.
	DefaultTags map[string]string
}

func (c Configuration) internal(basePath string) (*internal.Configuration, error) {
	cspURL := "https://csp.infoblox.com"
	if v, ok := os.LookupEnv(ENVBloxOneCSPURL); ok {
		cspURL = v
	}
	if len(c.CSPURL) > 0 {
		cspURL = c.CSPURL
	}
	cspURL = cspURL + basePath

	apiKey := ""
	if v, ok := os.LookupEnv(ENVBloxOneAPIKey); ok {
		apiKey = v
	}
	if len(c.APIKey) > 0 {
		apiKey = c.APIKey
	}
	if len(apiKey) == 0 {
		return nil, errors.New("APIKey is required")
	}

	if len(c.ClientName) == 0 {
		return nil, errors.New("ClientName is required")
	}

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	defaultHeaders := map[string]string{
		HeaderAuthorization: "Token " + apiKey,
		HeaderClient:        c.ClientName,
		HeaderSDK:           sdkIdentifier,
	}

	userAgent := fmt.Sprintf("bloxone-%s/%s", sdkIdentifier, version)

	ic := &internal.Configuration{
		DefaultHeader:    defaultHeaders,
		UserAgent:        userAgent,
		Debug:            false,
		OperationServers: nil,
		Servers:          []internal.ServerConfiguration{{URL: cspURL}},
		HTTPClient:       httpClient,
		DefaultTags:      make(map[string]string),
	}
	// Add default tags set
	if c.DefaultTags != nil {
		ic.AddDefaultTags(c.DefaultTags)
	}

	// setting up custom tag to identify the client
	dfTags := make(map[string]string)
	// Extract client from ClientName string
	// Format: <client>/version#commit
	dfTags[clientIdentifier] = strings.Split(c.ClientName, "/")[0]
	ic.AddDefaultTags(dfTags)

	return ic, nil
}
