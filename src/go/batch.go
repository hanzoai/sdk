// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// BatchService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchService] method instead.
type BatchService struct {
	Options []option.RequestOption
	Cancel  *BatchCancelService
}

// NewBatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBatchService(opts ...option.RequestOption) (r *BatchService) {
	r = &BatchService{}
	r.Options = opts
	r.Cancel = NewBatchCancelService(opts...)
	return
}

// Create large batches of API requests for asynchronous processing. This is the
// equivalent of POST https://api.openai.com/v1/batch Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/batch
//
// # Example Curl
//
// ```
//
//	curl http://localhost:4000/v1/batches         -H "Authorization: Bearer sk-1234"         -H "Content-Type: application/json"         -d '{
//	        "input_file_id": "file-abc123",
//	        "endpoint": "/v1/chat/completions",
//	        "completion_window": "24h"
//	}'
//
// ```
func (r *BatchService) New(ctx context.Context, body BatchNewParams, opts ...option.RequestOption) (res *BatchNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a batch. This is the equivalent of GET
// https://api.openai.com/v1/batches/{batch_id} Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/batch/retrieve
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/batches/batch_abc123     -H "Authorization: Bearer sk-1234"     -H "Content-Type: application/json"
// ```
func (r *BatchService) Get(ctx context.Context, batchID string, query BatchGetParams, opts ...option.RequestOption) (res *BatchGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/batches/%s", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Lists This is the equivalent of GET https://api.openai.com/v1/batches/ Supports
// Identical Params as: https://platform.openai.com/docs/api-reference/batch/list
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/batches?limit=2     -H "Authorization: Bearer sk-1234"     -H "Content-Type: application/json"
// ```
func (r *BatchService) List(ctx context.Context, query BatchListParams, opts ...option.RequestOption) (res *BatchListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/batches"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancel a batch. This is the equivalent of POST
// https://api.openai.com/v1/batches/{batch_id}/cancel
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/batch/cancel
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/batches/batch_abc123/cancel         -H "Authorization: Bearer sk-1234"         -H "Content-Type: application/json"         -X POST
//
// ```
func (r *BatchService) CancelWithProvider(ctx context.Context, provider string, batchID string, opts ...option.RequestOption) (res *BatchCancelWithProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/batches/%s/cancel", provider, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create large batches of API requests for asynchronous processing. This is the
// equivalent of POST https://api.openai.com/v1/batch Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/batch
//
// # Example Curl
//
// ```
//
//	curl http://localhost:4000/v1/batches         -H "Authorization: Bearer sk-1234"         -H "Content-Type: application/json"         -d '{
//	        "input_file_id": "file-abc123",
//	        "endpoint": "/v1/chat/completions",
//	        "completion_window": "24h"
//	}'
//
// ```
func (r *BatchService) NewWithProvider(ctx context.Context, provider string, opts ...option.RequestOption) (res *BatchNewWithProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/batches", provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Lists This is the equivalent of GET https://api.openai.com/v1/batches/ Supports
// Identical Params as: https://platform.openai.com/docs/api-reference/batch/list
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/batches?limit=2     -H "Authorization: Bearer sk-1234"     -H "Content-Type: application/json"
// ```
func (r *BatchService) ListWithProvider(ctx context.Context, provider string, query BatchListWithProviderParams, opts ...option.RequestOption) (res *BatchListWithProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/batches", provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves a batch. This is the equivalent of GET
// https://api.openai.com/v1/batches/{batch_id} Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/batch/retrieve
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/batches/batch_abc123     -H "Authorization: Bearer sk-1234"     -H "Content-Type: application/json"
// ```
func (r *BatchService) GetWithProvider(ctx context.Context, provider string, batchID string, opts ...option.RequestOption) (res *BatchGetWithProviderResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/batches/%s", provider, batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BatchNewResponse = interface{}

type BatchGetResponse = interface{}

type BatchListResponse = interface{}

type BatchCancelWithProviderResponse = interface{}

type BatchNewWithProviderResponse = interface{}

type BatchListWithProviderResponse = interface{}

type BatchGetWithProviderResponse = interface{}

type BatchNewParams struct {
	Provider param.Field[string] `query:"provider"`
}

// URLQuery serializes [BatchNewParams]'s query parameters as `url.Values`.
func (r BatchNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchGetParams struct {
	Provider param.Field[string] `query:"provider"`
}

// URLQuery serializes [BatchGetParams]'s query parameters as `url.Values`.
func (r BatchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchListParams struct {
	After    param.Field[string] `query:"after"`
	Limit    param.Field[int64]  `query:"limit"`
	Provider param.Field[string] `query:"provider"`
}

// URLQuery serializes [BatchListParams]'s query parameters as `url.Values`.
func (r BatchListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type BatchListWithProviderParams struct {
	After param.Field[string] `query:"after"`
	Limit param.Field[int64]  `query:"limit"`
}

// URLQuery serializes [BatchListWithProviderParams]'s query parameters as
// `url.Values`.
func (r BatchListWithProviderParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
