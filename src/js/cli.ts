#!/usr/bin/env node
/**
 * Hanzo CLI - TypeScript implementation
 */

import { Command } from 'commander';
import chalk from 'chalk';
import { nodeCommands } from './commands/node';
import { agentCommands } from './commands/agent';
import { mcpCommands } from './commands/mcp';
import { netCommands } from './commands/net';
import { devCommands } from './commands/dev';
import { authCommands } from './commands/auth';
import { configCommands } from './commands/config';

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