// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// ImageGenerationService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageGenerationService] method instead.
type ImageGenerationService struct {
	Options []option.RequestOption
}

// NewImageGenerationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageGenerationService(opts ...option.RequestOption) (r *ImageGenerationService) {
	r = &ImageGenerationService{}
	r.Options = opts
	return
}

// Image Generation
func (r *ImageGenerationService) New(ctx context.Context, opts ...option.RequestOption) (res *ImageGenerationNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/images/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ImageGenerationNewResponse = interface{}
