/**
 * Hanzo AI SDK - Unified TypeScript implementation
 */

export * from './client';
export * from './agent';
export * from './mcp';
export * from './node';
export * from './cli';

// Try to load Rust bindings if available
let rustBindings: any = null;
try {
  rustBindings = require('../../rs/target/release/hanzo.node');
} catch (e) {
  // Rust bindings not available, fall back to JS implementation
  console.debug('Using JavaScript implementation (Rust bindings not found)');
}

export const useRust = () => rustBindings !== null;
export const getRustBindings = () => rustBindings;