#!/usr/bin/env python3
"""Hanzo SDK Unified Conformance Test Harness & Mock Server.

Validates OpenAPI specifications, tests conformance fixtures, checks wire format
compatibility, and serves a local mock API server for cross-language SDK testing.
Uses Python 3 standard library only (zero external dependencies).
"""

from __future__ import annotations

import argparse
import http.server
import json
import os
import re
import socketserver
import sys
import threading
import time
import urllib.error
import urllib.request
from pathlib import Path

ROOT_DIR = Path(__file__).resolve().parent.parent
SPEC_FILE = ROOT_DIR / "spec" / "openapi.json"
FIXTURES_DIR = ROOT_DIR / "conformance" / "fixtures"


class MockApiHandler(http.server.BaseHTTPRequestHandler):
    """Mock HTTP handler serving canonical conformance fixtures."""

    def log_message(self, format: str, *args: object) -> None:
        # Suppress noisy default logging during automated test runs
        if os.environ.get("CONFORMANCE_VERBOSE"):
            super().log_message(format, *args)

    def _send_json(self, status_code: int, data: dict | list) -> None:
        body = json.dumps(data, indent=2).encode("utf-8")
        self.send_response(status_code)
        self.send_header("Content-Type", "application/json; charset=utf-8")
        self.send_header("Content-Length", str(len(body)))
        self.send_header("Access-Control-Allow-Origin", "*")
        self.send_header("Access-Control-Allow-Headers", "*")
        self.send_header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        self.end_headers()
        self.wfile.write(body)

    def do_OPTIONS(self) -> None:
        self.send_response(204)
        self.send_header("Access-Control-Allow-Origin", "*")
        self.send_header("Access-Control-Allow-Headers", "*")
        self.send_header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        self.end_headers()

    def do_GET(self) -> None:
        path = self.path.split("?")[0]

        if path in ("/v1/models", "/models"):
            # Catalog listing requires no credentials
            fixture = FIXTURES_DIR / "models_list.json"
            if fixture.exists():
                with open(fixture, "r", encoding="utf-8") as f:
                    self._send_json(200, json.load(f))
                return
            self._send_json(500, {"error": "Missing models_list.json fixture"})
            return

        if path.startswith("/v1/models/") or path.startswith("/models/"):
            model_id = path.split("/")[-1]
            self._send_json(200, {
                "id": model_id,
                "object": "model",
                "created": 1715000000,
                "owned_by": "hanzo",
                "pricing": {"input": 0.15, "output": 0.60}
            })
            return

        self._send_json(404, {"error": {"message": f"Route {path} not found", "type": "not_found"}})

    def do_POST(self) -> None:
        auth = self.headers.get("Authorization", "")
        if not auth or auth == "Bearer invalid":
            with open(FIXTURES_DIR / "error_auth.json", "r", encoding="utf-8") as f:
                self._send_json(401, json.load(f))
            return

        content_length = int(self.headers.get("Content-Length", 0))
        post_data = self.rfile.read(content_length)
        payload = {}
        if post_data:
            try:
                payload = json.loads(post_data.decode("utf-8"))
            except Exception:
                pass

        path = self.path.split("?")[0]

        if path in ("/v1/chat/completions", "/chat/completions"):
            # Check for streaming request
            if payload.get("stream") is True:
                sse_fixture = FIXTURES_DIR / "chat_completion_stream.sse"
                if sse_fixture.exists():
                    raw_sse = sse_fixture.read_bytes()
                    self.send_response(200)
                    self.send_header("Content-Type", "text/event-stream; charset=utf-8")
                    self.send_header("Cache-Control", "no-cache")
                    self.send_header("Content-Length", str(len(raw_sse)))
                    self.send_header("Connection", "close")
                    self.end_headers()
                    self.wfile.write(raw_sse)
                    self.wfile.flush()
                    return

            # Check if tools are passed
            if payload.get("tools"):
                tool_fixture = FIXTURES_DIR / "tool_call_response.json"
                with open(tool_fixture, "r", encoding="utf-8") as f:
                    self._send_json(200, json.load(f))
                return

            # Standard chat completion response
            chat_fixture = FIXTURES_DIR / "chat_completion_response.json"
            with open(chat_fixture, "r", encoding="utf-8") as f:
                self._send_json(200, json.load(f))
            return

        if path in ("/v1/embeddings", "/embeddings"):
            emb_fixture = FIXTURES_DIR / "embeddings_response.json"
            with open(emb_fixture, "r", encoding="utf-8") as f:
                self._send_json(200, json.load(f))
            return

        if path in ("/v1/decisions", "/decisions"):
            kai_fixture = FIXTURES_DIR / "kai_decision_response.json"
            with open(kai_fixture, "r", encoding="utf-8") as f:
                self._send_json(200, json.load(f))
            return

        self._send_json(404, {"error": {"message": f"Route {path} not found", "type": "not_found"}})


def run_mock_server(host: str = "127.0.0.1", port: int = 8080) -> socketserver.TCPServer:
    class ReusableServer(socketserver.TCPServer):
        allow_reuse_address = True

    server = ReusableServer((host, port), MockApiHandler)
    return server


def validate_fixtures() -> list[str]:
    """Validates that all JSON fixtures are syntactically sound and non-empty."""
    errors = []
    print(" [1/3] Validating conformance fixtures...")
    if not FIXTURES_DIR.exists():
        return [f"Fixtures directory missing: {FIXTURES_DIR}"]

    required_fixtures = [
        "auth.json",
        "models_list.json",
        "chat_completion_request.json",
        "chat_completion_response.json",
        "chat_completion_stream.sse",
        "tool_call_request.json",
        "tool_call_response.json",
        "kai_decision_request.json",
        "kai_decision_response.json",
        "embeddings_request.json",
        "embeddings_response.json",
        "error_auth.json",
        "error_rate_limit.json",
        "error_bad_request.json",
    ]

    for name in required_fixtures:
        fixture_path = FIXTURES_DIR / name
        if not fixture_path.exists():
            errors.append(f"Missing required fixture: {name}")
            continue

        if name.endswith(".json"):
            try:
                with open(fixture_path, "r", encoding="utf-8") as f:
                    data = json.load(f)
                if not data:
                    errors.append(f"Empty JSON fixture: {name}")
                else:
                    print(f"    ✓ {name} (valid JSON)")
            except Exception as e:
                errors.append(f"Invalid JSON in {name}: {e}")
        elif name.endswith(".sse"):
            content = fixture_path.read_text(encoding="utf-8")
            if "data: [DONE]" not in content:
                errors.append(f"SSE fixture {name} must terminate with data: [DONE]")
            else:
                print(f"    ✓ {name} (valid SSE stream format)")

    return errors


def validate_openapi_spec() -> list[str]:
    """Validates the canonical OpenAPI 3.1 specification."""
    errors = []
    print("\n [2/3] Validating OpenAPI 3.1 specification...")
    if not SPEC_FILE.exists():
        return [f"OpenAPI spec file missing: {SPEC_FILE}"]

    try:
        with open(SPEC_FILE, "r", encoding="utf-8") as f:
            spec = json.load(f)

        assert spec.get("openapi", "").startswith("3."), "Spec openapi version must be 3.x"
        assert "paths" in spec, "Spec must define paths"
        paths = spec["paths"]
        expected_paths = ["/models", "/chat/completions", "/embeddings", "/decisions"]
        for p in expected_paths:
            if p not in paths:
                errors.append(f"Spec missing canonical path: {p}")
            else:
                print(f"    ✓ Path defined: {p}")

        assert "components" in spec and "schemas" in spec["components"], "Spec must define schemas"
        print(f"    ✓ Total paths: {len(paths)}, Total schemas: {len(spec['components']['schemas'])}")
    except Exception as e:
        errors.append(f"OpenAPI validation error: {e}")

    return errors


def test_mock_endpoints(host: str = "127.0.0.1", port: int = 8080) -> list[str]:
    """Starts local mock server and executes client verification against all routes."""
    errors = []
    print(f"\n [3/3] Testing Wire Conformance against Mock Gateway (http://{host}:{port}/v1)...")

    server = run_mock_server(host, port)
    thread = threading.Thread(target=server.serve_forever, daemon=True)
    thread.start()

    base_url = f"http://{host}:{port}/v1"
    time.sleep(0.1)

    try:
        # Test 1: GET /v1/models (No Auth required)
        req = urllib.request.Request(f"{base_url}/models")
        with urllib.request.urlopen(req) as resp:
            assert resp.status == 200
            data = json.loads(resp.read().decode())
            assert data["object"] == "list"
            assert any(m["id"] == "zen5" for m in data["data"])
            print("    ✓ GET  /v1/models (200 OK, Catalog returned without auth)")

        # Test 2: POST /v1/chat/completions without Auth -> Should 401
        req = urllib.request.Request(
            f"{base_url}/chat/completions",
            data=json.dumps({"model": "zen5", "messages": [{"role": "user", "content": "hi"}]}).encode(),
            headers={"Content-Type": "application/json"}
        )
        try:
            with urllib.request.urlopen(req) as resp:
                errors.append("Expected 401 Unauthorized for unauthenticated POST /v1/chat/completions")
        except urllib.error.HTTPError as e:
            assert e.code == 401
            err_data = json.loads(e.read().decode())
            assert "error" in err_data
            print("    ✓ POST /v1/chat/completions [Unauthenticated] (401 Unauthorized as expected)")

        # Test 3: POST /v1/chat/completions (Unary)
        req = urllib.request.Request(
            f"{base_url}/chat/completions",
            data=json.dumps({"model": "zen5", "messages": [{"role": "user", "content": "hi"}]}).encode(),
            headers={
                "Authorization": "Bearer sk-test-key-conformance",
                "Content-Type": "application/json"
            }
        )
        with urllib.request.urlopen(req) as resp:
            assert resp.status == 200
            data = json.loads(resp.read().decode())
            assert data["object"] == "chat.completion"
            assert len(data["choices"]) > 0
            print("    ✓ POST /v1/chat/completions [Unary] (200 OK, Completion parsed)")

        # Test 4: POST /v1/chat/completions (SSE Streaming)
        req = urllib.request.Request(
            f"{base_url}/chat/completions",
            data=json.dumps({
                "model": "zen5",
                "messages": [{"role": "user", "content": "hi"}],
                "stream": True
            }).encode(),
            headers={
                "Authorization": "Bearer sk-test-key-conformance",
                "Content-Type": "application/json"
            }
        )
        with urllib.request.urlopen(req) as resp:
            assert resp.status == 200
            assert "text/event-stream" in resp.headers.get("Content-Type", "")
            chunks = resp.read().decode()
            assert "data: [DONE]" in chunks
            print("    ✓ POST /v1/chat/completions [Stream] (200 OK, Valid SSE stream chunked)")

        # Test 5: POST /v1/decisions (Kai Governance API)
        req = urllib.request.Request(
            f"{base_url}/decisions",
            data=json.dumps({
                "agent_id": "test-agent",
                "policy_id": "pol-sec",
                "action": {"type": "test", "payload": {}}
            }).encode(),
            headers={
                "Authorization": "Bearer sk-test-key-conformance",
                "Content-Type": "application/json"
            }
        )
        with urllib.request.urlopen(req) as resp:
            assert resp.status == 200
            data = json.loads(resp.read().decode())
            assert data.get("verdict") in ("ALLOW", "DENY", "MODIFY")
            print(f"    ✓ POST /v1/decisions [Kai Policy] (200 OK, Verdict: {data['verdict']})")

        # Test 6: POST /v1/embeddings
        req = urllib.request.Request(
            f"{base_url}/embeddings",
            data=json.dumps({"model": "zen-embed", "input": "test embedding"}).encode(),
            headers={
                "Authorization": "Bearer sk-test-key-conformance",
                "Content-Type": "application/json"
            }
        )
        with urllib.request.urlopen(req) as resp:
            assert resp.status == 200
            data = json.loads(resp.read().decode())
            assert data["object"] == "list"
            assert len(data["data"][0]["embedding"]) > 0
            print("    ✓ POST /v1/embeddings (200 OK, Dense vector received)")

    except Exception as e:
        errors.append(f"Mock endpoint test failed: {e}")
    finally:
        server.shutdown()
        server.server_close()

    return errors


def main() -> None:
    parser = argparse.ArgumentParser(description="Hanzo SDK Conformance Harness")
    parser.add_argument("--server", action="store_true", help="Run standalone mock server")
    parser.add_argument("--port", type=int, default=8080, help="Port for mock server (default: 8080)")
    parser.add_argument("--host", default="127.0.0.1", help="Host for mock server (default: 127.0.0.1)")
    parser.add_argument("--test-mock", action="store_true", help="Run wire test suite against mock server")
    parser.add_argument("--validate-all", action="store_true", help="Validate fixtures, specs, and mock wire")

    args = parser.parse_args()

    if args.server:
        print(f"Starting Hanzo Mock API Gateway at http://{args.host}:{args.port}/v1 ...")
        print("Press Ctrl+C to stop.")
        server = run_mock_server(args.host, args.port)
        try:
            server.serve_forever()
        except KeyboardInterrupt:
            print("\nShutting down mock server.")
            server.shutdown()
            server.server_close()
        return

    all_errors = []
    all_errors.extend(validate_fixtures())
    all_errors.extend(validate_openapi_spec())

    if args.validate_all or args.test_mock or len(sys.argv) == 1:
        all_errors.extend(test_mock_endpoints(args.host, args.port))

    if all_errors:
        print(f"\n❌ Conformance verification FAILED with {len(all_errors)} error(s):")
        for err in all_errors:
            print(f"  - {err}")
        sys.exit(1)
    else:
        print("\n✅ All Conformance and Specification Checks PASSED successfully!\n")
        sys.exit(0)


if __name__ == "__main__":
    main()
