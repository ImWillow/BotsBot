package main

import (
	//Internal modules
	"BotsBot/internal/bot"
	"BotsBot/internal/logging"

	// External modules
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Init system signals
	shutdown := make(chan bool)
	//create a notification channel to shutdown
	sigChan := make(chan os.Signal, 1)
	//register for interupt (Ctrl+C) and SIGTERM (docker)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// Start goroutine for receiving signals
	go func() {
		// Wait any of signal in Signal Channel
		sig := <-sigChan
		log.Info("Signal received: ", sig.String())
		// Set shutdown = true
		shutdown <- true
	}()
	// Init logger
	logging.InitLogger()

	log.SetLevel(logging.ParseLogLevel())

	bot.Start()

	log.Debug("Bot is running!")
	<-shutdown
}
