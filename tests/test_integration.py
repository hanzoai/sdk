#!/usr/bin/env python3
"""Integration tests for Hanzo SDK across all languages."""

import subprocess
import json
import os
import sys
from pathlib import Path

class TestRunner:
    def __init__(self):
        self.sdk_root = Path(__file__).parent.parent
        self.results = []
        
    def run_command(self, cmd, expected_output=None, should_fail=False):
        """Run a command and check output."""
        try:
            result = subprocess.run(
                cmd,
                shell=True,
                capture_output=True,
                text=True,
                timeout=10
            )
            
            if should_fail and result.returncode == 0:
                return False, f"Command should have failed: {cmd}"
            elif not should_fail and result.returncode != 0:
                return False, f"Command failed: {cmd}\n{result.stderr}"
            
            if expected_output and expected_output not in result.stdout:
                return False, f"Output mismatch for {cmd}\nExpected: {expected_output}\nGot: {result.stdout}"
            
            return True, "OK"
        except subprocess.TimeoutExpired:
            return False, f"Command timed out: {cmd}"
        except Exception as e:
            return False, f"Error running {cmd}: {e}"
    
    def test_python_cli(self):
        """Test Python CLI commands."""
        print("Testing Python CLI...")
        tests = [
            ("python -m hanzo.cli --version", "0.1.0"),
            ("python -m hanzo.cli --help", "Hanzo AI SDK"),
            ("python -m hanzo.cli node --help", "Manage Hanzo AI nodes"),
            ("python -m hanzo.cli agent --help", "Manage AI agents"),
            ("python -m hanzo.cli mcp --help", "Model Context Protocol"),
        ]
        
        for cmd, expected in tests:
            success, msg = self.run_command(cmd, expected)
            self.results.append(("Python", cmd, success, msg))
    
    def test_javascript_cli(self):
        """Test JavaScript CLI commands."""
        print("Testing JavaScript CLI...")
        
        # Build first
        build_success, build_msg = self.run_command("npm run build 2>/dev/null")
        if not build_success:
            self.results.append(("JavaScript", "npm run build", False, build_msg))
            return
        
        tests = [
            ("node dist/cli.js --version", "0.1.0"),
            ("node dist/cli.js --help", "Hanzo AI SDK"),
            ("node dist/cli.js node --help", "Manage Hanzo AI nodes"),
            ("node dist/cli.js agent --help", "Manage AI agents"),
        ]
        
        for cmd, expected in tests:
            success, msg = self.run_command(cmd, expected)
            self.results.append(("JavaScript", cmd, success, msg))
    
    def test_rust_cli(self):
        """Test Rust CLI commands."""
        print("Testing Rust CLI...")
        
        # Build first
        build_success, build_msg = self.run_command("cd src/rs && cargo build --release --quiet 2>/dev/null")
        if not build_success:
            self.results.append(("Rust", "cargo build", False, "Build failed (expected for now)"))
            return
        
        tests = [
            ("src/rs/target/release/hanzo --version", "0.1.0"),
            ("src/rs/target/release/hanzo --help", "Hanzo"),
        ]
        
        for cmd, expected in tests:
            success, msg = self.run_command(cmd, expected)
            self.results.append(("Rust", cmd, success, msg))
    
    def test_cross_language_compatibility(self):
        """Test that all languages produce compatible outputs."""
        print("Testing cross-language compatibility...")
        
        # Test that config files are compatible
        config_test = {
            "test_key": "test_value",
            "api_key": "sk-test-123"
        }
        
        config_path = Path.home() / ".hanzo" / "test_config.json"
        config_path.parent.mkdir(parents=True, exist_ok=True)
        config_path.write_text(json.dumps(config_test))
        
        # Each language should be able to read the same config
        tests = [
            ("Python config read", "python -c \"import json; print(json.load(open('" + str(config_path) + "'))['test_key'])\"", "test_value"),
            ("JS config read", "node -e \"console.log(JSON.parse(require('fs').readFileSync('" + str(config_path) + "')).test_key)\"", "test_value"),
        ]
        
        for name, cmd, expected in tests:
            success, msg = self.run_command(cmd, expected)
            self.results.append(("Cross-language", name, success, msg))
        
        # Clean up
        if config_path.exists():
            config_path.unlink()
    
    def print_results(self):
        """Print test results."""
        print("\n" + "="*60)
        print("TEST RESULTS")
        print("="*60)
        
        passed = 0
        failed = 0
        
        for language, test, success, msg in self.results:
            status = "✅ PASS" if success else "❌ FAIL"
            print(f"{status} [{language}] {test[:50]}")
            if not success and msg != "OK":
                print(f"    {msg}")
            
            if success:
                passed += 1
            else:
                failed += 1
        
        print("\n" + "-"*60)
        print(f"Total: {passed + failed} tests")
        print(f"Passed: {passed}")
        print(f"Failed: {failed}")
        print(f"Success Rate: {passed/(passed+failed)*100:.1f}%")
        
        return failed == 0
    
    def run_all_tests(self):
        """Run all tests."""
        print("Running Hanzo SDK Integration Tests")
        print("="*60)
        
        self.test_python_cli()
        self.test_javascript_cli()
        self.test_rust_cli()
        self.test_cross_language_compatibility()
        
        return self.print_results()


if __name__ == "__main__":
    runner = TestRunner()
    success = runner.run_all_tests()
    sys.exit(0 if success else 1)