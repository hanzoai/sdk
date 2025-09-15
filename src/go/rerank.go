// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// RerankService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRerankService] method instead.
type RerankService struct {
	Options []option.RequestOption
}

// NewRerankService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRerankService(opts ...option.RequestOption) (r *RerankService) {
	r = &RerankService{}
	r.Options = opts
	return
}

// Rerank
func (r *RerankService) New(ctx context.Context, opts ...option.RequestOption) (res *RerankNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Rerank
func (r *RerankService) NewV1(ctx context.Context, opts ...option.RequestOption) (res *RerankNewV1Response, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Rerank
func (r *RerankService) NewV2(ctx context.Context, opts ...option.RequestOption) (res *RerankNewV2Response, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/rerank"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type RerankNewResponse = interface{}

type RerankNewV1Response = interface{}

type RerankNewV2Response = interface{}
