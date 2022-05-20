package eventide

import (
	"time"
)

// https://discord.com/developers/docs/resources/guild-template#guild-template-object-guild-template-structure
type GuildTemplate struct {
	// Template name
	Name string `json:"name"`

	// The description for the template
	Description string `json:"description"`

	// Number of times this template has been used
	UsageCount int `json:"usage_count"`

	// The ID of the user who created the template
	CreatorID string `json:"creator_id"`

	// The user who created the template
	Creator *User `json:"creator"`

	// When this template was created
	CreatedAt time.Time `json:"created_at"`

	// When this template was last synced to the source guild
	UpdatedAt time.Time `json:"updated_at"`

	// The ID of the guild this template is based on
	SourceGuildID string `json:"source_guild_id"`

	// The guild snapshot this template contains
	SerializedSourceGuild *Guild `json:"serialized_source_guild"`

	// Whether the template has unsynced changes
	IsDirty bool `json:"is_dirty"`
}
