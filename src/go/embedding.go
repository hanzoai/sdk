// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// EmbeddingService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbeddingService] method instead.
type EmbeddingService struct {
	Options []option.RequestOption
}

// NewEmbeddingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmbeddingService(opts ...option.RequestOption) (r *EmbeddingService) {
	r = &EmbeddingService{}
	r.Options = opts
	return
}

// Follows the exact same API spec as
// `OpenAI's Embeddings API https://platform.openai.com/docs/api-reference/embeddings`
//
// ```bash
// curl -X POST http://localhost:4000/v1/embeddings
// -H "Content-Type: application/json"
// -H "Authorization: Bearer sk-1234"
//
//	-d '{
//	    "model": "text-embedding-ada-002",
//	    "input": "The quick brown fox jumps over the lazy dog"
//	}'
//
// ```
func (r *EmbeddingService) New(ctx context.Context, body EmbeddingNewParams, opts ...option.RequestOption) (res *EmbeddingNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "embeddings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type EmbeddingNewResponse = interface{}

type EmbeddingNewParams struct {
	Model param.Field[string] `query:"model"`
}

// URLQuery serializes [EmbeddingNewParams]'s query parameters as `url.Values`.
func (r EmbeddingNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
