package connection

import "time"

type Connection struct {
	ID          uint64 `json:"id"`
	ServiceName string
	Endpoint    string
	CreatedAt   *time.Time `gorm:"type:timestamp;" json:"created_at"`
}
