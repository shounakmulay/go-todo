package config

import (
	"fmt"
	"os"

	"go-todo/internal/env"
)

const JwtContextKey string = "USER"

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
			URL:           env.GetString("DB_SQL_URL"),
			LogQueries:    env.GetBool("DB_LOG_QUERIES"),
			Timeout:       env.GetInt("DB_TIMEOUT_SECONDS"),
			SlowThreshold: env.GetInt("DB_SLOW_THRESHOLD_MS"),
		},
		JWT: &JWT{
			Secret:                 env.GetString("JWT_SECRET"),
			RefreshSecret:          env.GetString("JWT_REFRESH_SECRET"),
			MinSecretLength:        env.GetInt("JWT_MIN_SECRET_LENGTH"),
			DurationMinutes:        env.GetInt("JWT_DURATION_MINS"),
			RefreshDurationMinutes: env.GetInt("JWT_REFRESH_MINS"),
			SigningAlgorithm:       env.GetString("JWT_ALGO"),
			ContextKey:             JwtContextKey,
		},
		Redis: &Redis{
			URL:  env.GetString("REDIS_URL"),
			Port: env.GetString("REDIS_PORT"),
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
