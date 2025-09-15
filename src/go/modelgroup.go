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

// ModelGroupService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelGroupService] method instead.
type ModelGroupService struct {
	Options []option.RequestOption
}

// NewModelGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelGroupService(opts ...option.RequestOption) (r *ModelGroupService) {
	r = &ModelGroupService{}
	r.Options = opts
	return
}

// Get information about all the deployments on llm proxy, including config.yaml
// descriptions (except api key and api base)
//
//   - /model_group/info returns all model groups. End users of proxy should use
//     /model_group/info since those models will be used for /chat/completions,
//     /embeddings, etc.
//   - /model_group/info?model_group=rerank-english-v3.0 returns all model groups for
//     a specific model group (`model_name` in config.yaml)
//
// Example Request (All Models):
//
// ```shell
// curl -X 'GET'     'http://localhost:4000/model_group/info'     -H 'accept: application/json'     -H 'x-api-key: sk-1234'
// ```
//
// Example Request (Specific Model Group):
//
// ```shell
// curl -X 'GET'     'http://localhost:4000/model_group/info?model_group=rerank-english-v3.0'     -H 'accept: application/json'     -H 'Authorization: Bearer sk-1234'
// ```
//
// Example Request (Specific Wildcard Model Group): (e.g. `model_name: openai/*` on
// config.yaml)
//
// ```shell
// curl -X 'GET'     'http://localhost:4000/model_group/info?model_group=openai/tts-1'
// -H 'accept: application/json'     -H 'Authorization: Bearersk-1234'
// ```
//
// Learn how to use and set wildcard models
// [here](https://docs.hanzo.ai/docs/wildcard_routing)
//
// Example Response:
//
// ```json
//
//	{
//	  "data": [
//	    {
//	      "model_group": "rerank-english-v3.0",
//	      "providers": ["cohere"],
//	      "max_input_tokens": null,
//	      "max_output_tokens": null,
//	      "input_cost_per_token": 0.0,
//	      "output_cost_per_token": 0.0,
//	      "mode": null,
//	      "tpm": null,
//	      "rpm": null,
//	      "supports_parallel_function_calling": false,
//	      "supports_vision": false,
//	      "supports_function_calling": false,
//	      "supported_openai_params": [
//	        "stream",
//	        "temperature",
//	        "max_tokens",
//	        "logit_bias",
//	        "top_p",
//	        "frequency_penalty",
//	        "presence_penalty",
//	        "stop",
//	        "n",
//	        "extra_headers"
//	      ]
//	    },
//	    {
//	      "model_group": "gpt-3.5-turbo",
//	      "providers": ["openai"],
//	      "max_input_tokens": 16385.0,
//	      "max_output_tokens": 4096.0,
//	      "input_cost_per_token": 1.5e-6,
//	      "output_cost_per_token": 2e-6,
//	      "mode": "chat",
//	      "tpm": null,
//	      "rpm": null,
//	      "supports_parallel_function_calling": false,
//	      "supports_vision": false,
//	      "supports_function_calling": true,
//	      "supported_openai_params": [
//	        "frequency_penalty",
//	        "logit_bias",
//	        "logprobs",
//	        "top_logprobs",
//	        "max_tokens",
//	        "max_completion_tokens",
//	        "n",
//	        "presence_penalty",
//	        "seed",
//	        "stop",
//	        "stream",
//	        "stream_options",
//	        "temperature",
//	        "top_p",
//	        "tools",
//	        "tool_choice",
//	        "function_call",
//	        "functions",
//	        "max_retries",
//	        "extra_headers",
//	        "parallel_tool_calls",
//	        "response_format"
//	      ]
//	    },
//	    {
//	      "model_group": "llava-hf",
//	      "providers": ["openai"],
//	      "max_input_tokens": null,
//	      "max_output_tokens": null,
//	      "input_cost_per_token": 0.0,
//	      "output_cost_per_token": 0.0,
//	      "mode": null,
//	      "tpm": null,
//	      "rpm": null,
//	      "supports_parallel_function_calling": false,
//	      "supports_vision": true,
//	      "supports_function_calling": false,
//	      "supported_openai_params": [
//	        "frequency_penalty",
//	        "logit_bias",
//	        "logprobs",
//	        "top_logprobs",
//	        "max_tokens",
//	        "max_completion_tokens",
//	        "n",
//	        "presence_penalty",
//	        "seed",
//	        "stop",
//	        "stream",
//	        "stream_options",
//	        "temperature",
//	        "top_p",
//	        "tools",
//	        "tool_choice",
//	        "function_call",
//	        "functions",
//	        "max_retries",
//	        "extra_headers",
//	        "parallel_tool_calls",
//	        "response_format"
//	      ]
//	    }
//	  ]
//	}
//
// ```
func (r *ModelGroupService) GetInfo(ctx context.Context, query ModelGroupGetInfoParams, opts ...option.RequestOption) (res *ModelGroupGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "model_group/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ModelGroupGetInfoResponse = interface{}

type ModelGroupGetInfoParams struct {
	ModelGroup param.Field[string] `query:"model_group"`
}

// URLQuery serializes [ModelGroupGetInfoParams]'s query parameters as
// `url.Values`.
func (r ModelGroupGetInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
