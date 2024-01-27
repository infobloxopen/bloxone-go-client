package client

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/infobloxopen/bloxone-go-client/internal"
)

func TestConfiguration_internal(t *testing.T) {
	type fields struct {
		ClientName  string
		CSPURL      string
		APIKey      string
		HTTPClient  *http.Client
		DefaultTags map[string]string
	}
	type args struct {
		basePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *internal.Configuration
		wantErr bool
	}{
		{
			"empty API key",
			fields{
				APIKey: "",
			},
			args{basePath: ""},
			nil,
			true,
		},
		{"empty clientName",
			fields{
				ClientName: "",
			},
			args{basePath: ""},
			nil,
			true,
		},
		{
			"empty DefaultTags",
			fields{
				ClientName: "terraform/v1.1#yug278872h",
				APIKey:     "12323455",
			},
			args{basePath: ""},
			&internal.Configuration{
				DefaultHeader: map[string]string{
					HeaderAuthorization: "Token 12323455",
					HeaderClient:        "terraform/v1.1#yug278872h",
					HeaderSDK:           sdkIdentifier,
				},
				Debug:      false,
				UserAgent:  fmt.Sprintf("bloxone-%s/%s", sdkIdentifier, version),
				Servers:    []internal.ServerConfiguration{{URL: "https://csp.infoblox.com"}},
				HTTPClient: http.DefaultClient,
				DefaultTags: map[string]string{
					clientIdentifier: "terraform",
				},
			},
			false,
		},
		{
			"DefaultTags provided",
			fields{
				CSPURL:     "https://csp.infoblox.com",
				ClientName: "terraformv1.1#yug278872h",
				APIKey:     "12323455",
				DefaultTags: map[string]string{
					"site": "A",
					"env":  "test",
				},
			},
			args{basePath: ""},
			&internal.Configuration{
				DefaultHeader: map[string]string{
					HeaderAuthorization: "Token 12323455",
					HeaderClient:        "terraformv1.1#yug278872h",
					HeaderSDK:           sdkIdentifier,
				},
				Debug:      false,
				UserAgent:  fmt.Sprintf("bloxone-%s/%s", sdkIdentifier, version),
				Servers:    []internal.ServerConfiguration{{URL: "https://csp.infoblox.com"}},
				HTTPClient: http.DefaultClient,
				DefaultTags: map[string]string{
					clientIdentifier: "terraformv1.1#yug278872h",
					"site":           "A",
					"env":            "test",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Configuration{
				ClientName:  tt.fields.ClientName,
				CSPURL:      tt.fields.CSPURL,
				APIKey:      tt.fields.APIKey,
				HTTPClient:  tt.fields.HTTPClient,
				DefaultTags: tt.fields.DefaultTags,
			}
			got, err := c.internal(tt.args.basePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("internal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("internal() got = %v, want %v", got, tt.want)
			}
		})
	}
}
