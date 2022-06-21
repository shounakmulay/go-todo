package env

import (
	"fmt"
	"os"

	"go-todo/internal/convert"
	errorutl "go-todo/internal/error"
	"go-todo/internal/log"

	"github.com/joho/godotenv"
)

func GetString(key string) string {
	value := os.Getenv(key)
	if value == "" {
		keyNotFound(key)
	}

	return value
}

func GetInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		keyNotFound(key)
		return 0
	}
	return convert.StringToInt(value)
}

func GetBool(key string) bool {
	value := os.Getenv(key)
	if value == "" {
		keyNotFound(key)
		return false
	}
	return convert.StringToBool(value)
}

func keyNotFound(key string) {
	log.Logger.Infof("Key %s not found in %s Returning default value.", key, FileName())
}

func FileName() string {
	environment := os.Getenv("ENVIRONMENT_NAME")
	var envFileName string

	if environment == "" {
		envFileName = ".env"
	} else {
		envFileName = fmt.Sprintf(".env.%s", environment)
	}
	return envFileName
}

func Load() {
	const envName = "ENVIRONMENT_NAME"
	env := os.Getenv(envName)

	if env == "" {
		env = "local"
		errorutl.Log(os.Setenv(envName, env))
	}

	errorutl.Fatal(
		godotenv.Load(fmt.Sprintf(".env.%s", env)),
	)
}
