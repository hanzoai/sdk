/**
 * Node management commands
 */

import { Command } from 'commander';
import { spawn, execSync } from 'child_process';
import chalk from 'chalk';
import ora from 'ora';
import { getRustBindings, useRust } from '../index';

export function nodeCommands(program: Command) {
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
      const spinner = ora('Starting Hanzo node...').start();
      
      try {
        // Try Rust implementation first
        if (useRust()) {
          const rust = getRustBindings();
          await rust.startNode(options);
        } else {
          // Fallback to JS implementation
          const args = ['start', '--port', options.port];
          if (options.config) args.push('--config', options.config);
          if (options.detach) args.push('--detach');
          
          if (options.detach) {
            spawn('hanzo-node', args, { 
              detached: true, 
              stdio: 'ignore' 
            }).unref();
          } else {
            execSync(`hanzo-node ${args.join(' ')}`, { stdio: 'inherit' });
          }
        }
        
        spinner.succeed(chalk.green(`Node started on port ${options.port}`));
      } catch (error) {
        spinner.fail(chalk.red('Failed to start node'));
        console.error(error);
        process.exit(1);
      }
    });

  node
    .command('stop')
    .description('Stop the running node')
    .action(async () => {
      const spinner = ora('Stopping Hanzo node...').start();
      
      try {
        if (useRust()) {
          const rust = getRustBindings();
          await rust.stopNode();
        } else {
          execSync('hanzo-node stop', { stdio: 'inherit' });
        }
        
        spinner.succeed(chalk.green('Node stopped successfully'));
      } catch (error) {
        spinner.fail(chalk.red('Failed to stop node'));
        console.error(error);
        process.exit(1);
      }
    });

  node
    .command('status')
    .description('Check node status')
    .action(async () => {
      try {
        if (useRust()) {
          const rust = getRustBindings();
          const status = await rust.getNodeStatus();
          console.log(status);
        } else {
          execSync('hanzo-node status', { stdio: 'inherit' });
        }
      } catch (error) {
        console.log(chalk.yellow('Node is not running'));
      }
    });

  node
    .command('load <model>')
    .description('Load a model into the node')
    .action(async (model) => {
      const spinner = ora(`Loading model: ${model}`).start();
      
      try {
        if (useRust()) {
          const rust = getRustBindings();
          await rust.loadModel(model);
        } else {
          execSync(`hanzo-node load ${model}`, { stdio: 'inherit' });
        }
        
        spinner.succeed(chalk.green(`Model ${model} loaded successfully`));
      } catch (error) {
        spinner.fail(chalk.red('Failed to load model'));
        console.error(error);
        process.exit(1);
      }
    });

  node
    .command('list')
    .description('List available models')
    .action(async () => {
      try {
        if (useRust()) {
          const rust = getRustBindings();
          const models = await rust.listModels();
          console.log(models);
        } else {
          execSync('hanzo-node list', { stdio: 'inherit' });
        }
      } catch (error) {
        console.error(chalk.red('Failed to list models'));
        process.exit(1);
      }
    });
}