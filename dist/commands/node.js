"use strict";
/**
 * Node management commands
 */
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.nodeCommands = nodeCommands;
const child_process_1 = require("child_process");
const chalk_1 = __importDefault(require("chalk"));
const ora_1 = __importDefault(require("ora"));
const index_1 = require("../index");
function nodeCommands(program) {
    const node = program
        .command('node')
        .description('Manage Hanzo AI nodes');
    node
        .command('start')
        .description('Start a local AI node')
        .option('-p, --port <port>', 'Port to run the node on', '4000')
        .option('-c, --config <path>', 'Configuration file')
        .option('-d, --detach', 'Run in background')
        .action(async (options) => {
        const spinner = (0, ora_1.default)('Starting Hanzo node...').start();
        try {
            // Try Rust implementation first
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.startNode(options);
            }
            else {
                // Fallback to JS implementation
                const args = ['start', '--port', options.port];
                if (options.config)
                    args.push('--config', options.config);
                if (options.detach)
                    args.push('--detach');
                if (options.detach) {
                    (0, child_process_1.spawn)('hanzo-node', args, {
                        detached: true,
                        stdio: 'ignore'
                    }).unref();
                }
                else {
                    (0, child_process_1.execSync)(`hanzo-node ${args.join(' ')}`, { stdio: 'inherit' });
                }
            }
            spinner.succeed(chalk_1.default.green(`Node started on port ${options.port}`));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to start node'));
            console.error(error);
            process.exit(1);
        }
    });
    node
        .command('stop')
        .description('Stop the running node')
        .action(async () => {
        const spinner = (0, ora_1.default)('Stopping Hanzo node...').start();
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.stopNode();
            }
            else {
                (0, child_process_1.execSync)('hanzo-node stop', { stdio: 'inherit' });
            }
            spinner.succeed(chalk_1.default.green('Node stopped successfully'));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to stop node'));
            console.error(error);
            process.exit(1);
        }
    });
    node
        .command('status')
        .description('Check node status')
        .action(async () => {
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                const status = await rust.getNodeStatus();
                console.log(status);
            }
            else {
                (0, child_process_1.execSync)('hanzo-node status', { stdio: 'inherit' });
            }
        }
        catch (error) {
            console.log(chalk_1.default.yellow('Node is not running'));
        }
    });
    node
        .command('load <model>')
        .description('Load a model into the node')
        .action(async (model) => {
        const spinner = (0, ora_1.default)(`Loading model: ${model}`).start();
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.loadModel(model);
            }
            else {
                (0, child_process_1.execSync)(`hanzo-node load ${model}`, { stdio: 'inherit' });
            }
            spinner.succeed(chalk_1.default.green(`Model ${model} loaded successfully`));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to load model'));
            console.error(error);
            process.exit(1);
        }
    });
    node
        .command('list')
        .description('List available models')
        .action(async () => {
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                const models = await rust.listModels();
                console.log(models);
            }
            else {
                (0, child_process_1.execSync)('hanzo-node list', { stdio: 'inherit' });
            }
        }
        catch (error) {
            console.error(chalk_1.default.red('Failed to list models'));
            process.exit(1);
        }
    });
}
//# sourceMappingURL=node.js.map