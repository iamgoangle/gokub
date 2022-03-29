package required

import (
	"log"
	"os"
)

// GetEnv get the environment variable
//
// panic the main application immediately if the env value is empty
func GetEnv(env string) string {
	e := os.Getenv(env)
	if len(e) == 0 {
		log.Panicf("[GetEnv]: unable to get value from %s", env)
	}

	return e
}
