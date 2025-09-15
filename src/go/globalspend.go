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

// GlobalSpendService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGlobalSpendService] method instead.
type GlobalSpendService struct {
	Options []option.RequestOption
}

// NewGlobalSpendService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGlobalSpendService(opts ...option.RequestOption) (r *GlobalSpendService) {
	r = &GlobalSpendService{}
	r.Options = opts
	return
}

// LLM Enterprise - View Spend Per Request Tag. Used by LLM UI
//
// Example Request:
//
// ```
// curl -X GET "http://0.0.0.0:4000/spend/tags" -H "Authorization: Bearer sk-1234"
// ```
//
// # Spend with Start Date and End Date
//
// ```
// curl -X GET "http://0.0.0.0:4000/spend/tags?start_date=2022-01-01&end_date=2022-02-01" -H "Authorization: Bearer sk-1234"
// ```
func (r *GlobalSpendService) ListTags(ctx context.Context, query GlobalSpendListTagsParams, opts ...option.RequestOption) (res *[]GlobalSpendListTagsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global/spend/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// ADMIN ONLY / MASTER KEY Only Endpoint
//
// Globally reset spend for All API Keys and Teams, maintain LLM_SpendLogs
//
//  1. LLM_SpendLogs will maintain the logs on spend, no data gets deleted from
//     there
//  2. LLM_VerificationTokens spend will be set = 0
//  3. LLM_TeamTable spend will be set = 0
func (r *GlobalSpendService) Reset(ctx context.Context, opts ...option.RequestOption) (res *GlobalSpendResetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global/spend/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get Daily Spend per Team, based on specific startTime and endTime. Per team,
// view usage by each key, model [ { "group-by-day": "2024-05-10", "teams": [ {
// "team_name": "team-1" "spend": 10, "keys": [ "key": "1213", "usage": {
// "model-1": { "cost": 12.50, "input_tokens": 1000, "output_tokens": 5000,
// "requests": 100 }, "audio-modelname1": { "cost": 25.50, "seconds": 25,
// "requests": 50 }, } } ] ] }
func (r *GlobalSpendService) GetReport(ctx context.Context, query GlobalSpendGetReportParams, opts ...option.RequestOption) (res *[]GlobalSpendGetReportResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "global/spend/report"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type GlobalSpendListTagsResponse struct {
	APIKey             string                                    `json:"api_key,required"`
	CallType           string                                    `json:"call_type,required"`
	EndTime            GlobalSpendListTagsResponseEndTimeUnion   `json:"endTime,required,nullable" format:"date-time"`
	Messages           GlobalSpendListTagsResponseMessagesUnion  `json:"messages,required,nullable"`
	RequestID          string                                    `json:"request_id,required"`
	Response           GlobalSpendListTagsResponseResponseUnion  `json:"response,required,nullable"`
	StartTime          GlobalSpendListTagsResponseStartTimeUnion `json:"startTime,required,nullable" format:"date-time"`
	APIBase            string                                    `json:"api_base,nullable"`
	CacheHit           string                                    `json:"cache_hit,nullable"`
	CacheKey           string                                    `json:"cache_key,nullable"`
	CompletionTokens   int64                                     `json:"completion_tokens,nullable"`
	Metadata           interface{}                               `json:"metadata"`
	Model              string                                    `json:"model,nullable"`
	PromptTokens       int64                                     `json:"prompt_tokens,nullable"`
	RequestTags        interface{}                               `json:"request_tags"`
	RequesterIPAddress string                                    `json:"requester_ip_address,nullable"`
	Spend              float64                                   `json:"spend,nullable"`
	TotalTokens        int64                                     `json:"total_tokens,nullable"`
	User               string                                    `json:"user,nullable"`
	JSON               globalSpendListTagsResponseJSON           `json:"-"`
}

// globalSpendListTagsResponseJSON contains the JSON metadata for the struct
// [GlobalSpendListTagsResponse]
type globalSpendListTagsResponseJSON struct {
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

func (r *GlobalSpendListTagsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSpendListTagsResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type GlobalSpendListTagsResponseEndTimeUnion interface {
	ImplementsGlobalSpendListTagsResponseEndTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendListTagsResponseEndTimeUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or
// [GlobalSpendListTagsResponseMessagesArray].
type GlobalSpendListTagsResponseMessagesUnion interface {
	ImplementsGlobalSpendListTagsResponseMessagesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendListTagsResponseMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GlobalSpendListTagsResponseMessagesArray{}),
		},
	)
}

type GlobalSpendListTagsResponseMessagesArray []interface{}

func (r GlobalSpendListTagsResponseMessagesArray) ImplementsGlobalSpendListTagsResponseMessagesUnion() {
}

// Union satisfied by [shared.UnionString] or
// [GlobalSpendListTagsResponseResponseArray].
type GlobalSpendListTagsResponseResponseUnion interface {
	ImplementsGlobalSpendListTagsResponseResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendListTagsResponseResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GlobalSpendListTagsResponseResponseArray{}),
		},
	)
}

type GlobalSpendListTagsResponseResponseArray []interface{}

func (r GlobalSpendListTagsResponseResponseArray) ImplementsGlobalSpendListTagsResponseResponseUnion() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type GlobalSpendListTagsResponseStartTimeUnion interface {
	ImplementsGlobalSpendListTagsResponseStartTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendListTagsResponseStartTimeUnion)(nil)).Elem(),
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

type GlobalSpendResetResponse = interface{}

type GlobalSpendGetReportResponse struct {
	APIKey             string                                     `json:"api_key,required"`
	CallType           string                                     `json:"call_type,required"`
	EndTime            GlobalSpendGetReportResponseEndTimeUnion   `json:"endTime,required,nullable" format:"date-time"`
	Messages           GlobalSpendGetReportResponseMessagesUnion  `json:"messages,required,nullable"`
	RequestID          string                                     `json:"request_id,required"`
	Response           GlobalSpendGetReportResponseResponseUnion  `json:"response,required,nullable"`
	StartTime          GlobalSpendGetReportResponseStartTimeUnion `json:"startTime,required,nullable" format:"date-time"`
	APIBase            string                                     `json:"api_base,nullable"`
	CacheHit           string                                     `json:"cache_hit,nullable"`
	CacheKey           string                                     `json:"cache_key,nullable"`
	CompletionTokens   int64                                      `json:"completion_tokens,nullable"`
	Metadata           interface{}                                `json:"metadata"`
	Model              string                                     `json:"model,nullable"`
	PromptTokens       int64                                      `json:"prompt_tokens,nullable"`
	RequestTags        interface{}                                `json:"request_tags"`
	RequesterIPAddress string                                     `json:"requester_ip_address,nullable"`
	Spend              float64                                    `json:"spend,nullable"`
	TotalTokens        int64                                      `json:"total_tokens,nullable"`
	User               string                                     `json:"user,nullable"`
	JSON               globalSpendGetReportResponseJSON           `json:"-"`
}

// globalSpendGetReportResponseJSON contains the JSON metadata for the struct
// [GlobalSpendGetReportResponse]
type globalSpendGetReportResponseJSON struct {
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

func (r *GlobalSpendGetReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r globalSpendGetReportResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type GlobalSpendGetReportResponseEndTimeUnion interface {
	ImplementsGlobalSpendGetReportResponseEndTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendGetReportResponseEndTimeUnion)(nil)).Elem(),
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

// Union satisfied by [shared.UnionString] or
// [GlobalSpendGetReportResponseMessagesArray].
type GlobalSpendGetReportResponseMessagesUnion interface {
	ImplementsGlobalSpendGetReportResponseMessagesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendGetReportResponseMessagesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GlobalSpendGetReportResponseMessagesArray{}),
		},
	)
}

type GlobalSpendGetReportResponseMessagesArray []interface{}

func (r GlobalSpendGetReportResponseMessagesArray) ImplementsGlobalSpendGetReportResponseMessagesUnion() {
}

// Union satisfied by [shared.UnionString] or
// [GlobalSpendGetReportResponseResponseArray].
type GlobalSpendGetReportResponseResponseUnion interface {
	ImplementsGlobalSpendGetReportResponseResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendGetReportResponseResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GlobalSpendGetReportResponseResponseArray{}),
		},
	)
}

type GlobalSpendGetReportResponseResponseArray []interface{}

func (r GlobalSpendGetReportResponseResponseArray) ImplementsGlobalSpendGetReportResponseResponseUnion() {
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type GlobalSpendGetReportResponseStartTimeUnion interface {
	ImplementsGlobalSpendGetReportResponseStartTimeUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GlobalSpendGetReportResponseStartTimeUnion)(nil)).Elem(),
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

type GlobalSpendListTagsParams struct {
	// Time till which to view key spend
	EndDate param.Field[string] `query:"end_date"`
	// Time from which to start viewing key spend
	StartDate param.Field[string] `query:"start_date"`
	// comman separated tags to filter on
	Tags param.Field[string] `query:"tags"`
}

// URLQuery serializes [GlobalSpendListTagsParams]'s query parameters as
// `url.Values`.
func (r GlobalSpendListTagsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GlobalSpendGetReportParams struct {
	// View spend for a specific api_key. Example api_key='sk-1234
	APIKey param.Field[string] `query:"api_key"`
	// View spend for a specific customer_id. Example customer_id='1234. Can be used in
	// conjunction with team_id as well.
	CustomerID param.Field[string] `query:"customer_id"`
	// Time till which to view spend
	EndDate param.Field[string] `query:"end_date"`
	// Group spend by internal team or customer or api_key
	GroupBy param.Field[GlobalSpendGetReportParamsGroupBy] `query:"group_by"`
	// View spend for a specific internal_user_id. Example internal_user_id='1234
	InternalUserID param.Field[string] `query:"internal_user_id"`
	// Time from which to start viewing spend
	StartDate param.Field[string] `query:"start_date"`
	// View spend for a specific team_id. Example team_id='1234
	TeamID param.Field[string] `query:"team_id"`
}

// URLQuery serializes [GlobalSpendGetReportParams]'s query parameters as
// `url.Values`.
func (r GlobalSpendGetReportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Group spend by internal team or customer or api_key
type GlobalSpendGetReportParamsGroupBy string

const (
	GlobalSpendGetReportParamsGroupByTeam     GlobalSpendGetReportParamsGroupBy = "team"
	GlobalSpendGetReportParamsGroupByCustomer GlobalSpendGetReportParamsGroupBy = "customer"
	GlobalSpendGetReportParamsGroupByAPIKey   GlobalSpendGetReportParamsGroupBy = "api_key"
)

func (r GlobalSpendGetReportParamsGroupBy) IsKnown() bool {
	switch r {
	case GlobalSpendGetReportParamsGroupByTeam, GlobalSpendGetReportParamsGroupByCustomer, GlobalSpendGetReportParamsGroupByAPIKey:
		return true
	}
	return false
}
