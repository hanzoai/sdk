/**
 * Hanzo API client
 */
import axios from 'axios';
export class HanzoClient {
    constructor(apiKey, baseURL = 'https://api.hanzo.ai') {
        this.client = axios.create({
            baseURL,
            headers: apiKey ? { Authorization: `Bearer ${apiKey}` } : {}
        });
    }
    async chatCompletion(options) {
        const response = await this.client.post('/v1/chat/completions', {
            model: options.model || 'gpt-4',
            messages: options.messages,
            temperature: options.temperature,
            max_tokens: options.max_tokens
        });
        return response.data;
    }
    async listModels() {
        const response = await this.client.get('/v1/models');
        return response.data.data.map((m) => m.id);
    }
}
//# sourceMappingURL=client.js.map