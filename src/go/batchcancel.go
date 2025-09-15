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

// BatchCancelService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBatchCancelService] method instead.
type BatchCancelService struct {
	Options []option.RequestOption
}

// NewBatchCancelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBatchCancelService(opts ...option.RequestOption) (r *BatchCancelService) {
	r = &BatchCancelService{}
	r.Options = opts
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
func (r *BatchCancelService) Cancel(ctx context.Context, batchID string, body BatchCancelCancelParams, opts ...option.RequestOption) (res *BatchCancelCancelResponse, err error) {
	opts = append(r.Options[:], opts...)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("batches/%s/cancel", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BatchCancelCancelResponse = interface{}

type BatchCancelCancelParams struct {
	Provider param.Field[string] `query:"provider"`
}

// URLQuery serializes [BatchCancelCancelParams]'s query parameters as
// `url.Values`.
func (r BatchCancelCancelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
