package eventide

// https://discord.com/developers/docs/topics/certified-devices#models-device-object
type Device struct {
	// The type of device
	Type *DeviceType `json:"type"`

	// The device's Windows UUID
	ID string `json:"id"`

	// The hardware vendor
	Vendor *Vendor `json:"vendor"`

	// The model of the product
	Model *Model `json:"model"`

	// UUIDs of related devices
	Related []string `json:"related"`

	// If the device's native echo cancellation is enabled
	EchoCancellation bool `json:"echo_cancellation"`

	// If the device's native noise suppression is enabled
	NoiseSuppression bool `json:"noise_suppression"`

	// If the device's native automatic gain control is enabled
	AutomaticGainControl bool `json:"automatic_gain_control"`

	// If the device is hardware muted
	HardwareMute bool `json:"hardware_mute"`
}

// https://discord.com/developers/docs/topics/certified-devices#models-vendor-object
type Vendor struct {
	// Name of the vendor
	Name string `json:"name"`

	// URL for the vendor
	URL string `json:"url"`
}

// https://discord.com/developers/docs/topics/certified-devices#models-model-object
type Model struct {
	// Name of the model
	Name string `json:"name"`

	// URL for the model
	URL string `json:"url"`
}
