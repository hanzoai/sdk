"use strict";
/**
 * Configuration commands
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
exports.configCommands = configCommands;
const chalk_1 = __importDefault(require("chalk"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
function configCommands(program) {
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
        let configData = {};
        if (fs.existsSync(configPath)) {
            configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
        }
        configData[key] = value;
        fs.writeFileSync(configPath, JSON.stringify(configData, null, 2));
        console.log(chalk_1.default.green(`✅ Set ${key} = ${value}`));
    });
    config
        .command('get <key>')
        .description('Get a configuration value')
        .action((key) => {
        const configPath = path.join(process.env.HOME || '', '.hanzo', 'config.json');
        if (!fs.existsSync(configPath)) {
            console.error(chalk_1.default.red(`Configuration key '${key}' not found`));
            process.exit(1);
        }
        const configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
        const value = configData[key];
        if (value === undefined) {
            console.error(chalk_1.default.red(`Configuration key '${key}' not found`));
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
            console.error(chalk_1.default.red(`Configuration key '${key}' not found`));
            process.exit(1);
        }
        const configData = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
        if (!(key in configData)) {
            console.error(chalk_1.default.red(`Configuration key '${key}' not found`));
            process.exit(1);
        }
        delete configData[key];
        fs.writeFileSync(configPath, JSON.stringify(configData, null, 2));
        console.log(chalk_1.default.green(`✅ Removed ${key}`));
    });
}
//# sourceMappingURL=config.js.map