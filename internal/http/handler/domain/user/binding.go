package user

type (
	AuthMeta struct {
		ID             int64  `json:"id"  db:"id"`
		Username       string `json:"username"  db:"username"`
		Email          string `json:"email"  db:"email"`
		Password       string `json:"password" db:"password"`
		Nickname       string `json:"nickname" db:"nickname"`
		ProfilePicture string `json:"profile_picture" db:"profile_picture"`
		CreatedAt      string `json:"created_at" db:"created_at"`
		UpdatedAt      string `json:"updated_at" db:"updated_at"`
	}

	User struct {
		ID             int64  `json:"id"  db:"id"`
		Username       string `json:"username"  db:"username"`
		Email          string `json:"email"  db:"email"`
		Password       string `json:"password" db:"password"`
		Nickname       string `json:"nickname" db:"nickname"`
		ProfilePicture string `json:"profile_picture" db:"profile_picture"`
		CreatedAt      string `json:"created_at" db:"created_at"`
		UpdatedAt      string `json:"updated_at" db:"updated_at"`
	}

	Profile struct {
		Username       string `json:"username" `
		Email          string `json:"email"  `
		Nickname       string `json:"nickname" `
		ProfilePicture string `json:"profile_picture" `
	}
)
