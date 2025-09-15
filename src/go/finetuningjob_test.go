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
	"github.com/hanzoai/go-sdk/shared"
)

func TestFineTuningJobNewWithOptionalParams(t *testing.T) {
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
	_, err := client.FineTuning.Jobs.New(context.TODO(), hanzoai.FineTuningJobNewParams{
		CustomLlmProvider: hanzoai.F(hanzoai.FineTuningJobNewParamsCustomLlmProviderOpenAI),
		Model:             hanzoai.F("model"),
		TrainingFile:      hanzoai.F("training_file"),
		Hyperparameters: hanzoai.F(hanzoai.FineTuningJobNewParamsHyperparameters{
			BatchSize:              hanzoai.F[hanzoai.FineTuningJobNewParamsHyperparametersBatchSizeUnion](shared.UnionString("string")),
			LearningRateMultiplier: hanzoai.F[hanzoai.FineTuningJobNewParamsHyperparametersLearningRateMultiplierUnion](shared.UnionString("string")),
			NEpochs:                hanzoai.F[hanzoai.FineTuningJobNewParamsHyperparametersNEpochsUnion](shared.UnionString("string")),
		}),
		Integrations:   hanzoai.F([]string{"string"}),
		Seed:           hanzoai.F(int64(0)),
		Suffix:         hanzoai.F("suffix"),
		ValidationFile: hanzoai.F("validation_file"),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFineTuningJobGet(t *testing.T) {
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
	_, err := client.FineTuning.Jobs.Get(
		context.TODO(),
		"fine_tuning_job_id",
		hanzoai.FineTuningJobGetParams{
			CustomLlmProvider: hanzoai.F(hanzoai.FineTuningJobGetParamsCustomLlmProviderOpenAI),
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

func TestFineTuningJobListWithOptionalParams(t *testing.T) {
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
	_, err := client.FineTuning.Jobs.List(context.TODO(), hanzoai.FineTuningJobListParams{
		CustomLlmProvider: hanzoai.F(hanzoai.FineTuningJobListParamsCustomLlmProviderOpenAI),
		After:             hanzoai.F("after"),
		Limit:             hanzoai.F(int64(0)),
	})
	if err != nil {
		var apierr *hanzoai.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
