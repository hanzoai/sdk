//! Hanzo API client

use anyhow::Result;
use reqwest::Client;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone)]
pub struct HanzoClient {
    client: Client,
    base_url: String,
    api_key: Option<String>,
}

#[derive(Debug, Serialize, Deserialize)]
pub struct ChatMessage {
    pub role: String,
    pub content: String,
}

impl HanzoClient {
    pub fn new() -> Self {
        Self {
            client: Client::new(),
            base_url: "https://api.hanzo.ai".to_string(),
            api_key: None,
        }
    }

    pub fn with_api_key(mut self, api_key: String) -> Self {
        self.api_key = Some(api_key);
        self
    }

    pub async fn chat_completion(&self, messages: Vec<ChatMessage>) -> Result<String> {
        // Implementation would make API call
        Ok("Response from Hanzo API".to_string())
    }
}