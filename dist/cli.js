#!/usr/bin/env node
"use strict";
/**
 * Hanzo CLI - TypeScript implementation
 */
Object.defineProperty(exports, "__esModule", { value: true });
const commander_1 = require("commander");
const node_1 = require("./commands/node");
const agent_1 = require("./commands/agent");
const mcp_1 = require("./commands/mcp");
const net_1 = require("./commands/net");
const dev_1 = require("./commands/dev");
const auth_1 = require("./commands/auth");
const config_1 = require("./commands/config");
const program = new commander_1.Command();
program
    .name('hanzo')
    .description('Hanzo AI SDK - Unified CLI for all Hanzo services')
    .version('0.1.0');
// Add command groups
(0, node_1.nodeCommands)(program);
(0, agent_1.agentCommands)(program);
(0, mcp_1.mcpCommands)(program);
(0, net_1.netCommands)(program);
(0, dev_1.devCommands)(program);
(0, auth_1.authCommands)(program);
(0, config_1.configCommands)(program);
// Parse arguments
program.parse(process.argv);
// Show help if no command provided
if (!process.argv.slice(2).length) {
    program.outputHelp();
}
//# sourceMappingURL=cli.js.map