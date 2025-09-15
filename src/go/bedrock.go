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

// BedrockService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBedrockService] method instead.
type BedrockService struct {
	Options []option.RequestOption
}

// NewBedrockService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBedrockService(opts ...option.RequestOption) (r *BedrockService) {
	r = &BedrockService{}
	r.Options = opts
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/bedrock)
func (r *BedrockService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *BedrockNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("bedrock/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/bedrock)
func (r *BedrockService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *BedrockGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("bedrock/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/bedrock)
func (r *BedrockService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *BedrockUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("bedrock/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/bedrock)
func (r *BedrockService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *BedrockDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("bedrock/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/bedrock)
func (r *BedrockService) Patch(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *BedrockPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("bedrock/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type BedrockNewResponse = interface{}

type BedrockGetResponse = interface{}

type BedrockUpdateResponse = interface{}

type BedrockDeleteResponse = interface{}

type BedrockPatchResponse = interface{}
