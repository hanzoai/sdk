/**
 * Agent framework
 */
export class Agent {
    constructor(options) {
        this.name = options.name;
        this.model = options.model || 'gpt-4';
        this.systemPrompt = options.systemPrompt;
        this.tools = options.tools || [];
    }
    async run(task) {
        // Implementation would execute the agent
        return `Task '${task}' completed by ${this.name}`;
    }
    addTool(tool) {
        this.tools.push(tool);
    }
    toJSON() {
        return {
            name: this.name,
            model: this.model,
            systemPrompt: this.systemPrompt,
            tools: this.tools
        };
    }
}
//# sourceMappingURL=agent.js.map