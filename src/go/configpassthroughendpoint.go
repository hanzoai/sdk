// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ConfigPassThroughEndpointService contains methods and other services that help
// with interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigPassThroughEndpointService] method instead.
type ConfigPassThroughEndpointService struct {
	Options []option.RequestOption
}

// NewConfigPassThroughEndpointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConfigPassThroughEndpointService(opts ...option.RequestOption) (r *ConfigPassThroughEndpointService) {
	r = &ConfigPassThroughEndpointService{}
	r.Options = opts
	return
}

// Create new pass-through endpoint
func (r *ConfigPassThroughEndpointService) New(ctx context.Context, body ConfigPassThroughEndpointNewParams, opts ...option.RequestOption) (res *ConfigPassThroughEndpointNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "config/pass_through_endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update a pass-through endpoint
func (r *ConfigPassThroughEndpointService) Update(ctx context.Context, endpointID string, opts ...option.RequestOption) (res *ConfigPassThroughEndpointUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpointID == "" {
		err = errors.New("missing required endpoint_id parameter")
		return
	}
	path := fmt.Sprintf("config/pass_through_endpoint/%s", endpointID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// GET configured pass through endpoint.
//
// If no endpoint_id given, return all configured endpoints.
func (r *ConfigPassThroughEndpointService) List(ctx context.Context, query ConfigPassThroughEndpointListParams, opts ...option.RequestOption) (res *PassThroughEndpointResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "config/pass_through_endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a pass-through endpoint
//
// Returns - the deleted endpoint
func (r *ConfigPassThroughEndpointService) Delete(ctx context.Context, body ConfigPassThroughEndpointDeleteParams, opts ...option.RequestOption) (res *PassThroughEndpointResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "config/pass_through_endpoint"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type PassThroughEndpointResponse struct {
	Endpoints []PassThroughGenericEndpoint    `json:"endpoints,required"`
	JSON      passThroughEndpointResponseJSON `json:"-"`
}

// passThroughEndpointResponseJSON contains the JSON metadata for the struct
// [PassThroughEndpointResponse]
type passThroughEndpointResponseJSON struct {
	Endpoints   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PassThroughEndpointResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r passThroughEndpointResponseJSON) RawJSON() string {
	return r.raw
}

type PassThroughGenericEndpoint struct {
	// Key-value pairs of headers to be forwarded with the request. You can set any key
	// value pair here and it will be forwarded to your target endpoint
	Headers interface{} `json:"headers,required"`
	// The route to be added to the LLM Proxy Server.
	Path string `json:"path,required"`
	// The URL to which requests for this path should be forwarded.
	Target string                         `json:"target,required"`
	JSON   passThroughGenericEndpointJSON `json:"-"`
}

// passThroughGenericEndpointJSON contains the JSON metadata for the struct
// [PassThroughGenericEndpoint]
type passThroughGenericEndpointJSON struct {
	Headers     apijson.Field
	Path        apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PassThroughGenericEndpoint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r passThroughGenericEndpointJSON) RawJSON() string {
	return r.raw
}

type PassThroughGenericEndpointParam struct {
	// Key-value pairs of headers to be forwarded with the request. You can set any key
	// value pair here and it will be forwarded to your target endpoint
	Headers param.Field[interface{}] `json:"headers,required"`
	// The route to be added to the LLM Proxy Server.
	Path param.Field[string] `json:"path,required"`
	// The URL to which requests for this path should be forwarded.
	Target param.Field[string] `json:"target,required"`
}

func (r PassThroughGenericEndpointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConfigPassThroughEndpointNewResponse = interface{}

type ConfigPassThroughEndpointUpdateResponse = interface{}

type ConfigPassThroughEndpointNewParams struct {
	PassThroughGenericEndpoint PassThroughGenericEndpointParam `json:"pass_through_generic_endpoint,required"`
}

func (r ConfigPassThroughEndpointNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PassThroughGenericEndpoint)
}

type ConfigPassThroughEndpointListParams struct {
	EndpointID param.Field[string] `query:"endpoint_id"`
}

// URLQuery serializes [ConfigPassThroughEndpointListParams]'s query parameters as
// `url.Values`.
func (r ConfigPassThroughEndpointListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConfigPassThroughEndpointDeleteParams struct {
	EndpointID param.Field[string] `query:"endpoint_id,required"`
}

// URLQuery serializes [ConfigPassThroughEndpointDeleteParams]'s query parameters
// as `url.Values`.
func (r ConfigPassThroughEndpointDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
