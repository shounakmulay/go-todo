package error

import "go-todo/internal/util/log"

func Log(e error) {
	if e != nil {
		log.Logger.Error(e)
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
