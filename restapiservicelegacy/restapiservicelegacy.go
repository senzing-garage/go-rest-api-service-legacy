package restapiservicelegacy

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"sync"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// RestApiServiceLegacyImpl is...
type RestApiServiceLegacyImpl struct {
	JarFile         string
	ProxyTemplate   string
	CustomTransport http.RoundTripper
}

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

var (
	restApiServiceSyncOnce sync.Once
)

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

// Handle an HTTP request
func (restApiServiceLegacyImpl *RestApiServiceLegacyImpl) handleRequest(w http.ResponseWriter, r *http.Request) {

	// Create a new HTTP request to the proxied server with the same method and body as the original request.

	proxyUrl := fmt.Sprintf(restApiServiceLegacyImpl.ProxyTemplate, r.URL)
	proxyReq, err := http.NewRequest(r.Method, proxyUrl, r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	// Copy the headers from the original request to the proxy request.

	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	// Send the proxy request using the custom transport.

	resp, err := restApiServiceLegacyImpl.CustomTransport.RoundTrip(proxyReq)
	if err != nil {
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	// Copy the headers from the proxy response to the original response.

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Set the status code of the original response to the status code of the proxy response.

	w.WriteHeader(resp.StatusCode)

	// Copy the body of the proxy response to the original response.

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Could not copy HTTP body", http.StatusInternalServerError)
		return
	}
}

// Run Java jar file at most once.
func (restApiServiceLegacyImpl *RestApiServiceLegacyImpl) runJava() {
	restApiServiceSyncOnce.Do(func() {
		cmd, err := exec.Command("java", "-jar", restApiServiceLegacyImpl.JarFile).CombinedOutput() // #nosec 204
		if err != nil {
			panic(fmt.Sprintf("%v: %v", cmd, err))
		}
	})
}

// ----------------------------------------------------------------------------
// Public methods
// ----------------------------------------------------------------------------

func (restApiServiceLegacyImpl *RestApiServiceLegacyImpl) Handler(ctx context.Context) *http.ServeMux {

	// Run Java jar file.

	go restApiServiceLegacyImpl.runJava()

	// Proxy HTTP requests.

	rootMux := http.NewServeMux()
	rootMux.HandleFunc("/", restApiServiceLegacyImpl.handleRequest)
	return rootMux
}
