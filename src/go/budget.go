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

// BudgetService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBudgetService] method instead.
type BudgetService struct {
	Options []option.RequestOption
}

// NewBudgetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBudgetService(opts ...option.RequestOption) (r *BudgetService) {
	r = &BudgetService{}
	r.Options = opts
	return
}

// Create a new budget object. Can apply this to teams, orgs, end-users, keys.
//
// Parameters:
//
//   - budget_duration: Optional[str] - Budget reset period ("30d", "1h", etc.)
//   - budget_id: Optional[str] - The id of the budget. If not provided, a new id
//     will be generated.
//   - max_budget: Optional[float] - The max budget for the budget.
//   - soft_budget: Optional[float] - The soft budget for the budget.
//   - max_parallel_requests: Optional[int] - The max number of parallel requests for
//     the budget.
//   - tpm_limit: Optional[int] - The tokens per minute limit for the budget.
//   - rpm_limit: Optional[int] - The requests per minute limit for the budget.
//   - model_max_budget: Optional[dict] - Specify max budget for a given model.
//     Example: {"openai/gpt-4o-mini": {"max_budget": 100.0, "budget_duration": "1d",
//     "tpm_limit": 100000, "rpm_limit": 100000}}
func (r *BudgetService) New(ctx context.Context, body BudgetNewParams, opts ...option.RequestOption) (res *BudgetNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "budget/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an existing budget object.
//
// Parameters:
//
//   - budget_duration: Optional[str] - Budget reset period ("30d", "1h", etc.)
//   - budget_id: Optional[str] - The id of the budget. If not provided, a new id
//     will be generated.
//   - max_budget: Optional[float] - The max budget for the budget.
//   - soft_budget: Optional[float] - The soft budget for the budget.
//   - max_parallel_requests: Optional[int] - The max number of parallel requests for
//     the budget.
//   - tpm_limit: Optional[int] - The tokens per minute limit for the budget.
//   - rpm_limit: Optional[int] - The requests per minute limit for the budget.
//   - model_max_budget: Optional[dict] - Specify max budget for a given model.
//     Example: {"openai/gpt-4o-mini": {"max_budget": 100.0, "budget_duration": "1d",
//     "tpm_limit": 100000, "rpm_limit": 100000}}
func (r *BudgetService) Update(ctx context.Context, body BudgetUpdateParams, opts ...option.RequestOption) (res *BudgetUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "budget/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all the created budgets in proxy db. Used on Admin UI.
func (r *BudgetService) List(ctx context.Context, opts ...option.RequestOption) (res *BudgetListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "budget/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete budget
//
// Parameters:
//
// - id: str - The budget id to delete
func (r *BudgetService) Delete(ctx context.Context, body BudgetDeleteParams, opts ...option.RequestOption) (res *BudgetDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "budget/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the budget id specific information
//
// Parameters:
//
// - budgets: List[str] - The list of budget ids to get information for
func (r *BudgetService) Info(ctx context.Context, body BudgetInfoParams, opts ...option.RequestOption) (res *BudgetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "budget/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get list of configurable params + current value for a budget item + description
// of each field
//
// Used on Admin UI.
//
// Query Parameters:
//
// - budget_id: str - The budget id to get information for
func (r *BudgetService) Settings(ctx context.Context, query BudgetSettingsParams, opts ...option.RequestOption) (res *BudgetSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "budget/settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type BudgetNewParam struct {
	// Max duration budget should be set for (e.g. '1hr', '1d', '28d')
	BudgetDuration param.Field[string] `json:"budget_duration"`
	// The unique budget id.
	BudgetID param.Field[string] `json:"budget_id"`
	// Requests will fail if this budget (in USD) is exceeded.
	MaxBudget param.Field[float64] `json:"max_budget"`
	// Max concurrent requests allowed for this budget id.
	MaxParallelRequests param.Field[int64] `json:"max_parallel_requests"`
	// Max budget for each model (e.g. {'gpt-4o': {'max_budget': '0.0000001',
	// 'budget_duration': '1d', 'tpm_limit': 1000, 'rpm_limit': 1000}})
	ModelMaxBudget param.Field[map[string]BudgetNewModelMaxBudgetParam] `json:"model_max_budget"`
	// Max requests per minute, allowed for this budget id.
	RpmLimit param.Field[int64] `json:"rpm_limit"`
	// Requests will NOT fail if this is exceeded. Will fire alerting though.
	SoftBudget param.Field[float64] `json:"soft_budget"`
	// Max tokens per minute, allowed for this budget id.
	TpmLimit param.Field[int64] `json:"tpm_limit"`
}

func (r BudgetNewParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetNewModelMaxBudgetParam struct {
	BudgetDuration param.Field[string]  `json:"budget_duration"`
	MaxBudget      param.Field[float64] `json:"max_budget"`
	RpmLimit       param.Field[int64]   `json:"rpm_limit"`
	TpmLimit       param.Field[int64]   `json:"tpm_limit"`
}

func (r BudgetNewModelMaxBudgetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetNewResponse = interface{}

type BudgetUpdateResponse = interface{}

type BudgetListResponse = interface{}

type BudgetDeleteResponse = interface{}

type BudgetInfoResponse = interface{}

type BudgetSettingsResponse = interface{}

type BudgetNewParams struct {
	BudgetNew BudgetNewParam `json:"budget_new,required"`
}

func (r BudgetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BudgetNew)
}

type BudgetUpdateParams struct {
	BudgetNew BudgetNewParam `json:"budget_new,required"`
}

func (r BudgetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BudgetNew)
}

type BudgetDeleteParams struct {
	ID param.Field[string] `json:"id,required"`
}

func (r BudgetDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetInfoParams struct {
	Budgets param.Field[[]string] `json:"budgets,required"`
}

func (r BudgetInfoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BudgetSettingsParams struct {
	BudgetID param.Field[string] `query:"budget_id,required"`
}

// URLQuery serializes [BudgetSettingsParams]'s query parameters as `url.Values`.
func (r BudgetSettingsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
