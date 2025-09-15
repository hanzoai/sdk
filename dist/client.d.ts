/**
 * Hanzo API client
 */
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
export declare class HanzoClient {
    private client;
    constructor(apiKey?: string, baseURL?: string);
    chatCompletion(options: ChatCompletionOptions): Promise<any>;
    listModels(): Promise<string[]>;
}
//# sourceMappingURL=client.d.ts.map