package logger

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
	ALPHA Environment = "alpha"

	// BETA represent alpha environment
	BETA Environment = "beta"

	// STAGING represent alpha environment
	STAGING Environment = "staging"

	// PRODUCTION represent alpha environment
	PRODUCTION Environment = "production"
)

// Logger represent a logging wrapper type
// the library does not need to coupling with vendor locking so this library is going to expose
// the standard receiver method to logging with structuring logger manner.
type Logger struct {
	wrapLogger *zap.Logger
}

// LoggerOpts represents logger options
type LoggerOpts struct {
	Environment string
}

type Fields []Field
type Field struct {
	Key string
	Val interface{}
}

func init() {
	logger = Logger{
		wrapLogger: initZapLogger(
			setLogLevelByEnv(PRODUCTION),
		),
	}
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
