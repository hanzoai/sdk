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

// ResponseService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseService] method instead.
type ResponseService struct {
	Options    []option.RequestOption
	InputItems *ResponseInputItemService
}

// NewResponseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResponseService(opts ...option.RequestOption) (r *ResponseService) {
	r = &ResponseService{}
	r.Options = opts
	r.InputItems = NewResponseInputItemService(opts...)
	return
}

// Follows the OpenAI Responses API spec:
// https://platform.openai.com/docs/api-reference/responses
//
// ```bash
//
//	curl -X POST http://localhost:4000/v1/responses     -H "Content-Type: application/json"     -H "Authorization: Bearer sk-1234"     -d '{
//	    "model": "gpt-4o",
//	    "input": "Tell me about AI"
//	}'
//
// ```
func (r *ResponseService) New(ctx context.Context, opts ...option.RequestOption) (res *ResponseNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/responses"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Get a response by ID.
//
// Follows the OpenAI Responses API spec:
// https://platform.openai.com/docs/api-reference/responses/get
//
// ```bash
// curl -X GET http://localhost:4000/v1/responses/resp_abc123     -H "Authorization: Bearer sk-1234"
// ```
func (r *ResponseService) Get(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a response by ID.
//
// Follows the OpenAI Responses API spec:
// https://platform.openai.com/docs/api-reference/responses/delete
//
// ```bash
// curl -X DELETE http://localhost:4000/v1/responses/resp_abc123     -H "Authorization: Bearer sk-1234"
// ```
func (r *ResponseService) Delete(ctx context.Context, responseID string, opts ...option.RequestOption) (res *ResponseDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("v1/responses/%s", responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ResponseNewResponse = interface{}

type ResponseGetResponse = interface{}

type ResponseDeleteResponse = interface{}
