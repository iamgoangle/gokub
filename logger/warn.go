package logger

// Warn logs a message at WarnLevel.
func Warn(m string) {
	logger.wrapLogger.Warn(m)
}

// Warnf log a message with format.
func Warnf(m string, v interface{}) {
	logger.wrapLogger.Sugar().Warnf(m, v)
}
