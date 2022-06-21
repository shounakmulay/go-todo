package config

type Database struct {
	URL           string
	LogQueries    bool
	Timeout       int
	SlowThreshold int
}
