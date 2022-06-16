package error

import (
	"go-todo/utl/log"
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
