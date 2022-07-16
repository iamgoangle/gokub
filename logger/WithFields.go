package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func WithFields(fs ...interface{}) *Logger {
	var fields []zap.Field
	// TODO: loop assign
	fields = append(fields, zap.Field{Key: "key", Type: zapcore.StringType, String: "test tst"})

	log := &Logger{
		wrapLogger: logger.wrapLogger.With(fields...),
	}
	return log
}

func (l *Logger) Info(m string) {
	l.wrapLogger.Info(m)
}
