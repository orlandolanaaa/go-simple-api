package auth

type (
	RegisterUserRequest struct {
		Username string `json:"username" validate:"required,min=1,max=16" db:"username"`
		Email    string `json:"email" validate:"required,email" db:"email"`
		Password string `json:"password" validate:"required,min=1" db:"password"`
	}

	LoginRequest struct {
		Username string `json:"username" validate:"required,min=1,max=16" db:"username"`
		Password string `json:"password" validate:"required,min=8" db:"password"`
	}
)
