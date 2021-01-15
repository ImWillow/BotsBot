package bot

import (
	"BotsBot/internal/configuration"
	"BotsBot/internal/handlers"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Start - start bot
func Start() {
	log.Trace("Bot Start: Start!")
	DiscordSession, err := discordgo.New("Bot " + configuration.Token)
	if err != nil {
		log.Fatal(err)
		return
	}

	handlers.InitHandlers(DiscordSession)

	err = DiscordSession.Open()
	if err != nil {
		log.Fatal("Discord init error: ", err)
	}

	log.Trace("Bot Start: Finish!")
}
