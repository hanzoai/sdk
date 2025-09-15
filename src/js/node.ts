/**
 * Node management
 */

export interface NodeStatus {
  running: boolean;
  port: number;
  models: string[];
}

export class Node {
  private port: number;
  private running: boolean;
  private models: string[];

  constructor(port: number = 4000) {
    this.port = port;
    this.running = false;
    this.models = [];
  }

  async start(): Promise<void> {
    this.running = true;
    console.log(`Node started on port ${this.port}`);
  }

  async stop(): Promise<void> {
    this.running = false;
    console.log('Node stopped');
  }

  getStatus(): NodeStatus {
    return {
      running: this.running,
      port: this.port,
      models: this.models
    };
  }

  async loadModel(model: string): Promise<void> {
    this.models.push(model);
    console.log(`Model ${model} loaded`);
  }

  listModels(): string[] {
    return this.models;
  }
}