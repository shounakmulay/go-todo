package reqmodel

type TodoStatus struct {
	Done int `query:"done" validate:"oneof=0 1"`
}
