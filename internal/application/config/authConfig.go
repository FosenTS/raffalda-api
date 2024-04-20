package config

import (
	"path"
	"raffalda-api/pkg/mysync"
	"time"
)

const authCfgFilename = "auth.config.yaml"

type AuthConfig struct {
	Salt         string
	SecretJWTKey string

	jwtLiveTimeSeconds     uint `yaml:"jwtLiveTimeSeconds" env-required:"true"`
	refreshLiveTimeSeconds uint `yaml:"refreshLiveTimeSeconds" env-required:"true"`

	JwtLiveTime     time.Duration
	RefreshLiveTime time.Duration

	FastAuth struct {
		Scheme string `yaml:"scheme" env-required:"true"`
		Url    string `yaml:"url" env-required:"true"`
		Path   string `yaml:"path" env-required:"true"`
	} `yaml:"fastAuth" env-required:"true"`
}

var (
	authConfigInst     = &AuthConfig{}
	loadAuthConfigOnce = mysync.NewOnce()
)

func Auth() AuthConfig {
	loadAuthConfigOnce.Do(func() {
		env := Env()
		readConfig(path.Join(env.ConfigAbsPath, authCfgFilename), authConfigInst)

		// TODO: fix token expires time
		authConfigInst.JwtLiveTime = time.Second * time.Duration(authConfigInst.jwtLiveTimeSeconds)

		authConfigInst.JwtLiveTime = time.Second * 10000

		authConfigInst.RefreshLiveTime = time.Duration(authConfigInst.refreshLiveTimeSeconds) * time.Second

		authConfigInst.RefreshLiveTime = time.Hour * 100
	})

	return *authConfigInst
}
