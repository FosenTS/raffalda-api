package config

import (
	"path"
	"raffalda-api/pkg/mysync"
)

const httpConfigFilename = "http.config.yaml"

type HTTPConfig struct {
	Host                    string
	Port                    string
	UseCache                bool `yaml:"useCache"`
	MaxConcurrentConnection uint `yaml:"maxConcurrentConnection" env-required:"true`
}

var (
	httpConfigInst     = &HTTPConfig{}
	loadHTTPConfigOnce = mysync.NewOnce()
)

func HTTP() HTTPConfig {
	loadHTTPConfigOnce.Do(func() {
		env := Env()

		httpConfigInst.Host = env.IpAddress
		httpConfigInst.Port = env.ApiPort
		readConfig(path.Join(env.ConfigAbsPath, httpConfigFilename), httpConfigInst)
	})

	return *httpConfigInst
}
