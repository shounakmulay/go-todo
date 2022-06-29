package config

type Redis struct {
	URL                          string
	Port                         string
	AuthRateLimitCount           int
	AuthRateLimitDurationSeconds int
	DisableRateLimit             bool
}
