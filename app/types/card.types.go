package types

type (
	CardDTO struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
		AssignedTo  uint   `json:"assignedTo"`
		DueDate     string `json:"dueDate"`
	}

	AssignCardToUserDTO struct {
		UserId uint `json:"userId" validate:"required"`
	}
)
