package discord

// https://discord.com/developers/docs/topics/gateway#payloads-gateway-payload-structure
type GatewayPayload struct {
	// Opcode for the payload
	Op int `json:"op"`

	// Event data
	D interface{} `json:"d"`

	// Sequence number, used for resuming sessions and heartbeats
	S int `json:"s"`

	// The event name for this payload
	T string `json:"t"`
}

// https://discord.com/developers/docs/topics/gateway#identify-identify-structure
type Identify struct {
	// Authentication token
	Token string `json:"token"`

	// Connection properties
	Properties *IdentifyConnectionProperties `json:"properties"`

	// Whether this connection supports compression of packets
	Compress bool `json:"compress,omitempty"`

	// Value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	LargeThreshold int `json:"large_threshold,omitempty"`

	// Used for Guild Sharding
	Shard []int `json:"shard,omitempty"`

	// Presence structure for initial presence information
	Presence *GatewayPresenceUpdate `json:"presence,omitempty"`

	// The Gateway Intents you wish to receive
	Intents int `json:"intents"`
}

// https://discord.com/developers/docs/topics/gateway#identify-identify-connection-properties
type IdentifyConnectionProperties struct {
	// Your operating system
	OS string `json:"$os"`

	// Your library name
	Browser string `json:"$browser"`

	// Your library name
	Device string `json:"$device"`
}

// https://discord.com/developers/docs/topics/gateway#resume-resume-structure
type Resume struct {
	// Session token
	Token string `json:"token"`

	// Session ID
	SessionID string `json:"session_id"`

	// Last sequence number received
	Seq int `json:"seq"`
}

// https://discord.com/developers/docs/topics/gateway#request-guild-members-guild-request-members-structure
type GuildRequestMembers struct {
	// ID of the guild to get members for
	GuildID string `json:"guild_id"`

	// String that username starts with, or an empty string to return all members
	Query string `json:"query,omitempty"`

	// Maximum number of members to send matching the query`; a limit of 0 can be used with an empty string query to return all members
	Limit int `json:"limit"`

	// Used to specify if we want the presences of the matched members
	Presences bool `json:"presences,omitempty"`

	// Used to specify which users you wish to fetch
	UserIDs []string `json:"user_ids,omitempty"`

	// Nonce to identify the Guild Members Chunk response
	Nonce string `json:"nonce,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#update-voice-state-gateway-voice-state-update-structure
type GatewayVoiceStateUpdate struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// ID of the voice channel client wants to join (null if disconnecting)
	ChannelID string `json:"channel_id"`

	// Is the client muted
	SelfMute bool `json:"self_mute"`

	// Is the client deafened
	SelfDeaf bool `json:"self_deaf"`
}

// https://discord.com/developers/docs/topics/gateway#update-presence-gateway-presence-update-structure
type GatewayPresenceUpdate struct {
	// Unix time (in milliseconds) of when the client went idle, or null if the client is not idle
	Since int `json:"since"`

	// The user's activities
	Activities []*Activity `json:"activities"`

	// The user's new status
	Status string `json:"status"`

	// Whether or not the client is AFK
	AFK bool `json:"afk"`
}

// https://discord.com/developers/docs/topics/gateway#hello-hello-structure
type Hello struct {
	// The interval (in milliseconds) the client should heartbeat with
	HeartbeatInterval int `json:"heartbeat_interval"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-structure
type Activity struct {
	// The activity's name
	Name string `json:"name"`

	// Activity type
	Type ActivityType `json:"type"`

	// Stream URL, is validated when type is 1
	URL string `json:"url,omitempty"`

	// Unix timestamp (in milliseconds) of when the activity was added to the user's session
	CreatedAt int `json:"created_at"`

	// Unix timestamps for start and/or end of the game
	Timestamps *ActivityTimestamps `json:"timestamps,omitempty"`

	// Application ID for the game
	ApplicationID string `json:"application_id,omitempty"`

	// What the player is currently doing
	Details string `json:"details,omitempty"`

	// The user's current party status
	State string `json:"state,omitempty"`

	// The emoji used for a custom status
	Emoji *Emoji `json:"emoji,omitempty"`

	// Information for the current party of the player
	Party *ActivityParty `json:"party,omitempty"`

	// Images for the presence and their hover texts
	Assets *ActivityAssets `json:"assets,omitempty"`

	// Secrets for Rich Presence joining and spectating
	Secrets *ActivitySecrets `json:"secrets,omitempty"`

	// Whether or not the activity is an instanced game session
	Instance bool `json:"instance,omitempty"`

	// Activity flags OR'd together, describes what the payload includes
	Flags ActivityFlags `json:"flags,omitempty"`

	// The custom buttons shown in the Rich Presence (max 2)
	Buttons []*ActivityButton `json:"buttons,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-types
type ActivityType int

const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom
	ActivityTypeCompeting
)

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-timestamps
type ActivityTimestamps struct {
	// Unix time (in milliseconds) of when the activity started
	Start int `json:"start,omitempty"`

	// Unix time (in milliseconds) of when the activity ends
	End int `json:"end,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-emoji
type ActivityEmoji struct {
	// The name of the emoji
	Name string `json:"name"`

	// The ID of the emoji
	ID string `json:"id,omitempty"`

	// Whether the emoji is animated
	Animated bool `json:"animated,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-party
type ActivityParty struct {
	// The ID of the party
	ID string `json:"id,omitempty"`

	// Array of two integers (current_size, max_size)
	Size []int `json:"size,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-assets
type ActivityAssets struct {
	// See Activity Asset Image
	LargeImage string `json:"large_image,omitempty"`

	// Text displayed when hovering over the large image of the activity
	LargeText string `json:"large_text,omitempty"`

	// see Activity Asset Image
	SmallImage string `json:"small_image,omitempty"`

	// Text displayed when hovering over the small image of the activity
	SmallText string `json:"small_text,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-secrets
type ActivitySecrets struct {
	// The secret for joining a party
	Join string `json:"join,omitempty"`

	// The secret for spectating a game
	Spectate string `json:"spectate,omitempty"`

	// The secret for a specific instanced match
	Match string `json:"match,omitempty"`
}

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-flags
type ActivityFlags int

const (
	ActivityFlagsInstance ActivityFlags = 1 << iota
	ActivityFlagsJoin
	ActivityFlagsSpectate
	ActivityFlagsJoinRequest
	ActivityFlagsSync
	ActivityFlagsPlay
	ActivityFlagsPartyPrivacyFriends
	ActivityFlagsPartyPrivacyVoiceChannel
	ActivityFlagsEmbedded
)

// https://discord.com/developers/docs/topics/gateway#activity-object-activity-buttons
type ActivityButton struct {
	// The text shown on the button (1-32 characters)
	Label string `json:"label"`

	// The URL opened when clicking the button (1-512 characters)
	URL string `json:"url"`
}

// https://discord.com/developers/docs/topics/gateway#session-start-limit-object-session-start-limit-structure
type SessionStartLimit struct {
	// The total number of session starts the current user is allowed
	Total int `json:"total"`

	// The remaining number of session starts the current user is allowed
	Remaining int `json:"remaining"`

	// The number of milliseconds after which the limit resets
	ResetAfter int `json:"reset_after"`

	// The number of identify requests allowed per 5 seconds
	MaxConcurrency int `json:"max_concurrency"`
}
