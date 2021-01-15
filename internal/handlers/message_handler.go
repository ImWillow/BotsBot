package handlers

import (
	"BotsBot/internal/commands"
	"BotsBot/internal/configuration"
	"BotsBot/internal/tools"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageHandler - handler for catch messages
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages that don't have the !checkers prefix
	if !strings.HasPrefix(m.Content, configuration.Prefix) {
		return
	}

	switch command := m.Content; command {
	case configuration.Prefix + "I?":
		commands.WhoIAm(s, m)
	case configuration.Prefix + "Addme" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"Addme"):
		commands.AddMe(s, m)
	case configuration.Prefix + "GameList":
		commands.GamesList(s, m)
	}

}
