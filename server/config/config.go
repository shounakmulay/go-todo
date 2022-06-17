package config

import (
	"fmt"
	"go-todo/internal/env"
	"os"
)

func Load() (*Configuration, error) {
	envLoadErr := checkEnvKeysPresent(
		"DB_TIMEOUT_SECONDS",
		"DB_SLOW_THRESHOLD_MS",
		"SERVER_PORT",
		"SERVER_READ_TIMEOUT",
		"SERVER_WRITE_TIMEOUT",
	)
	if envLoadErr != nil {
		return nil, envLoadErr
	}

	cfg := &Configuration{
		Server: &Server{
			Port:                fmt.Sprintf(":%d", env.GetInt("SERVER_PORT")),
			Debug:               env.GetBool("SERVER_DEBUG"),
			ReadTimeoutSeconds:  env.GetInt("SERVER_READ_TIMEOUT"),
			WriteTimeoutSeconds: env.GetInt("SERVER_WRITE_TIMEOUT"),
			SkipLogs:            env.GetBool("SERVER_SKIP_LOGS"),
			SkipBodyDump:        env.GetBool("SERVER_SKIP_BODY_DUMP"),
		},
		Database: &Database{
			Url:           env.GetString("DB_SQL_URL"),
			LogQueries:    env.GetBool("DB_LOG_QUERIES"),
			Timeout:       env.GetInt("DB_TIMEOUT_SECONDS"),
			SlowThreshold: env.GetInt("DB_SLOW_THRESHOLD_MS"),
		},
	}

	return cfg, nil
}

func checkEnvKeysPresent(keys ...string) error {
	for _, key := range keys {
		value := os.Getenv(key)
		if value == "" {
			envFile := env.FileName()
			return fmt.Errorf(
				"error loading REQUIRED KEY %s from %s file. "+
					"Make sure you have declared it in %s", key, envFile, envFile,
			)
		}
	}
	return nil
}
