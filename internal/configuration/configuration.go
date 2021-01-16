package configuration

const (
	// Token - token for discord bot
	Token = "NzEyMDUyMjk4NDczMjc1NDMz.XsL8YQ.RtyPXWTWcrwXo2Rmktw6MlZfL7k"

	// ServerID - discord server ID
	ServerID = "355652238611578890"

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
