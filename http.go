package eventide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/thefakequake/eventide/discord"
)

type HTTPError struct {
	Request      *http.Request
	Response     *http.Response
	ResponseBody []byte
	Message      *ErrorMessage
}

type ErrorMessage struct {
	Code int    `json:"code"`
	Text string `json:"message"`
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("http %d: %s", e.Response.StatusCode, e.ResponseBody)
}

func (c *Client) Request(method string, url string, body interface{}) ([]byte, error) {
	var err error
	var reader io.Reader

	if body != nil {
		dat, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reader = bytes.NewBuffer(dat)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusNoContent:
	default:
		e := HTTPError{
			Request:      req,
			Response:     resp,
			ResponseBody: respBody,
		}

		var mes ErrorMessage
		if err := json.Unmarshal(respBody, &mes); err == nil {
			e.Message = &mes
		}
		err = e
	}

	return respBody, err
}

func (c *Client) GetGateway() (string, error) {
	var err error

	body, err := c.Request("GET", discord.EndpointGateway, nil)
	if err != nil {
		return "", err
	}

	var data discord.GetGateway
	err = json.Unmarshal(body, &data)

	v := url.Values{}
	v.Add("v", discord.APIVersion)
	v.Add("encoding", "json")

	return data.URL + "?" + v.Encode(), err
}

// https://discord.com/developers/docs/resources/audit-log#get-guild-audit-log
func (c *Client) GetGuildAuditLog(guildID string) (*discord.AuditLog, error) {
	body, err := c.Request("GET", discord.EndpointGuildAuditLog(guildID), nil)
	if err != nil {
		return nil, err
	}

	var auditLog discord.AuditLog
	err = json.Unmarshal(body, &auditLog)

	return &auditLog, err
}

// https://discord.com/developers/docs/resources/channel#get-channel
func (c *Client) GetChannel(channelID string) (*discord.Channel, error) {
	body, err := c.Request("GET", discord.EndpointChannel(channelID), nil)
	if err != nil {
		return nil, err
	}

	var channel discord.Channel
	err = json.Unmarshal(body, &channel)

	return &channel, err
}

// https://discord.com/developers/docs/resources/channel#modify-channel
func (c *Client) ModifyChannel(channelID string, params *discord.ModifyChannel) (*discord.Channel, error) {
	body, err := c.Request("PATCH", discord.EndpointChannel(channelID), params)
	if err != nil {
		return nil, err
	}

	var channel discord.Channel
	err = json.Unmarshal(body, &channel)

	return &channel, err
}

// https://discord.com/developers/docs/resources/channel#deleteclose-channel
func (c *Client) DeleteChannel(channelID string) (*discord.Channel, error) {
	body, err := c.Request("DELETE", discord.EndpointChannel(channelID), nil)
	if err != nil {
		return nil, err
	}

	var channel discord.Channel
	err = json.Unmarshal(body, &channel)

	return &channel, err
}

// https://discord.com/developers/docs/resources/channel#get-channel-messages
func (c *Client) GetChannelMessages(channelID string, params *discord.GetChannelMessages) ([]*discord.Message, error) {
	body, err := c.Request("GET", discord.EndpointChannelMessages(channelID), params)
	if err != nil {
		return nil, err
	}

	var messages []*discord.Message
	err = json.Unmarshal(body, &messages)

	return messages, err
}

// https://discord.com/developers/docs/resources/channel#get-channel-message
func (c *Client) GetChannelMessage(channelID string, messageID string) (*discord.Message, error) {
	body, err := c.Request("GET", discord.EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(body, &message)

	return &message, err
}

// https://discord.com/developers/docs/resources/channel#create-message
func (c *Client) CreateMessage(channelID string, params *discord.CreateMessage) (*discord.Message, error) {
	body, err := c.Request("POST", discord.EndpointChannelMessages(channelID), params)
	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(body, &message)

	return &message, err
}

// https://discord.com/developers/docs/resources/channel#crosspost-message
func (c *Client) CrosspostMessage(channelID string, messageID string) (*discord.Message, error) {
	body, err := c.Request("POST", discord.EndpointCrosspostMessage(channelID, messageID), nil)
	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(body, &message)

	return &message, err
}

// https://discord.com/developers/docs/resources/channel#create-reaction
func (c *Client) CreateReaction(channelID string, messageID string, emoji string) error {
	_, err := c.Request("PUT", discord.EndpointOwnReaction(channelID, messageID, emoji), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#delete-own-reaction
func (c *Client) DeleteOwnReaction(channelID string, messageID string, emoji string) error {
	_, err := c.Request("DELETE", discord.EndpointOwnReaction(channelID, messageID, emoji), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#delete-user-reaction
func (c *Client) DeleteUserReaction(channelID string, messageID string, emoji string, userID string) error {
	_, err := c.Request("DELETE", discord.EndpointUserReaction(channelID, messageID, emoji, userID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#get-reactions
func (c *Client) GetReactions(channelID string, messageID string, emoji string, params *discord.GetReactions) ([]*discord.User, error) {
	body, err := c.Request("GET", discord.EndpointReactionsEmoji(channelID, messageID, emoji), params)
	if err != nil {
		return nil, err
	}

	var users []*discord.User
	err = json.Unmarshal(body, &users)

	return users, err
}

// https://discord.com/developers/docs/resources/channel#delete-all-reactions
func (c *Client) DeleteAllReactions(channelID string, messageID string) error {
	_, err := c.Request("DELETE", discord.EndpointReactions(channelID, messageID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#delete-all-reactions
func (c *Client) DeleteAllReactionsForEmoji(channelID string, messageID string, emoji string) error {
	_, err := c.Request("DELETE", discord.EndpointReactionsEmoji(channelID, messageID, emoji), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#edit-message
func (c *Client) EditMessage(channelID string, messageID string, params *discord.EditMessage) (*discord.Message, error) {
	body, err := c.Request("PATCH", discord.EndpointChannelMessage(channelID, messageID), params)
	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(body, &message)

	return &message, err
}

// https://discord.com/developers/docs/resources/channel#delete-message
func (c *Client) DeleteMessage(channelID string, messageID string) (*discord.Message, error) {
	body, err := c.Request("DELETE", discord.EndpointChannelMessage(channelID, messageID), nil)
	if err != nil {
		return nil, err
	}

	var message discord.Message
	err = json.Unmarshal(body, &message)

	return &message, err
}

// https://discord.com/developers/docs/resources/channel#bulk-delete-messages
func (c *Client) BulkDeleteMessages(channelID string, params *discord.BulkDeleteMessages) error {
	_, err := c.Request("POST", discord.EndpointBulkDeleteMessages(channelID), params)
	return err
}

// https://discord.com/developers/docs/resources/channel#edit-channel-permissions
func (c *Client) EditChannelPermissions(channelID string, overwriteID string, params discord.EditChannelPermissions) error {
	_, err := c.Request("PUT", discord.EndpointChannelPermission(channelID, overwriteID), params)
	return err
}

// https://discord.com/developers/docs/resources/channel#get-channel-invites
func (c *Client) GetChannelInvites(channelID string) ([]*discord.Invite, error) {
	body, err := c.Request("GET", discord.EndpointChannelInvites(channelID), nil)
	if err != nil {
		return nil, err
	}

	var invites []*discord.Invite
	err = json.Unmarshal(body, &invites)

	return invites, err
}

// https://discord.com/developers/docs/resources/channel#create-channel-invite
func (c *Client) CreateChannelInvite(channelID string, params *discord.CreateChannelInvite) (*discord.Invite, error) {
	body, err := c.Request("POST", discord.EndpointChannelInvites(channelID), params)
	if err != nil {
		return nil, err
	}

	var invite discord.Invite
	err = json.Unmarshal(body, &invite)

	return &invite, err
}

// https://discord.com/developers/docs/resources/channel#delete-channel-permission
func (c *Client) DeleteChannelPermission(channelID string, overwriteID string) error {
	_, err := c.Request("DELETE", discord.EndpointChannelPermission(channelID, overwriteID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#follow-news-channel
func (c *Client) FollowNewsChannel(channelID string, params *discord.FollowNewsChannel) (*discord.FollowedChannel, error) {
	body, err := c.Request("POST", discord.EndpointFollowNewsChannel(channelID), nil)
	if err != nil {
		return nil, err
	}

	var channel discord.FollowedChannel
	err = json.Unmarshal(body, &channel)

	return &channel, err
}

// https://discord.com/developers/docs/resources/channel#trigger-typing-indicator
func (c *Client) TriggerTypingIndicator(channelID string) error {
	_, err := c.Request("POST", discord.EndpointTyping(channelID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#get-pinned-messages
func (c *Client) GetPinnedMessages(channelID string) ([]*discord.Message, error) {
	body, err := c.Request("GET", discord.EndpointPinnedMessages(channelID), nil)
	if err != nil {
		return nil, err
	}

	var messages []*discord.Message
	err = json.Unmarshal(body, &messages)

	return messages, err
}

// https://discord.com/developers/docs/resources/channel#pin-message
func (c *Client) PinMessage(channelID string, messageID string) error {
	_, err := c.Request("PUT", discord.EndpointPinnedMessage(channelID, messageID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#unpin-message
func (c *Client) UnpinMessage(channelID string, messageID string) error {
	_, err := c.Request("DELETE", discord.EndpointPinnedMessage(channelID, messageID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#start-thread-from-message
func (c *Client) StartThreadFromMessage(channelID string, messageID string, params *discord.StartThreadFromMessage) (*discord.Channel, error) {
	body, err := c.Request("POST", discord.EndpointMessageThreads(channelID, messageID), params)
	if err != nil {
		return nil, err
	}

	var channel discord.Channel
	err = json.Unmarshal(body, &channel)

	return &channel, err
}

// https://discord.com/developers/docs/resources/channel#start-thread-without-message
func (c *Client) StartThreadWithoutMessage(channelID string, params *discord.StartThreadWithoutMessage) (*discord.Channel, error) {
	body, err := c.Request("POST", discord.EndpointChannelThreads(channelID), params)
	if err != nil {
		return nil, err
	}

	var channel discord.Channel
	err = json.Unmarshal(body, &channel)

	return &channel, err
}

// https://discord.com/developers/docs/resources/channel#start-thread-in-forum-channel
func (c *Client) StartThreadInForumChannel(channelID string, params *discord.StartThreadInForumChannel) (*discord.ForumChannelThreadCreate, error) {
	body, err := c.Request("POST", discord.EndpointChannelThreads(channelID), params)
	if err != nil {
		return nil, err
	}

	var thread discord.ForumChannelThreadCreate
	err = json.Unmarshal(body, &thread)

	return &thread, err
}

// https://discord.com/developers/docs/resources/channel#join-thread
func (c *Client) JoinThread(channelID string) error {
	_, err := c.Request("PUT", discord.EndpointThreadMemberSelf(channelID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#join-thread
func (c *Client) AddThreadMember(channelID string, userID string) error {
	_, err := c.Request("PUT", discord.EndpointThreadMember(channelID, userID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#leave-thread
func (c *Client) LeaveThread(channelID string) error {
	_, err := c.Request("DELETE", discord.EndpointThreadMemberSelf(channelID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#remove-thread-member
func (c *Client) RemoveThreadMember(channelID string, userID string) error {
	_, err := c.Request("DELETE", discord.EndpointThreadMember(channelID, userID), nil)
	return err
}

// https://discord.com/developers/docs/resources/channel#get-thread-member
func (c *Client) GetThreadMember(channelID string, userID string) (*discord.ThreadMember, error) {
	body, err := c.Request("GET", discord.EndpointThreadMember(channelID, userID), nil)
	if err != nil {
		return nil, err
	}

	var member discord.ThreadMember
	err = json.Unmarshal(body, &member)

	return &member, err
}

// https://discord.com/developers/docs/resources/channel#list-thread-members
func (c *Client) ListThreadMembers(channelID string) ([]*discord.ThreadMember, error) {
	body, err := c.Request("GET", discord.EndpointThreadMembers(channelID), nil)
	if err != nil {
		return nil, err
	}

	var members []*discord.ThreadMember
	err = json.Unmarshal(body, &members)

	return members, err
}

// https://discord.com/developers/docs/resources/channel#list-public-archived-threads
func (c *Client) ListPublicArchivedThreads(channelID string, params *discord.ListArchivedThreads) (*discord.ArchivedThreads, error) {
	body, err := c.Request("GET", discord.EndpointArchivedThreadsPublic(channelID), params)
	if err != nil {
		return nil, err
	}

	var threads *discord.ArchivedThreads
	err = json.Unmarshal(body, &threads)

	return threads, err
}

// https://discord.com/developers/docs/resources/channel#list-private-archived-threads
func (c *Client) ListPrivateArchivedThreads(channelID string, params *discord.ListArchivedThreads) (*discord.ArchivedThreads, error) {
	body, err := c.Request("GET", discord.EndpointArchivedThreadsPrivate(channelID), params)
	if err != nil {
		return nil, err
	}

	var threads *discord.ArchivedThreads
	err = json.Unmarshal(body, &threads)

	return threads, err
}

// https://discord.com/developers/docs/resources/channel#list-joined-private-archived-threads
func (c *Client) ListJoinedPrivateArchivedThreads(channelID string, params *discord.ListArchivedThreads) (*discord.ArchivedThreads, error) {
	body, err := c.Request("GET", discord.EndpointJoinedArchivedThreads(channelID), params)
	if err != nil {
		return nil, err
	}

	var threads *discord.ArchivedThreads
	err = json.Unmarshal(body, &threads)

	return threads, err
}

// https://discord.com/developers/docs/resources/emoji#list-guild-emojis
func (c *Client) ListGuildEmojis(guildID string) ([]*discord.Emoji, error) {
	body, err := c.Request("GET", discord.EndpointGuildEmojis(guildID), nil)
	if err != nil {
		return nil, err
	}

	var emojis []*discord.Emoji
	err = json.Unmarshal(body, &emojis)

	return emojis, err
}

// https://discord.com/developers/docs/resources/emoji#get-guild-emoji
func (c *Client) GetGuildEmoji(guildID string, emojiID string) (*discord.Emoji, error) {
	body, err := c.Request("GET", discord.EndpointGuildEmoji(guildID, emojiID), nil)
	if err != nil {
		return nil, err
	}

	var emoji discord.Emoji
	err = json.Unmarshal(body, &emoji)

	return &emoji, err
}

// https://discord.com/developers/docs/resources/emoji#create-guild-emoji
func (c *Client) CreateGuildEmoji(guildID string, params *discord.CreateGuildEmoji) (*discord.Emoji, error) {
	body, err := c.Request("POST", discord.EndpointGuildEmojis(guildID), params)
	if err != nil {
		return nil, err
	}

	var emoji discord.Emoji
	err = json.Unmarshal(body, &emoji)

	return &emoji, err
}

// https://discord.com/developers/docs/resources/emoji#modify-guild-emoji
func (c *Client) ModifyGuildEmoji(guildID string, emojiID string, params *discord.ModifyGuildEmoji) (*discord.Emoji, error) {
	body, err := c.Request("PATCH", discord.EndpointGuildEmoji(guildID, emojiID), params)
	if err != nil {
		return nil, err
	}

	var emoji discord.Emoji
	err = json.Unmarshal(body, &emoji)

	return &emoji, err
}

// https://discord.com/developers/docs/resources/emoji#delete-guild-emoji
func (c *Client) DeleteGuildEmoji(guildID string, emojiID string) error {
	_, err := c.Request("DELETE", discord.EndpointGuildEmoji(guildID, emojiID), nil)
	return err
}
