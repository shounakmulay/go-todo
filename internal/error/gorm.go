package error

import (
	"errors"
	"fmt"
	"go-todo/server/model/resmodel"
	"gorm.io/gorm"
)

func GormToResErr(err error, id any) *resmodel.ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resmodel.NotFound(fmt.Sprintf("No record found for id = %v", id))
	}
	return resmodel.InternalServerError(err)
}
