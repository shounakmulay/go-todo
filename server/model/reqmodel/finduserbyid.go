package reqmodel

type UserID struct {
	ID int `param:"id" validate:"required"`
}
