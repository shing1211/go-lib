package config

// ServerConfig
type ServerConfig struct {
	ServerPort     string // public port of server
	ServerMode     string
	ServerLogLevel uint64
	ServerJWT      struct {
		Key    string
		Expire int
	}
	ServerHashPass struct {
		Memory      uint32
		Iterations  uint32
		Parallelism uint8
		SaltLength  uint32
		KeyLength   uint32
	}
}
