// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// DeleteService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeleteService] method instead.
type DeleteService struct {
	Options []option.RequestOption
}

// NewDeleteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDeleteService(opts ...option.RequestOption) (r *DeleteService) {
	r = &DeleteService{}
	r.Options = opts
	return
}

// Delete Allowed Ip
func (r *DeleteService) NewAllowedIP(ctx context.Context, body DeleteNewAllowedIPParams, opts ...option.RequestOption) (res *DeleteNewAllowedIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "delete/allowed_ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DeleteNewAllowedIPResponse = interface{}

type DeleteNewAllowedIPParams struct {
	IPAddress IPAddressParam `json:"ip_address,required"`
}

func (r DeleteNewAllowedIPParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IPAddress)
}
