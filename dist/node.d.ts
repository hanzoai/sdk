/**
 * Node management
 */
export interface NodeStatus {
    running: boolean;
    port: number;
    models: string[];
}
export declare class Node {
    private port;
    private running;
    private models;
    constructor(port?: number);
    start(): Promise<void>;
    stop(): Promise<void>;
    getStatus(): NodeStatus;
    loadModel(model: string): Promise<void>;
    listModels(): string[];
}
//# sourceMappingURL=node.d.ts.map