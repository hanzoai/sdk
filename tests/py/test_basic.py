"""Basic tests for Hanzo Python SDK."""

import pytest
import sys
import os

# Add src/py to path
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', '..', 'src', 'py'))


def test_import():
    """Test that we can import hanzo module."""
    import hanzo
    assert hanzo.__version__


def test_cli_import():
    """Test that we can import CLI module."""
    from hanzo import cli
    assert cli is not None
