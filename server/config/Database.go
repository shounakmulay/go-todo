package config

type Database struct {
	Url           string
	LogQueries    bool
	Timeout       int
	SlowThreshold int
}
