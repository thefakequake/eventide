package eventide

import (
	"encoding/json"
	"fmt"

	"github.com/thefakequake/eventide/discord"
)

var eventCodec = NewEventCodec(
	new(discord.ReadyEvent),
	new(discord.ResumedEvent),
	new(discord.ApplicationCommandPermissionsUpdateEvent),
	new(discord.ChannelCreateEvent),
	new(discord.ChannelUpdateEvent),
	new(discord.ChannelDeleteEvent),
	new(discord.ThreadCreateEvent),
	new(discord.ThreadUpdateEvent),
	new(discord.ThreadDeleteEvent),
	new(discord.ThreadListSyncEvent),
	new(discord.ThreadMemberUpdateEvent),
	new(discord.ThreadMembersUpdateEvent),
	new(discord.ChannelPinsUpdateEvent),
	new(discord.GuildCreateEvent),
	new(discord.GuildUpdateEvent),
	new(discord.GuildDeleteEvent),
	new(discord.GuildBanAddEvent),
	new(discord.GuildBanRemoveEvent),
	new(discord.GuildEmojisUpdateEvent),
	new(discord.GuildStickersUpdateEvent),
	new(discord.GuildIntegrationsUpdateEvent),
	new(discord.GuildMemberAddEvent),
	new(discord.GuildMemberRemoveEvent),
	new(discord.GuildMemberUpdateEvent),
	new(discord.GuildMembersChunkEvent),
	new(discord.GuildRoleCreateEvent),
	new(discord.GuildRoleUpdateEvent),
	new(discord.GuildRoleDeleteEvent),
	new(discord.GuildScheduledEventCreateEvent),
	new(discord.GuildScheduledEventUpdateEvent),
	new(discord.GuildScheduledEventDeleteEvent),
	new(discord.GuildScheduledEventUserAddEvent),
	new(discord.GuildScheduledEventUserRemoveEvent),
	new(discord.IntegrationCreateEvent),
	new(discord.IntegrationUpdateEvent),
	new(discord.IntegrationDeleteEvent),
	new(discord.InviteCreateEvent),
	new(discord.InviteDeleteEvent),
	new(discord.MessageCreateEvent),
	new(discord.MessageUpdateEvent),
	new(discord.MessageDeleteEvent),
	new(discord.MessageDeleteBulkEvent),
	new(discord.MessageReactionAddEvent),
	new(discord.MessageReactionRemoveEvent),
	new(discord.MessageReactionRemoveAllEvent),
	new(discord.MessageReactionRemoveEmojiEvent),
	new(discord.PresenceUpdateEvent),
	new(discord.TypingStartEvent),
	new(discord.UserUpdateEvent),
	new(discord.VoiceStateUpdateEvent),
	new(discord.VoiceServerUpdateEvent),
	new(discord.WebhooksUpdateEvent),
	new(discord.InteractionCreateEvent),
	new(discord.StageInstanceCreateEvent),
	new(discord.StageInstanceUpdateEvent),
	new(discord.StageInstanceDeleteEvent),
)

type EventCodec struct {
	events map[string]discord.Event
}

func NewEventCodec(events ...discord.Event) *EventCodec {
	c := &EventCodec{events: make(map[string]discord.Event)}
	for _, e := range events {
		c.events[e.EventType()] = e
	}
	return c
}

func (c *EventCodec) DecodeEvent(op discord.GatewayPayload[json.RawMessage]) (discord.Event, error) {
	e, ok := c.events[op.Type]
	fmt.Printf("%s called: %v\n", op.Type, e)

	if !ok {
		return nil, fmt.Errorf("unknown event: %s", op.Type)
	}
	err := json.Unmarshal(op.Data, &e)

	return e, err
}
