package eventide

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
	Id                string   `json:"id"`
	ChannelId         string   `json:"channel_id"`
	GuildId           string   `json:"guild_id"`
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
	ReferencedMessage *Message `json:"message"`
	Interaction       any      `json:"interaction"`
	Thread            any      `json:"thread"`
	Components        any      `json:"components"`
	StickerItems      any      `json:"sticker_items"`
	Stickers          any      `json:"stickers"`
}
