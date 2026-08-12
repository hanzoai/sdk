#!/usr/bin/env node
/**
 * Hanzo CLI - TypeScript implementation
 */

import { Command } from 'commander';
import chalk from 'chalk';
import { nodeCommands } from './commands/node.js';
import { agentCommands } from './commands/agent.js';
import { mcpCommands } from './commands/mcp.js';
import { netCommands } from './commands/net.js';
import { devCommands } from './commands/dev.js';
import { authCommands } from './commands/auth.js';
import { configCommands } from './commands/config.js';

const program = new Command();

program
  .name('hanzo')
  .description('Hanzo AI SDK - Unified CLI for all Hanzo services')
  .version('0.1.0');

// Add command groups
nodeCommands(program);
agentCommands(program);
mcpCommands(program);
netCommands(program);
devCommands(program);
authCommands(program);
configCommands(program);

// Parse arguments
program.parse(process.argv);

// Show help if no command provided
if (!process.argv.slice(2).length) {
  program.outputHelp();
}