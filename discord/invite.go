package discord

import (
	"time"
)

// https://discord.com/developers/docs/resources/invite#invite-object-invite-structure
type Invite struct {
	// The invite code (unique ID)
	Code string `json:"code"`

	// The guild this invite is for
	Guild *Guild `json:"guild,omitempty"`

	// The channel this invite is for
	Channel *Channel `json:"channel"`

	// The user who created the invite
	Inviter *User `json:"inviter,omitempty"`

	// The type of target for this voice channel invite
	TargetType InviteTargetType `json:"target_type,omitempty"`

	// The user whose stream to display for this voice channel stream invite
	TargetUser *User `json:"target_user,omitempty"`

	// The embedded application to open for this voice channel embedded application invite
	TargetApplication *Application `json:"target_application,omitempty"`

	// Approximate count of online members, returned from the GET /invites/<code> endpoint when with_counts is true
	ApproximatePresenceCount int `json:"approximate_presence_count,omitempty"`

	// Approximate count of total members, returned from the GET /invites/<code> endpoint when with_counts is true
	ApproximateMemberCount int `json:"approximate_member_count,omitempty"`

	// The expiration date of this invite, returned from the GET /invites/<code> endpoint when with_expiration is true
	ExpiresAt time.Time `json:"expires_at,omitempty"`

	// Stage instance data if there is a public Stage instance in the Stage channel this invite is for (deprecated)
	StageInstance *InviteStageInstance `json:"stage_instance,omitempty"`

	// Guild scheduled event data, only included if guild_scheduled_event_id contains a valid guild scheduled event ID
	GuildScheduledEvent *GuildScheduledEvent `json:"guild_scheduled_event,omitempty"`
}

// https://discord.com/developers/docs/resources/invite#invite-object-invite-target-types
type InviteTargetType int

const (
	InviteTargetTypeStream InviteTargetType = iota + 1
	InviteTargetTypeEmbeddedApplication
)

// https://discord.com/developers/docs/resources/invite#invite-metadata-object-invite-metadata-structure
type InviteMetadataStructure struct {
	// Number of times this invite has been used
	Uses int `json:"uses"`

	// Max number of times this invite can be used
	MaxUses int `json:"max_uses"`

	// Duration (in seconds) after which the invite expires
	MaxAge int `json:"max_age"`

	// Whether this invite only grants temporary membership
	Temporary bool `json:"temporary"`

	// When this invite was created
	CreatedAt time.Time `json:"created_at"`
}

// https://discord.com/developers/docs/resources/invite#invite-stage-instance-object-invite-stage-instance-structure
type InviteStageInstance struct {
	// The members speaking in the Stage
	Members []*GuildMember `json:"members"`

	// The number of users in the Stage
	ParticipantCount int `json:"participant_count"`

	// The number of users speaking in the Stage
	SpeakerCount int `json:"speaker_count"`

	// The topic of the Stage instance (1-120 characters)
	Topic string `json:"topic"`
}
