/**
 * Hanzo AI SDK - Unified TypeScript implementation
 */

import { createRequire } from 'node:module';

export * from './client.js';
export * from './agent.js';
export * from './mcp.js';
export * from './node.js';
export * from './cli.js';

// Try to load Rust bindings if available. A native .node addon has no ESM
// loader, so it is required — createRequire is how an ES module does that.
const requireAddon = createRequire(import.meta.url);

let rustBindings: any = null;
try {
  rustBindings = requireAddon('../../rs/target/release/hanzo.node');
} catch (e) {
  // Rust bindings not available, fall back to JS implementation
  console.debug('Using JavaScript implementation (Rust bindings not found)');
}

export const useRust = () => rustBindings !== null;
export const getRustBindings = () => rustBindings;