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

// AzureService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAzureService] method instead.
type AzureService struct {
	Options []option.RequestOption
}

// NewAzureService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAzureService(opts ...option.RequestOption) (r *AzureService) {
	r = &AzureService{}
	r.Options = opts
	return
}

// Call any azure endpoint using the proxy.
//
// Just use `{PROXY_BASE_URL}/azure/{endpoint:path}`
func (r *AzureService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AzureNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("azure/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Call any azure endpoint using the proxy.
//
// Just use `{PROXY_BASE_URL}/azure/{endpoint:path}`
func (r *AzureService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AzureUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("azure/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Call any azure endpoint using the proxy.
//
// Just use `{PROXY_BASE_URL}/azure/{endpoint:path}`
func (r *AzureService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AzureDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("azure/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Call any azure endpoint using the proxy.
//
// Just use `{PROXY_BASE_URL}/azure/{endpoint:path}`
func (r *AzureService) Call(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AzureCallResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("azure/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Call any azure endpoint using the proxy.
//
// Just use `{PROXY_BASE_URL}/azure/{endpoint:path}`
func (r *AzureService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AzurePatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("azure/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type AzureNewResponse = interface{}

type AzureUpdateResponse = interface{}

type AzureDeleteResponse = interface{}

type AzureCallResponse = interface{}

type AzurePatchResponse = interface{}
