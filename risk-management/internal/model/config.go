package model

import "time"

type Config struct {
	ServiceName            string        `mapstructure:"service-name"`
	HostName               string        `mapstructure:"hostname"`
	Level                  string        `mapstructure:"level"`
	HttpPort               int           `mapstructure:"http-port"`
	HttpReadTimeout        time.Duration `mapstructure:"http-read-timeout"`
	HttpWriteTimetout      time.Duration `mapstructure:"http-write-timeout"`
	ServerShutdownTimeout  time.Duration `mapstructure:"server-shutdown-timeout"`
	StorageBackend         string        `mapstructure:"storage-backend"`
	DefaultPaginationLimit int           `mapstructure:"default-pagination-limit"`
}
