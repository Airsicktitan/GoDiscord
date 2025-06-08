package main

import (
	"testing"
)

// handleCommand simulates command handling for testing purposes.
func handleCommand(input, user string) string {
	switch input {
	case "!ping":
		return "Pong!"
	case "!hello":
		return "Hello, " + user + "!"
	case "!bye":
		return "Goodbye, " + user + "!"
	case "!help":
		return "Available commands:\n!ping - Responds with Pong!\n!hello - Greets the user\n!bye - Says goodbye\n!whoami - Tells who the bot is\n!google - Provides a Google search link\n!joke - Tells a joke\n!delete - Deletes a message"
	case "!whoami":
		return "I am a friendly bot created to assist you!"
	case "!google":
		return "https://www.google.com/search?q="
	case "!joke":
		return "Why did the chicken cross the road? To get to the other side!"
	case "!delete":
		return "Message deleted."
	default:
		return "Unknown command"
	}
}

func TestCommandHandlers(t *testing.T) {
	tests := []struct {
		input    string
		user     string
		expected string
	}{
		{"!ping", "Adam", "Pong!"},
		{"!hello", "Adam", "Hello, Adam!"},
		{"!bye", "Adam", "Goodbye, Adam!"},
		{"!help", "Adam", "Available commands:\n!ping - Responds with Pong!\n!hello - Greets the user\n!bye - Says goodbye\n!whoami - Tells who the bot is\n!google - Provides a Google search link\n!joke - Tells a joke\n!delete - Deletes a message"},
		{"!whoami", "Adam", "I am a friendly bot created to assist you!"},
		{"!google", "Adam", "https://www.google.com/search?q="},
		{"!joke", "Adam", "Why did the chicken cross the road? To get to the other side!"},
		{"!delete", "Adam", "Message deleted."},
	}

	for _, test := range tests {
		result := handleCommand(test.input, test.user)
		if result != test.expected {
			t.Errorf("For input '%s' expected '%s' but got '%s'", test.input, test.expected, result)
		}
	}
}