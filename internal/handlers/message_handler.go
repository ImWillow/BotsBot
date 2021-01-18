package handlers

import (
	moders "BotsBot/internal/commands/moderators"
	users "BotsBot/internal/commands/users"
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
		users.WhoIAm(s, m)
	case configuration.Prefix + "addme" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"addme"):
		if !tools.CheckPlayerRole(m.Member.Roles) {
			users.AddMe(s, m)
			tools.CommandLog(s, m)
		} else {
			tools.SendToUser(s, m, "Ты  уже добавлен как игрок сервера! Радуйся!")
			tools.CommandLog(s, m)
		}
	case configuration.Prefix + "gameList":
		users.GamesList(s, m)
		tools.CommandLog(s, m)
	case configuration.Prefix + "clear" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"clear"):
		if tools.AdminCheck(m.Author.ID) {
			moders.Clear(s, m)
			tools.CommandLog(s, m)
		}
	case configuration.Prefix + "commandsList":
		users.ListCommands(s, m)
		tools.CommandLog(s, m)
	case configuration.Prefix + "send" + tools.DeletePrefix(m.Message.Content, configuration.Prefix+"send"):
		if tools.AdminCheck(m.Author.ID) {
			moders.ReplaceMessage(s, m)
			tools.CommandLog(s, m)
		}
	case configuration.Prefix + "ping":
		tools.CommandLog(s, m)
	}

}
