package handlers

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// InitHandlers - function for init handlers
func InitHandlers(ds *discordgo.Session) {
	log.Debug("InitHandlers: Start!")

	log.Debug("Init MessageHandler")
	ds.AddHandler(MessageHandler)

	log.Debug("InitHandlers: Finish!")
}
