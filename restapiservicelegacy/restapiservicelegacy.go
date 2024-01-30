package restapiservicelegacy

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"os"
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

// --- Logging ----------------------------------------------------------------

// --- Errors -----------------------------------------------------------------

// --- Services ---------------------------------------------------------------

func (restApiServiceLegacyImpl *RestApiServiceLegacyImpl) Handler(ctx context.Context) *http.ServeMux {
	rootMux := http.NewServeMux()

	// Start Java

	os.Setenv("PATH", "/home/sdk/jdk-11.0.16/bin:/home/temp/jtreg/bin:$PATH")
	cmd, err := exec.Command("java", "-version").CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(cmd))

	fmt.Print(">>>>>>>>>>>>>>>>>>>>>> HERE")

	return rootMux
}
