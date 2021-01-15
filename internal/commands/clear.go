package commands

import (
	"BotsBot/internal/configuration"
	"BotsBot/internal/tools"
	"strconv"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Clear - command for clear text channel
func Clear(s *discordgo.Session, m *discordgo.MessageCreate) {
	messageCount, err := strconv.Atoi(tools.DeletePrefix(m.Content, configuration.Prefix+"clear "))
	if err != nil {
		log.Error("Send message error: ", err)
		return
	}

	messages, err := s.ChannelMessages(m.ChannelID, messageCount, "", "", "")
	if err != nil {
		log.Error("Send message error: ", err)
		return
	}

	for _, message := range messages {
		err = s.ChannelMessageDelete(m.ChannelID, message.ID)
		if err != nil {
			log.Error("Send message error: ", err)
			return
		}
	}

	log.Info(messageCount, " messages delete successful")
}
