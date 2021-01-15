package configuration

const (
	// GamesList - list with games on discord server
	GamesList = "LostArk,LeagueOfLegends,Valorant,Dota2,Warframe,Terraria"
)

// NewUserMessage - structure of new user messages
type NewUserMessage struct {
	Username string
	Age      string
	Games    string
	Gender   string
}

// Roles - structure of roles from server
type Roles struct {
	Bots            string
	Saint           string
	Administration  string
	Moderator       string
	Drinker         string
	Streamer        string
	LostArk         string
	LeagueOfLegends string
	Valorant        string
	Dota2           string
	Warframe        string
	Terraria        string
	ServerPlayers   string
	FirstCategory   string
}

// TextChannels - structure of server text channels
// TODO: Доделать текстовые каналы.
type TextChannels struct {
}

// VoiceChannels - structure of server voice channels
// TODO: Доделать голосовые каналы.
type VoiceChannels struct {
}
