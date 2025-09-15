/**
 * Configuration commands
 */

import { Command } from 'commander';
import chalk from 'chalk';
import * as fs from 'fs';
import * as path from 'path';

export function configCommands(program: Command) {
  const config = program
    .command('config')
    .description('Configuration management');

  config
    .command('set <key> <value>')
    .description('Set a configuration value')
    .action((key, value) => {
      const configPath = path.join(process.env.HOME || '', '.hanzo', 'config.json');
      const configDir = path.dirname(configPath);
      fs.mkdirSync(configDir, { recursive: true });
      
      let configData: any = {};
      if (fs.existsSync(configPath)) {
        configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
      }
      
      configData[key] = value;
      fs.writeFileSync(configPath, JSON.stringify(configData, null, 2));
      console.log(chalk.green(`✅ Set ${key} = ${value}`));
    });

  config
    .command('get <key>')
    .description('Get a configuration value')
    .action((key) => {
      const configPath = path.join(process.env.HOME || '', '.hanzo', 'config.json');
      if (!fs.existsSync(configPath)) {
        console.error(chalk.red(`Configuration key '${key}' not found`));
        process.exit(1);
      }
      
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
      const value = configData[key];
      if (value === undefined) {
        console.error(chalk.red(`Configuration key '${key}' not found`));
        process.exit(1);
      }
      
      console.log(value);
    });

  config
    .command('list')
    .description('List all configuration values')
    .action(() => {
      const configPath = path.join(process.env.HOME || '', '.hanzo', 'config.json');
      if (!fs.existsSync(configPath)) {
        console.log('No configuration found');
        return;
      }
      
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
      if (Object.keys(configData).length === 0) {
        console.log('No configuration found');
        return;
      }
      
      console.log('Configuration:');
      Object.entries(configData).forEach(([key, value]) => {
        console.log(`  ${key}: ${value}`);
      });
    });

  config
    .command('unset <key>')
    .description('Remove a configuration value')
    .action((key) => {
      const configPath = path.join(process.env.HOME || '', '.hanzo', 'config.json');
      if (!fs.existsSync(configPath)) {
        console.error(chalk.red(`Configuration key '${key}' not found`));
        process.exit(1);
      }
      
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
      if (!(key in configData)) {
        console.error(chalk.red(`Configuration key '${key}' not found`));
        process.exit(1);
      }
      
      delete configData[key];
      fs.writeFileSync(configPath, JSON.stringify(configData, null, 2));
      console.log(chalk.green(`✅ Removed ${key}`));
    });
}