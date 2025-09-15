// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// CacheRediService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCacheRediService] method instead.
type CacheRediService struct {
	Options []option.RequestOption
}

// NewCacheRediService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCacheRediService(opts ...option.RequestOption) (r *CacheRediService) {
	r = &CacheRediService{}
	r.Options = opts
	return
}

// Endpoint for getting /redis/info
func (r *CacheRediService) GetInfo(ctx context.Context, opts ...option.RequestOption) (res *CacheRediGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cache/redis/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CacheRediGetInfoResponse = interface{}
