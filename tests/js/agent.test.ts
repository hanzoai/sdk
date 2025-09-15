import { Agent } from '../../src/js/agent';

describe('Agent', () => {
  it('should create agent instance', () => {
    const agent = new Agent({ name: 'test-agent' });
    expect(agent).toBeDefined();
    expect(agent.name).toBe('test-agent');
  });

  it('should use default model', () => {
    const agent = new Agent({ name: 'test-agent' });
    expect(agent.model).toBe('gpt-4');
  });

  it('should add tools', () => {
    const agent = new Agent({ name: 'test-agent' });
    agent.addTool('filesystem');
    agent.addTool('shell');
    expect(agent.tools).toHaveLength(2);
    expect(agent.tools).toContain('filesystem');
  });

  it('should serialize to JSON', () => {
    const agent = new Agent({ 
      name: 'test-agent',
      model: 'claude-3',
      tools: ['filesystem']
    });
    const json = agent.toJSON();
    expect(json.name).toBe('test-agent');
    expect(json.model).toBe('claude-3');
    expect(json.tools).toContain('filesystem');
  });

  it('should run tasks', async () => {
    const agent = new Agent({ name: 'test-agent' });
    const result = await agent.run('test task');
    expect(result).toContain('test task');
    expect(result).toContain('test-agent');
  });
});