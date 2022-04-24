package eventide

var (
	ApiVer          = "v9"
	EndpointAPI     = "https://discord.com/api/" + ApiVer + "/"
	EndpointGateway = EndpointAPI + "gateway"

	EndpointChannels        = EndpointAPI + "channels"
	EndpointChannel         = func(cID string) string { return EndpointChannels + "/" + cID }
	EndpointChannelMessages = func(cID string) string { return EndpointChannel(cID) + "/" + "messages" }
	EndpointChannelMessage  = func(cID, mID string) string { return EndpointChannelMessages(cID) + "/" + mID }
)
