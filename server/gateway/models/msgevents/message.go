package msgevents

import "time"

// Message defines struct message
type Message struct {
	ID        int64
	ChannelID int64
	Body      string
	CreatedAt *time.Time
	Creator   int64
	EditedAt  *time.Time
}
