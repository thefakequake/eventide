package eventide

import (
	"time"
)

type GatewayURL struct {
	URL string `json:"url"`
}


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

type ActivityType int

const (
	ActivityTypeGame ActivityType = iota
	ActivityTypeStreaming
	ActivityTypeListening
	ActivityTypeWatching
	ActivityTypeCustom
	ActivityTypeCompeting
)

type ActivityEmoji struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	Animated bool   `json:"animated"`
}

type ActivityTimestamp struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type ActivityParty struct {
	ID   string `json:"id"`
	Size [2]int `json:"size"`
}

type ActivityAssets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

type ActivitySecrets struct {
	Join     string `json:"join"`
	Spectate string `json:"spectate"`
	Match    string `json:"match"`
}

type ActivityButton struct {
	Label string `json:"label"`
	URL   string `json:"url"`
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

type PrivacyLevel int

const (
	PrivacyLevelPublic PrivacyLevel = iota
	PrivacyLevelGuildOnly
)

type StickerType int

const (
	StickerTypeStandard StickerType = iota
	StickerTypeGuild
)

type StickerFormatType int

const (
	StickerFormatTypePNG    StickerFormatType = iota
	StickerFormatTypeAPNG   StickerFormatType = iota
	StickerFormatTypeLottie StickerFormatType = iota
)

type GuildScheduledEventPrivacyLevel int

const (
	GuildScheduledEventPrivacyLevelGuildOnly GuildScheduledEventPrivacyLevel = 2
)

type GuildScheduledEventEntityType int

const (
	GuildScheduledEventEntityTypeStageInstance GuildScheduledEventEntityType = iota
	GuildScheduledEventEntityTypeVoice
	GuildScheduledEventEntityTypeExternal
)

type GuildScheduledEventStatus int

const (
	GuildScheduledEventStatusScheduled GuildScheduledEventStatus = iota
	GuildScheduledEventStatusActive
	GuildScheduledEventStatusCompleted
	GuildScheduledEventStatusCanceled
)

type EntityMetadata struct {
	Location string `json:"location"`
}

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

type MessageSend struct {
	Content string   `json:"content,omitempty"`
	Embeds  []*Embed `json:"embeds"`
	TTS     bool     `json:"tts"`
}

type MessageActivityType int

const (
	MessageActivityTypeJoin MessageActivityType = iota
	MessageActivityTypeSpectate
	MessageActivityTypeListen
	MessageActivityTypeJoinRequest
)
