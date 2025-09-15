// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// TeamCallbackService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamCallbackService] method instead.
type TeamCallbackService struct {
	Options []option.RequestOption
}

// NewTeamCallbackService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamCallbackService(opts ...option.RequestOption) (r *TeamCallbackService) {
	r = &TeamCallbackService{}
	r.Options = opts
	return
}

// Get the success/failure callbacks and variables for a team
//
// Parameters:
//
// - team_id (str, required): The unique identifier for the team
//
// Example curl:
//
// ```
// curl -X GET 'http://localhost:4000/team/dbe2f686-a686-4896-864a-4c3924458709/callback'         -H 'Authorization: Bearer sk-1234'
// ```
//
// This will return the callback settings for the team with id
// dbe2f686-a686-4896-864a-4c3924458709
//
// Returns { "status": "success", "data": { "team_id": team_id,
// "success_callbacks": team_callback_settings_obj.success_callback,
// "failure_callbacks": team_callback_settings_obj.failure_callback,
// "callback_vars": team_callback_settings_obj.callback_vars, }, }
func (r *TeamCallbackService) Get(ctx context.Context, teamID string, opts ...option.RequestOption) (res *TeamCallbackGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	path := fmt.Sprintf("team/%s/callback", teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add a success/failure callback to a team
//
// Use this if if you want different teams to have different success/failure
// callbacks
//
// Parameters:
//
//   - callback_name (Literal["langfuse", "langsmith", "gcs"], required): The name of
//     the callback to add
//   - callback_type (Literal["success", "failure", "success_and_failure"],
//     required): The type of callback to add. One of:
//   - "success": Callback for successful LLM calls
//   - "failure": Callback for failed LLM calls
//   - "success_and_failure": Callback for both successful and failed LLM calls
//   - callback_vars (StandardCallbackDynamicParams, required): A dictionary of
//     variables to pass to the callback
//   - langfuse_public_key: The public key for the Langfuse callback
//   - langfuse_secret_key: The secret key for the Langfuse callback
//   - langfuse_secret: The secret for the Langfuse callback
//   - langfuse_host: The host for the Langfuse callback
//   - gcs_bucket_name: The name of the GCS bucket
//   - gcs_path_service_account: The path to the GCS service account
//   - langsmith_api_key: The API key for the Langsmith callback
//   - langsmith_project: The project for the Langsmith callback
//   - langsmith_base_url: The base URL for the Langsmith callback
//
// Example curl:
//
// ```
//
//	curl -X POST 'http:/localhost:4000/team/dbe2f686-a686-4896-864a-4c3924458709/callback'         -H 'Content-Type: application/json'         -H 'Authorization: Bearer sk-1234'         -d '{
//	    "callback_name": "langfuse",
//	    "callback_type": "success",
//	    "callback_vars": {"langfuse_public_key": "pk-lf-xxxx1", "langfuse_secret_key": "sk-xxxxx"}
//
// }'
// ```
//
// This means for the team where team_id = dbe2f686-a686-4896-864a-4c3924458709,
// all LLM calls will be logged to langfuse using the public key pk-lf-xxxx1 and
// the secret key sk-xxxxx
func (r *TeamCallbackService) Add(ctx context.Context, teamID string, params TeamCallbackAddParams, opts ...option.RequestOption) (res *TeamCallbackAddResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	if teamID == "" {
		err = errors.New("missing required team_id parameter")
		return
	}
	path := fmt.Sprintf("team/%s/callback", teamID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TeamCallbackGetResponse = interface{}

type TeamCallbackAddResponse = interface{}

type TeamCallbackAddParams struct {
	CallbackName param.Field[string]                            `json:"callback_name,required"`
	CallbackVars param.Field[map[string]string]                 `json:"callback_vars,required"`
	CallbackType param.Field[TeamCallbackAddParamsCallbackType] `json:"callback_type"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r TeamCallbackAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamCallbackAddParamsCallbackType string

const (
	TeamCallbackAddParamsCallbackTypeSuccess           TeamCallbackAddParamsCallbackType = "success"
	TeamCallbackAddParamsCallbackTypeFailure           TeamCallbackAddParamsCallbackType = "failure"
	TeamCallbackAddParamsCallbackTypeSuccessAndFailure TeamCallbackAddParamsCallbackType = "success_and_failure"
)

func (r TeamCallbackAddParamsCallbackType) IsKnown() bool {
	switch r {
	case TeamCallbackAddParamsCallbackTypeSuccess, TeamCallbackAddParamsCallbackTypeFailure, TeamCallbackAddParamsCallbackTypeSuccessAndFailure:
		return true
	}
	return false
}
