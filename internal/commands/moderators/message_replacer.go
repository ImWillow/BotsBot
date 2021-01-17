package commands

import (
	"BotsBot/internal/tools"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// ReplaceMessage - command for replace author message as bot message
func ReplaceMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
	if err != nil {
		log.Error("Send message error: ", err)
		return
	}

	content := tools.DeletePrefix(m.Message.Content, "?send")
	_, err = s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		log.Debug("Send message error: ", err)
	}
}
