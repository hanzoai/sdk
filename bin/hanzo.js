#!/usr/bin/env node

/**
 * Hanzo CLI - Universal entry point
 * Attempts to use Rust implementation for performance,
 * falls back to JavaScript implementation if unavailable
 */

import { execSync } from 'node:child_process';
import 'dotenv/config';
const args = process.argv.slice(2);

// For certain commands, prefer Python implementation
const pythonPreferredCommands = ['mcp', 'agent', 'repl'];
const shouldUsePython = pythonPreferredCommands.some(cmd => args.includes(cmd));

if (shouldUsePython) {
  try {
    // Try to use Python implementation for these commands
    execSync(`python -m hanzo.cli ${args.join(' ')}`, { 
      stdio: 'inherit',
      env: process.env 
    });
    process.exit(0);
  } catch (error) {
    // Fall back to JS implementation
    console.debug('Python implementation not available, using JavaScript');
  }
}

// Load the compiled CLI.
await import('../dist/cli.js');