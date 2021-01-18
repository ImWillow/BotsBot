package tools

import (
	"BotsBot/internal/configuration"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// CommandLog - tool for send logs of using commands
func CommandLog(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := "Пользователь " + m.Author.Username + " использовал комманду: " + m.Message.Content
	_, err := s.ChannelMessageSend(configuration.ServerTextChannels.BotLog, content)
	if err != nil {
		log.Error("Send message error: ", err)
	}

	embed := &discordgo.MessageEmbed{}
	embed.Description = "Описание тута"
	author := &discordgo.MessageEmbedAuthor{}
	author.Name = m.Author.Username
	author.IconURL = m.Author.AvatarURL("")
	footer := &discordgo.MessageEmbedFooter{}
	// footer.Text = m.Timestamp.Parse(time.Now())
	embed.Footer = footer
	embed.Author = author
	_, err = s.ChannelMessageSendEmbed(configuration.ServerTextChannels.BotLog, embed)
	if err != nil {
		log.Error("Send message error: ", err)
	}
}
