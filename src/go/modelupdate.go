// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ModelUpdateService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelUpdateService] method instead.
type ModelUpdateService struct {
	Options []option.RequestOption
}

// NewModelUpdateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewModelUpdateService(opts ...option.RequestOption) (r *ModelUpdateService) {
	r = &ModelUpdateService{}
	r.Options = opts
	return
}

// Edit existing model params
func (r *ModelUpdateService) Full(ctx context.Context, body ModelUpdateFullParams, opts ...option.RequestOption) (res *ModelUpdateFullResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "model/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// PATCH Endpoint for partial model updates.
//
// Only updates the fields specified in the request while preserving other existing
// values. Follows proper PATCH semantics by only modifying provided fields.
//
// Args: model_id: The ID of the model to update patch_data: The fields to update
// and their new values user_api_key_dict: User authentication information
//
// Returns: Updated model information
//
// Raises: ProxyException: For various error conditions including authentication
// and database errors
func (r *ModelUpdateService) Partial(ctx context.Context, modelID string, body ModelUpdatePartialParams, opts ...option.RequestOption) (res *ModelUpdatePartialResponse, err error) {
	opts = append(r.Options[:], opts...)
	if modelID == "" {
		err = errors.New("missing required model_id parameter")
		return
	}
	path := fmt.Sprintf("model/%s/update", modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type UpdateDeploymentParam struct {
	LlmParams param.Field[UpdateDeploymentLlmParamsParam] `json:"llm_params"`
	ModelInfo param.Field[ModelInfoParam]                 `json:"model_info"`
	ModelName param.Field[string]                         `json:"model_name"`
}

func (r UpdateDeploymentParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UpdateDeploymentLlmParamsParam struct {
	APIBase                          param.Field[string]                                                                `json:"api_base"`
	APIKey                           param.Field[string]                                                                `json:"api_key"`
	APIVersion                       param.Field[string]                                                                `json:"api_version"`
	AwsAccessKeyID                   param.Field[string]                                                                `json:"aws_access_key_id"`
	AwsRegionName                    param.Field[string]                                                                `json:"aws_region_name"`
	AwsSecretAccessKey               param.Field[string]                                                                `json:"aws_secret_access_key"`
	BudgetDuration                   param.Field[string]                                                                `json:"budget_duration"`
	ConfigurableClientsideAuthParams param.Field[[]UpdateDeploymentLlmParamsConfigurableClientsideAuthParamsUnionParam] `json:"configurable_clientside_auth_params"`
	CustomLlmProvider                param.Field[string]                                                                `json:"custom_llm_provider"`
	InputCostPerSecond               param.Field[float64]                                                               `json:"input_cost_per_second"`
	InputCostPerToken                param.Field[float64]                                                               `json:"input_cost_per_token"`
	LlmTraceID                       param.Field[string]                                                                `json:"llm_trace_id"`
	MaxBudget                        param.Field[float64]                                                               `json:"max_budget"`
	MaxFileSizeMB                    param.Field[float64]                                                               `json:"max_file_size_mb"`
	MaxRetries                       param.Field[int64]                                                                 `json:"max_retries"`
	MergeReasoningContentInChoices   param.Field[bool]                                                                  `json:"merge_reasoning_content_in_choices"`
	Model                            param.Field[string]                                                                `json:"model"`
	ModelInfo                        param.Field[interface{}]                                                           `json:"model_info"`
	Organization                     param.Field[string]                                                                `json:"organization"`
	OutputCostPerSecond              param.Field[float64]                                                               `json:"output_cost_per_second"`
	OutputCostPerToken               param.Field[float64]                                                               `json:"output_cost_per_token"`
	RegionName                       param.Field[string]                                                                `json:"region_name"`
	Rpm                              param.Field[int64]                                                                 `json:"rpm"`
	StreamTimeout                    param.Field[UpdateDeploymentLlmParamsStreamTimeoutUnionParam]                      `json:"stream_timeout"`
	Timeout                          param.Field[UpdateDeploymentLlmParamsTimeoutUnionParam]                            `json:"timeout"`
	Tpm                              param.Field[int64]                                                                 `json:"tpm"`
	UseInPassThrough                 param.Field[bool]                                                                  `json:"use_in_pass_through"`
	VertexCredentials                param.Field[interface{}]                                                           `json:"vertex_credentials"`
	VertexLocation                   param.Field[string]                                                                `json:"vertex_location"`
	VertexProject                    param.Field[string]                                                                `json:"vertex_project"`
	WatsonxRegionName                param.Field[string]                                                                `json:"watsonx_region_name"`
	ExtraFields                      map[string]interface{}                                                             `json:"-,extras"`
}

func (r UpdateDeploymentLlmParamsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [ConfigurableClientsideParamsCustomAuth].
type UpdateDeploymentLlmParamsConfigurableClientsideAuthParamsUnionParam interface {
	ImplementsUpdateDeploymentLlmParamsConfigurableClientsideAuthParamsUnionParam()
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type UpdateDeploymentLlmParamsStreamTimeoutUnionParam interface {
	ImplementsUpdateDeploymentLlmParamsStreamTimeoutUnionParam()
}

// Satisfied by [shared.UnionFloat], [shared.UnionString].
type UpdateDeploymentLlmParamsTimeoutUnionParam interface {
	ImplementsUpdateDeploymentLlmParamsTimeoutUnionParam()
}

type ModelUpdateFullResponse = interface{}

type ModelUpdatePartialResponse = interface{}

type ModelUpdateFullParams struct {
	UpdateDeployment UpdateDeploymentParam `json:"update_deployment,required"`
}

func (r ModelUpdateFullParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateDeployment)
}

type ModelUpdatePartialParams struct {
	UpdateDeployment UpdateDeploymentParam `json:"update_deployment,required"`
}

func (r ModelUpdatePartialParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.UpdateDeployment)
}
