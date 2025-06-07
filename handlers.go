package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	// This prevents the bot from responding to its own messages
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch m.Content {
	case "!ping":
		handlePing(s, m)
	case "!hello":
		handleHello(s, m)
	case "!bye":
		handleBye(s, m)
	case "!help":
		handleHelp(s, m)
	case "!whoami":
		handleWhoami(s, m)
	case "!google":
		handleGoogle(s, m)
	case "!joke":
		handleJoke(s, m)
	default:
		handleUnknown(s, m)
	}
}

func handlePing(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Pong!")
}

func handleHello(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Hello, "+m.Author.Username+"!")
}

func handleBye(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Goodbye, "+m.Author.Username+"!")
}

func handleHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	helpMessage := "Available commands:\n"
	for _, cmd := range commands {
		helpMessage += "!" + cmd.Name + " - " + cmd.Description + "\n"
	}
	s.ChannelMessageSend(m.ChannelID, helpMessage)
}

func handleWhoami(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "I am a friendly bot created to assist you!")
}

func handleGoogle(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Here is a link to Google: https://www.google.com")
}

func handleJoke(s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Sorry, I couldn't fetch a joke right now.")
		return
	}
	defer resp.Body.Close()

	var joke struct {
		Setup     string `json:"setup"`
		Punchline string `json:"punchline"`
	}

	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Sorry, I couldn't find a joke. ☹️")
		return
	}

	s.ChannelMessageSend(m.ChannelID, joke.Setup)
	time.Sleep(2 * time.Second) // Wait for 2 seconds before sending the punchline
	s.ChannelMessageSend(m.ChannelID, joke.Punchline)
}

func handleUnknown(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Unknown command. Type !help for a list of commands.")
}
