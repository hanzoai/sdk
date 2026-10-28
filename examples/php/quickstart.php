<?php
/**
 * PHP Quickstart for Hanzo AI SDK
 *
 * Install: composer require hanzoai/hanzoai
 * Repository: https://github.com/hanzo-php/sdk (mirror: https://github.com/hanzoai/php-sdk)
 */

require_once __DIR__ . '/vendor/autoload.php';

use Hanzo\Client;

$apiKey = getenv('HANZO_API_KEY') ?: 'sk-test-key';
$client = new Client(['apiKey' => $apiKey]);

// 1. List Models (Public Gateway Catalog)
echo "--- 1. List Models ---\n";
$models = $client->models->list();
foreach (array_slice($models['data'], 0, 5) as $model) {
    echo "Model: " . $model['id'] . " (" . ($model['owned_by'] ?? 'hanzo') . ")\n";
}

// 2. Chat Completion
echo "\n--- 2. Chat Completion ---\n";
$response = $client->chat->completions->create([
    'model' => 'zen5',
    'messages' => [
        ['role' => 'system', 'content' => 'You are a concise AI assistant.'],
        ['role' => 'user', 'content' => 'Say hello from PHP!']
    ]
]);

echo "Reply: " . $response['choices'][0]['message']['content'] . "\n";

// 3. Kai Agent Decision Evaluation
echo "\n--- 3. Kai Decision Evaluation ---\n";
$decision = $client->decisions->evaluate([
    'agent_id' => 'agent-php-01',
    'policy_id' => 'pol-safety-core',
    'action' => ['type' => 'read_file', 'payload' => ['path' => 'data.json']]
]);
echo "Verdict: " . $decision['verdict'] . " (Score: " . $decision['confidence'] . ")\n";
