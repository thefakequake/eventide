package discord

// https://discord.com/developers/docs/topics/rate-limits#exceeding-a-rate-limit-rate-limit-response-structure
type RateLimitResponse struct {
	// A message saying you are being rate limited.
	Message string `json:"message"`

	// The number of seconds to wait before submitting another request.
	RetryAfter float64 `json:"retry_after"`

	// A value indicating if you are being globally rate limited or not
	Global bool `json:"global"`
}
