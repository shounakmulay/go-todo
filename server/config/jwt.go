package config

type JWT struct {
	Secret                 string
	RefreshSecret          string
	MinSecretLength        int
	DurationMinutes        int
	RefreshDurationMinutes int
	SigningAlgorithm       string
}
