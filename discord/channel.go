package discord

import (
	"time"
)

// https://discord.com/developers/docs/resources/channel#channel-object-channel-structure
type Channel struct {
	// The ID of this channel
	ID string `json:"id"`

	// The type of channel
	Type ChannelType `json:"type"`

	// The ID of the guild (may be missing for some channel objects received over gateway guild dispatches)
	GuildID string `json:"guild_id,omitempty"`

	// Sorting position of the channel
	Position int `json:"position,omitempty"`

	// Explicit permission overwrites for members and roles
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"`

	// The name of the channel (1-100 characters)
	Name string `json:"name,omitempty"`

	// The channel topic (0-1024 characters)
	Topic string `json:"topic,omitempty"`

	// Whether the channel is NSFW
	NSFW bool `json:"nsfw,omitempty"`

	// The ID of the last message sent in this channel (or thread for GUILD_FORUM channels) (may not point to an existing or valid message or thread)
	LastMessageID string `json:"last_message_id,omitempty"`

	// The bitrate (in bits) of the voice channel
	Bitrate int `json:"bitrate,omitempty"`

	// The user limit of the voice channel
	UserLimit int `json:"user_limit,omitempty"`

	// Amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages or manage_channel`, are unaffected
	RateLimitPerUser int `json:"rate_limit_per_user"`

	// The recipients of the DM
	Recipients []*User `json:"recipients,omitempty"`

	// Icon hash of the group DM
	Icon string `json:"icon,omitempty"`

	// ID of the creator of the group DM or thread
	OwnerID string `json:"owner_id,omitempty"`

	// Application ID of the group DM creator if it is bot-created
	ApplicationID string `json:"application_id,omitempty"`

	// For guild channels: ID of the parent category for a channel (each parent category can contain up to 50 channels), for threads: ID of the text channel this thread was created
	ParentID string `json:"parent_id,omitempty"`

	// When the last pinned message was pinned. This may be null in events such as GUILD_CREATE when a message is not pinned.
	LastPinTimestamp time.Time `json:"last_pin_timestamp,omitempty"`

	// Voice region ID for the voice channel, automatic when set to null
	RTCRegion string `json:"rtc_region,omitempty"`

	// The camera video quality mode of the voice channel, 1 when not present
	VideoQualityMode VideoQualityMode `json:"video_quality_mode,omitempty"`

	// An approximate count of messages in a thread, stops counting at 50
	MessageCount int `json:"message_count,omitempty"`

	// An approximate count of users in a thread, stops counting at 50
	MemberCount int `json:"member_count,omitempty"`

	// Thread-specific fields not needed by other channels
	ThreadMetadata *ThreadMetadata `json:"thread_metadata,omitempty"`

	// Thread member object for the current user, if they have joined the thread, only included on certain API endpoints
	Member *ThreadMember `json:"member,omitempty"`

	// Default duration that the clients (not the API) will use for newly created threads, in minutes, to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration,omitempty"`

	// Computed permissions for the invoking user in the channel, including overwrites, only included when part of the resolved data received on a slash command interaction
	Permissions string `json:"permissions,omitempty"`

	// Channel flags combined as a bitfield
	Flags ChannelFlags `json:"flags,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#channel-object-channel-types
type ChannelType int

const (
	ChannelTypeGuildText ChannelType = iota
	ChannelTypeDM
	ChannelTypeGuildVoice
	ChannelTypeGroupDM
	ChannelTypeGuildCategory
	ChannelTypeGuildNews
	ChannelTypeGuildNewsThread
	ChannelTypeGuildPublicThread
	ChannelTypeGuildPrivateThread
	ChannelTypeGuildStageVoice
	ChannelTypeDirectory
	ChannelTypeGuildForum
)

// https://discord.com/developers/docs/resources/channel#channel-object-video-quality-modes
type VideoQualityMode int

const (
	VideoQualityModeAuto VideoQualityMode = iota + 1
	VideoQualityModeFull
)

// https://discord.com/developers/docs/resources/channel#channel-object-channel-flags
type ChannelFlags int

const (
	ChannelFlagsPinned ChannelFlags = 1 << iota
)

// https://discord.com/developers/docs/resources/channel#message-object-message-structure
type Message struct {
	// ID of the message
	ID string `json:"id"`

	// ID of the channel the message was sent in
	ChannelID string `json:"channel_id"`

	// ID of the guild the message was sent in
	GuildID string `json:"guild_id"`

	// The author of this message (not guaranteed to be a valid user, see below)
	Author *User `json:"author"`

	// Member properties for this message's author
	Member *GuildMember `json:"member"`

	// Contents of the message
	Content string `json:"content"`

	// When this message was sent
	Timestamp time.Time `json:"timestamp"`

	// When this message was edited (or null if never)
	EditedTimestamp time.Time `json:"edited_timestamp"`

	// Whether this was a TTS message
	TTS bool `json:"tts"`

	// Whether this message mentions everyone
	MentionEveryone bool `json:"mention_everyone"`

	// Users specifically mentioned in the message
	Mentions []*MemberMention `json:"mentions"`

	// Roles specifically mentioned in this message
	MentionRoles []string `json:"mention_roles"`

	// Channels specifically mentioned in this message
	MentionChannels []*ChannelMention `json:"mention_channels"`

	// Any attached files
	Attachments []*Attachment `json:"attachments"`

	// Any embedded content
	Embeds []*Embed `json:"embeds"`

	// Reactions to the message
	Reactions []*Reaction `json:"reactions,omitempty"`

	// Used for validating a message was sent
	Nonce string `json:"nonce,omitempty"`

	// Whether this message is pinned
	Pinned bool `json:"pinned"`

	// If the message is generated by a webhook, this is the webhook's ID
	WebhookID string `json:"webhook_id,omitempty"`

	// Type of message
	Type MessageType `json:"type"`

	// Sent with Rich Presence-related chat embeds
	Activity *MessageActivity `json:"activity,omitempty"`

	// Sent with Rich Presence-related chat embeds
	Application *Application `json:"application,omitempty"`

	// If the message is an Interaction or application-owned webhook, this is the ID of the application
	ApplicationID string `json:"application_id,omitempty"`

	// Data showing the source of a crosspost, channel follow add, pin, or reply message
	MessageReference *MessageReference `json:"message_reference,omitempty"`

	// Message flags combined as a bitfield
	Flags MessageFlags `json:"flags,omitempty"`

	// The message associated with the message_reference
	ReferencedMessage *Message `json:"referenced_message"`

	// Sent if the message is a response to an Interaction
	Interaction *MessageInteraction `json:"interaction,omitempty"`

	// The thread that was started from this message, includes thread member object
	Thread *Channel `json:"thread,omitempty"`

	// Sent if the message contains components like buttons, action rows, or other interactive components
	// Components []*MessageComponent `json:"components,omitempty"`

	// Sent if the message contains stickers
	StickerItems []*StickerItem `json:"sticker_items,omitempty"`

	// Deprecated** the stickers sent with the message
	Stickers []*Sticker `json:"stickers,omitempty"`
}

type MemberMention struct {
	*User
	Member *GuildMember `json:"member"`
}

// https://discord.com/developers/docs/resources/channel#message-object-message-types
type MessageType int

const (
	MessageTypeDefault MessageType = iota
	MessageTypeRecipientAdd
	MessageTypeRecipientRemove
	MessageTypeCall
	MessageTypeChannelNameChange
	MessageTypeChannelIconChange
	MessageTypePinnedMessage
	MessageTypeGuildMemberJoin
	MessageTypeUserPremiumGuildSubscription
	MessageTypeUserPremiumGuildSubscriptionTier1
	MessageTypeUserPremiumGuildSubscriptionTier2
	MessageTypeUserPremiumGuildSubscriptionTier3
	MessageTypeChannelFollowAdd
	MessageTypeGuildDiscoveryDisqualified
	MessageTypeGuildDiscoveryRequalified
	MessageTypeGuildDiscoveryGracePeriodInitialWarning
	MessageTypeGuildDiscoveryGracePeriodFinalWarning
	MessageTypeThreadCreated
	MessageTypeReply
	MessageTypeChatInputCommand
	MessageTypeThreadStarterMessage
	MessageTypeGuildInviteReminder
	MessageTypeContextMenuCommand
)

// https://discord.com/developers/docs/resources/channel#message-object-message-activity-structure
type MessageActivity struct {
	// Type of message activity
	Type MessageActivityType `json:"type"`

	// Party_id from a Rich Presence event
	PartyID string `json:"party_id,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#message-object-message-activity-types
type MessageActivityType int

const (
	MessageActivityTypeJoin MessageActivityType = iota + 1
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	_
	MessageActivityTypeJoinRequest
)

// https://discord.com/developers/docs/resources/channel#message-object-message-flags
type MessageFlags int

const (
	MessageFlagsCrossposted MessageFlags = 1 << iota
	MessageFlagsIsCrosspost
	MessageFlagsSuppressEmbeds
	MessageFlagsSourceMessageDeleted
	MessageFlagsUrgent
	MessageFlagsHasThread
	MessageFlagsEphemeral
	MessageFlagsLoading
	MessageFlagsFailedToMentionSomeRolesInThread
)

// https://discord.com/developers/docs/resources/channel#message-reference-object-message-reference-structure
type MessageReference struct {
	// ID of the originating message
	MessageID string `json:"message_id,omitempty"`

	// ID of the originating message's channel
	ChannelID string `json:"channel_id"`

	// ID of the originating message's guild
	GuildID string `json:"guild_id,omitempty"`

	// When sending, whether to error if the referenced message doesn't exist instead of sending as a normal (non-reply) message, default true
	FailIfNotExists bool `json:"fail_if_not_exists,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#followed-channel-object-followed-channel-structure
type FollowedChannel struct {
	// Source channel ID
	ChannelID string `json:"channel_id"`

	// Created target webhook ID
	WebhookID string `json:"webhook_id"`
}

// https://discord.com/developers/docs/resources/channel#reaction-object-reaction-structure
type Reaction struct {
	// Times this emoji has been used to react
	Count int `json:"count"`

	// Whether the current user reacted using this emoji
	Me bool `json:"me"`

	// Emoji information
	Emoji *Emoji `json:"emoji"`
}

// https://discord.com/developers/docs/resources/channel#overwrite-object-overwrite-structure
type Overwrite struct {
	// Role or user ID
	ID string `json:"id"`

	// Either 0 (role) or 1 (member)
	Type OverwriteType `json:"type"`

	// Permission bit set
	Allow string `json:"allow"`

	// Permission bit set
	Deny string `json:"deny"`
}

type OverwriteType int

const (
	OverwriteTypeRole = iota
	OverwriteTypeMember
)

// https://discord.com/developers/docs/resources/channel#thread-metadata-object-thread-metadata-structure
type ThreadMetadata struct {
	// Whether the thread is archived
	Archived bool `json:"archived"`

	// Duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration"`

	// Timestamp when the thread's archive status was last changed, used for calculating recent activity
	ArchiveTimestamp time.Time `json:"archive_timestamp"`

	// Whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
	Locked bool `json:"locked"`

	// Whether non-moderators can add other non-moderators to a thread; only available on private threads
	Invitable bool `json:"invitable,omitempty"`

	// Timestamp when the thread was created; only populated for threads created after 2022-01-09
	CreateTimestamp time.Time `json:"create_timestamp,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#thread-member-object-thread-member-structure
type ThreadMember struct {
	// The ID of the thread
	ID string `json:"id"`

	// The ID of the user
	UserID string `json:"user_id"`

	// The time the current user last joined the thread
	JoinTimestamp time.Time `json:"join_timestamp"`

	// Any user-thread settings, currently only used for notifications
	Flags int `json:"flags"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-structure
type Embed struct {
	// Title of embed
	Title string `json:"title,omitempty"`

	// Type of embed (always "rich" for webhook embeds)
	Type string `json:"type,omitempty"`

	// Description of embed
	Description string `json:"description,omitempty"`

	// URL of embed
	URL string `json:"url,omitempty"`

	// Timestamp of embed content
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Color code of the embed
	Color int `json:"color,omitempty"`

	// Footer information
	Footer *EmbedFooter `json:"footer,omitempty"`

	// Image information
	Image *EmbedImage `json:"image,omitempty"`

	// Thumbnail information
	Thumbnail *EmbedThumbnail `json:"thumbnail,omitempty"`

	// Video information
	Video *EmbedVideo `json:"video,omitempty"`

	// Provider information
	Provider *EmbedProvider `json:"provider,omitempty"`

	// Author information
	Author *EmbedAuthor `json:"author,omitempty"`

	// Fields information
	Fields []*EmbedField `json:"fields,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-thumbnail-structure
type EmbedThumbnail struct {
	// Source URL of thumbnail (only supports HTTP(S) and attachments)
	URL string `json:"url"`

	// A proxied URL of the thumbnail
	ProxyURL string `json:"proxy_url,omitempty"`

	// Height of thumbnail
	Height int `json:"height,omitempty"`

	// Width of thumbnail
	Width int `json:"width,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-video-structure
type EmbedVideo struct {
	// Source URL of video
	URL string `json:"url,omitempty"`

	// A proxied URL of the video
	ProxyURL string `json:"proxy_url,omitempty"`

	// Height of video
	Height int `json:"height,omitempty"`

	// Width of video
	Width int `json:"width,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-image-structure
type EmbedImage struct {
	// Source URL of image (only supports HTTP(S) and attachments)
	URL string `json:"url"`

	// A proxied URL of the image
	ProxyURL string `json:"proxy_url,omitempty"`

	// Height of image
	Height int `json:"height,omitempty"`

	// Width of image
	Width int `json:"width,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-provider-structure
type EmbedProvider struct {
	// Name of provider
	Name string `json:"name,omitempty"`

	// URL of provider
	URL string `json:"url,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-author-structure
type EmbedAuthor struct {
	// Name of author
	Name string `json:"name"`

	// URL of author
	URL string `json:"url,omitempty"`

	// URL of author icon (only supports HTTP(S) and attachments)
	IconURL string `json:"icon_url,omitempty"`

	// A proxied URL of author icon
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-footer-structure
type EmbedFooter struct {
	// Footer text
	Text string `json:"text"`

	// URL of footer icon (only supports HTTP(S) and attachments)
	IconURL string `json:"icon_url,omitempty"`

	// A proxied URL of footer icon
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#embed-object-embed-field-structure
type EmbedField struct {
	// Name of the field
	Name string `json:"name"`

	// Value of the field
	Value string `json:"value"`

	// Whether or not this field should display inline
	Inline bool `json:"inline,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#attachment-object-attachment-structure
type Attachment struct {
	// Attachment ID
	ID string `json:"id"`

	// Name of file attached
	Filename string `json:"filename"`

	// Description for the file
	Description string `json:"description,omitempty"`

	// The attachment's media type
	ContentType string `json:"content_type,omitempty"`

	// Size of file in bytes
	Size int `json:"size"`

	// Source URL of file
	URL string `json:"url"`

	// A proxied URL of file
	ProxyURL string `json:"proxy_url"`

	// Height of file (if image)
	Height int `json:"height,omitempty"`

	// Width of file (if image)
	Width int `json:"width,omitempty"`

	// Whether this attachment is ephemeral
	Ephemeral bool `json:"ephemeral"`
}

// https://discord.com/developers/docs/resources/channel#channel-mention-object-channel-mention-structure
type ChannelMention struct {
	// ID of the channel
	ID string `json:"id"`

	// ID of the guild containing the channel
	GuildID string `json:"guild_id"`

	// The type of channel
	Type ChannelType `json:"type"`

	// The name of the channel
	Name string `json:"name"`
}

// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mentions-structure
type AllowedMentions struct {
	// An array of allowed mention types to parse from the content.
	Parse []AllowedMentionType `json:"parse"`

	// Array of role_ids to mention (Max size of 100)
	Roles []string `json:"roles"`

	// Array of user_ids to mention (Max size of 100)
	Users []string `json:"users"`

	// For replies, whether to mention the author of the message being replied to (default false)
	RepliedUser bool `json:"replied_user"`
}

// https://discord.com/developers/docs/resources/channel#allowed-mentions-object-allowed-mention-types
type AllowedMentionType string

const (
	AllowedMentionTypeRoleMentions     AllowedMentionType = "roles"
	AllowedMentionTypeUserMentions     AllowedMentionType = "users"
	AllowedMentionTypeEveryoneMentions AllowedMentionType = "everyone"
)
