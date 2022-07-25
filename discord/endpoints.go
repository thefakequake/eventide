package discord

var (
	APIVersion      = "10"
	EndpointAPI     = "https://discord.com/api/v" + APIVersion
	EndpointGateway = EndpointAPI + "/gateway"

	EndpointGuilds = EndpointAPI + "/guilds"
	EndpointGuild  = func(gID string) string { return EndpointGuilds + "/" + gID }

	EndpointGuildAuditLog = func(gID string) string { return EndpointGuild(gID) + "/audit-logs" }

	EndpointChannels               = EndpointAPI + "/channels"
	EndpointChannel                = func(cID string) string { return EndpointChannels + "/" + cID }
	EndpointChannelMessages        = func(cID string) string { return EndpointChannel(cID) + "/" + "messages" }
	EndpointChannelMessage         = func(cID, mID string) string { return EndpointChannelMessages(cID) + "/" + mID }
	EndpointCrosspostMessage       = func(cID, mID string) string { return EndpointChannelMessage(cID, mID) + "/crosspost" }
	EndpointReactions              = func(cID, mID string) string { return EndpointChannelMessage(cID, mID) + "/reactions" }
	EndpointReactionsEmoji         = func(cID, mID, e string) string { return EndpointReactions(cID, mID) + "/" + e }
	EndpointOwnReaction            = func(cID, mID, e string) string { return EndpointReactionsEmoji(cID, mID, e) + "/@me" }
	EndpointUserReaction           = func(cID, mID, e, uID string) string { return EndpointReactionsEmoji(cID, mID, e) + "/" + uID }
	EndpointBulkDeleteMessages     = func(cID string) string { return EndpointChannelMessages(cID) + "/bulk-delete" }
	EndpointChannelPermission      = func(cID, oID string) string { return EndpointChannel(cID) + "/permissions/" + oID }
	EndpointChannelInvites         = func(cID string) string { return EndpointChannel(cID) + "/invites" }
	EndpointFollowNewsChannel      = func(cID string) string { return EndpointChannel(cID) + "/followers" }
	EndpointTyping                 = func(cID string) string { return EndpointChannel(cID) + "/typing" }
	EndpointPinnedMessages         = func(cID string) string { return EndpointChannel(cID) + "/pins" }
	EndpointPinnedMessage          = func(cID string, mID string) string { return EndpointPinnedMessages(cID) + "/" + mID }
	EndpointMessageThreads         = func(cID, mID string) string { return EndpointChannelMessage(cID, mID) + "/threads" }
	EndpointChannelThreads         = func(cID string) string { return EndpointChannel(cID) + "/threads" }
	EndpointThreadMembers          = func(cID string) string { return EndpointChannel(cID) + "/thread-members" }
	EndpointThreadMemberSelf       = func(cID string) string { return EndpointThreadMembers(cID) + "/@me" }
	EndpointThreadMember           = func(cID, uID string) string { return EndpointThreadMembers(cID) + "/" + uID }
	EndpointArchivedThreads        = func(cID string) string { return EndpointChannelThreads(cID) + "/archived" }
	EndpointArchivedThreadsPrivate = func(cID string) string { return EndpointArchivedThreads(cID) + "/private" }
	EndpointArchivedThreadsPublic  = func(cID string) string { return EndpointArchivedThreads(cID) + "/public" }
	EndpointJoinedArchivedThreads  = func(cID string) string { return EndpointChannel(cID) + "/users/@me/threads/archived/private" }

	EndpointGuildEmojis = func(gID string) string { return EndpointGuild(gID) + "/emojis" }
	EndpointGuildEmoji  = func(gID, eID string) string { return EndpointGuildEmojis(gID) + "/" + eID }
)
