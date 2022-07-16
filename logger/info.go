package logger

// Info logs a message at InfoLevel.
func Info(m string) {
	logger.wrapLogger.Info(m)
}

// Infof log a message with format.
func Infof(m string, v interface{}) {
	logger.wrapLogger.Sugar().Infof(m, v)
}

// InfoWithFields log a message with fields
func InfoWithFields(m string, fs Fields) {
	var fields []interface{}

	for _, field := range fs {
		fields = append(fields, field.Key, field.Val)
	}

	logger.wrapLogger.Sugar().Infow(m, fields...)
}
