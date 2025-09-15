"use strict";
/**
 * Node management
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.Node = void 0;
class Node {
    constructor(port = 4000) {
        this.port = port;
        this.running = false;
        this.models = [];
    }
    async start() {
        this.running = true;
        console.log(`Node started on port ${this.port}`);
    }
    async stop() {
        this.running = false;
        console.log('Node stopped');
    }
    getStatus() {
        return {
            running: this.running,
            port: this.port,
            models: this.models
        };
    }
    async loadModel(model) {
        this.models.push(model);
        console.log(`Model ${model} loaded`);
    }
    listModels() {
        return this.models;
    }
}
exports.Node = Node;
//# sourceMappingURL=node.js.map