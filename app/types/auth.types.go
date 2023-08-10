package types

type (
	LoginDTO struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"password"`
	}

	SignupDTO struct {
		LoginDTO
		Name string `json:"name" validate:"required,min=3"`
	}
)
