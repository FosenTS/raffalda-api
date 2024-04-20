package config

import (
	"path"
	"raffalda-api/pkg/gconfig"
	"raffalda-api/pkg/mysync"

	log "github.com/sirupsen/logrus"
)

const (
	envLocalFilename  = ".env.local"
	envDeployFilename = ".env.deploy"
)

// EnvConfig Env
type EnvConfig struct {
	ProjectAbsPath string `env:"PROJECT_ABS_PATH" env-required:"true"`
	ConfigPath     string `env:"CONFIG_PATH" env-required:"true"`

	IpAddress  string `env:"IP_ADDRESS" env-required:"true"`
	ApiPort    string `env:"API_PORT" env-required:"true"`
	MetricPort string `env:"METRIC_PORT" env-required:"true"`

	// Postgres
	PostgresHost         string `env:"POSTGRES_HOST" env-required:"true"`
	PostgresPort         string `env:"POSTGRES_PORT" env-required:"true"`
	PostgresUsername     string `env:"POSTGRES_USERNAME" env-required:"true"`
	PostgresPassword     string `env:"POSTGRES_PASSWORD" env-required:"true"`
	PostgresDatabaseName string `env:"POSTGRES_DATABASE_NAME" env-required:"true"`
	PostgresUseCA        bool   `env:"POSTGRES_USE_CA" env-required:"true"`
	PostgresCaPath       string `env:"POSTGRES_CA_PATH" env_required:"true"`
	PostgresTimeZone     string `env:"POSTGRES_TIME_ZONE" env-required:"true"`

	// Secrets
	PasswordHashSalt string `env:"PASSWORD_HASH_SALT" env-required:"true"`
	JWTSecret        string `env:"JWT_SECRET" env-required:"true"`

	ConfigAbsPath string
}

var (
	envCfgInst  = &EnvConfig{}
	loadEnvOnce = mysync.NewOnce()
)

func LoadEnv(mode string) {
	envFilename := ""
	switch mode {
	case Mode.Local():
		envFilename = envLocalFilename
	case Mode.Deploy():
		envFilename = envDeployFilename

	default:
		panic("invalid mode")
	}
	loadEnvOnce.Do(func() {
		err := gconfig.ReadEnv(envFilename, envCfgInst)
		if err != nil {
			log.Fatalf("fatal reading env: %s\n", err)
		}

		envCfgInst.ConfigAbsPath = path.Join(envCfgInst.ProjectAbsPath, envCfgInst.ConfigPath)

		log.Infoln("Env successfully read")
	})
}

func Env() EnvConfig {
	if !loadEnvOnce.IsDone() {
		log.Fatalln("env must be loaded before use!")
	}
	return *envCfgInst
}
