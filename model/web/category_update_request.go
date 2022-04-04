package web

type CategoryUpdateRequest struct {
	ID   int64	`json:"id" validate:"required"`
	Name string	`json:"name" validate:"required, max=200, min=1"`
}
