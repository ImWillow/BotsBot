package tools

import (
	"BotsBot/internal/configuration"
	"strings"
)

// AdminCheck - tool for check user. If user is admin return true.
func AdminCheck(id string) bool {
	if strings.Contains(configuration.AdminList, id) {
		return true
	}

	return false
}
