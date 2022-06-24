package config

type Configuration struct {
	Server   *Server
	Database *Database
	JWT      *JWT
	Redis    *Redis
}
