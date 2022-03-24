package entities

import (
	"time"
)

type (
	UserToken struct {
		ID        int64      `json:"id"  db:"id"`
		UserID    int64      `json:"user_id"  db:"user_id"`
		Token     string     `json:"token"  db:"token"`
		ExpiredAt *time.Time `json:"expired_at" db:"expired_at"`
		CreatedAt *time.Time `json:"created_at" db:"created_at"`
		UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	}
)
