package config

import ()

// Application Application
type Application struct {
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	KeyExpiration int    `mapstructure:"key_expiration" json:"key_expiration" yaml:"key_expiration"`
}
