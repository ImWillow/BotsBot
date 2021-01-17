package commands

import (
	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
)

// WhoIAm - command for testing
func WhoIAm(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := m.Author.Username + " - pidr"
	_, err := s.ChannelMessageSend(m.ChannelID, content)
	if err != nil {
		log.Debug("Send message error: ", err)
	}
}
