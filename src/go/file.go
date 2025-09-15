// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/hanzoai/go-sdk/internal/apiform"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// FileService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFileService] method instead.
type FileService struct {
	Options []option.RequestOption
	Content *FileContentService
}

// NewFileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFileService(opts ...option.RequestOption) (r *FileService) {
	r = &FileService{}
	r.Options = opts
	r.Content = NewFileContentService(opts...)
	return
}

// Upload a file that can be used across - Assistants API, Batch API This is the
// equivalent of POST https://api.openai.com/v1/files
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/files/create
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/files         -H "Authorization: Bearer sk-1234"         -F purpose="batch"         -F file="@mydata.jsonl"
//
// ```
func (r *FileService) New(ctx context.Context, provider string, body FileNewParams, opts ...option.RequestOption) (res *FileNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/files", provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns information about a specific file. that can be used across - Assistants
// API, Batch API This is the equivalent of GET
// https://api.openai.com/v1/files/{file_id}
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/files/retrieve
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/files/file-abc123         -H "Authorization: Bearer sk-1234"
//
// ```
func (r *FileService) Get(ctx context.Context, provider string, fileID string, opts ...option.RequestOption) (res *FileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/files/%s", provider, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns information about a specific file. that can be used across - Assistants
// API, Batch API This is the equivalent of GET https://api.openai.com/v1/files/
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/files/list
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/files        -H "Authorization: Bearer sk-1234"
//
// ```
func (r *FileService) List(ctx context.Context, provider string, query FileListParams, opts ...option.RequestOption) (res *FileListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/files", provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a specified file. that can be used across - Assistants API, Batch API
// This is the equivalent of DELETE https://api.openai.com/v1/files/{file_id}
//
// Supports Identical Params as:
// https://platform.openai.com/docs/api-reference/files/delete
//
// # Example Curl
//
// ```
// curl http://localhost:4000/v1/files/file-abc123     -X DELETE     -H "Authorization: Bearer $OPENAI_API_KEY"
//
// ```
func (r *FileService) Delete(ctx context.Context, provider string, fileID string, opts ...option.RequestOption) (res *FileDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	if fileID == "" {
		err = errors.New("missing required file_id parameter")
		return
	}
	path := fmt.Sprintf("%s/v1/files/%s", provider, fileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type FileNewResponse = interface{}

type FileGetResponse = interface{}

type FileListResponse = interface{}

type FileDeleteResponse = interface{}

type FileNewParams struct {
	File              param.Field[io.Reader] `json:"file,required" format:"binary"`
	Purpose           param.Field[string]    `json:"purpose,required"`
	CustomLlmProvider param.Field[string]    `json:"custom_llm_provider"`
}

func (r FileNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type FileListParams struct {
	Purpose param.Field[string] `query:"purpose"`
}

// URLQuery serializes [FileListParams]'s query parameters as `url.Values`.
func (r FileListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
