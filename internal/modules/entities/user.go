package entities

import (
	"time"
)

type User struct {
	ID             int64      `json:"id"  db:"id"`
	Username       string     `json:"username"  db:"username"`
	Email          string     `json:"email"  db:"email"`
	Password       string     `json:"password" db:"password"`
	Nickname       *string    `json:"nickname" db:"nickname"`
	ProfilePicture *string    `json:"profile_picture" db:"profile_picture"`
	CreatedAt      *time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}
