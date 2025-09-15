// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// CredentialService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCredentialService] method instead.
type CredentialService struct {
	Options []option.RequestOption
}

// NewCredentialService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCredentialService(opts ...option.RequestOption) (r *CredentialService) {
	r = &CredentialService{}
	r.Options = opts
	return
}

// [BETA] endpoint. This might change unexpectedly. Stores credential in DB.
// Reloads credentials in memory.
func (r *CredentialService) New(ctx context.Context, body CredentialNewParams, opts ...option.RequestOption) (res *CredentialNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// [BETA] endpoint. This might change unexpectedly.
func (r *CredentialService) List(ctx context.Context, opts ...option.RequestOption) (res *CredentialListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "credentials"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// [BETA] endpoint. This might change unexpectedly.
func (r *CredentialService) Delete(ctx context.Context, credentialName string, opts ...option.RequestOption) (res *CredentialDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if credentialName == "" {
		err = errors.New("missing required credential_name parameter")
		return
	}
	path := fmt.Sprintf("credentials/%s", credentialName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CredentialNewResponse = interface{}

type CredentialListResponse = interface{}

type CredentialDeleteResponse = interface{}

type CredentialNewParams struct {
	CredentialInfo   param.Field[interface{}] `json:"credential_info,required"`
	CredentialName   param.Field[string]      `json:"credential_name,required"`
	CredentialValues param.Field[interface{}] `json:"credential_values"`
	ModelID          param.Field[string]      `json:"model_id"`
}

func (r CredentialNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
