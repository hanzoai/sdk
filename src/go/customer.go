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

// CustomerService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerService] method instead.
type CustomerService struct {
	Options []option.RequestOption
}

// NewCustomerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerService(opts ...option.RequestOption) (r *CustomerService) {
	r = &CustomerService{}
	r.Options = opts
	return
}

// Allow creating a new Customer
//
// Parameters:
//
//   - user_id: str - The unique identifier for the user.
//   - alias: Optional[str] - A human-friendly alias for the user.
//   - blocked: bool - Flag to allow or disallow requests for this end-user. Default
//     is False.
//   - max_budget: Optional[float] - The maximum budget allocated to the user. Either
//     'max_budget' or 'budget_id' should be provided, not both.
//   - budget_id: Optional[str] - The identifier for an existing budget allocated to
//     the user. Either 'max_budget' or 'budget_id' should be provided, not both.
//   - allowed_model_region: Optional[Union[Literal["eu"], Literal["us"]]] - Require
//     all user requests to use models in this specific region.
//   - default_model: Optional[str] - If no equivalent model in the allowed region,
//     default all requests to this model.
//   - metadata: Optional[dict] = Metadata for customer, store information for
//     customer. Example metadata = {"data_training_opt_out": True}
//   - budget_duration: Optional[str] - Budget is reset at the end of specified
//     duration. If not set, budget is never reset. You can set duration as seconds
//     ("30s"), minutes ("30m"), hours ("30h"), days ("30d").
//   - tpm_limit: Optional[int] - [Not Implemented Yet] Specify tpm limit for a given
//     customer (Tokens per minute)
//   - rpm_limit: Optional[int] - [Not Implemented Yet] Specify rpm limit for a given
//     customer (Requests per minute)
//   - model_max_budget: Optional[dict] - [Not Implemented Yet] Specify max budget
//     for a given model. Example: {"openai/gpt-4o-mini": {"max_budget": 100.0,
//     "budget_duration": "1d"}}
//   - max_parallel_requests: Optional[int] - [Not Implemented Yet] Specify max
//     parallel requests for a given customer.
//   - soft_budget: Optional[float] - [Not Implemented Yet] Get alerts when customer
//     crosses given budget, doesn't block requests.
//
// - Allow specifying allowed regions
// - Allow specifying default model
//
// Example curl:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/customer/new'         --header 'Authorization: Bearer sk-1234'         --header 'Content-Type: application/json'         --data '{
//	        "user_id" : "z-jaff-3",
//	        "allowed_region": "eu",
//	        "budget_id": "free_tier",
//	        "default_model": "azure/gpt-3.5-turbo-eu" <- all calls from this user, use this model?
//	    }'
//
//	    # return end-user object
//
// ```
//
// NOTE: This used to be called `/end_user/new`, we will still be maintaining
// compatibility for /end_user/XXX for these endpoints
func (r *CustomerService) New(ctx context.Context, body CustomerNewParams, opts ...option.RequestOption) (res *CustomerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Example curl
//
// Parameters:
//
//   - user_id: str
//   - alias: Optional[str] = None # human-friendly alias
//   - blocked: bool = False # allow/disallow requests for this end-user
//   - max_budget: Optional[float] = None
//   - budget_id: Optional[str] = None # give either a budget_id or max_budget
//   - allowed_model_region: Optional[AllowedModelRegion] = ( None # require all user
//     requests to use models in this specific region )
//   - default_model: Optional[str] = ( None # if no equivalent model in allowed
//     region - default all requests to this model )
//
// Example curl:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/customer/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "user_id": "test-llm-user-4",
//	    "budget_id": "paid_tier"
//	}'
//
// See below for all params
// ```
func (r *CustomerService) Update(ctx context.Context, body CustomerUpdateParams, opts ...option.RequestOption) (res *CustomerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [Admin-only] List all available customers
//
// Example curl:
//
// ```
// curl --location --request GET 'http://0.0.0.0:4000/customer/list'         --header 'Authorization: Bearer sk-1234'
// ```
func (r *CustomerService) List(ctx context.Context, opts ...option.RequestOption) (res *[]CustomerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete multiple end-users.
//
// Parameters:
//
// - user_ids (List[str], required): The unique `user_id`s for the users to delete
//
// Example curl:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/customer/delete'         --header 'Authorization: Bearer sk-1234'         --header 'Content-Type: application/json'         --data '{
//	        "user_ids" :["z-jaff-5"]
//	}'
//
// See below for all params
// ```
func (r *CustomerService) Delete(ctx context.Context, body CustomerDeleteParams, opts ...option.RequestOption) (res *CustomerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [BETA] Reject calls with this end-user id
//
// Parameters:
//
// - user_ids (List[str], required): The unique `user_id`s for the users to block
//
//	(any /chat/completion call with this user={end-user-id} param, will be
//	rejected.)
//
//	```
//	curl -X POST "http://0.0.0.0:8000/user/block"
//	-H "Authorization: Bearer sk-1234"
//	-d '{
//	"user_ids": [<user_id>, ...]
//	}'
//	```
func (r *CustomerService) Block(ctx context.Context, body CustomerBlockParams, opts ...option.RequestOption) (res *CustomerBlockResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/block"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about an end-user. An `end_user` is a customer (external user)
// of the proxy.
//
// Parameters:
//
// - end_user_id (str, required): The unique identifier for the end-user
//
// Example curl:
//
// ```
// curl -X GET 'http://localhost:4000/customer/info?end_user_id=test-llm-user-4'         -H 'Authorization: Bearer sk-1234'
// ```
func (r *CustomerService) GetInfo(ctx context.Context, query CustomerGetInfoParams, opts ...option.RequestOption) (res *CustomerGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// [BETA] Unblock calls with this user id
//
// # Example
//
// ```
// curl -X POST "http://0.0.0.0:8000/user/unblock"
// -H "Authorization: Bearer sk-1234"
// -d '{
// "user_ids": [<user_id>, ...]
// }'
// ```
func (r *CustomerService) Unblock(ctx context.Context, body CustomerUnblockParams, opts ...option.RequestOption) (res *CustomerUnblockResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "customer/unblock"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockUsersParam struct {
	UserIDs param.Field[[]string] `json:"user_ids,required"`
}

func (r BlockUsersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewResponse = interface{}

type CustomerUpdateResponse = interface{}

type CustomerListResponse struct {
	Blocked            bool                                   `json:"blocked,required"`
	UserID             string                                 `json:"user_id,required"`
	Alias              string                                 `json:"alias,nullable"`
	AllowedModelRegion CustomerListResponseAllowedModelRegion `json:"allowed_model_region,nullable"`
	DefaultModel       string                                 `json:"default_model,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable CustomerListResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                            `json:"spend"`
	JSON           customerListResponseJSON           `json:"-"`
}

// customerListResponseJSON contains the JSON metadata for the struct
// [CustomerListResponse]
type customerListResponseJSON struct {
	Blocked            apijson.Field
	UserID             apijson.Field
	Alias              apijson.Field
	AllowedModelRegion apijson.Field
	DefaultModel       apijson.Field
	LlmBudgetTable     apijson.Field
	Spend              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerListResponseAllowedModelRegion string

const (
	CustomerListResponseAllowedModelRegionEu CustomerListResponseAllowedModelRegion = "eu"
	CustomerListResponseAllowedModelRegionUs CustomerListResponseAllowedModelRegion = "us"
)

func (r CustomerListResponseAllowedModelRegion) IsKnown() bool {
	switch r {
	case CustomerListResponseAllowedModelRegionEu, CustomerListResponseAllowedModelRegionUs:
		return true
	}
	return false
}

// Represents user-controllable params for a LLM_BudgetTable record
type CustomerListResponseLlmBudgetTable struct {
	BudgetDuration      string                                 `json:"budget_duration,nullable"`
	MaxBudget           float64                                `json:"max_budget,nullable"`
	MaxParallelRequests int64                                  `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                            `json:"model_max_budget,nullable"`
	RpmLimit            int64                                  `json:"rpm_limit,nullable"`
	SoftBudget          float64                                `json:"soft_budget,nullable"`
	TpmLimit            int64                                  `json:"tpm_limit,nullable"`
	JSON                customerListResponseLlmBudgetTableJSON `json:"-"`
}

// customerListResponseLlmBudgetTableJSON contains the JSON metadata for the struct
// [CustomerListResponseLlmBudgetTable]
type customerListResponseLlmBudgetTableJSON struct {
	BudgetDuration      apijson.Field
	MaxBudget           apijson.Field
	MaxParallelRequests apijson.Field
	ModelMaxBudget      apijson.Field
	RpmLimit            apijson.Field
	SoftBudget          apijson.Field
	TpmLimit            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerListResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerListResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type CustomerDeleteResponse = interface{}

type CustomerBlockResponse = interface{}

type CustomerGetInfoResponse struct {
	Blocked            bool                                      `json:"blocked,required"`
	UserID             string                                    `json:"user_id,required"`
	Alias              string                                    `json:"alias,nullable"`
	AllowedModelRegion CustomerGetInfoResponseAllowedModelRegion `json:"allowed_model_region,nullable"`
	DefaultModel       string                                    `json:"default_model,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable CustomerGetInfoResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                               `json:"spend"`
	JSON           customerGetInfoResponseJSON           `json:"-"`
}

// customerGetInfoResponseJSON contains the JSON metadata for the struct
// [CustomerGetInfoResponse]
type customerGetInfoResponseJSON struct {
	Blocked            apijson.Field
	UserID             apijson.Field
	Alias              apijson.Field
	AllowedModelRegion apijson.Field
	DefaultModel       apijson.Field
	LlmBudgetTable     apijson.Field
	Spend              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CustomerGetInfoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerGetInfoResponseJSON) RawJSON() string {
	return r.raw
}

type CustomerGetInfoResponseAllowedModelRegion string

const (
	CustomerGetInfoResponseAllowedModelRegionEu CustomerGetInfoResponseAllowedModelRegion = "eu"
	CustomerGetInfoResponseAllowedModelRegionUs CustomerGetInfoResponseAllowedModelRegion = "us"
)

func (r CustomerGetInfoResponseAllowedModelRegion) IsKnown() bool {
	switch r {
	case CustomerGetInfoResponseAllowedModelRegionEu, CustomerGetInfoResponseAllowedModelRegionUs:
		return true
	}
	return false
}

// Represents user-controllable params for a LLM_BudgetTable record
type CustomerGetInfoResponseLlmBudgetTable struct {
	BudgetDuration      string                                    `json:"budget_duration,nullable"`
	MaxBudget           float64                                   `json:"max_budget,nullable"`
	MaxParallelRequests int64                                     `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                               `json:"model_max_budget,nullable"`
	RpmLimit            int64                                     `json:"rpm_limit,nullable"`
	SoftBudget          float64                                   `json:"soft_budget,nullable"`
	TpmLimit            int64                                     `json:"tpm_limit,nullable"`
	JSON                customerGetInfoResponseLlmBudgetTableJSON `json:"-"`
}

// customerGetInfoResponseLlmBudgetTableJSON contains the JSON metadata for the
// struct [CustomerGetInfoResponseLlmBudgetTable]
type customerGetInfoResponseLlmBudgetTableJSON struct {
	BudgetDuration      apijson.Field
	MaxBudget           apijson.Field
	MaxParallelRequests apijson.Field
	ModelMaxBudget      apijson.Field
	RpmLimit            apijson.Field
	SoftBudget          apijson.Field
	TpmLimit            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *CustomerGetInfoResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customerGetInfoResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type CustomerUnblockResponse = interface{}

type CustomerNewParams struct {
	UserID             param.Field[string]                              `json:"user_id,required"`
	Alias              param.Field[string]                              `json:"alias"`
	AllowedModelRegion param.Field[CustomerNewParamsAllowedModelRegion] `json:"allowed_model_region"`
	Blocked            param.Field[bool]                                `json:"blocked"`
	// Max duration budget should be set for (e.g. '1hr', '1d', '28d')
	BudgetDuration param.Field[string] `json:"budget_duration"`
	BudgetID       param.Field[string] `json:"budget_id"`
	DefaultModel   param.Field[string] `json:"default_model"`
	// Requests will fail if this budget (in USD) is exceeded.
	MaxBudget param.Field[float64] `json:"max_budget"`
	// Max concurrent requests allowed for this budget id.
	MaxParallelRequests param.Field[int64] `json:"max_parallel_requests"`
	// Max budget for each model (e.g. {'gpt-4o': {'max_budget': '0.0000001',
	// 'budget_duration': '1d', 'tpm_limit': 1000, 'rpm_limit': 1000}})
	ModelMaxBudget param.Field[map[string]CustomerNewParamsModelMaxBudget] `json:"model_max_budget"`
	// Max requests per minute, allowed for this budget id.
	RpmLimit param.Field[int64] `json:"rpm_limit"`
	// Requests will NOT fail if this is exceeded. Will fire alerting though.
	SoftBudget param.Field[float64] `json:"soft_budget"`
	// Max tokens per minute, allowed for this budget id.
	TpmLimit param.Field[int64] `json:"tpm_limit"`
}

func (r CustomerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerNewParamsAllowedModelRegion string

const (
	CustomerNewParamsAllowedModelRegionEu CustomerNewParamsAllowedModelRegion = "eu"
	CustomerNewParamsAllowedModelRegionUs CustomerNewParamsAllowedModelRegion = "us"
)

func (r CustomerNewParamsAllowedModelRegion) IsKnown() bool {
	switch r {
	case CustomerNewParamsAllowedModelRegionEu, CustomerNewParamsAllowedModelRegionUs:
		return true
	}
	return false
}

type CustomerNewParamsModelMaxBudget struct {
	BudgetDuration param.Field[string]  `json:"budget_duration"`
	MaxBudget      param.Field[float64] `json:"max_budget"`
	RpmLimit       param.Field[int64]   `json:"rpm_limit"`
	TpmLimit       param.Field[int64]   `json:"tpm_limit"`
}

func (r CustomerNewParamsModelMaxBudget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParams struct {
	UserID             param.Field[string]                                 `json:"user_id,required"`
	Alias              param.Field[string]                                 `json:"alias"`
	AllowedModelRegion param.Field[CustomerUpdateParamsAllowedModelRegion] `json:"allowed_model_region"`
	Blocked            param.Field[bool]                                   `json:"blocked"`
	BudgetID           param.Field[string]                                 `json:"budget_id"`
	DefaultModel       param.Field[string]                                 `json:"default_model"`
	MaxBudget          param.Field[float64]                                `json:"max_budget"`
}

func (r CustomerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerUpdateParamsAllowedModelRegion string

const (
	CustomerUpdateParamsAllowedModelRegionEu CustomerUpdateParamsAllowedModelRegion = "eu"
	CustomerUpdateParamsAllowedModelRegionUs CustomerUpdateParamsAllowedModelRegion = "us"
)

func (r CustomerUpdateParamsAllowedModelRegion) IsKnown() bool {
	switch r {
	case CustomerUpdateParamsAllowedModelRegionEu, CustomerUpdateParamsAllowedModelRegionUs:
		return true
	}
	return false
}

type CustomerDeleteParams struct {
	UserIDs param.Field[[]string] `json:"user_ids,required"`
}

func (r CustomerDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomerBlockParams struct {
	BlockUsers BlockUsersParam `json:"block_users,required"`
}

func (r CustomerBlockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockUsers)
}

type CustomerGetInfoParams struct {
	// End User ID in the request parameters
	EndUserID param.Field[string] `query:"end_user_id,required"`
}

// URLQuery serializes [CustomerGetInfoParams]'s query parameters as `url.Values`.
func (r CustomerGetInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type CustomerUnblockParams struct {
	BlockUsers BlockUsersParam `json:"block_users,required"`
}

func (r CustomerUnblockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockUsers)
}
