package error

import (
	"errors"
	"fmt"

	"go-todo/server/model/resmodel"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

func GormToResErr(err error, id any) *resmodel.ErrorResponse {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return resmodel.NotFound(fmt.Sprintf("No record found for %v", id))
	}

	queryError, ok := err.(*QueryError)
	if ok {
		return resmodel.BadRequest(queryError)
	}

	mysqlErr, ok := err.(*mysql.MySQLError)
	if ok {
		switch mysqlErr.Number {
		// Error codes that need to passed to the user, i.e. client errors.
		case 1062:
		case 1452:
			return resmodel.BadRequest(mysqlErr)
		}
	}

	return resmodel.InternalServerError(err)
}

type QueryError struct {
	message string
}

func (q *QueryError) Error() string {
	return q.message
}

func NewQueryError(msg string) *QueryError {
	return &QueryError{
		message: msg,
	}
}
