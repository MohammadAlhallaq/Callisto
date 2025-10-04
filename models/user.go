package models

import "time"

type User struct {
	ID        int64     `json:"id,omitempty"`
	Email     string     `json:"email"`
	Password  string     `json:"password,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
