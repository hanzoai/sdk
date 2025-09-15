// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// OrganizationInfoService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationInfoService] method instead.
type OrganizationInfoService struct {
	Options []option.RequestOption
}

// NewOrganizationInfoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrganizationInfoService(opts ...option.RequestOption) (r *OrganizationInfoService) {
	r = &OrganizationInfoService{}
	r.Options = opts
	return
}

// Get the org specific information
func (r *OrganizationInfoService) Get(ctx context.Context, query OrganizationInfoGetParams, opts ...option.RequestOption) (res *OrganizationInfoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// DEPRECATED: Use GET /organization/info instead
func (r *OrganizationInfoService) Deprecated(ctx context.Context, body OrganizationInfoDeprecatedParams, opts ...option.RequestOption) (res *OrganizationInfoDeprecatedResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returned by the /organization/info endpoint and /organization/list endpoint
type OrganizationInfoGetResponse struct {
	BudgetID  string    `json:"budget_id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy string    `json:"created_by,required"`
	Models    []string  `json:"models,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	UpdatedBy string    `json:"updated_by,required"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable    OrganizationInfoGetResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Members           []OrganizationInfoGetResponseMember       `json:"members"`
	Metadata          interface{}                               `json:"metadata,nullable"`
	OrganizationAlias string                                    `json:"organization_alias,nullable"`
	OrganizationID    string                                    `json:"organization_id,nullable"`
	Spend             float64                                   `json:"spend"`
	Teams             []OrganizationInfoGetResponseTeam         `json:"teams"`
	JSON              organizationInfoGetResponseJSON           `json:"-"`
}

// organizationInfoGetResponseJSON contains the JSON metadata for the struct
// [OrganizationInfoGetResponse]
type organizationInfoGetResponseJSON struct {
	BudgetID          apijson.Field
	CreatedAt         apijson.Field
	CreatedBy         apijson.Field
	Models            apijson.Field
	UpdatedAt         apijson.Field
	UpdatedBy         apijson.Field
	LlmBudgetTable    apijson.Field
	Members           apijson.Field
	Metadata          apijson.Field
	OrganizationAlias apijson.Field
	OrganizationID    apijson.Field
	Spend             apijson.Field
	Teams             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationInfoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInfoGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationInfoGetResponseLlmBudgetTable struct {
	BudgetDuration      string                                        `json:"budget_duration,nullable"`
	MaxBudget           float64                                       `json:"max_budget,nullable"`
	MaxParallelRequests int64                                         `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                   `json:"model_max_budget,nullable"`
	RpmLimit            int64                                         `json:"rpm_limit,nullable"`
	SoftBudget          float64                                       `json:"soft_budget,nullable"`
	TpmLimit            int64                                         `json:"tpm_limit,nullable"`
	JSON                organizationInfoGetResponseLlmBudgetTableJSON `json:"-"`
}

// organizationInfoGetResponseLlmBudgetTableJSON contains the JSON metadata for the
// struct [OrganizationInfoGetResponseLlmBudgetTable]
type organizationInfoGetResponseLlmBudgetTableJSON struct {
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

func (r *OrganizationInfoGetResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInfoGetResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationInfoGetResponseMember struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationInfoGetResponseMembersLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                          `json:"spend"`
	User           interface{}                                      `json:"user"`
	UserRole       string                                           `json:"user_role,nullable"`
	JSON           organizationInfoGetResponseMemberJSON            `json:"-"`
}

// organizationInfoGetResponseMemberJSON contains the JSON metadata for the struct
// [OrganizationInfoGetResponseMember]
type organizationInfoGetResponseMemberJSON struct {
	CreatedAt      apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	UserID         apijson.Field
	BudgetID       apijson.Field
	LlmBudgetTable apijson.Field
	Spend          apijson.Field
	User           apijson.Field
	UserRole       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationInfoGetResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInfoGetResponseMemberJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationInfoGetResponseMembersLlmBudgetTable struct {
	BudgetDuration      string                                               `json:"budget_duration,nullable"`
	MaxBudget           float64                                              `json:"max_budget,nullable"`
	MaxParallelRequests int64                                                `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                          `json:"model_max_budget,nullable"`
	RpmLimit            int64                                                `json:"rpm_limit,nullable"`
	SoftBudget          float64                                              `json:"soft_budget,nullable"`
	TpmLimit            int64                                                `json:"tpm_limit,nullable"`
	JSON                organizationInfoGetResponseMembersLlmBudgetTableJSON `json:"-"`
}

// organizationInfoGetResponseMembersLlmBudgetTableJSON contains the JSON metadata
// for the struct [OrganizationInfoGetResponseMembersLlmBudgetTable]
type organizationInfoGetResponseMembersLlmBudgetTableJSON struct {
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

func (r *OrganizationInfoGetResponseMembersLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInfoGetResponseMembersLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationInfoGetResponseTeam struct {
	TeamID              string                                        `json:"team_id,required"`
	Admins              []interface{}                                 `json:"admins"`
	Blocked             bool                                          `json:"blocked"`
	BudgetDuration      string                                        `json:"budget_duration,nullable"`
	BudgetResetAt       time.Time                                     `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt           time.Time                                     `json:"created_at,nullable" format:"date-time"`
	LlmModelTable       OrganizationInfoGetResponseTeamsLlmModelTable `json:"llm_model_table,nullable"`
	MaxBudget           float64                                       `json:"max_budget,nullable"`
	MaxParallelRequests int64                                         `json:"max_parallel_requests,nullable"`
	Members             []interface{}                                 `json:"members"`
	MembersWithRoles    []Member                                      `json:"members_with_roles"`
	Metadata            interface{}                                   `json:"metadata,nullable"`
	ModelID             int64                                         `json:"model_id,nullable"`
	Models              []interface{}                                 `json:"models"`
	OrganizationID      string                                        `json:"organization_id,nullable"`
	RpmLimit            int64                                         `json:"rpm_limit,nullable"`
	Spend               float64                                       `json:"spend,nullable"`
	TeamAlias           string                                        `json:"team_alias,nullable"`
	TpmLimit            int64                                         `json:"tpm_limit,nullable"`
	JSON                organizationInfoGetResponseTeamJSON           `json:"-"`
}

// organizationInfoGetResponseTeamJSON contains the JSON metadata for the struct
// [OrganizationInfoGetResponseTeam]
type organizationInfoGetResponseTeamJSON struct {
	TeamID              apijson.Field
	Admins              apijson.Field
	Blocked             apijson.Field
	BudgetDuration      apijson.Field
	BudgetResetAt       apijson.Field
	CreatedAt           apijson.Field
	LlmModelTable       apijson.Field
	MaxBudget           apijson.Field
	MaxParallelRequests apijson.Field
	Members             apijson.Field
	MembersWithRoles    apijson.Field
	Metadata            apijson.Field
	ModelID             apijson.Field
	Models              apijson.Field
	OrganizationID      apijson.Field
	RpmLimit            apijson.Field
	Spend               apijson.Field
	TeamAlias           apijson.Field
	TpmLimit            apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OrganizationInfoGetResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInfoGetResponseTeamJSON) RawJSON() string {
	return r.raw
}

type OrganizationInfoGetResponseTeamsLlmModelTable struct {
	CreatedBy    string                                            `json:"created_by,required"`
	UpdatedBy    string                                            `json:"updated_by,required"`
	ModelAliases interface{}                                       `json:"model_aliases,nullable"`
	JSON         organizationInfoGetResponseTeamsLlmModelTableJSON `json:"-"`
}

// organizationInfoGetResponseTeamsLlmModelTableJSON contains the JSON metadata for
// the struct [OrganizationInfoGetResponseTeamsLlmModelTable]
type organizationInfoGetResponseTeamsLlmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ModelAliases apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationInfoGetResponseTeamsLlmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationInfoGetResponseTeamsLlmModelTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationInfoDeprecatedResponse = interface{}

type OrganizationInfoGetParams struct {
	OrganizationID param.Field[string] `query:"organization_id,required"`
}

// URLQuery serializes [OrganizationInfoGetParams]'s query parameters as
// `url.Values`.
func (r OrganizationInfoGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrganizationInfoDeprecatedParams struct {
	Organizations param.Field[[]string] `json:"organizations,required"`
}

func (r OrganizationInfoDeprecatedParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
