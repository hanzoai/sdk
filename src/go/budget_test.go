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

func TestBudgetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Budget.New(context.TODO(), hanzoai.BudgetNewParams{
		BudgetNew: hanzoai.BudgetNewParam{
			BudgetDuration:      hanzoai.F("budget_duration"),
			BudgetID:            hanzoai.F("budget_id"),
			MaxBudget:           hanzoai.F(0.000000),
			MaxParallelRequests: hanzoai.F(int64(0)),
			ModelMaxBudget: hanzoai.F(map[string]hanzoai.BudgetNewModelMaxBudgetParam{
				"foo": {
					BudgetDuration: hanzoai.F("budget_duration"),
					MaxBudget:      hanzoai.F(0.000000),
					RpmLimit:       hanzoai.F(int64(0)),
					TpmLimit:       hanzoai.F(int64(0)),
				},
			}),
			RpmLimit:   hanzoai.F(int64(0)),
			SoftBudget: hanzoai.F(0.000000),
			TpmLimit:   hanzoai.F(int64(0)),
		},
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Budget.Update(context.TODO(), hanzoai.BudgetUpdateParams{
		BudgetNew: hanzoai.BudgetNewParam{
			BudgetDuration:      hanzoai.F("budget_duration"),
			BudgetID:            hanzoai.F("budget_id"),
			MaxBudget:           hanzoai.F(0.000000),
			MaxParallelRequests: hanzoai.F(int64(0)),
			ModelMaxBudget: hanzoai.F(map[string]hanzoai.BudgetNewModelMaxBudgetParam{
				"foo": {
					BudgetDuration: hanzoai.F("budget_duration"),
					MaxBudget:      hanzoai.F(0.000000),
					RpmLimit:       hanzoai.F(int64(0)),
					TpmLimit:       hanzoai.F(int64(0)),
				},
			}),
			RpmLimit:   hanzoai.F(int64(0)),
			SoftBudget: hanzoai.F(0.000000),
			TpmLimit:   hanzoai.F(int64(0)),
		},
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetList(t *testing.T) {
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
	_, err := client.Budget.List(context.TODO())
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetDelete(t *testing.T) {
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
	_, err := client.Budget.Delete(context.TODO(), hanzoai.BudgetDeleteParams{
		ID: hanzoai.F("id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetInfo(t *testing.T) {
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
	_, err := client.Budget.Info(context.TODO(), hanzoai.BudgetInfoParams{
		Budgets: hanzoai.F([]string{"string"}),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestBudgetSettings(t *testing.T) {
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
	_, err := client.Budget.Settings(context.TODO(), hanzoai.BudgetSettingsParams{
		BudgetID: hanzoai.F("budget_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
