package user

import "database/sql"

type User struct {
	ID             int64          `json:"id"  db:"id"`
	Username       string         `json:"username"  db:"username"`
	Email          string         `json:"email"  db:"email"`
	Password       string         `json:"password" db:"password"`
	Nickname       sql.NullString `json:"nickname" db:"nickname"`
	ProfilePicture sql.NullString `json:"profile_picture" db:"profile_picture"`
	CreatedAt      sql.NullTime   `json:"created_at" db:"created_at"`
	UpdatedAt      sql.NullTime   `json:"updated_at" db:"updated_at"`
	DeletedAt      sql.NullTime   `json:"deleted_at" db:"deleted_at"`
}
