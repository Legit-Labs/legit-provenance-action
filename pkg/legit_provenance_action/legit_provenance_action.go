package legit_provenance_action

import (
	"os"
	"strings"
)

func GetEnv() map[string]string {
	envStrings := os.Environ()
	env := make(map[string]string, len(envStrings))
	for _, kv := range envStrings {
		pair := strings.SplitN(kv, "=", 2)
		key := pair[0]
		value := pair[1]
		env[key] = value
	}
	return env
}
