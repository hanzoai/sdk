#!/usr/bin/env ruby
# Ruby Quickstart for Hanzo AI SDK
#
# Install: gem install hanzoai
# Repository: https://github.com/hanzo-ruby/sdk (mirror: https://github.com/hanzoai/ruby-sdk)

require 'hanzoai'

client = Hanzoai::Client.new(
  api_key: ENV['HANZO_API_KEY'] || 'sk-test-key',
  base_url: ENV['HANZO_BASE_URL'] || 'https://api.hanzo.ai/v1'
)

puts "--- 1. List Models ---"
models = client.models.list
models.data.first(5).each do |m|
  puts "Model: #{m.id} (#{m.owned_by})"
end

puts "\n--- 2. Chat Completion ---"
reply = client.chat.completions.create(
  model: 'zen5',
  messages: [
    { role: 'user', content: 'Say hello from Ruby!' }
  ]
)
puts "Reply: #{reply.choices.first.message.content}"
