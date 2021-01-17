package commands

import (
	"BotsBot/internal/configuration"
	"BotsBot/internal/tools"

	"github.com/bwmarrin/discordgo"
)

// ListCommands - command for send commands list
func ListCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	if tools.AdminCheck(m.Author.ID) {
		content := configuration.UserCommandsList + "\n" + configuration.AdminCommandsList
		tools.SendToUser(s, m, content)
	} else {
		content := configuration.UserCommandsList
		tools.SendToUser(s, m, content)
	}
}
