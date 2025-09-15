// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// AssistantService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAssistantService] method instead.
type AssistantService struct {
	Options []option.RequestOption
}

// NewAssistantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAssistantService(opts ...option.RequestOption) (r *AssistantService) {
	r = &AssistantService{}
	r.Options = opts
	return
}

// Create assistant
//
// API Reference docs -
// https://platform.openai.com/docs/api-reference/assistants/createAssistant
func (r *AssistantService) New(ctx context.Context, opts ...option.RequestOption) (res *AssistantNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Returns a list of assistants.
//
// API Reference docs -
// https://platform.openai.com/docs/api-reference/assistants/listAssistants
func (r *AssistantService) List(ctx context.Context, opts ...option.RequestOption) (res *AssistantListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/assistants"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete assistant
//
// API Reference docs -
// https://platform.openai.com/docs/api-reference/assistants/createAssistant
func (r *AssistantService) Delete(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *AssistantDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return
	}
	path := fmt.Sprintf("v1/assistants/%s", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AssistantNewResponse = interface{}

type AssistantListResponse = interface{}

type AssistantDeleteResponse = interface{}
