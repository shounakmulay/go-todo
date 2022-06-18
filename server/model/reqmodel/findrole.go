package reqmodel

type FindRole struct {
	ID int `param:"id" validate:"required"`
}
