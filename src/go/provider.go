// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ProviderService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProviderService] method instead.
type ProviderService struct {
	Options []option.RequestOption
}

// NewProviderService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProviderService(opts ...option.RequestOption) (r *ProviderService) {
	r = &ProviderService{}
	r.Options = opts
	return
}

// Provider Budget Routing - Get Budget, Spend Details
// https://docs.hanzo.ai/docs/proxy/provider_budget_routing
//
// Use this endpoint to check current budget, spend and budget reset time for a
// provider
//
// # Example Request
//
// ```bash
// curl -X GET http://localhost:4000/provider/budgets     -H "Content-Type: application/json"     -H "Authorization: Bearer sk-1234"
// ```
//
// # Example Response
//
// ```json
//
//	{
//	  "providers": {
//	    "openai": {
//	      "budget_limit": 1e-12,
//	      "time_period": "1d",
//	      "spend": 0.0,
//	      "budget_reset_at": null
//	    },
//	    "azure": {
//	      "budget_limit": 100.0,
//	      "time_period": "1d",
//	      "spend": 0.0,
//	      "budget_reset_at": null
//	    },
//	    "anthropic": {
//	      "budget_limit": 100.0,
//	      "time_period": "10d",
//	      "spend": 0.0,
//	      "budget_reset_at": null
//	    },
//	    "vertex_ai": {
//	      "budget_limit": 100.0,
//	      "time_period": "12d",
//	      "spend": 0.0,
//	      "budget_reset_at": null
//	    }
//	  }
//	}
//
// ```
func (r *ProviderService) ListBudgets(ctx context.Context, opts ...option.RequestOption) (res *ProviderListBudgetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "provider/budgets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Complete provider budget configuration and status. Maps provider names to their
// budget configs.
type ProviderListBudgetsResponse struct {
	Providers map[string]ProviderListBudgetsResponseProvider `json:"providers"`
	JSON      providerListBudgetsResponseJSON                `json:"-"`
}

// providerListBudgetsResponseJSON contains the JSON metadata for the struct
// [ProviderListBudgetsResponse]
type providerListBudgetsResponseJSON struct {
	Providers   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProviderListBudgetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListBudgetsResponseJSON) RawJSON() string {
	return r.raw
}

// Configuration for a single provider's budget settings
type ProviderListBudgetsResponseProvider struct {
	BudgetLimit   float64                                 `json:"budget_limit,required,nullable"`
	TimePeriod    string                                  `json:"time_period,required,nullable"`
	BudgetResetAt string                                  `json:"budget_reset_at,nullable"`
	Spend         float64                                 `json:"spend,nullable"`
	JSON          providerListBudgetsResponseProviderJSON `json:"-"`
}

// providerListBudgetsResponseProviderJSON contains the JSON metadata for the
// struct [ProviderListBudgetsResponseProvider]
type providerListBudgetsResponseProviderJSON struct {
	BudgetLimit   apijson.Field
	TimePeriod    apijson.Field
	BudgetResetAt apijson.Field
	Spend         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ProviderListBudgetsResponseProvider) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r providerListBudgetsResponseProviderJSON) RawJSON() string {
	return r.raw
}
