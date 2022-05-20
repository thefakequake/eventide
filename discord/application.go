package eventide

// https://discord.com/developers/docs/resources/application#application-object-application-structure
type Application struct {
	// The ID of the app
	ID string `json:"id"`

	// The name of the app
	Name string `json:"name"`

	// The icon hash of the app
	Icon string `json:"icon"`

	// The description of the app
	Description string `json:"description"`

	// An array of RPC origin urls, if RPC is enabled
	RPCOrigins []string `json:"rpc_origins,omitempty"`

	// When false only app owner can join the app's bot to guilds
	BotPublic bool `json:"bot_public"`

	// When true the app's bot will only join upon completion of the full oauth2 code grant flow
	BotRequireCodeGrant bool `json:"bot_require_code_grant"`

	// The URL of the app's terms of service
	TermsOfServiceURL string `json:"terms_of_service_url,omitempty"`

	// The URL of the app's privacy policy
	PrivacyPolicyURL string `json:"privacy_policy_url,omitempty"`

	// Partial user object containing info on the owner of the application
	Owner *User `json:"owner,omitempty"`

	// The hex encoded key for verification in interactions and the GameSDK's GetTicket
	VerifyKey string `json:"verify_key"`

	// If the application belongs to a team, this will be a list of the members of that team
	Team *Team `json:"team"`

	// If this application is a game sold on Discord, this field will be the guild to which it has been linked
	GuildID string `json:"guild_id,omitempty"`

	// If this application is a game sold on Discord, this field will be the ID of the "Game SKU" that is created, if exists
	PrimarySkuID string `json:"primary_sku_id,omitempty"`

	// If this application is a game sold on Discord, this field will be the URL slug that links to the store page
	Slug string `json:"slug,omitempty"`

	// The application's default rich presence invite cover image hash
	CoverImage string `json:"cover_image,omitempty"`

	// The application's public flags
	Flags int `json:"flags,omitempty"`

	// Up to 5 tags describing the content and functionality of the application
	Tags []string `json:"tags,omitempty"`

	// Settings for the application's default in-app authorization link, if enabled
	InstallParams *InstallParams `json:"install_params,omitempty"`

	// The application's default custom authorization link, if enabled
	CustomInstallURL string `json:"custom_install_url,omitempty"`
}

// https://discord.com/developers/docs/resources/application#install-params-object-install-params-structure
type InstallParams struct {
	// The scopes to add the application to the server with
	Scopes []string `json:"scopes"`

	// The permissions to request for the bot role
	Permissions string `json:"permissions"`
}
