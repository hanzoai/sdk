// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"
	"net/url"

	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// HealthService contains methods and other services that help with interacting
// with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r *HealthService) {
	r = &HealthService{}
	r.Options = opts
	return
}

// 🚨 USE `/health/liveliness` to health check the proxy 🚨
//
// See more 👉 https://docs.hanzo.ai/docs/proxy/health
//
// # Check the health of all the endpoints in config.yaml
//
// To run health checks in the background, add this to config.yaml:
//
// ```
// general_settings:
//
//	# ... other settings
//	background_health_checks: True
//
// ```
//
// else, the health checks will be run on models when /health is called.
func (r *HealthService) CheckAll(ctx context.Context, query HealthCheckAllParams, opts ...option.RequestOption) (res *HealthCheckAllResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Unprotected endpoint for checking if worker is alive
func (r *HealthService) CheckLiveliness(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckLivelinessResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/liveliness"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unprotected endpoint for checking if worker is alive
func (r *HealthService) CheckLiveness(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckLivenessResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/liveness"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Unprotected endpoint for checking if worker can receive requests
func (r *HealthService) CheckReadiness(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckReadinessResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/readiness"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Use this admin-only endpoint to check if the service is healthy.
//
// Example:
//
// ```
// curl -L -X GET 'http://0.0.0.0:4000/health/services?service=datadog'     -H 'Authorization: Bearer sk-1234'
// ```
func (r *HealthService) CheckServices(ctx context.Context, query HealthCheckServicesParams, opts ...option.RequestOption) (res *HealthCheckServicesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HealthCheckAllResponse = interface{}

type HealthCheckLivelinessResponse = interface{}

type HealthCheckLivenessResponse = interface{}

type HealthCheckReadinessResponse = interface{}

type HealthCheckServicesResponse = interface{}

type HealthCheckAllParams struct {
	// Specify the model name (optional)
	Model param.Field[string] `query:"model"`
}

// URLQuery serializes [HealthCheckAllParams]'s query parameters as `url.Values`.
func (r HealthCheckAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type HealthCheckServicesParams struct {
	// Specify the service being hit.
	Service param.Field[HealthCheckServicesParamsService] `query:"service,required"`
}

// URLQuery serializes [HealthCheckServicesParams]'s query parameters as
// `url.Values`.
func (r HealthCheckServicesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specify the service being hit.
type HealthCheckServicesParamsService string

const (
	HealthCheckServicesParamsServiceSlackBudgetAlerts HealthCheckServicesParamsService = "slack_budget_alerts"
	HealthCheckServicesParamsServiceLangfuse          HealthCheckServicesParamsService = "langfuse"
	HealthCheckServicesParamsServiceSlack             HealthCheckServicesParamsService = "slack"
	HealthCheckServicesParamsServiceOpenmeter         HealthCheckServicesParamsService = "openmeter"
	HealthCheckServicesParamsServiceWebhook           HealthCheckServicesParamsService = "webhook"
	HealthCheckServicesParamsServiceEmail             HealthCheckServicesParamsService = "email"
	HealthCheckServicesParamsServiceBraintrust        HealthCheckServicesParamsService = "braintrust"
	HealthCheckServicesParamsServiceDatadog           HealthCheckServicesParamsService = "datadog"
)

func (r HealthCheckServicesParamsService) IsKnown() bool {
	switch r {
	case HealthCheckServicesParamsServiceSlackBudgetAlerts, HealthCheckServicesParamsServiceLangfuse, HealthCheckServicesParamsServiceSlack, HealthCheckServicesParamsServiceOpenmeter, HealthCheckServicesParamsServiceWebhook, HealthCheckServicesParamsServiceEmail, HealthCheckServicesParamsServiceBraintrust, HealthCheckServicesParamsServiceDatadog:
		return true
	}
	return false
}
