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

// EuAssemblyaiService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEuAssemblyaiService] method instead.
type EuAssemblyaiService struct {
	Options []option.RequestOption
}

// NewEuAssemblyaiService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEuAssemblyaiService(opts ...option.RequestOption) (r *EuAssemblyaiService) {
	r = &EuAssemblyaiService{}
	r.Options = opts
	return
}

// Assemblyai Proxy Route
func (r *EuAssemblyaiService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *EuAssemblyaiNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("eu.assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *EuAssemblyaiService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *EuAssemblyaiGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("eu.assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *EuAssemblyaiService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *EuAssemblyaiUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("eu.assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *EuAssemblyaiService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *EuAssemblyaiDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("eu.assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Assemblyai Proxy Route
func (r *EuAssemblyaiService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *EuAssemblyaiPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("eu.assemblyai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type EuAssemblyaiNewResponse = interface{}

type EuAssemblyaiGetResponse = interface{}

type EuAssemblyaiUpdateResponse = interface{}

type EuAssemblyaiDeleteResponse = interface{}

type EuAssemblyaiPatchResponse = interface{}
