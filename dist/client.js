"use strict";
/**
 * Hanzo API client
 */
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.HanzoClient = void 0;
const axios_1 = __importDefault(require("axios"));
class HanzoClient {
    constructor(apiKey, baseURL = 'https://api.hanzo.ai') {
        this.client = axios_1.default.create({
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
exports.HanzoClient = HanzoClient;
//# sourceMappingURL=client.js.map