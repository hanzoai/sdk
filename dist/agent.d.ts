/**
 * Agent framework
 */
export interface AgentOptions {
    name: string;
    model?: string;
    systemPrompt?: string;
    tools?: string[];
}
export declare class Agent {
    name: string;
    model: string;
    systemPrompt?: string;
    tools: string[];
    constructor(options: AgentOptions);
    run(task: string): Promise<string>;
    addTool(tool: string): void;
    toJSON(): any;
}
//# sourceMappingURL=agent.d.ts.map