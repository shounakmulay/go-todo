package reqmodel

type FindUserByUsername struct {
	Username string `param:"username" validate:"required"`
}
