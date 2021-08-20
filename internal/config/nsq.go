package config

import ()

// Nsq Nsq配置参数
type Nsq struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}
