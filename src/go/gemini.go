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

// GeminiService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeminiService] method instead.
type GeminiService struct {
	Options []option.RequestOption
}

// NewGeminiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGeminiService(opts ...option.RequestOption) (r *GeminiService) {
	r = &GeminiService{}
	r.Options = opts
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/google_ai_studio)
func (r *GeminiService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *GeminiNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("gemini/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/google_ai_studio)
func (r *GeminiService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *GeminiGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("gemini/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/google_ai_studio)
func (r *GeminiService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *GeminiUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("gemini/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/google_ai_studio)
func (r *GeminiService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *GeminiDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("gemini/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/google_ai_studio)
func (r *GeminiService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *GeminiPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("gemini/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type GeminiNewResponse = interface{}

type GeminiGetResponse = interface{}

type GeminiUpdateResponse = interface{}

type GeminiDeleteResponse = interface{}

type GeminiPatchResponse = interface{}
