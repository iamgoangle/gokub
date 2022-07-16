package logger

// Error logs a message at ErrorLevel.
func Error(m string) {
	logger.wrapLogger.Error(m)
}

// Errorf log a message with format.
func Errorf(m string, v interface{}) {
	logger.wrapLogger.Sugar().Errorf(m, v)
}
