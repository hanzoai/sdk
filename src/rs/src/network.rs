//! Network operations

use anyhow::Result;

pub async fn status() -> Result<String> {
    Ok("Network online".to_string())
}

pub async fn list_peers() -> Result<Vec<String>> {
    Ok(vec!["peer1".to_string(), "peer2".to_string()])
}

pub async fn connect(address: &str) -> Result<()> {
    println!("Connecting to {}", address);
    Ok(())
}

pub async fn disconnect(peer_id: &str) -> Result<()> {
    println!("Disconnecting from {}", peer_id);
    Ok(())
}