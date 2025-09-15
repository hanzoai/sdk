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

// AnthropicService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnthropicService] method instead.
type AnthropicService struct {
	Options []option.RequestOption
}

// NewAnthropicService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnthropicService(opts ...option.RequestOption) (r *AnthropicService) {
	r = &AnthropicService{}
	r.Options = opts
	return
}

// [Docs](https://docs.hanzo.ai/docs/anthropic_completion)
func (r *AnthropicService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AnthropicNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("anthropic/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/anthropic_completion)
func (r *AnthropicService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AnthropicGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("anthropic/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/anthropic_completion)
func (r *AnthropicService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AnthropicUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("anthropic/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/anthropic_completion)
func (r *AnthropicService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AnthropicDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("anthropic/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/anthropic_completion)
func (r *AnthropicService) Modify(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *AnthropicModifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("anthropic/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type AnthropicNewResponse = interface{}

type AnthropicGetResponse = interface{}

type AnthropicUpdateResponse = interface{}

type AnthropicDeleteResponse = interface{}

type AnthropicModifyResponse = interface{}
