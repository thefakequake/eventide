package discord

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-structure
type Interaction struct {
	// ID of the interaction
	ID string `json:"id"`

	// ID of the application this interaction is for
	ApplicationID string `json:"application_id"`

	// The type of interaction
	Type InteractionType `json:"type"`

	// The command data payload
	Data *InteractionData `json:"data"`

	// The guild it was sent from
	GuildID string `json:"guild_id,omitempty"`

	// The channel it was sent from
	ChannelID string `json:"channel_id,omitempty"`

	// Guild member data for the invoking user, including permissions
	Member *GuildMember `json:"member"`

	// User object for the invoking user, if invoked in a DM
	User *User `json:"user,omitempty"`

	// A continuation token for responding to the interaction
	Token string `json:"token"`

	// Read-only property, always 1
	Version int `json:"version"`

	// For components, the message they were attached to
	Message *Message `json:"message,omitempty"`

	// The selected language of the invoking user
	Locale string `json:"locale"`

	// The guild's preferred locale, if invoked in a guild
	GuildLocale string `json:"guild_locale,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-type
type InteractionType int

const (
	InteractionTypePing InteractionType = iota + 1
	InteractionTypeApplicationCommmand
	InteractionTypeMessageComponent
	InteractionTypeApplicationCommandAutocomplete
	InteractionTypeModalSubmit
)

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-interaction-data-structure
type InteractionData struct {
	// The ID of the invoked command
	ID string `json:"id"`

	// The name of the invoked command
	Name string `json:"name"`

	// The type of the invoked command
	Type ApplicationCommandType `json:"type"`

	// Converted users + roles + channels + attachments
	Resolved *ResolvedData `json:"resolved,omitempty"`

	// The params + values from the user
	Options []*ApplicationCommandInteractionDataOption `json:"options,omitempty"`

	// The ID of the guild the command is registered to
	GuildID string `json:"guild_id,omitempty"`

	// The custom_id of the component
	CustomID string `json:"custom_id,omitempty"`

	// The type of the component
	ComponentType ComponentType `json:"component_type,omitempty"`

	// The values the user selected
	Values []*SelectOption `json:"values,omitempty"`

	// ID of the user or message targeted by a user or message command
	TargetID string `json:"target_id,omitempty"`

	// The values submitted by the user
	// Components []*MessageComponent `json:"components,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-object-resolved-data-structure
type ResolvedData struct {
	// The IDs and User objects
	Users map[string]*User `json:"users,omitempty"`

	// The IDs and partial Member objects
	Members map[string]*GuildMember `json:"members"`

	// The IDs and Role objects
	Roles map[string]*Role `json:"roles,omitempty"`

	// The IDs and partial Channel objects
	Channels map[string]*Channel `json:"channels"`

	// The IDs and partial Message objects
	Messages map[string]*Message `json:"messages,omitempty"`

	// The IDs and attachment objects
	Attachments map[string]*Attachment `json:"attachments,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#message-interaction-object-message-interaction-structure
type MessageInteraction struct {
	// ID of the interaction
	ID string `json:"id"`

	// The type of interaction
	Type *InteractionType `json:"type"`

	// The name of the application command
	Name string `json:"name"`

	// The user who invoked the interaction
	User *User `json:"user"`

	// The member who invoked the interaction in the guild
	Member *GuildMember `json:"member,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-response-structure
type InteractionResponse struct {
	// The type of response
	Type InteractionCallbackType `json:"type"`

	// An optional response message
	Data *InteractionCallbackData `json:"data,omitempty"`
}

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-type
type InteractionCallbackType int

const (
	InteractionCallbackTypePong InteractionCallbackType = iota + 1
	_
	_
	InteractionCallbackTypeChannelMessageWithSource
	InteractionCallbackTypeDeferredChannelMessageWithSource
	InteractionCallbackTypeDeferredUpdateMessage
	InteractionCallbackTypeUpdateMessage
	InteractionCallbackTypeApplicationCommandAutocompleteResult
	InteractionCallbackTypeModal
)

// https://discord.com/developers/docs/interactions/receiving-and-responding#interaction-response-object-interaction-callback-data-structure
type InteractionCallbackData struct {
	// Is the response TTS
	TTS bool `json:"tts,omitempty"`

	// Message content
	Content string `json:"content,omitempty"`

	// Supports up to 10 embeds
	Embeds []*Embed `json:"embeds,omitempty"`

	// Allowed mentions object
	AllowedMentions *AllowedMentions `json:"allowed_mentions,omitempty"`

	// Message flags combined as a bitfield (only SUPPRESS_EMBEDS and EPHEMERAL can be set)
	Flags int `json:"flags,omitempty"`

	// Message components
	// Components []*Component `json:"components,omitempty"`

	// Attachment objects with filename and description
	Attachments []*Attachment `json:"attachments"`
}
