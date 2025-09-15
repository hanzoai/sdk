// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/hanzoai/go-sdk"
	"github.com/hanzoai/go-sdk/internal/testutil"
	"github.com/hanzoai/go-sdk/option"
)

func TestKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Update(context.TODO(), hanzoai.KeyUpdateParams{
		Key:                  hanzoai.F("key"),
		Aliases:              hanzoai.F[any](map[string]interface{}{}),
		AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
		Blocked:              hanzoai.F(true),
		BudgetDuration:       hanzoai.F("budget_duration"),
		BudgetID:             hanzoai.F("budget_id"),
		Config:               hanzoai.F[any](map[string]interface{}{}),
		Duration:             hanzoai.F("duration"),
		EnforcedParams:       hanzoai.F([]string{"string"}),
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
		Spend:                hanzoai.F(0.000000),
		Tags:                 hanzoai.F([]string{"string"}),
		TeamID:               hanzoai.F("team_id"),
		TempBudgetExpiry:     hanzoai.F(time.Now()),
		TempBudgetIncrease:   hanzoai.F(0.000000),
		TpmLimit:             hanzoai.F(int64(0)),
		UserID:               hanzoai.F("user_id"),
		LlmChangedBy:         hanzoai.F("llm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyListWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.List(context.TODO(), hanzoai.KeyListParams{
		IncludeTeamKeys:  hanzoai.F(true),
		KeyAlias:         hanzoai.F("key_alias"),
		OrganizationID:   hanzoai.F("organization_id"),
		Page:             hanzoai.F(int64(1)),
		ReturnFullObject: hanzoai.F(true),
		Size:             hanzoai.F(int64(1)),
		TeamID:           hanzoai.F("team_id"),
		UserID:           hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Delete(context.TODO(), hanzoai.KeyDeleteParams{
		KeyAliases:   hanzoai.F([]string{"string"}),
		Keys:         hanzoai.F([]string{"string"}),
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

func TestKeyBlockWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Block(context.TODO(), hanzoai.KeyBlockParams{
		BlockKeyRequest: hanzoai.BlockKeyRequestParam{
			Key: hanzoai.F("key"),
		},
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

func TestKeyCheckHealth(t *testing.T) {
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
	_, err := client.Key.CheckHealth(context.TODO())
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Generate(context.TODO(), hanzoai.KeyGenerateParams{
		Aliases:              hanzoai.F[any](map[string]interface{}{}),
		AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
		Blocked:              hanzoai.F(true),
		BudgetDuration:       hanzoai.F("budget_duration"),
		BudgetID:             hanzoai.F("budget_id"),
		Config:               hanzoai.F[any](map[string]interface{}{}),
		Duration:             hanzoai.F("duration"),
		EnforcedParams:       hanzoai.F([]string{"string"}),
		Guardrails:           hanzoai.F([]string{"string"}),
		Key:                  hanzoai.F("key"),
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
		SoftBudget:           hanzoai.F(0.000000),
		Spend:                hanzoai.F(0.000000),
		Tags:                 hanzoai.F([]string{"string"}),
		TeamID:               hanzoai.F("team_id"),
		TpmLimit:             hanzoai.F(int64(0)),
		UserID:               hanzoai.F("user_id"),
		LlmChangedBy:         hanzoai.F("llm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyRegenerateByKeyWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.RegenerateByKey(
		context.TODO(),
		"key",
		hanzoai.KeyRegenerateByKeyParams{
			RegenerateKeyRequest: hanzoai.RegenerateKeyRequestParam{
				Aliases:              hanzoai.F[any](map[string]interface{}{}),
				AllowedCacheControls: hanzoai.F([]interface{}{map[string]interface{}{}}),
				Blocked:              hanzoai.F(true),
				BudgetDuration:       hanzoai.F("budget_duration"),
				BudgetID:             hanzoai.F("budget_id"),
				Config:               hanzoai.F[any](map[string]interface{}{}),
				Duration:             hanzoai.F("duration"),
				EnforcedParams:       hanzoai.F([]string{"string"}),
				Guardrails:           hanzoai.F([]string{"string"}),
				Key:                  hanzoai.F("key"),
				KeyAlias:             hanzoai.F("key_alias"),
				MaxBudget:            hanzoai.F(0.000000),
				MaxParallelRequests:  hanzoai.F(int64(0)),
				Metadata:             hanzoai.F[any](map[string]interface{}{}),
				ModelMaxBudget:       hanzoai.F[any](map[string]interface{}{}),
				ModelRpmLimit:        hanzoai.F[any](map[string]interface{}{}),
				ModelTpmLimit:        hanzoai.F[any](map[string]interface{}{}),
				Models:               hanzoai.F([]interface{}{map[string]interface{}{}}),
				NewMasterKey:         hanzoai.F("new_master_key"),
				Permissions:          hanzoai.F[any](map[string]interface{}{}),
				RpmLimit:             hanzoai.F(int64(0)),
				SendInviteEmail:      hanzoai.F(true),
				SoftBudget:           hanzoai.F(0.000000),
				Spend:                hanzoai.F(0.000000),
				Tags:                 hanzoai.F([]string{"string"}),
				TeamID:               hanzoai.F("team_id"),
				TpmLimit:             hanzoai.F(int64(0)),
				UserID:               hanzoai.F("user_id"),
			},
			LlmChangedBy: hanzoai.F("llm-changed-by"),
		},
	)
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyGetInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.GetInfo(context.TODO(), hanzoai.KeyGetInfoParams{
		Key: hanzoai.F("key"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestKeyUnblockWithOptionalParams(t *testing.T) {
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
	_, err := client.Key.Unblock(context.TODO(), hanzoai.KeyUnblockParams{
		BlockKeyRequest: hanzoai.BlockKeyRequestParam{
			Key: hanzoai.F("key"),
		},
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
