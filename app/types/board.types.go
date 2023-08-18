package types

type (
	BoardDTO struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
	}
)
