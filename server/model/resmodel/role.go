package resmodel

type Role struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AccessLevel int    `json:"accessLevel"`
}
