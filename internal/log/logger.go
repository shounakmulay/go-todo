package log

import "go.uber.org/zap"

var Logger = initLogger()

type logger struct {
	*zap.SugaredLogger
	// error occurred during creation of logger
	InitError error
}

func initLogger() *logger {
	zapLogger, err := zap.NewDevelopment()

	return &logger{
		zapLogger.Sugar(),
		err,
	}
}
