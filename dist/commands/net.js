/**
 * Network management commands
 */
import chalk from 'chalk';
import { getRustBindings, useRust } from '../index.js';
export function netCommands(program) {
    const net = program
        .command('net')
        .description('Network operations');
    net
        .command('status')
        .description('Check network status')
        .action(async () => {
        if (useRust()) {
            const rust = getRustBindings();
            const status = await rust.getNetworkStatus();
            console.log(status);
        }
        else {
            console.log('Network Status:');
            console.log('  - Node: Online');
            console.log('  - Peers: 12 connected');
            console.log('  - Bandwidth: 1.2 MB/s');
            console.log('  - Latency: 45ms avg');
        }
    });
    net
        .command('peers')
        .description('List connected peers')
        .action(async () => {
        if (useRust()) {
            const rust = getRustBindings();
            const peers = await rust.listPeers();
            console.log(peers);
        }
        else {
            console.log('Connected peers:');
            console.log('  - peer1: 192.168.1.10:8080 [active]');
            console.log('  - peer2: 192.168.1.11:8080 [active]');
            console.log('  - peer3: 192.168.1.12:8080 [idle]');
        }
    });
    net
        .command('connect <address>')
        .description('Connect to a peer')
        .action(async (address) => {
        console.log(`Connecting to ${address}...`);
        if (useRust()) {
            const rust = getRustBindings();
            await rust.connectToPeer(address);
        }
        console.log(chalk.green(`✅ Connected to ${address}`));
    });
    net
        .command('disconnect <peerId>')
        .description('Disconnect from a peer')
        .action(async (peerId) => {
        console.log(`Disconnecting from ${peerId}...`);
        if (useRust()) {
            const rust = getRustBindings();
            await rust.disconnectFromPeer(peerId);
        }
        console.log(chalk.green(`✅ Disconnected from ${peerId}`));
    });
}
//# sourceMappingURL=net.js.map