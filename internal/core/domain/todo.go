package domain

// Todo is the representation of a item
type Todo struct {
	ID        string
	Title     string
	Completed bool
	Order     int
	URL       string
}

// TodoResponse is the struct
type TodoResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
	URL       string `json:"url"`
}

// TodoCreationRequest is the body of a todo creation
type TodoCreationRequest struct {
	Title string `json:"title"`
}

// TodoUpdateRequest is the body of a todo update
type TodoUpdateRequest struct {
	Title     string `json:"title" validate:"required"`
	Completed bool   `json:"completed" validate:"required"`
	Order     int    `json:"order" validate:"required"`
}

// TodoPartialUpdateRequest is the body of a partial todo update
// All arguments are optional
type TodoPartialUpdateRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
}

// ToResponse maps a Todo to a TodoResponse
func (t Todo) ToResponse() *TodoResponse {
	return &TodoResponse{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		Order:     t.Order,
		URL:       t.URL,
	}
}
