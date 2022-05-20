package eventide

// https://discord.com/developers/docs/resources/webhook#webhook-object-webhook-structure
type Webhook struct {
	// The ID of the webhook
	ID string `json:"id"`

	// The type of the webhook
	Type int `json:"type"`

	// The guild ID this webhook is for, if any
	GuildID string `json:"guild_id,omitempty"`

	// The channel ID this webhook is for, if any
	ChannelID string `json:"channel_id"`

	// The user this webhook was created by (not returned when getting a webhook with its token)
	User *User `json:"user,omitempty"`

	// The default name of the webhook
	Name string `json:"name"`

	// The default user avatar hash of the webhook
	Avatar string `json:"avatar"`

	// The secure token of the webhook (returned for Incoming Webhooks)
	Token string `json:"token,omitempty"`

	// The bot/OAuth2 application that created this webhook
	ApplicationID string `json:"application_id"`

	// The guild of the channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceGuild *Guild `json:"source_guild,omitempty"`

	// The channel that this webhook is following (returned for Channel Follower Webhooks)
	SourceChannel *Channel `json:"source_channel,omitempty"`

	// The URL used for executing the webhook (returned by the webhooks OAuth2 flow)
	URL string `json:"url,omitempty"`
}
