/**
 * TypeScript Quickstart for Hanzo AI SDK.
 *
 * Install: npm install @hanzo/ai (or npm install hanzoai)
 */

import { Hanzo } from '@hanzo/ai';

const client = new Hanzo({
  apiKey: process.env.HANZO_API_KEY || 'sk-test-key',
  baseURL: process.env.HANZO_BASE_URL || 'https://api.hanzo.ai/v1',
});

async function main() {
  console.log('--- 1. List Available Models (Catalog) ---');
  const models = await client.models.list();
  for (const m of models.data.slice(0, 5)) {
    console.log(`Model: ${m.id} (Owner: ${m.owned_by})`);
  }

  console.log('\n--- 2. Chat Completion ---');
  const completion = await client.chat.completions.create({
    model: 'zen5',
    messages: [
      { role: 'system', content: 'You are a helpful assistant.' },
      { role: 'user', content: 'Say hello in 3 words.' },
    ],
  });
  console.log('Reply:', completion.choices[0]?.message?.content);

  console.log('\n--- 3. Kai Agent Decision Evaluation ---');
  const decision = await client.decisions.evaluate({
    agent_id: 'agent-ts-01',
    policy_id: 'pol-safety-core',
    action: { type: 'api_dispatch', payload: { endpoint: '/v1/status' } },
  });
  console.log(`Decision: ${decision.verdict} (Confidence: ${decision.confidence})`);
}

main().catch(console.error);
