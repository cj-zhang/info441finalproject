package msgevents

import "time"

// Channel defines a channel
type Channel struct {
	ID          int64
	Name        string
	Description string
	Private     bool
	Members     []int64
	CreatedAt   *time.Time
	Creator     int64
	editedAt    *time.Time
}
