//! Hanzo SDK - Rust implementation

pub mod client;
pub mod node;
pub mod agent;
pub mod network;
pub mod crypto;

use napi_derive::napi;

#[napi]
pub struct HanzoSDK {
    client: client::HanzoClient,
}

#[napi]
impl HanzoSDK {
    #[napi(constructor)]
    pub fn new() -> Self {
        Self {
            client: client::HanzoClient::new(),
        }
    }

    #[napi]
    pub async fn start_node(&self, port: u16) -> napi::Result<()> {
        node::start(port).await.map_err(|e| napi::Error::from_reason(e.to_string()))
    }

    #[napi]
    pub async fn stop_node(&self) -> napi::Result<()> {
        node::stop().await.map_err(|e| napi::Error::from_reason(e.to_string()))
    }

    #[napi]
    pub async fn get_node_status(&self) -> napi::Result<String> {
        node::status().await.map_err(|e| napi::Error::from_reason(e.to_string()))
    }

    #[napi]
    pub async fn load_model(&self, model: String) -> napi::Result<()> {
        node::load_model(&model).await.map_err(|e| napi::Error::from_reason(e.to_string()))
    }

    #[napi]
    pub async fn list_models(&self) -> napi::Result<Vec<String>> {
        node::list_models().await.map_err(|e| napi::Error::from_reason(e.to_string()))
    }
}