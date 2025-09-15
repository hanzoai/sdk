"""Agent framework."""

from typing import Optional, Dict, Any, List
from dataclasses import dataclass


@dataclass
class Agent:
    """AI Agent."""
    
    name: str
    model: str = "gpt-4"
    system_prompt: Optional[str] = None
    tools: List[str] = None
    
    def __post_init__(self):
        if self.tools is None:
            self.tools = []
    
    async def run(self, task: str) -> str:
        """Run a task."""
        # Implementation would execute the agent
        return f"Task '{task}' completed by {self.name}"
    
    def add_tool(self, tool: str):
        """Add a tool to the agent."""
        self.tools.append(tool)
    
    def to_dict(self) -> Dict[str, Any]:
        """Convert to dictionary."""
        return {
            "name": self.name,
            "model": self.model,
            "system_prompt": self.system_prompt,
            "tools": self.tools
        }