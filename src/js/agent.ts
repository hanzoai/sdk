/**
 * Agent framework
 */

export interface AgentOptions {
  name: string;
  model?: string;
  systemPrompt?: string;
  tools?: string[];
}

export class Agent {
  name: string;
  model: string;
  systemPrompt?: string;
  tools: string[];

  constructor(options: AgentOptions) {
    this.name = options.name;
    this.model = options.model || 'gpt-4';
    this.systemPrompt = options.systemPrompt;
    this.tools = options.tools || [];
  }

  async run(task: string): Promise<string> {
    // Implementation would execute the agent
    return `Task '${task}' completed by ${this.name}`;
  }

  addTool(tool: string): void {
    this.tools.push(tool);
  }

  toJSON(): any {
    return {
      name: this.name,
      model: this.model,
      systemPrompt: this.systemPrompt,
      tools: this.tools
    };
  }
}