// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// AddService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddService] method instead.
type AddService struct {
	Options []option.RequestOption
}

// NewAddService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddService(opts ...option.RequestOption) (r *AddService) {
	r = &AddService{}
	r.Options = opts
	return
}

// Add Allowed Ip
func (r *AddService) AddAllowedIP(ctx context.Context, body AddAddAllowedIPParams, opts ...option.RequestOption) (res *AddAddAllowedIPResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "add/allowed_ip"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type IPAddressParam struct {
	IP param.Field[string] `json:"ip,required"`
}

func (r IPAddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddAddAllowedIPResponse = interface{}

type AddAddAllowedIPParams struct {
	IPAddress IPAddressParam `json:"ip_address,required"`
}

func (r AddAddAllowedIPParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IPAddress)
}
