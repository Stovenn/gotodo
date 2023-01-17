package domain

type Todo struct {
	id        string
	title     string
	completed bool
	order     int
	url       string
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
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Order     int    `json:"order"`
}

type TodoPartialUpdateRequest struct {
	Title     string `json:"title omitempty"`
	Completed bool   `json:"completed omitempty"`
	Order     int    `json:"order omitempty"`
}

func (t Todo) toResponse() TodoResponse {
	return TodoResponse{
		ID:        t.id,
		Title:     t.title,
		Completed: t.completed,
		Order:     t.order,
		Url:       t.url,
	}
}
