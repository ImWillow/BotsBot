package configuration

const (
	// Token - token for discord bot
	Token = "Place_Your_Token_Here"

	// ServerID - discord server ID
	ServerID = "Place_YourServer_ID_Here"

	// LoggingLevel - level of logs
	LoggingLevel = "TRACE"

	// Prefix - prefix for using bot commands
	Prefix = "?"
)

// ServerRoles - inited server roles
var ServerRoles = InitRoles()

// ServerTextChannels - inited server text channels
var ServerTextChannels = InitTextChannels()

// ServerVoiceChannels - inited server voice channels
var ServerVoiceChannels = InitVoiceChannels()
