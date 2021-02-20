package commands

import (
	"BotsBot/internal/tools"

	"github.com/bwmarrin/discordgo"
)

// HelpList - command for help with commands
func HelpList(s *discordgo.Session, m *discordgo.MessageCreate, command string) {
	switch command {
	case "?addme":
		content := "Комманда '?addme' выдает роли указанные после второго '/'. \nНапример: '?addme user1/-/LeagueOfLegend/-' добавит вам роль LoLеры"
		tools.SendToUser(s, m, content)
	}
}
