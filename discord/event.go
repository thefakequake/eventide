package discord

import "time"

type Event interface {
	EventType() string
}

// https://discord.com/developers/docs/topics/gateway#ready
type ReadyEvent struct {
	// Gateway version
	Version int `json:"v"`

	// Information about the user including email
	User *User `json:"user"`

	// The guilds the user is in
	Guilds any `json:"guilds"`

	// Used for resuming connections
	SessionID string `json:"session_id"`

	// The shard information associated with this session, if sent when identifying
	Shard []int `json:"shard"`

	// Contains ID and flags
	Application *Application `json:"application"`
}

func (r *ReadyEvent) EventType() string { return "READY" }

// https://discord.com/developers/docs/topics/gateway#resumed
type ResumedEvent struct {
}

func (r *ResumedEvent) EventType() string { return "RESUMED" }

// https://discord.com/developers/docs/topics/gateway#application-command-permissions-update
type ApplicationCommandPermissionsUpdateEvent struct {
	*ApplicationCommandPermissions
}

func (a *ApplicationCommandPermissionsUpdateEvent) EventType() string {
	return "APPLICATION_COMMAND_PERMISSIONS_UPDATE"
}

// https://discord.com/developers/docs/topics/gateway#channel-create
type ChannelCreateEvent struct {
	*Channel
}

func (c *ChannelCreateEvent) EventType() string { return "CHANNEL_CREATE" }

// https://discord.com/developers/docs/topics/gateway#channel-update
type ChannelUpdateEvent struct {
	*Channel
}

func (c *ChannelUpdateEvent) EventType() string { return "CHANNEL_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#channel-delete
type ChannelDeleteEvent struct {
	*Channel
}

func (c *ChannelDeleteEvent) EventType() string { return "CHANNEL_DELETE" }

// https://discord.com/developers/docs/topics/gateway#channel-delete
type ThreadCreateEvent struct {
	*Channel

	// True if the thread is newly created
	NewlyCreated bool `json:"newly_created,omitempty"`
}

func (t *ThreadCreateEvent) EventType() string { return "THREAD_CREATE" }

// https://discord.com/developers/docs/topics/gateway#thread-update
type ThreadUpdateEvent struct {
	*Channel
}

func (t *ThreadUpdateEvent) EventType() string { return "THREAD_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#thread-delete
type ThreadDeleteEvent struct {
	// The ID of this channel
	ID string `json:"id"`

	// The ID of the guild
	GuildID string `json:"guild_id"`

	// ID of the text channel this thread was created
	ParentID string `json:"parent_id"`

	// The type of channel
	Type ChannelType `json:"type"`
}

func (t *ThreadDeleteEvent) EventType() string { return "THREAD_DELETE" }

// https://discord.com/developers/docs/topics/gateway#thread-list-sync
type ThreadListSyncEvent struct {
	// The ID of the guild
	GuildID string `json:"guild_id"`

	// The parent channel IDs whose threads are being synced. If omitted, then threads were synced for the entire guild. This array may contain channel_ids that have no active threads as well, so you know to clear that data.
	ChannelIDs []string `json:"channel_ids,omitempty"`

	// All active threads in the given channels that the current user can access
	Threads []*Channel `json:"threads"`

	// All thread member objects from the synced threads for the current user, indicating which threads the current user has been added to
	Members []*ThreadMember `json:"members"`
}

func (t *ThreadListSyncEvent) EventType() string { return "THREAD_LIST_SYNC" }

// https://discord.com/developers/docs/topics/gateway#thread-member-update
type ThreadMemberUpdateEvent struct {
	*ThreadMember

	// The ID of the guild
	GuildID string `json:"guild_id"`
}

func (t *ThreadMemberUpdateEvent) EventType() string { return "THREAD_MEMBER_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#thread-members-update
type ThreadMembersUpdateEvent struct {
	// The ID of the thread
	ID string `json:"id"`

	// The ID of the guild
	GuildID string `json:"guild_id"`

	// The approximate number of members in the thread, capped at 50
	MemberCount int `json:"member_count"`
}

func (t *ThreadMembersUpdateEvent) EventType() string { return "THREAD_MEMBERS_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#channel-pins-update
type ChannelPinsUpdateEvent struct {
	// The ID of the guild,
	GuildID string `json:"guild_id,omitempty"`

	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The time at which the most recent pinned message was pinned
	LastPinTimestamp time.Time `json:"last_pin_timestamp,omitempty"`
}

func (c *ChannelPinsUpdateEvent) EventType() string { return "CHANNEL_PINS_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-create
type GuildCreateEvent struct {
	*Guild

	// When this guild was joined at
	JoinedAt time.Time `json:"joined_at"`

	// True if this is considered a large guild
	Large bool `json:"large"`

	// True if this guild is unavailable due to an outage
	Unavailable bool `json:"unavailable"`

	// Total number of members in this guild
	MemberCount int `json:"member_count"`

	// States of members currently in voice channels; lacks the guild_id key
	VoiceStates []*VoiceState `json:"voice_states"`

	// Users in the guild
	Members []*GuildMember `json:"member"`

	// Channels in the guild
	Channels []*Channel `json:"channels"`

	// All active threads in the guild that current user has permission to view
	Threads []*Channel `json:"threads"`

	// Presences of the members in the guild, will only include non-offline members if the size is greater than large threshold
	Presences []*GatewayPresenceUpdate `json:"presences"`

	// Stage instances in the guild
	StageInstances []*StageInstance `json:"stage_instances"`

	// The scheduled events in the guild
	GuildScheduledEvents []*GuildScheduledEvent `json:"guild_scheduled_events"`
}

func (g *GuildCreateEvent) EventType() string { return "GUILD_CREATE" }

// https://discord.com/developers/docs/topics/gateway#guild-update
type GuildUpdateEvent struct {
	*Guild
}

func (g *GuildUpdateEvent) EventType() string { return "GUILD_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-delete
type GuildDeleteEvent struct {
	*UnavailableGuild
}

func (g *GuildDeleteEvent) EventType() string { return "GUILD_DELETE" }

// https://discord.com/developers/docs/topics/gateway#guild-ban-add
type GuildBanAddEvent struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// The banned user
	User *User `json:"user"`
}

func (g *GuildBanAddEvent) EventType() string { return "GUILD_BAN_ADD" }

// https://discord.com/developers/docs/topics/gateway#guild-ban-remove
type GuildBanRemoveEvent struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// The unbanned user
	User *User `json:"user"`
}

func (g *GuildBanRemoveEvent) EventType() string { return "GUILD_BAN_REMOVE" }

// https://discord.com/developers/docs/topics/gateway#guild-emojis-update
type GuildEmojisUpdateEvent struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// Array of emojis
	Emojis []*Emoji `json:"emojis"`
}

func (g *GuildEmojisUpdateEvent) EventType() string { return "GUILD_EMOJIS_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-stickers-update
type GuildStickersUpdateEvent struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// Array of stickers
	Stickers []*Sticker
}

func (g *GuildStickersUpdateEvent) EventType() string { return "GUILD_STICKERS_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-integrations-update
type GuildIntegrationsUpdateEvent struct {
	// ID of the guild whose integrations were updated
	GuildID string `json:"guild_id"`
}

func (g *GuildIntegrationsUpdateEvent) EventType() string { return "GUILD_INTEGRATIONS_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-member-add
type GuildMemberAddEvent struct {
	*GuildMember

	// ID of the guild
	GuildID string `json:"guild_id"`
}

func (g *GuildMemberAddEvent) EventType() string { return "GUILD_MEMBER_ADD" }

// https://discord.com/developers/docs/topics/gateway#guild-member-remove
type GuildMemberRemoveEvent struct {
	// The ID of the guild
	GuildID string `json:"guild_id"`

	// The user who was removed
	User *User `json:"user"`
}

func (g *GuildMemberRemoveEvent) EventType() string { return "GUILD_MEMBER_REMOVE" }

// https://discord.com/developers/docs/topics/gateway#guild-member-update
type GuildMemberUpdateEvent struct {
	// The ID of the guild
	GuildID string `json:"guild_id"`

	// User role IDs
	Roles []string `json:"roles"`

	// The user
	User *User `json:"user"`

	// Nickname of the user in the guild
	Nick string `json:"nick,omitempty"`

	// The member's guild avatar hash
	Avatar string `json:"avatar"`

	// When the user joined the guild
	JoinedAt time.Time `json:"joined_at"`

	// When the user starting boosting the guild
	PremiumSince time.Time `json:"premium_since,omitempty"`

	// Whether the user is deafened in voice channels
	Deaf bool `json:"deaf,omitempty"`

	// Whether the user is muted in voice channels
	Mute bool `json:"mute,omitempty"`

	// Whether the user has not yet passed the guild's Membership Screening requirements
	Pending bool `json:"pending,omitempty"`

	CommunicationDisabledUntil time.Time `json:"communication_disabled_until,omitempty"`
}

func (g *GuildMemberUpdateEvent) EventType() string { return "GUILD_MEMBER_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-members-chunk
type GuildMembersChunkEvent struct {
	// The ID of the guild
	GuildID string `json:"guild_id"`

	// Set of guild members
	Members []*GuildMember `json:"members"`

	// The chunk index in the expected chunks for this response (0 <= chunk_index < chunk_count)
	ChunkIndex int64 `json:"chunk_index"`

	// The total number of expected chunks for this response
	ChunkCount int64 `json:"chunk_count"`

	// If passing an invalid id to REQUEST_GUILD_MEMBERS, it will be returned here
	NotFound []string `json:"not_found,omitempty"`

	// If passing true to REQUEST_GUILD_MEMBERS, presences of the returned members will be here
	Presences []*PresenceUpdateEvent `json:"presences,omitempty"`

	// The nonce used in the Guild Members Request
	Nonce string `json:"nonce,omitempty"`
}

func (g *GuildMembersChunkEvent) EventType() string { return "GUILD_MEMBERS_CHUNK" }

// https://discord.com/developers/docs/topics/gateway#guild-role-create
type GuildRoleCreateEvent struct {
	// The ID of the guild
	GuildID string `json:"guild_id"`

	// The role created
	Role *Role `json:"role"`
}

func (g *GuildRoleCreateEvent) EventType() string { return "GUILD_ROLE_CREATE" }

// https://discord.com/developers/docs/topics/gateway#guild-role-update
type GuildRoleUpdateEvent struct {
	// The ID of the guild
	GuildID string `json:"guild_id"`

	// The role updated
	Role *Role `json:"role"`
}

func (g *GuildRoleUpdateEvent) EventType() string { return "GUILD_ROLE_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-role-delete
type GuildRoleDeleteEvent struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// ID of the role
	RoleID string `json:"role_id"`
}

func (g *GuildRoleDeleteEvent) EventType() string { return "GUILD_ROLE_DELETE" }

// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-create
type GuildScheduledEventCreateEvent struct {
	*GuildScheduledEvent
}

func (g *GuildScheduledEventCreateEvent) EventType() string { return "GUILD_SCHEDULED_EVENT_CREATE" }

// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-update
type GuildScheduledEventUpdateEvent struct {
	*GuildScheduledEvent
}

func (g *GuildScheduledEventUpdateEvent) EventType() string { return "GUILD_SCHEDULED_EVENT_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-delete
type GuildScheduledEventDeleteEvent struct {
	*GuildScheduledEvent
}

func (g *GuildScheduledEventDeleteEvent) EventType() string { return "GUILD_SCHEDULED_EVENT_DELETE" }

// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-add
type GuildScheduledEventUserAddEvent struct {
	// ID of the guild scheduled event
	GuildScheduledEventID string `json:"guild_scheduled_event_id"`

	// ID of the user
	UserID string `json:"user_id"`

	// ID of the guild
	GuildID string `json:"guild_id"`
}

func (g *GuildScheduledEventUserAddEvent) EventType() string { return "GUILD_SCHEDULED_EVENT_USER_ADD" }

// https://discord.com/developers/docs/topics/gateway#guild-scheduled-event-user-add
type GuildScheduledEventUserRemoveEvent struct {
	// ID of the guild scheduled event
	GuildScheduledEventID string `json:"guild_scheduled_event_id"`

	// ID of the user
	UserID string `json:"user_id"`

	// ID of the guild
	GuildID string `json:"guild_id"`
}

func (g *GuildScheduledEventUserRemoveEvent) EventType() string {
	return "GUILD_SCHEDULED_EVENT_USER_REMOVE"
}

// https://discord.com/developers/docs/topics/gateway#integration-create
type IntegrationCreateEvent struct {
	*Integration

	// ID of the guild
	GuildID string `json:"guild_id"`
}

func (i *IntegrationCreateEvent) EventType() string { return "INTEGRATION_CREATE" }

// https://discord.com/developers/docs/topics/gateway#integration-update
type IntegrationUpdateEvent struct {
	*Integration

	// ID of the guild
	GuildID string `json:"guild_id"`
}

func (i *IntegrationUpdateEvent) EventType() string { return "INTEGRATION_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#integration-delete
type IntegrationDeleteEvent struct {
	// Integration ID
	ID string `json:"id"`

	// ID of the guild
	GuildID string `json:"guild_id"`

	// ID of the bot/OAuth2 application for this discord integration
	ApplicationID string `json:"application_id,omitempty"`
}

func (i *IntegrationDeleteEvent) EventType() string { return "INTEGRATION_DELETE" }

// https://discord.com/developers/docs/topics/gateway#invite-create
type InviteCreateEvent struct {
	// The channel the invite is for
	ChannelID string `json:"channel_id"`

	// The unique invite code
	Code string `json:"code"`

	// The time at which the invite was created
	CreatedAt time.Time `json:"created_at"`

	// The guild of the invite
	GuildID string `json:"guild_id,omitempty"`

	// The user that created the invite
	Inviter *User `json:"inviter,omitempty"`

	// How long the invite is valid for (in seconds)
	MaxAge int `json:"max_age"`

	// The maximum amount of times the invite can be used
	MaxUses int `json:"max_uses"`

	// The type of target for this voice channel invite
	TargetType int `json:"target_type,omitempty"`

	// The user whose stream to display for this voice channel stream invite
	TargetUser *User `json:"target_user"`

	// The embedded application to open for this voice channel embedded application invite
	TargetApplication *Application `json:"target_application,omitempty"`

	// Whether or not the invite is temporary (invited users will be kicked on disconnect unless they're assigned a role)
	Temporary bool `json:"temporary"`

	// How many times the invite has been used (always will be 0)
	Uses int `json:"uses"`
}

func (i *InviteCreateEvent) EventType() string { return "INVITE_CREATE" }

// https://discord.com/developers/docs/topics/gateway#invite-delete
type InviteDeleteEvent struct {
	// The channel of the invite
	ChannelID string `json:"channel_id"`

	// The guild of the invite
	GuildID string `json:"guild_id,omitempty"`

	// The unique invite code
	Code string `json:"code"`
}

func (i *InviteDeleteEvent) EventType() string { return "INVITE_DELETE" }

// https://discord.com/developers/docs/topics/gateway#message-create
type MessageCreateEvent struct {
	Message
}

func (m *MessageCreateEvent) EventType() string { return "MESSAGE_CREATE" }

// https://discord.com/developers/docs/topics/gateway#message-update
type MessageUpdateEvent struct {
	*Message
}

func (m *MessageUpdateEvent) EventType() string { return "MESSAGE_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#message-delete
type MessageDeleteEvent struct {
	// The ID of the message
	ID string `json:"id"`

	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The ID of the guild
	GuildID string `json:"guild_id,omitempty"`
}

func (m *MessageDeleteEvent) EventType() string { return "MESSAGE_DELETE" }

// https://discord.com/developers/docs/topics/gateway#message-delete-bulk
type MessageDeleteBulkEvent struct {
	// The IDs of the messages
	IDs []string `json:"ids"`

	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The ID of the guild
	GuildID string `json:"guild_id,omitempty"`
}

func (m *MessageDeleteBulkEvent) EventType() string { return "MESSAGE_DELETE_BULK" }

// https://discord.com/developers/docs/topics/gateway#message-reaction-add
type MessageReactionAddEvent struct {
	// The ID of the user
	UserID string `json:"user_id"`

	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The ID of the message
	MessageID string `json:"message_id"`

	// The ID of the guild
	GuildID string `json:"guild_id,omitempty"`

	// The member who reacted if this happened in a guild
	Member *GuildMember `json:"member,omitempty"`

	// The emoji used to react
	Emoji *Emoji `json:"emoji"`
}

func (m *MessageReactionAddEvent) EventType() string { return "MESSAGE_REACTION_ADD" }

// https://discord.com/developers/docs/topics/gateway#message-reaction-remove
type MessageReactionRemoveEvent struct {
	// The ID of the user
	UserID string `json:"user_id"`

	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The ID of the message
	MessageID string `json:"message_id"`

	// The ID of the guild
	GuildID string `json:"guild_id,omitempty"`

	// The emoji used to react
	Emoji *Emoji `json:"emoji"`
}

func (m *MessageReactionRemoveEvent) EventType() string { return "MESSAGE_REACTION_REMOVE" }

// https://discord.com/developers/docs/topics/gateway#message-reaction-remove-all
type MessageReactionRemoveAllEvent struct {
	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The ID of the message
	MessageID string `json:"message_id"`

	// The ID of the guild
	GuildID string `json:"guild_id,omitempty"`
}

func (m *MessageReactionRemoveAllEvent) EventType() string { return "MESSAGE_REACTION_REMOVE_ALL" }

// https://discord.com/developers/docs/topics/gateway#message-reaction-remove-emoji
type MessageReactionRemoveEmojiEvent struct {
	// The ID of the channel
	ChannelID string `json:"channel_id"`

	// The ID of the guild
	GuildID string `json:"guild_id,omitempty"`

	// The ID of the message
	MessageID string `json:"message_id"`

	// The emoji that was removed
	Emoji *Emoji `json:"emoji"`
}

func (m *MessageReactionRemoveEmojiEvent) EventType() string { return "MESSAGE_REACTION_REMOVE_EMOJI" }

// https://discord.com/developers/docs/topics/gateway#presence-update
type PresenceUpdateEvent struct {
	// The user presence is being updated for
	User *User `json:"user"`

	// ID of the guild
	GuildID string `json:"guild_id"`

	// Either "idle", "dnd", "online", or "offline"
	Status string `json:"status"`

	// User's current activities
	Activities []*Activity `json:"activities"`

	// User's platform-dependent status
	ClientStatus *ClientStatus `json:"client_status"`
}

func (p *PresenceUpdateEvent) EventType() string { return "PRESENCE_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#typing-start
type TypingStartEvent struct {
	// ID of the channel
	ChannelID string `json:"channel_id"`

	// ID of the guild
	GuildID string `json:"guild_id,omitempty"`

	// ID of the user
	UserID string `json:"user_id"`

	// Unix time (in seconds) of when the user started typing
	Timestamp int `json:"timestamp"`

	// The member who started typing if this happened in a guild
	Member *GuildMember `json:"member,omitempty"`
}

func (t *TypingStartEvent) EventType() string { return "TYPING_START" }

// https://discord.com/developers/docs/topics/gateway#user-update
type UserUpdateEvent struct {
	*User
}

func (u *UserUpdateEvent) EventType() string { return "USER_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#voice-state-update
type VoiceStateUpdateEvent struct {
	*VoiceState
}

func (v *VoiceStateUpdateEvent) EventType() string { return "VOICE_STATE_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#voice-server-update
type VoiceServerUpdateEvent struct {
	// Voice connection token
	Token string `json:"token"`

	// The guild this voice server update is for
	GuildID string `json:"guild_id"`

	// The voice server host
	Endpoint string `json:"endpoint"`
}

func (v *VoiceServerUpdateEvent) EventType() string { return "VOICE_SERVER_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#webhooks-update
type WebhooksUpdateEvent struct {
	// ID of the guild
	GuildID string `json:"guild_id"`

	// ID of the channel
	ChannelID string `json:"channel_id"`
}

func (v *WebhooksUpdateEvent) EventType() string { return "WEBHOOKS_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#interaction-create
type InteractionCreateEvent struct {
	*Interaction
}

func (i *InteractionCreateEvent) EventType() string { return "INTERACTION_CREATE" }

// https://discord.com/developers/docs/topics/gateway#stage-instance-create
type StageInstanceCreateEvent struct {
	*StageInstance
}

func (s *StageInstanceCreateEvent) EventType() string { return "STAGE_INSTANCE_CREATE" }

// https://discord.com/developers/docs/topics/gateway#stage-instance-update
type StageInstanceUpdateEvent struct {
	*StageInstance
}

func (s *StageInstanceUpdateEvent) EventType() string { return "STAGE_INSTANCE_UPDATE" }

// https://discord.com/developers/docs/topics/gateway#stage-instance-delete
type StageInstanceDeleteEvent struct {
	*StageInstance
}

func (s *StageInstanceDeleteEvent) EventType() string { return "STAGE_INSTANCE_DELETE" }
