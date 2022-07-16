package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zapcore.DebugLevel

	// InfoLevel is the default logging priority.
	InfoLevel = zapcore.InfoLevel

	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zapcore.WarnLevel

	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zapcore.ErrorLevel

	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zapcore.DPanicLevel

	// PanicLevel logs a message, then panics.
	PanicLevel = zapcore.PanicLevel

	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zapcore.FatalLevel
)

func initZapLogger(options ...func(*zap.Logger) error) *zap.Logger {
	// initialize uber/zap logger
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder

	// entry fields
	fields := zap.Fields(
		zap.String("env", os.Getenv("APP_ENV")),
		zap.String("app_name", os.Getenv("APP_NAME")),
	)

	z, err := loggerConfig.Build(fields)
	if err != nil {
		panic("[init]: unable to initialize logger")
	}
	defer z.Sync() // flushes buffer

	for _, op := range options {
		err := op(z)
		if err != nil {
			panic(err)
		}
	}

	return z
}
