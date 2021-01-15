package tools

import (
	"BotsBot/internal/configuration"
	"strings"

	log "github.com/sirupsen/logrus"
)

// CheckAddMeCommand - tool for validate new user messages
func CheckAddMeCommand(message string) bool {
	log.Debug("Message: ", message)
	if len(message) == 0 {
		return false
	}

	splittedMessage := strings.Split(message, "/")
	log.Debug("Splitted message: ", splittedMessage)
	if len(splittedMessage) != 4 {
		return false
	}

	for _, i := range splittedMessage {
		if len(i) == 0 {
			return false
		}
	}

	return true
}

// GetRoles - tools for get roles of user games list
func GetRoles(gamelist string) (roles []string) {
	splittedGamesList := strings.Split(gamelist, ", ")
	for _, game := range splittedGamesList {
		switch game {
		case "LostArk":
			roles = append(roles, configuration.ServerRoles.LostArk)

		case "LeagueOfLegends":
			roles = append(roles, configuration.ServerRoles.LeagueOfLegends)

		case "Valorant":
			roles = append(roles, configuration.ServerRoles.Valorant)

		case "Dota2":
			roles = append(roles, configuration.ServerRoles.Dota2)

		case "Warframe":
			roles = append(roles, configuration.ServerRoles.Warframe)

		case "Terraria":
			roles = append(roles, configuration.ServerRoles.Terraria)
		}
	}

	roles = append(roles, configuration.ServerRoles.ServerPlayers)

	return
}
