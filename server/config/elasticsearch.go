package config

import "fmt"

type Elasticsearch struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}

func (m *Elasticsearch) Dsn() string {
	return fmt.Sprintf("http://%s:%d", m.Path, m.Port)
}
