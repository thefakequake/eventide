package discord

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-structure
type ApplicationCommand struct {
	// Unique ID of command
	ID string `json:"id"`

	// Type of command, defaults to 1
	Type ApplicationCommandType `json:"type,omitempty"`

	// ID of the parent application
	ApplicationID string `json:"application_id"`

	// Guild ID of the command, if not global
	GuildID string `json:"guild_id,omitempty"`

	// Name of command, 1-32 characters
	Name string `json:"name"`

	// Localization dictionary for name field. Values follow the same restrictions as name
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`

	// Description for CHAT_INPUT commands, 1-100 characters. Empty string for USER and MESSAGE commands
	Description string `json:"description"`

	// Localization dictionary for description field. Values follow the same restrictions as description
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`

	// Parameters for the command, max of 25
	Options []*ApplicationCommandOption `json:"options,omitempty"`

	// Set of permissions represented as a bit set
	DefaultMemberPermissions string `json:"default_member_permissions,omitempty"`

	// Indicates whether the command is available in DMs with the app, only for globally-scoped commands. By default, commands are visible.
	DmPermission bool `json:"dm_permission,omitempty"`

	// Not recommended for use as field will soon be deprecated. Indicates whether the command is enabled by default when the app is added to a guild, defaults to true
	DefaultPermission bool `json:"default_permission,omitempty"`

	// Autoincrementing version identifier updated during substantial record changes
	Version string `json:"version"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-types
type ApplicationCommandType int

const (
	ApplicationCommandTypeChatInput ApplicationCommandType = iota + 1
	ApplicationCommandTypeUser
	ApplicationCommandTypeMessage
)

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-structure
type ApplicationCommandOption struct {
	// Type of option
	Type ApplicationCommandOptionType `json:"type"`

	// 1-32 character name
	Name string `json:"name"`

	// Localization dictionary for the name field. Values follow the same restrictions as name
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`

	// 1-100 character description
	Description string `json:"description"`

	// Localization dictionary for the description field. Values follow the same restrictions as description
	DescriptionLocalizations map[string]string `json:"description_localizations,omitempty"`

	// If the parameter is required or optional--default false
	Required bool `json:"required,omitempty"`

	// Choices for STRING, INTEGER, and NUMBER types for the user to pick from, max 25
	Choices []*ApplicationCommandOptionChoice `json:"choices,omitempty"`

	// If the option is a subcommand or subcommand group type, these nested options will be the parameters
	Options []*ApplicationCommandOption `json:"options,omitempty"`

	// If the option is a channel type, the channels shown will be restricted to these types
	ChannelTypes []ChannelType `json:"channel_types,omitempty"`

	// If the option is an INTEGER or NUMBER type, the minimum value permitted
	MinValue float64 `json:"min_value,omitempty"`

	// If the option is an INTEGER or NUMBER type, the maximum value permitted
	MaxValue float64 `json:"max_value,omitempty"`

	// If autocomplete interactions are enabled for this STRING, INTEGER, or NUMBER type option
	Autocomplete bool `json:"autocomplete"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-type
type ApplicationCommandOptionType int

const (
	ApplicationCommandOptionTypeSubCommand ApplicationCommandOptionType = iota + 1
	ApplicationCommandOptionTypeSubCommandGroup
	ApplicationCommandOptionTypeString
	ApplicationCommandOptionTypeInteger
	ApplicationCommandOptionTypeBoolean
	ApplicationCommandOptionTypeUser
	ApplicationCommandOptionTypeChannel
	ApplicationCommandOptionTypeRole
	ApplicationCommandOptionTypeMentionable
	ApplicationCommandOptionTypeNumber
	ApplicationCommandOptionTypeAttachment
)

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-option-choice-structure
type ApplicationCommandOptionChoice struct {
	// 1-100 character choice name
	Name string `json:"name"`

	// Localization dictionary for the name field. Values follow the same restrictions as name
	NameLocalizations map[string]string `json:"name_localizations,omitempty"`

	// Value for the choice, can be string, integer or float and up to 100 characters if string
	Value interface{} `json:"value"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-object-application-command-interaction-data-option-structure
type ApplicationCommandInteractionDataOption struct {
	// Name of the parameter
	Name string `json:"name"`

	// Value of application command option type
	Type int `json:"type"`

	// Value of the option resulting from user input, can be string, integer or float
	Value interface{} `json:"value,omitempty"`

	// Present if this option is a group or subcommand
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`

	// True if this option is the currently focused option for autocomplete
	Focused bool `json:"focused,omitempty"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-guild-application-command-permissions-structure
type GuildApplicationCommandPermissions struct {
	// ID of the command
	ID string `json:"id"`

	// ID of the application the command belongs to
	ApplicationID string `json:"application_id"`

	// ID of the guild
	GuildID string `json:"guild_id"`

	// Permissions for the command in the guild, max of 100
	Permissions []*ApplicationCommandPermissions `json:"permissions"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permissions-structure
type ApplicationCommandPermissions struct {
	// ID of the role, user, or channel. It can also be a permission constant
	ID string `json:"id"`

	// Role (`1`), user (`2`), or channel (`3`)
	Type *ApplicationCommandPermissionType `json:"type"`

	// True to allow, false`, to disallow
	Permission bool `json:"permission"`
}

// https://discord.com/developers/docs/interactions/application-commands#application-command-permissions-object-application-command-permission-type
type ApplicationCommandPermissionType int

const (
	ApplicationCommandPermissionTypeRole ApplicationCommandPermissionType = iota + 1
	ApplicationCommandPermissionTypeUser
	ApplicationCommandPermissionTypeChannel
)
