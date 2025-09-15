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

// OpenAIDeploymentChatService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIDeploymentChatService] method instead.
type OpenAIDeploymentChatService struct {
	Options []option.RequestOption
}

// NewOpenAIDeploymentChatService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIDeploymentChatService(opts ...option.RequestOption) (r *OpenAIDeploymentChatService) {
	r = &OpenAIDeploymentChatService{}
	r.Options = opts
	return
}

// Follows the exact same API spec as
// `OpenAI's Chat API https://platform.openai.com/docs/api-reference/chat`
//
// ```bash
// curl -X POST http://localhost:4000/v1/chat/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer sk-1234"
//
//	-d '{
//	    "model": "gpt-4o",
//	    "messages": [
//	        {
//	            "role": "user",
//	            "content": "Hello!"
//	        }
//	    ]
//	}'
//
// ```
func (r *OpenAIDeploymentChatService) Complete(ctx context.Context, model string, opts ...option.RequestOption) (res *OpenAIDeploymentChatCompleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if model == "" {
		err = errors.New("missing required model parameter")
		return
	}
	path := fmt.Sprintf("openai/deployments/%s/chat/completions", model)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type OpenAIDeploymentChatCompleteResponse = interface{}
