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

// LangfuseService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLangfuseService] method instead.
type LangfuseService struct {
	Options []option.RequestOption
}

// NewLangfuseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLangfuseService(opts ...option.RequestOption) (r *LangfuseService) {
	r = &LangfuseService{}
	r.Options = opts
	return
}

// Call Langfuse via LLM proxy. Works with Langfuse SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/langfuse)
func (r *LangfuseService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *LangfuseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("langfuse/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Call Langfuse via LLM proxy. Works with Langfuse SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/langfuse)
func (r *LangfuseService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *LangfuseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("langfuse/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Call Langfuse via LLM proxy. Works with Langfuse SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/langfuse)
func (r *LangfuseService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *LangfuseUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("langfuse/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Call Langfuse via LLM proxy. Works with Langfuse SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/langfuse)
func (r *LangfuseService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *LangfuseDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("langfuse/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Call Langfuse via LLM proxy. Works with Langfuse SDK.
//
// [Docs](https://docs.hanzo.ai/docs/pass_through/langfuse)
func (r *LangfuseService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *LangfusePatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("langfuse/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type LangfuseNewResponse = interface{}

type LangfuseGetResponse = interface{}

type LangfuseUpdateResponse = interface{}

type LangfuseDeleteResponse = interface{}

type LangfusePatchResponse = interface{}
