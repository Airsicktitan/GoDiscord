package main

import (
	"github.com/bwmarrin/discordgo"
)

// Define the commands that the bot will respond to
var (
	commands = []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "Replies with Pong!",
		},
		{
			Name:        "hello",
			Description: "Replies with a greeting!",
		},
		{
			Name:        "bye",
			Description: "Replies with a goodbye message!",
		},
		{
			Name:        "help",
			Description: "Lists all available commands!",
		},
		{
			Name:        "whoami",
			Description: "I am a friendly bot created to assist you!",
		},
		{
			Name:        "google",
			Description: "Link to Google",
		},
		{
			Name:        "joke",
			Description: "Tells a random joke!",
		},
		{
			Name: "reminder",
			Description: "Sets a reminder for a specified time",
		},
	}
)
