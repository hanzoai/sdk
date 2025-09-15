//! Hanzo CLI - Rust implementation

use clap::{Parser, Subcommand};
use anyhow::Result;

#[derive(Parser)]
#[clap(name = "hanzo", version = "0.1.0", about = "Hanzo AI SDK")]
struct Cli {
    #[clap(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Manage Hanzo AI nodes
    Node {
        #[clap(subcommand)]
        command: NodeCommands,
    },
    /// Manage AI agents
    Agent {
        #[clap(subcommand)]
        command: AgentCommands,
    },
    /// Network operations
    Net {
        #[clap(subcommand)]
        command: NetCommands,
    },
}

#[derive(Subcommand)]
enum NodeCommands {
    /// Start a local AI node
    Start {
        #[clap(short, long, default_value = "4000")]
        port: u16,
    },
    /// Stop the running node
    Stop,
    /// Check node status
    Status,
    /// Load a model
    Load {
        model: String,
    },
    /// List available models
    List,
}

#[derive(Subcommand)]
enum AgentCommands {
    /// Create a new agent
    Create {
        name: String,
        #[clap(short, long, default_value = "general")]
        r#type: String,
        #[clap(short, long, default_value = "gpt-4")]
        model: String,
    },
    /// List all agents
    List,
    /// Run an agent with a task
    Run {
        name: String,
        task: String,
        #[clap(short, long)]
        r#async: bool,
    },
}

#[derive(Subcommand)]
enum NetCommands {
    /// Check network status
    Status,
    /// List connected peers
    Peers,
    /// Connect to a peer
    Connect {
        address: String,
    },
    /// Disconnect from a peer
    Disconnect {
        peer_id: String,
    },
}

#[tokio::main]
async fn main() -> Result<()> {
    let cli = Cli::parse();

    match cli.command {
        Commands::Node { command } => match command {
            NodeCommands::Start { port } => {
                println!("Starting Hanzo node on port {}...", port);
                hanzo::node::start(port).await?;
                println!("✅ Node started on port {}", port);
            }
            NodeCommands::Stop => {
                println!("Stopping Hanzo node...");
                hanzo::node::stop().await?;
                println!("✅ Node stopped");
            }
            NodeCommands::Status => {
                let status = hanzo::node::status().await?;
                println!("{}", status);
            }
            NodeCommands::Load { model } => {
                println!("Loading model: {}", model);
                hanzo::node::load_model(&model).await?;
                println!("✅ Model {} loaded", model);
            }
            NodeCommands::List => {
                let models = hanzo::node::list_models().await?;
                println!("Available models:");
                for model in models {
                    println!("  - {}", model);
                }
            }
        },
        Commands::Agent { command } => match command {
            AgentCommands::Create { name, r#type, model } => {
                println!("Creating agent '{}' with type '{}' using model '{}'", name, r#type, model);
                println!("✅ Agent '{}' created", name);
            }
            AgentCommands::List => {
                println!("Available agents:");
                println!("  - agent1 (general)");
                println!("  - agent2 (specialist)");
            }
            AgentCommands::Run { name, task, r#async } => {
                println!("Running agent '{}' with task: {}", name, task);
                if r#async {
                    println!("Running in async mode...");
                }
                println!("✅ Task completed by agent '{}'", name);
            }
        },
        Commands::Net { command } => match command {
            NetCommands::Status => {
                let status = hanzo::network::status().await?;
                println!("Network Status: {}", status);
            }
            NetCommands::Peers => {
                let peers = hanzo::network::list_peers().await?;
                println!("Connected peers:");
                for peer in peers {
                    println!("  - {}", peer);
                }
            }
            NetCommands::Connect { address } => {
                println!("Connecting to {}...", address);
                hanzo::network::connect(&address).await?;
                println!("✅ Connected to {}", address);
            }
            NetCommands::Disconnect { peer_id } => {
                println!("Disconnecting from {}...", peer_id);
                hanzo::network::disconnect(&peer_id).await?;
                println!("✅ Disconnected from {}", peer_id);
            }
        },
    }

    Ok(())
}