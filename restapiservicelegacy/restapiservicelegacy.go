package restapiservicelegacy

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os/exec"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SenzingRestServiceImpl is...
type RestApiServiceLegacyImpl struct {
	UrlRoutePrefix string
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// internal methods
// ----------------------------------------------------------------------------

var customTransport = http.DefaultTransport

func init() {
	// Here, you can customize the transport, e.g., set timeouts or enable/disable keep-alive
}

func runJava() {
	// os.Setenv("PATH", "/home/sdk/jdk-11.0.16/bin:/home/temp/jtreg/bin:$PATH")
	cmd, err := exec.Command("java", "-jar", "/tmp/senzing-poc-server.jar").CombinedOutput()
	if err != nil {
		fmt.Println(">>>>>> Java Error:", err)
	}
	fmt.Println(">>>>>>>>>>>>>>", string(cmd))
}

// --- xxxxxx -----------------------------------------------------------------

func handleRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Printf(">>>>>> r = %+v\n", r)

	// Create a new HTTP request with the same method, URL, and body as the original request

	fmt.Printf(">>>>>> r.URL = %+v\n", r.URL)

	proxyUrl := fmt.Sprintf("http://localhost:8250%s", r.URL)

	fmt.Printf(">>>>>> proxyURL = %s\n", proxyUrl)

	proxyReq, err := http.NewRequest(r.Method, proxyUrl, r.Body)
	if err != nil {
		http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
		return
	}

	// Copy the headers from the original request to the proxy request
	for name, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(name, value)
		}
	}

	// Send the proxy request using the custom transport
	resp, err := customTransport.RoundTrip(proxyReq)
	if err != nil {
		http.Error(w, "Error sending proxy request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copy the headers from the proxy response to the original response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Set the status code of the original response to the status code of the proxy response
	w.WriteHeader(resp.StatusCode)

	// Copy the body of the proxy response to the original response
	io.Copy(w, resp.Body)
}

// --- Services ---------------------------------------------------------------

func (restApiServiceLegacyImpl *RestApiServiceLegacyImpl) Handler(ctx context.Context) *http.ServeMux {
	rootMux := http.NewServeMux()

	// Start Java

	go runJava()

	// Proxy HTTP requests.

	rootMux.HandleFunc("/", handleRequest)
	return rootMux
}
