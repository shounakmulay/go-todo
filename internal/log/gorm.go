package log

type GormLogger struct {
	logger *logger
}

func NewGormLogger() GormLogger {
	return GormLogger{
		logger: Logger,
	}
}

func (gl GormLogger) Printf(template string, args ...interface{}) {
	gl.logger.Named("GORM: ").Debugf(template, args)
}
