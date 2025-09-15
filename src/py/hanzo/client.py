"""Hanzo API client."""

import httpx
from typing import Optional, Dict, Any, List


class HanzoClient:
    """Hanzo API client for Python."""
    
    def __init__(self, api_key: Optional[str] = None, base_url: str = "https://api.hanzo.ai"):
        self.api_key = api_key
        self.base_url = base_url
        self.client = httpx.Client(
            base_url=base_url,
            headers={"Authorization": f"Bearer {api_key}"} if api_key else {}
        )
    
    def chat_completion(self, messages: List[Dict[str, str]], model: str = "gpt-4") -> Dict[str, Any]:
        """Create a chat completion."""
        response = self.client.post(
            "/v1/chat/completions",
            json={"messages": messages, "model": model}
        )
        response.raise_for_status()
        return response.json()
    
    def list_models(self) -> List[str]:
        """List available models."""
        response = self.client.get("/v1/models")
        response.raise_for_status()
        return [m["id"] for m in response.json()["data"]]
    
    def __enter__(self):
        return self
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        self.client.close()