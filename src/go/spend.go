// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"
	"reflect"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
	"github.com/hanzoai/go-sdk/shared"
	"github.com/tidwall/gjson"
)

// SpendService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpendService] method instead.
type SpendService struct {
	Options []option.RequestOption
}

// NewSpendService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSpendService(opts ...option.RequestOption) (r *SpendService) {
	r = &SpendService{}
	r.Options = opts
	return
}

// Accepts all the params of completion_cost.
//
// Calculate spend **before** making call:
//
// Note: If you see a spend of $0.0 you need to set custom_pricing for your model:
// https://docs.hanzo.ai/docs/proxy/custom_pricing
//
// ```
// curl --location 'http://localhost:4000/spend/calculate'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data '{
//	    "model": "anthropic.claude-v2",
//	    "messages": [{"role": "user", "content": "Hey, how'''s it going?"}]
//	}'
//
// ```
//
// Calculate spend **after** making call:
//
// ```
// curl --location 'http://localhost:4000/spend/calculate'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data '{
//	    "completion_response": {
//	        "id": "chatcmpl-123",
//	        "object": "chat.completion",
//	        "created": 1677652288,
//	        "model": "gpt-3.5-turbo-0125",
//	        "system_fingerprint": "fp_44709d6fcb",
//	        "choices": [{
//	            "index": 0,
//	            "message": {
//	                "role": "assistant",
//	                "content": "Hello there, how may I assist you today?"
//	            },
//	            "logprobs": null,
//	            "finish_reason": "stop"
//	        }]
//	        "usage": {
//	            "prompt_tokens": 9,
//	            "completion_tokens": 12,
//	            "total_tokens": 21
//	        }
//	    }
//	}'
//
// ```
func (r *SpendService) CalculateSpend(ctx context.Context, body SpendCalculateSpendParams, opts ...option.RequestOption) (res *SpendCalculateSpendResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "spend/calculate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// View all spend logs, if request_id is provided, only logs for that request_id
// will be returned
//
// # Example Request for all logs
//
// ```
// curl -X GET "http://0.0.0.0:8000/spend/logs" -H "Authorization: Bearer sk-1234"
// ```
//
// Example Request for specific request_id
//
// ```
// curl -X GET "http://0.0.0.0:8000/spend/logs?request_id=chatcmpl-6dcb2540-d3d7-4e49-bb27-291f863f112e" -H "Authorization: Bearer sk-1234"
// ```
//
// Example Request for specific api_key
//
// ```
// curl -X GET "http://0.0.0.0:8000/spend/logs?api_key=sk-Fn8Ej39NkBQmUagFEoUWPQ" -H "Authorization: Bearer sk-1234"
// ```
//
// Example Request for specific user_id
//
// ```
// curl -X GET "http://0.0.0.0:8000/spend/logs?user_id=z@hanzo.ai" -H "Authorization: Bearer sk-1234"
// ```
func (r *SpendService) ListLogs(ctx context.Context, query SpendListLogsParams, opts ...option.RequestOption) (res *[]SpendListLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "spend/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// LLM Enterprise - View Spend Per Request Tag
//
// Example Request:
//
// ```
// curl -X GET "http://0.0.0.0:8000/spend/tags" -H "Authorization: Bearer sk-1234"
// ```
//
// # Spend with Start Date and End Date
//
// ```
// curl -X GET "http://0.0.0.0:8000/spend/tags?start_date=2022-01-01&end_date=2022-02-01" -H "Authorization: Bearer sk-1234"
// ```
func (r *SpendService) ListTags(ctx context.Context, query SpendListTagsParams, opts ...option.RequestOption) (res *[]SpendListTagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "spend/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SpendCalculateSpendResponse = interface{}

type SpendListLogsResponse struct {
	APIKey             string                              `json:"api_key,required"`
	CallType           string                              `json:"call_type,required"`
	EndTime            SpendListLogsResponseEndTimeUnion   `json:"endTime,required,nullable" format:"date-time"`
	Messages           SpendListLogsResponseMessagesUnion  `json:"messages,required,nullable"`
	RequestID          string                              `json:"request_id,required"`
	Response           SpendListLogsResponseResponseUnion  `json:"response,required,nullable"`
	StartTime          SpendListLogsResponseStartTimeUnion `json:"startTime,required,nullable" format:"date-time"`
	APIBase            string                              `json:"api_base,nullable"`
	CacheHit           string                              `json:"cache_hit,nullable"`
	CacheKey           string                              `json:"cache_key,nullable"`
	CompletionTokens   int64                               `json:"completion_tokens,nullable"`
	Metadata           interface{}                         `json:"metadata"`
	Model              string                              `json:"model,nullable"`
	PromptTokens       int64                               `json:"prompt_tokens,nullable"`
	RequestTags        interface{}                         `json:"request_tags"`
	RequesterIPAddress string                              `json:"requester_ip_address,nullable"`
	Spend              float64                             `json:"spend,nullable"`
	TotalTokens        int64                               `json:"total_tokens,nullable"`
	User               string                              `json:"user,nullable"`
	JSON               spendListLogsResponseJSON           `json:"-"`
}

// spendListLogsResponseJSON contains the JSON metadata for the struct
// [SpendListLogsResponse]
type spendListLogsResponseJSON struct {
	APIKey             apijson.Field
	CallType           apijson.Field
	EndTime            apijson.Field
	Messages           apijson.Field
	RequestID          apijson.Field
	Response           apijson.Field
	StartTime          apijson.Field
	APIBase            apijson.Field
	CacheHit           apijson.Field
	CacheKey           apijson.Field
	CompletionTokens   apijson.Field
	Metadata           apijson.Field
	Model              apijson.Field
	PromptTokens       apijson.Field
	RequestTags        apijson.Field
	RequesterIPAddress apijson.Field
	Spend              apijson.Field
	TotalTokens        apijson.Field
	User               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendListLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendListLogsResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type SpendListLogsResponseEndTimeUnion interface {
	ImplementsSpendListLogsResponseEndTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListLogsResponseEndTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

// Union satisfied by [shared.UnionString] or [SpendListLogsResponseMessagesArray].
type SpendListLogsResponseMessagesUnion interface {
	ImplementsSpendListLogsResponseMessagesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListLogsResponseMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SpendListLogsResponseMessagesArray{}),
		},
	)
}

type SpendListLogsResponseMessagesArray []interface{}

func (r SpendListLogsResponseMessagesArray) ImplementsSpendListLogsResponseMessagesUnion() {}

// Union satisfied by [shared.UnionString] or [SpendListLogsResponseResponseArray].
type SpendListLogsResponseResponseUnion interface {
	ImplementsSpendListLogsResponseResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListLogsResponseResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SpendListLogsResponseResponseArray{}),
		},
	)
}

type SpendListLogsResponseResponseArray []interface{}

func (r SpendListLogsResponseResponseArray) ImplementsSpendListLogsResponseResponseUnion() {}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type SpendListLogsResponseStartTimeUnion interface {
	ImplementsSpendListLogsResponseStartTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListLogsResponseStartTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

type SpendListTagsResponse struct {
	APIKey             string                              `json:"api_key,required"`
	CallType           string                              `json:"call_type,required"`
	EndTime            SpendListTagsResponseEndTimeUnion   `json:"endTime,required,nullable" format:"date-time"`
	Messages           SpendListTagsResponseMessagesUnion  `json:"messages,required,nullable"`
	RequestID          string                              `json:"request_id,required"`
	Response           SpendListTagsResponseResponseUnion  `json:"response,required,nullable"`
	StartTime          SpendListTagsResponseStartTimeUnion `json:"startTime,required,nullable" format:"date-time"`
	APIBase            string                              `json:"api_base,nullable"`
	CacheHit           string                              `json:"cache_hit,nullable"`
	CacheKey           string                              `json:"cache_key,nullable"`
	CompletionTokens   int64                               `json:"completion_tokens,nullable"`
	Metadata           interface{}                         `json:"metadata"`
	Model              string                              `json:"model,nullable"`
	PromptTokens       int64                               `json:"prompt_tokens,nullable"`
	RequestTags        interface{}                         `json:"request_tags"`
	RequesterIPAddress string                              `json:"requester_ip_address,nullable"`
	Spend              float64                             `json:"spend,nullable"`
	TotalTokens        int64                               `json:"total_tokens,nullable"`
	User               string                              `json:"user,nullable"`
	JSON               spendListTagsResponseJSON           `json:"-"`
}

// spendListTagsResponseJSON contains the JSON metadata for the struct
// [SpendListTagsResponse]
type spendListTagsResponseJSON struct {
	APIKey             apijson.Field
	CallType           apijson.Field
	EndTime            apijson.Field
	Messages           apijson.Field
	RequestID          apijson.Field
	Response           apijson.Field
	StartTime          apijson.Field
	APIBase            apijson.Field
	CacheHit           apijson.Field
	CacheKey           apijson.Field
	CompletionTokens   apijson.Field
	Metadata           apijson.Field
	Model              apijson.Field
	PromptTokens       apijson.Field
	RequestTags        apijson.Field
	RequesterIPAddress apijson.Field
	Spend              apijson.Field
	TotalTokens        apijson.Field
	User               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SpendListTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r spendListTagsResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type SpendListTagsResponseEndTimeUnion interface {
	ImplementsSpendListTagsResponseEndTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListTagsResponseEndTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

// Union satisfied by [shared.UnionString] or [SpendListTagsResponseMessagesArray].
type SpendListTagsResponseMessagesUnion interface {
	ImplementsSpendListTagsResponseMessagesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListTagsResponseMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SpendListTagsResponseMessagesArray{}),
		},
	)
}

type SpendListTagsResponseMessagesArray []interface{}

func (r SpendListTagsResponseMessagesArray) ImplementsSpendListTagsResponseMessagesUnion() {}

// Union satisfied by [shared.UnionString] or [SpendListTagsResponseResponseArray].
type SpendListTagsResponseResponseUnion interface {
	ImplementsSpendListTagsResponseResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListTagsResponseResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(SpendListTagsResponseResponseArray{}),
		},
	)
}

type SpendListTagsResponseResponseArray []interface{}

func (r SpendListTagsResponseResponseArray) ImplementsSpendListTagsResponseResponseUnion() {}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type SpendListTagsResponseStartTimeUnion interface {
	ImplementsSpendListTagsResponseStartTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpendListTagsResponseStartTimeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

type SpendCalculateSpendParams struct {
	CompletionResponse param.Field[interface{}]   `json:"completion_response"`
	Messages           param.Field[[]interface{}] `json:"messages"`
	Model              param.Field[string]        `json:"model"`
}

func (r SpendCalculateSpendParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SpendListLogsParams struct {
	// Get spend logs based on api key
	APIKey param.Field[string] `query:"api_key"`
	// Time till which to view key spend
	EndDate param.Field[string] `query:"end_date"`
	// request_id to get spend logs for specific request_id. If none passed then pass
	// spend logs for all requests
	RequestID param.Field[string] `query:"request_id"`
	// Time from which to start viewing key spend
	StartDate param.Field[string] `query:"start_date"`
	// Get spend logs based on user_id
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [SpendListLogsParams]'s query parameters as `url.Values`.
func (r SpendListLogsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpendListTagsParams struct {
	// Time till which to view key spend
	EndDate param.Field[string] `query:"end_date"`
	// Time from which to start viewing key spend
	StartDate param.Field[string] `query:"start_date"`
}

// URLQuery serializes [SpendListTagsParams]'s query parameters as `url.Values`.
func (r SpendListTagsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
