package eventide

// https://discord.com/developers/docs/resources/user#user-object-user-structure
type User struct {
	// The user's ID
	ID string `json:"id"`

	// The user's username, not unique across the platform
	Username string `json:"username"`

	// The user's 4-digit discord-tag
	Discriminator string `json:"discriminator"`

	// The user's avatar hash
	Avatar string `json:"avatar"`

	// Whether the user belongs to an OAuth2 application
	Bot bool `json:"bot,omitempty"`

	// Whether the user is an Official Discord System user (part of the urgent message system)
	System bool `json:"system,omitempty"`

	// Whether the user has two factor enabled on their account
	MFAEnabled bool `json:"mfa_enabled,omitempty"`

	// The user's banner hash
	Banner string `json:"banner,omitempty"`

	// The user's banner color encoded as an integer representation of hexadecimal color code
	AccentColor int `json:"accent_color,omitempty"`

	// The user's chosen language option
	Locale string `json:"locale,omitempty"`

	// Whether the email on this account has been verified
	Verified bool `json:"verified,omitempty"`

	// The user's email
	Email string `json:"email,omitempty"`

	// The flags on a user's account
	Flags int `json:"flags,omitempty"`

	// The type of Nitro subscription on a user's account
	PremiumType int `json:"premium_type,omitempty"`

	// The public flags on a user's account
	PublicFlags int `json:"public_flags,omitempty"`
}

// https://discord.com/developers/docs/resources/user#connection-object-connection-structure
type Connection struct {
	// ID of the connection account
	ID string `json:"id"`

	// The username of the connection account
	Name string `json:"name"`

	// The service of the connection (twitch, youtube)
	Type string `json:"type"`

	// Whether the connection is revoked
	Revoked bool `json:"revoked,omitempty"`

	// An array of partial server integrations
	Integrations *Integration `json:"integrations,omitempty"`

	// Whether the connection is verified
	Verified bool `json:"verified"`

	// Whether friend sync is enabled for this connection
	FriendSync bool `json:"friend_sync"`

	// Whether activities related to this connection will be shown in presence updates
	ShowActivity bool `json:"show_activity"`

	// Visibility of this connection
	Visibility int `json:"visibility"`
}
