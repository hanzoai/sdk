//! Rust Quickstart for Hanzo AI SDK
//!
//! Add to Cargo.toml:
//! [dependencies]
//! hanzo-client = "1.0"
//! tokio = { version = "1", features = ["full"] }

use hanzo_client::{Client, ChatMessage, Role};
use std::env;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let api_key = env::var("HANZO_API_KEY").unwrap_or_else(|_| "sk-test-key".into());
    let client = Client::new(&api_key);

    println!("--- 1. List Models ---");
    let models = client.models().list().await?;
    println!("Retrieved {} models", models.data.len());

    println!("\n--- 2. Chat Completion ---");
    let response = client
        .chat()
        .create_completion("zen5", vec![
            ChatMessage {
                role: Role::User,
                content: "Hello from Rust!".into(),
            }
        ])
        .await?;

    if let Some(choice) = response.choices.first() {
        println!("Reply: {}", choice.message.content);
    }

    Ok(())
}
