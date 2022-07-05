package env

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

	if env == "develop" {
		type copilotSecrets struct {
			Username string `json:"username"`
			Host     string `json:"host"`
			DBName   string `json:"dbname"`
			Password string `json:"password"`
			Port     int    `json:"port"`
		}
		secrets := &copilotSecrets{}
		errorutl.Fatal(json.Unmarshal([]byte(os.Getenv("SHGOTODOSERVICECLUSTER_SECRET")), secrets))

		os.Setenv(
			"DB_SQL_URL",
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				secrets.Username,
				secrets.Password,
				secrets.Host,
				secrets.Port,
				secrets.DBName,
			))
		os.Setenv("DB_NAME", secrets.DBName)
		os.Setenv("DB_PORT", strconv.Itoa(secrets.Port))
		os.Setenv("DB_HOST", secrets.Host)
		os.Setenv("DB_USER", secrets.Username)
		os.Setenv("DB_PASS", secrets.Password)

		redisURL := os.Getenv("SH_GO_TODO_REDIS_ENDPOINT")
		redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
		if err != nil {
			redisPort = 6379
		}
		os.Setenv("REDIS_URL", fmt.Sprintf("%s:%d", redisURL, redisPort))
	}
}
