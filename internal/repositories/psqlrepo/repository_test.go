package psqlrepo

//var r *todoRepository
//
//func init() {
//	r = NewTodoRepository()
//}
//
//func TestTodoRepository_Create(t *testing.T) {
//	createTodo(t)
//}
//
//func createTodo(t *testing.T) *domain.Todo {
//	arg := domain.Todo{ID: "", Title: util.RandomString(15), Order: 0, Completed: false, URL: ""}
//	createdTodo, err := r.Create(arg)
//	expected := &domain.Todo{ID: "", Title: "new todo", Order: , Completed: false, URL: ""}
//
//	assertCreation(t, expected, createdTodo, err)
//
//	return createdTodo
//}
//
//func assertCreation(t *testing.T, expected, got *domain.Todo, err error) {
//	assert.NotEmpty(t, got)
//	assert.NoError(t, err)
//
//	assert.Equal(t, expected.Title, got.Title)
//	assert.Equal(t, expected.Order, got.Order)
//	assert.NotZero(t, got.ID)
//}
