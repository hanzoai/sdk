// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// TeamService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options  []option.RequestOption
	Model    *TeamModelService
	Callback *TeamCallbackService
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r *TeamService) {
	r = &TeamService{}
	r.Options = opts
	r.Model = NewTeamModelService(opts...)
	r.Callback = NewTeamCallbackService(opts...)
	return
}

// Allow users to create a new team. Apply user permissions to their team.
//
// 👉
// [Detailed Doc on setting team budgets](https://docs.hanzo.ai/docs/proxy/team_budgets)
//
// Parameters:
//
//   - team_alias: Optional[str] - User defined team alias
//   - team_id: Optional[str] - The team id of the user. If none passed, we'll
//     generate it.
//   - members_with_roles: List[{"role": "admin" or "user", "user_id":
//     "<user-id>"}] - A list of users and their roles in the team. Get user_id when
//     making a new user via `/user/new`.
//   - metadata: Optional[dict] - Metadata for team, store information for team.
//     Example metadata = {"extra_info": "some info"}
//   - tpm_limit: Optional[int] - The TPM (Tokens Per Minute) limit for this team -
//     all keys with this team_id will have at max this TPM limit
//   - rpm_limit: Optional[int] - The RPM (Requests Per Minute) limit for this team -
//     all keys associated with this team_id will have at max this RPM limit
//   - max_budget: Optional[float] - The maximum budget allocated to the team - all
//     keys for this team_id will have at max this max_budget
//   - budget_duration: Optional[str] - The duration of the budget for the team. Doc
//     [here](https://docs.hanzo.ai/docs/proxy/team_budgets)
//   - models: Optional[list] - A list of models associated with the team - all keys
//     for this team_id will have at most, these models. If empty, assumes all models
//     are allowed.
//   - blocked: bool - Flag indicating if the team is blocked or not - will stop all
//     calls from keys with this team_id.
//   - members: Optional[List] - Control team members via `/team/member/add` and
//     `/team/member/delete`.
//   - tags: Optional[List[str]] - Tags for
//     [tracking spend](https://llm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://llm.vercel.app/docs/proxy/tag_routing).
//   - organization_id: Optional[str] - The organization id of the team. Default is
//     None. Create via `/organization/new`.
//   - model_aliases: Optional[dict] - Model aliases for the team.
//     [Docs](https://docs.hanzo.ai/docs/proxy/team_based_routing#create-team-with-model-alias)
//   - guardrails: Optional[List[str]] - Guardrails for the team.
//     [Docs](https://docs.hanzo.ai/docs/proxy/guardrails) Returns:
//   - team_id: (str) Unique team id - used for tracking spend across multiple keys
//     for same team id.
//
// \_deprecated_params:
//
// - admins: list - A list of user_id's for the admin role
// - users: list - A list of user_id's for the user role
//
// Example Request:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/new'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	  "team_alias": "my-new-team_2",
//	  "members_with_roles": [{"role": "admin", "user_id": "user-1234"},
//	    {"role": "user", "user_id": "user-2434"}]
//	}'
//
// ```
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/new'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	           "team_alias": "QA Prod Bot",
//	           "max_budget": 0.000000001,
//	           "budget_duration": "1d"
//	       }'
//
// ```
func (r *TeamService) New(ctx context.Context, params TeamNewParams, opts ...option.RequestOption) (res *TeamNewResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "team/new"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Use `/team/member_add` AND `/team/member/delete` to add/remove new team members
//
// You can now update team budget / rate limits via /team/update
//
// Parameters:
//
//   - team_id: str - The team id of the user. Required param.
//   - team_alias: Optional[str] - User defined team alias
//   - metadata: Optional[dict] - Metadata for team, store information for team.
//     Example metadata = {"team": "core-infra", "app": "app2", "email": "z@hanzo.ai"
//     }
//   - tpm_limit: Optional[int] - The TPM (Tokens Per Minute) limit for this team -
//     all keys with this team_id will have at max this TPM limit
//   - rpm_limit: Optional[int] - The RPM (Requests Per Minute) limit for this team -
//     all keys associated with this team_id will have at max this RPM limit
//   - max_budget: Optional[float] - The maximum budget allocated to the team - all
//     keys for this team_id will have at max this max_budget
//   - budget_duration: Optional[str] - The duration of the budget for the team. Doc
//     [here](https://docs.hanzo.ai/docs/proxy/team_budgets)
//   - models: Optional[list] - A list of models associated with the team - all keys
//     for this team_id will have at most, these models. If empty, assumes all models
//     are allowed.
//   - blocked: bool - Flag indicating if the team is blocked or not - will stop all
//     calls from keys with this team_id.
//   - tags: Optional[List[str]] - Tags for
//     [tracking spend](https://llm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://llm.vercel.app/docs/proxy/tag_routing).
//   - organization_id: Optional[str] - The organization id of the team. Default is
//     None. Create via `/organization/new`.
//   - model_aliases: Optional[dict] - Model aliases for the team.
//     [Docs](https://docs.hanzo.ai/docs/proxy/team_based_routing#create-team-with-model-alias)
//   - guardrails: Optional[List[str]] - Guardrails for the team.
//     [Docs](https://docs.hanzo.ai/docs/proxy/guardrails) Example - update team TPM
//     Limit
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "team_id": "8d916b1c-510d-4894-a334-1c16a93344f5",
//	    "tpm_limit": 100
//	}'
//
// ```
//
// Example - Update Team `max_budget` budget
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "team_id": "8d916b1c-510d-4894-a334-1c16a93344f5",
//	    "max_budget": 10
//	}'
//
// ```
func (r *TeamService) Update(ctx context.Context, params TeamUpdateParams, opts ...option.RequestOption) (res *TeamUpdateResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "team/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// ```
// curl --location --request GET 'http://0.0.0.0:4000/team/list'         --header 'Authorization: Bearer sk-1234'
// ```
//
// Parameters:
//
//   - user_id: str - Optional. If passed will only return teams that the user_id is
//     a member of.
//   - organization_id: str - Optional. If passed will only return teams that belong
//     to the organization_id. Pass 'default_organization' to get all teams without
//     organization_id.
func (r *TeamService) List(ctx context.Context, query TeamListParams, opts ...option.RequestOption) (res *TeamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// delete team and associated team keys
//
// Parameters:
//
//   - team_ids: List[str] - Required. List of team IDs to delete. Example:
//     ["team-1234", "team-5678"]
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/delete'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "team_ids": ["8d916b1c-510d-4894-a334-1c16a93344f5"]
//	}'
//
// ```
func (r *TeamService) Delete(ctx context.Context, params TeamDeleteParams, opts ...option.RequestOption) (res *TeamDeleteResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "team/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// [BETA]
//
// Add new members (either via user_email or user_id) to a team
//
// If user doesn't exist, new user row will also be added to User Table
//
// Only proxy_admin or admin of team, allowed to access this endpoint.
//
// ```
//
// curl -X POST 'http://0.0.0.0:4000/team/member_add'     -H 'Authorization: Bearer sk-1234'     -H 'Content-Type: application/json'     -d '{"team_id": "45e3e396-ee08-4a61-a88e-16b3ce7e0849", "member": {"role": "user", "user_id": "dev247652@hanzo.ai"}}'
//
// ```
func (r *TeamService) AddMember(ctx context.Context, body TeamAddMemberParams, opts ...option.RequestOption) (res *TeamAddMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/member_add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Blocks all calls from keys with this team id.
//
// Parameters:
//
// - team_id: str - Required. The unique identifier of the team to block.
//
// Example:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/block'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "team_id": "team-1234"
//	}'
//
// ```
//
// Returns:
//
// - The updated team record with blocked=True
func (r *TeamService) Block(ctx context.Context, body TeamBlockParams, opts ...option.RequestOption) (res *TeamBlockResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/block"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disable all logging callbacks for a team
//
// Parameters:
//
// - team_id (str, required): The unique identifier for the team
//
// Example curl:
//
// ```
// curl -X POST 'http://localhost:4000/team/dbe2f686-a686-4896-864a-4c3924458709/disable_logging'         -H 'Authorization: Bearer sk-1234'
// ```
func (r *TeamService) DisableLogging(ctx context.Context, teamID string, opts ...option.RequestOption) (res *TeamDisableLoggingResponse, err error) {
	opts = append(r.Options[:], opts...)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	path := fmt.Sprintf("team/%s/disable_logging", teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// List Available Teams
func (r *TeamService) ListAvailable(ctx context.Context, query TeamListAvailableParams, opts ...option.RequestOption) (res *TeamListAvailableResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// [BETA]
//
// delete members (either via user_email or user_id) from a team
//
// If user doesn't exist, an exception will be raised
//
// ```
// curl -X POST 'http://0.0.0.0:8000/team/member_delete'
// -H 'Authorization: Bearer sk-1234'
// -H 'Content-Type: application/json'
//
//	-d '{
//	    "team_id": "45e3e396-ee08-4a61-a88e-16b3ce7e0849",
//	    "user_id": "dev247652@hanzo.ai"
//	}'
//
// ```
func (r *TeamService) RemoveMember(ctx context.Context, body TeamRemoveMemberParams, opts ...option.RequestOption) (res *TeamRemoveMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/member_delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// get info on team + related keys
//
// Parameters:
//
// - team_id: str - Required. The unique identifier of the team to get info on.
//
// ```
// curl --location 'http://localhost:4000/team/info?team_id=your_team_id_here'     --header 'Authorization: Bearer your_api_key_here'
// ```
func (r *TeamService) GetInfo(ctx context.Context, query TeamGetInfoParams, opts ...option.RequestOption) (res *TeamGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Blocks all calls from keys with this team id.
//
// Parameters:
//
// - team_id: str - Required. The unique identifier of the team to unblock.
//
// Example:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/unblock'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "team_id": "team-1234"
//	}'
//
// ```
func (r *TeamService) Unblock(ctx context.Context, body TeamUnblockParams, opts ...option.RequestOption) (res *TeamUnblockResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/unblock"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [BETA]
//
// Update team member budgets and team member role
func (r *TeamService) UpdateMember(ctx context.Context, body TeamUpdateMemberParams, opts ...option.RequestOption) (res *TeamUpdateMemberResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/member_update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type BlockTeamRequestParam struct {
	TeamID param.Field[string] `json:"team_id,required"`
}

func (r BlockTeamRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Member struct {
	Role      MemberRole `json:"role,required"`
	UserEmail string     `json:"user_email,nullable"`
	UserID    string     `json:"user_id,nullable"`
	JSON      memberJSON `json:"-"`
}

// memberJSON contains the JSON metadata for the struct [Member]
type memberJSON struct {
	Role        apijson.Field
	UserEmail   apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Member) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberJSON) RawJSON() string {
	return r.raw
}

type MemberRole string

const (
	MemberRoleAdmin MemberRole = "admin"
	MemberRoleUser  MemberRole = "user"
)

func (r MemberRole) IsKnown() bool {
	switch r {
	case MemberRoleAdmin, MemberRoleUser:
		return true
	}
	return false
}

type MemberParam struct {
	Role      param.Field[MemberRole] `json:"role,required"`
	UserEmail param.Field[string]     `json:"user_email"`
	UserID    param.Field[string]     `json:"user_id"`
}

func (r MemberParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberParam) implementsTeamAddMemberParamsMemberUnion() {}

type TeamNewResponse struct {
	TeamID              string                       `json:"team_id,required"`
	Admins              []interface{}                `json:"admins"`
	Blocked             bool                         `json:"blocked"`
	BudgetDuration      string                       `json:"budget_duration,nullable"`
	BudgetResetAt       time.Time                    `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt           time.Time                    `json:"created_at,nullable" format:"date-time"`
	LlmModelTable       TeamNewResponseLlmModelTable `json:"llm_model_table,nullable"`
	MaxBudget           float64                      `json:"max_budget,nullable"`
	MaxParallelRequests int64                        `json:"max_parallel_requests,nullable"`
	Members             []interface{}                `json:"members"`
	MembersWithRoles    []Member                     `json:"members_with_roles"`
	Metadata            interface{}                  `json:"metadata,nullable"`
	ModelID             int64                        `json:"model_id,nullable"`
	Models              []interface{}                `json:"models"`
	OrganizationID      string                       `json:"organization_id,nullable"`
	RpmLimit            int64                        `json:"rpm_limit,nullable"`
	Spend               float64                      `json:"spend,nullable"`
	TeamAlias           string                       `json:"team_alias,nullable"`
	TpmLimit            int64                        `json:"tpm_limit,nullable"`
	JSON                teamNewResponseJSON          `json:"-"`
}

// teamNewResponseJSON contains the JSON metadata for the struct [TeamNewResponse]
type teamNewResponseJSON struct {
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

func (r *TeamNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamNewResponseJSON) RawJSON() string {
	return r.raw
}

type TeamNewResponseLlmModelTable struct {
	CreatedBy    string                           `json:"created_by,required"`
	UpdatedBy    string                           `json:"updated_by,required"`
	ModelAliases interface{}                      `json:"model_aliases,nullable"`
	JSON         teamNewResponseLlmModelTableJSON `json:"-"`
}

// teamNewResponseLlmModelTableJSON contains the JSON metadata for the struct
// [TeamNewResponseLlmModelTable]
type teamNewResponseLlmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ModelAliases apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamNewResponseLlmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamNewResponseLlmModelTableJSON) RawJSON() string {
	return r.raw
}

type TeamUpdateResponse = interface{}

type TeamListResponse = interface{}

type TeamDeleteResponse = interface{}

type TeamAddMemberResponse struct {
	TeamID                 string                                       `json:"team_id,required"`
	UpdatedTeamMemberships []TeamAddMemberResponseUpdatedTeamMembership `json:"updated_team_memberships,required"`
	UpdatedUsers           []TeamAddMemberResponseUpdatedUser           `json:"updated_users,required"`
	Admins                 []interface{}                                `json:"admins"`
	Blocked                bool                                         `json:"blocked"`
	BudgetDuration         string                                       `json:"budget_duration,nullable"`
	BudgetResetAt          time.Time                                    `json:"budget_reset_at,nullable" format:"date-time"`
	CreatedAt              time.Time                                    `json:"created_at,nullable" format:"date-time"`
	LlmModelTable          TeamAddMemberResponseLlmModelTable           `json:"llm_model_table,nullable"`
	MaxBudget              float64                                      `json:"max_budget,nullable"`
	MaxParallelRequests    int64                                        `json:"max_parallel_requests,nullable"`
	Members                []interface{}                                `json:"members"`
	MembersWithRoles       []Member                                     `json:"members_with_roles"`
	Metadata               interface{}                                  `json:"metadata,nullable"`
	ModelID                int64                                        `json:"model_id,nullable"`
	Models                 []interface{}                                `json:"models"`
	OrganizationID         string                                       `json:"organization_id,nullable"`
	RpmLimit               int64                                        `json:"rpm_limit,nullable"`
	Spend                  float64                                      `json:"spend,nullable"`
	TeamAlias              string                                       `json:"team_alias,nullable"`
	TpmLimit               int64                                        `json:"tpm_limit,nullable"`
	JSON                   teamAddMemberResponseJSON                    `json:"-"`
}

// teamAddMemberResponseJSON contains the JSON metadata for the struct
// [TeamAddMemberResponse]
type teamAddMemberResponseJSON struct {
	TeamID                 apijson.Field
	UpdatedTeamMemberships apijson.Field
	UpdatedUsers           apijson.Field
	Admins                 apijson.Field
	Blocked                apijson.Field
	BudgetDuration         apijson.Field
	BudgetResetAt          apijson.Field
	CreatedAt              apijson.Field
	LlmModelTable          apijson.Field
	MaxBudget              apijson.Field
	MaxParallelRequests    apijson.Field
	Members                apijson.Field
	MembersWithRoles       apijson.Field
	Metadata               apijson.Field
	ModelID                apijson.Field
	Models                 apijson.Field
	OrganizationID         apijson.Field
	RpmLimit               apijson.Field
	Spend                  apijson.Field
	TeamAlias              apijson.Field
	TpmLimit               apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TeamAddMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseJSON) RawJSON() string {
	return r.raw
}

type TeamAddMemberResponseUpdatedTeamMembership struct {
	BudgetID string `json:"budget_id,required"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable TeamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTable `json:"llm_budget_table,required,nullable"`
	TeamID         string                                                    `json:"team_id,required"`
	UserID         string                                                    `json:"user_id,required"`
	JSON           teamAddMemberResponseUpdatedTeamMembershipJSON            `json:"-"`
}

// teamAddMemberResponseUpdatedTeamMembershipJSON contains the JSON metadata for
// the struct [TeamAddMemberResponseUpdatedTeamMembership]
type teamAddMemberResponseUpdatedTeamMembershipJSON struct {
	BudgetID       apijson.Field
	LlmBudgetTable apijson.Field
	TeamID         apijson.Field
	UserID         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TeamAddMemberResponseUpdatedTeamMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedTeamMembershipJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type TeamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTable struct {
	BudgetDuration      string                                                        `json:"budget_duration,nullable"`
	MaxBudget           float64                                                       `json:"max_budget,nullable"`
	MaxParallelRequests int64                                                         `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                                   `json:"model_max_budget,nullable"`
	RpmLimit            int64                                                         `json:"rpm_limit,nullable"`
	SoftBudget          float64                                                       `json:"soft_budget,nullable"`
	TpmLimit            int64                                                         `json:"tpm_limit,nullable"`
	JSON                teamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTableJSON `json:"-"`
}

// teamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTableJSON contains the JSON
// metadata for the struct
// [TeamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTable]
type teamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTableJSON struct {
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

func (r *TeamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedTeamMembershipsLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type TeamAddMemberResponseUpdatedUser struct {
	UserID                  string                                                    `json:"user_id,required"`
	BudgetDuration          string                                                    `json:"budget_duration,nullable"`
	BudgetResetAt           time.Time                                                 `json:"budget_reset_at,nullable" format:"date-time"`
	MaxBudget               float64                                                   `json:"max_budget,nullable"`
	Metadata                interface{}                                               `json:"metadata,nullable"`
	ModelMaxBudget          interface{}                                               `json:"model_max_budget,nullable"`
	ModelSpend              interface{}                                               `json:"model_spend,nullable"`
	Models                  []interface{}                                             `json:"models"`
	OrganizationMemberships []TeamAddMemberResponseUpdatedUsersOrganizationMembership `json:"organization_memberships,nullable"`
	RpmLimit                int64                                                     `json:"rpm_limit,nullable"`
	Spend                   float64                                                   `json:"spend"`
	SSOUserID               string                                                    `json:"sso_user_id,nullable"`
	Teams                   []string                                                  `json:"teams"`
	TpmLimit                int64                                                     `json:"tpm_limit,nullable"`
	UserEmail               string                                                    `json:"user_email,nullable"`
	UserRole                string                                                    `json:"user_role,nullable"`
	JSON                    teamAddMemberResponseUpdatedUserJSON                      `json:"-"`
}

// teamAddMemberResponseUpdatedUserJSON contains the JSON metadata for the struct
// [TeamAddMemberResponseUpdatedUser]
type teamAddMemberResponseUpdatedUserJSON struct {
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

func (r *TeamAddMemberResponseUpdatedUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedUserJSON) RawJSON() string {
	return r.raw
}

// This is the table that track what organizations a user belongs to and users
// spend within the organization
type TeamAddMemberResponseUpdatedUsersOrganizationMembership struct {
	CreatedAt      time.Time `json:"created_at,required" format:"date-time"`
	OrganizationID string    `json:"organization_id,required"`
	UpdatedAt      time.Time `json:"updated_at,required" format:"date-time"`
	UserID         string    `json:"user_id,required"`
	BudgetID       string    `json:"budget_id,nullable"`
	// Represents user-controllable params for a LLM_BudgetTable record
	LlmBudgetTable TeamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable `json:"llm_budget_table,nullable"`
	Spend          float64                                                                `json:"spend"`
	User           interface{}                                                            `json:"user"`
	UserRole       string                                                                 `json:"user_role,nullable"`
	JSON           teamAddMemberResponseUpdatedUsersOrganizationMembershipJSON            `json:"-"`
}

// teamAddMemberResponseUpdatedUsersOrganizationMembershipJSON contains the JSON
// metadata for the struct
// [TeamAddMemberResponseUpdatedUsersOrganizationMembership]
type teamAddMemberResponseUpdatedUsersOrganizationMembershipJSON struct {
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

func (r *TeamAddMemberResponseUpdatedUsersOrganizationMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedUsersOrganizationMembershipJSON) RawJSON() string {
	return r.raw
}

// Represents user-controllable params for a LLM_BudgetTable record
type TeamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable struct {
	BudgetDuration      string                                                                     `json:"budget_duration,nullable"`
	MaxBudget           float64                                                                    `json:"max_budget,nullable"`
	MaxParallelRequests int64                                                                      `json:"max_parallel_requests,nullable"`
	ModelMaxBudget      interface{}                                                                `json:"model_max_budget,nullable"`
	RpmLimit            int64                                                                      `json:"rpm_limit,nullable"`
	SoftBudget          float64                                                                    `json:"soft_budget,nullable"`
	TpmLimit            int64                                                                      `json:"tpm_limit,nullable"`
	JSON                teamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON `json:"-"`
}

// teamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON
// contains the JSON metadata for the struct
// [TeamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable]
type teamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON struct {
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

func (r *TeamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseUpdatedUsersOrganizationMembershipsLlmBudgetTableJSON) RawJSON() string {
	return r.raw
}

type TeamAddMemberResponseLlmModelTable struct {
	CreatedBy    string                                 `json:"created_by,required"`
	UpdatedBy    string                                 `json:"updated_by,required"`
	ModelAliases interface{}                            `json:"model_aliases,nullable"`
	JSON         teamAddMemberResponseLlmModelTableJSON `json:"-"`
}

// teamAddMemberResponseLlmModelTableJSON contains the JSON metadata for the struct
// [TeamAddMemberResponseLlmModelTable]
type teamAddMemberResponseLlmModelTableJSON struct {
	CreatedBy    apijson.Field
	UpdatedBy    apijson.Field
	ModelAliases apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamAddMemberResponseLlmModelTable) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamAddMemberResponseLlmModelTableJSON) RawJSON() string {
	return r.raw
}

type TeamBlockResponse = interface{}

type TeamDisableLoggingResponse = interface{}

type TeamListAvailableResponse = interface{}

type TeamRemoveMemberResponse = interface{}

type TeamGetInfoResponse = interface{}

type TeamUnblockResponse = interface{}

type TeamUpdateMemberResponse struct {
	TeamID          string                       `json:"team_id,required"`
	UserID          string                       `json:"user_id,required"`
	MaxBudgetInTeam float64                      `json:"max_budget_in_team,nullable"`
	UserEmail       string                       `json:"user_email,nullable"`
	JSON            teamUpdateMemberResponseJSON `json:"-"`
}

// teamUpdateMemberResponseJSON contains the JSON metadata for the struct
// [TeamUpdateMemberResponse]
type teamUpdateMemberResponseJSON struct {
	TeamID          apijson.Field
	UserID          apijson.Field
	MaxBudgetInTeam apijson.Field
	UserEmail       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TeamUpdateMemberResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamUpdateMemberResponseJSON) RawJSON() string {
	return r.raw
}

type TeamNewParams struct {
	Admins           param.Field[[]interface{}] `json:"admins"`
	Blocked          param.Field[bool]          `json:"blocked"`
	BudgetDuration   param.Field[string]        `json:"budget_duration"`
	Guardrails       param.Field[[]string]      `json:"guardrails"`
	MaxBudget        param.Field[float64]       `json:"max_budget"`
	Members          param.Field[[]interface{}] `json:"members"`
	MembersWithRoles param.Field[[]MemberParam] `json:"members_with_roles"`
	Metadata         param.Field[interface{}]   `json:"metadata"`
	ModelAliases     param.Field[interface{}]   `json:"model_aliases"`
	Models           param.Field[[]interface{}] `json:"models"`
	OrganizationID   param.Field[string]        `json:"organization_id"`
	RpmLimit         param.Field[int64]         `json:"rpm_limit"`
	Tags             param.Field[[]interface{}] `json:"tags"`
	TeamAlias        param.Field[string]        `json:"team_alias"`
	TeamID           param.Field[string]        `json:"team_id"`
	TpmLimit         param.Field[int64]         `json:"tpm_limit"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r TeamNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamUpdateParams struct {
	TeamID         param.Field[string]        `json:"team_id,required"`
	Blocked        param.Field[bool]          `json:"blocked"`
	BudgetDuration param.Field[string]        `json:"budget_duration"`
	Guardrails     param.Field[[]string]      `json:"guardrails"`
	MaxBudget      param.Field[float64]       `json:"max_budget"`
	Metadata       param.Field[interface{}]   `json:"metadata"`
	ModelAliases   param.Field[interface{}]   `json:"model_aliases"`
	Models         param.Field[[]interface{}] `json:"models"`
	OrganizationID param.Field[string]        `json:"organization_id"`
	RpmLimit       param.Field[int64]         `json:"rpm_limit"`
	Tags           param.Field[[]interface{}] `json:"tags"`
	TeamAlias      param.Field[string]        `json:"team_alias"`
	TpmLimit       param.Field[int64]         `json:"tpm_limit"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r TeamUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamListParams struct {
	OrganizationID param.Field[string] `query:"organization_id"`
	// Only return teams which this 'user_id' belongs to
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [TeamListParams]'s query parameters as `url.Values`.
func (r TeamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamDeleteParams struct {
	TeamIDs param.Field[[]string] `json:"team_ids,required"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r TeamDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamAddMemberParams struct {
	Member          param.Field[TeamAddMemberParamsMemberUnion] `json:"member,required"`
	TeamID          param.Field[string]                         `json:"team_id,required"`
	MaxBudgetInTeam param.Field[float64]                        `json:"max_budget_in_team"`
}

func (r TeamAddMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [TeamAddMemberParamsMemberArray], [MemberParam].
type TeamAddMemberParamsMemberUnion interface {
	implementsTeamAddMemberParamsMemberUnion()
}

type TeamAddMemberParamsMemberArray []MemberParam

func (r TeamAddMemberParamsMemberArray) implementsTeamAddMemberParamsMemberUnion() {}

type TeamBlockParams struct {
	BlockTeamRequest BlockTeamRequestParam `json:"block_team_request,required"`
}

func (r TeamBlockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockTeamRequest)
}

type TeamListAvailableParams struct {
	ResponseModel param.Field[interface{}] `query:"response_model"`
}

// URLQuery serializes [TeamListAvailableParams]'s query parameters as
// `url.Values`.
func (r TeamListAvailableParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamRemoveMemberParams struct {
	TeamID    param.Field[string] `json:"team_id,required"`
	UserEmail param.Field[string] `json:"user_email"`
	UserID    param.Field[string] `json:"user_id"`
}

func (r TeamRemoveMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamGetInfoParams struct {
	// Team ID in the request parameters
	TeamID param.Field[string] `query:"team_id"`
}

// URLQuery serializes [TeamGetInfoParams]'s query parameters as `url.Values`.
func (r TeamGetInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamUnblockParams struct {
	BlockTeamRequest BlockTeamRequestParam `json:"block_team_request,required"`
}

func (r TeamUnblockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockTeamRequest)
}

type TeamUpdateMemberParams struct {
	TeamID          param.Field[string]                     `json:"team_id,required"`
	MaxBudgetInTeam param.Field[float64]                    `json:"max_budget_in_team"`
	Role            param.Field[TeamUpdateMemberParamsRole] `json:"role"`
	UserEmail       param.Field[string]                     `json:"user_email"`
	UserID          param.Field[string]                     `json:"user_id"`
}

func (r TeamUpdateMemberParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamUpdateMemberParamsRole string

const (
	TeamUpdateMemberParamsRoleAdmin TeamUpdateMemberParamsRole = "admin"
	TeamUpdateMemberParamsRoleUser  TeamUpdateMemberParamsRole = "user"
)

func (r TeamUpdateMemberParamsRole) IsKnown() bool {
	switch r {
	case TeamUpdateMemberParamsRoleAdmin, TeamUpdateMemberParamsRoleUser:
		return true
	}
	return false
}
