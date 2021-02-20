package configuration

import (
	"github.com/bwmarrin/discordgo"
)

const (
	// GamesList - list with games on discord server
	GamesList = "LostArk,LeagueOfLegends,Valorant,Dota2,Warframe,Terraria,WorldOfWarcraft"

	// AdminList - list with all administration IDs
	AdminList = "264355075944611840,501043582656839682,221621445971804160,288022850093318146,215819877116674048"

	// UserCommandsList - list with commands for any users
	UserCommandsList = "User commands: ?addme, ?gameList."

	// AdminCommandsList - list with commands for users with admin roles
	AdminCommandsList = "Admin commands: ?clear, ?send."
)

// NewUserMessage - structure of new user messages
type NewUserMessage struct {
	Username string
	Age      string
	Games    string
	Gender   string
}

// EmbedMessage - structure for send embed message
type EmbedMessage struct {
	Fields      []*discordgo.MessageEmbedField
	Description string
	Image       *discordgo.MessageEmbedImage
	Provider    *discordgo.MessageEmbedProvider
	Thumbnail   *discordgo.MessageEmbedThumbnail
	Video       *discordgo.MessageEmbedVideo
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
	WoWers          string
	Dota2           string
	Warframe        string
	Terraria        string
	ServerPlayers   string
	FirstCategory   string
}

// TextChannels - structure of server text channels
// TODO: Доделать текстовые каналы.
type TextChannels struct {
	BotLog       string
	Bids         string
	ModeratorLog string
}

// VoiceChannels - structure of server voice channels
// TODO: Доделать голосовые каналы.
type VoiceChannels struct {
}
