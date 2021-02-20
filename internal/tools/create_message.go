package tools

import (
	"BotsBot/internal/configuration"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// CreateMessage - tool for create structed message
func CreateMessage(m *discordgo.MessageCreate, embedMessage configuration.EmbedMessage) (message discordgo.MessageEmbed) {
	messageTime := ParseMessageTime(m)

	author := &discordgo.MessageEmbedAuthor{}
	footer := &discordgo.MessageEmbedFooter{}

	// Author
	author.Name = m.Author.String()
	author.IconURL = m.Author.AvatarURL("")

	if embedMessage.Fields != nil {
		message.Fields = embedMessage.Fields
	}

	// Footer
	footer.Text = messageTime

	// Image
	if embedMessage.Image != nil {
		message.Image = embedMessage.Image
	}

	// Provider
	if embedMessage.Provider != nil {
		message.Provider = embedMessage.Provider
	}

	// Thumbnail
	if embedMessage.Thumbnail != nil {
		message.Thumbnail = embedMessage.Thumbnail
	}

	// Video
	if embedMessage.Video != nil {
		message.Video = embedMessage.Video
	}

	message.Author = author
	message.Description = embedMessage.Description
	message.Footer = footer

	return
}

// ParseMessageTime - tool for parse message time
func ParseMessageTime(m *discordgo.MessageCreate) (messageTimeString string) {
	messageTime, err := m.Timestamp.Parse()
	if err != nil {
		log.Error("Time parse error: ", err)
	}
	messageTimeString = messageTime.Format("2006-01-02 15:04")

	return
}
