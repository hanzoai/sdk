#!/usr/bin/env python3
"""Python Quickstart for Hanzo AI SDK.

Demonstrates model listing (public), unary chat completion, and Kai decision policy evaluation.
Requires: pip install hanzoai (or pip install hanzo)
"""

import os
from hanzoai import Hanzo

api_key = os.getenv("HANZO_API_KEY", "sk-test-key")
base_url = os.getenv("HANZO_BASE_URL", "https://api.hanzo.ai/v1")

client = Hanzo(api_key=api_key, base_url=base_url)

def main():
    print("--- 1. List Available Models (Catalog) ---")
    models = client.models.list()
    for m in models.data[:5]:
        print(f"Model: {m.id} (Owner: {m.owned_by})")

    print("\n--- 2. Chat Completion ---")
    completion = client.chat.completions.create(
        model="zen5",
        messages=[
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "What is the frontier of agentic AI?"}
        ],
        temperature=0.7
    )
    print("Reply:", completion.choices[0].message.content)

    print("\n--- 3. Kai Agent Decision Evaluation ---")
    decision = client.decisions.evaluate(
        agent_id="agent-01",
        policy_id="pol-safety-core",
        action={"type": "query_database", "payload": {"target": "public_records"}}
    )
    print(f"Decision Verdict: {decision.verdict} (Confidence: {decision.confidence})")
    print(f"Reasoning: {decision.reasoning}")

if __name__ == "__main__":
    main()
