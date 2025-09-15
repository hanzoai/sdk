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

func TestOrganizationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.New(context.TODO(), hanzoai.OrganizationNewParams{
		OrganizationAlias:   hanzoai.F("organization_alias"),
		BudgetDuration:      hanzoai.F("budget_duration"),
		BudgetID:            hanzoai.F("budget_id"),
		MaxBudget:           hanzoai.F(0.000000),
		MaxParallelRequests: hanzoai.F(int64(0)),
		Metadata:            hanzoai.F[any](map[string]interface{}{}),
		ModelMaxBudget:      hanzoai.F[any](map[string]interface{}{}),
		Models:              hanzoai.F([]interface{}{map[string]interface{}{}}),
		OrganizationID:      hanzoai.F("organization_id"),
		RpmLimit:            hanzoai.F(int64(0)),
		SoftBudget:          hanzoai.F(0.000000),
		TpmLimit:            hanzoai.F(int64(0)),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.Update(context.TODO(), hanzoai.OrganizationUpdateParams{
		BudgetID:          hanzoai.F("budget_id"),
		Metadata:          hanzoai.F[any](map[string]interface{}{}),
		Models:            hanzoai.F([]string{"string"}),
		OrganizationAlias: hanzoai.F("organization_alias"),
		OrganizationID:    hanzoai.F("organization_id"),
		Spend:             hanzoai.F(0.000000),
		UpdatedBy:         hanzoai.F("updated_by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationList(t *testing.T) {
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
	_, err := client.Organization.List(context.TODO())
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationDelete(t *testing.T) {
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
	_, err := client.Organization.Delete(context.TODO(), hanzoai.OrganizationDeleteParams{
		OrganizationIDs: hanzoai.F([]string{"string"}),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationAddMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.AddMember(context.TODO(), hanzoai.OrganizationAddMemberParams{
		Member: hanzoai.F[hanzoai.OrganizationAddMemberParamsMemberUnion](hanzoai.OrganizationAddMemberParamsMemberArray([]hanzoai.OrgMemberParam{{
			Role:      hanzoai.F(hanzoai.OrgMemberRoleOrgAdmin),
			UserEmail: hanzoai.F("user_email"),
			UserID:    hanzoai.F("user_id"),
		}})),
		OrganizationID:          hanzoai.F("organization_id"),
		MaxBudgetInOrganization: hanzoai.F(0.000000),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationDeleteMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.DeleteMember(context.TODO(), hanzoai.OrganizationDeleteMemberParams{
		OrganizationID: hanzoai.F("organization_id"),
		UserEmail:      hanzoai.F("user_email"),
		UserID:         hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrganizationUpdateMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Organization.UpdateMember(context.TODO(), hanzoai.OrganizationUpdateMemberParams{
		OrganizationID:          hanzoai.F("organization_id"),
		MaxBudgetInOrganization: hanzoai.F(0.000000),
		Role:                    hanzoai.F(hanzoai.OrganizationUpdateMemberParamsRoleProxyAdmin),
		UserEmail:               hanzoai.F("user_email"),
		UserID:                  hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
