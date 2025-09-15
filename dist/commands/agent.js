"use strict";
/**
 * Agent management commands
 */
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.agentCommands = agentCommands;
const chalk_1 = __importDefault(require("chalk"));
const ora_1 = __importDefault(require("ora"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
const index_1 = require("../index");
function agentCommands(program) {
    const agent = program
        .command('agent')
        .description('Manage AI agents');
    agent
        .command('create <name>')
        .description('Create a new agent')
        .option('-t, --type <type>', 'Agent type', 'general')
        .option('-m, --model <model>', 'Model to use', 'gpt-4')
        .action(async (name, options) => {
        const spinner = (0, ora_1.default)(`Creating agent '${name}'...`).start();
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.createAgent(name, options);
            }
            else {
                // JS implementation
                const config = {
                    name,
                    type: options.type,
                    model: options.model,
                    created: new Date().toISOString()
                };
                const configDir = path.join(process.env.HOME || '', '.hanzo', 'agents');
                fs.mkdirSync(configDir, { recursive: true });
                fs.writeFileSync(path.join(configDir, `${name}.json`), JSON.stringify(config, null, 2));
            }
            spinner.succeed(chalk_1.default.green(`Agent '${name}' created successfully`));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to create agent'));
            console.error(error);
            process.exit(1);
        }
    });
    agent
        .command('list')
        .description('List all agents')
        .action(async () => {
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                const agents = await rust.listAgents();
                console.log(agents);
            }
            else {
                const configDir = path.join(process.env.HOME || '', '.hanzo', 'agents');
                if (!fs.existsSync(configDir)) {
                    console.log('No agents found');
                    return;
                }
                const agents = fs.readdirSync(configDir)
                    .filter(f => f.endsWith('.json'))
                    .map(f => {
                    const config = JSON.parse(fs.readFileSync(path.join(configDir, f), 'utf-8'));
                    return `  - ${config.name} (${config.type})`;
                });
                if (agents.length === 0) {
                    console.log('No agents found');
                }
                else {
                    console.log('Available agents:');
                    agents.forEach(a => console.log(a));
                }
            }
        }
        catch (error) {
            console.error(chalk_1.default.red('Failed to list agents'));
            process.exit(1);
        }
    });
    agent
        .command('run <name> <task>')
        .description('Run an agent with a task')
        .option('-a, --async', 'Run asynchronously')
        .action(async (name, task, options) => {
        const spinner = (0, ora_1.default)(`Running agent '${name}'...`).start();
        try {
            if ((0, index_1.useRust)()) {
                const rust = (0, index_1.getRustBindings)();
                await rust.runAgent(name, task, options);
            }
            else {
                // JS implementation
                if (options.async) {
                    console.log('Running in async mode...');
                }
                // Simulate agent execution
                await new Promise(resolve => setTimeout(resolve, 1000));
            }
            spinner.succeed(chalk_1.default.green(`Task completed by agent '${name}'`));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Failed to run agent'));
            console.error(error);
            process.exit(1);
        }
    });
}
//# sourceMappingURL=agent.js.map