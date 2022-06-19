package reqmodel

type FindUserByID struct {
	ID int `param:"id" validate:"required"`
}
