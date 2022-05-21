package eventide

import (
	"time"
)

type GatewayURL struct {
	URL string `json:"url"`
}

type OverwriteType int

const (
	OverwriteRole OverwriteType = iota
	OverwriteMember
)

type ChannelFlags int

const (
	ChannelFlagsPinned ChannelFlags = iota
)

type Activity struct {
	Name          string             `json:"activity"`
	Type          ActivityType       `json:"type"`
	URL           string             `json:"url"`
	CreatedAt     int                `json:"created_at"`
	Timestamps    *ActivityTimestamp `json:"timestamps"`
	ApplicationID string             `json:"application_id"`
	Details       string             `json:"details"`
	State         string             `json:"state"`
	Emoji         *ActivityEmoji     `json:"emoji"`
	Party         *ActivityParty     `json:"party"`
	Assets        *ActivityAssets    `json:"assets"`
	Secrets       *ActivitySecrets   `json:"secrets"`
	Instance      bool               `json:"instance"`
	Buttons       []*ActivityButton  `json:"buttons"`
}

type WelcomeScreen struct {
	Description     *string                 `json:"description"`
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelID   string `json:"channel_id"`
	Description string `json:"description"`
	EmojiID     string `json:"emoji_id"`
	EmojiName   string `json:"emoji_name"`
}

type StageInstance struct {
	ID                    string       `json:"id"`
	GuildID               string       `json:"guild_id"`
	ChannelID             string       `json:"channel_id"`
	Topic                 string       `json:"topic"`
	PrivacyLevel          PrivacyLevel `json:"privacy_level"`
	DiscoverableDisabled  bool         `json:"discoverable_disabled"`
	GuildScheduledEventID string       `json:"guild_scheduled_event_id"`
}

type MessageSend struct {
	Content string   `json:"content,omitempty"`
	Embeds  []*Embed `json:"embeds"`
	TTS     bool     `json:"tts"`
}
