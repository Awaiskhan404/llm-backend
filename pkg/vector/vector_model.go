package vector

import "time"

type Vector struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	Description string `json:"description"`
	ConnectionString string `json:"connection_string"`
	CreatedAt time.Time `json:"created_at"`
}
