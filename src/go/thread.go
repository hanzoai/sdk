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

// ThreadService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadService] method instead.
type ThreadService struct {
	Options  []option.RequestOption
	Messages *ThreadMessageService
	Runs     *ThreadRunService
}

// NewThreadService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewThreadService(opts ...option.RequestOption) (r *ThreadService) {
	r = &ThreadService{}
	r.Options = opts
	r.Messages = NewThreadMessageService(opts...)
	r.Runs = NewThreadRunService(opts...)
	return
}

// Create a thread.
//
// API Reference -
// https://platform.openai.com/docs/api-reference/threads/createThread
func (r *ThreadService) New(ctx context.Context, opts ...option.RequestOption) (res *ThreadNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/threads"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retrieves a thread.
//
// API Reference - https://platform.openai.com/docs/api-reference/threads/getThread
func (r *ThreadService) Get(ctx context.Context, threadID string, opts ...option.RequestOption) (res *ThreadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("v1/threads/%s", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreadNewResponse = interface{}

type ThreadGetResponse = interface{}
