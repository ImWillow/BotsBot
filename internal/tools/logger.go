package tools

import (
	"BotsBot/internal/configuration"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// CommandLog - tool for send logs of using commands
func CommandLog(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Create content
	content := "Использовал комманду: " + m.Message.Content

	messageContent := configuration.EmbedMessage{}
	messageContent.Description = content

	message := CreateMessage(m, messageContent)
	_, err := s.ChannelMessageSendEmbed(configuration.ServerTextChannels.BotLog, &message)
	if err != nil {
		log.Error("Send message error: ", err)
	}
}
