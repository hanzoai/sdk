// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// OrganizationService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	Info    *OrganizationInfoService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.Info = NewOrganizationInfoService(opts...)
	return
}

// Allow orgs to own teams
//
// Set org level budgets + model access.
//
// Only admins can create orgs.
//
// # Parameters
//
//   - organization_alias: _str_ - The name of the organization.
//   - models: _List_ - The models the organization has access to.
//   - budget_id: _Optional[str]_ - The id for a budget (tpm/rpm/max budget) for the
//     organization.
//
// ### IF NO BUDGET ID - CREATE ONE WITH THESE PARAMS
//
//   - max_budget: _Optional[float]_ - Max budget for org
//   - tpm_limit: _Optional[int]_ - Max tpm limit for org
//   - rpm_limit: _Optional[int]_ - Max rpm limit for org
//   - max_parallel_requests: _Optional[int]_ - [Not Implemented Yet] Max parallel
//     requests for org
//   - soft_budget: _Optional[float]_ - [Not Implemented Yet] Get a slack alert when
//     this soft budget is reached. Don't block requests.
//   - model_max_budget: _Optional[dict]_ - Max budget for a specific model
//   - budget_duration: _Optional[str]_ - Frequency of reseting org budget
//   - metadata: _Optional[dict]_ - Metadata for organization, store information for
//     organization. Example metadata - {"extra_info": "some info"}
//   - blocked: _bool_ - Flag indicating if the org is blocked or not - will stop all
//     calls from keys with this org_id.
//   - tags: _Optional[List[str]]_ - Tags for
//     [tracking spend](https://llm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://llm.vercel.app/docs/proxy/tag_routing).
//   - organization_id: _Optional[str]_ - The organization id of the team. Default is
//     None. Create via `/organization/new`.
//   - model_aliases: Optional[dict] - Model aliases for the team.
//     [Docs](https://docs.hanzo.ai/docs/proxy/team_based_routing#create-team-with-model-alias)
//
// Case 1: Create new org **without** a budget_id
//
// ```bash
// curl --location 'http://0.0.0.0:4000/organization/new'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data '{
//	    "organization_alias": "my-secret-org",
//	    "models": ["model1", "model2"],
//	    "max_budget": 100
//	}'
//
// ```
//
// Case 2: Create new org **with** a budget_id
//
// ```bash
// curl --location 'http://0.0.0.0:4000/organization/new'
// --header 'Authorization: Bearer sk-1234'
// --header 'Content-Type: application/json'
//
//	--data '{
//	    "organization_alias": "my-secret-org",
//	    "models": ["model1", "model2"],
//	    "budget_id": "428eeaa8-f3ac-4e85-a8fb-7dc8d7aa8689"
//	}'
//
// ```
func (r *OrganizationService) New(ctx context.Context, body OrganizationNewParams, opts ...option.RequestOption) (res *OrganizationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Update an organization
func (r *OrganizationService) Update(ctx context.Context, body OrganizationUpdateParams, opts ...option.RequestOption) (res *OrganizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// ```
// curl --location --request GET 'http://0.0.0.0:4000/organization/list'         --header 'Authorization: Bearer sk-1234'
// ```
func (r *OrganizationService) List(ctx context.Context, opts ...option.RequestOption) (res *[]OrganizationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an organization
//
// # Parameters:
//
// - organization_ids: List[str] - The organization ids to delete.
func (r *OrganizationService) Delete(ctx context.Context, body OrganizationDeleteParams, opts ...option.RequestOption) (res *[]OrganizationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// [BETA]
//
// Add new members (either via user_email or user_id) to an organization
//
// If user doesn't exist, new user row will also be added to User Table
//
// Only proxy_admin or org_admin of organization, allowed to access this endpoint.
//
// # Parameters:
//
// - organization_id: str (required)
// - member: Union[List[Member], Member] (required)
//   - role: Literal[LLMUserRoles] (required)
//   - user_id: Optional[str]
//   - user_email: Optional[str]
//
// Note: Either user_id or user_email must be provided for each member.
//
// Example:
//
// ```
//
//	curl -X POST 'http://0.0.0.0:4000/organization/member_add'     -H 'Authorization: Bearer sk-1234'     -H 'Content-Type: application/json'     -d '{
//	    "organization_id": "45e3e396-ee08-4a61-a88e-16b3ce7e0849",
//	    "member": {
//	        "role": "internal_user",
//	        "user_id": "dev247652@hanzo.ai"
//	    },
//	    "max_budget_in_organization": 100.0
//	}'
//
// ```
//
// The following is executed in this function:
//
//  1. Check if organization exists
//  2. Creates a new Internal User if the user_id or user_email is not found in
//     LLM_UserTable
//  3. Add Internal User to the `LLM_OrganizationMembership` table
func (r *OrganizationService) AddMember(ctx context.Context, body OrganizationAddMemberParams, opts ...option.RequestOption) (res *OrganizationAddMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/member_add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a member from an organization
func (r *OrganizationService) DeleteMember(ctx context.Context, body OrganizationDeleteMemberParams, opts ...option.RequestOption) (res *OrganizationDeleteMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/member_delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update a member's role in an organization
func (r *OrganizationService) UpdateMember(ctx context.Context, body OrganizationUpdateMemberParams, opts ...option.RequestOption) (res *OrganizationUpdateMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "organization/member_update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

type OrgMemberParam struct {
	Role      param.Field[OrgMemberRole] `json:"role,required"`
	UserEmail param.Field[string]        `json:"user_email"`
	UserID    param.Field[string]        `json:"user_id"`
}

func (r OrgMemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r OrgMemberParam) implementsOrganizationAddMemberParamsMemberUnion() {}

type OrgMemberRole string

const (
	OrgMemberRoleOrgAdmin           OrgMemberRole = "org_admin"
	OrgMemberRoleInternalUser       OrgMemberRole = "internal_user"
	OrgMemberRoleInternalUserViewer OrgMemberRole = "internal_user_viewer"
)

func (r OrgMemberRole) IsKnown() bool {
	switch r {
	case OrgMemberRoleOrgAdmin, OrgMemberRoleInternalUser, OrgMemberRoleInternalUserViewer:
		return true
	}
	return false
}

type OrganizationNewResponse struct {
	BudgetID          string                      `json:"budget_id,required"`
	CreatedAt         time.Time                   `json:"created_at,required" format:"date-time"`
	CreatedBy         string                      `json:"created_by,required"`
	Models            []string                    `json:"models,required"`
	OrganizationID    string                      `json:"organization_id,required"`
	UpdatedAt         time.Time                   `json:"updated_at,required" format:"date-time"`
	UpdatedBy         string                      `json:"updated_by,required"`
	Metadata          interface{}                 `json:"metadata,nullable"`
	OrganizationAlias string                      `json:"organization_alias,nullable"`
	Spend             float64                     `json:"spend"`
	JSON              organizationNewResponseJSON `json:"-"`
}

// organizationNewResponseJSON contains the JSON metadata for the struct
// [OrganizationNewResponse]
type organizationNewResponseJSON struct {
	BudgetID          apijson.Field
	CreatedAt         apijson.Field
	CreatedBy         apijson.Field
	Models            apijson.Field
	OrganizationID    apijson.Field
	UpdatedAt         apijson.Field
	UpdatedBy         apijson.Field
	Metadata          apijson.Field
	OrganizationAlias apijson.Field
	Spend             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrganizationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseJSON) RawJSON() string {
	return r.raw
}

// Returned by the /organization/info endpoint and /organization/list endpoint
type OrganizationUpdateResponse struct {
	BudgetID  string    `json:"budget_id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy string    `json:"created_by,required"`
	Models    []string  `json:"models,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	UpdatedBy string    `json:"updated_by,required"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable    OrganizationUpdateResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Members           []OrganizationUpdateResponseMember       `json:"members"`
	Metadata          interface{}                              `json:"metadata,nullable"`
	OrganizationAlias string                                   `json:"organization_alias,nullable"`
	OrganizationID    string                                   `json:"organization_id,nullable"`
	Spend             float64                                  `json:"spend"`
	Teams             []OrganizationUpdateResponseTeam         `json:"teams"`
	JSON              organizationUpdateResponseJSON           `json:"-"`
}

// organizationUpdateResponseJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponse]
type organizationUpdateResponseJSON struct {
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

func (r *OrganizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationUpdateResponseLlmBudgetTable struct {
	BudgetDuration      string                                       `json:"budget_duration,nullable"`
	MaxBudget           float64                                      `json:"max_budget,nullable"`
	MaxParallelRequests int64                                        `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                  `json:"model_max_budget,nullable"`
	RpmLimit            int64                                        `json:"rpm_limit,nullable"`
	SoftBudget          float64                                      `json:"soft_budget,nullable"`
	TpmLimit            int64                                        `json:"tpm_limit,nullable"`
	JSON                organizationUpdateResponseLlmBudgetTableJSON `json:"-"`
}

// organizationUpdateResponseLlmBudgetTableJSON contains the JSON metadata for the
// struct [OrganizationUpdateResponseLlmBudgetTable]
type organizationUpdateResponseLlmBudgetTableJSON struct {
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

func (r *OrganizationUpdateResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationUpdateResponseMember struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationUpdateResponseMembersLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                         `json:"spend"`
	User           interface{}                                     `json:"user"`
	UserRole       string                                          `json:"user_role,nullable"`
	JSON           organizationUpdateResponseMemberJSON            `json:"-"`
}

// organizationUpdateResponseMemberJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseMember]
type organizationUpdateResponseMemberJSON struct {
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

func (r *OrganizationUpdateResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseMemberJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationUpdateResponseMembersLlmBudgetTable struct {
	BudgetDuration      string                                              `json:"budget_duration,nullable"`
	MaxBudget           float64                                             `json:"max_budget,nullable"`
	MaxParallelRequests int64                                               `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                         `json:"model_max_budget,nullable"`
	RpmLimit            int64                                               `json:"rpm_limit,nullable"`
	SoftBudget          float64                                             `json:"soft_budget,nullable"`
	TpmLimit            int64                                               `json:"tpm_limit,nullable"`
	JSON                organizationUpdateResponseMembersLlmBudgetTableJSON `json:"-"`
}

// organizationUpdateResponseMembersLlmBudgetTableJSON contains the JSON metadata
// for the struct [OrganizationUpdateResponseMembersLlmBudgetTable]
type organizationUpdateResponseMembersLlmBudgetTableJSON struct {
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

func (r *OrganizationUpdateResponseMembersLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseMembersLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseTeam struct {
	TeamID              string                                       `json:"team_id,required"`
	Admins              []interface{}                                `json:"admins"`
	Blocked             bool                                         `json:"blocked"`
	BudgetDuration      string                                       `json:"budget_duration,nullable"`
	BudgetResetAt       time.Time                                    `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt           time.Time                                    `json:"created_at,nullable" format:"date-time"`
	LlmModelTable       OrganizationUpdateResponseTeamsLlmModelTable `json:"llm_model_table,nullable"`
	MaxBudget           float64                                      `json:"max_budget,nullable"`
	MaxParallelRequests int64                                        `json:"max_parallel_requests,nullable"`
	Members             []interface{}                                `json:"members"`
	MembersWithRoles    []Member                                     `json:"members_with_roles"`
	Metadata            interface{}                                  `json:"metadata,nullable"`
	ModelID             int64                                        `json:"model_id,nullable"`
	Models              []interface{}                                `json:"models"`
	OrganizationID      string                                       `json:"organization_id,nullable"`
	RpmLimit            int64                                        `json:"rpm_limit,nullable"`
	Spend               float64                                      `json:"spend,nullable"`
	TeamAlias           string                                       `json:"team_alias,nullable"`
	TpmLimit            int64                                        `json:"tpm_limit,nullable"`
	JSON                organizationUpdateResponseTeamJSON           `json:"-"`
}

// organizationUpdateResponseTeamJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseTeam]
type organizationUpdateResponseTeamJSON struct {
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

func (r *OrganizationUpdateResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseTeamJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseTeamsLlmModelTable struct {
	CreatedBy    string                                           `json:"created_by,required"`
	UpdatedBy    string                                           `json:"updated_by,required"`
	ModelAliases interface{}                                      `json:"model_aliases,nullable"`
	JSON         organizationUpdateResponseTeamsLlmModelTableJSON `json:"-"`
}

// organizationUpdateResponseTeamsLlmModelTableJSON contains the JSON metadata for
// the struct [OrganizationUpdateResponseTeamsLlmModelTable]
type organizationUpdateResponseTeamsLlmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ModelAliases apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationUpdateResponseTeamsLlmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseTeamsLlmModelTableJSON) RawJSON() string {
	return r.raw
}

// Returned by the /organization/info endpoint and /organization/list endpoint
type OrganizationListResponse struct {
	BudgetID  string    `json:"budget_id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy string    `json:"created_by,required"`
	Models    []string  `json:"models,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	UpdatedBy string    `json:"updated_by,required"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable    OrganizationListResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Members           []OrganizationListResponseMember       `json:"members"`
	Metadata          interface{}                            `json:"metadata,nullable"`
	OrganizationAlias string                                 `json:"organization_alias,nullable"`
	OrganizationID    string                                 `json:"organization_id,nullable"`
	Spend             float64                                `json:"spend"`
	Teams             []OrganizationListResponseTeam         `json:"teams"`
	JSON              organizationListResponseJSON           `json:"-"`
}

// organizationListResponseJSON contains the JSON metadata for the struct
// [OrganizationListResponse]
type organizationListResponseJSON struct {
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

func (r *OrganizationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationListResponseLlmBudgetTable struct {
	BudgetDuration      string                                     `json:"budget_duration,nullable"`
	MaxBudget           float64                                    `json:"max_budget,nullable"`
	MaxParallelRequests int64                                      `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                `json:"model_max_budget,nullable"`
	RpmLimit            int64                                      `json:"rpm_limit,nullable"`
	SoftBudget          float64                                    `json:"soft_budget,nullable"`
	TpmLimit            int64                                      `json:"tpm_limit,nullable"`
	JSON                organizationListResponseLlmBudgetTableJSON `json:"-"`
}

// organizationListResponseLlmBudgetTableJSON contains the JSON metadata for the
// struct [OrganizationListResponseLlmBudgetTable]
type organizationListResponseLlmBudgetTableJSON struct {
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

func (r *OrganizationListResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationListResponseMember struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationListResponseMembersLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                       `json:"spend"`
	User           interface{}                                   `json:"user"`
	UserRole       string                                        `json:"user_role,nullable"`
	JSON           organizationListResponseMemberJSON            `json:"-"`
}

// organizationListResponseMemberJSON contains the JSON metadata for the struct
// [OrganizationListResponseMember]
type organizationListResponseMemberJSON struct {
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

func (r *OrganizationListResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseMemberJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationListResponseMembersLlmBudgetTable struct {
	BudgetDuration      string                                            `json:"budget_duration,nullable"`
	MaxBudget           float64                                           `json:"max_budget,nullable"`
	MaxParallelRequests int64                                             `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                       `json:"model_max_budget,nullable"`
	RpmLimit            int64                                             `json:"rpm_limit,nullable"`
	SoftBudget          float64                                           `json:"soft_budget,nullable"`
	TpmLimit            int64                                             `json:"tpm_limit,nullable"`
	JSON                organizationListResponseMembersLlmBudgetTableJSON `json:"-"`
}

// organizationListResponseMembersLlmBudgetTableJSON contains the JSON metadata for
// the struct [OrganizationListResponseMembersLlmBudgetTable]
type organizationListResponseMembersLlmBudgetTableJSON struct {
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

func (r *OrganizationListResponseMembersLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseMembersLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseTeam struct {
	TeamID              string                                     `json:"team_id,required"`
	Admins              []interface{}                              `json:"admins"`
	Blocked             bool                                       `json:"blocked"`
	BudgetDuration      string                                     `json:"budget_duration,nullable"`
	BudgetResetAt       time.Time                                  `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt           time.Time                                  `json:"created_at,nullable" format:"date-time"`
	LlmModelTable       OrganizationListResponseTeamsLlmModelTable `json:"llm_model_table,nullable"`
	MaxBudget           float64                                    `json:"max_budget,nullable"`
	MaxParallelRequests int64                                      `json:"max_parallel_requests,nullable"`
	Members             []interface{}                              `json:"members"`
	MembersWithRoles    []Member                                   `json:"members_with_roles"`
	Metadata            interface{}                                `json:"metadata,nullable"`
	ModelID             int64                                      `json:"model_id,nullable"`
	Models              []interface{}                              `json:"models"`
	OrganizationID      string                                     `json:"organization_id,nullable"`
	RpmLimit            int64                                      `json:"rpm_limit,nullable"`
	Spend               float64                                    `json:"spend,nullable"`
	TeamAlias           string                                     `json:"team_alias,nullable"`
	TpmLimit            int64                                      `json:"tpm_limit,nullable"`
	JSON                organizationListResponseTeamJSON           `json:"-"`
}

// organizationListResponseTeamJSON contains the JSON metadata for the struct
// [OrganizationListResponseTeam]
type organizationListResponseTeamJSON struct {
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

func (r *OrganizationListResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseTeamJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseTeamsLlmModelTable struct {
	CreatedBy    string                                         `json:"created_by,required"`
	UpdatedBy    string                                         `json:"updated_by,required"`
	ModelAliases interface{}                                    `json:"model_aliases,nullable"`
	JSON         organizationListResponseTeamsLlmModelTableJSON `json:"-"`
}

// organizationListResponseTeamsLlmModelTableJSON contains the JSON metadata for
// the struct [OrganizationListResponseTeamsLlmModelTable]
type organizationListResponseTeamsLlmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ModelAliases apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationListResponseTeamsLlmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseTeamsLlmModelTableJSON) RawJSON() string {
	return r.raw
}

// Returned by the /organization/info endpoint and /organization/list endpoint
type OrganizationDeleteResponse struct {
	BudgetID  string    `json:"budget_id,required"`
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	CreatedBy string    `json:"created_by,required"`
	Models    []string  `json:"models,required"`
	UpdatedAt time.Time `json:"updated_at,required" format:"date-time"`
	UpdatedBy string    `json:"updated_by,required"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable    OrganizationDeleteResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Members           []OrganizationDeleteResponseMember       `json:"members"`
	Metadata          interface{}                              `json:"metadata,nullable"`
	OrganizationAlias string                                   `json:"organization_alias,nullable"`
	OrganizationID    string                                   `json:"organization_id,nullable"`
	Spend             float64                                  `json:"spend"`
	Teams             []OrganizationDeleteResponseTeam         `json:"teams"`
	JSON              organizationDeleteResponseJSON           `json:"-"`
}

// organizationDeleteResponseJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponse]
type organizationDeleteResponseJSON struct {
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

func (r *OrganizationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationDeleteResponseLlmBudgetTable struct {
	BudgetDuration      string                                       `json:"budget_duration,nullable"`
	MaxBudget           float64                                      `json:"max_budget,nullable"`
	MaxParallelRequests int64                                        `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                  `json:"model_max_budget,nullable"`
	RpmLimit            int64                                        `json:"rpm_limit,nullable"`
	SoftBudget          float64                                      `json:"soft_budget,nullable"`
	TpmLimit            int64                                        `json:"tpm_limit,nullable"`
	JSON                organizationDeleteResponseLlmBudgetTableJSON `json:"-"`
}

// organizationDeleteResponseLlmBudgetTableJSON contains the JSON metadata for the
// struct [OrganizationDeleteResponseLlmBudgetTable]
type organizationDeleteResponseLlmBudgetTableJSON struct {
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

func (r *OrganizationDeleteResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationDeleteResponseMember struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationDeleteResponseMembersLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                         `json:"spend"`
	User           interface{}                                     `json:"user"`
	UserRole       string                                          `json:"user_role,nullable"`
	JSON           organizationDeleteResponseMemberJSON            `json:"-"`
}

// organizationDeleteResponseMemberJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponseMember]
type organizationDeleteResponseMemberJSON struct {
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

func (r *OrganizationDeleteResponseMember) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseMemberJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationDeleteResponseMembersLlmBudgetTable struct {
	BudgetDuration      string                                              `json:"budget_duration,nullable"`
	MaxBudget           float64                                             `json:"max_budget,nullable"`
	MaxParallelRequests int64                                               `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                         `json:"model_max_budget,nullable"`
	RpmLimit            int64                                               `json:"rpm_limit,nullable"`
	SoftBudget          float64                                             `json:"soft_budget,nullable"`
	TpmLimit            int64                                               `json:"tpm_limit,nullable"`
	JSON                organizationDeleteResponseMembersLlmBudgetTableJSON `json:"-"`
}

// organizationDeleteResponseMembersLlmBudgetTableJSON contains the JSON metadata
// for the struct [OrganizationDeleteResponseMembersLlmBudgetTable]
type organizationDeleteResponseMembersLlmBudgetTableJSON struct {
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

func (r *OrganizationDeleteResponseMembersLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseMembersLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteResponseTeam struct {
	TeamID              string                                       `json:"team_id,required"`
	Admins              []interface{}                                `json:"admins"`
	Blocked             bool                                         `json:"blocked"`
	BudgetDuration      string                                       `json:"budget_duration,nullable"`
	BudgetResetAt       time.Time                                    `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt           time.Time                                    `json:"created_at,nullable" format:"date-time"`
	LlmModelTable       OrganizationDeleteResponseTeamsLlmModelTable `json:"llm_model_table,nullable"`
	MaxBudget           float64                                      `json:"max_budget,nullable"`
	MaxParallelRequests int64                                        `json:"max_parallel_requests,nullable"`
	Members             []interface{}                                `json:"members"`
	MembersWithRoles    []Member                                     `json:"members_with_roles"`
	Metadata            interface{}                                  `json:"metadata,nullable"`
	ModelID             int64                                        `json:"model_id,nullable"`
	Models              []interface{}                                `json:"models"`
	OrganizationID      string                                       `json:"organization_id,nullable"`
	RpmLimit            int64                                        `json:"rpm_limit,nullable"`
	Spend               float64                                      `json:"spend,nullable"`
	TeamAlias           string                                       `json:"team_alias,nullable"`
	TpmLimit            int64                                        `json:"tpm_limit,nullable"`
	JSON                organizationDeleteResponseTeamJSON           `json:"-"`
}

// organizationDeleteResponseTeamJSON contains the JSON metadata for the struct
// [OrganizationDeleteResponseTeam]
type organizationDeleteResponseTeamJSON struct {
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

func (r *OrganizationDeleteResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseTeamJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteResponseTeamsLlmModelTable struct {
	CreatedBy    string                                           `json:"created_by,required"`
	UpdatedBy    string                                           `json:"updated_by,required"`
	ModelAliases interface{}                                      `json:"model_aliases,nullable"`
	JSON         organizationDeleteResponseTeamsLlmModelTableJSON `json:"-"`
}

// organizationDeleteResponseTeamsLlmModelTableJSON contains the JSON metadata for
// the struct [OrganizationDeleteResponseTeamsLlmModelTable]
type organizationDeleteResponseTeamsLlmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ModelAliases apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OrganizationDeleteResponseTeamsLlmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationDeleteResponseTeamsLlmModelTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationAddMemberResponse struct {
	OrganizationID                 string                                                       `json:"organization_id,required"`
	UpdatedOrganizationMemberships []OrganizationAddMemberResponseUpdatedOrganizationMembership `json:"updated_organization_memberships,required"`
	UpdatedUsers                   []OrganizationAddMemberResponseUpdatedUser                   `json:"updated_users,required"`
	JSON                           organizationAddMemberResponseJSON                            `json:"-"`
}

// organizationAddMemberResponseJSON contains the JSON metadata for the struct
// [OrganizationAddMemberResponse]
type organizationAddMemberResponseJSON struct {
	OrganizationID                 apijson.Field
	UpdatedOrganizationMemberships apijson.Field
	UpdatedUsers                   apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *OrganizationAddMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationAddMemberResponseUpdatedOrganizationMembership struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                                                   `json:"spend"`
	User           interface{}                                                               `json:"user"`
	UserRole       string                                                                    `json:"user_role,nullable"`
	JSON           organizationAddMemberResponseUpdatedOrganizationMembershipJSON            `json:"-"`
}

// organizationAddMemberResponseUpdatedOrganizationMembershipJSON contains the JSON
// metadata for the struct
// [OrganizationAddMemberResponseUpdatedOrganizationMembership]
type organizationAddMemberResponseUpdatedOrganizationMembershipJSON struct {
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

func (r *OrganizationAddMemberResponseUpdatedOrganizationMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedOrganizationMembershipJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTable struct {
	BudgetDuration      string                                                                        `json:"budget_duration,nullable"`
	MaxBudget           float64                                                                       `json:"max_budget,nullable"`
	MaxParallelRequests int64                                                                         `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                                                   `json:"model_max_budget,nullable"`
	RpmLimit            int64                                                                         `json:"rpm_limit,nullable"`
	SoftBudget          float64                                                                       `json:"soft_budget,nullable"`
	TpmLimit            int64                                                                         `json:"tpm_limit,nullable"`
	JSON                organizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTableJSON `json:"-"`
}

// organizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTableJSON
// contains the JSON metadata for the struct
// [OrganizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTable]
type organizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTableJSON struct {
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

func (r *OrganizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedOrganizationMembershipsLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationAddMemberResponseUpdatedUser struct {
	UserID                  string                                                            `json:"user_id,required"`
	BudgetDuration          string                                                            `json:"budget_duration,nullable"`
	BudgetResetAt           time.Time                                                         `json:"budget_reset_at,nullable" format:"date-time"`
	MaxBudget               float64                                                           `json:"max_budget,nullable"`
	Metadata                interface{}                                                       `json:"metadata,nullable"`
	ModelMaxBudget          interface{}                                                       `json:"model_max_budget,nullable"`
	ModelSpend              interface{}                                                       `json:"model_spend,nullable"`
	Models                  []interface{}                                                     `json:"models"`
	OrganizationMemberships []OrganizationAddMemberResponseUpdatedUsersOrganizationMembership `json:"organization_memberships,nullable"`
	RpmLimit                int64                                                             `json:"rpm_limit,nullable"`
	Spend                   float64                                                           `json:"spend"`
	SSOUserID               string                                                            `json:"sso_user_id,nullable"`
	Teams                   []string                                                          `json:"teams"`
	TpmLimit                int64                                                             `json:"tpm_limit,nullable"`
	UserEmail               string                                                            `json:"user_email,nullable"`
	UserRole                string                                                            `json:"user_role,nullable"`
	JSON                    organizationAddMemberResponseUpdatedUserJSON                      `json:"-"`
}

// organizationAddMemberResponseUpdatedUserJSON contains the JSON metadata for the
// struct [OrganizationAddMemberResponseUpdatedUser]
type organizationAddMemberResponseUpdatedUserJSON struct {
	UserID                  apijson.Field
	BudgetDuration          apijson.Field
	BudgetResetAt           apijson.Field
	MaxBudget               apijson.Field
	Metadata                apijson.Field
	ModelMaxBudget          apijson.Field
	ModelSpend              apijson.Field
	Models                  apijson.Field
	OrganizationMemberships apijson.Field
	RpmLimit                apijson.Field
	Spend                   apijson.Field
	SSOUserID               apijson.Field
	Teams                   apijson.Field
	TpmLimit                apijson.Field
	UserEmail               apijson.Field
	UserRole                apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrganizationAddMemberResponseUpdatedUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedUserJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationAddMemberResponseUpdatedUsersOrganizationMembership struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                                                        `json:"spend"`
	User           interface{}                                                                    `json:"user"`
	UserRole       string                                                                         `json:"user_role,nullable"`
	JSON           organizationAddMemberResponseUpdatedUsersOrganizationMembershipJSON            `json:"-"`
}

// organizationAddMemberResponseUpdatedUsersOrganizationMembershipJSON contains the
// JSON metadata for the struct
// [OrganizationAddMemberResponseUpdatedUsersOrganizationMembership]
type organizationAddMemberResponseUpdatedUsersOrganizationMembershipJSON struct {
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

func (r *OrganizationAddMemberResponseUpdatedUsersOrganizationMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedUsersOrganizationMembershipJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable struct {
	BudgetDuration      string                                                                             `json:"budget_duration,nullable"`
	MaxBudget           float64                                                                            `json:"max_budget,nullable"`
	MaxParallelRequests int64                                                                              `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                                                        `json:"model_max_budget,nullable"`
	RpmLimit            int64                                                                              `json:"rpm_limit,nullable"`
	SoftBudget          float64                                                                            `json:"soft_budget,nullable"`
	TpmLimit            int64                                                                              `json:"tpm_limit,nullable"`
	JSON                organizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON `json:"-"`
}

// organizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON
// contains the JSON metadata for the struct
// [OrganizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable]
type organizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON struct {
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

func (r *OrganizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationDeleteMemberResponse = interface{}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type OrganizationUpdateMemberResponse struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable OrganizationUpdateMemberResponseLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                        `json:"spend"`
	User           interface{}                                    `json:"user"`
	UserRole       string                                         `json:"user_role,nullable"`
	JSON           organizationUpdateMemberResponseJSON           `json:"-"`
}

// organizationUpdateMemberResponseJSON contains the JSON metadata for the struct
// [OrganizationUpdateMemberResponse]
type organizationUpdateMemberResponseJSON struct {
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

func (r *OrganizationUpdateMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateMemberResponseJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type OrganizationUpdateMemberResponseLlmBudgetTable struct {
	BudgetDuration      string                                             `json:"budget_duration,nullable"`
	MaxBudget           float64                                            `json:"max_budget,nullable"`
	MaxParallelRequests int64                                              `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                        `json:"model_max_budget,nullable"`
	RpmLimit            int64                                              `json:"rpm_limit,nullable"`
	SoftBudget          float64                                            `json:"soft_budget,nullable"`
	TpmLimit            int64                                              `json:"tpm_limit,nullable"`
	JSON                organizationUpdateMemberResponseLlmBudgetTableJSON `json:"-"`
}

// organizationUpdateMemberResponseLlmBudgetTableJSON contains the JSON metadata
// for the struct [OrganizationUpdateMemberResponseLlmBudgetTable]
type organizationUpdateMemberResponseLlmBudgetTableJSON struct {
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

func (r *OrganizationUpdateMemberResponseLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateMemberResponseLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewParams struct {
	OrganizationAlias   param.Field[string]        `json:"organization_alias,required"`
	BudgetDuration      param.Field[string]        `json:"budget_duration"`
	BudgetID            param.Field[string]        `json:"budget_id"`
	MaxBudget           param.Field[float64]       `json:"max_budget"`
	MaxParallelRequests param.Field[int64]         `json:"max_parallel_requests"`
	Metadata            param.Field[interface{}]   `json:"metadata"`
	ModelMaxBudget      param.Field[interface{}]   `json:"model_max_budget"`
	Models              param.Field[[]interface{}] `json:"models"`
	OrganizationID      param.Field[string]        `json:"organization_id"`
	RpmLimit            param.Field[int64]         `json:"rpm_limit"`
	SoftBudget          param.Field[float64]       `json:"soft_budget"`
	TpmLimit            param.Field[int64]         `json:"tpm_limit"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParams struct {
	BudgetID          param.Field[string]      `json:"budget_id"`
	Metadata          param.Field[interface{}] `json:"metadata"`
	Models            param.Field[[]string]    `json:"models"`
	OrganizationAlias param.Field[string]      `json:"organization_alias"`
	OrganizationID    param.Field[string]      `json:"organization_id"`
	Spend             param.Field[float64]     `json:"spend"`
	UpdatedBy         param.Field[string]      `json:"updated_by"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationDeleteParams struct {
	OrganizationIDs param.Field[[]string] `json:"organization_ids,required"`
}

func (r OrganizationDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationAddMemberParams struct {
	Member                  param.Field[OrganizationAddMemberParamsMemberUnion] `json:"member,required"`
	OrganizationID          param.Field[string]                                 `json:"organization_id,required"`
	MaxBudgetInOrganization param.Field[float64]                                `json:"max_budget_in_organization"`
}

func (r OrganizationAddMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [OrganizationAddMemberParamsMemberArray], [OrgMemberParam].
type OrganizationAddMemberParamsMemberUnion interface {
	implementsOrganizationAddMemberParamsMemberUnion()
}

type OrganizationAddMemberParamsMemberArray []OrgMemberParam

func (r OrganizationAddMemberParamsMemberArray) implementsOrganizationAddMemberParamsMemberUnion() {}

type OrganizationDeleteMemberParams struct {
	OrganizationID param.Field[string] `json:"organization_id,required"`
	UserEmail      param.Field[string] `json:"user_email"`
	UserID         param.Field[string] `json:"user_id"`
}

func (r OrganizationDeleteMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateMemberParams struct {
	OrganizationID          param.Field[string]  `json:"organization_id,required"`
	MaxBudgetInOrganization param.Field[float64] `json:"max_budget_in_organization"`
	// Admin Roles: PROXY_ADMIN: admin over the platform PROXY_ADMIN_VIEW_ONLY: can
	// login, view all own keys, view all spend ORG_ADMIN: admin over a specific
	// organization, can create teams, users only within their organization
	//
	// Internal User Roles: INTERNAL_USER: can login, view/create/delete their own
	// keys, view their spend INTERNAL_USER_VIEW_ONLY: can login, view their own keys,
	// view their own spend
	//
	// Team Roles: TEAM: used for JWT auth
	//
	// Customer Roles: CUSTOMER: External users -> these are customers
	Role      param.Field[OrganizationUpdateMemberParamsRole] `json:"role"`
	UserEmail param.Field[string]                             `json:"user_email"`
	UserID    param.Field[string]                             `json:"user_id"`
}

func (r OrganizationUpdateMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Admin Roles: PROXY_ADMIN: admin over the platform PROXY_ADMIN_VIEW_ONLY: can
// login, view all own keys, view all spend ORG_ADMIN: admin over a specific
// organization, can create teams, users only within their organization
//
// Internal User Roles: INTERNAL_USER: can login, view/create/delete their own
// keys, view their spend INTERNAL_USER_VIEW_ONLY: can login, view their own keys,
// view their own spend
//
// Team Roles: TEAM: used for JWT auth
//
// Customer Roles: CUSTOMER: External users -> these are customers
type OrganizationUpdateMemberParamsRole string

const (
	OrganizationUpdateMemberParamsRoleProxyAdmin         OrganizationUpdateMemberParamsRole = "proxy_admin"
	OrganizationUpdateMemberParamsRoleProxyAdminViewer   OrganizationUpdateMemberParamsRole = "proxy_admin_viewer"
	OrganizationUpdateMemberParamsRoleOrgAdmin           OrganizationUpdateMemberParamsRole = "org_admin"
	OrganizationUpdateMemberParamsRoleInternalUser       OrganizationUpdateMemberParamsRole = "internal_user"
	OrganizationUpdateMemberParamsRoleInternalUserViewer OrganizationUpdateMemberParamsRole = "internal_user_viewer"
	OrganizationUpdateMemberParamsRoleTeam               OrganizationUpdateMemberParamsRole = "team"
	OrganizationUpdateMemberParamsRoleCustomer           OrganizationUpdateMemberParamsRole = "customer"
)

func (r OrganizationUpdateMemberParamsRole) IsKnown() bool {
	switch r {
	case OrganizationUpdateMemberParamsRoleProxyAdmin, OrganizationUpdateMemberParamsRoleProxyAdminViewer, OrganizationUpdateMemberParamsRoleOrgAdmin, OrganizationUpdateMemberParamsRoleInternalUser, OrganizationUpdateMemberParamsRoleInternalUserViewer, OrganizationUpdateMemberParamsRoleTeam, OrganizationUpdateMemberParamsRoleCustomer:
		return true
	}
	return false
}
