/**
 * Agent management commands
 */
import chalk from 'chalk';
import ora from 'ora';
import * as fs from 'fs';
import * as path from 'path';
import { getRustBindings, useRust } from '../index.js';
export function agentCommands(program) {
    const agent = program
        .command('agent')
        .description('Manage AI agents');
    agent
        .command('create <name>')
        .description('Create a new agent')
        .option('-t, --type <type>', 'Agent type', 'general')
        .option('-m, --model <model>', 'Model to use', 'gpt-4')
        .action(async (name, options) => {
        const spinner = ora(`Creating agent '${name}'...`).start();
        try {
            if (useRust()) {
                const rust = getRustBindings();
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
            spinner.succeed(chalk.green(`Agent '${name}' created successfully`));
        }
        catch (error) {
            spinner.fail(chalk.red('Failed to create agent'));
            console.error(error);
            process.exit(1);
        }
    });
    agent
        .command('list')
        .description('List all agents')
        .action(async () => {
        try {
            if (useRust()) {
                const rust = getRustBindings();
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
            console.error(chalk.red('Failed to list agents'));
            process.exit(1);
        }
    });
    agent
        .command('run <name> <task>')
        .description('Run an agent with a task')
        .option('-a, --async', 'Run asynchronously')
        .action(async (name, task, options) => {
        const spinner = ora(`Running agent '${name}'...`).start();
        try {
            if (useRust()) {
                const rust = getRustBindings();
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
            spinner.succeed(chalk.green(`Task completed by agent '${name}'`));
        }
        catch (error) {
            spinner.fail(chalk.red('Failed to run agent'));
            console.error(error);
            process.exit(1);
        }
    });
}
//# sourceMappingURL=agent.js.map