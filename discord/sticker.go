package eventide

// https://discord.com/developers/docs/resources/sticker#sticker-object-sticker-structure
type Sticker struct {
	// ID of the sticker
	ID string `json:"id"`

	// For standard stickers, ID of the pack the sticker is from
	PackID string `json:"pack_id,omitempty"`

	// Name of the sticker
	Name string `json:"name"`

	// Description of the sticker
	Description string `json:"description"`

	// Autocomplete/suggestion tags for the sticker (max 200 characters)
	Tags string `json:"tags"`

	// Deprecated** previously the sticker asset hash, now an empty string
	Asset string `json:"asset,omitempty"`

	// Type of sticker
	Type int `json:"type"`

	// Type of sticker format
	FormatType int `json:"format_type"`

	// Whether this guild sticker can be used, may be false due to loss of Server Boosts
	Available bool `json:"available,omitempty"`

	// ID of the guild that owns this sticker
	GuildID string `json:"guild_id,omitempty"`

	// The user that uploaded the guild sticker
	User *User `json:"user,omitempty"`

	// The standard sticker's sort order within its pack
	SortValue int `json:"sort_value,omitempty"`
}

// https://discord.com/developers/docs/resources/sticker#sticker-item-object-sticker-item-structure
type StickerItem struct {
	// ID of the sticker
	ID string `json:"id"`

	// Name of the sticker
	Name string `json:"name"`

	// Type of sticker format
	FormatType int `json:"format_type"`
}

// https://discord.com/developers/docs/resources/sticker#sticker-pack-object-sticker-pack-structure
type StickerPack struct {
	// ID of the sticker pack
	ID string `json:"id"`

	// The stickers in the pack
	Stickers []*Sticker `json:"stickers"`

	// Name of the sticker pack
	Name string `json:"name"`

	// ID of the pack's SKU
	SkuID string `json:"sku_id"`

	// ID of a sticker in the pack which is shown as the pack's icon
	CoverStickerID string `json:"cover_sticker_id,omitempty"`

	// Description of the sticker pack
	Description string `json:"description"`

	// ID of the sticker pack's banner image
	BannerAssetID string `json:"banner_asset_id,omitempty"`
}

// https://discord.com/developers/docs/resources/sticker#list-nitro-sticker-packs-%-get-/sticker-packs-response-structure
type Response struct {
	// Name of the sticker (2-30 characters)
	Name string `json:"name"`

	// Description of the sticker (empty or 2-100 characters)
	Description string `json:"description"`

	// Autocomplete/suggestion tags for the sticker (max 200 characters)
	Tags string `json:"tags"`

	// The sticker file to upload, must be a PNG, APNG, or Lottie JSON file, max 500 KB
	File *FileContents `json:"file"`
}
