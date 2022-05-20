package eventide

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
	Properties *Object `json:"properties"`

	// Whether this connection supports compression of packets
	Compress bool `json:"compress,omitempty"`

	// Value between 50 and 250, total number of members where the gateway will stop sending offline members in the guild member list
	LargeThreshold int `json:"large_threshold,omitempty"`

	// Used for Guild Sharding
	Shard []int `json:"shard,omitempty"`

	// Presence structure for initial presence information
	Presence *UpdatePresence `json:"presence,omitempty"`

	// The Gateway Intents you wish to receive
	Intents int `json:"intents"`
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
	UserIDs *SnowflakeOrArrayOfSnowflakes `json:"user_ids,omitempty"`

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
	Type int `json:"type"`

	// Stream url, is validated when type is 1
	URL string `json:"url,omitempty"`

	// Unix timestamp (in milliseconds) of when the activity was added to the user's session
	CreatedAt int `json:"created_at"`

	// Unix timestamps for start and/or end of the game
	Timestamps *Timestamps `json:"timestamps,omitempty"`

	// Application ID for the game
	ApplicationID string `json:"application_id,omitempty"`

	// What the player is currently doing
	Details string `json:"details,omitempty"`

	// The user's current party status
	State string `json:"state,omitempty"`

	// The emoji used for a custom status
	Emoji *Emoji `json:"emoji,omitempty"`

	// Information for the current party of the player
	Party *Party `json:"party,omitempty"`

	// Images for the presence and their hover texts
	Assets *Assets `json:"assets,omitempty"`

	// Secrets for Rich Presence joining and spectating
	Secrets *Secrets `json:"secrets,omitempty"`

	// Whether or not the activity is an instanced game session
	Instance bool `json:"instance,omitempty"`

	// Activity flags OR`d together, describes what the payload includes
	Flags int `json:"flags,omitempty"`

	// The custom buttons shown in the Rich Presence (max 2)
	Buttons []*Button `json:"buttons,omitempty"`
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
