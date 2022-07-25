package discord

import (
	"time"
)

// https://discord.com/developers/docs/resources/guild#guild-object-guild-structure
type Guild struct {
	// Guild ID
	ID string `json:"id"`

	// Guild name (2-100 characters, excluding trailing and leading whitespace)
	Name string `json:"name"`

	// Icon hash
	Icon string `json:"icon"`

	// Icon hash, returned when in the template object
	IconHash string `json:"icon_hash,omitempty"`

	// Splash hash
	Splash string `json:"splash"`

	// Discovery splash hash; only present for guilds with the "DISCOVERABLE" feature
	DiscoverySplash string `json:"discovery_splash"`

	// True if the user is the owner of the guild
	Owner bool `json:"owner"`

	// ID of owner
	OwnerID string `json:"owner_id"`

	// Total permissions for the user in the guild (excludes overwrites)
	Permissions string `json:"permissions"`

	// ID of AFK channel
	AFKChannelID string `json:"afk_channel_id"`

	// AFK timeout in seconds
	AFKTimeout int `json:"afk_timeout"`

	// True if the server widget is enabled
	WidgetEnabled bool `json:"widget_enabled,omitempty"`

	// The channel ID that the widget will generate an invite to, or null if set to no invite
	WidgetChannelID string `json:"widget_channel_id,omitempty"`

	// Verification level required for the guild
	VerificationLevel VerificationLevel `json:"verification_level"`

	// Default message notifications level
	DefaultMessageNotifications DefaultMessageNotifications `json:"default_message_notifications"`

	// Explicit content filter level
	ExplicitContentFilter ExplicitContentFilter `json:"explicit_content_filter"`

	// Roles in the guild
	Roles []*Role `json:"roles"`

	// Custom guild emojis
	Emojis []*Emoji `json:"emojis"`

	// Enabled guild features
	Features []GuildFeature `json:"features"`

	// Required MFA level for the guild
	MFALevel MFALevel `json:"mfa_level"`

	// Application ID of the guild creator if it is bot-created
	ApplicationID string `json:"application_id"`

	// The ID of the channel where guild notices such as welcome messages and boost events are posted
	SystemChannelID string `json:"system_channel_id"`

	// System channel flags
	SystemChannelFlags SystemChannelFlags `json:"system_channel_flags"`

	// The ID of the channel where Community guilds can display rules and/or guidelines
	RulesChannelID string `json:"rules_channel_id"`

	// The maximum number of presences for the guild (`null is always returned, apart from the largest of guilds)
	MaxPresences int `json:"max_presences,omitempty"`

	// The maximum number of members for the guild
	MaxMembers int `json:"max_members,omitempty"`

	// The vanity URL code for the guild
	VanityURLCode string `json:"vanity_url_code"`

	// The description of a guild
	Description string `json:"description"`

	// Banner hash
	Banner string `json:"banner"`

	// Premium tier (Server Boost level)
	PremiumTier PremiumTier `json:"premium_tier"`

	// The number of boosts this guild currently has
	PremiumSubscriptionCount int `json:"premium_subscription_count,omitempty"`

	// The preferred locale of a Community guild; used in server discovery and notices from Discord, and sent in interactions; defaults to "en-US"
	PreferredLocale string `json:"preferred_locale"`

	// The ID of the channel where admins and moderators of Community guilds receive notices from Discord
	PublicUpdatesChannelID string `json:"public_updates_channel_id"`

	// The maximum amount of users in a video channel
	MaxVideoChannelUsers int `json:"max_video_channel_users,omitempty"`

	// Approximate number of members in this guild, returned from the GET /guilds/<id> endpoint when with_counts is true
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`

	// Approximate number of non-offline members in this guild, returned from the GET /guilds/<id> endpoint when with_counts is true
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`

	// The welcome screen of a Community guild, shown to new members, returned in an Invite's guild object
	WelcomeScreen *WelcomeScreen `json:"welcome_screen,omitempty"`

	// Guild NSFW level
	NSFWLevel NSFWLevel `json:"nsfw_level"`

	// Custom guild stickers
	Stickers []*Sticker `json:"stickers,omitempty"`

	// Whether the guild has the boost progress bar enabled
	PremiumProgressBarEnabled bool `json:"premium_progress_bar_enabled"`
}

// https://discord.com/developers/docs/resources/guild#guild-object-default-message-notification-level
type DefaultMessageNotifications int

const (
	DefaultMessageNotificationsAllMessages DefaultMessageNotifications = iota + 1
	DefaultMessageNotificationsOnlyMentions
)

// https://discord.com/developers/docs/resources/guild#guild-object-explicit-content-filter-level
type ExplicitContentFilter int

const (
	ExplicitContentFilterDisabled ExplicitContentFilter = iota + 1
	ExplicitContentFilterMembersWithoutRoles
	ExplicitContentFilterAllMembers
)

// https://discord.com/developers/docs/resources/guild#guild-object-mfa-level
type MFALevel int

const (
	MFALevelNone MFALevel = iota + 1
	MFALevelElevated
)

// https://discord.com/developers/docs/resources/guild#guild-object-verification-level
type VerificationLevel int

const (
	VerificationLevelNone VerificationLevel = iota + 1
	VerificationLevelLow
	VerificationLevelMedium
	VerificationLevelHigh
	VerificationLevelVeryHigh
)

// https://discord.com/developers/docs/resources/guild#guild-object-guild-nsfw-level
type NSFWLevel int

const (
	NSFWLevelDefault NSFWLevel = iota + 1
	NSFWLevelExplicit
	NSFWLevelSafe
	NSFWLevelAgeRestricted
)

// https://discord.com/developers/docs/resources/guild#guild-object-premium-tier
type PremiumTier int

const (
	PremiumTierNone PremiumTier = iota + 1
	PremiumTier1
	PremiumTier2
	PremiumTier3
)

// https://discord.com/developers/docs/resources/guild#guild-object-system-channel-flags
type SystemChannelFlags int

const (
	SystemChannelFlagsSuppressJoinNotifications SystemChannelFlags = 1 << iota
	SystemChannelFlagsSuppressPremiumNotifications
	SystemChannelFlagsSuppressGuildReminderNotifications
	SystemChannelFlagsSuppressJoinNotificationReplies
)

// https://discord.com/developers/docs/resources/guild#guild-object-guild-features
type GuildFeature string

const (
	GuildFeatureAnimatedBanner                GuildFeature = "ANIMATED_BANNER"
	GuildFeatureAnimatedIcon                  GuildFeature = "ANIMATED_ICON"
	GuildFeatureBanner                        GuildFeature = "BANNER"
	GuildFeatureCommerce                      GuildFeature = "COMMERCE"
	GuildFeatureCommunity                     GuildFeature = "COMMUNITY"
	GuildFeatureDiscoverable                  GuildFeature = "DISCOVERABLE"
	GuildFeatureFeaturable                    GuildFeature = "FEATURABLE"
	GuildFeatureInviteSplash                  GuildFeature = "INVITE_SPLASH"
	GuildFeatureMemberVerificationGateEnabled GuildFeature = "MEMBER_VERIFICATION_GATE_ENABLED"
	GuildFeatureMonetizationEnabled           GuildFeature = "MONETIZATION_ENABLED"
	GuildFeatureMoreStickers                  GuildFeature = "MORE_STICKERS"
	GuildFeatureNews                          GuildFeature = "NEWS"
	GuildFeaturePartnered                     GuildFeature = "PARTNERED"
	GuildFeaturePreviewEnabled                GuildFeature = "PREVIEW_ENABLED"
	GuildFeaturePrivateThreads                GuildFeature = "PRIVATE_THREADS"
	GuildFeatureRoleIcons                     GuildFeature = "ROLE_ICONS"
	GuildFeatureSevenDayThreadArchive         GuildFeature = "SEVEN_DAY_THREAD_ARCHIVE"
	GuildFeatureThreeDayThreadArchieve        GuildFeature = "THREE_DAY_THREAD_ARCHIVE"
	GuildFeatureTicketedEventsEnabled         GuildFeature = "TICKETED_EVENTS_ENABLED"
	GuildFeatureVanityURL                     GuildFeature = "VANITY_URL"
	GuildFeatureVerified                      GuildFeature = "VERIFIED"
	GuildFeatureVIPRegions                    GuildFeature = "VIP_REGIONS"
	GuildFeatureWelcomeScreenEnabled          GuildFeature = "WELCOME_SCREEN_ENABLED"
)

// https://discord.com/developers/docs/resources/guild#unavailable-guild-object-example-unavailable-guild
type UnavailableGuild struct {
	// Guild ID
	ID string `json:"id"`

	// True if the guild is unavailable
	Unavailable bool `json:"unavailable"`
}

// https://discord.com/developers/docs/resources/guild#guild-preview-object-guild-preview-structure
type GuildPreview struct {
	// Guild ID
	ID string `json:"id"`

	// Guild name (2-100 characters)
	Name string `json:"name"`

	// Icon hash
	Icon string `json:"icon"`

	// Splash hash
	Splash string `json:"splash"`

	// Discovery splash hash
	DiscoverySplash string `json:"discovery_splash"`

	// Custom guild emojis
	Emojis []*Emoji `json:"emojis"`

	// Enabled guild features
	Features []string `json:"features"`

	// Approximate number of members in this guild
	ApproximateMemberCount int `json:"approximate_member_count"`

	// Approximate number of online members in this guild
	ApproximatePresenceCount int `json:"approximate_presence_count"`

	// The description for the guild
	Description string `json:"description"`

	// Custom guild stickers
	Stickers []*Sticker `json:"stickers"`
}

// https://discord.com/developers/docs/resources/guild#guild-widget-settings-object-guild-widget-settings-structure
type GuildWidgetSettings struct {
	// Whether the widget is enabled
	Enabled bool `json:"enabled"`

	// The widget channel ID
	ChannelID string `json:"channel_id"`
}

// https://discord.com/developers/docs/resources/guild#get-guild-widget-object-get-guild-widget-structure
type GetGuildWidget struct {
	// Guild ID
	ID string `json:"id"`

	// Guild name (2-100 characters)
	Name string `json:"name"`

	// Instant invite for the guilds specified widget invite channel
	InstantInvite string `json:"instant_invite"`

	// Voice and stage channels which are accessible by @everyone
	Channels []*Channel `json:"channels"`

	// Special widget user objects that includes users presence (Limit 100)
	Members []*User `json:"members"`

	// Number of online members in this guild
	PresenceCount int `json:"presence_count"`
}

// https://discord.com/developers/docs/resources/guild#guild-member-object-guild-member-structure
type GuildMember struct {
	// The user this guild member represents
	User *User `json:"user,omitempty"`

	// This user's guild nickname
	Nick string `json:"nick,omitempty"`

	// The member's guild avatar hash
	Avatar string `json:"avatar,omitempty"`

	// Array of role object IDs
	Roles []string `json:"roles"`

	// When the user joined the guild
	JoinedAt time.Time `json:"joined_at"`

	// When the user started boosting the guild
	PremiumSince time.Time `json:"premium_since,omitempty"`

	// Whether the user is deafened in voice channels
	Deaf bool `json:"deaf"`

	// Whether the user is muted in voice channels
	Mute bool `json:"mute"`

	// Whether the user has not yet passed the guild's Membership Screening requirements
	Pending bool `json:"pending,omitempty"`

	// Total permissions of the member in the channel, including overwrites, returned when in the interaction object
	Permissions string `json:"permissions,omitempty"`

	// When the user's timeout will expire and the user will be able to communicate in the guild again, null or a time in the past if the user is not timed out
	CommunicationDisabledUntil time.Time `json:"communication_disabled_until,omitempty"`
}

// https://discord.com/developers/docs/resources/guild#integration-object-integration-structure
type Integration struct {
	// Integration ID
	ID string `json:"id"`

	// Integration name
	Name string `json:"name"`

	// Integration type (twitch, youtube, or discord)
	Type string `json:"type"`

	// Is this integration enabled
	Enabled bool `json:"enabled"`

	// Is this integration syncing
	Syncing bool `json:"syncing"`

	// ID that this integration uses for "subscribers"
	RoleID string `json:"role_id"`

	// Whether emoticons should be synced for this integration (twitch only currently)
	EnableEmoticons bool `json:"enable_emoticons"`

	// The behavior of expiring subscribers
	ExpireBehavior IntegrationExpireBehavior `json:"expire_behavior"`

	// The grace period (in days) before expiring subscribers
	ExpireGracePeriod int `json:"expire_grace_period"`

	// User for this integration
	User *User `json:"user"`

	// Integration account information
	Account *IntegrationAccount `json:"account"`

	// When this integration was last synced
	SyncedAt time.Time `json:"synced_at"`

	// How many subscribers this integration has
	SubscriberCount int `json:"subscriber_count"`

	// Has this integration been revoked
	Revoked bool `json:"revoked"`

	// The bot/OAuth2 application for discord integrations
	Application *Application `json:"application,omitempty"`
}

// https://discord.com/developers/docs/resources/guild#integration-object-integration-expire-behaviors
type IntegrationExpireBehavior int

const (
	IntegrationExpireBehaviorRemoveRole IntegrationExpireBehavior = iota
	IntegrationExpireBehaviorKick
)

// https://discord.com/developers/docs/resources/guild#integration-account-object-integration-account-structure
type IntegrationAccount struct {
	// ID of the account
	ID string `json:"id"`

	// Name of the account
	Name string `json:"name"`
}

// https://discord.com/developers/docs/resources/guild#integration-application-object-integration-application-structure
type IntegrationApplication struct {
	// The ID of the app
	ID string `json:"id"`

	// The name of the app
	Name string `json:"name"`

	// The icon hash of the app
	Icon string `json:"icon"`

	// The description of the app
	Description string `json:"description"`

	// The bot associated with this application
	Bot *User `json:"bot,omitempty"`
}

// https://discord.com/developers/docs/resources/guild#ban-object-ban-structure
type Ban struct {
	// The reason for the ban
	Reason string `json:"reason"`

	// The banned user
	User *User `json:"user"`
}

// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-structure
type WelcomeScreen struct {
	// The server description shown in the welcome screen
	Description string `json:"description"`

	// The channels shown in the welcome screen, up to 5
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels"`
}

// https://discord.com/developers/docs/resources/guild#welcome-screen-object-welcome-screen-channel-structure
type WelcomeScreenChannel struct {
	// The channel's ID
	ChannelID string `json:"channel_id"`

	// The description shown for the channel
	Description string `json:"description"`

	// The emoji ID, if the emoji is custom
	EmojiID string `json:"emoji_id"`

	// The emoji name if custom, the unicode character if standard, or null if no emoji is set
	EmojiName string `json:"emoji_name"`
}
