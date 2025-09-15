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

func TestTeamNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.New(context.TODO(), hanzoai.TeamNewParams{
		Admins:         hanzoai.F([]interface{}{map[string]interface{}{}}),
		Blocked:        hanzoai.F(true),
		BudgetDuration: hanzoai.F("budget_duration"),
		Guardrails:     hanzoai.F([]string{"string"}),
		MaxBudget:      hanzoai.F(0.000000),
		Members:        hanzoai.F([]interface{}{map[string]interface{}{}}),
		MembersWithRoles: hanzoai.F([]hanzoai.MemberParam{{
			Role:      hanzoai.F(hanzoai.MemberRoleAdmin),
			UserEmail: hanzoai.F("user_email"),
			UserID:    hanzoai.F("user_id"),
		}}),
		Metadata:       hanzoai.F[any](map[string]interface{}{}),
		ModelAliases:   hanzoai.F[any](map[string]interface{}{}),
		Models:         hanzoai.F([]interface{}{map[string]interface{}{}}),
		OrganizationID: hanzoai.F("organization_id"),
		RpmLimit:       hanzoai.F(int64(0)),
		Tags:           hanzoai.F([]interface{}{map[string]interface{}{}}),
		TeamAlias:      hanzoai.F("team_alias"),
		TeamID:         hanzoai.F("team_id"),
		TpmLimit:       hanzoai.F(int64(0)),
		LlmChangedBy:   hanzoai.F("llm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.Update(context.TODO(), hanzoai.TeamUpdateParams{
		TeamID:         hanzoai.F("team_id"),
		Blocked:        hanzoai.F(true),
		BudgetDuration: hanzoai.F("budget_duration"),
		Guardrails:     hanzoai.F([]string{"string"}),
		MaxBudget:      hanzoai.F(0.000000),
		Metadata:       hanzoai.F[any](map[string]interface{}{}),
		ModelAliases:   hanzoai.F[any](map[string]interface{}{}),
		Models:         hanzoai.F([]interface{}{map[string]interface{}{}}),
		OrganizationID: hanzoai.F("organization_id"),
		RpmLimit:       hanzoai.F(int64(0)),
		Tags:           hanzoai.F([]interface{}{map[string]interface{}{}}),
		TeamAlias:      hanzoai.F("team_alias"),
		TpmLimit:       hanzoai.F(int64(0)),
		LlmChangedBy:   hanzoai.F("llm-changed-by"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamListWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.List(context.TODO(), hanzoai.TeamListParams{
		OrganizationID: hanzoai.F("organization_id"),
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

func TestTeamDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.Delete(context.TODO(), hanzoai.TeamDeleteParams{
		TeamIDs:      hanzoai.F([]string{"string"}),
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

func TestTeamAddMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.AddMember(context.TODO(), hanzoai.TeamAddMemberParams{
		Member: hanzoai.F[hanzoai.TeamAddMemberParamsMemberUnion](hanzoai.TeamAddMemberParamsMemberArray([]hanzoai.MemberParam{{
			Role:      hanzoai.F(hanzoai.MemberRoleAdmin),
			UserEmail: hanzoai.F("user_email"),
			UserID:    hanzoai.F("user_id"),
		}})),
		TeamID:          hanzoai.F("team_id"),
		MaxBudgetInTeam: hanzoai.F(0.000000),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamBlock(t *testing.T) {
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
	_, err := client.Team.Block(context.TODO(), hanzoai.TeamBlockParams{
		BlockTeamRequest: hanzoai.BlockTeamRequestParam{
			TeamID: hanzoai.F("team_id"),
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

func TestTeamDisableLogging(t *testing.T) {
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
	_, err := client.Team.DisableLogging(context.TODO(), "team_id")
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamListAvailableWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.ListAvailable(context.TODO(), hanzoai.TeamListAvailableParams{
		ResponseModel: hanzoai.F[any](map[string]interface{}{}),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamRemoveMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.RemoveMember(context.TODO(), hanzoai.TeamRemoveMemberParams{
		TeamID:    hanzoai.F("team_id"),
		UserEmail: hanzoai.F("user_email"),
		UserID:    hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamGetInfoWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.GetInfo(context.TODO(), hanzoai.TeamGetInfoParams{
		TeamID: hanzoai.F("team_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTeamUnblock(t *testing.T) {
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
	_, err := client.Team.Unblock(context.TODO(), hanzoai.TeamUnblockParams{
		BlockTeamRequest: hanzoai.BlockTeamRequestParam{
			TeamID: hanzoai.F("team_id"),
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

func TestTeamUpdateMemberWithOptionalParams(t *testing.T) {
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
	_, err := client.Team.UpdateMember(context.TODO(), hanzoai.TeamUpdateMemberParams{
		TeamID:          hanzoai.F("team_id"),
		MaxBudgetInTeam: hanzoai.F(0.000000),
		Role:            hanzoai.F(hanzoai.TeamUpdateMemberParamsRoleAdmin),
		UserEmail:       hanzoai.F("user_email"),
		UserID:          hanzoai.F("user_id"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
