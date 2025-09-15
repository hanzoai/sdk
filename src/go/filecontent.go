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

// FileContentService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileContentService] method instead.
type FileContentService struct {
	Options []option.RequestOption
}

// NewFileContentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFileContentService(opts ...option.RequestOption) (r *FileContentService) {
	r = &FileContentService{}
	r.Options = opts
	return
}

// Returns information about a specific file. that can be used across - Assistants
// API, Batch API This is the equivalent of GET
// https://api.openai.com/v1/files/{file_id}/content
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/files/retrieve-contents
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/files/file-abc123/content         -H "Authorization: Bearer sk-1234"
//
// ```
func (r *FileContentService) Get(ctx context.Context, provider string, fileID string, opts ...option.RequestOption) (res *FileContentGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/files/%s/content", provider, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FileContentGetResponse = interface{}
