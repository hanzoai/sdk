#!/usr/bin/env bash
# Quickstart using cURL and the Hanzo API Gateway

set -euo pipefail

HANZO_API_KEY="${HANZO_API_KEY:-sk-test-key}"
HANZO_BASE_URL="${HANZO_BASE_URL:-https://api.hanzo.ai/v1}"

echo "=== 1. List Models (Public, No Auth Required) ==="
curl -s "${HANZO_BASE_URL}/models" | jq '.data[0:5]'

echo -e "\n=== 2. Unary Chat Completion ==="
curl -s "${HANZO_BASE_URL}/chat/completions" \
  -H "Authorization: Bearer ${HANZO_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "zen5",
    "messages": [
      {"role": "system", "content": "You are a helpful assistant."},
      {"role": "user", "content": "Explain quantum computing in 10 words."}
    ],
    "temperature": 0.7
  }' | jq '.choices[0].message.content'

echo -e "\n=== 3. Streaming SSE Chat Completion ==="
curl -N -s "${HANZO_BASE_URL}/chat/completions" \
  -H "Authorization: Bearer ${HANZO_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "zen5",
    "messages": [
      {"role": "user", "content": "Stream count 1 to 5."}
    ],
    "stream": true
  }'

echo -e "\n=== 4. Kai Agent Policy Decision Evaluation ==="
curl -s "${HANZO_BASE_URL}/decisions" \
  -H "Authorization: Bearer ${HANZO_API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{
    "agent_id": "agent-curl-01",
    "policy_id": "pol-safety-baseline",
    "action": {
      "type": "database_write",
      "payload": {"table": "logs", "record": {"event": "login"}}
    }
  }' | jq '.'
