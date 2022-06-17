package error

import (
	"errors"
	"fmt"
	"go-todo/internal/log"
)

func Log(e error) {
	if e != nil {
		log.Logger.Error(e.Error())
	}
}

func Panic(e error) {
	if e != nil {
		log.Logger.Panic(e)
	}
}

func Fatal(e error) {
	if e != nil {
		log.Logger.Fatal(e)
	}
}

func Format(msg string, e error) error {
	errMsg := fmt.Sprintf("%s: %v", msg, e)
	return errors.New(errMsg)
}
