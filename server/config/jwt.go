package config

type JWT struct {
	Secret           string
	MinSecretLength  int
	DurationMinutes  int
	SigningAlgorithm string
}
