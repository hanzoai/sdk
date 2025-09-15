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

// ThreadMessageService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreadMessageService] method instead.
type ThreadMessageService struct {
	Options []option.RequestOption
}

// NewThreadMessageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreadMessageService(opts ...option.RequestOption) (r *ThreadMessageService) {
	r = &ThreadMessageService{}
	r.Options = opts
	return
}

// Create a message.
//
// API Reference -
// https://platform.openai.com/docs/api-reference/messages/createMessage
func (r *ThreadMessageService) New(ctx context.Context, threadID string, opts ...option.RequestOption) (res *ThreadMessageNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("v1/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Returns a list of messages for a given thread.
//
// API Reference -
// https://platform.openai.com/docs/api-reference/messages/listMessages
func (r *ThreadMessageService) List(ctx context.Context, threadID string, opts ...option.RequestOption) (res *ThreadMessageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return
	}
	path := fmt.Sprintf("v1/threads/%s/messages", threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreadMessageNewResponse = interface{}

type ThreadMessageListResponse = interface{}
