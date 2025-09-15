// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// TestService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTestService] method instead.
type TestService struct {
	Options []option.RequestOption
}

// NewTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTestService(opts ...option.RequestOption) (r *TestService) {
	r = &TestService{}
	r.Options = opts
	return
}

// [DEPRECATED] use `/health/liveliness` instead.
//
// A test endpoint that pings the proxy server to check if it's healthy.
//
// Parameters: request (Request): The incoming request.
//
// Returns: dict: A dictionary containing the route of the request URL.
func (r *TestService) Ping(ctx context.Context, opts ...option.RequestOption) (res *TestPingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "test"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TestPingResponse = interface{}
