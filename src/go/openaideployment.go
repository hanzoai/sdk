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

// OpenAIDeploymentService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIDeploymentService] method instead.
type OpenAIDeploymentService struct {
	Options []option.RequestOption
	Chat    *OpenAIDeploymentChatService
}

// NewOpenAIDeploymentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOpenAIDeploymentService(opts ...option.RequestOption) (r *OpenAIDeploymentService) {
	r = &OpenAIDeploymentService{}
	r.Options = opts
	r.Chat = NewOpenAIDeploymentChatService(opts...)
	return
}

// Follows the exact same API spec as
// `OpenAI's Completions API https://platform.openai.com/docs/api-reference/completions`
//
// ```bash
// curl -X POST http://localhost:4000/v1/completions
// -H "Content-Type: application/json"
// -H "Authorization: Bearer sk-1234"
//
//	-d '{
//	    "model": "gpt-3.5-turbo-instruct",
//	    "prompt": "Once upon a time",
//	    "max_tokens": 50,
//	    "temperature": 0.7
//	}'
//
// ```
func (r *OpenAIDeploymentService) Complete(ctx context.Context, model string, opts ...option.RequestOption) (res *OpenAIDeploymentCompleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if model == "" {
		err = errors.New("missing required model parameter")
		return
	}
	path := fmt.Sprintf("openai/deployments/%s/completions", model)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Follows the exact same API spec as
// `OpenAI's Embeddings API https://platform.openai.com/docs/api-reference/embeddings`
//
// ```bash
// curl -X POST http://localhost:4000/v1/embeddings
// -H "Content-Type: application/json"
// -H "Authorization: Bearer sk-1234"
//
//	-d '{
//	    "model": "text-embedding-ada-002",
//	    "input": "The quick brown fox jumps over the lazy dog"
//	}'
//
// ```
func (r *OpenAIDeploymentService) Embed(ctx context.Context, model string, opts ...option.RequestOption) (res *OpenAIDeploymentEmbedResponse, err error) {
	opts = append(r.Options[:], opts...)
	if model == "" {
		err = errors.New("missing required model parameter")
		return
	}
	path := fmt.Sprintf("openai/deployments/%s/embeddings", model)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type OpenAIDeploymentCompleteResponse = interface{}

type OpenAIDeploymentEmbedResponse = interface{}
