package tools

import (
	"BotsBot/internal/configuration"
	"strings"

	log "github.com/sirupsen/logrus"
)

// AdminCheck - tool for check user. If user is admin return true.
func AdminCheck(id string) bool {
	if strings.Contains(configuration.AdminList, id) {
		return true
	}

	return false
}

// CheckPlayerRole - tool for chek player role
func CheckPlayerRole(roles []string) bool {
	log.Debug("Roles: ", roles)
	for _, role := range roles {
		if role == configuration.ServerRoles.ServerPlayers {
			return true
		}
	}

	return false
}
