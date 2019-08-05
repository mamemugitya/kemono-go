package e621

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

func NewMockServer() (*http.ServeMux, *url.URL) {
	mux := http.NewServeMux()
	testServer := httptest.NewServer(mux)
	testServerURL, _ := url.Parse(testServer.URL)
	return mux, testServerURL
}

func NewTestClient(testServerURL *url.URL) *Client {
	endpointURL := testServerURL.String()
	client, _ := NewClient(endpointURL, "test user", nil)
	return client
}
