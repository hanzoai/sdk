//! Agent framework

use anyhow::Result;
use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize)]
pub struct Agent {
    pub name: String,
    pub model: String,
    pub system_prompt: Option<String>,
    pub tools: Vec<String>,
}

impl Agent {
    pub fn new(name: String) -> Self {
        Self {
            name,
            model: "gpt-4".to_string(),
            system_prompt: None,
            tools: Vec::new(),
        }
    }

    pub async fn run(&self, task: &str) -> Result<String> {
        Ok(format!("Task '{}' completed by {}", task, self.name))
    }

    pub fn add_tool(&mut self, tool: String) {
        self.tools.push(tool);
    }
}