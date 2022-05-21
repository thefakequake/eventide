package discord

// https://discord.com/developers/docs/resources/stage-instance#stage-instance-object-stage-instance-structure
type StageInstance struct {
	// The ID of this Stage instance
	ID string `json:"id"`

	// The guild ID of the associated Stage channel
	GuildID string `json:"guild_id"`

	// The ID of the associated Stage channel
	ChannelID string `json:"channel_id"`

	// The topic of the Stage instance (1-120 characters)
	Topic string `json:"topic"`

	// The privacy level of the Stage instance
	PrivacyLevel int `json:"privacy_level"`

	// Whether or not Stage Discovery is disabled (deprecated)
	DiscoverableDisabled bool `json:"discoverable_disabled"`

	// The ID of the scheduled event for this Stage instance
	GuildScheduledEventID string `json:"guild_scheduled_event_id"`
}
