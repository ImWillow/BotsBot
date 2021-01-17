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
	case configuration.Prefix + "addme" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"addme"):
		if !tools.CheckPlayerRole(m.Member.Roles) {
			commands.AddMe(s, m)
		} else {
			tools.SendToUser(s, m, "Ты  уже добавлен как игрок сервера! Радуйся!")
		}
	case configuration.Prefix + "gameList":
		commands.GamesList(s, m)
	case configuration.Prefix + "clear" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"clear"):
		if tools.AdminCheck(m.Author.ID) {
			commands.Clear(s, m)
		}
	case configuration.Prefix + "commandsList":
		commands.ListCommands(s, m)
	}

}
