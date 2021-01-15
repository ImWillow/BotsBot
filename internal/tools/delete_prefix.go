package tools

import (
	"strings"
)

// DeletePrefix - tool for delete prefix on executin commands
func DeletePrefix(messageWithPrefix string, prefix string) (messageWithoutPrefix string) {
	messageWithoutPrefix = strings.TrimPrefix(messageWithPrefix, prefix)
	return
}
