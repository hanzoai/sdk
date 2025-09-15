// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/hanzoai/go-sdk/internal/apijson"
	"github.com/hanzoai/go-sdk/internal/apiquery"
	"github.com/hanzoai/go-sdk/internal/param"
	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
	"github.com/hanzoai/go-sdk/shared"
	"github.com/tidwall/gjson"
)

// KeyService contains methods and other services that help with interacting with
// the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKeyService] method instead.
type KeyService struct {
	Options    []option.RequestOption
	Regenerate *KeyRegenerateService
}

// NewKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewKeyService(opts ...option.RequestOption) (r *KeyService) {
	r = &KeyService{}
	r.Options = opts
	r.Regenerate = NewKeyRegenerateService(opts...)
	return
}

// Update an existing API key's parameters.
//
// Parameters:
//
//   - key: str - The key to update
//   - key_alias: Optional[str] - User-friendly key alias
//   - user_id: Optional[str] - User ID associated with key
//   - team_id: Optional[str] - Team ID associated with key
//   - budget_id: Optional[str] - The budget id associated with the key. Created by
//     calling `/budget/new`.
//   - models: Optional[list] - Model_name's a user is allowed to call
//   - tags: Optional[List[str]] - Tags for organizing keys (Enterprise only)
//   - enforced_params: Optional[List[str]] - List of enforced params for the key
//     (Enterprise only).
//     [Docs](https://docs.hanzo.ai/docs/proxy/enterprise#enforce-required-params-for-llm-requests)
//   - spend: Optional[float] - Amount spent by key
//   - max_budget: Optional[float] - Max budget for key
//   - model_max_budget: Optional[Dict[str, BudgetConfig]] - Model-specific budgets
//     {"gpt-4": {"budget_limit": 0.0005, "time_period": "30d"}}
//   - budget_duration: Optional[str] - Budget reset period ("30d", "1h", etc.)
//   - soft_budget: Optional[float] - [TODO] Soft budget limit (warning vs. hard
//     stop). Will trigger a slack alert when this soft budget is reached.
//   - max_parallel_requests: Optional[int] - Rate limit for parallel requests
//   - metadata: Optional[dict] - Metadata for key. Example {"team": "core-infra",
//     "app": "app2"}
//   - tpm_limit: Optional[int] - Tokens per minute limit
//   - rpm_limit: Optional[int] - Requests per minute limit
//   - model_rpm_limit: Optional[dict] - Model-specific RPM limits {"gpt-4": 100,
//     "claude-v1": 200}
//   - model_tpm_limit: Optional[dict] - Model-specific TPM limits {"gpt-4": 100000,
//     "claude-v1": 200000}
//   - allowed_cache_controls: Optional[list] - List of allowed cache control values
//   - duration: Optional[str] - Key validity duration ("30d", "1h", etc.)
//   - permissions: Optional[dict] - Key-specific permissions
//   - send_invite_email: Optional[bool] - Send invite email to user_id
//   - guardrails: Optional[List[str]] - List of active guardrails for the key
//   - blocked: Optional[bool] - Whether the key is blocked
//   - aliases: Optional[dict] - Model aliases for the key -
//     [Docs](https://llm.vercel.app/docs/proxy/virtual_keys#model-aliases)
//   - config: Optional[dict] - [DEPRECATED PARAM] Key-specific config.
//   - temp_budget_increase: Optional[float] - Temporary budget increase for the key
//     (Enterprise only).
//   - temp_budget_expiry: Optional[str] - Expiry time for the temporary budget
//     increase (Enterprise only).
//
// Example:
//
// ```bash
//
//	curl --location 'http://0.0.0.0:4000/key/update'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "key": "sk-1234",
//	    "key_alias": "my-key",
//	    "user_id": "user-1234",
//	    "team_id": "team-1234",
//	    "max_budget": 100,
//	    "metadata": {"any_key": "any-val"},
//	}'
//
// ```
func (r *KeyService) Update(ctx context.Context, params KeyUpdateParams, opts ...option.RequestOption) (res *KeyUpdateResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "key/update"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all keys for a given user / team / organization.
//
// Returns: { "keys": List[str] or List[UserAPIKeyAuth], "total_count": int,
// "current_page": int, "total_pages": int, }
func (r *KeyService) List(ctx context.Context, query KeyListParams, opts ...option.RequestOption) (res *KeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "key/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a key from the key management system.
//
// Parameters::
//
//   - keys (List[str]): A list of keys or hashed keys to delete. Example {"keys":
//     ["sk-QWrxEynunsNpV1zT48HIrw",
//     "837e17519f44683334df5291321d97b8bf1098cd490e49e215f6fea935aa28be"]}
//   - key_aliases (List[str]): A list of key aliases to delete. Can be passed
//     instead of `keys`.Example {"key_aliases": ["alias1", "alias2"]}
//
// Returns:
//
//   - deleted_keys (List[str]): A list of deleted keys. Example {"deleted_keys":
//     ["sk-QWrxEynunsNpV1zT48HIrw",
//     "837e17519f44683334df5291321d97b8bf1098cd490e49e215f6fea935aa28be"]}
//
// Example:
//
// ```bash
//
//	curl --location 'http://0.0.0.0:4000/key/delete'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "keys": ["sk-QWrxEynunsNpV1zT48HIrw"]
//	}'
//
// ```
//
// Raises: HTTPException: If an error occurs during key deletion.
func (r *KeyService) Delete(ctx context.Context, params KeyDeleteParams, opts ...option.RequestOption) (res *KeyDeleteResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "key/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Block an Virtual key from making any requests.
//
// Parameters:
//
//   - key: str - The key to block. Can be either the unhashed key (sk-...) or the
//     hashed key value
//
// Example:
//
// ```bash
//
//	curl --location 'http://0.0.0.0:4000/key/block'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "key": "sk-Fn8Ej39NxjAXrvpUGKghGw"
//	}'
//
// ```
//
// Note: This is an admin-only endpoint. Only proxy admins can block keys.
func (r *KeyService) Block(ctx context.Context, params KeyBlockParams, opts ...option.RequestOption) (res *KeyBlockResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "key/block"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Check the health of the key
//
// Checks:
//
// - If key based logging is configured correctly - sends a test log
//
// # Usage
//
// # Pass the key in the request header
//
// ```bash
// curl -X POST "http://localhost:4000/key/health"      -H "Authorization: Bearer sk-1234"      -H "Content-Type: application/json"
// ```
//
// Response when logging callbacks are setup correctly:
//
// ```json
//
//	{
//	  "key": "healthy",
//	  "logging_callbacks": {
//	    "callbacks": ["gcs_bucket"],
//	    "status": "healthy",
//	    "details": "No logger exceptions triggered, system is healthy. Manually check if logs were sent to ['gcs_bucket']"
//	  }
//	}
//
// ```
//
// Response when logging callbacks are not setup correctly:
//
// ```json
//
//	{
//	  "key": "unhealthy",
//	  "logging_callbacks": {
//	    "callbacks": ["gcs_bucket"],
//	    "status": "unhealthy",
//	    "details": "Logger exceptions triggered, system is unhealthy: Failed to load vertex credentials. Check to see if credentials containing partial/invalid information."
//	  }
//	}
//
// ```
func (r *KeyService) CheckHealth(ctx context.Context, opts ...option.RequestOption) (res *KeyCheckHealthResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "key/health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Generate an API key based on the provided data.
//
// Docs: https://docs.hanzo.ai/docs/proxy/virtual_keys
//
// Parameters:
//
//   - duration: Optional[str] - Specify the length of time the token is valid for.
//     You can set duration as seconds ("30s"), minutes ("30m"), hours ("30h"), days
//     ("30d").
//   - key_alias: Optional[str] - User defined key alias
//   - key: Optional[str] - User defined key value. If not set, a 16-digit unique
//     sk-key is created for you.
//   - team_id: Optional[str] - The team id of the key
//   - user_id: Optional[str] - The user id of the key
//   - budget_id: Optional[str] - The budget id associated with the key. Created by
//     calling `/budget/new`.
//   - models: Optional[list] - Model_name's a user is allowed to call. (if empty,
//     key is allowed to call all models)
//   - aliases: Optional[dict] - Any alias mappings, on top of anything in the
//     config.yaml model list. -
//     https://docs.hanzo.ai/docs/proxy/virtual_keys#managing-auth---upgradedowngrade-models
//   - config: Optional[dict] - any key-specific configs, overrides config in
//     config.yaml
//   - spend: Optional[int] - Amount spent by key. Default is 0. Will be updated by
//     proxy whenever key is used.
//     https://docs.hanzo.ai/docs/proxy/virtual_keys#managing-auth---tracking-spend
//   - send_invite_email: Optional[bool] - Whether to send an invite email to the
//     user_id, with the generate key
//   - max_budget: Optional[float] - Specify max budget for a given key.
//   - budget_duration: Optional[str] - Budget is reset at the end of specified
//     duration. If not set, budget is never reset. You can set duration as seconds
//     ("30s"), minutes ("30m"), hours ("30h"), days ("30d").
//   - max_parallel_requests: Optional[int] - Rate limit a user based on the number
//     of parallel requests. Raises 429 error, if user's parallel requests > x.
//   - metadata: Optional[dict] - Metadata for key, store information for key.
//     Example metadata = {"team": "core-infra", "app": "app2", "email": "z@hanzo.ai"
//     }
//   - guardrails: Optional[List[str]] - List of active guardrails for the key
//   - permissions: Optional[dict] - key-specific permissions. Currently just used
//     for turning off pii masking (if connected). Example - {"pii": false}
//   - model_max_budget: Optional[Dict[str, BudgetConfig]] - Model-specific budgets
//     {"gpt-4": {"budget_limit": 0.0005, "time_period": "30d"}}}. IF null or {} then
//     no model specific budget.
//   - model_rpm_limit: Optional[dict] - key-specific model rpm limit. Example -
//     {"text-davinci-002": 1000, "gpt-3.5-turbo": 1000}. IF null or {} then no model
//     specific rpm limit.
//   - model_tpm_limit: Optional[dict] - key-specific model tpm limit. Example -
//     {"text-davinci-002": 1000, "gpt-3.5-turbo": 1000}. IF null or {} then no model
//     specific tpm limit.
//   - allowed_cache_controls: Optional[list] - List of allowed cache control values.
//     Example - ["no-cache", "no-store"]. See all values -
//     https://docs.hanzo.ai/docs/proxy/caching#turn-on--off-caching-per-request
//   - blocked: Optional[bool] - Whether the key is blocked.
//   - rpm_limit: Optional[int] - Specify rpm limit for a given key (Requests per
//     minute)
//   - tpm_limit: Optional[int] - Specify tpm limit for a given key (Tokens per
//     minute)
//   - soft_budget: Optional[float] - Specify soft budget for a given key. Will
//     trigger a slack alert when this soft budget is reached.
//   - tags: Optional[List[str]] - Tags for
//     [tracking spend](https://llm.vercel.app/docs/proxy/enterprise#tracking-spend-for-custom-tags)
//     and/or doing
//     [tag-based routing](https://llm.vercel.app/docs/proxy/tag_routing).
//   - enforced_params: Optional[List[str]] - List of enforced params for the key
//     (Enterprise only).
//     [Docs](https://docs.hanzo.ai/docs/proxy/enterprise#enforce-required-params-for-llm-requests)
//
// Examples:
//
// 1. Allow users to turn on/off pii masking
//
// ```bash
//
//	curl --location 'http://0.0.0.0:4000/key/generate'         --header 'Authorization: Bearer sk-1234'         --header 'Content-Type: application/json'         --data '{
//	        "permissions": {"allow_pii_controls": true}
//	}'
//
// ```
//
// Returns:
//
//   - key: (str) The generated api key
//   - expires: (datetime) Datetime object for when key expires.
//   - user_id: (str) Unique user id - used for tracking spend across multiple keys
//     for same user id.
func (r *KeyService) Generate(ctx context.Context, params KeyGenerateParams, opts ...option.RequestOption) (res *GenerateKeyResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "key/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Regenerate an existing API key while optionally updating its parameters.
//
// Parameters:
//
//   - key: str (path parameter) - The key to regenerate
//   - data: Optional[RegenerateKeyRequest] - Request body containing optional
//     parameters to update
//   - key_alias: Optional[str] - User-friendly key alias
//   - user_id: Optional[str] - User ID associated with key
//   - team_id: Optional[str] - Team ID associated with key
//   - models: Optional[list] - Model_name's a user is allowed to call
//   - tags: Optional[List[str]] - Tags for organizing keys (Enterprise only)
//   - spend: Optional[float] - Amount spent by key
//   - max_budget: Optional[float] - Max budget for key
//   - model_max_budget: Optional[Dict[str, BudgetConfig]] - Model-specific budgets
//     {"gpt-4": {"budget_limit": 0.0005, "time_period": "30d"}}
//   - budget_duration: Optional[str] - Budget reset period ("30d", "1h", etc.)
//   - soft_budget: Optional[float] - Soft budget limit (warning vs. hard stop).
//     Will trigger a slack alert when this soft budget is reached.
//   - max_parallel_requests: Optional[int] - Rate limit for parallel requests
//   - metadata: Optional[dict] - Metadata for key. Example {"team": "core-infra",
//     "app": "app2"}
//   - tpm_limit: Optional[int] - Tokens per minute limit
//   - rpm_limit: Optional[int] - Requests per minute limit
//   - model_rpm_limit: Optional[dict] - Model-specific RPM limits {"gpt-4": 100,
//     "claude-v1": 200}
//   - model_tpm_limit: Optional[dict] - Model-specific TPM limits {"gpt-4":
//     100000, "claude-v1": 200000}
//   - allowed_cache_controls: Optional[list] - List of allowed cache control
//     values
//   - duration: Optional[str] - Key validity duration ("30d", "1h", etc.)
//   - permissions: Optional[dict] - Key-specific permissions
//   - guardrails: Optional[List[str]] - List of active guardrails for the key
//   - blocked: Optional[bool] - Whether the key is blocked
//
// Returns:
//
// - GenerateKeyResponse containing the new key and its updated parameters
//
// Example:
//
// ```bash
//
//	curl --location --request POST 'http://localhost:4000/key/sk-1234/regenerate'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data-raw '{
//	    "max_budget": 100,
//	    "metadata": {"team": "core-infra"},
//	    "models": ["gpt-4", "gpt-3.5-turbo"]
//	}'
//
// ```
//
// Note: This is an Enterprise feature. It requires a premium license to use.
func (r *KeyService) RegenerateByKey(ctx context.Context, key string, params KeyRegenerateByKeyParams, opts ...option.RequestOption) (res *GenerateKeyResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	if key == "" {
		err = errors.New("missing required key parameter")
		return
	}
	path := fmt.Sprintf("key/%s/regenerate", key)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve information about a key. Parameters: key: Optional[str] = Query
// parameter representing the key in the request user_api_key_dict: UserAPIKeyAuth
// = Dependency representing the user's API key Returns: Dict containing the key
// and its associated information
//
// Example Curl:
//
// ```
// curl -X GET "http://0.0.0.0:4000/key/info?key=sk-02Wr4IAlN3NvPXvL5JVvDA" -H "Authorization: Bearer sk-1234"
// ```
//
// Example Curl - if no key is passed, it will use the Key Passed in Authorization
// Header
//
// ```
// curl -X GET "http://0.0.0.0:4000/key/info" -H "Authorization: Bearer sk-02Wr4IAlN3NvPXvL5JVvDA"
// ```
func (r *KeyService) GetInfo(ctx context.Context, query KeyGetInfoParams, opts ...option.RequestOption) (res *KeyGetInfoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "key/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Unblock a Virtual key to allow it to make requests again.
//
// Parameters:
//
//   - key: str - The key to unblock. Can be either the unhashed key (sk-...) or the
//     hashed key value
//
// Example:
//
// ```bash
//
//	curl --location 'http://0.0.0.0:4000/key/unblock'     --header 'Authorization: Bearer sk-1234'     --header 'Content-Type: application/json'     --data '{
//	    "key": "sk-Fn8Ej39NxjAXrvpUGKghGw"
//	}'
//
// ```
//
// Note: This is an admin-only endpoint. Only proxy admins can unblock keys.
func (r *KeyService) Unblock(ctx context.Context, params KeyUnblockParams, opts ...option.RequestOption) (res *KeyUnblockResponse, err error) {
	if params.LlmChangedBy.Present {
		opts = append(opts, option.WithHeader("llm-changed-by", fmt.Sprintf("%s", params.LlmChangedBy)))
	}
	opts = append(r.Options[:], opts...)
	path := "key/unblock"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BlockKeyRequestParam struct {
	Key param.Field[string] `json:"key,required"`
}

func (r BlockKeyRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GenerateKeyResponse struct {
	Expires              time.Time               `json:"expires,required,nullable" format:"date-time"`
	Key                  string                  `json:"key,required"`
	Token                string                  `json:"token,nullable"`
	Aliases              interface{}             `json:"aliases,nullable"`
	AllowedCacheControls []interface{}           `json:"allowed_cache_controls,nullable"`
	Blocked              bool                    `json:"blocked,nullable"`
	BudgetDuration       string                  `json:"budget_duration,nullable"`
	BudgetID             string                  `json:"budget_id,nullable"`
	Config               interface{}             `json:"config,nullable"`
	CreatedBy            string                  `json:"created_by,nullable"`
	Duration             string                  `json:"duration,nullable"`
	EnforcedParams       []string                `json:"enforced_params,nullable"`
	Guardrails           []string                `json:"guardrails,nullable"`
	KeyAlias             string                  `json:"key_alias,nullable"`
	KeyName              string                  `json:"key_name,nullable"`
	LlmBudgetTable       interface{}             `json:"llm_budget_table"`
	MaxBudget            float64                 `json:"max_budget,nullable"`
	MaxParallelRequests  int64                   `json:"max_parallel_requests,nullable"`
	Metadata             interface{}             `json:"metadata,nullable"`
	ModelMaxBudget       interface{}             `json:"model_max_budget,nullable"`
	ModelRpmLimit        interface{}             `json:"model_rpm_limit,nullable"`
	ModelTpmLimit        interface{}             `json:"model_tpm_limit,nullable"`
	Models               []interface{}           `json:"models,nullable"`
	Permissions          interface{}             `json:"permissions,nullable"`
	RpmLimit             int64                   `json:"rpm_limit,nullable"`
	Spend                float64                 `json:"spend,nullable"`
	Tags                 []string                `json:"tags,nullable"`
	TeamID               string                  `json:"team_id,nullable"`
	TokenID              string                  `json:"token_id,nullable"`
	TpmLimit             int64                   `json:"tpm_limit,nullable"`
	UpdatedBy            string                  `json:"updated_by,nullable"`
	UserID               string                  `json:"user_id,nullable"`
	JSON                 generateKeyResponseJSON `json:"-"`
}

// generateKeyResponseJSON contains the JSON metadata for the struct
// [GenerateKeyResponse]
type generateKeyResponseJSON struct {
	Expires              apijson.Field
	Key                  apijson.Field
	Token                apijson.Field
	Aliases              apijson.Field
	AllowedCacheControls apijson.Field
	Blocked              apijson.Field
	BudgetDuration       apijson.Field
	BudgetID             apijson.Field
	Config               apijson.Field
	CreatedBy            apijson.Field
	Duration             apijson.Field
	EnforcedParams       apijson.Field
	Guardrails           apijson.Field
	KeyAlias             apijson.Field
	KeyName              apijson.Field
	LlmBudgetTable       apijson.Field
	MaxBudget            apijson.Field
	MaxParallelRequests  apijson.Field
	Metadata             apijson.Field
	ModelMaxBudget       apijson.Field
	ModelRpmLimit        apijson.Field
	ModelTpmLimit        apijson.Field
	Models               apijson.Field
	Permissions          apijson.Field
	RpmLimit             apijson.Field
	Spend                apijson.Field
	Tags                 apijson.Field
	TeamID               apijson.Field
	TokenID              apijson.Field
	TpmLimit             apijson.Field
	UpdatedBy            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *GenerateKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r generateKeyResponseJSON) RawJSON() string {
	return r.raw
}

type KeyUpdateResponse = interface{}

type KeyListResponse struct {
	CurrentPage int64                      `json:"current_page,nullable"`
	Keys        []KeyListResponseKeysUnion `json:"keys"`
	TotalCount  int64                      `json:"total_count,nullable"`
	TotalPages  int64                      `json:"total_pages,nullable"`
	JSON        keyListResponseJSON        `json:"-"`
}

// keyListResponseJSON contains the JSON metadata for the struct [KeyListResponse]
type keyListResponseJSON struct {
	CurrentPage apijson.Field
	Keys        apijson.Field
	TotalCount  apijson.Field
	TotalPages  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyListResponseJSON) RawJSON() string {
	return r.raw
}

// Return the row in the db
//
// Union satisfied by [shared.UnionString] or [KeyListResponseKeysUserAPIKeyAuth].
type KeyListResponseKeysUnion interface {
	ImplementsKeyListResponseKeysUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KeyListResponseKeysUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(KeyListResponseKeysUserAPIKeyAuth{}),
		},
	)
}

// Return the row in the db
type KeyListResponseKeysUserAPIKeyAuth struct {
	Token                string                                              `json:"token,nullable"`
	Aliases              interface{}                                         `json:"aliases"`
	AllowedCacheControls []interface{}                                       `json:"allowed_cache_controls,nullable"`
	AllowedModelRegion   KeyListResponseKeysUserAPIKeyAuthAllowedModelRegion `json:"allowed_model_region,nullable"`
	APIKey               string                                              `json:"api_key,nullable"`
	Blocked              bool                                                `json:"blocked,nullable"`
	BudgetDuration       string                                              `json:"budget_duration,nullable"`
	BudgetResetAt        time.Time                                           `json:"budget_reset_at,nullable" format:"date-time"`
	Config               interface{}                                         `json:"config"`
	CreatedAt            time.Time                                           `json:"created_at,nullable" format:"date-time"`
	CreatedBy            string                                              `json:"created_by,nullable"`
	EndUserID            string                                              `json:"end_user_id,nullable"`
	EndUserMaxBudget     float64                                             `json:"end_user_max_budget,nullable"`
	EndUserRpmLimit      int64                                               `json:"end_user_rpm_limit,nullable"`
	EndUserTpmLimit      int64                                               `json:"end_user_tpm_limit,nullable"`
	Expires              KeyListResponseKeysUserAPIKeyAuthExpiresUnion       `json:"expires,nullable" format:"date-time"`
	KeyAlias             string                                              `json:"key_alias,nullable"`
	KeyName              string                                              `json:"key_name,nullable"`
	LastRefreshedAt      float64                                             `json:"last_refreshed_at,nullable"`
	LlmBudgetTable       interface{}                                         `json:"llm_budget_table,nullable"`
	MaxBudget            float64                                             `json:"max_budget,nullable"`
	MaxParallelRequests  int64                                               `json:"max_parallel_requests,nullable"`
	Metadata             interface{}                                         `json:"metadata"`
	ModelMaxBudget       interface{}                                         `json:"model_max_budget"`
	ModelSpend           interface{}                                         `json:"model_spend"`
	Models               []interface{}                                       `json:"models"`
	OrgID                string                                              `json:"org_id,nullable"`
	ParentOtelSpan       interface{}                                         `json:"parent_otel_span"`
	Permissions          interface{}                                         `json:"permissions"`
	RpmLimit             int64                                               `json:"rpm_limit,nullable"`
	RpmLimitPerModel     map[string]int64                                    `json:"rpm_limit_per_model,nullable"`
	SoftBudget           float64                                             `json:"soft_budget,nullable"`
	SoftBudgetCooldown   bool                                                `json:"soft_budget_cooldown"`
	Spend                float64                                             `json:"spend"`
	TeamAlias            string                                              `json:"team_alias,nullable"`
	TeamBlocked          bool                                                `json:"team_blocked"`
	TeamID               string                                              `json:"team_id,nullable"`
	TeamMaxBudget        float64                                             `json:"team_max_budget,nullable"`
	TeamMember           Member                                              `json:"team_member,nullable"`
	TeamMemberSpend      float64                                             `json:"team_member_spend,nullable"`
	TeamMetadata         interface{}                                         `json:"team_metadata,nullable"`
	TeamModelAliases     interface{}                                         `json:"team_model_aliases,nullable"`
	TeamModels           []interface{}                                       `json:"team_models"`
	TeamRpmLimit         int64                                               `json:"team_rpm_limit,nullable"`
	TeamSpend            float64                                             `json:"team_spend,nullable"`
	TeamTpmLimit         int64                                               `json:"team_tpm_limit,nullable"`
	TpmLimit             int64                                               `json:"tpm_limit,nullable"`
	TpmLimitPerModel     map[string]int64                                    `json:"tpm_limit_per_model,nullable"`
	UpdatedAt            time.Time                                           `json:"updated_at,nullable" format:"date-time"`
	UpdatedBy            string                                              `json:"updated_by,nullable"`
	UserEmail            string                                              `json:"user_email,nullable"`
	UserID               string                                              `json:"user_id,nullable"`
	// Admin Roles: PROXY_ADMIN: admin over the platform PROXY_ADMIN_VIEW_ONLY: can
	// login, view all own keys, view all spend ORG_ADMIN: admin over a specific
	// organization, can create teams, users only within their organization
	//
	// Internal User Roles: INTERNAL_USER: can login, view/create/delete their own
	// keys, view their spend INTERNAL_USER_VIEW_ONLY: can login, view their own keys,
	// view their own spend
	//
	// Team Roles: TEAM: used for JWT auth
	//
	// Customer Roles: CUSTOMER: External users -> these are customers
	UserRole     KeyListResponseKeysUserAPIKeyAuthUserRole `json:"user_role,nullable"`
	UserRpmLimit int64                                     `json:"user_rpm_limit,nullable"`
	UserTpmLimit int64                                     `json:"user_tpm_limit,nullable"`
	JSON         keyListResponseKeysUserAPIKeyAuthJSON     `json:"-"`
}

// keyListResponseKeysUserAPIKeyAuthJSON contains the JSON metadata for the struct
// [KeyListResponseKeysUserAPIKeyAuth]
type keyListResponseKeysUserAPIKeyAuthJSON struct {
	Token                apijson.Field
	Aliases              apijson.Field
	AllowedCacheControls apijson.Field
	AllowedModelRegion   apijson.Field
	APIKey               apijson.Field
	Blocked              apijson.Field
	BudgetDuration       apijson.Field
	BudgetResetAt        apijson.Field
	Config               apijson.Field
	CreatedAt            apijson.Field
	CreatedBy            apijson.Field
	EndUserID            apijson.Field
	EndUserMaxBudget     apijson.Field
	EndUserRpmLimit      apijson.Field
	EndUserTpmLimit      apijson.Field
	Expires              apijson.Field
	KeyAlias             apijson.Field
	KeyName              apijson.Field
	LastRefreshedAt      apijson.Field
	LlmBudgetTable       apijson.Field
	MaxBudget            apijson.Field
	MaxParallelRequests  apijson.Field
	Metadata             apijson.Field
	ModelMaxBudget       apijson.Field
	ModelSpend           apijson.Field
	Models               apijson.Field
	OrgID                apijson.Field
	ParentOtelSpan       apijson.Field
	Permissions          apijson.Field
	RpmLimit             apijson.Field
	RpmLimitPerModel     apijson.Field
	SoftBudget           apijson.Field
	SoftBudgetCooldown   apijson.Field
	Spend                apijson.Field
	TeamAlias            apijson.Field
	TeamBlocked          apijson.Field
	TeamID               apijson.Field
	TeamMaxBudget        apijson.Field
	TeamMember           apijson.Field
	TeamMemberSpend      apijson.Field
	TeamMetadata         apijson.Field
	TeamModelAliases     apijson.Field
	TeamModels           apijson.Field
	TeamRpmLimit         apijson.Field
	TeamSpend            apijson.Field
	TeamTpmLimit         apijson.Field
	TpmLimit             apijson.Field
	TpmLimitPerModel     apijson.Field
	UpdatedAt            apijson.Field
	UpdatedBy            apijson.Field
	UserEmail            apijson.Field
	UserID               apijson.Field
	UserRole             apijson.Field
	UserRpmLimit         apijson.Field
	UserTpmLimit         apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *KeyListResponseKeysUserAPIKeyAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyListResponseKeysUserAPIKeyAuthJSON) RawJSON() string {
	return r.raw
}

func (r KeyListResponseKeysUserAPIKeyAuth) ImplementsKeyListResponseKeysUnion() {}

type KeyListResponseKeysUserAPIKeyAuthAllowedModelRegion string

const (
	KeyListResponseKeysUserAPIKeyAuthAllowedModelRegionEu KeyListResponseKeysUserAPIKeyAuthAllowedModelRegion = "eu"
	KeyListResponseKeysUserAPIKeyAuthAllowedModelRegionUs KeyListResponseKeysUserAPIKeyAuthAllowedModelRegion = "us"
)

func (r KeyListResponseKeysUserAPIKeyAuthAllowedModelRegion) IsKnown() bool {
	switch r {
	case KeyListResponseKeysUserAPIKeyAuthAllowedModelRegionEu, KeyListResponseKeysUserAPIKeyAuthAllowedModelRegionUs:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type KeyListResponseKeysUserAPIKeyAuthExpiresUnion interface {
	ImplementsKeyListResponseKeysUserAPIKeyAuthExpiresUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KeyListResponseKeysUserAPIKeyAuthExpiresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

// Admin Roles: PROXY_ADMIN: admin over the platform PROXY_ADMIN_VIEW_ONLY: can
// login, view all own keys, view all spend ORG_ADMIN: admin over a specific
// organization, can create teams, users only within their organization
//
// Internal User Roles: INTERNAL_USER: can login, view/create/delete their own
// keys, view their spend INTERNAL_USER_VIEW_ONLY: can login, view their own keys,
// view their own spend
//
// Team Roles: TEAM: used for JWT auth
//
// Customer Roles: CUSTOMER: External users -> these are customers
type KeyListResponseKeysUserAPIKeyAuthUserRole string

const (
	KeyListResponseKeysUserAPIKeyAuthUserRoleProxyAdmin         KeyListResponseKeysUserAPIKeyAuthUserRole = "proxy_admin"
	KeyListResponseKeysUserAPIKeyAuthUserRoleProxyAdminViewer   KeyListResponseKeysUserAPIKeyAuthUserRole = "proxy_admin_viewer"
	KeyListResponseKeysUserAPIKeyAuthUserRoleOrgAdmin           KeyListResponseKeysUserAPIKeyAuthUserRole = "org_admin"
	KeyListResponseKeysUserAPIKeyAuthUserRoleInternalUser       KeyListResponseKeysUserAPIKeyAuthUserRole = "internal_user"
	KeyListResponseKeysUserAPIKeyAuthUserRoleInternalUserViewer KeyListResponseKeysUserAPIKeyAuthUserRole = "internal_user_viewer"
	KeyListResponseKeysUserAPIKeyAuthUserRoleTeam               KeyListResponseKeysUserAPIKeyAuthUserRole = "team"
	KeyListResponseKeysUserAPIKeyAuthUserRoleCustomer           KeyListResponseKeysUserAPIKeyAuthUserRole = "customer"
)

func (r KeyListResponseKeysUserAPIKeyAuthUserRole) IsKnown() bool {
	switch r {
	case KeyListResponseKeysUserAPIKeyAuthUserRoleProxyAdmin, KeyListResponseKeysUserAPIKeyAuthUserRoleProxyAdminViewer, KeyListResponseKeysUserAPIKeyAuthUserRoleOrgAdmin, KeyListResponseKeysUserAPIKeyAuthUserRoleInternalUser, KeyListResponseKeysUserAPIKeyAuthUserRoleInternalUserViewer, KeyListResponseKeysUserAPIKeyAuthUserRoleTeam, KeyListResponseKeysUserAPIKeyAuthUserRoleCustomer:
		return true
	}
	return false
}

type KeyDeleteResponse = interface{}

type KeyBlockResponse struct {
	Token                string                       `json:"token,nullable"`
	Aliases              interface{}                  `json:"aliases"`
	AllowedCacheControls []interface{}                `json:"allowed_cache_controls,nullable"`
	Blocked              bool                         `json:"blocked,nullable"`
	BudgetDuration       string                       `json:"budget_duration,nullable"`
	BudgetResetAt        time.Time                    `json:"budget_reset_at,nullable" format:"date-time"`
	Config               interface{}                  `json:"config"`
	CreatedAt            time.Time                    `json:"created_at,nullable" format:"date-time"`
	CreatedBy            string                       `json:"created_by,nullable"`
	Expires              KeyBlockResponseExpiresUnion `json:"expires,nullable" format:"date-time"`
	KeyAlias             string                       `json:"key_alias,nullable"`
	KeyName              string                       `json:"key_name,nullable"`
	LlmBudgetTable       interface{}                  `json:"llm_budget_table,nullable"`
	MaxBudget            float64                      `json:"max_budget,nullable"`
	MaxParallelRequests  int64                        `json:"max_parallel_requests,nullable"`
	Metadata             interface{}                  `json:"metadata"`
	ModelMaxBudget       interface{}                  `json:"model_max_budget"`
	ModelSpend           interface{}                  `json:"model_spend"`
	Models               []interface{}                `json:"models"`
	OrgID                string                       `json:"org_id,nullable"`
	Permissions          interface{}                  `json:"permissions"`
	RpmLimit             int64                        `json:"rpm_limit,nullable"`
	SoftBudgetCooldown   bool                         `json:"soft_budget_cooldown"`
	Spend                float64                      `json:"spend"`
	TeamID               string                       `json:"team_id,nullable"`
	TpmLimit             int64                        `json:"tpm_limit,nullable"`
	UpdatedAt            time.Time                    `json:"updated_at,nullable" format:"date-time"`
	UpdatedBy            string                       `json:"updated_by,nullable"`
	UserID               string                       `json:"user_id,nullable"`
	JSON                 keyBlockResponseJSON         `json:"-"`
}

// keyBlockResponseJSON contains the JSON metadata for the struct
// [KeyBlockResponse]
type keyBlockResponseJSON struct {
	Token                apijson.Field
	Aliases              apijson.Field
	AllowedCacheControls apijson.Field
	Blocked              apijson.Field
	BudgetDuration       apijson.Field
	BudgetResetAt        apijson.Field
	Config               apijson.Field
	CreatedAt            apijson.Field
	CreatedBy            apijson.Field
	Expires              apijson.Field
	KeyAlias             apijson.Field
	KeyName              apijson.Field
	LlmBudgetTable       apijson.Field
	MaxBudget            apijson.Field
	MaxParallelRequests  apijson.Field
	Metadata             apijson.Field
	ModelMaxBudget       apijson.Field
	ModelSpend           apijson.Field
	Models               apijson.Field
	OrgID                apijson.Field
	Permissions          apijson.Field
	RpmLimit             apijson.Field
	SoftBudgetCooldown   apijson.Field
	Spend                apijson.Field
	TeamID               apijson.Field
	TpmLimit             apijson.Field
	UpdatedAt            apijson.Field
	UpdatedBy            apijson.Field
	UserID               apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *KeyBlockResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyBlockResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString] or [shared.UnionTime].
type KeyBlockResponseExpiresUnion interface {
	ImplementsKeyBlockResponseExpiresUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*KeyBlockResponseExpiresUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionTime(shared.UnionTime{})),
		},
	)
}

type KeyCheckHealthResponse struct {
	Key              KeyCheckHealthResponseKey              `json:"key"`
	LoggingCallbacks KeyCheckHealthResponseLoggingCallbacks `json:"logging_callbacks,nullable"`
	JSON             keyCheckHealthResponseJSON             `json:"-"`
}

// keyCheckHealthResponseJSON contains the JSON metadata for the struct
// [KeyCheckHealthResponse]
type keyCheckHealthResponseJSON struct {
	Key              apijson.Field
	LoggingCallbacks apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *KeyCheckHealthResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyCheckHealthResponseJSON) RawJSON() string {
	return r.raw
}

type KeyCheckHealthResponseKey string

const (
	KeyCheckHealthResponseKeyHealthy   KeyCheckHealthResponseKey = "healthy"
	KeyCheckHealthResponseKeyUnhealthy KeyCheckHealthResponseKey = "unhealthy"
)

func (r KeyCheckHealthResponseKey) IsKnown() bool {
	switch r {
	case KeyCheckHealthResponseKeyHealthy, KeyCheckHealthResponseKeyUnhealthy:
		return true
	}
	return false
}

type KeyCheckHealthResponseLoggingCallbacks struct {
	Callbacks []string                                     `json:"callbacks"`
	Details   string                                       `json:"details,nullable"`
	Status    KeyCheckHealthResponseLoggingCallbacksStatus `json:"status"`
	JSON      keyCheckHealthResponseLoggingCallbacksJSON   `json:"-"`
}

// keyCheckHealthResponseLoggingCallbacksJSON contains the JSON metadata for the
// struct [KeyCheckHealthResponseLoggingCallbacks]
type keyCheckHealthResponseLoggingCallbacksJSON struct {
	Callbacks   apijson.Field
	Details     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeyCheckHealthResponseLoggingCallbacks) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keyCheckHealthResponseLoggingCallbacksJSON) RawJSON() string {
	return r.raw
}

type KeyCheckHealthResponseLoggingCallbacksStatus string

const (
	KeyCheckHealthResponseLoggingCallbacksStatusHealthy   KeyCheckHealthResponseLoggingCallbacksStatus = "healthy"
	KeyCheckHealthResponseLoggingCallbacksStatusUnhealthy KeyCheckHealthResponseLoggingCallbacksStatus = "unhealthy"
)

func (r KeyCheckHealthResponseLoggingCallbacksStatus) IsKnown() bool {
	switch r {
	case KeyCheckHealthResponseLoggingCallbacksStatusHealthy, KeyCheckHealthResponseLoggingCallbacksStatusUnhealthy:
		return true
	}
	return false
}

type KeyGetInfoResponse = interface{}

type KeyUnblockResponse = interface{}

type KeyUpdateParams struct {
	Key                  param.Field[string]        `json:"key,required"`
	Aliases              param.Field[interface{}]   `json:"aliases"`
	AllowedCacheControls param.Field[[]interface{}] `json:"allowed_cache_controls"`
	Blocked              param.Field[bool]          `json:"blocked"`
	BudgetDuration       param.Field[string]        `json:"budget_duration"`
	BudgetID             param.Field[string]        `json:"budget_id"`
	Config               param.Field[interface{}]   `json:"config"`
	Duration             param.Field[string]        `json:"duration"`
	EnforcedParams       param.Field[[]string]      `json:"enforced_params"`
	Guardrails           param.Field[[]string]      `json:"guardrails"`
	KeyAlias             param.Field[string]        `json:"key_alias"`
	MaxBudget            param.Field[float64]       `json:"max_budget"`
	MaxParallelRequests  param.Field[int64]         `json:"max_parallel_requests"`
	Metadata             param.Field[interface{}]   `json:"metadata"`
	ModelMaxBudget       param.Field[interface{}]   `json:"model_max_budget"`
	ModelRpmLimit        param.Field[interface{}]   `json:"model_rpm_limit"`
	ModelTpmLimit        param.Field[interface{}]   `json:"model_tpm_limit"`
	Models               param.Field[[]interface{}] `json:"models"`
	Permissions          param.Field[interface{}]   `json:"permissions"`
	RpmLimit             param.Field[int64]         `json:"rpm_limit"`
	Spend                param.Field[float64]       `json:"spend"`
	Tags                 param.Field[[]string]      `json:"tags"`
	TeamID               param.Field[string]        `json:"team_id"`
	TempBudgetExpiry     param.Field[time.Time]     `json:"temp_budget_expiry" format:"date-time"`
	TempBudgetIncrease   param.Field[float64]       `json:"temp_budget_increase"`
	TpmLimit             param.Field[int64]         `json:"tpm_limit"`
	UserID               param.Field[string]        `json:"user_id"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r KeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeyListParams struct {
	// Include all keys for teams that user is an admin of.
	IncludeTeamKeys param.Field[bool] `query:"include_team_keys"`
	// Filter keys by key alias
	KeyAlias param.Field[string] `query:"key_alias"`
	// Filter keys by organization ID
	OrganizationID param.Field[string] `query:"organization_id"`
	// Page number
	Page param.Field[int64] `query:"page"`
	// Return full key object
	ReturnFullObject param.Field[bool] `query:"return_full_object"`
	// Page size
	Size param.Field[int64] `query:"size"`
	// Filter keys by team ID
	TeamID param.Field[string] `query:"team_id"`
	// Filter keys by user ID
	UserID param.Field[string] `query:"user_id"`
}

// URLQuery serializes [KeyListParams]'s query parameters as `url.Values`.
func (r KeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type KeyDeleteParams struct {
	KeyAliases param.Field[[]string] `json:"key_aliases"`
	Keys       param.Field[[]string] `json:"keys"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r KeyDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeyBlockParams struct {
	BlockKeyRequest BlockKeyRequestParam `json:"block_key_request,required"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r KeyBlockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockKeyRequest)
}

type KeyGenerateParams struct {
	Aliases              param.Field[interface{}]   `json:"aliases"`
	AllowedCacheControls param.Field[[]interface{}] `json:"allowed_cache_controls"`
	Blocked              param.Field[bool]          `json:"blocked"`
	BudgetDuration       param.Field[string]        `json:"budget_duration"`
	BudgetID             param.Field[string]        `json:"budget_id"`
	Config               param.Field[interface{}]   `json:"config"`
	Duration             param.Field[string]        `json:"duration"`
	EnforcedParams       param.Field[[]string]      `json:"enforced_params"`
	Guardrails           param.Field[[]string]      `json:"guardrails"`
	Key                  param.Field[string]        `json:"key"`
	KeyAlias             param.Field[string]        `json:"key_alias"`
	MaxBudget            param.Field[float64]       `json:"max_budget"`
	MaxParallelRequests  param.Field[int64]         `json:"max_parallel_requests"`
	Metadata             param.Field[interface{}]   `json:"metadata"`
	ModelMaxBudget       param.Field[interface{}]   `json:"model_max_budget"`
	ModelRpmLimit        param.Field[interface{}]   `json:"model_rpm_limit"`
	ModelTpmLimit        param.Field[interface{}]   `json:"model_tpm_limit"`
	Models               param.Field[[]interface{}] `json:"models"`
	Permissions          param.Field[interface{}]   `json:"permissions"`
	RpmLimit             param.Field[int64]         `json:"rpm_limit"`
	SendInviteEmail      param.Field[bool]          `json:"send_invite_email"`
	SoftBudget           param.Field[float64]       `json:"soft_budget"`
	Spend                param.Field[float64]       `json:"spend"`
	Tags                 param.Field[[]string]      `json:"tags"`
	TeamID               param.Field[string]        `json:"team_id"`
	TpmLimit             param.Field[int64]         `json:"tpm_limit"`
	UserID               param.Field[string]        `json:"user_id"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r KeyGenerateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeyRegenerateByKeyParams struct {
	RegenerateKeyRequest RegenerateKeyRequestParam `json:"regenerate_key_request"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r KeyRegenerateByKeyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RegenerateKeyRequest)
}

type KeyGetInfoParams struct {
	// Key in the request parameters
	Key param.Field[string] `query:"key"`
}

// URLQuery serializes [KeyGetInfoParams]'s query parameters as `url.Values`.
func (r KeyGetInfoParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type KeyUnblockParams struct {
	BlockKeyRequest BlockKeyRequestParam `json:"block_key_request,required"`
	// The llm-changed-by header enables tracking of actions performed by authorized
	// users on behalf of other users, providing an audit trail for accountability
	LlmChangedBy param.Field[string] `header:"llm-changed-by"`
}

func (r KeyUnblockParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BlockKeyRequest)
}
