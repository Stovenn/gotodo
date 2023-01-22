package domain

type Todo struct {
	ID        string
	Title     string
	Completed bool
	Order     int
	Url       string
}

type TodoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	Url       string `json:"url"`
}

type TodoCreationRequest struct {
	Title string `json:"title"`
}

type TodoUpdateRequest struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
	Order     int    `json:"order" validate:"required"`
}

type TodoPartialUpdateRequest struct {
	Title     string `json:"title omitempty"`
	Completed bool   `json:"completed omitempty"`
	Order     int    `json:"order omitempty"`
}

func (t Todo) ToResponse() *TodoResponse {
	return &TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		Order:     t.Order,
		Url:       t.Url,
	}
}
