package config

// Case 全局Setup config
var (
	Case Configs
)

// Configs Configs
type Configs struct {
	Application Application `mapstructure:"application" json:"application" yaml:"application"`

	Nsq Nsq `mapstructure:"nsq" json:"nsq" yaml:"nsq"`

	Redis Redis `mapstructure:"redis" json:"redis" yaml:"redis"`
}
