// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hanzoai

import (
	"context"
	"net/http"

	"github.com/hanzoai/go-sdk/internal/requestconfig"
	"github.com/hanzoai/go-sdk/option"
)

// AudioSpeechService contains methods and other services that help with
// interacting with the Hanzo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioSpeechService] method instead.
type AudioSpeechService struct {
	Options []option.RequestOption
}

// NewAudioSpeechService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioSpeechService(opts ...option.RequestOption) (r *AudioSpeechService) {
	r = &AudioSpeechService{}
	r.Options = opts
	return
}

// Same params as:
//
// https://platform.openai.com/docs/api-reference/audio/createSpeech
func (r *AudioSpeechService) New(ctx context.Context, opts ...option.RequestOption) (res *AudioSpeechNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/audio/speech"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AudioSpeechNewResponse = interface{}
