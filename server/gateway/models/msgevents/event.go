package msgevents

// Event defines an event
type Event struct {
	Type      string   `json:"type"`
	Channel   *Channel `json:"channel"`
	ChannelID int64    `json:"channelID"`
	Message   *Message `json:"message"`
	MessageID int64    `json:"messageID"`
	UserIDs   []int64  `json:"userIDs"`
}
