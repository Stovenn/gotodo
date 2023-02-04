package domain

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// Todo is the representation of a item
type Todo struct {
	ID         string         `db:"id"`
	Title      string         `db:"title"`
	Completed  bool           `db:"completed"`
	Order      int            `db:"item_order"`
	AssignedTo sql.NullString `db:"assigned_to"`
	CreatedAt  time.Time      `db:"created_at"`
	UpdatedAt  time.Time      `db:"updated_at"`
}

// TodoResponse is the struct
type TodoResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	Order      int    `json:"order"`
	AssignedTo string `json:"assigned_to"`
	URL        string `json:"url"`
}

func (tr TodoResponse) JSON() []byte {
	var b []byte
	var err error

	b, err = json.Marshal(&tr)
	if err != nil {
		return nil
	}
	return b
}

// TodoCreationRequest is the body of a todo creation
type TodoCreationRequest struct {
	Title      string `json:"title" validate:"required"`
	AssignedTo string `json:"assigned_to"`
}

// TodoUpdateRequest is the body of a todo update
type TodoUpdateRequest struct {
	Title      string `json:"title" validate:"required"`
	Completed  bool   `json:"completed" validate:"required"`
	Order      int    `json:"order" validate:"required"`
	AssignedTo string `json:"assigned_to" validate:"required"`
}

// TodoPartialUpdateRequest is the body of a partial todo update
// All arguments are optional
type TodoPartialUpdateRequest struct {
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	Order      int    `json:"order"`
	AssignedTo string `json:"assigned_to"`
}

// ToResponse maps a Todo to a TodoResponse
func (t Todo) ToResponse() *TodoResponse {
	return &TodoResponse{
		ID:         t.ID,
		Title:      t.Title,
		Completed:  t.Completed,
		Order:      t.Order,
		AssignedTo: t.AssignedTo.String,
		URL:        fmt.Sprintf("http://%s:%s/api/todos/%s", viper.Get("HOST"), viper.Get("PORT"), t.ID),
	}
}
