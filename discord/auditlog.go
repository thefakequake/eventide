package discord

// https://discord.com/developers/docs/resources/audit-log#audit-log-object-audit-log-structure
type AuditLog struct {
	// List of audit log entries, sorted from most to least recent
	AuditLogEntries []*AuditLogEntry `json:"audit_log_entries"`

	// List of guild scheduled events found in the audit log
	GuildScheduledEvents []*GuildScheduledEvent `json:"guild_scheduled_events"`

	// List of partial integration objects
	Integrations []*Integration `json:"integrations"`

	// List of threads found in the audit log
	Threads []*Channel `json:"threads"`

	// List of users found in the audit log
	Users []*User `json:"users"`

	// List of webhooks found in the audit log
	Webhooks []*Webhook `json:"webhooks"`
}

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-entry-structure
type AuditLogEntry struct {
	// ID of the affected entity (webhook, user, role, etc.)
	TargetID string `json:"target_id"`

	// Changes made to the target_id
	Changes []*AuditLogChange `json:"changes,omitempty"`

	// User or app that made the changes
	UserID string `json:"user_id"`

	// ID of the entry
	ID string `json:"id"`

	// Type of action that occurred
	ActionType AuditLogEvent `json:"action_type"`

	// Additional info for certain event types
	Options *OptionalAuditEntryInfo `json:"options,omitempty"`

	// Reason for the change (1-512 characters)
	Reason string `json:"reason,omitempty"`
}

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-audit-log-events
type AuditLogEvent int

// TODO: finish this
const (
	AuditLogEventGuildUpdate                        AuditLogEvent = 1
	AuditLogEventChannelCreate                      AuditLogEvent = 10
	AuditLogEventChannelUpdate                      AuditLogEvent = 11
	AuditLogEventChannelDelete                      AuditLogEvent = 12
	AuditLogEventOverwriteCreate                    AuditLogEvent = 13
	AuditLogEventChannelOverwriteCreate             AuditLogEvent = 13
	AuditLogEventChannelOverwriteUpdate             AuditLogEvent = 14
	AuditLogEventChannelOverwriteDelete             AuditLogEvent = 15
	AuditLogEventMemberKick                         AuditLogEvent = 20
	AuditLogEventMemberPrune                        AuditLogEvent = 21
	AuditLogEventMemberBanAdd                       AuditLogEvent = 22
	AuditLogEventMemberUpdate                       AuditLogEvent = 24
	AuditLogEventMemberRoleUpdate                   AuditLogEvent = 25
	AuditLogEventMemberMove                         AuditLogEvent = 26
	AuditLogEventMemberDisconnect                   AuditLogEvent = 27
	AuditLogEventBotAdd                             AuditLogEvent = 28
	AuditLogEventRoleCreate                         AuditLogEvent = 30
	AuditLogEventRoleUpdate                         AuditLogEvent = 31
	AuditLogEventRoleDelete                         AuditLogEvent = 32
	AuditLogEventInviteCreate                       AuditLogEvent = 40
	AuditLogEventInviteUpdate                       AuditLogEvent = 41
	AuditLogEventInviteDelete                       AuditLogEvent = 42
	AuditLogEventWebhookCreate                      AuditLogEvent = 50
	AuditLogEventWebhookUpdate                      AuditLogEvent = 51
	AuditLogEventWebhookDelete                      AuditLogEvent = 52
	AuditLogEventEmojiCreate                        AuditLogEvent = 60
	AuditLogEventEmojiUpdate                        AuditLogEvent = 61
	AuditLogEventEmojiDelete                        AuditLogEvent = 62
	AuditLogEventMessageDelete                      AuditLogEvent = 72
	AuditLogEventMessageBulkDelete                  AuditLogEvent = 73
	AuditLogEventMessagePin                         AuditLogEvent = 74
	AuditLogEventMessageUnpin                       AuditLogEvent = 75
	AuditLogEventIntegrationCreate                  AuditLogEvent = 80
	AuditLogEventIntegrationUpdate                  AuditLogEvent = 81
	AuditLogEventIntegrationDelete                  AuditLogEvent = 82
	AuditLogEventStageInstanceCreate                AuditLogEvent = 83
	AuditLogEventStageInstanceUpdate                AuditLogEvent = 84
	AuditLogEventStageInstanceDelete                AuditLogEvent = 85
	AuditLogEventStickerCreate                      AuditLogEvent = 90
	AuditLogEventStickerUpdate                      AuditLogEvent = 91
	AuditLogEventStickerDelete                      AuditLogEvent = 92
	AuditLogEventGuildScheduledEventCreate          AuditLogEvent = 100
	AuditLogEventGuildScheduledEventUpdate          AuditLogEvent = 101
	AuditLogEventGuildScheduledEventDelete          AuditLogEvent = 102
	AuditLogEventThreadCreate                       AuditLogEvent = 110
	AuditLogEventThreadUpdate                       AuditLogEvent = 111
	AuditLogEventThreadDelete                       AuditLogEvent = 112
	AuditLogEventApplicationCommandPermissionUpdate AuditLogEvent = 121
)

// https://discord.com/developers/docs/resources/audit-log#audit-log-entry-object-optional-audit-entry-info
type OptionalAuditEntryInfo struct {
	// ID of the app whose permissions were targeted
	ApplicationID string `json:"application_id"`

	// Channel in which the entities were targeted
	ChannelID string `json:"channel_id"`

	// Number of entities that were targeted
	Count string `json:"count"`

	// Number of days after which inactive members were kicked
	DeleteMemberDays string `json:"delete_member_days"`

	// ID of the overwritten entity
	ID string `json:"id"`

	// Number of members removed by the prune
	MembersRemoved string `json:"members_removed"`

	// ID of the message that was targeted
	MessageID string `json:"message_id"`

	// Name of the role if type is "0" (not present if type is "1"`)
	RoleName string `json:"role_name"`

	// Type of overwritten entity - role (`"0"`) or member (`"1"`)
	Type string `json:"type"`
}

// https://discord.com/developers/docs/resources/audit-log#audit-log-change-object-audit-log-change-structure
type AuditLogChange struct {
	// New value of the key
	NewValue interface{} `json:"new_value,omitempty"`

	// Old value of the key
	OldValue interface{} `json:"old_value,omitempty"`

	// Name of the changed entity, with a few exceptions
	Key string `json:"key"`
}
