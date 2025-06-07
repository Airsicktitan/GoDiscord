package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

var commandHandlers = map[string]func(*discordgo.Session, *discordgo.MessageCreate){
	"!ping":   handlePing,
	"!hello":  handleHello,
	"!bye":    handleBye,
	"!help":   handleHelp,
	"!whoami": handleWhoami,
	"!google": handleGoogle,
	"!joke":   handleJoke,
	"!delete": handleDelete,
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	// This prevents the bot from responding to its own messages
	if m.Author.ID == s.State.User.ID {
		return
	}
	
	// Check if the message starts with a command prefix
	handler, exists := commandHandlers[m.Content]
	if exists {
		handler(s, m)
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

func handleDelete(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if the message is from the owner of the server
	if m.Author.ID != "245711794344034305" {
		s.ChannelMessageSend(m.ChannelID, "You can only delete your own messages.")
		return
	}

	// Delete the message
	messages, err := s.ChannelMessages(m.ChannelID, 100, "", "", "")
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Failed to delete the message.")
		return
	}

	var messageIDs []string
	for _, msg := range messages {
		if time.Since(msg.Timestamp).Hours() < 24*14 {
			messageIDs = append(messageIDs, msg.ID)
		}
	}

	if len(messageIDs) > 0 {
		err := s.ChannelMessagesBulkDelete(m.ChannelID, messageIDs)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Failed to delete the message.")
			return
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "No messages found to delete.")
	}
}
