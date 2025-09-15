"use strict";
/**
 * Authentication commands
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
exports.authCommands = authCommands;
const chalk_1 = __importDefault(require("chalk"));
const fs = __importStar(require("fs"));
const path = __importStar(require("path"));
const readline = __importStar(require("readline"));
function authCommands(program) {
    const auth = program
        .command('auth')
        .description('Authentication management');
    auth
        .command('login')
        .description('Login to Hanzo')
        .action(async () => {
        const rl = readline.createInterface({
            input: process.stdin,
            output: process.stdout
        });
        const email = await new Promise((resolve) => {
            rl.question('Email: ', resolve);
        });
        const password = await new Promise((resolve) => {
            rl.question('Password: ', resolve);
        });
        rl.close();
        console.log(`Logging in as ${email}...`);
        // Simulate authentication
        const token = 'sk-hanzo-example-token';
        const configDir = path.join(process.env.HOME || '', '.hanzo');
        fs.mkdirSync(configDir, { recursive: true });
        fs.writeFileSync(path.join(configDir, 'auth.json'), JSON.stringify({ email, token }, null, 2));
        console.log(chalk_1.default.green('✅ Successfully logged in'));
    });
    auth
        .command('logout')
        .description('Logout from Hanzo')
        .action(() => {
        const configPath = path.join(process.env.HOME || '', '.hanzo', 'auth.json');
        if (fs.existsSync(configPath)) {
            fs.unlinkSync(configPath);
        }
        console.log(chalk_1.default.green('✅ Successfully logged out'));
    });
    auth
        .command('status')
        .description('Check authentication status')
        .action(() => {
        const configPath = path.join(process.env.HOME || '', '.hanzo', 'auth.json');
        if (!fs.existsSync(configPath)) {
            console.log('Not logged in');
            return;
        }
        const config = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
        console.log(`Logged in as: ${config.email}`);
    });
    auth
        .command('token')
        .description('Display current auth token')
        .action(() => {
        const configPath = path.join(process.env.HOME || '', '.hanzo', 'auth.json');
        if (!fs.existsSync(configPath)) {
            console.error(chalk_1.default.red('Not logged in'));
            process.exit(1);
        }
        const config = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
        console.log(config.token);
    });
}
//# sourceMappingURL=auth.js.map