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
}

func (m *Message) Reference() *MessageReference {
	return &MessageReference{
		MessageID: m.ID,
		ChannelID: m.ChannelID,
		GuildID:   m.GuildID,
	}
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
	_
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
	ChannelID string `json:"channel_id,omitempty"`

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

// https://discord.com/developers/docs/resources/channel#modify-channel
type ModifyChannel struct {
	// 1-100 character channel name
	Name string `json:"name,omitempty"`

	// Base64 encoded icon, only for group DMs
	Icon string `json:"icon,omitempty"`

	// The type of channel; only conversion between text and news is supported and only in guilds with the "NEWS" feature
	Type ChannelType `json:"type,omitempty"`

	// The position of the channel in the left-hand listing
	Position int `json:"position,omitempty"`

	// 0-1024 character channel topic
	Topic string `json:"topic,omitempty"`

	// Whether the channel is NSFW
	NSFW bool `json:"nsfw,omitempty"`

	// Amount of seconds a user has to wait before sending another message (0-21600); bots, as well as users with the permission manage_messages or manage_channel, are unaffected
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`

	// The bitrate (in bits) of the voice or stage channel; min 8000
	Bitrate int `json:"bitrate,omitempty"`

	// The user limit of the voice channel; 0 refers to no limit, 1 to 99 refers to a user limit
	UserLimit int `json:"user_limit,omitempty"`

	// Channel or category-specific permissions
	PermissionOverwrites []*Overwrite `json:"permission_overwrites,omitempty"`

	// ID of the new parent category for a channel
	ParentID string `json:"parent_id,omitempty"`

	// Channel voice region ID, automatic when set to null
	RTCRegion string `json:"rtc_region,omitempty"`

	// The camera video quality mode of the voice channel
	VoiceQualityMode VideoQualityMode `json:"voice_quality_mode,omitempty"`

	// The default duration that the clients use (not the API) for newly created threads in the channel, in minutes, to automatically archive the thread after recent activity
	DefaultAutoArchiveDuration int `json:"default_auto_archive_duration,omitempty"`

	// BELOW ARE ALL THREAD ONLY

	// Whether the thread is archived
	Archived bool `json:"archived,omitempty"`

	// Duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// Whether the thread is locked; when a thread is locked, only users with MANAGE_THREADS can unarchive it
	Locked bool `json:"locked,omitempty"`

	// Whether non-moderators can add other non-moderators to a thread; only available on private threads
	Invitable bool `json:"invitable,omitempty"`

	// Channel flags combined as a bitfield; PINNED can only be set for threads in forum channels
	Flags ChannelFlags `json:"flags"`
}

// https://discord.com/developers/docs/resources/channel#get-channel-messages
type GetChannelMessages struct {
	// Get messages around this message ID
	Around string `json:"around,omitempty"`

	// Get messages before this message ID
	Before string `json:"before,omitempty"`

	// Get messages after this message ID
	After string `json:"after,omitempty"`

	// Max number of messages to return (1-100)
	Limit int `json:"limit,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#create-message
type CreateMessage struct {
	// Message contents (up to 2000 characters)
	Content string `json:"content,omitempty"`

	// True if this is a TTS message
	TTS bool `json:"tts,omitempty"`

	// Embedded rich content (up to 6000 characters)
	Embeds []*Embed `json:"embeds,omitempty"`

	// Allowed mentions for the message
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	// Include to make your message a reply
	MessageReference *MessageReference `json:"message_reference,omitempty"`

	// Components to include with the message
	// Components []*MessageComponent `json:"components"`

	// IDs of up to 3 stickers in the server to send in the message
	StickerIDs []string `json:"sticker_ids,omitempty"`

	// Contents of the file being sent. See Uploading Files
	// Files []*File `json:"files"`

	// JSON-encoded body of non-file params, only for multipart/form-data requests
	PayloadJSON string `json:"payload_json,omitempty"`

	// Attachment objects with filename and description.
	Attachments []*Attachment `json:"attachments,omitempty"`

	// Message flags combined as a bitfield (only SUPPRESS_EMBEDS can be set)
	Flags MessageFlags `json:"flags,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#get-reactions
type GetReactions struct {
	// Get users after this user ID
	After string `json:"after,omitempty"`

	// Max number of users to return (1-100)
	Limit int `json:"limit,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#edit-message
type EditMessage struct {
	// Message contents (up to 2000 characters)
	Content string `json:"content,omitempty"`

	// Embedded rich content (up to 6000 characters)
	Embeds []*Embed `json:"embeds,omitempty"`

	// Edit the flags of a message (only SUPPRESS_EMBEDS can currently be set/unset)
	Flags MessageFlags `json:"flags,omitempty"`

	// Allowed mentions for the message
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	// Components to include with the message
	// Components []*MessageComponent `json:"components,omitempty"`

	// Contents of the file being sent/edited
	// Files []*File `json:"files,omitempty"`

	// JSON-encoded body of non-file params (multipart/form-data only)
	PayloadJSON string `json:"payload_json,omitempty"`

	// 	Attached files to keep and possible descriptions for new files
	Attachments []*Attachment `json:"attachments,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#bulk-delete-messages
type BulkDeleteMessages struct {
	// An array of message IDs to delete (2-100)
	Messages []string `json:"messages"`
}

// https://discord.com/developers/docs/resources/channel#edit-channel-permissions
type EditChannelPermissions struct {
	// The bitwise value of all allowed permissions (default "0")
	Allow string `json:"allow,omitempty"`
	// The bitwise value of all disallowed permissions (default "0")
	Deny string `json:"deny,omitempty"`
	// 0 for a role or 1 for a member
	Type OverwriteType `json:"type"`
}

// https://discord.com/developers/docs/resources/channel#create-channel-invite
type CreateChannelInvite struct {
	// Duration of invite in seconds before expiry, or 0 for never. between 0 and 604800 (7 days)
	MaxAge int `json:"max_age,omitempty"`

	// Max number of uses or 0 for unlimited. between 0 and 100
	MaxUses int `json:"max_uses,omitempty"`

	// Whether this invite only grants temporary membership
	Temporary bool `json:"temporary,omitempty"`

	// If true, don't try to reuse a similar invite (useful for creating many unique one time use invites)
	Unique bool `json:"unique,omitempty"`

	// The type of target for this voice channel invite
	TargetType InviteTargetType `json:"target_type,omitempty"`

	// The ID of the user whose stream to display for this invite, required if target_type is 1, the user must be streaming in the channel
	TargetUserID string `json:"target_user_id,omitempty"`

	// The ID of the embedded application to open for this invite, required if target_type is 2, the application must have the EMBEDDED flag
	TargetApplicationID string `json:"target_application_id,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#follow-news-channel
type FollowNewsChannel struct {
	// ID of target channel
	WebhookChannelID string `json:"webhook_channel_id"`
}

// https://discord.com/developers/docs/resources/channel#start-thread-from-message
type StartThreadFromMessage struct {
	// 1-100 character channel name
	Name string `json:"name"`

	// Duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// Amount of seconds a user has to wait before sending another message (0-21600)
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#start-thread-without-message
type StartThreadWithoutMessage struct {
	// 1-100 character channel name
	Name string `json:"name"`

	// Duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// The type of thread to create
	Type ChannelType `json:"type,omitempty"`

	// Whether non-moderators can add other non-moderators to a thread; only available when creating a private thread
	Invitable bool `json:"invitable,omitempty"`

	// Amount of seconds a user has to wait before sending another message (0-21600)
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel
type StartThreadInForumChannel struct {
	// 1-100 character channel name
	Name string `json:"name"`

	// Duration in minutes to automatically archive the thread after recent activity, can be set to: 60, 1440, 4320, 10080
	AutoArchiveDuration int `json:"auto_archive_duration,omitempty"`

	// Amount of seconds a user has to wait before sending another message (0-21600)
	RateLimitPerUser int `json:"rate_limit_per_user,omitempty"`

	// Contents of the first message in the forum thread
	Message *ForumThreadMessageParams
}

type ForumChannelThreadCreate struct {
	*Channel

	// The message that was sent inside the thread on creation
	Message *Message `json:"message"`
}

// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel-forum-thread-message-params-object
type ForumThreadMessageParams struct {
	// Message contents (up to 2000 characters)
	Content string `json:"content"`

	// Embedded rich content (up to 6000 characters)
	Embeds []*Embed `json:"embeds,omitempty"`

	// Allowed mentions for the message
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	// Components to include with the message
	// Components []*MessageComponent `json:"components,omitempty"`

	// IDs of up to 3 stickers in the server to send in the message
	StickerIDs []string `json:"sticker_ids"`

	// Contents of the file being sent. See Uploading Files
	// Files[N] *FileContents `json:"files[n]"`

	// JSON-encoded body of non-file params, only for multipart/form-data requests. See Uploading Files
	PayloadJSON string `json:"payload_json,omitempty"`

	// Attachment objects with filename and description`. See Uploading Files
	Attachments []*Attachment `json:"attachments,omitempty"`

	// Message flags combined as a bitfield (only SUPPRESS_EMBEDS can be set)
	Flags int `json:"flags,omitempty"`
}

// https://discord.com/developers/docs/resources/channel#list-public-archived-threads
type ListArchivedThreads struct {
	// Returns threads before this timestamp
	Before time.Time `json:"before,omitempty"`

	// Optional maximum number of threads to return
	Limit int `json:"limit,omitempty"`
}

type ArchivedThreads struct {
	// The public, archived threads
	Threads []*Channel `json:"threads"`

	// A thread member object for each returned thread the current user has joined
	Members []*ThreadMember `json:"members"`

	// Whether there are potentially additional threads that could be returned on a subsequent call
	HasMore bool `json:"has_more"`
}
