// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ModelService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
	Info    *ModelInfoService
	Update  *ModelUpdateService
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	r.Info = NewModelInfoService(opts...)
	r.Update = NewModelUpdateService(opts...)
	return
}

// Allows adding new models to the model list in the config.yaml
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *ModelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "model/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Allows deleting models in the model list in the config.yaml
func (r *ModelService) Delete(ctx context.Context, body ModelDeleteParams, opts ...option.RequestOption) (res *ModelDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "model/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ConfigurableClientsideParamsCustomAuth struct {
	APIBase param.Field[string] `json:"api_base,required"`
}

func (r ConfigurableClientsideParamsCustomAuth) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ConfigurableClientsideParamsCustomAuth) ImplementsModelNewParamsLlmParamsConfigurableClientsideAuthParamUnion() {
}

func (r ConfigurableClientsideParamsCustomAuth) ImplementsUpdateDeploymentLlmParamsConfigurableClientsideAuthParamsUnionParam() {
}

type ModelInfoParam struct {
	ID                  param.Field[string]        `json:"id,required"`
	BaseModel           param.Field[string]        `json:"base_model"`
	CreatedAt           param.Field[time.Time]     `json:"created_at" format:"date-time"`
	CreatedBy           param.Field[string]        `json:"created_by"`
	DBModel             param.Field[bool]          `json:"db_model"`
	TeamID              param.Field[string]        `json:"team_id"`
	TeamPublicModelName param.Field[string]        `json:"team_public_model_name"`
	Tier                param.Field[ModelInfoTier] `json:"tier"`
	UpdatedAt           param.Field[time.Time]     `json:"updated_at" format:"date-time"`
	UpdatedBy           param.Field[string]        `json:"updated_by"`
	ExtraFields         map[string]interface{}     `json:"-,extras"`
}

func (r ModelInfoParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ModelInfoTier string

const (
	ModelInfoTierFree ModelInfoTier = "free"
	ModelInfoTierPaid ModelInfoTier = "paid"
)

func (r ModelInfoTier) IsKnown() bool {
	switch r {
	case ModelInfoTierFree, ModelInfoTierPaid:
		return true
	}
	return false
}

type ModelNewResponse = interface{}

type ModelDeleteResponse = interface{}

type ModelNewParams struct {
	// LLM Params with 'model' requirement - used for completions
	LlmParams param.Field[ModelNewParamsLlmParams] `json:"llm_params,required"`
	ModelInfo param.Field[ModelInfoParam]          `json:"model_info,required"`
	ModelName param.Field[string]                  `json:"model_name,required"`
}

func (r ModelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// LLM Params with 'model' requirement - used for completions
type ModelNewParamsLlmParams struct {
	Model                            param.Field[string]                                                        `json:"model,required"`
	APIBase                          param.Field[string]                                                        `json:"api_base"`
	APIKey                           param.Field[string]                                                        `json:"api_key"`
	APIVersion                       param.Field[string]                                                        `json:"api_version"`
	AwsAccessKeyID                   param.Field[string]                                                        `json:"aws_access_key_id"`
	AwsRegionName                    param.Field[string]                                                        `json:"aws_region_name"`
	AwsSecretAccessKey               param.Field[string]                                                        `json:"aws_secret_access_key"`
	BudgetDuration                   param.Field[string]                                                        `json:"budget_duration"`
	ConfigurableClientsideAuthParams param.Field[[]ModelNewParamsLlmParamsConfigurableClientsideAuthParamUnion] `json:"configurable_clientside_auth_params"`
	CustomLlmProvider                param.Field[string]                                                        `json:"custom_llm_provider"`
	InputCostPerSecond               param.Field[float64]                                                       `json:"input_cost_per_second"`
	InputCostPerToken                param.Field[float64]                                                       `json:"input_cost_per_token"`
	LlmTraceID                       param.Field[string]                                                        `json:"llm_trace_id"`
	MaxBudget                        param.Field[float64]                                                       `json:"max_budget"`
	MaxFileSizeMB                    param.Field[float64]                                                       `json:"max_file_size_mb"`
	MaxRetries                       param.Field[int64]                                                         `json:"max_retries"`
	MergeReasoningContentInChoices   param.Field[bool]                                                          `json:"merge_reasoning_content_in_choices"`
	ModelInfo                        param.Field[interface{}]                                                   `json:"model_info"`
	Organization                     param.Field[string]                                                        `json:"organization"`
	OutputCostPerSecond              param.Field[float64]                                                       `json:"output_cost_per_second"`
	OutputCostPerToken               param.Field[float64]                                                       `json:"output_cost_per_token"`
	RegionName                       param.Field[string]                                                        `json:"region_name"`
	Rpm                              param.Field[int64]                                                         `json:"rpm"`
	StreamTimeout                    param.Field[ModelNewParamsLlmParamsStreamTimeoutUnion]                     `json:"stream_timeout"`
	Timeout                          param.Field[ModelNewParamsLlmParamsTimeoutUnion]                           `json:"timeout"`
	Tpm                              param.Field[int64]                                                         `json:"tpm"`
	UseInPassThrough                 param.Field[bool]                                                          `json:"use_in_pass_through"`
	VertexCredentials                param.Field[interface{}]                                                   `json:"vertex_credentials"`
	VertexLocation                   param.Field[string]                                                        `json:"vertex_location"`
	VertexProject                    param.Field[string]                                                        `json:"vertex_project"`
	WatsonxRegionName                param.Field[string]                                                        `json:"watsonx_region_name"`
	ExtraFields                      map[string]interface{}                                                     `json:"-,extras"`
}

func (r ModelNewParamsLlmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [ConfigurableClientsideParamsCustomAuth].
type ModelNewParamsLlmParamsConfigurableClientsideAuthParamUnion interface {
	ImplementsModelNewParamsLlmParamsConfigurableClientsideAuthParamUnion()
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type ModelNewParamsLlmParamsStreamTimeoutUnion interface {
	ImplementsModelNewParamsLlmParamsStreamTimeoutUnion()
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type ModelNewParamsLlmParamsTimeoutUnion interface {
	ImplementsModelNewParamsLlmParamsTimeoutUnion()
}

type ModelDeleteParams struct {
	ID param.Field[string] `json:"id,required"`
}

func (r ModelDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
