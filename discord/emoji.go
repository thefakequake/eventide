package discord

// https://discord.com/developers/docs/resources/emoji#emoji-object-emoji-structure
type Emoji struct {
	// Emoji ID
	ID string `json:"id"`

	// Emoji name
	Name string `json:"name"`

	// Roles allowed to use this emoji
	Roles []string `json:"roles,omitempty"`

	// User that created this emoji
	User *User `json:"user,omitempty"`

	// Whether this emoji must be wrapped in colons
	RequireColons bool `json:"require_colons,omitempty"`

	// Whether this emoji is managed
	Managed bool `json:"managed,omitempty"`

	// Whether this emoji is animated
	Animated bool `json:"animated,omitempty"`

	// Whether this emoji can be used, may be false due to loss of Server Boosts
	Available bool `json:"available,omitempty"`
}

// https://discord.com/developers/docs/resources/emoji#create-guild-emoji
type CreateGuildEmoji struct {
	// Name of the emoji
	Name string `json:"name"`

	// The 128x128 emoji image
	Image string `json:"image"`

	// Roles allowed to use this emoji
	Roles []string `json:"roles"`
}

// https://discord.com/developers/docs/resources/emoji#modify-guild-emoji
type ModifyGuildEmoji struct {
	// Name of the emoji
	Name string `json:"name,omitempty"`

	// Roles allowed to use this emoji
	Roles []string `json:"roles,omitempty"`
}
