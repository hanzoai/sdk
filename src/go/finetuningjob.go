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

// FineTuningJobService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFineTuningJobService] method instead.
type FineTuningJobService struct {
	Options []option.RequestOption
	Cancel  *FineTuningJobCancelService
}

// NewFineTuningJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFineTuningJobService(opts ...option.RequestOption) (r *FineTuningJobService) {
	r = &FineTuningJobService{}
	r.Options = opts
	r.Cancel = NewFineTuningJobCancelService(opts...)
	return
}

// Creates a fine-tuning job which begins the process of creating a new model from
// a given dataset. This is the equivalent of POST
// https://api.openai.com/v1/fine_tuning/jobs
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/fine-tuning/create
//
// Example Curl:
//
// ```
//
//	curl http://localhost:4000/v1/fine_tuning/jobs       -H "Content-Type: application/json"       -H "Authorization: Bearer sk-1234"       -d '{
//	    "model": "gpt-3.5-turbo",
//	    "training_file": "file-abc123",
//	    "hyperparameters": {
//	      "n_epochs": 4
//	    }
//	  }'
//
// ```
func (r *FineTuningJobService) New(ctx context.Context, body FineTuningJobNewParams, opts ...option.RequestOption) (res *FineTuningJobNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/fine_tuning/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a fine-tuning job. This is the equivalent of GET
// https://api.openai.com/v1/fine_tuning/jobs/{fine_tuning_job_id}
//
// Supported Query Params:
//
// - `custom_llm_provider`: Name of the LLM provider
// - `fine_tuning_job_id`: The ID of the fine-tuning job to retrieve.
func (r *FineTuningJobService) Get(ctx context.Context, fineTuningJobID string, query FineTuningJobGetParams, opts ...option.RequestOption) (res *FineTuningJobGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fineTuningJobID == "" {
		err = errors.New("missing required fine_tuning_job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/fine_tuning/jobs/%s", fineTuningJobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Lists fine-tuning jobs for the organization. This is the equivalent of GET
// https://api.openai.com/v1/fine_tuning/jobs
//
// Supported Query Params:
//
// - `custom_llm_provider`: Name of the LLM provider
// - `after`: Identifier for the last job from the previous pagination request.
// - `limit`: Number of fine-tuning jobs to retrieve (default is 20).
func (r *FineTuningJobService) List(ctx context.Context, query FineTuningJobListParams, opts ...option.RequestOption) (res *FineTuningJobListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/fine_tuning/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type FineTuningJobNewResponse = interface{}

type FineTuningJobGetResponse = interface{}

type FineTuningJobListResponse = interface{}

type FineTuningJobNewParams struct {
	CustomLlmProvider param.Field[FineTuningJobNewParamsCustomLlmProvider] `json:"custom_llm_provider,required"`
	Model             param.Field[string]                                  `json:"model,required"`
	TrainingFile      param.Field[string]                                  `json:"training_file,required"`
	Hyperparameters   param.Field[FineTuningJobNewParamsHyperparameters]   `json:"hyperparameters"`
	Integrations      param.Field[[]string]                                `json:"integrations"`
	Seed              param.Field[int64]                                   `json:"seed"`
	Suffix            param.Field[string]                                  `json:"suffix"`
	ValidationFile    param.Field[string]                                  `json:"validation_file"`
}

func (r FineTuningJobNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FineTuningJobNewParamsCustomLlmProvider string

const (
	FineTuningJobNewParamsCustomLlmProviderOpenAI   FineTuningJobNewParamsCustomLlmProvider = "openai"
	FineTuningJobNewParamsCustomLlmProviderAzure    FineTuningJobNewParamsCustomLlmProvider = "azure"
	FineTuningJobNewParamsCustomLlmProviderVertexAI FineTuningJobNewParamsCustomLlmProvider = "vertex_ai"
)

func (r FineTuningJobNewParamsCustomLlmProvider) IsKnown() bool {
	switch r {
	case FineTuningJobNewParamsCustomLlmProviderOpenAI, FineTuningJobNewParamsCustomLlmProviderAzure, FineTuningJobNewParamsCustomLlmProviderVertexAI:
		return true
	}
	return false
}

type FineTuningJobNewParamsHyperparameters struct {
	BatchSize              param.Field[FineTuningJobNewParamsHyperparametersBatchSizeUnion]              `json:"batch_size"`
	LearningRateMultiplier param.Field[FineTuningJobNewParamsHyperparametersLearningRateMultiplierUnion] `json:"learning_rate_multiplier"`
	NEpochs                param.Field[FineTuningJobNewParamsHyperparametersNEpochsUnion]                `json:"n_epochs"`
}

func (r FineTuningJobNewParamsHyperparameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionInt].
type FineTuningJobNewParamsHyperparametersBatchSizeUnion interface {
	ImplementsFineTuningJobNewParamsHyperparametersBatchSizeUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type FineTuningJobNewParamsHyperparametersLearningRateMultiplierUnion interface {
	ImplementsFineTuningJobNewParamsHyperparametersLearningRateMultiplierUnion()
}

// Satisfied by [shared.UnionString], [shared.UnionInt].
type FineTuningJobNewParamsHyperparametersNEpochsUnion interface {
	ImplementsFineTuningJobNewParamsHyperparametersNEpochsUnion()
}

type FineTuningJobGetParams struct {
	CustomLlmProvider param.Field[FineTuningJobGetParamsCustomLlmProvider] `query:"custom_llm_provider,required"`
}

// URLQuery serializes [FineTuningJobGetParams]'s query parameters as `url.Values`.
func (r FineTuningJobGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FineTuningJobGetParamsCustomLlmProvider string

const (
	FineTuningJobGetParamsCustomLlmProviderOpenAI FineTuningJobGetParamsCustomLlmProvider = "openai"
	FineTuningJobGetParamsCustomLlmProviderAzure  FineTuningJobGetParamsCustomLlmProvider = "azure"
)

func (r FineTuningJobGetParamsCustomLlmProvider) IsKnown() bool {
	switch r {
	case FineTuningJobGetParamsCustomLlmProviderOpenAI, FineTuningJobGetParamsCustomLlmProviderAzure:
		return true
	}
	return false
}

type FineTuningJobListParams struct {
	CustomLlmProvider param.Field[FineTuningJobListParamsCustomLlmProvider] `query:"custom_llm_provider,required"`
	After             param.Field[string]                                   `query:"after"`
	Limit             param.Field[int64]                                    `query:"limit"`
}

// URLQuery serializes [FineTuningJobListParams]'s query parameters as
// `url.Values`.
func (r FineTuningJobListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FineTuningJobListParamsCustomLlmProvider string

const (
	FineTuningJobListParamsCustomLlmProviderOpenAI FineTuningJobListParamsCustomLlmProvider = "openai"
	FineTuningJobListParamsCustomLlmProviderAzure  FineTuningJobListParamsCustomLlmProvider = "azure"
)

func (r FineTuningJobListParamsCustomLlmProvider) IsKnown() bool {
	switch r {
	case FineTuningJobListParamsCustomLlmProviderOpenAI, FineTuningJobListParamsCustomLlmProviderAzure:
		return true
	}
	return false
}
