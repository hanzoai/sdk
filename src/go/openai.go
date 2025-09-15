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

// OpenAIService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOpenAIService] method instead.
type OpenAIService struct {
	Options     []option.RequestOption
	Deployments *OpenAIDeploymentService
}

// NewOpenAIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOpenAIService(opts ...option.RequestOption) (r *OpenAIService) {
	r = &OpenAIService{}
	r.Options = opts
	r.Deployments = NewOpenAIDeploymentService(opts...)
	return
}

// Simple pass-through for OpenAI. Use this if you want to directly send a request
// to OpenAI.
func (r *OpenAIService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *OpenAINewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("openai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Simple pass-through for OpenAI. Use this if you want to directly send a request
// to OpenAI.
func (r *OpenAIService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *OpenAIGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("openai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Simple pass-through for OpenAI. Use this if you want to directly send a request
// to OpenAI.
func (r *OpenAIService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *OpenAIUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("openai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Simple pass-through for OpenAI. Use this if you want to directly send a request
// to OpenAI.
func (r *OpenAIService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *OpenAIDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("openai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Simple pass-through for OpenAI. Use this if you want to directly send a request
// to OpenAI.
func (r *OpenAIService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *OpenAIPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("openai/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type OpenAINewResponse = interface{}

type OpenAIGetResponse = interface{}

type OpenAIUpdateResponse = interface{}

type OpenAIDeleteResponse = interface{}

type OpenAIPatchResponse = interface{}
