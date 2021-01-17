package commands

import (
	"BotsBot/internal/configuration"
	"BotsBot/internal/tools"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/bwmarrin/discordgo"
)

// AddMe - command for add roles to new user
func AddMe(s *discordgo.Session, m *discordgo.MessageCreate) {
	if !tools.CheckAddMeCommand(m.Content) {
		log.Debug("Unexpected message!")
		return
	}

	var message configuration.NewUserMessage

	splittedContent := strings.Split(m.Content, "/")

	message.Username = tools.DeletePrefix(splittedContent[0], configuration.Prefix+"addme ")
	err := s.GuildMemberNickname(configuration.ServerID, m.Author.ID, message.Username)
	if err != nil {
		log.Debug("Send message error: ", err)
	}

	// TODO: Придумай че с этим можно сделать.
	// message.Age = splittedContent[1]
	// message.Gender = splittedContent[3]

	message.Games = splittedContent[2]

	roles := tools.GetRoles(message.Games)
	for _, role := range roles {
		err := s.GuildMemberRoleAdd(configuration.ServerID, m.Author.ID, role)
		if err != nil {
			log.Debug("Send message error: ", err)
		}
	}

	content := "Пользователь " + m.Author.Username + " одобрен!"
	_, err = s.ChannelMessageSend(configuration.ServerTextChannels.Bids, content)
	if err != nil {
		log.Debug("Send message error: ", err)
	}

	log.Debug("User ", m.Author.Username, " added.")
}

// GamesList - command for send games list to user
func GamesList(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := "List of games: " + configuration.GamesList
	tools.SendToUser(s, m, content)
}
