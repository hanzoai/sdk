"""Tests for Python client."""

import pytest
from hanzo.client import HanzoClient


def test_client_creation():
    """Test client can be created."""
    client = HanzoClient()
    assert client is not None
    assert client.base_url == "https://api.hanzo.ai"


def test_client_with_api_key():
    """Test client with API key."""
    client = HanzoClient(api_key="test-key")
    assert client.api_key == "test-key"


def test_client_context_manager():
    """Test client works as context manager."""
    with HanzoClient() as client:
        assert client is not None