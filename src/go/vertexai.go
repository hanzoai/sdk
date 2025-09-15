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

// VertexAIService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVertexAIService] method instead.
type VertexAIService struct {
	Options []option.RequestOption
}

// NewVertexAIService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVertexAIService(opts ...option.RequestOption) (r *VertexAIService) {
	r = &VertexAIService{}
	r.Options = opts
	return
}

// Call LLM proxy via Vertex AI SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/vertex_ai)
func (r *VertexAIService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *VertexAINewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("vertex_ai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Call LLM proxy via Vertex AI SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/vertex_ai)
func (r *VertexAIService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *VertexAIGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("vertex_ai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Call LLM proxy via Vertex AI SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/vertex_ai)
func (r *VertexAIService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *VertexAIUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("vertex_ai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Call LLM proxy via Vertex AI SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/vertex_ai)
func (r *VertexAIService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *VertexAIDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("vertex_ai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Call LLM proxy via Vertex AI SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/vertex_ai)
func (r *VertexAIService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *VertexAIPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("vertex_ai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type VertexAINewResponse = interface{}

type VertexAIGetResponse = interface{}

type VertexAIUpdateResponse = interface{}

type VertexAIDeleteResponse = interface{}

type VertexAIPatchResponse = interface{}
