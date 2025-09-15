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

// CohereService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCohereService] method instead.
type CohereService struct {
	Options []option.RequestOption
}

// NewCohereService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCohereService(opts ...option.RequestOption) (r *CohereService) {
	r = &CohereService{}
	r.Options = opts
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/cohere)
func (r *CohereService) New(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *CohereNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("cohere/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/cohere)
func (r *CohereService) Get(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *CohereGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("cohere/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/cohere)
func (r *CohereService) Update(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *CohereUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("cohere/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/cohere)
func (r *CohereService) Delete(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *CohereDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("cohere/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// [Docs](https://docs.hanzo.ai/docs/pass_through/cohere)
func (r *CohereService) Modify(ctx context.Context, endpoint string, opts ...option.RequestOption) (res *CohereModifyResponse, err error) {
	opts = append(r.Options[:], opts...)
	if endpoint == "" {
		err = errors.New("missing required endpoint parameter")
		return
	}
	path := fmt.Sprintf("cohere/%s", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &res, opts...)
	return
}

type CohereNewResponse = interface{}

type CohereGetResponse = interface{}

type CohereUpdateResponse = interface{}

type CohereDeleteResponse = interface{}

type CohereModifyResponse = interface{}
