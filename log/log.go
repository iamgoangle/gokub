package log

import (
	"go.uber.org/zap"
)

var (
	logger Logger

	environment string
)

// Environment represent environment type
type Environment string

const (
	// ALPHA represent alpha environment
	ALPHA      Environment = "alpha"

	// BETA represent alpha environment
	BETA       Environment = "beta"

	// STAGING represent alpha environment
	STAGING    Environment = "staging"

	// PRODUCTION represent alpha environment
	PRODUCTION Environment = "production"
)

// Logger represent a logging wrapper type
// the library does not need to coupling with vendor locking so this library is going to expose
// the standard receiver method to logging with structuring log manner.
type Logger struct {
	wrapLogger *zap.Logger
}

type LoggerOpts struct {
	Environment string
}

func (l *Logger) Info(m string) {
	l.wrapLogger.Info(m)
}

func initZapLogger(options ...func(*zap.Logger) error) *zap.Logger {
	// initialize uber/zap logger
	z, err := zap.NewProduction()
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

// setLogLevelByEnv automatic set logger level by environment and above logging
func setLogLevelByEnv(env Environment) func(*zap.Logger) error {
	return func(z *zap.Logger) error {
		atom := zap.NewAtomicLevel()
		if env == PRODUCTION {
			atom.SetLevel(ErrorLevel)
		}

		atom.SetLevel(DebugLevel)

		return nil
	}
}

func init() {
	logger = Logger{
		wrapLogger: initZapLogger(
			setLogLevelByEnv(PRODUCTION),
		),
	}
}

func Info(m string) {
	logger.wrapLogger.Info(m)
}

func Debug(m string) {
	logger.wrapLogger.Debug(m)
}