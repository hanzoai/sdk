// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"reflect"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
	"github.com/hanzoai/go-sdk/shared"
	"github.com/tidwall/gjson"
)

// GuardrailService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGuardrailService] method instead.
type GuardrailService struct {
	Options []option.RequestOption
}

// NewGuardrailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGuardrailService(opts ...option.RequestOption) (r *GuardrailService) {
	r = &GuardrailService{}
	r.Options = opts
	return
}

// List the guardrails that are available on the proxy server
//
// 👉 [Guardrail docs](https://docs.hanzo.ai/docs/proxy/guardrails/quick_start)
//
// Example Request:
//
// ```bash
// curl -X GET "http://localhost:4000/guardrails/list" -H "Authorization: Bearer <your_api_key>"
// ```
//
// Example Response:
//
// ```json
//
//	{
//	  "guardrails": [
//	    {
//	      "guardrail_name": "bedrock-pre-guard",
//	      "guardrail_info": {
//	        "params": [
//	          {
//	            "name": "toxicity_score",
//	            "type": "float",
//	            "description": "Score between 0-1 indicating content toxicity level"
//	          },
//	          {
//	            "name": "pii_detection",
//	            "type": "boolean"
//	          }
//	        ]
//	      }
//	    }
//	  ]
//	}
//
// ```
func (r *GuardrailService) List(ctx context.Context, opts ...option.RequestOption) (res *GuardrailListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "guardrails/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type GuardrailListResponse struct {
	Guardrails []GuardrailListResponseGuardrail `json:"guardrails,required"`
	JSON       guardrailListResponseJSON        `json:"-"`
}

// guardrailListResponseJSON contains the JSON metadata for the struct
// [GuardrailListResponse]
type guardrailListResponseJSON struct {
	Guardrails  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuardrailListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseJSON) RawJSON() string {
	return r.raw
}

type GuardrailListResponseGuardrail struct {
	GuardrailInfo interface{} `json:"guardrail_info,required,nullable"`
	GuardrailName string      `json:"guardrail_name,required"`
	// The returned LLM Params object for /guardrails/list
	LlmParams GuardrailListResponseGuardrailsLlmParams `json:"llm_params,required"`
	JSON      guardrailListResponseGuardrailJSON       `json:"-"`
}

// guardrailListResponseGuardrailJSON contains the JSON metadata for the struct
// [GuardrailListResponseGuardrail]
type guardrailListResponseGuardrailJSON struct {
	GuardrailInfo apijson.Field
	GuardrailName apijson.Field
	LlmParams     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GuardrailListResponseGuardrail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseGuardrailJSON) RawJSON() string {
	return r.raw
}

// The returned LLM Params object for /guardrails/list
type GuardrailListResponseGuardrailsLlmParams struct {
	Guardrail string                                            `json:"guardrail,required"`
	Mode      GuardrailListResponseGuardrailsLlmParamsModeUnion `json:"mode,required"`
	DefaultOn bool                                              `json:"default_on"`
	JSON      guardrailListResponseGuardrailsLlmParamsJSON      `json:"-"`
}

// guardrailListResponseGuardrailsLlmParamsJSON contains the JSON metadata for the
// struct [GuardrailListResponseGuardrailsLlmParams]
type guardrailListResponseGuardrailsLlmParamsJSON struct {
	Guardrail   apijson.Field
	Mode        apijson.Field
	DefaultOn   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GuardrailListResponseGuardrailsLlmParams) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r guardrailListResponseGuardrailsLlmParamsJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or
// [GuardrailListResponseGuardrailsLlmParamsModeArray].
type GuardrailListResponseGuardrailsLlmParamsModeUnion interface {
	ImplementsGuardrailListResponseGuardrailsLlmParamsModeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GuardrailListResponseGuardrailsLlmParamsModeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GuardrailListResponseGuardrailsLlmParamsModeArray{}),
		},
	)
}

type GuardrailListResponseGuardrailsLlmParamsModeArray []string

func (r GuardrailListResponseGuardrailsLlmParamsModeArray) ImplementsGuardrailListResponseGuardrailsLlmParamsModeUnion() {
}
