package logger

// Panic logs a message at PanicLevel.
func Panic(m string) {
	logger.wrapLogger.Panic(m)
}

// Panicf log a message with format.
func Panicf(m string, v interface{}) {
	logger.wrapLogger.Sugar().Panicf(m, v)
}
