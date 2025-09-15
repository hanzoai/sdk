// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// UtilService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUtilService] method instead.
type UtilService struct {
	Options []option.RequestOption
}

// NewUtilService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUtilService(opts ...option.RequestOption) (r *UtilService) {
	r = &UtilService{}
	r.Options = opts
	return
}

// Returns supported openai params for a given llm model name
//
// e.g. `gpt-4` vs `gpt-3.5-turbo`
//
// Example curl:
//
// ```
// curl -X GET --location 'http://localhost:4000/utils/supported_openai_params?model=gpt-3.5-turbo-16k'         --header 'Authorization: Bearer sk-1234'
// ```
func (r *UtilService) GetSupportedOpenAIParams(ctx context.Context, query UtilGetSupportedOpenAIParamsParams, opts ...option.RequestOption) (res *UtilGetSupportedOpenAIParamsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "utils/supported_openai_params"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Token Counter
func (r *UtilService) TokenCounter(ctx context.Context, body UtilTokenCounterParams, opts ...option.RequestOption) (res *UtilTokenCounterResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "utils/token_counter"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Transform Request
func (r *UtilService) TransformRequest(ctx context.Context, body UtilTransformRequestParams, opts ...option.RequestOption) (res *UtilTransformRequestResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "utils/transform_request"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UtilGetSupportedOpenAIParamsResponse = interface{}

type UtilTokenCounterResponse struct {
	ModelUsed     string                       `json:"model_used,required"`
	RequestModel  string                       `json:"request_model,required"`
	TokenizerType string                       `json:"tokenizer_type,required"`
	TotalTokens   int64                        `json:"total_tokens,required"`
	JSON          utilTokenCounterResponseJSON `json:"-"`
}

// utilTokenCounterResponseJSON contains the JSON metadata for the struct
// [UtilTokenCounterResponse]
type utilTokenCounterResponseJSON struct {
	ModelUsed     apijson.Field
	RequestModel  apijson.Field
	TokenizerType apijson.Field
	TotalTokens   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UtilTokenCounterResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r utilTokenCounterResponseJSON) RawJSON() string {
	return r.raw
}

type UtilTransformRequestResponse struct {
	Error             string                           `json:"error,nullable"`
	RawRequestAPIBase string                           `json:"raw_request_api_base,nullable"`
	RawRequestBody    interface{}                      `json:"raw_request_body,nullable"`
	RawRequestHeaders interface{}                      `json:"raw_request_headers,nullable"`
	JSON              utilTransformRequestResponseJSON `json:"-"`
}

// utilTransformRequestResponseJSON contains the JSON metadata for the struct
// [UtilTransformRequestResponse]
type utilTransformRequestResponseJSON struct {
	Error             apijson.Field
	RawRequestAPIBase apijson.Field
	RawRequestBody    apijson.Field
	RawRequestHeaders apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UtilTransformRequestResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r utilTransformRequestResponseJSON) RawJSON() string {
	return r.raw
}

type UtilGetSupportedOpenAIParamsParams struct {
	Model param.Field[string] `query:"model,required"`
}

// URLQuery serializes [UtilGetSupportedOpenAIParamsParams]'s query parameters as
// `url.Values`.
func (r UtilGetSupportedOpenAIParamsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UtilTokenCounterParams struct {
	Model    param.Field[string]        `json:"model,required"`
	Messages param.Field[[]interface{}] `json:"messages"`
	Prompt   param.Field[string]        `json:"prompt"`
}

func (r UtilTokenCounterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UtilTransformRequestParams struct {
	CallType    param.Field[UtilTransformRequestParamsCallType] `json:"call_type,required"`
	RequestBody param.Field[interface{}]                        `json:"request_body,required"`
}

func (r UtilTransformRequestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UtilTransformRequestParamsCallType string

const (
	UtilTransformRequestParamsCallTypeEmbedding              UtilTransformRequestParamsCallType = "embedding"
	UtilTransformRequestParamsCallTypeAembedding             UtilTransformRequestParamsCallType = "aembedding"
	UtilTransformRequestParamsCallTypeCompletion             UtilTransformRequestParamsCallType = "completion"
	UtilTransformRequestParamsCallTypeAcompletion            UtilTransformRequestParamsCallType = "acompletion"
	UtilTransformRequestParamsCallTypeAtextCompletion        UtilTransformRequestParamsCallType = "atext_completion"
	UtilTransformRequestParamsCallTypeTextCompletion         UtilTransformRequestParamsCallType = "text_completion"
	UtilTransformRequestParamsCallTypeImageGeneration        UtilTransformRequestParamsCallType = "image_generation"
	UtilTransformRequestParamsCallTypeAimageGeneration       UtilTransformRequestParamsCallType = "aimage_generation"
	UtilTransformRequestParamsCallTypeModeration             UtilTransformRequestParamsCallType = "moderation"
	UtilTransformRequestParamsCallTypeAmoderation            UtilTransformRequestParamsCallType = "amoderation"
	UtilTransformRequestParamsCallTypeAtranscription         UtilTransformRequestParamsCallType = "atranscription"
	UtilTransformRequestParamsCallTypeTranscription          UtilTransformRequestParamsCallType = "transcription"
	UtilTransformRequestParamsCallTypeAspeech                UtilTransformRequestParamsCallType = "aspeech"
	UtilTransformRequestParamsCallTypeSpeech                 UtilTransformRequestParamsCallType = "speech"
	UtilTransformRequestParamsCallTypeRerank                 UtilTransformRequestParamsCallType = "rerank"
	UtilTransformRequestParamsCallTypeArerank                UtilTransformRequestParamsCallType = "arerank"
	UtilTransformRequestParamsCallType_Arealtime             UtilTransformRequestParamsCallType = "_arealtime"
	UtilTransformRequestParamsCallTypeCreateBatch            UtilTransformRequestParamsCallType = "create_batch"
	UtilTransformRequestParamsCallTypeAcreateBatch           UtilTransformRequestParamsCallType = "acreate_batch"
	UtilTransformRequestParamsCallTypeAretrieveBatch         UtilTransformRequestParamsCallType = "aretrieve_batch"
	UtilTransformRequestParamsCallTypeRetrieveBatch          UtilTransformRequestParamsCallType = "retrieve_batch"
	UtilTransformRequestParamsCallTypePassThroughEndpoint    UtilTransformRequestParamsCallType = "pass_through_endpoint"
	UtilTransformRequestParamsCallTypeAnthropicMessages      UtilTransformRequestParamsCallType = "anthropic_messages"
	UtilTransformRequestParamsCallTypeGetAssistants          UtilTransformRequestParamsCallType = "get_assistants"
	UtilTransformRequestParamsCallTypeAgetAssistants         UtilTransformRequestParamsCallType = "aget_assistants"
	UtilTransformRequestParamsCallTypeCreateAssistants       UtilTransformRequestParamsCallType = "create_assistants"
	UtilTransformRequestParamsCallTypeAcreateAssistants      UtilTransformRequestParamsCallType = "acreate_assistants"
	UtilTransformRequestParamsCallTypeDeleteAssistant        UtilTransformRequestParamsCallType = "delete_assistant"
	UtilTransformRequestParamsCallTypeAdeleteAssistant       UtilTransformRequestParamsCallType = "adelete_assistant"
	UtilTransformRequestParamsCallTypeAcreateThread          UtilTransformRequestParamsCallType = "acreate_thread"
	UtilTransformRequestParamsCallTypeCreateThread           UtilTransformRequestParamsCallType = "create_thread"
	UtilTransformRequestParamsCallTypeAgetThread             UtilTransformRequestParamsCallType = "aget_thread"
	UtilTransformRequestParamsCallTypeGetThread              UtilTransformRequestParamsCallType = "get_thread"
	UtilTransformRequestParamsCallTypeAAddMessage            UtilTransformRequestParamsCallType = "a_add_message"
	UtilTransformRequestParamsCallTypeAddMessage             UtilTransformRequestParamsCallType = "add_message"
	UtilTransformRequestParamsCallTypeAgetMessages           UtilTransformRequestParamsCallType = "aget_messages"
	UtilTransformRequestParamsCallTypeGetMessages            UtilTransformRequestParamsCallType = "get_messages"
	UtilTransformRequestParamsCallTypeArunThread             UtilTransformRequestParamsCallType = "arun_thread"
	UtilTransformRequestParamsCallTypeRunThread              UtilTransformRequestParamsCallType = "run_thread"
	UtilTransformRequestParamsCallTypeArunThreadStream       UtilTransformRequestParamsCallType = "arun_thread_stream"
	UtilTransformRequestParamsCallTypeRunThreadStream        UtilTransformRequestParamsCallType = "run_thread_stream"
	UtilTransformRequestParamsCallTypeAfileRetrieve          UtilTransformRequestParamsCallType = "afile_retrieve"
	UtilTransformRequestParamsCallTypeFileRetrieve           UtilTransformRequestParamsCallType = "file_retrieve"
	UtilTransformRequestParamsCallTypeAfileDelete            UtilTransformRequestParamsCallType = "afile_delete"
	UtilTransformRequestParamsCallTypeFileDelete             UtilTransformRequestParamsCallType = "file_delete"
	UtilTransformRequestParamsCallTypeAfileList              UtilTransformRequestParamsCallType = "afile_list"
	UtilTransformRequestParamsCallTypeFileList               UtilTransformRequestParamsCallType = "file_list"
	UtilTransformRequestParamsCallTypeAcreateFile            UtilTransformRequestParamsCallType = "acreate_file"
	UtilTransformRequestParamsCallTypeCreateFile             UtilTransformRequestParamsCallType = "create_file"
	UtilTransformRequestParamsCallTypeAfileContent           UtilTransformRequestParamsCallType = "afile_content"
	UtilTransformRequestParamsCallTypeFileContent            UtilTransformRequestParamsCallType = "file_content"
	UtilTransformRequestParamsCallTypeCreateFineTuningJob    UtilTransformRequestParamsCallType = "create_fine_tuning_job"
	UtilTransformRequestParamsCallTypeAcreateFineTuningJob   UtilTransformRequestParamsCallType = "acreate_fine_tuning_job"
	UtilTransformRequestParamsCallTypeAcancelFineTuningJob   UtilTransformRequestParamsCallType = "acancel_fine_tuning_job"
	UtilTransformRequestParamsCallTypeCancelFineTuningJob    UtilTransformRequestParamsCallType = "cancel_fine_tuning_job"
	UtilTransformRequestParamsCallTypeAlistFineTuningJobs    UtilTransformRequestParamsCallType = "alist_fine_tuning_jobs"
	UtilTransformRequestParamsCallTypeListFineTuningJobs     UtilTransformRequestParamsCallType = "list_fine_tuning_jobs"
	UtilTransformRequestParamsCallTypeAretrieveFineTuningJob UtilTransformRequestParamsCallType = "aretrieve_fine_tuning_job"
	UtilTransformRequestParamsCallTypeRetrieveFineTuningJob  UtilTransformRequestParamsCallType = "retrieve_fine_tuning_job"
	UtilTransformRequestParamsCallTypeResponses              UtilTransformRequestParamsCallType = "responses"
	UtilTransformRequestParamsCallTypeAresponses             UtilTransformRequestParamsCallType = "aresponses"
)

func (r UtilTransformRequestParamsCallType) IsKnown() bool {
	switch r {
	case UtilTransformRequestParamsCallTypeEmbedding, UtilTransformRequestParamsCallTypeAembedding, UtilTransformRequestParamsCallTypeCompletion, UtilTransformRequestParamsCallTypeAcompletion, UtilTransformRequestParamsCallTypeAtextCompletion, UtilTransformRequestParamsCallTypeTextCompletion, UtilTransformRequestParamsCallTypeImageGeneration, UtilTransformRequestParamsCallTypeAimageGeneration, UtilTransformRequestParamsCallTypeModeration, UtilTransformRequestParamsCallTypeAmoderation, UtilTransformRequestParamsCallTypeAtranscription, UtilTransformRequestParamsCallTypeTranscription, UtilTransformRequestParamsCallTypeAspeech, UtilTransformRequestParamsCallTypeSpeech, UtilTransformRequestParamsCallTypeRerank, UtilTransformRequestParamsCallTypeArerank, UtilTransformRequestParamsCallType_Arealtime, UtilTransformRequestParamsCallTypeCreateBatch, UtilTransformRequestParamsCallTypeAcreateBatch, UtilTransformRequestParamsCallTypeAretrieveBatch, UtilTransformRequestParamsCallTypeRetrieveBatch, UtilTransformRequestParamsCallTypePassThroughEndpoint, UtilTransformRequestParamsCallTypeAnthropicMessages, UtilTransformRequestParamsCallTypeGetAssistants, UtilTransformRequestParamsCallTypeAgetAssistants, UtilTransformRequestParamsCallTypeCreateAssistants, UtilTransformRequestParamsCallTypeAcreateAssistants, UtilTransformRequestParamsCallTypeDeleteAssistant, UtilTransformRequestParamsCallTypeAdeleteAssistant, UtilTransformRequestParamsCallTypeAcreateThread, UtilTransformRequestParamsCallTypeCreateThread, UtilTransformRequestParamsCallTypeAgetThread, UtilTransformRequestParamsCallTypeGetThread, UtilTransformRequestParamsCallTypeAAddMessage, UtilTransformRequestParamsCallTypeAddMessage, UtilTransformRequestParamsCallTypeAgetMessages, UtilTransformRequestParamsCallTypeGetMessages, UtilTransformRequestParamsCallTypeArunThread, UtilTransformRequestParamsCallTypeRunThread, UtilTransformRequestParamsCallTypeArunThreadStream, UtilTransformRequestParamsCallTypeRunThreadStream, UtilTransformRequestParamsCallTypeAfileRetrieve, UtilTransformRequestParamsCallTypeFileRetrieve, UtilTransformRequestParamsCallTypeAfileDelete, UtilTransformRequestParamsCallTypeFileDelete, UtilTransformRequestParamsCallTypeAfileList, UtilTransformRequestParamsCallTypeFileList, UtilTransformRequestParamsCallTypeAcreateFile, UtilTransformRequestParamsCallTypeCreateFile, UtilTransformRequestParamsCallTypeAfileContent, UtilTransformRequestParamsCallTypeFileContent, UtilTransformRequestParamsCallTypeCreateFineTuningJob, UtilTransformRequestParamsCallTypeAcreateFineTuningJob, UtilTransformRequestParamsCallTypeAcancelFineTuningJob, UtilTransformRequestParamsCallTypeCancelFineTuningJob, UtilTransformRequestParamsCallTypeAlistFineTuningJobs, UtilTransformRequestParamsCallTypeListFineTuningJobs, UtilTransformRequestParamsCallTypeAretrieveFineTuningJob, UtilTransformRequestParamsCallTypeRetrieveFineTuningJob, UtilTransformRequestParamsCallTypeResponses, UtilTransformRequestParamsCallTypeAresponses:
		return true
	}
	return false
}
