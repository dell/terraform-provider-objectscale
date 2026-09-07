package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httputil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"terraform-provider-objectscale/internal/clientgen"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Client type is to hold objectscale client.
type Client struct {
	GenClient *clientgen.APIClient
}

// NewClient returns the objectscale client.
func NewClient(endpoint string, username string, password string, insecure bool, timeout int64) (*Client, error) {
	genClient, err := newClientGen(context.Background(), endpoint, username, password, insecure, timeout)
	if err != nil {
		return nil, fmt.Errorf("cannot create client: %w", err)
	}

	var client = Client{
		GenClient: genClient,
	}
	return &client, nil
}

// newClientGen returns the generated objectscale client.
func newClientGen(ctx context.Context, endpoint string, username string, password string, insecure bool, timeout int64) (*clientgen.APIClient, error) {

	// Setup a User-Agent for your API client (replace the provider name for yours):
	userAgent := "terraform-objectscale-provider/1.0.0"
	jar, err := cookiejar.New(nil)
	if err != nil {
		tflog.Error(ctx, "Got error while creating cookie jar")
	}

	httpclient := &http.Client{
		Timeout: (time.Duration(timeout) * time.Second),
		Jar:     jar,
	}
	if insecure {
		/* #nosec */
		httpclient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				InsecureSkipVerify: true,
			},
		}
	} else {
		// Loading system certs by default if insecure is set to false
		pool, err := x509.SystemCertPool()
		if err != nil {
			errSysCerts := errors.New("unable to initialize cert pool from system")
			return nil, errSysCerts
		}
		httpclient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				MinVersion:         tls.VersionTLS12,
				RootCAs:            pool,
				InsecureSkipVerify: false,
			},
		}
	}

	url, _ := strings.CutSuffix(endpoint, "/")
	basicAuthString := basicAuth(username, password)

	// Debug mode is disabled by default for security. Enable via TF_OBJECTSCALE_DEBUG=true environment variable.
	debug, _ := strconv.ParseBool(os.Getenv("TF_OBJECTSCALE_DEBUG"))

	// Wrap transport with custom debug logging that redacts sensitive headers
	if debug {
		httpclient.Transport = &debugTransport{
			transport: httpclient.Transport,
			debug:     true,
		}
	}

	cfg := &clientgen.Configuration{
		HTTPClient: httpclient,
		// Host:          url,
		DefaultHeader: make(map[string]string),
		UserAgent:     userAgent,
		Debug:         false, // Always false - custom transport handles debug logging with redaction
		Servers: clientgen.ServerConfigurations{
			{
				URL:         url,
				Description: url,
			},
		},
		OperationServers: map[string]clientgen.ServerConfigurations{},
	}
	cfg.AddDefaultHeader("Authorization", "Basic "+basicAuthString)

	apiClient := clientgen.NewAPIClient(cfg)

	_, resp, err := apiClient.AuthenticationApi.AuthenticationResourceGetLoginToken(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("error during login: %w", err)
	}

	// get the X-SDS-AUTH-TOKEN header from the response
	token := resp.Header.Get("X-SDS-AUTH-TOKEN")
	if len(token) != 0 {
		cfg.AddDefaultHeader("X-SDS-AUTH-TOKEN", token)
		apiClient = clientgen.NewAPIClient(cfg)
	} else {
		return nil, errors.New("no token returned during login")
	}

	return apiClient, nil

}

// Generate the base 64 Authorization string from username / password.
func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// debugTransport wraps an http.RoundTripper and adds debug logging with sensitive header redaction
type debugTransport struct {
	transport http.RoundTripper
	debug     bool
}

func (t *debugTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err == nil {
			redacted := replaceSensitiveHeadersObjectScale(dump)
			log.Printf("\n%s\n", redacted)
		}
	}

	resp, err := t.transport.RoundTrip(req)
	if err != nil {
		return resp, err
	}

	if t.debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err == nil {
			redacted := replaceSensitiveHeadersObjectScale(dump)
			log.Printf("\n%s\n", redacted)
		}
	}
	return resp, err
}

// replaceSensitiveHeadersObjectScale redacts sensitive headers from HTTP dumps for ObjectScale
func replaceSensitiveHeadersObjectScale(dump []byte) string {
	// Redact Authorization, X-SDS-AUTH-TOKEN, and Cookie headers
	sensitiveDataRegexp := regexp.MustCompile(
		`(?m)(Authorization: |X-SDS-AUTH-TOKEN: )([^\n]+)|(auth_cookie=)([^;]+)`)
	return sensitiveDataRegexp.ReplaceAllString(string(dump), "$1$3******")
}
