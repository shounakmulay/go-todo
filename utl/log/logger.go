package log

import "go.uber.org/zap"

var Logger *logger

type logger struct {
	*zap.SugaredLogger
}

func InitLogger() {
	if Logger != nil {
		return
	}

	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		// TODO: Handle logger creation error
		panic(err)
	}

	Logger = &logger{
		zapLogger.Sugar(),
	}
}
