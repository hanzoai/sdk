// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// CacheService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCacheService] method instead.
type CacheService struct {
	Options []option.RequestOption
	Redis   *CacheRediService
}

// NewCacheService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCacheService(opts ...option.RequestOption) (r *CacheService) {
	r = &CacheService{}
	r.Options = opts
	r.Redis = NewCacheRediService(opts...)
	return
}

// Endpoint for deleting a key from the cache. All responses from llm proxy have
// `x-llm-cache-key` in the headers
//
// Parameters:
//
//   - **keys**: _Optional[List[str]]_ - A list of keys to delete from the cache.
//     Example {"keys": ["key1", "key2"]}
//
// ```shell
// curl -X POST "http://0.0.0.0:4000/cache/delete"     -H "Authorization: Bearer sk-1234"     -d '{"keys": ["key1", "key2"]}'
// ```
func (r *CacheService) Delete(ctx context.Context, opts ...option.RequestOption) (res *CacheDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cache/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// A function to flush all items from the cache. (All items will be deleted from
// the cache with this) Raises HTTPException if the cache is not initialized or if
// the cache type does not support flushing. Returns a dictionary with the status
// of the operation.
//
// Usage:
//
// ```
// curl -X POST http://0.0.0.0:4000/cache/flushall -H "Authorization: Bearer sk-1234"
// ```
func (r *CacheService) FlushAll(ctx context.Context, opts ...option.RequestOption) (res *CacheFlushAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cache/flushall"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Endpoint for checking if cache can be pinged
func (r *CacheService) Ping(ctx context.Context, opts ...option.RequestOption) (res *CachePingResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "cache/ping"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CacheDeleteResponse = interface{}

type CacheFlushAllResponse = interface{}

type CachePingResponse struct {
	CacheType              string                `json:"cache_type,required"`
	Status                 string                `json:"status,required"`
	HealthCheckCacheParams interface{}           `json:"health_check_cache_params,nullable"`
	LlmCacheParams         string                `json:"llm_cache_params,nullable"`
	PingResponse           bool                  `json:"ping_response,nullable"`
	SetCacheResponse       string                `json:"set_cache_response,nullable"`
	JSON                   cachePingResponseJSON `json:"-"`
}

// cachePingResponseJSON contains the JSON metadata for the struct
// [CachePingResponse]
type cachePingResponseJSON struct {
	CacheType              apijson.Field
	Status                 apijson.Field
	HealthCheckCacheParams apijson.Field
	LlmCacheParams         apijson.Field
	PingResponse           apijson.Field
	SetCacheResponse       apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *CachePingResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cachePingResponseJSON) RawJSON() string {
	return r.raw
}
