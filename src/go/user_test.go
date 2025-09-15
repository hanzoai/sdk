// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/internal/testutil"
	"github.com/hanzoai/go-sdk/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.New(context.TODO(), hanzoai.UserNewParams{
		Aliases:              hanzoai.F[any](map[string]interface{}{}),
		AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
		AutoCreateKey:        hanzoai.F(true),
		Blocked:              hanzoai.F(true),
		BudgetDuration:       hanzoai.F("budget_duration"),
		Config:               hanzoai.F[any](map[string]interface{}{}),
		Duration:             hanzoai.F("duration"),
		Guardrails:           hanzoai.F([]string{"string"}),
		KeyAlias:             hanzoai.F("key_alias"),
		MaxBudget:            hanzoai.F(0.000000),
		MaxParallelRequests:  hanzoai.F(int64(0)),
		Metadata:             hanzoai.F[any](map[string]interface{}{}),
		ModelMaxBudget:       hanzoai.F[any](map[string]interface{}{}),
		ModelRpmLimit:        hanzoai.F[any](map[string]interface{}{}),
		ModelTpmLimit:        hanzoai.F[any](map[string]interface{}{}),
		Models:               hanzoai.F([]interface{}{map[string]interface{}{}}),
		Permissions:          hanzoai.F[any](map[string]interface{}{}),
		RpmLimit:             hanzoai.F(int64(0)),
		SendInviteEmail:      hanzoai.F(true),
		Spend:                hanzoai.F(0.000000),
		TeamID:               hanzoai.F("team_id"),
		Teams:                hanzoai.F([]interface{}{map[string]interface{}{}}),
		TpmLimit:             hanzoai.F(int64(0)),
		UserAlias:            hanzoai.F("user_alias"),
		UserEmail:            hanzoai.F("user_email"),
		UserID:               hanzoai.F("user_id"),
		UserRole:             hanzoai.F(hanzoai.UserNewParamsUserRoleProxyAdmin),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.Update(context.TODO(), hanzoai.UserUpdateParams{
		Aliases:              hanzoai.F[any](map[string]interface{}{}),
		AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
		Blocked:              hanzoai.F(true),
		BudgetDuration:       hanzoai.F("budget_duration"),
		Config:               hanzoai.F[any](map[string]interface{}{}),
		Duration:             hanzoai.F("duration"),
		Guardrails:           hanzoai.F([]string{"string"}),
		KeyAlias:             hanzoai.F("key_alias"),
		MaxBudget:            hanzoai.F(0.000000),
		MaxParallelRequests:  hanzoai.F(int64(0)),
		Metadata:             hanzoai.F[any](map[string]interface{}{}),
		ModelMaxBudget:       hanzoai.F[any](map[string]interface{}{}),
		ModelRpmLimit:        hanzoai.F[any](map[string]interface{}{}),
		ModelTpmLimit:        hanzoai.F[any](map[string]interface{}{}),
		Models:               hanzoai.F([]interface{}{map[string]interface{}{}}),
		Password:             hanzoai.F("password"),
		Permissions:          hanzoai.F[any](map[string]interface{}{}),
		RpmLimit:             hanzoai.F(int64(0)),
		Spend:                hanzoai.F(0.000000),
		TeamID:               hanzoai.F("team_id"),
		TpmLimit:             hanzoai.F(int64(0)),
		UserEmail:            hanzoai.F("user_email"),
		UserID:               hanzoai.F("user_id"),
		UserRole:             hanzoai.F(hanzoai.UserUpdateParamsUserRoleProxyAdmin),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserListWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.List(context.TODO(), hanzoai.UserListParams{
		Page:     hanzoai.F(int64(1)),
		PageSize: hanzoai.F(int64(1)),
		Role:     hanzoai.F("role"),
		UserIDs:  hanzoai.F("user_ids"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserDeleteWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.Delete(context.TODO(), hanzoai.UserDeleteParams{
		UserIDs:      hanzoai.F([]string{"string"}),
		LlmChangedBy: hanzoai.F("llm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserGetInfoWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hanzoai.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.User.GetInfo(context.TODO(), hanzoai.UserGetInfoParams{
		UserID: hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
