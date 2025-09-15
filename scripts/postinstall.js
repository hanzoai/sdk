#!/usr/bin/env node

console.log('Hanzo SDK installed successfully!');
console.log('Run "hanzo --help" to get started.');

// Try to build Rust bindings if cargo is available
const { execSync } = require('child_process');

try {
  execSync('cargo --version', { stdio: 'ignore' });
  console.log('Building Rust bindings...');
  execSync('cargo build --release --manifest-path=src/rs/Cargo.toml', { stdio: 'ignore' });
  console.log('Rust bindings built successfully.');
} catch (e) {
  console.log('Rust not available, using JavaScript fallback.');
}
