// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/option"
)

// KeyRegenerateService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyRegenerateService] method instead.
type KeyRegenerateService struct {
	Options []option.RequestOption
}

// NewKeyRegenerateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewKeyRegenerateService(opts ...option.RequestOption) (r *KeyRegenerateService) {
	r = &KeyRegenerateService{}
	r.Options = opts
	return
}

type RegenerateKeyRequestParam struct {
	Aliases              param.Field[interface{}]   `json:"aliases"`
	AllowedCacheControls param.Field[[]interface{}] `json:"allowed_cache_controls"`
	Blocked              param.Field[bool]          `json:"blocked"`
	BudgetDuration       param.Field[string]        `json:"budget_duration"`
	BudgetID             param.Field[string]        `json:"budget_id"`
	Config               param.Field[interface{}]   `json:"config"`
	Duration             param.Field[string]        `json:"duration"`
	EnforcedParams       param.Field[[]string]      `json:"enforced_params"`
	Guardrails           param.Field[[]string]      `json:"guardrails"`
	Key                  param.Field[string]        `json:"key"`
	KeyAlias             param.Field[string]        `json:"key_alias"`
	MaxBudget            param.Field[float64]       `json:"max_budget"`
	MaxParallelRequests  param.Field[int64]         `json:"max_parallel_requests"`
	Metadata             param.Field[interface{}]   `json:"metadata"`
	ModelMaxBudget       param.Field[interface{}]   `json:"model_max_budget"`
	ModelRpmLimit        param.Field[interface{}]   `json:"model_rpm_limit"`
	ModelTpmLimit        param.Field[interface{}]   `json:"model_tpm_limit"`
	Models               param.Field[[]interface{}] `json:"models"`
	NewMasterKey         param.Field[string]        `json:"new_master_key"`
	Permissions          param.Field[interface{}]   `json:"permissions"`
	RpmLimit             param.Field[int64]         `json:"rpm_limit"`
	SendInviteEmail      param.Field[bool]          `json:"send_invite_email"`
	SoftBudget           param.Field[float64]       `json:"soft_budget"`
	Spend                param.Field[float64]       `json:"spend"`
	Tags                 param.Field[[]string]      `json:"tags"`
	TeamID               param.Field[string]        `json:"team_id"`
	TpmLimit             param.Field[int64]         `json:"tpm_limit"`
	UserID               param.Field[string]        `json:"user_id"`
}

func (r RegenerateKeyRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
