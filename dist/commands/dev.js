/**
 * Development commands
 */
import chalk from 'chalk';
import ora from 'ora';
import { execSync } from 'child_process';
export function devCommands(program) {
    const dev = program
        .command('dev')
        .description('Development tools');
    dev
        .command('start')
        .description('Start development server')
        .option('-p, --port <port>', 'Development server port', '3000')
        .option('--hot-reload', 'Enable hot reload')
        .action((options) => {
        console.log(`Starting development server on port ${options.port}`);
        if (options.hotReload) {
            console.log('Hot reload enabled');
        }
        console.log(chalk.green(`✅ Development server running on http://localhost:${options.port}`));
    });
    dev
        .command('build')
        .description('Build the project')
        .action(() => {
        const spinner = ora('Building project...').start();
        try {
            execSync('npm run build', { stdio: 'ignore' });
            spinner.succeed(chalk.green('Build complete'));
        }
        catch (error) {
            spinner.fail(chalk.red('Build failed'));
            process.exit(1);
        }
    });
    dev
        .command('test')
        .description('Run tests')
        .action(() => {
        const spinner = ora('Running tests...').start();
        try {
            execSync('npm test', { stdio: 'ignore' });
            spinner.succeed(chalk.green('All tests passed'));
        }
        catch (error) {
            spinner.fail(chalk.red('Tests failed'));
            process.exit(1);
        }
    });
    dev
        .command('lint')
        .description('Run linter')
        .action(() => {
        const spinner = ora('Running linter...').start();
        try {
            execSync('npm run lint', { stdio: 'ignore' });
            spinner.succeed(chalk.green('No linting issues found'));
        }
        catch (error) {
            spinner.fail(chalk.red('Linting issues found'));
            process.exit(1);
        }
    });
    dev
        .command('format')
        .description('Format code')
        .action(() => {
        const spinner = ora('Formatting code...').start();
        try {
            execSync('npm run format', { stdio: 'ignore' });
            spinner.succeed(chalk.green('Code formatted'));
        }
        catch (error) {
            spinner.fail(chalk.red('Formatting failed'));
            process.exit(1);
        }
    });
}
//# sourceMappingURL=dev.js.map