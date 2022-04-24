package eventide

type GatewayURL struct {
	URL string `json:"url"`
}

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	System        bool   `json:"system"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Banner        string `json:"banner"`
	AccentColor   int    `json:"accent_color"`
	Locale        string `json:"locale"`
	Verified      bool   `json:"verified"`
	Email         string `json:"email"`
	Flags         int    `json:"flags"`
	PremiumType   int    `json:"premium_type"`
	PublicFlags   int    `json:"public_flags"`
}



type Message struct {
	ID                string   `json:"id"`
	ChannelID         string   `json:"channel_id"`
	GuildId           string   `json:"guild_id,omitempty"`
	Author            *User    `json:"author"`
	Member            any      `json:"member"`
	Content           string   `json:"content"`
	Timestamp         string   `json:"timestamp"`
	EditedTimestamp   string   `json:"edited_timestamp"`
	TTS               bool     `json:"tts"`
	MentionEveryone   bool     `json:"mention_everyone"`
	Mentions          []*User  `json:"mentions"`
	MentionRoles      any      `json:"mention_roles"`
	MentionChannels   any      `json:"mention_channels"`
	Attachments       any      `json:"attachments"`
	Embeds            any      `json:"embeds"`
	Reactions         any      `json:"reactions"`
	Pinned            bool     `json:"pinned"`
	WebhookID         string   `json:"webhook_id"`
	Type              int      `json:"type"`
	Activity          any      `json:"activity"`
	Application       any      `json:"application"`
	ApplicationID     string   `json:"application_id"`
	MessageReference  any      `json:"message_reference"`
	Flags             int      `json:"flags"`
	ReferencedMessage *Message `json:"referenced_message"`
	Interaction       any      `json:"interaction"`
	Thread            any      `json:"thread,omitempty"`
	Components        any      `json:"components"`
	StickerItems      any      `json:"sticker_items"`
	Stickers          any      `json:"stickers"`
}

type MessageReference struct {
	MessageID string `json:"message_id"`
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id,omitempty"`
}

type MessageSend struct {
	Content string   `json:"content,omitempty"`
	Embeds  []*Embed `json:"embeds"`
	TTS     bool     `json:"tts"`
}

type Embed struct {
	Title       string          `json:"title,omitempty"`
	Type        string          `json:"type,omitempty"`
	Description string          `json:"description,omitempty"`
	URL         string          `json:"url,omitempty"`
	Timestamp   string          `json:"timestamp,omitempty"`
	Color       int             `json:"color,omitempty"`
	Footer      *EmbedFooter    `json:"footer,omitempty"`
	Image       *EmbedImage     `json:"image,omitempty"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail,omitempty"`
	Video       *EmbedVideo     `json:"video,omitempty"`
	Provider    *EmbedProvider  `json:"provider,omitempty"`
	Author      *EmbedAuthor    `json:"author,omitempty"`
	Fields      []*EmbedField   `json:"fields,omitempty"`
}

type EmbedFooter struct {
	Text         string `json:"text,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedImage struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type EmbedThumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type EmbedVideo struct {
	URL    string `json:"url,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type EmbedProvider struct {
	URL  string `json:"url,omitempty"`
	Name string `json:"name,omitempty"`
}

type EmbedAuthor struct {
	URL          string `json:"url,omitempty"`
	Name         string `json:"name"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type EmbedField struct {
	Name   string `json:"name,omitempty"`
	Value  string `json:"value,omitempty"`
	Inline bool   `json:"inline,omitempty"`
}
