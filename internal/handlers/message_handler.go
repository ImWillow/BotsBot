package handlers

import (
	"BotsBot/internal/commands"
	"BotsBot/internal/configuration"
	"BotsBot/internal/tools"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
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
	case configuration.Prefix + "addme" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"addme"):
		commands.AddMe(s, m)
	case configuration.Prefix + "gameList":
		commands.GamesList(s, m)
	case configuration.Prefix + "clear" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"clear"):
		if strings.Contains(configuration.AdminList, m.Author.ID) {
			commands.Clear(s, m)
		} else {
			log.Debug("У чела недостаточно прав")
		}
	}

}
