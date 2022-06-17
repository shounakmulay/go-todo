package log

type gormLogger struct {
	logger *logger
}

func NewGormLogger() gormLogger {
	return gormLogger{
		logger: Logger,
	}
}

func (gl gormLogger) Printf(template string, args ...interface{}) {
	gl.logger.Named("GORM: ").Debugf(template, args)
}
