package types

type (
	ListDTO struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
	}
)
