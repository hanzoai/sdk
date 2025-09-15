#!/usr/bin/env python3
"""Integration tests for Hanzo SDK."""

import sys
import os

# Add src/py to path
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'src', 'py'))

def test_integration():
    """Basic integration test."""
    print("Running integration tests...")
    
    # Test Python import
    import hanzo
    print(f"✓ Python SDK version: {hanzo.__version__}")
    
    # Test that Node.js is available
    import subprocess
    result = subprocess.run(['node', '--version'], capture_output=True, text=True)
    if result.returncode == 0:
        print(f"✓ Node.js version: {result.stdout.strip()}")
    
    print("✓ Integration tests passed")
    return True

if __name__ == "__main__":
    success = test_integration()
    sys.exit(0 if success else 1)
