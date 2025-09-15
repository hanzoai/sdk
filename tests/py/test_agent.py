"""Tests for Agent framework."""

import pytest
import asyncio
from hanzo.agent import Agent


def test_agent_creation():
    """Test agent can be created."""
    agent = Agent(name="test-agent")
    assert agent.name == "test-agent"
    assert agent.model == "gpt-4"
    assert agent.tools == []


def test_agent_with_tools():
    """Test agent with tools."""
    agent = Agent(name="test-agent")
    agent.add_tool("filesystem")
    agent.add_tool("shell")
    assert len(agent.tools) == 2
    assert "filesystem" in agent.tools


@pytest.mark.asyncio
async def test_agent_run():
    """Test agent run method."""
    agent = Agent(name="test-agent")
    result = await agent.run("test task")
    assert "test task" in result
    assert "test-agent" in result


def test_agent_to_dict():
    """Test agent serialization."""
    agent = Agent(name="test-agent", model="claude-3")
    data = agent.to_dict()
    assert data["name"] == "test-agent"
    assert data["model"] == "claude-3"