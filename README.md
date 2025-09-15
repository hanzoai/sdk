# Hanzo Unified SDK

Unified multi-language SDK for Hanzo AI infrastructure, combining Python, JavaScript/TypeScript, Rust, and Go implementations.

## Structure

```
sdk/
├── src/           # Source code by language
│   ├── py/        # Python implementation
│   ├── js/        # JavaScript/TypeScript implementation
│   ├── rs/        # Rust implementation
│   └── go/        # Go implementation
├── docs/          # Documentation
├── tests/         # Cross-language tests
├── scripts/       # Build and deployment scripts
└── bin/           # Compiled binaries
```

## Installation

### Python
```bash
pip install hanzo
```

### Node.js
```bash
npm install -g @hanzo/cli
```

### Rust
```bash
cargo install hanzo-cli
```

### Go
```bash
go install github.com/hanzoai/sdk/cmd/hanzo@latest
```

## Commands

All commands work identically across all language implementations:

- `hanzo node start` - Start local AI node
- `hanzo agent run` - Run AI agents
- `hanzo mcp serve` - Start MCP server
- `hanzo net status` - Network status
- `hanzo dev` - Development tools

## Development

Each language implementation maintains feature parity while leveraging language-specific strengths:

- **Python**: ML/AI operations, data processing
- **JavaScript**: Web integration, browser compatibility
- **Rust**: System operations, performance-critical paths
- **Go**: Network operations, blockchain integration