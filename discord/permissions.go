package eventide

// https://discord.com/developers/docs/topics/permissions#role-object-role-structure
type Role struct {
	// Role ID
	ID string `json:"id"`

	// Role name
	Name string `json:"name"`

	// Integer representation of hexadecimal color code
	Color int `json:"color"`

	// If this role is pinned in the user listing
	Hoist bool `json:"hoist"`

	// Role icon hash
	Icon string `json:"icon,omitempty"`

	// Role unicode emoji
	UnicodeEmoji string `json:"unicode_emoji,omitempty"`

	// Position of this role
	Position int `json:"position"`

	// Permission bit set
	Permissions string `json:"permissions"`

	// Whether this role is managed by an integration
	Managed bool `json:"managed"`

	// Whether this role is mentionable
	Mentionable bool `json:"mentionable"`

	// The tags this role has
	Tags *RoleTags `json:"tags,omitempty"`
}

// https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	// The ID of the bot this role belongs to
	BotID string `json:"bot_id,omitempty"`

	// The ID of the integration this role belongs to
	IntegrationID string `json:"integration_id,omitempty"`

	// Whether this is the guild's premium subscriber role
	PremiumSubscriber bool `json:"premium_subscriber,omitempty"`
}
