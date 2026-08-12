/**
 * Authentication commands
 */
import chalk from 'chalk';
import * as fs from 'fs';
import * as path from 'path';
import * as readline from 'readline';
export function authCommands(program) {
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
        console.log(chalk.green('✅ Successfully logged in'));
    });
    auth
        .command('logout')
        .description('Logout from Hanzo')
        .action(() => {
        const configPath = path.join(process.env.HOME || '', '.hanzo', 'auth.json');
        if (fs.existsSync(configPath)) {
            fs.unlinkSync(configPath);
        }
        console.log(chalk.green('✅ Successfully logged out'));
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
            console.error(chalk.red('Not logged in'));
            process.exit(1);
        }
        const config = JSON.parse(fs.readFileSync(configPath, 'utf-8'));
        console.log(config.token);
    });
}
//# sourceMappingURL=auth.js.map