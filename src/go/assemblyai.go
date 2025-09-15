// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// AssemblyaiService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssemblyaiService] method instead.
type AssemblyaiService struct {
	Options []option.RequestOption
}

// NewAssemblyaiService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssemblyaiService(opts ...option.RequestOption) (r *AssemblyaiService) {
	r = &AssemblyaiService{}
	r.Options = opts
	return
}

// Assemblyai Proxy Route
func (r *AssemblyaiService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AssemblyaiNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *AssemblyaiService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AssemblyaiGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *AssemblyaiService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AssemblyaiUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *AssemblyaiService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AssemblyaiDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *AssemblyaiService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AssemblyaiPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type AssemblyaiNewResponse = interface{}

type AssemblyaiGetResponse = interface{}

type AssemblyaiUpdateResponse = interface{}

type AssemblyaiDeleteResponse = interface{}

type AssemblyaiPatchResponse = interface{}
