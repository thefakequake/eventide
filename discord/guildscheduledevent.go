package eventide

import (
	"time"
)

// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-structure
type GuildScheduledEvent struct {
	// The ID of the scheduled event
	ID string `json:"id"`

	// The guild ID which the scheduled event belongs to
	GuildID string `json:"guild_id"`

	// The channel ID in which the scheduled event will be hosted, or null if scheduled entity type is EXTERNAL
	ChannelID string `json:"channel_id"`

	// The ID of the user that created the scheduled event
	CreatorID string `json:"creator_id"`

	// The name of the scheduled event (1-100 characters)
	Name string `json:"name"`

	// The description of the scheduled event (1-1000 characters)
	Description string `json:"description,omitempty"`

	// The time the scheduled event will start
	ScheduledStartTime time.Time `json:"scheduled_start_time"`

	// The time the scheduled event will end, required if entity_type is EXTERNAL
	ScheduledEndTime time.Time `json:"scheduled_end_time"`

	// The privacy level of the scheduled event
	PrivacyLevel *PrivacyLevel `json:"privacy_level"`

	// The status of the scheduled event
	Status *EventStatus `json:"status"`

	// The type of the scheduled event
	EntityType *ScheduledEntityType `json:"entity_type"`

	// The ID of an entity associated with a guild scheduled event
	EntityID string `json:"entity_id"`

	// Additional metadata for the guild scheduled event
	EntityMetadata *EntityMetadata `json:"entity_metadata"`

	// The user that created the scheduled event
	Creator *User `json:"creator,omitempty"`

	// The number of users subscribed to the scheduled event
	UserCount int `json:"user_count,omitempty"`

	// The cover image hash of the scheduled event
	Image string `json:"image,omitempty"`
}

// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-object-guild-scheduled-event-entity-metadata
type GuildScheduledEventEntityMetadata struct {
	// Location of the event (1-100 characters)
	Location string `json:"location"`
}

// https://discord.com/developers/docs/resources/guild-scheduled-event#guild-scheduled-event-user-object-guild-scheduled-event-user-structure
type GuildScheduledEventUser struct {
	// The scheduled event ID which the user subscribed to
	GuildScheduledEventID string `json:"guild_scheduled_event_id"`

	// User which subscribed to an event
	User *User `json:"user"`

	// Guild member data for this user for the guild which this event belongs to, if any
	Member *GuildMember `json:"member,omitempty"`
}
