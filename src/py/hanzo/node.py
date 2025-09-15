"""Node management."""

import asyncio
from typing import Optional, List, Dict, Any


class Node:
    """Hanzo AI Node."""
    
    def __init__(self, port: int = 4000):
        self.port = port
        self.running = False
        self.models = []
    
    async def start(self):
        """Start the node."""
        self.running = True
        print(f"Node started on port {self.port}")
    
    async def stop(self):
        """Stop the node."""
        self.running = False
        print("Node stopped")
    
    def status(self) -> Dict[str, Any]:
        """Get node status."""
        return {
            "running": self.running,
            "port": self.port,
            "models": self.models
        }
    
    async def load_model(self, model: str):
        """Load a model."""
        self.models.append(model)
        print(f"Model {model} loaded")
    
    def list_models(self) -> List[str]:
        """List loaded models."""
        return self.models


# Module-level functions for CLI
async def start(port: int = 4000):
    """Start a node."""
    node = Node(port)
    await node.start()
    return node


async def stop():
    """Stop the node."""
    # Implementation would stop running node
    return "Node stopped"


async def status() -> str:
    """Get node status."""
    # Implementation would check actual status
    return "Node running on port 4000"


async def load_model(model: str):
    """Load a model."""
    # Implementation would load model
    return f"Model {model} loaded"


async def list_models() -> List[str]:
    """List available models."""
    # Implementation would list actual models
    return ["gpt-4", "claude-3", "llama-2"]