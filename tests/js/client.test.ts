import { HanzoClient } from '../../src/js/client';

describe('HanzoClient', () => {
  it('should create client instance', () => {
    const client = new HanzoClient();
    expect(client).toBeDefined();
  });

  it('should create client with API key', () => {
    const client = new HanzoClient('test-api-key');
    expect(client).toBeDefined();
  });

  it('should have chat completion method', () => {
    const client = new HanzoClient();
    expect(client.chatCompletion).toBeDefined();
    expect(typeof client.chatCompletion).toBe('function');
  });

  it('should have list models method', () => {
    const client = new HanzoClient();
    expect(client.listModels).toBeDefined();
    expect(typeof client.listModels).toBe('function');
  });
});