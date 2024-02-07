package restapiservicelegacy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	senzingRestServiceSingleton RestApiServiceLegacy
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) RestApiServiceLegacy {
	if senzingRestServiceSingleton == nil {
		senzingRestServiceSingleton = &RestApiServiceLegacyImpl{}
	}
	return senzingRestServiceSingleton
}

func testError(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------
