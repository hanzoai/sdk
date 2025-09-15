#!/usr/bin/env node
/**
 * Post-install script to attempt building Rust bindings
 */

const { execSync } = require('child_process');
const path = require('path');
const fs = require('fs');

console.log('Checking for Rust installation...');

try {
  execSync('cargo --version', { stdio: 'ignore' });
  console.log('Rust found, attempting to build native bindings...');
  
  const rustDir = path.join(__dirname, '..', 'src', 'rs');
  if (fs.existsSync(rustDir)) {
    try {
      execSync('cargo build --release', { 
        cwd: rustDir,
        stdio: 'inherit'
      });
      console.log('✅ Rust bindings built successfully');
    } catch (e) {
      console.log('⚠️  Rust build failed, using JavaScript fallback');
    }
  }
} catch (e) {
  console.log('ℹ️  Rust not installed, using JavaScript implementation');
}

console.log('Post-install complete');