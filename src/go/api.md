# hanzoai

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GetHomeResponse">GetHomeResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HanzoaiService.GetHome">GetHome</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GetHomeResponse">GetHomeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelListResponse">ModelListResponse</a>

Methods:

- <code title="get /v1/models">client.Models.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelListParams">ModelListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelListResponse">ModelListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# OpenAI

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAINewResponse">OpenAINewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIGetResponse">OpenAIGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIUpdateResponse">OpenAIUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeleteResponse">OpenAIDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIPatchResponse">OpenAIPatchResponse</a>

Methods:

- <code title="post /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAINewResponse">OpenAINewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIGetResponse">OpenAIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIUpdateResponse">OpenAIUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeleteResponse">OpenAIDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /openai/{endpoint}">client.OpenAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIPatchResponse">OpenAIPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Deployments

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentCompleteResponse">OpenAIDeploymentCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentEmbedResponse">OpenAIDeploymentEmbedResponse</a>

Methods:

- <code title="post /openai/deployments/{model}/completions">client.OpenAI.Deployments.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentCompleteResponse">OpenAIDeploymentCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /openai/deployments/{model}/embeddings">client.OpenAI.Deployments.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentService.Embed">Embed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentEmbedResponse">OpenAIDeploymentEmbedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentChatCompleteResponse">OpenAIDeploymentChatCompleteResponse</a>

Methods:

- <code title="post /openai/deployments/{model}/chat/completions">client.OpenAI.Deployments.Chat.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentChatService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OpenAIDeploymentChatCompleteResponse">OpenAIDeploymentChatCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Engines

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineCompleteResponse">EngineCompleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineEmbedResponse">EngineEmbedResponse</a>

Methods:

- <code title="post /engines/{model}/completions">client.Engines.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineCompleteResponse">EngineCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /engines/{model}/embeddings">client.Engines.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineService.Embed">Embed</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineEmbedResponse">EngineEmbedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineChatCompleteResponse">EngineChatCompleteResponse</a>

Methods:

- <code title="post /engines/{model}/chat/completions">client.Engines.Chat.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineChatService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, model <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EngineChatCompleteResponse">EngineChatCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Chat

## Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ChatCompletionNewResponse">ChatCompletionNewResponse</a>

Methods:

- <code title="post /v1/chat/completions">client.Chat.Completions.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ChatCompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ChatCompletionNewParams">ChatCompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ChatCompletionNewResponse">ChatCompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Completions

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CompletionNewResponse">CompletionNewResponse</a>

Methods:

- <code title="post /completions">client.Completions.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CompletionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CompletionNewParams">CompletionNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CompletionNewResponse">CompletionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Embeddings

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EmbeddingNewResponse">EmbeddingNewResponse</a>

Methods:

- <code title="post /embeddings">client.Embeddings.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EmbeddingService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EmbeddingNewParams">EmbeddingNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EmbeddingNewResponse">EmbeddingNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

## Generations

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ImageGenerationNewResponse">ImageGenerationNewResponse</a>

Methods:

- <code title="post /v1/images/generations">client.Images.Generations.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ImageGenerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ImageGenerationNewResponse">ImageGenerationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Audio

## Speech

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioSpeechNewResponse">AudioSpeechNewResponse</a>

Methods:

- <code title="post /v1/audio/speech">client.Audio.Speech.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioSpeechService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioSpeechNewResponse">AudioSpeechNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transcriptions

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioTranscriptionNewResponse">AudioTranscriptionNewResponse</a>

Methods:

- <code title="post /v1/audio/transcriptions">client.Audio.Transcriptions.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioTranscriptionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioTranscriptionNewParams">AudioTranscriptionNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AudioTranscriptionNewResponse">AudioTranscriptionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assistants

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantNewResponse">AssistantNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantListResponse">AssistantListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantDeleteResponse">AssistantDeleteResponse</a>

Methods:

- <code title="post /v1/assistants">client.Assistants.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantNewResponse">AssistantNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/assistants">client.Assistants.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantListResponse">AssistantListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/assistants/{assistant_id}">client.Assistants.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, assistantID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssistantDeleteResponse">AssistantDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Threads

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadNewResponse">ThreadNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadGetResponse">ThreadGetResponse</a>

Methods:

- <code title="post /v1/threads">client.Threads.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadNewResponse">ThreadNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/threads/{thread_id}">client.Threads.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadGetResponse">ThreadGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Messages

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadMessageNewResponse">ThreadMessageNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadMessageListResponse">ThreadMessageListResponse</a>

Methods:

- <code title="post /v1/threads/{thread_id}/messages">client.Threads.Messages.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadMessageService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadMessageNewResponse">ThreadMessageNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/threads/{thread_id}/messages">client.Threads.Messages.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadMessageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadMessageListResponse">ThreadMessageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Runs

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadRunNewResponse">ThreadRunNewResponse</a>

Methods:

- <code title="post /v1/threads/{thread_id}/runs">client.Threads.Runs.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadRunService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, threadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ThreadRunNewResponse">ThreadRunNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Moderations

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModerationNewResponse">ModerationNewResponse</a>

Methods:

- <code title="post /v1/moderations">client.Moderations.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModerationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModerationNewResponse">ModerationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Utils

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilGetSupportedOpenAIParamsResponse">UtilGetSupportedOpenAIParamsResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilTokenCounterResponse">UtilTokenCounterResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilTransformRequestResponse">UtilTransformRequestResponse</a>

Methods:

- <code title="get /utils/supported_openai_params">client.Utils.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilService.GetSupportedOpenAIParams">GetSupportedOpenAIParams</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilGetSupportedOpenAIParamsParams">UtilGetSupportedOpenAIParamsParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilGetSupportedOpenAIParamsResponse">UtilGetSupportedOpenAIParamsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /utils/token_counter">client.Utils.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilService.TokenCounter">TokenCounter</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilTokenCounterParams">UtilTokenCounterParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilTokenCounterResponse">UtilTokenCounterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /utils/transform_request">client.Utils.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilService.TransformRequest">TransformRequest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilTransformRequestParams">UtilTransformRequestParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UtilTransformRequestResponse">UtilTransformRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Model

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigurableClientsideParamsCustomAuth">ConfigurableClientsideParamsCustomAuth</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelInfoParam">ModelInfoParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelNewResponse">ModelNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelDeleteResponse">ModelDeleteResponse</a>

Methods:

- <code title="post /model/new">client.Model.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelNewParams">ModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelNewResponse">ModelNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /model/delete">client.Model.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelDeleteParams">ModelDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelDeleteResponse">ModelDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelInfoListResponse">ModelInfoListResponse</a>

Methods:

- <code title="get /model/info">client.Model.Info.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelInfoService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelInfoListParams">ModelInfoListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelInfoListResponse">ModelInfoListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Update

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UpdateDeploymentParam">UpdateDeploymentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdateFullResponse">ModelUpdateFullResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdatePartialResponse">ModelUpdatePartialResponse</a>

Methods:

- <code title="post /model/update">client.Model.Update.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdateService.Full">Full</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdateFullParams">ModelUpdateFullParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdateFullResponse">ModelUpdateFullResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /model/{model_id}/update">client.Model.Update.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdateService.Partial">Partial</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdatePartialParams">ModelUpdatePartialParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelUpdatePartialResponse">ModelUpdatePartialResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ModelGroup

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelGroupGetInfoResponse">ModelGroupGetInfoResponse</a>

Methods:

- <code title="get /model_group/info">client.ModelGroup.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelGroupService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelGroupGetInfoParams">ModelGroupGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ModelGroupGetInfoResponse">ModelGroupGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Routes

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RouteListResponse">RouteListResponse</a>

Methods:

- <code title="get /routes">client.Routes.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RouteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RouteListResponse">RouteListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Responses

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseNewResponse">ResponseNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseGetResponse">ResponseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseDeleteResponse">ResponseDeleteResponse</a>

Methods:

- <code title="post /v1/responses">client.Responses.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseNewResponse">ResponseNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseGetResponse">ResponseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v1/responses/{response_id}">client.Responses.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseDeleteResponse">ResponseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## InputItems

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseInputItemListResponse">ResponseInputItemListResponse</a>

Methods:

- <code title="get /v1/responses/{response_id}/input_items">client.Responses.InputItems.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseInputItemService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, responseID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ResponseInputItemListResponse">ResponseInputItemListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Batches

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchNewResponse">BatchNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchGetResponse">BatchGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchListResponse">BatchListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchCancelWithProviderResponse">BatchCancelWithProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchNewWithProviderResponse">BatchNewWithProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchListWithProviderResponse">BatchListWithProviderResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchGetWithProviderResponse">BatchGetWithProviderResponse</a>

Methods:

- <code title="post /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchNewParams">BatchNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchNewResponse">BatchNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches/{batch_id}">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchGetParams">BatchGetParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchGetResponse">BatchGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchListParams">BatchListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchListResponse">BatchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{provider}/v1/batches/{batch_id}/cancel">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.CancelWithProvider">CancelWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchCancelWithProviderResponse">BatchCancelWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /{provider}/v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.NewWithProvider">NewWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchNewWithProviderResponse">BatchNewWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/batches">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.ListWithProvider">ListWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchListWithProviderParams">BatchListWithProviderParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchListWithProviderResponse">BatchListWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/batches/{batch_id}">client.Batches.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchService.GetWithProvider">GetWithProvider</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchGetWithProviderResponse">BatchGetWithProviderResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cancel

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchCancelCancelResponse">BatchCancelCancelResponse</a>

Methods:

- <code title="post /batches/{batch_id}/cancel">client.Batches.Cancel.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchCancelService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, batchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchCancelCancelParams">BatchCancelCancelParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BatchCancelCancelResponse">BatchCancelCancelResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Rerank

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankNewResponse">RerankNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankNewV1Response">RerankNewV1Response</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankNewV2Response">RerankNewV2Response</a>

Methods:

- <code title="post /rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankNewResponse">RerankNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankService.NewV1">NewV1</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankNewV1Response">RerankNewV1Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/rerank">client.Rerank.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankService.NewV2">NewV2</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RerankNewV2Response">RerankNewV2Response</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# FineTuning

## Jobs

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobNewResponse">FineTuningJobNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobGetResponse">FineTuningJobGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobListResponse">FineTuningJobListResponse</a>

Methods:

- <code title="post /v1/fine_tuning/jobs">client.FineTuning.Jobs.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobNewParams">FineTuningJobNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobNewResponse">FineTuningJobNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/fine_tuning/jobs/{fine_tuning_job_id}">client.FineTuning.Jobs.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fineTuningJobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobGetParams">FineTuningJobGetParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobGetResponse">FineTuningJobGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/fine_tuning/jobs">client.FineTuning.Jobs.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobListParams">FineTuningJobListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobListResponse">FineTuningJobListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Cancel

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobCancelNewResponse">FineTuningJobCancelNewResponse</a>

Methods:

- <code title="post /v1/fine_tuning/jobs/{fine_tuning_job_id}/cancel">client.FineTuning.Jobs.Cancel.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobCancelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fineTuningJobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FineTuningJobCancelNewResponse">FineTuningJobCancelNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Credentials

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialNewResponse">CredentialNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialListResponse">CredentialListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialDeleteResponse">CredentialDeleteResponse</a>

Methods:

- <code title="post /credentials">client.Credentials.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialNewParams">CredentialNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialNewResponse">CredentialNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /credentials">client.Credentials.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialListResponse">CredentialListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /credentials/{credential_name}">client.Credentials.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, credentialName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CredentialDeleteResponse">CredentialDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VertexAI

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAINewResponse">VertexAINewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIGetResponse">VertexAIGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIUpdateResponse">VertexAIUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIDeleteResponse">VertexAIDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIPatchResponse">VertexAIPatchResponse</a>

Methods:

- <code title="post /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAINewResponse">VertexAINewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIGetResponse">VertexAIGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIUpdateResponse">VertexAIUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIDeleteResponse">VertexAIDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /vertex_ai/{endpoint}">client.VertexAI.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#VertexAIPatchResponse">VertexAIPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Gemini

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiNewResponse">GeminiNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiGetResponse">GeminiGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiUpdateResponse">GeminiUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiDeleteResponse">GeminiDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiPatchResponse">GeminiPatchResponse</a>

Methods:

- <code title="post /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiNewResponse">GeminiNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiGetResponse">GeminiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiUpdateResponse">GeminiUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiDeleteResponse">GeminiDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /gemini/{endpoint}">client.Gemini.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GeminiPatchResponse">GeminiPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cohere

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereNewResponse">CohereNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereGetResponse">CohereGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereUpdateResponse">CohereUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereDeleteResponse">CohereDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereModifyResponse">CohereModifyResponse</a>

Methods:

- <code title="post /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereNewResponse">CohereNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereGetResponse">CohereGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereUpdateResponse">CohereUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereDeleteResponse">CohereDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /cohere/{endpoint}">client.Cohere.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereService.Modify">Modify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CohereModifyResponse">CohereModifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Anthropic

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicNewResponse">AnthropicNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicGetResponse">AnthropicGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicUpdateResponse">AnthropicUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicDeleteResponse">AnthropicDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicModifyResponse">AnthropicModifyResponse</a>

Methods:

- <code title="post /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicNewResponse">AnthropicNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicGetResponse">AnthropicGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicUpdateResponse">AnthropicUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicDeleteResponse">AnthropicDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /anthropic/{endpoint}">client.Anthropic.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicService.Modify">Modify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AnthropicModifyResponse">AnthropicModifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Bedrock

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockNewResponse">BedrockNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockGetResponse">BedrockGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockUpdateResponse">BedrockUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockDeleteResponse">BedrockDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockPatchResponse">BedrockPatchResponse</a>

Methods:

- <code title="post /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockNewResponse">BedrockNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockGetResponse">BedrockGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockUpdateResponse">BedrockUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockDeleteResponse">BedrockDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /bedrock/{endpoint}">client.Bedrock.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BedrockPatchResponse">BedrockPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# EuAssemblyai

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiNewResponse">EuAssemblyaiNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiGetResponse">EuAssemblyaiGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiUpdateResponse">EuAssemblyaiUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiDeleteResponse">EuAssemblyaiDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiPatchResponse">EuAssemblyaiPatchResponse</a>

Methods:

- <code title="post /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiNewResponse">EuAssemblyaiNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiGetResponse">EuAssemblyaiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiUpdateResponse">EuAssemblyaiUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiDeleteResponse">EuAssemblyaiDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /eu.assemblyai/{endpoint}">client.EuAssemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#EuAssemblyaiPatchResponse">EuAssemblyaiPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Assemblyai

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiNewResponse">AssemblyaiNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiGetResponse">AssemblyaiGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiUpdateResponse">AssemblyaiUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiDeleteResponse">AssemblyaiDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiPatchResponse">AssemblyaiPatchResponse</a>

Methods:

- <code title="post /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiNewResponse">AssemblyaiNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiGetResponse">AssemblyaiGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiUpdateResponse">AssemblyaiUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiDeleteResponse">AssemblyaiDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /assemblyai/{endpoint}">client.Assemblyai.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AssemblyaiPatchResponse">AssemblyaiPatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Azure

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureNewResponse">AzureNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureUpdateResponse">AzureUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureDeleteResponse">AzureDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureCallResponse">AzureCallResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzurePatchResponse">AzurePatchResponse</a>

Methods:

- <code title="post /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureNewResponse">AzureNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureUpdateResponse">AzureUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureDeleteResponse">AzureDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureService.Call">Call</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureCallResponse">AzureCallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /azure/{endpoint}">client.Azure.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzureService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AzurePatchResponse">AzurePatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Langfuse

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseNewResponse">LangfuseNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseGetResponse">LangfuseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseUpdateResponse">LangfuseUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseDeleteResponse">LangfuseDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfusePatchResponse">LangfusePatchResponse</a>

Methods:

- <code title="post /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseNewResponse">LangfuseNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseGetResponse">LangfuseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseUpdateResponse">LangfuseUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseDeleteResponse">LangfuseDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /langfuse/{endpoint}">client.Langfuse.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfuseService.Patch">Patch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpoint <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#LangfusePatchResponse">LangfusePatchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Config

## PassThroughEndpoint

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#PassThroughGenericEndpointParam">PassThroughGenericEndpointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#PassThroughEndpointResponse">PassThroughEndpointResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#PassThroughGenericEndpoint">PassThroughGenericEndpoint</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointNewResponse">ConfigPassThroughEndpointNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointUpdateResponse">ConfigPassThroughEndpointUpdateResponse</a>

Methods:

- <code title="post /config/pass_through_endpoint">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointNewParams">ConfigPassThroughEndpointNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointNewResponse">ConfigPassThroughEndpointNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /config/pass_through_endpoint/{endpoint_id}">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, endpointID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointUpdateResponse">ConfigPassThroughEndpointUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /config/pass_through_endpoint">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointListParams">ConfigPassThroughEndpointListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#PassThroughEndpointResponse">PassThroughEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /config/pass_through_endpoint">client.Config.PassThroughEndpoint.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ConfigPassThroughEndpointDeleteParams">ConfigPassThroughEndpointDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#PassThroughEndpointResponse">PassThroughEndpointResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Test

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TestPingResponse">TestPingResponse</a>

Methods:

- <code title="get /test">client.Test.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TestService.Ping">Ping</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TestPingResponse">TestPingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckAllResponse">HealthCheckAllResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckLivelinessResponse">HealthCheckLivelinessResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckLivenessResponse">HealthCheckLivenessResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckReadinessResponse">HealthCheckReadinessResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckServicesResponse">HealthCheckServicesResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthService.CheckAll">CheckAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckAllParams">HealthCheckAllParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckAllResponse">HealthCheckAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/liveliness">client.Health.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthService.CheckLiveliness">CheckLiveliness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckLivelinessResponse">HealthCheckLivelinessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/liveness">client.Health.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthService.CheckLiveness">CheckLiveness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckLivenessResponse">HealthCheckLivenessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/readiness">client.Health.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthService.CheckReadiness">CheckReadiness</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckReadinessResponse">HealthCheckReadinessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /health/services">client.Health.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthService.CheckServices">CheckServices</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckServicesParams">HealthCheckServicesParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#HealthCheckServicesResponse">HealthCheckServicesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Active

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ActiveListCallbacksResponse">ActiveListCallbacksResponse</a>

Methods:

- <code title="get /active/callbacks">client.Active.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ActiveService.ListCallbacks">ListCallbacks</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ActiveListCallbacksResponse">ActiveListCallbacksResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SettingGetResponse">SettingGetResponse</a>

Methods:

- <code title="get /settings">client.Settings.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SettingGetResponse">SettingGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Key

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BlockKeyRequestParam">BlockKeyRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GenerateKeyResponse">GenerateKeyResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyUpdateResponse">KeyUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyListResponse">KeyListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyDeleteResponse">KeyDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyBlockResponse">KeyBlockResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyCheckHealthResponse">KeyCheckHealthResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyGetInfoResponse">KeyGetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyUnblockResponse">KeyUnblockResponse</a>

Methods:

- <code title="post /key/update">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyUpdateParams">KeyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyUpdateResponse">KeyUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /key/list">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyListParams">KeyListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyListResponse">KeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/delete">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyDeleteParams">KeyDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyDeleteResponse">KeyDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/block">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.Block">Block</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyBlockParams">KeyBlockParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyBlockResponse">KeyBlockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/health">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.CheckHealth">CheckHealth</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyCheckHealthResponse">KeyCheckHealthResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/generate">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyGenerateParams">KeyGenerateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GenerateKeyResponse">GenerateKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/{key}/regenerate">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.RegenerateByKey">RegenerateByKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, key <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyRegenerateByKeyParams">KeyRegenerateByKeyParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GenerateKeyResponse">GenerateKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /key/info">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyGetInfoParams">KeyGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyGetInfoResponse">KeyGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /key/unblock">client.Key.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyService.Unblock">Unblock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyUnblockParams">KeyUnblockParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#KeyUnblockResponse">KeyUnblockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Regenerate

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#RegenerateKeyRequestParam">RegenerateKeyRequestParam</a>

# User

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserNewResponse">UserNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserUpdateResponse">UserUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserListResponse">UserListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserDeleteResponse">UserDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserGetInfoResponse">UserGetInfoResponse</a>

Methods:

- <code title="post /user/new">client.User.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserNewParams">UserNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserNewResponse">UserNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /user/update">client.User.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserUpdateParams">UserUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserUpdateResponse">UserUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/get_users">client.User.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserListParams">UserListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserListResponse">UserListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /user/delete">client.User.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserDeleteParams">UserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserDeleteResponse">UserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /user/info">client.User.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserGetInfoParams">UserGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#UserGetInfoResponse">UserGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Team

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BlockTeamRequestParam">BlockTeamRequestParam</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#MemberParam">MemberParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#Member">Member</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamNewResponse">TeamNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUpdateResponse">TeamUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamListResponse">TeamListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamDeleteResponse">TeamDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamAddMemberResponse">TeamAddMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamBlockResponse">TeamBlockResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamDisableLoggingResponse">TeamDisableLoggingResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamListAvailableResponse">TeamListAvailableResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamRemoveMemberResponse">TeamRemoveMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamGetInfoResponse">TeamGetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUnblockResponse">TeamUnblockResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUpdateMemberResponse">TeamUpdateMemberResponse</a>

Methods:

- <code title="post /team/new">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamNewParams">TeamNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamNewResponse">TeamNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/update">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUpdateParams">TeamUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUpdateResponse">TeamUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team/list">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamListParams">TeamListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamListResponse">TeamListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/delete">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamDeleteParams">TeamDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamDeleteResponse">TeamDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/member_add">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.AddMember">AddMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamAddMemberParams">TeamAddMemberParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamAddMemberResponse">TeamAddMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/block">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.Block">Block</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamBlockParams">TeamBlockParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamBlockResponse">TeamBlockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/{team_id}/disable_logging">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.DisableLogging">DisableLogging</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamDisableLoggingResponse">TeamDisableLoggingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team/available">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamListAvailableParams">TeamListAvailableParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamListAvailableResponse">TeamListAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/member_delete">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.RemoveMember">RemoveMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamRemoveMemberParams">TeamRemoveMemberParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamRemoveMemberResponse">TeamRemoveMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team/info">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamGetInfoParams">TeamGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamGetInfoResponse">TeamGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/unblock">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.Unblock">Unblock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUnblockParams">TeamUnblockParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUnblockResponse">TeamUnblockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/member_update">client.Team.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamService.UpdateMember">UpdateMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUpdateMemberParams">TeamUpdateMemberParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamUpdateMemberResponse">TeamUpdateMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Model

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelAddResponse">TeamModelAddResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelRemoveResponse">TeamModelRemoveResponse</a>

Methods:

- <code title="post /team/model/add">client.Team.Model.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelAddParams">TeamModelAddParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelAddResponse">TeamModelAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/model/delete">client.Team.Model.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelRemoveParams">TeamModelRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamModelRemoveResponse">TeamModelRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Callback

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackGetResponse">TeamCallbackGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackAddResponse">TeamCallbackAddResponse</a>

Methods:

- <code title="get /team/{team_id}/callback">client.Team.Callback.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackGetResponse">TeamCallbackGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /team/{team_id}/callback">client.Team.Callback.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackAddParams">TeamCallbackAddParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#TeamCallbackAddResponse">TeamCallbackAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organization

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrgMemberParam">OrgMemberParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationNewResponse">OrganizationNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationUpdateResponse">OrganizationUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationListResponse">OrganizationListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationDeleteResponse">OrganizationDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationAddMemberResponse">OrganizationAddMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationDeleteMemberResponse">OrganizationDeleteMemberResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationUpdateMemberResponse">OrganizationUpdateMemberResponse</a>

Methods:

- <code title="post /organization/new">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationNewParams">OrganizationNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationNewResponse">OrganizationNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /organization/update">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationUpdateParams">OrganizationUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationUpdateResponse">OrganizationUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /organization/list">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationListResponse">OrganizationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /organization/delete">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationDeleteParams">OrganizationDeleteParams</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationDeleteResponse">OrganizationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /organization/member_add">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.AddMember">AddMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationAddMemberParams">OrganizationAddMemberParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationAddMemberResponse">OrganizationAddMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /organization/member_delete">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.DeleteMember">DeleteMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationDeleteMemberParams">OrganizationDeleteMemberParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationDeleteMemberResponse">OrganizationDeleteMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /organization/member_update">client.Organization.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationService.UpdateMember">UpdateMember</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationUpdateMemberParams">OrganizationUpdateMemberParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationUpdateMemberResponse">OrganizationUpdateMemberResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Info

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoGetResponse">OrganizationInfoGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoDeprecatedResponse">OrganizationInfoDeprecatedResponse</a>

Methods:

- <code title="get /organization/info">client.Organization.Info.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoGetParams">OrganizationInfoGetParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoGetResponse">OrganizationInfoGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /organization/info">client.Organization.Info.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoService.Deprecated">Deprecated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoDeprecatedParams">OrganizationInfoDeprecatedParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#OrganizationInfoDeprecatedResponse">OrganizationInfoDeprecatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customer

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BlockUsersParam">BlockUsersParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerNewResponse">CustomerNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerUpdateResponse">CustomerUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerListResponse">CustomerListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerDeleteResponse">CustomerDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerBlockResponse">CustomerBlockResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerGetInfoResponse">CustomerGetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerUnblockResponse">CustomerUnblockResponse</a>

Methods:

- <code title="post /customer/new">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerNewParams">CustomerNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerNewResponse">CustomerNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/update">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerUpdateParams">CustomerUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerUpdateResponse">CustomerUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customer/list">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/delete">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerDeleteParams">CustomerDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerDeleteResponse">CustomerDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/block">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.Block">Block</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerBlockParams">CustomerBlockParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerBlockResponse">CustomerBlockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customer/info">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerGetInfoParams">CustomerGetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerGetInfoResponse">CustomerGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customer/unblock">client.Customer.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerService.Unblock">Unblock</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerUnblockParams">CustomerUnblockParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CustomerUnblockResponse">CustomerUnblockResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Spend

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendCalculateSpendResponse">SpendCalculateSpendResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendListLogsResponse">SpendListLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendListTagsResponse">SpendListTagsResponse</a>

Methods:

- <code title="post /spend/calculate">client.Spend.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendService.CalculateSpend">CalculateSpend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendCalculateSpendParams">SpendCalculateSpendParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendCalculateSpendResponse">SpendCalculateSpendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /spend/logs">client.Spend.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendService.ListLogs">ListLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendListLogsParams">SpendListLogsParams</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendListLogsResponse">SpendListLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /spend/tags">client.Spend.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendService.ListTags">ListTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendListTagsParams">SpendListTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#SpendListTagsResponse">SpendListTagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Global

## Spend

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendListTagsResponse">GlobalSpendListTagsResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendResetResponse">GlobalSpendResetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendGetReportResponse">GlobalSpendGetReportResponse</a>

Methods:

- <code title="get /global/spend/tags">client.Global.Spend.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendService.ListTags">ListTags</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendListTagsParams">GlobalSpendListTagsParams</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendListTagsResponse">GlobalSpendListTagsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /global/spend/reset">client.Global.Spend.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendService.Reset">Reset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendResetResponse">GlobalSpendResetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /global/spend/report">client.Global.Spend.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendService.GetReport">GetReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendGetReportParams">GlobalSpendGetReportParams</a>) ([]<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GlobalSpendGetReportResponse">GlobalSpendGetReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Provider

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ProviderListBudgetsResponse">ProviderListBudgetsResponse</a>

Methods:

- <code title="get /provider/budgets">client.Provider.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ProviderService.ListBudgets">ListBudgets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#ProviderListBudgetsResponse">ProviderListBudgetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Cache

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheDeleteResponse">CacheDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheFlushAllResponse">CacheFlushAllResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CachePingResponse">CachePingResponse</a>

Methods:

- <code title="post /cache/delete">client.Cache.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheDeleteResponse">CacheDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /cache/flushall">client.Cache.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheService.FlushAll">FlushAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheFlushAllResponse">CacheFlushAllResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /cache/ping">client.Cache.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheService.Ping">Ping</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CachePingResponse">CachePingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Redis

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheRediGetInfoResponse">CacheRediGetInfoResponse</a>

Methods:

- <code title="get /cache/redis/info">client.Cache.Redis.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheRediService.GetInfo">GetInfo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#CacheRediGetInfoResponse">CacheRediGetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Guardrails

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GuardrailListResponse">GuardrailListResponse</a>

Methods:

- <code title="get /guardrails/list">client.Guardrails.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GuardrailService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#GuardrailListResponse">GuardrailListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Add

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#IPAddressParam">IPAddressParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AddAddAllowedIPResponse">AddAddAllowedIPResponse</a>

Methods:

- <code title="post /add/allowed_ip">client.Add.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AddService.AddAllowedIP">AddAllowedIP</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AddAddAllowedIPParams">AddAddAllowedIPParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#AddAddAllowedIPResponse">AddAddAllowedIPResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Delete

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#DeleteNewAllowedIPResponse">DeleteNewAllowedIPResponse</a>

Methods:

- <code title="post /delete/allowed_ip">client.Delete.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#DeleteService.NewAllowedIP">NewAllowedIP</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#DeleteNewAllowedIPParams">DeleteNewAllowedIPParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#DeleteNewAllowedIPResponse">DeleteNewAllowedIPResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Files

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileNewResponse">FileNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileGetResponse">FileGetResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileListResponse">FileListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileDeleteResponse">FileDeleteResponse</a>

Methods:

- <code title="post /{provider}/v1/files">client.Files.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileNewParams">FileNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileNewResponse">FileNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileGetResponse">FileGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /{provider}/v1/files">client.Files.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileListParams">FileListParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileListResponse">FileListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /{provider}/v1/files/{file_id}">client.Files.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileDeleteResponse">FileDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Content

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileContentGetResponse">FileContentGetResponse</a>

Methods:

- <code title="get /{provider}/v1/files/{file_id}/content">client.Files.Content.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileContentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, provider <a href="https://pkg.go.dev/builtin#string">string</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#FileContentGetResponse">FileContentGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Budget

Params Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetNewParam">BudgetNewParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetNewResponse">BudgetNewResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetUpdateResponse">BudgetUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetListResponse">BudgetListResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetDeleteResponse">BudgetDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetInfoResponse">BudgetInfoResponse</a>
- <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetSettingsResponse">BudgetSettingsResponse</a>

Methods:

- <code title="post /budget/new">client.Budget.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetNewParams">BudgetNewParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetNewResponse">BudgetNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /budget/update">client.Budget.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetUpdateParams">BudgetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetUpdateResponse">BudgetUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budget/list">client.Budget.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetListResponse">BudgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /budget/delete">client.Budget.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetDeleteParams">BudgetDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetDeleteResponse">BudgetDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /budget/info">client.Budget.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetService.Info">Info</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetInfoParams">BudgetInfoParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetInfoResponse">BudgetInfoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budget/settings">client.Budget.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetService.Settings">Settings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetSettingsParams">BudgetSettingsParams</a>) (<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk">hanzoai</a>.<a href="https://pkg.go.dev/github.com/hanzoai/go-sdk#BudgetSettingsResponse">BudgetSettingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
