package logger

// Debug logs a message at DebugLevel.
func Debug(m string) {
	logger.wrapLogger.Debug(m)
}

// Debugf log a message with format.
func Debugf(m string, v interface{}) {
	logger.wrapLogger.Sugar().Debugf(m, v)
}
