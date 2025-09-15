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

// FineTuningJobCancelService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFineTuningJobCancelService] method instead.
type FineTuningJobCancelService struct {
	Options []option.RequestOption
}

// NewFineTuningJobCancelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFineTuningJobCancelService(opts ...option.RequestOption) (r *FineTuningJobCancelService) {
	r = &FineTuningJobCancelService{}
	r.Options = opts
	return
}

// Cancel a fine-tuning job.
//
// This is the equivalent of POST
// https://api.openai.com/v1/fine_tuning/jobs/{fine_tuning_job_id}/cancel
//
// Supported Query Params:
//
// - `custom_llm_provider`: Name of the LLM provider
// - `fine_tuning_job_id`: The ID of the fine-tuning job to cancel.
func (r *FineTuningJobCancelService) New(ctx context.Context, fineTuningJobID string, opts ...option.RequestOption) (res *FineTuningJobCancelNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if fineTuningJobID == "" {
		err = errors.New("missing required fine_tuning_job_id parameter")
		return
	}
	path := fmt.Sprintf("v1/fine_tuning/jobs/%s/cancel", fineTuningJobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type FineTuningJobCancelNewResponse = interface{}
