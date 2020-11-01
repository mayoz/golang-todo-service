package todo

import "time"

type Todo struct {
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
