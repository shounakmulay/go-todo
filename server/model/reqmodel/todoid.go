package reqmodel

type TodoID struct {
	ID int `param:"id" validate:"required"`
}
