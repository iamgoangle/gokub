package log

import (
	"go.uber.org/zap"
)

var (
	logger Logger
)

// Logger represent a logging wrapper type
// the library does not need to coupling with vendor locking so this library is going to expose
// the standard receiver method to logging with structuring log manner.
type Logger struct {
	wrapLogger *zap.Logger
}

func (l *Logger) Info(m string) {
	l.wrapLogger.Info(m)
}

func init() {
	// initialize uber/zap logger
	z, err := zap.NewProduction()
	if err != nil {
		panic("[init]: unable to initialize logger")
	}

	logger = Logger{wrapLogger: z}
}

func Info(m string){
	logger.wrapLogger.Info(m)
}