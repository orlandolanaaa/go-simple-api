package auth

import "database/sql"

type (
	UserToken struct {
		ID        int64        `json:"id"  db:"id"`
		UserID    int64        `json:"user_id"  db:"user_id"`
		Token     string       `json:"token"  db:"token"`
		ExpiredAt string       `json:"expired_at" db:"expired_at"`
		CreatedAt sql.NullTime `json:"created_at" db:"created_at"`
		UpdatedAt sql.NullTime `json:"updated_at" db:"updated_at"`
		DeletedAt sql.NullTime `json:"deleted_at" db:"deleted_at"`
	}
)
