package config

type Server struct {
	Port                string
	Debug               bool
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	SkipLogs            bool
	SkipBodyDump        bool
}
