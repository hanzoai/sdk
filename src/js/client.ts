/**
 * Hanzo API client
 */

import axios, { AxiosInstance } from 'axios';

export interface ChatMessage {
  role: 'system' | 'user' | 'assistant';
  content: string;
}

export interface ChatCompletionOptions {
  messages: ChatMessage[];
  model?: string;
  temperature?: number;
  max_tokens?: number;
}

export class HanzoClient {
  private client: AxiosInstance;

  constructor(apiKey?: string, baseURL: string = 'https://api.hanzo.ai') {
    this.client = axios.create({
      baseURL,
      headers: apiKey ? { Authorization: `Bearer ${apiKey}` } : {}
    });
  }

  async chatCompletion(options: ChatCompletionOptions): Promise<any> {
    const response = await this.client.post('/v1/chat/completions', {
      model: options.model || 'gpt-4',
      messages: options.messages,
      temperature: options.temperature,
      max_tokens: options.max_tokens
    });
    return response.data;
  }

  async listModels(): Promise<string[]> {
    const response = await this.client.get('/v1/models');
    return response.data.data.map((m: any) => m.id);
  }
}