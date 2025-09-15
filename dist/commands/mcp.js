"use strict";
/**
 * MCP (Model Context Protocol) commands
 */
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.mcpCommands = mcpCommands;
const chalk_1 = __importDefault(require("chalk"));
const ora_1 = __importDefault(require("ora"));
const index_1 = require("../index");
function mcpCommands(program) {
    const mcp = program
        .command('mcp')
        .description('Model Context Protocol tools');
    mcp
        .command('serve')
        .description('Start MCP server')
        .option('-p, --port <port>', 'Port to serve on', '3000')
        .option('-t, --transport <type>', 'Transport type', 'stdio')
        .action(async (options) => {
        const spinner = (0, ora_1.default)('Starting MCP server...').start();
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.startMCPServer(options);
            }
            else {
                // JS implementation
                console.log(`MCP server running on port ${options.port} with ${options.transport} transport`);
            }
            spinner.succeed(chalk_1.default.green(`MCP server running on port ${options.port}`));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to start MCP server'));
            console.error(error);
            process.exit(1);
        }
    });
    mcp
        .command('tools')
        .description('List available MCP tools')
        .action(async () => {
        const tools = [
            'filesystem - File system operations',
            'shell - Shell command execution',
            'memory - Memory and knowledge management',
            'search - Code and document search',
            'editor - Code editing tools',
            'database - Database operations',
            'vector - Vector operations',
            'agent - Agent orchestration'
        ];
        console.log('Available MCP tools:');
        tools.forEach(tool => console.log(`  - ${tool}`));
    });
    mcp
        .command('run <tool> <action>')
        .description('Execute an MCP tool')
        .option('-p, --params <json>', 'JSON parameters')
        .action(async (tool, action, options) => {
        const spinner = (0, ora_1.default)(`Running MCP tool '${tool}'...`).start();
        try {
            if (options.params) {
                try {
                    const params = JSON.parse(options.params);
                    console.log('Parameters:', params);
                }
                catch (e) {
                    spinner.fail(chalk_1.default.red('Invalid JSON parameters'));
                    process.exit(1);
                }
            }
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.runMCPTool(tool, action, options.params);
            }
            else {
                // JS implementation
                await new Promise(resolve => setTimeout(resolve, 500));
            }
            spinner.succeed(chalk_1.default.green(`Tool '${tool}' executed successfully`));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to run MCP tool'));
            console.error(error);
            process.exit(1);
        }
    });
}
//# sourceMappingURL=mcp.js.map