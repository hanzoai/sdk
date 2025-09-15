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

// ModelInfoService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelInfoService] method instead.
type ModelInfoService struct {
	Options []option.RequestOption
}

// NewModelInfoService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelInfoService(opts ...option.RequestOption) (r *ModelInfoService) {
	r = &ModelInfoService{}
	r.Options = opts
	return
}

// Provides more info about each model in /models, including config.yaml
// descriptions (except api key and api base)
//
// Parameters: llm_model_id: Optional[str] = None (this is the value of
// `x-llm-model-id` returned in response headers)
//
//   - When llm_model_id is passed, it will return the info for that specific model
//   - When llm_model_id is not passed, it will return the info for all models
//
// Returns: Returns a dictionary containing information about each model.
//
// Example Response:
//
// ```json
//
//	{
//	  "data": [
//	    {
//	      "model_name": "fake-openai-endpoint",
//	      "llm_params": {
//	        "api_base": "https://exampleopenaiendpoint-production.up.railway.app/",
//	        "model": "openai/fake"
//	      },
//	      "model_info": {
//	        "id": "112f74fab24a7a5245d2ced3536dd8f5f9192c57ee6e332af0f0512e08bed5af",
//	        "db_model": false
//	      }
//	    }
//	  ]
//	}
//
// ```
func (r *ModelInfoService) List(ctx context.Context, query ModelInfoListParams, opts ...option.RequestOption) (res *ModelInfoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "model/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ModelInfoListResponse = interface{}

type ModelInfoListParams struct {
	LlmModelID param.Field[string] `query:"llm_model_id"`
}

// URLQuery serializes [ModelInfoListParams]'s query parameters as `url.Values`.
func (r ModelInfoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
