// Dart Quickstart for Hanzo AI SDK
//
// Install: dart pub add hanzoai
// Repository: https://github.com/hanzo-dart/sdk (mirror: https://github.com/hanzoai/dart-sdk)

import 'dart:io';
import 'package:hanzoai/hanzoai.dart';

void main() async {
  final apiKey = Platform.environment['HANZO_API_KEY'] ?? 'sk-test-key';
  final client = HanzoClient(apiKey: apiKey);

  print('--- 1. List Models ---');
  final models = await client.models.list();
  for (final m in models.data.take(5)) {
    print('Model: ${m.id} (${m.ownedBy})');
  }

  print('\n--- 2. Chat Completion ---');
  final response = await client.chat.createCompletion(
    ChatCompletionRequest(
      model: 'zen5',
      messages: [
        ChatMessage(role: 'user', content: 'Say hello from Dart!'),
      ],
    ),
  );
  print('Reply: ${response.choices.first.message.content}');
}
