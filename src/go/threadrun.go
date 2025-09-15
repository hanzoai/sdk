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

// ThreadRunService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadRunService] method instead.
type ThreadRunService struct {
	Options []option.RequestOption
}

// NewThreadRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreadRunService(opts ...option.RequestOption) (r *ThreadRunService) {
	r = &ThreadRunService{}
	r.Options = opts
	return
}

// Create a run.
//
// API Reference: https://platform.openai.com/docs/api-reference/runs/createRun
func (r *ThreadRunService) New(ctx context.Context, threadID string, opts ...option.RequestOption) (res *ThreadRunNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("v1/threads/%s/runs", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ThreadRunNewResponse = interface{}
