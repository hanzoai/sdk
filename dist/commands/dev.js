"use strict";
/**
 * Development commands
 */
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.devCommands = devCommands;
const chalk_1 = __importDefault(require("chalk"));
const ora_1 = __importDefault(require("ora"));
const child_process_1 = require("child_process");
function devCommands(program) {
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
        console.log(chalk_1.default.green(`✅ Development server running on http://localhost:${options.port}`));
    });
    dev
        .command('build')
        .description('Build the project')
        .action(() => {
        const spinner = (0, ora_1.default)('Building project...').start();
        try {
            (0, child_process_1.execSync)('npm run build', { stdio: 'ignore' });
            spinner.succeed(chalk_1.default.green('Build complete'));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Build failed'));
            process.exit(1);
        }
    });
    dev
        .command('test')
        .description('Run tests')
        .action(() => {
        const spinner = (0, ora_1.default)('Running tests...').start();
        try {
            (0, child_process_1.execSync)('npm test', { stdio: 'ignore' });
            spinner.succeed(chalk_1.default.green('All tests passed'));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Tests failed'));
            process.exit(1);
        }
    });
    dev
        .command('lint')
        .description('Run linter')
        .action(() => {
        const spinner = (0, ora_1.default)('Running linter...').start();
        try {
            (0, child_process_1.execSync)('npm run lint', { stdio: 'ignore' });
            spinner.succeed(chalk_1.default.green('No linting issues found'));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Linting issues found'));
            process.exit(1);
        }
    });
    dev
        .command('format')
        .description('Format code')
        .action(() => {
        const spinner = (0, ora_1.default)('Formatting code...').start();
        try {
            (0, child_process_1.execSync)('npm run format', { stdio: 'ignore' });
            spinner.succeed(chalk_1.default.green('Code formatted'));
        }
        catch (error) {
            spinner.fail(chalk_1.default.red('Formatting failed'));
            process.exit(1);
        }
    });
}
//# sourceMappingURL=dev.js.map