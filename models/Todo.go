package models

////////////////////////////////////////////
// Default todos
////////////////////////////////////////////
var defaultTodos = map[int]string{
	1: "Task one",
	2: "Task two",
}

////////////////////////////////////////////
// Todo struct
////////////////////////////////////////////
type Todo struct {
	todos map[int]string
}

////////////////////////////////////////////
// Create a new todo struct instance
////////////////////////////////////////////
func NewTodo() *Todo {
	return &Todo{
		todos: defaultTodos,
	}
}

// Get all todos
func (t *Todo) Get() map[int]string {
	return t.todos
}

// Find a todo by the given todo id
func (t *Todo) Find(id int) string {
	return t.todos[id]
}

// Insert a todo into todos list
func (t *Todo) Insert(todo string) bool {
	t.todos[len(t.todos)+1] = todo

	return true
}

// Update an existing todo
func (t *Todo) Update(id int, todo string) bool {
	t.todos[id] = todo

	return true
}

// Delete an existing todo by the given todo id
func (t *Todo) Delete(id int) bool {
	delete(t.todos, id)

	return true
}

// Get the total task in the todos
func (t *Todo) Count() int {
	return len(t.todos)
}
