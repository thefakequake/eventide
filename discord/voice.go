package discord

import (
	"time"
)

// https://discord.com/developers/docs/resources/voice#voice-state-object-voice-state-structure
type VoiceState struct {
	// The guild ID this voice state is for
	GuildID string `json:"guild_id,omitempty"`

	// The channel ID this user is connected to
	ChannelID string `json:"channel_id"`

	// The user ID this voice state is for
	UserID string `json:"user_id"`

	// The guild member this voice state is for
	Member *GuildMember `json:"member,omitempty"`

	// The session ID for this voice state
	SessionID string `json:"session_id"`

	// Whether this user is deafened by the server
	Deaf bool `json:"deaf"`

	// Whether this user is muted by the server
	Mute bool `json:"mute"`

	// Whether this user is locally deafened
	SelfDeaf bool `json:"self_deaf"`

	// Whether this user is locally muted
	SelfMute bool `json:"self_mute"`

	// Whether this user is streaming using "Go Live"
	SelfStream bool `json:"self_stream,omitempty"`

	// Whether this user's camera is enabled
	SelfVideo bool `json:"self_video"`

	// Whether this user is muted by the current user
	Suppress bool `json:"suppress"`

	// The time at which the user requested to speak
	RequestToSpeakTimestamp time.Time `json:"request_to_speak_timestamp"`
}

// https://discord.com/developers/docs/resources/voice#voice-region-object-voice-region-structure
type VoiceRegion struct {
	// Unique ID for the region
	ID string `json:"id"`

	// Name of the region
	Name string `json:"name"`

	// True for a single server that is closest to the current user's client
	Optimal bool `json:"optimal"`

	// Whether this is a custom voice region (used for events/etc)
	Custom bool `json:"custom"`
}
