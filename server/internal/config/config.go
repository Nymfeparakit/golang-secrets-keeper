package config

import (
	"flag"
)

// Config - stores server general settings.
type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:":8080"`
	DatabaseDSN   string `env:"DATABASE_DSN"`
	EnableHTTPS   bool   `env:"ENABLE_HTTPS" envDefault:"false" json:"enable_https"`
	CertFile      string `env:"CERTIFICATE_FILE" envDefault:"cert.pem"`
	TLSKey        string `env:"TLS_KEY" envDefault:"key.pem"`
}

// InitFlags - inits flags for server general settings.
func (c *Config) InitFlags() {
	flag.StringVar(
		&c.ServerAddress, "a", c.ServerAddress, "The address where the server will be started",
	)
	flag.StringVar(&c.DatabaseDSN, "d", c.DatabaseDSN, "Connection string for database storage")
	flag.BoolVar(&c.EnableHTTPS, "s", c.EnableHTTPS, "Should https be used")
	flag.StringVar(&c.CertFile, "c", c.CertFile, "Certificate file for tls connection")
	flag.StringVar(&c.TLSKey, "k", c.CertFile, "Private key file for tls connection")
}
