package discord

// https://discord.com/developers/docs/interactions/message-components#component-object-component-types
type ComponentType int

const (
	ComponentTypeActionRow = iota + 1
	ComponentTypeButton
	ComponentTypeSelectMenu
	ComponentTypeTextInput
)

// https://discord.com/developers/docs/interactions/message-components#button-object-button-structure
type Button struct {
	// 2 for a button
	Type ComponentType `json:"type"`

	// One of button styles
	Style int `json:"style"`

	// Text that appears on the button, max 80 characters
	Label string `json:"label,omitempty"`

	// Name`, id`, and animated
	Emoji *Emoji `json:"emoji,omitempty"`

	// A developer-defined identifier for the button, max 100 characters
	CustomID string `json:"custom_id,omitempty"`

	// A URL for link-style buttons
	URL string `json:"url,omitempty"`

	// Whether the button is disabled (default false`)
	Disabled bool `json:"disabled,omitempty"`
}

// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-menu-structure
type SelectMenu struct {
	// 3 for a select menu
	Type ComponentType `json:"type"`

	// A developer-defined identifier for the select menu, max 100 characters
	CustomID string `json:"custom_id"`

	// The choices in the select, max 25
	Options []*SelectOption `json:"options"`

	// Custom placeholder text if nothing is selected, max 150 characters
	Placeholder string `json:"placeholder,omitempty"`

	// The minimum number of items that must be chosen; default 1, min 0, max 25
	MinValues int `json:"min_values,omitempty"`

	// The maximum number of items that can be chosen; default 1, max 25
	MaxValues int `json:"max_values,omitempty"`

	// Disable the select, default false
	Disabled bool `json:"disabled,omitempty"`
}

// https://discord.com/developers/docs/interactions/message-components#select-menu-object-select-option-structure
type SelectOption struct {
	// The user-facing name of the option, max 100 characters
	Label string `json:"label"`

	// The dev-defined value of the option, max 100 characters
	Value string `json:"value"`

	// An additional description of the option, max 100 characters
	Description string `json:"description,omitempty"`

	// Id`, name`, and animated
	Emoji *Emoji `json:"emoji,omitempty"`

	// Will render this option as selected by default
	Default bool `json:"default,omitempty"`
}

// https://discord.com/developers/docs/interactions/message-components#text-inputs-text-input-structure
type TextInput struct {
	// 4 for a text input
	Type ComponentType `json:"type"`

	// A developer-defined identifier for the input, max 100 characters
	CustomID string `json:"custom_id"`

	// The Text Input Style
	Style int `json:"style"`

	// The label for this component, max 45 characters
	Label string `json:"label"`

	// The minimum input length for a text input, min 0, max 4000
	MinLength int `json:"min_length,omitempty"`

	// The maximum input length for a text input, min 1, max 4000
	MaxLength int `json:"max_length,omitempty"`

	// Whether this component is required to be filled, default true
	Required bool `json:"required,omitempty"`

	// A pre-filled value for this component, max 4000 characters
	Value string `json:"value,omitempty"`

	// Custom placeholder text if the input is empty, max 100 characters
	Placeholder string `json:"placeholder,omitempty"`
}
