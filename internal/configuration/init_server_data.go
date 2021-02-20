package configuration

// InitRoles - function for init server roles
func InitRoles() (roles Roles) {
	roles.Bots = "638474358372827176"
	roles.Saint = "567098582277292042"
	roles.Administration = "501059154132074496"
	roles.Moderator = "638463191239622686"
	roles.Drinker = "740901516982091896"
	roles.Streamer = "674960138939138048"
	roles.LostArk = "789812703543164958"
	roles.LeagueOfLegends = "638472018773147709"
	roles.Valorant = "737228757504491521"
	roles.WoWers = "812695343321448459"
	roles.Dota2 = "638824899259006996"
	roles.Warframe = "712989643179556924"
	roles.Terraria = "708622509003636828"
	roles.ServerPlayers = "379721328288333844"
	roles.FirstCategory = "639385681722081282"

	return
}

// InitTextChannels - function for init text channels
// TODO: Доделать текстовые каналы.
func InitTextChannels() (textChannels TextChannels) {
	textChannels.BotLog = "800759923843137537"
	textChannels.ModeratorLog = "379664015472852993"
	textChannels.Bids = "638470434685452328"
	return
}

// InitVoiceChannels - function for init voice channels
// TODO: Доделать текстовые каналы.
func InitVoiceChannels() (voiceChannels VoiceChannels) {
	return
}
