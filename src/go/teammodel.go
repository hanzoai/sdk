// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// TeamModelService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamModelService] method instead.
type TeamModelService struct {
	Options []option.RequestOption
}

// NewTeamModelService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamModelService(opts ...option.RequestOption) (r *TeamModelService) {
	r = &TeamModelService{}
	r.Options = opts
	return
}

// Add models to a team's allowed model list. Only proxy admin or team admin can
// add models.
//
// Parameters:
//
// - team_id: str - Required. The team to add models to
// - models: List[str] - Required. List of models to add to the team
//
// Example Request:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/model/add'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "team_id": "team-1234",
//	    "models": ["gpt-4", "claude-2"]
//	}'
//
// ```
func (r *TeamModelService) Add(ctx context.Context, body TeamModelAddParams, opts ...option.RequestOption) (res *TeamModelAddResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/model/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Remove models from a team's allowed model list. Only proxy admin or team admin
// can remove models.
//
// Parameters:
//
// - team_id: str - Required. The team to remove models from
// - models: List[str] - Required. List of models to remove from the team
//
// Example Request:
//
// ```
//
//	curl --location 'http://0.0.0.0:4000/team/model/delete'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "team_id": "team-1234",
//	    "models": ["gpt-4"]
//	}'
//
// ```
func (r *TeamModelService) Remove(ctx context.Context, body TeamModelRemoveParams, opts ...option.RequestOption) (res *TeamModelRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "team/model/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type TeamModelAddResponse = interface{}

type TeamModelRemoveResponse = interface{}

type TeamModelAddParams struct {
	Models param.Field[[]string] `json:"models,required"`
	TeamID param.Field[string]   `json:"team_id,required"`
}

func (r TeamModelAddParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TeamModelRemoveParams struct {
	Models param.Field[[]string] `json:"models,required"`
	TeamID param.Field[string]   `json:"team_id,required"`
}

func (r TeamModelRemoveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
